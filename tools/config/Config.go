package config

import (
	"gitlab/blackbird.ai/mongo-aws/tools/dependencies"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Config config `yaml:"config"`
}

type config struct {
	System system `yaml:"system"`
	Kubernetes kubernetes `yaml:"kubernetes"`
	AWS aws `yaml:"amazon-aws"`
}

type system struct {
	KubectlExecutable string `yaml:"kubectl-executable-path"`
	KopsExecutable string `yaml:"kops-executable-path"`
	TerraformExecutable string `yaml:"terraform-executable-path"`
	PipExecutable string `yaml:"pip-executable-path"`
	AwsCliExecutable string `yaml:"awscli-executable-path"`
}

type kubernetes struct {
	StateStore string `yaml:"state-store"`
	Zones []string `yaml:"zones"`
	NodeCount int `yaml:"node-count"`
	NodeSize string `yaml:"node-size"`
	MasterSize string `yaml:"master-size"`
	DNSZone string `yaml:"dns-zone"`
}

type aws struct {
	AccessKeyId string `yaml:"access-key-id"`
	SecretAccessKey string `yaml:"secret-access-key"`
	Region string `yaml:"region"`
}

func (c *Config) ParseConfigFile (filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return
	}

	err = dependencies.ValidateExecutables(c.Config.System.KubectlExecutable)
	if err != nil {
		dependencies.InstallDependencies("kubectl")
	}

	err = dependencies.ValidateExecutables(c.Config.System.KopsExecutable)
	if err != nil {
		dependencies.InstallDependencies("kops")
	}

	err = dependencies.ValidateExecutables(c.Config.System.TerraformExecutable)
	if err != nil {
		dependencies.InstallDependencies("terraform")
	}

	err = dependencies.ValidateExecutables(c.Config.System.PipExecutable)
	if err != nil {
		dependencies.InstallDependencies("pip")
	}

	err = dependencies.ValidateExecutables(c.Config.System.AwsCliExecutable)
	if err != nil {
		dependencies.InstallDependencies("aws")
	}

	err = nil
	return
}
