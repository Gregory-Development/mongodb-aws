package kubernetes

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type K8S struct {
	ApiVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   ClusterMetadata `yaml:"metadata"`
	Spec       ClusterSpec     `yaml:"spec"`
}

type ClusterMetadata struct {
	Name string `yaml:"name"`
}

type ClusterSpec struct {
	Api                 ClusterApi        `yaml:"api"`
	Authorization       ClusterAuth       `yaml:"authorization"`
	Channel             string            `yaml:"channel"`
	CloudProvider       string            `yaml:"cloudProvider"`
	ConfigBase          string            `yaml:"configBase"`
	EtcdClusters        []EtcdMembers     `yaml:"etcdClusters"`
	Iam                 Iam               `yaml:"iam"`
	Kubelet             Kubelet           `yaml:"kubelet"`
	KubernetesApiAccess []string          `yaml:"kubernetesApiAccess"`
	KubernetesVersion   string            `yaml:"kubernetesVersion"`
	MasterInternalName  string            `yaml:"masterInternalName"`
	MasterPublicName    string            `yaml:"masterPublicName"`
	NetworkCIDR         string            `yaml:"networkCIDR"`
	Networking          ClusterNetworking `yaml:"networking"`
	NonMasqueradeCIDR   string            `yaml:"nonMasqueradeCIDR"`
	SshAccess           []string          `yaml:"sshAccess"`
	Subnets             []Subnets         `yaml:"subnets"`
	Topology            ClusterTopology   `yaml:"topology"`
}

type ClusterApi struct {
	Dns string `yaml:"dns"`
}

type ClusterAuth struct {
	Rbac string `yaml:"rbac"`
}

type EtcdMembers struct {
	InstanceGroup []InstanceGroup `yaml:"instanceGroup"`
	Name          string          `yaml:"name"`
}

type InstanceGroup struct {
	Name string
}

type Iam struct {
	AllowContainerRegistry bool `yaml:"allowContainerRegistry"`
	Legacy                 bool `yaml:"legacy"`
}

type Kubelet struct {
	AnonymousAuth bool `yaml:"anonymousAuth"`
}

type ClusterNetworking struct {
	Kubelet string `yaml:"kubelet"`
}

type Subnets struct {
	CIDR string `yaml:"cidr"`
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Zone string `yaml:"zone"`
}

type ClusterTopology struct {
	Dns     ClusterDNS `yaml:"dns"`
	Masters string     `yaml:"masters"`
	Nodes   string     `yaml:"nodes"`
}

type ClusterDNS struct {
	Type string `yaml:"type"`
}

func (k *K8S) NewCluster(stateStoreBucketName string, clusterName string) {
	var out bytes.Buffer
	var ssbn = fmt.Sprintf("--state=s3://%v", stateStoreBucketName)
	var cn = fmt.Sprintf("--name=%v", clusterName)

	k.Spec.ConfigBase = fmt.Sprintf("s3://%v/%v", stateStoreBucketName, clusterName)

	install := exec.Command("kops", "create", "cluster", ssbn, cn, "--yes")
	install.Stdout = &out
	err := install.Run()
	if err != nil {
		log.Fatalf("unable to install %v, are you root?", pkg)
		os.Exit(1)
	}
}

func (k *K8S) ModifyCluster() {

}

func (k *K8S) DeleteCluster() {

}
