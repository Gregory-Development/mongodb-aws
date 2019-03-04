package main

import (
	"flag"
	"gitlab/blackbird.ai/mongo-aws/tools/config"
	"log"
	"os"
)

var c config.Config

func main () {
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
}
