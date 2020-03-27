package manager

import (
	"fmt"

	gfn "github.com/awslabs/goformation/cloudformation"
	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/cfn/outputs"
	"github.com/weaveworks/eksctl/pkg/vpc"
)

// FixClusterCompatibility adds any resources missing in the CloudFormation stack in order to support new features
// like Managed Nodegroups and Fargate
func (c *StackCollection) FixClusterCompatibility() error {
	logger.Info("checking cluster stack for missing resources")
	stack, err := c.DescribeClusterStack()
	if err != nil {
		return err
	}

	var (
		clusterDefaultSG string
		fargateRole      string
	)

	featureOutputs := map[string]outputs.Collector{
		// available on clusters created after Managed Nodes support was out
		outputs.ClusterDefaultSecurityGroup: func(v string) error {
			clusterDefaultSG = v
			return nil
		},
		// available on 1.14 clusters created after Fargate support was out
		outputs.FargatePodExecutionRoleARN: func(v string) error {
			fargateRole = v
			return nil
		},
	}

	if err := outputs.Collect(*stack, nil, featureOutputs); err != nil {
		return err
	}

	stackSupportsManagedNodes := false
	if clusterDefaultSG != "" {
		stackSupportsManagedNodes, err = c.hasManagedToUnmanagedSG()
		if err != nil {
			return err
		}
	}

	stackSupportsFargate := fargateRole != ""

	// Check public subnets
	publicSubnetsAreValid, err := c.isMapPublicIpOnLaunchEnabled()
	if err != nil {
		return err
	}

	if stackSupportsManagedNodes && stackSupportsFargate && publicSubnetsAreValid {
		logger.Info("cluster stack has all required resources")
		return nil
	}

	if !stackSupportsManagedNodes {
		logger.Info("cluster stack is missing resources for Managed Nodegroups")
	}
	if !stackSupportsFargate {
		logger.Info("cluster stack is missing resources for Fargate")
	}
	if !publicSubnetsAreValid {
		logger.Info("public subnets don't have MapPublicIpOnLaunch enabled")
	}

	logger.Info("adding missing resources to cluster stack")
	_, err = c.AppendNewClusterStackResource(false, true)
	return err
}

func (c *StackCollection) hasManagedToUnmanagedSG() (bool, error) {
	stackTemplate, err := c.GetStackTemplate(c.makeClusterStackName())
	if err != nil {
		return false, errors.Wrap(err, "error getting cluster stack template")
	}
	stackResources := gjson.Get(stackTemplate, resourcesRootPath)
	return builder.HasManagedNodesSG(&stackResources), nil
}

// ensureMapPublicIpOnLaunchEnabled sets this subnet property to true when it is not set or is set to false
// returns the modified template and the list of modified subnets
func (c *StackCollection) ensureMapPublicIpOnLaunchEnabled(currentTemplate string) (string, []string, error) {
	outputTemplate := gjson.Get(currentTemplate, outputsRootPath)
	publicSubnetsNames, err := getPublicSubnetResourceNames(outputTemplate.Raw)
	if err != nil {
		// Subnets do not appear in the stack -> they were imported -> check their configuration in EC2
		err = vpc.ValidateExistingPublicSubnets(c.provider, c.spec.PublicSubnetIDs())
		if err != nil {
			return "", nil, errors.Wrapf(err, "mis-configured imported subnets. Expected property MapPublicIpOnLaunch enabled")
		}
		return currentTemplate, make([]string, 0), nil
	}

	// Modify the subnets' properties in the stack
	logger.Debug("ensuring subnets have MapPublicIpOnLaunch enabled")
	modifiedResources := make([]string, 0)
	for _, subnet := range publicSubnetsNames {
		path := subnetResourcePath(subnet)

		currentValue := gjson.Get(currentTemplate, path)
		if !currentValue.Exists() || !currentValue.Bool() {
			currentTemplate, err = sjson.Set(currentTemplate, path, gfn.True())
			if err != nil {
				return "", nil, errors.Wrapf(err, "unable to set MapPublicIpOnLaunch property on subnet %q", path)
			}
			modifiedResources = append(modifiedResources, subnet)
		}
	}
	return currentTemplate, modifiedResources, nil
}

func (c *StackCollection) isMapPublicIpOnLaunchEnabled() (bool, error) {
	currentTemplate, err := c.GetStackTemplate(c.makeClusterStackName())
	if err != nil {
		return false, errors.Wrapf(err, "error getting stack template %s", c.makeClusterStackName())
	}

	outputTemplate := gjson.Get(currentTemplate, outputsRootPath)
	publicSubnetsNames, err := getPublicSubnetResourceNames(outputTemplate.Raw)
	if err != nil {
		// Subnets do not appear in the stack -> they were imported -> check their configuration in EC2
		err = vpc.ValidateExistingPublicSubnets(c.provider, c.spec.PublicSubnetIDs())
		if err != nil {
			return false, nil
		}
		return true, nil
	}

	// Subnets appear in the stack
	for _, subnet := range publicSubnetsNames {
		currentValue := gjson.Get(currentTemplate, subnetResourcePath(subnet))
		if !currentValue.Exists() || !currentValue.Bool() {
			return false, nil
		}
	}
	return true, nil
}

func subnetResourcePath(subnetName string) string {
	return fmt.Sprintf("Resources.%s.Properties.MapPublicIpOnLaunch", subnetName)
}
