package config

import (
	"errors"
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

}

type aws struct {
	AccessKeyId string `yaml:"access-key-id"`
	SecretAccessKey string `yaml:"secret-access-key"`
	Region string `yaml:"region"`
}

func (c *Config) ValidateConfigFile (filename string) (err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.New("unable to read the file provided")
		return
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		err = errors.New("invalid yaml file")
		return
	}

	err = nil
	return
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

	err = nil
	return
}