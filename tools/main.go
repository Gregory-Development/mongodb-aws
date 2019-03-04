package main

import (
	"flag"
	"os"
)

var akid  = os.Getenv("AWS_ACCESS_KEY_ID")
var sakid = os.Getenv("AWS_SECRET_ACCESS_KEY")
var rid = os.Getenv("AWS_DEFAULT_REGION")

func main () {
	flag.String("aws-access-key-id", akid, "the aws access key id to use")
	flag.String("aws-secret-access-key", sakid, "the aws secret access key to use")
	flag.String("aws-default-region", rid, "the aws region to use")
	flag.String("k8s-state-location", "", "the S3 bucket to store the k8s state")
	flag.String("aws-zone", "", "the region zone(s) to use during deployment")
	flag.Int("node-count", 0, "the number of nodes to deploy")
	flag.String("aws-k8s-node-instance-size", "t2.micro", "the size of the instance to create for k8s nodes")
	flag.String("aws-k8s-master-instance-size", "t2.micro", "the size of the instance to create for the k8s master node")
	flag.String("dns-zone", "k8s.blackbird.ai", "the dns zone to use for this cluster")
}
