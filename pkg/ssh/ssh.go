package ssh

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	"k8s.io/kops/pkg/pki"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha4"
	"github.com/weaveworks/eksctl/pkg/utils"
)

// LoadSSHKeyFromFile loads and imports a public SSH key from a file provided a path to that file
func LoadSSHKeyFromFile(filePath, clusterName, ngName string, provider api.ClusterProvider) error {
	if err := checkFileExists(filePath); err != nil {
		return errors.Wrap(err, fmt.Sprintf("SSH public key file %q not found", filePath))
	}

	expandedPath := utils.ExpandPath(filePath)
	fileContent, err := readFileContents(expandedPath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("reading SSH public key file %q", filePath))
	}

	key := string(fileContent)

	fingerprint, err := pki.ComputeAWSKeyFingerprint(key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("computing fingerprint for key %q", filePath))
	}
	keyName := getSSHKeyName(clusterName, ngName, fingerprint)

	logger.Info("using SSH public key %q as %q ", expandedPath, keyName)

	// Import SSH key in EC2
	if err := importSSHPublicKey(keyName, fingerprint, &key, provider); err != nil {
		return err
	}
	return nil
}

// LoadSSHKeyByContent loads and imports an SSH public key into EC2 if it doesn't exist
func LoadSSHKeyByContent(key *string, clusterName, ngName string, provider api.ClusterProvider) error {
	fingerprint, err := pki.ComputeAWSKeyFingerprint(*key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("computing fingerprint for key %q", *key))
	}
	keyName := getSSHKeyName(clusterName, ngName, fingerprint)

	logger.Info("using SSH public key %q ", *key)

	// Import SSH key in EC2
	if err := importSSHPublicKey(keyName, fingerprint, key, provider); err != nil {
		return err
	}
	return nil
}

// DeletePublicSSHKeys will delete the public SSH key, if it exists
func DeletePublicSSHKeys(clusterName string, provider api.ClusterProvider) {
	existing, err := provider.EC2().DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	if err != nil {
		logger.Debug("cannot describe keys: %v", err)
		return
	}
	var matching []*string
	prefix := getSSHKeyName(clusterName, "", "")
	logger.Debug("existing = %#v", existing)
	for _, e := range existing.KeyPairs {
		if !strings.HasPrefix(*e.KeyName, prefix) {
			continue
		}
		nameParts := strings.Split(*e.KeyName, "-")
		logger.Debug("existing key %q matches prefix", *e.KeyName)
		if nameParts[len(nameParts)-1] == *e.KeyFingerprint {
			logger.Debug("existing key %q matches fingerprint", *e.KeyName)
			matching = append(matching, e.KeyName)
		}
	}
	for i := range matching {
		input := &ec2.DeleteKeyPairInput{
			KeyName: matching[i],
		}
		logger.Debug("deleting key %q", *matching[i])
		if _, err := provider.EC2().DeleteKeyPair(input); err != nil {
			logger.Debug("key pair couldn't be deleted: %v", err)
		}
	}
}

// CheckKeyExistsInEc2 returns whether a public ssh key already exists in EC2 or error if it couldn't be checked
func CheckKeyExistsInEc2(sshPublicKeyName string, provider api.ClusterProvider) error {
	existing, err := findSSHKeyInEc2(sshPublicKeyName, provider)
	if err != nil {
		return errors.Wrap(err, "checking existing key pair")
	}

	if existing == nil {
		return fmt.Errorf("cannot find EC2 key pair %q", sshPublicKeyName)
	}

	return nil
}

// FileExists returns true if there is a local file with that path
func FileExists(filePath string) bool {
	extendedPath := utils.ExpandPath(filePath)
	if _, err := os.Stat(extendedPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func importSSHPublicKey(keyName, fingerprint string, keyContent *string, provider api.ClusterProvider) error {
	if existing, err := findSSHKeyInEc2(keyName, provider); err != nil {
		return err
	} else if existing != nil {
		if *existing.KeyFingerprint != fingerprint {
			return fmt.Errorf("SSH public key %s already exists, but fingerprints don't match (exected: %q, got: %q)", keyName, fingerprint, *existing.KeyFingerprint)
		}

		logger.Debug("SSH public key %s already exists", keyName)
		return nil
	}

	// Import it
	input := &ec2.ImportKeyPairInput{
		KeyName:           &keyName,
		PublicKeyMaterial: []byte(*keyContent),
	}
	logger.Debug("importing SSH public key %q", keyName)

	if _, err := provider.EC2().ImportKeyPair(input); err != nil {
		return errors.Wrap(err, "importing SSH public key")
	}
	return nil
}

// getSSHKeyName generates the name of an SSH key based on the cluster name, nodegroup name and fingerprint
// in the form "eksctl-<clusterName>-nodegroup-<nodeGroupName>-<fingerprint>"
func getSSHKeyName(clusterName, nodeGroupName, fingerprint string) string {
	keyNameParts := []string{"eksctl", clusterName}
	if nodeGroupName != "" {
		keyNameParts = append(keyNameParts, fmt.Sprintf("nodegroup-%s", nodeGroupName))
	}

	if fingerprint != "" {
		keyNameParts = append(keyNameParts, fingerprint)
	}

	return strings.Join(keyNameParts, "-")
}

func findSSHKeyInEc2(name string, provider api.ClusterProvider) (*ec2.KeyPairInfo, error) {
	input := &ec2.DescribeKeyPairsInput{
		KeyNames: aws.StringSlice([]string{name}),
	}
	output, err := provider.EC2().DescribeKeyPairs(input)

	if err != nil {
		awsError := err.(awserr.Error)
		if awsError.Code() == "InvalidKeyPair.NotFound" {
			return nil, nil
		}
		return nil, errors.Wrapf(err, fmt.Sprintf("searching for SSH public key %q in EC2", name))
	}

	if len(output.KeyPairs) != 1 {
		logger.Debug("output = %#v", output)
		return nil, fmt.Errorf("unexpected number of key pairs found (expected: 1, got: %d)", len(output.KeyPairs))
	}
	return output.KeyPairs[0], nil
}

func checkFileExists(filePath string) error {
	extendedPath := utils.ExpandPath(filePath)
	if _, err := os.Stat(extendedPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

func readFileContents(filePath string) ([]byte, error) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("reading SSH public key file %q", filePath))
	}
	return fileContents, nil
}