package config

import (
	"bitbucket.org/blackbird.ai/mongo-aws/tools/dependencies"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Config config `yaml:"config"`
}

type config struct {
	System     system     `yaml:"system"`
	Kubernetes kubernetes `yaml:"kubernetes"`
	AWS        aws        `yaml:"amazon-aws"`
}

type system struct {
	KubectlExecutable   string `yaml:"kubectl-executable-path"`
	KopsExecutable      string `yaml:"kops-executable-path"`
	TerraformExecutable string `yaml:"terraform-executable-path"`
	PipExecutable       string `yaml:"pip-executable-path"`
	AwsCliExecutable    string `yaml:"awscli-executable-path"`
}

type kubernetes struct {
	StateStore  string   `yaml:"state-store"`
	Zones       []string `yaml:"zones"`
	NodeCount   int      `yaml:"node-count"`
	NodeSize    string   `yaml:"node-size"`
	MasterSize  string   `yaml:"master-size"`
	DNSZone     string   `yaml:"dns-zone"`
	ClusterName string   `yaml:"cluster-name"`
}

type aws struct {
	AccessKeyId     string `yaml:"access-key-id"`
	SecretAccessKey string `yaml:"secret-access-key"`
	Region          string `yaml:"region"`
}

var s dependencies.SYS

func (c *Config) ParseConfigFile(filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return
	}
	err = nil
	return
}

func (c *Config) ValidateDependencies() (err error) {

	err = dependencies.ValidateExecutables(c.Config.System.KubectlExecutable)
	if err != nil {
		s.InstallDependencies("kubectl")
	}

	err = dependencies.ValidateExecutables(c.Config.System.KopsExecutable)
	if err != nil {
		s.InstallDependencies("kops")
	}

	err = dependencies.ValidateExecutables(c.Config.System.TerraformExecutable)
	if err != nil {
		s.InstallDependencies("terraform")
	}

	err = dependencies.ValidateExecutables(c.Config.System.PipExecutable)
	if err != nil {
		s.InstallDependencies("pip")
	}

	err = dependencies.ValidateExecutables(c.Config.System.AwsCliExecutable)
	if err != nil {
		s.InstallDependencies("aws")
	}

	err = nil
	return
}

func (c *Config) ValidateAWSConfig() (err error) {
	akid := os.Getenv("AWS_ACCESS_KEY_ID")
	sak := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")

	if akid == "" {
		err = os.Setenv("AWS_ACCESS_KEY_ID", c.Config.AWS.AccessKeyId)
		if err != nil {
			return
		}
	}
	if sak == "" {
		err = os.Setenv("AWS_SECRET_ACCESS_KEY", c.Config.AWS.SecretAccessKey)
		if err != nil {
			return
		}
	}
	if region == "" {
		err = os.Setenv("AWS_DEFAULT_REGION", c.Config.AWS.Region)
		if err != nil {
			return
		}
	}
	err = nil
	return
}
