package main

import (
	"bitbucket.org/blackbird.ai/mongo-aws/tools/cmd/kubernetes"
	"bitbucket.org/blackbird.ai/mongo-aws/tools/config"
	"flag"
	"log"
	"os"
)

var c config.Config
var k8s kubernetes.K8S

func main() {
	configFile := flag.String("config-file", "./configuration.yml", "The configuration file to use.")
	flag.Parse()

	if *configFile == "" {
		log.Fatalf("no config file set")
		os.Exit(1) // Exit application
	}

	err := c.ParseConfigFile(*configFile)
	if err != nil {
		log.Fatalf("opening the configuration file %v was not successful with error: %v", *configFile, err)
		os.Exit(1)
	}

	err = c.ValidateDependencies()
	if err != nil {
		log.Fatalf("error installing dependencies: %v", err)
		os.Exit(1)
	}

	err = c.ValidateAWSConfig()
	if err != nil {
		log.Fatalf("error setting AWS environment variables: %v", err)
	}

	k8s.NewCluster(c.Config.Kubernetes.StateStore, c.Config.Kubernetes.ClusterName)
}
