package v1alpha5

import (
	"fmt"
	"net"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/util/validation"
	kubeletapis "k8s.io/kubernetes/pkg/kubelet/apis"
)

var (
	// ErrClusterEndpointNoAccess indicates the config prevents API access
	ErrClusterEndpointNoAccess = errors.New("Kubernetes API access must have one of public or private clusterEndpoints enabled")

	// ErrClusterEndpointPrivateOnly warns private-only access requires changes
	// to AWS resource configuration in order to effectively use clients in the VPC
	ErrClusterEndpointPrivateOnly = errors.New("warning, having public access disallowed will subsequently interfere with some " +
		"features of eksctl. This will require running subsequent eksctl (and Kubernetes) " +
		"commands/API calls from within the VPC.  Running these in the VPC requires making " +
		"updates to some AWS resources.  See: " +
		"https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html#private-access " +
		"for more details")
)

// NOTE: we don't use k8s.io/apimachinery/pkg/util/sets here to keep API package free of dependencies
type nameSet map[string]struct{}

func (s nameSet) checkUnique(path, name string) (bool, error) {
	if _, notUnique := s[name]; notUnique {
		return false, fmt.Errorf("%s %q is not unique", path, name)
	}
	s[name] = struct{}{}
	return true, nil
}

// ValidateClusterConfig checks compatible fields of a given ClusterConfig
func ValidateClusterConfig(cfg *ClusterConfig) error {
	if IsDisabled(cfg.IAM.WithOIDC) && len(cfg.IAM.ServiceAccounts) > 0 {
		return fmt.Errorf("iam.withOIDC must be enabled explicitly for iam.serviceAccounts to be created")
	}

	saNames := nameSet{}
	for i, sa := range cfg.IAM.ServiceAccounts {
		path := fmt.Sprintf("iam.serviceAccounts[%d]", i)
		if sa.Name == "" {
			return fmt.Errorf("%s.name must be set", path)
		}
		if ok, err := saNames.checkUnique("<namespace>/<name> of "+path, sa.NameString()); !ok {
			return err
		}
		if len(sa.AttachPolicyARNs) == 0 && sa.AttachPolicy == nil {
			return fmt.Errorf("%s.attachPolicyARNs or %s.attachPolicy must be set", path, path)
		}
	}

	// names must be unique across both managed and unmanaged nodegroups
	ngNames := nameSet{}
	validateNg := func(ng *NodeGroupBase, path string) error {
		if ng.Name == "" {
			return fmt.Errorf("%s.name must be set", path)
		}
		if _, err := ngNames.checkUnique(path+".name", ng.Name); err != nil {
			return err
		}
		if cfg.PrivateCluster.Enabled && !ng.PrivateNetworking {
			return fmt.Errorf("%s.privateNetworking must be enabled for a fully-private cluster", path)
		}
		return nil
	}

	for i, ng := range cfg.NodeGroups {
		path := fmt.Sprintf("nodeGroups[%d]", i)
		if err := validateNg(ng.NodeGroupBase, path); err != nil {
			return err
		}
	}

	for i, ng := range cfg.ManagedNodeGroups {
		path := fmt.Sprintf("managedNodeGroups[%d]", i)
		if err := validateNg(ng.NodeGroupBase, path); err != nil {
			return err
		}
	}

	if cfg.HasClusterCloudWatchLogging() {
		for i, logType := range cfg.CloudWatch.ClusterLogging.EnableTypes {
			isUnknown := true
			for _, knownLogType := range SupportedCloudWatchClusterLogTypes() {
				if logType == knownLogType {
					isUnknown = false
				}
			}
			if isUnknown {
				return fmt.Errorf("log type %q (cloudWatch.clusterLogging.enableTypes[%d]) is unknown", logType, i)
			}
		}
	}

	if cfg.VPC != nil && len(cfg.VPC.PublicAccessCIDRs) > 0 {
		cidrs, err := validateCIDRs(cfg.VPC.PublicAccessCIDRs)
		if err != nil {
			return err
		}
		cfg.VPC.PublicAccessCIDRs = cidrs
	}

	if cfg.SecretsEncryption != nil && cfg.SecretsEncryption.KeyARN == nil {
		return errors.New("field secretsEncryption.keyARN is required for enabling secrets encryption")
	}

	return nil
}

// ValidateClusterEndpointConfig checks the endpoint configuration for potential issues
func (c *ClusterConfig) ValidateClusterEndpointConfig() error {
	if !c.HasClusterEndpointAccess() {
		return ErrClusterEndpointNoAccess
	}
	endpts := c.VPC.ClusterEndpoints
	if NoAccess(endpts) {
		return ErrClusterEndpointNoAccess
	}
	return nil
}

// ValidatePrivateCluster validates the private cluster config
func (c *ClusterConfig) ValidatePrivateCluster() error {
	if c.PrivateCluster.Enabled {
		if c.VPC != nil && c.VPC.ID != "" && len(c.VPC.Subnets.Private) == 0 {
			return errors.New("vpc.subnets.private must be specified in a fully-private cluster when a pre-existing VPC is supplied")
		}
		if additionalEndpoints := c.PrivateCluster.AdditionalEndpointServices; len(additionalEndpoints) > 0 {
			if err := ValidateAdditionalEndpointServices(additionalEndpoints); err != nil {
				return errors.Wrap(err, "invalid value in privateCluster.additionalEndpointServices")
			}
		}
		if c.VPC != nil && c.VPC.ClusterEndpoints == nil {
			c.VPC.ClusterEndpoints = &ClusterEndpoints{}
		}
		// public access is initially enabled to allow running operations that access the Kubernetes API
		c.VPC.ClusterEndpoints.PublicAccess = Enabled()
		c.VPC.ClusterEndpoints.PrivateAccess = Enabled()
	}
	return nil
}

// NoAccess returns true if neither public are private cluster endpoint access is enabled and false otherwise
func NoAccess(ces *ClusterEndpoints) bool {
	return !(*ces.PublicAccess || *ces.PrivateAccess)
}

// PrivateOnly returns true if public cluster endpoint access is disabled and private cluster endpoint access is enabled, and false otherwise
func PrivateOnly(ces *ClusterEndpoints) bool {
	return !*ces.PublicAccess && *ces.PrivateAccess
}

// ValidateNodeGroup checks compatible fields of a given nodegroup
func ValidateNodeGroup(i int, ng *NodeGroup) error {
	path := fmt.Sprintf("nodeGroups[%d]", i)

	if ng.VolumeSize == nil {
		errCantSet := func(field string) error {
			return fmt.Errorf("%s.%s cannot be set without %s.volumeSize", path, field, path)
		}
		if IsSetAndNonEmptyString(ng.VolumeType) {
			return errCantSet("volumeType")
		}
		if IsSetAndNonEmptyString(ng.VolumeName) {
			return errCantSet("volumeName")
		}
		if IsEnabled(ng.VolumeEncrypted) {
			return errCantSet("volumeEncrypted")
		}
		if IsSetAndNonEmptyString(ng.VolumeKmsKeyID) {
			return errCantSet("volumeKmsKeyID")
		}
	}

	if ng.VolumeType != nil && *ng.VolumeType == NodeVolumeTypeIO1 {
		if ng.VolumeIOPS == nil {
			return fmt.Errorf("%s.volumeIOPS is required for %s volume type", path, NodeVolumeTypeIO1)
		}
	} else if ng.VolumeIOPS != nil {
		return fmt.Errorf("%s.volumeIOPS is only supported for %s volume type", path, NodeVolumeTypeIO1)
	}

	if ng.VolumeEncrypted == nil || IsDisabled(ng.VolumeEncrypted) {
		if IsSetAndNonEmptyString(ng.VolumeKmsKeyID) {
			return fmt.Errorf("%s.VolumeKmsKeyID can not be set without %s.VolumeEncrypted enabled explicitly", path, path)
		}
	}

	if ng.IAM != nil {
		if err := validateNodeGroupIAM(ng.IAM, ng.IAM.InstanceProfileARN, "instanceProfileARN", path); err != nil {
			return err
		}
		if err := validateNodeGroupIAM(ng.IAM, ng.IAM.InstanceRoleARN, "instanceRoleARN", path); err != nil {
			return err
		}
		if attachPolicyARNs := ng.IAM.AttachPolicyARNs; len(attachPolicyARNs) > 0 {
			for _, policyARN := range attachPolicyARNs {
				if _, err := arn.Parse(policyARN); err != nil {
					return errors.Wrapf(err, "invalid ARN %q in %s.iam.attachPolicyARNs", policyARN, path)
				}

			}
		}
	}

	if err := ValidateNodeGroupLabels(ng.Labels); err != nil {
		return err
	}

	if ng.SSH != nil {
		if err := validateNodeGroupSSH(ng.SSH); err != nil {
			return err
		}
		if len(ng.SSH.SourceSecurityGroupIDs) > 0 {
			return fmt.Errorf("%s.sourceSecurityGroupIds is not supported for unmanaged nodegroups", path)
		}
	}

	if ng.Bottlerocket != nil && ng.AMIFamily != NodeImageFamilyBottlerocket {
		return fmt.Errorf(`bottlerocket config can only be used with amiFamily "Bottlerocket" but found %s (path=%s.bottlerocket)`,
			ng.AMIFamily, path)
	}

	if IsWindowsImage(ng.AMIFamily) || ng.AMIFamily == NodeImageFamilyBottlerocket {
		fieldNotSupported := func(field string) error {
			return fmt.Errorf("%s is not supported for %s nodegroups (path=%s.%s)", field, ng.AMIFamily, path, field)
		}
		if ng.KubeletExtraConfig != nil {
			return fieldNotSupported("kubeletExtraConfig")
		}
		if ng.AMIFamily == NodeImageFamilyBottlerocket {
			if ng.PreBootstrapCommands != nil {
				return fieldNotSupported("preBootstrapCommands")
			}
		}
		if ng.OverrideBootstrapCommand != nil {
			return fieldNotSupported("overrideBootstrapCommand")
		}

	} else if err := validateNodeGroupKubeletExtraConfig(ng.KubeletExtraConfig); err != nil {
		return err
	}

	if ng.AMIFamily == NodeImageFamilyBottlerocket && ng.Bottlerocket != nil {
		err := checkBottlerocketSettings(ng.Bottlerocket.Settings, path)
		if err != nil {
			return err
		}
	}

	if err := validateInstancesDistribution(ng); err != nil {
		return err
	}

	return nil
}

// ValidateNodeGroupLabels uses proper Kubernetes label validation,
// it's designed to make sure users don't pass weird labels to the
// nodes, which would prevent kubelets to startup properly
func ValidateNodeGroupLabels(labels map[string]string) error {
	// compact version based on:
	// - https://github.com/kubernetes/kubernetes/blob/v1.13.2/cmd/kubelet/app/options/options.go#L257-L267
	// - https://github.com/kubernetes/kubernetes/blob/v1.13.2/pkg/kubelet/apis/well_known_labels.go
	// we cannot import those packages because they break other dependencies

	unknownKubernetesLabels := []string{}

	for label := range labels {
		labelParts := strings.Split(label, "/")

		if len(labelParts) > 2 {
			return fmt.Errorf("node label key %q is of invalid format, can only use one '/' separator", label)
		}

		if errs := validation.IsQualifiedName(label); len(errs) > 0 {
			return fmt.Errorf("label %q is invalid - %v", label, errs)
		}
		if errs := validation.IsValidLabelValue(labels[label]); len(errs) > 0 {
			return fmt.Errorf("label %q has invalid value %q - %v", label, labels[label], errs)
		}

		if len(labelParts) == 2 {
			namespace := labelParts[0]
			if isKubernetesLabel(namespace) && !kubeletapis.IsKubeletLabel(label) {
				unknownKubernetesLabels = append(unknownKubernetesLabels, label)
			}
		}
	}

	if len(unknownKubernetesLabels) > 0 {
		return fmt.Errorf("unknown 'kubernetes.io' or 'k8s.io' labels were specified: %v", unknownKubernetesLabels)
	}
	return nil
}

func isKubernetesLabel(namespace string) bool {
	for _, domain := range []string{"kubernetes.io", "k8s.io"} {
		if namespace == domain || strings.HasSuffix(namespace, "."+domain) {
			return true
		}
	}
	return false
}

func validateNodeGroupIAM(iam *NodeGroupIAM, value, fieldName, path string) error {
	if value != "" {
		fmtFieldConflictErr := func(conflictingField string) error {
			return fmt.Errorf("%s.iam.%s and %s.iam.%s cannot be set at the same time", path, fieldName, path, conflictingField)
		}

		if iam.InstanceRoleName != "" {
			return fmtFieldConflictErr("instanceRoleName")
		}
		if len(iam.AttachPolicyARNs) != 0 {
			return fmtFieldConflictErr("attachPolicyARNs")
		}
		if iam.InstanceRolePermissionsBoundary != "" {
			return fmtFieldConflictErr("instanceRolePermissionsBoundary")
		}
		prefix := "withAddonPolicies."
		if IsEnabled(iam.WithAddonPolicies.AutoScaler) {
			return fmtFieldConflictErr(prefix + "autoScaler")
		}
		if IsEnabled(iam.WithAddonPolicies.ExternalDNS) {
			return fmtFieldConflictErr(prefix + "externalDNS")
		}
		if IsEnabled(iam.WithAddonPolicies.CertManager) {
			return fmtFieldConflictErr(prefix + "certManager")
		}
		if IsEnabled(iam.WithAddonPolicies.ImageBuilder) {
			return fmtFieldConflictErr(prefix + "imageBuilder")
		}
		if IsEnabled(iam.WithAddonPolicies.AppMesh) {
			return fmtFieldConflictErr(prefix + "appMesh")
		}
		if IsEnabled(iam.WithAddonPolicies.AppMeshPreview) {
			return fmtFieldConflictErr(prefix + "appMeshPreview")
		}
		if IsEnabled(iam.WithAddonPolicies.EBS) {
			return fmtFieldConflictErr(prefix + "ebs")
		}
		if IsEnabled(iam.WithAddonPolicies.FSX) {
			return fmtFieldConflictErr(prefix + "fsx")
		}
		if IsEnabled(iam.WithAddonPolicies.EFS) {
			return fmtFieldConflictErr(prefix + "efs")
		}
		if IsEnabled(iam.WithAddonPolicies.ALBIngress) {
			return fmtFieldConflictErr(prefix + "albIngress")
		}
		if IsEnabled(iam.WithAddonPolicies.XRay) {
			return fmtFieldConflictErr(prefix + "xRay")
		}
		if IsEnabled(iam.WithAddonPolicies.CloudWatch) {
			return fmtFieldConflictErr(prefix + "cloudWatch")
		}
	}
	return nil
}

// ValidateManagedNodeGroup validates a ManagedNodeGroup and sets some defaults
func ValidateManagedNodeGroup(ng *ManagedNodeGroup, index int) error {
	if ng.AMIFamily != NodeImageFamilyAmazonLinux2 {
		return fmt.Errorf("only %s is supported for Managed Nodegroups", NodeImageFamilyAmazonLinux2)
	}
	path := fmt.Sprintf("managedNodeGroups[%d]", index)
	if ng.IAM != nil {
		if err := validateNodeGroupIAM(ng.IAM, ng.IAM.InstanceRoleARN, "instanceRoleARN", path); err != nil {
			return err
		}

		errNotSupported := func(field string) error {
			return fmt.Errorf("%s is not supported for Managed Nodes (%s.%s)", field, path, field)
		}

		if ng.IAM.InstanceProfileARN != "" {
			return errNotSupported("instanceProfileARN")
		}
	}

	// TODO fix error messages to not use CLI flags
	if ng.MinSize == nil {
		if ng.DesiredCapacity == nil {
			defaultNodeCount := DefaultNodeCount
			ng.MinSize = &defaultNodeCount
		} else {
			ng.MinSize = ng.DesiredCapacity
		}
	} else if ng.DesiredCapacity != nil && *ng.DesiredCapacity < *ng.MinSize {
		return fmt.Errorf("cannot use --nodes-min=%d and --nodes=%d at the same time", *ng.MinSize, *ng.DesiredCapacity)
	}

	// Ensure MaxSize is set, as it is required by the ASG cfn resource
	if ng.MaxSize == nil {
		if ng.DesiredCapacity == nil {
			ng.MaxSize = ng.MinSize
		} else {
			ng.MaxSize = ng.DesiredCapacity
		}
	} else if ng.DesiredCapacity != nil && *ng.DesiredCapacity > *ng.MaxSize {
		return fmt.Errorf("cannot use --nodes-max=%d and --nodes=%d at the same time", *ng.MaxSize, *ng.DesiredCapacity)
	} else if *ng.MaxSize < *ng.MinSize {
		return fmt.Errorf("cannot use --nodes-min=%d and --nodes-max=%d at the same time", *ng.MinSize, *ng.MaxSize)
	}

	if ng.DesiredCapacity == nil {
		ng.DesiredCapacity = ng.MinSize
	}
	return nil
}

func validateInstancesDistribution(ng *NodeGroup) error {
	if ng.InstancesDistribution == nil {
		return nil
	}

	if ng.InstanceType != "" && ng.InstanceType != "mixed" {
		return fmt.Errorf("instanceType should be \"mixed\" or unset when using the mixed instances feature")
	}

	distribution := ng.InstancesDistribution
	if distribution.InstanceTypes == nil || len(distribution.InstanceTypes) == 0 {
		return fmt.Errorf("at least two instance types have to be specified for mixed nodegroups")
	}

	allInstanceTypes := make(map[string]bool)
	for _, instanceType := range distribution.InstanceTypes {
		allInstanceTypes[instanceType] = true
	}

	if len(allInstanceTypes) < 1 || len(allInstanceTypes) > 20 {
		return fmt.Errorf("mixed nodegroups should have between 1 and 20 different instance types")
	}

	if distribution.OnDemandBaseCapacity != nil && *distribution.OnDemandBaseCapacity < 0 {
		return fmt.Errorf("onDemandBaseCapacity should be 0 or more")
	}

	if distribution.OnDemandPercentageAboveBaseCapacity != nil && (*distribution.OnDemandPercentageAboveBaseCapacity < 0 || *distribution.OnDemandPercentageAboveBaseCapacity > 100) {
		return fmt.Errorf("percentageAboveBase should be between 0 and 100")
	}

	if distribution.SpotInstancePools != nil && (*distribution.SpotInstancePools < 1 || *distribution.SpotInstancePools > 20) {
		return fmt.Errorf("spotInstancePools should be between 1 and 20")
	}

	if distribution.SpotInstancePools != nil && distribution.SpotAllocationStrategy != nil && *distribution.SpotAllocationStrategy == SpotAllocationStrategyCapacityOptimized {
		return fmt.Errorf("spotInstancePools cannot be specified when also specifying spotAllocationStrategy: %s", SpotAllocationStrategyCapacityOptimized)
	}

	if distribution.SpotAllocationStrategy != nil {
		if !isSpotAllocationStrategySupported(*distribution.SpotAllocationStrategy) {
			return fmt.Errorf("spotAllocationStrategy should be one of: %v", strings.Join(supportedSpotAllocationStrategies(), ", "))
		}
	}

	return nil
}

func validateNodeGroupSSH(SSH *NodeGroupSSH) error {
	numSSHFlagsEnabled := countEnabledFields(
		SSH.PublicKeyPath,
		SSH.PublicKey,
		SSH.PublicKeyName)

	if numSSHFlagsEnabled > 1 {
		return errors.New("only one SSH public key can be specified per node-group")
	}
	return nil
}

func countEnabledFields(fields ...*string) int {
	count := 0
	for _, flag := range fields {
		if flag != nil && *flag != "" {
			count++
		}
	}
	return count
}

func validateNodeGroupKubeletExtraConfig(kubeletConfig *InlineDocument) error {
	if kubeletConfig == nil {
		return nil
	}

	var kubeletForbiddenFields = map[string]struct{}{
		"kind":               {},
		"apiVersion":         {},
		"address":            {},
		"clusterDomain":      {},
		"authentication":     {},
		"authorization":      {},
		"serverTLSBootstrap": {},
	}

	for k := range *kubeletConfig {
		if _, exists := kubeletForbiddenFields[k]; exists {
			return fmt.Errorf("cannot override %q in kubelet config, as it's critical to eksctl functionality", k)
		}
	}
	return nil
}

// IsWindowsImage reports whether the AMI family is for Windows
func IsWindowsImage(imageFamily string) bool {
	return imageFamily == NodeImageFamilyWindowsServer2019CoreContainer ||
		imageFamily == NodeImageFamilyWindowsServer2019FullContainer ||
		imageFamily == NodeImageFamilyWindowsServer1909CoreContainer
}

func validateCIDRs(cidrs []string) ([]string, error) {
	var validCIDRs []string
	for _, cidr := range cidrs {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
		validCIDRs = append(validCIDRs, ipNet.String())
	}
	return validCIDRs, nil
}

// ReservedProfileNamePrefix defines the Fargate profile name prefix reserved
// for AWS, and which therefore, cannot be used by users. AWS' API should
// reject the creation of profiles starting with this prefix, but we eagerly
// validate this client-side.
const ReservedProfileNamePrefix = "eks-"

// Validate validates this FargateProfile object.
func (fp FargateProfile) Validate() error {
	if fp.Name == "" {
		return errors.New("invalid Fargate profile: empty name")
	}
	if strings.HasPrefix(fp.Name, ReservedProfileNamePrefix) {
		return fmt.Errorf("invalid Fargate profile %q: name should NOT start with %q", fp.Name, ReservedProfileNamePrefix)
	}
	if len(fp.Selectors) == 0 {
		return fmt.Errorf("invalid Fargate profile %q: no profile selector", fp.Name)
	}
	for i, selector := range fp.Selectors {
		if err := selector.Validate(); err != nil {
			return errors.Wrapf(err, "invalid Fargate profile %q: invalid profile selector at index #%v", fp.Name, i)
		}
	}
	return nil
}

// Validate validates this FargateProfileSelector object.
func (fps FargateProfileSelector) Validate() error {
	if fps.Namespace == "" {
		return errors.New("empty namespace")
	}
	return nil
}

func checkBottlerocketSettings(doc *InlineDocument, path string) error {
	if doc == nil {
		return nil
	}

	overlapErr := func(key, ngField string) error {
		return errors.Errorf("invalid Bottlerocket setting: use %s.%s instead (path=%s)", path, ngField, key)
	}

	// Dig into kubernetes settings if provided.
	kubeVal, ok := (*doc)["kubernetes"]
	if !ok {
		return nil
	}

	kube, ok := kubeVal.(map[string]interface{})
	if !ok {
		return errors.New("invalid kubernetes settings provided: expected a map of settings")
	}

	checkMapping := map[string]string{
		"node-labels":    "labels",
		"node-taints":    "taints",
		"max-pods":       "maxPodsPerNode",
		"cluster-dns-ip": "clusterDNS",
	}

	for checkKey, shouldUse := range checkMapping {
		_, ok := kube[checkKey]
		if ok {
			return overlapErr(path+".kubernetes."+checkKey, shouldUse)
		}
	}

	return nil
}
