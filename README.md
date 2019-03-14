# MongoDB Clusters on Kubernetes on AWS via KOPS/HELM
---

## Prereqs

_GOLANG_
[https://golang.org/doc/install](https://golang.org/doc/install)

_JQ_
`sudo yum install jq -y` or `sudo dnf install jq -y` or `sudo apt-get install jq -y`

_KUBECTL_

`sudo curl --silent --location -o /usr/local/bin/kubectl "https://amazon-eks.s3-us-west-2.amazonaws.com/1.11.5/2018-12-06/bin/linux/amd64/kubectl"`

`sudo chmod +x /usr/local/bin/kubectl`

_AWS-IAM-Authenticator_

`go get -u -v github.com/kubernetes-sigs/aws-iam-authenticator/cmd/aws-iam-authenticator`

`sudo mv ~/go/bin/aws-iam-authenticator /usr/local/bin/aws-iam-authenticator`

_EKSCTL_

`curl --silent --location "https://github.com/weaveworks/eksctl/releases/download/latest_release/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp`

`sudo mv -v /tmp/eksctl /usr/local/bin`

_HELM_

`curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh`

`chmod +x get_helm.sh`

`./get_helm.sh`

## Starting a new EKS cluster and installing helm

First we need to build a new EKS (Kubernetes) cluster -- this will take roughly 15 minutes the first time

```bash
eksctl create cluster \
  --name=<unique name for your cluster - string> \
  --nodes=<number of worker nodes - int> \
  --node-type=<the ec2 instance size to use> \
  --node-ami=auto \
  --node-volume-size=<the size of the node volumes to create> \
  --node-volume-type=<the type of volume to create - e.g. gp2> \
  --region=<the aws region to deploy to - string> \
  --asg-access
```

By default, this command will create all of the necessary files to connect directly to you cluster (located in ./kubeconfig/config)
You can verify that it works with the command `kubectl get nodes`

We now need to enable role based access control on the cluster (RBAC), to do this, do the following:

```bash
cat <<EoF > ~/rbac.yaml
 ---
 apiVersion: v1
 kind: ServiceAccount
 metadata:
   name: tiller
   namespace: kube-system
 ---
 apiVersion: rbac.authorization.k8s.io/v1beta1
 kind: ClusterRoleBinding
 metadata:
   name: tiller
 roleRef:
   apiGroup: rbac.authorization.k8s.io
   kind: ClusterRole
   name: cluster-admin
 subjects:
   - kind: ServiceAccount
     name: tiller
     namespace: kube-system
 EoF
```
 
Now, we have to apply this config to our cluster and set up our local helm:
 
`kubectl apply -f ~/rbac.yaml`
 
`helm init --service-account tiller`

We now want to update all of our helm repositories:

`helm repo update`

Now we want to install MongoDB

`curl -O https://raw.githubusercontent.com/kubernetes/charts/master/stable/mongodb/values-production.yaml`

`helm install --name <name for the mongodb cluster> -f ./values-production.yaml stable/mongodb`

## Scaling

### EKS

To manually scale the worker nodes:

`eksctl scale nodegroup --cluster=<name of your cluster> --nodes=<desired count> --name=<the name of the nodegroup>`

### MongoDB

To scale mongo, simply run the following commands:

`kubectl scale statefulset <the name of your mongodb cluster>-mongodb-secondary --replicas=<number of replicas to spin up>`