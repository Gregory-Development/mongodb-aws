# MongoDB Clusters on Kubernetes on AWS via KOPS/HELM
---

### Prereqs

The following applications are required to be installed on the provisioning system:

* Kubectl (https://kubernetes.io/docs/tasks/tools/install-kubectl/)
* Kops (https://github.com/kubernetes/kops)
* Python pip (https://pypi.org/project/pip/)
* AWS CLI (https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)

### AWS Prereqs

* A Programmatic access user with the following permissions:
  * AmazonEC2FullAccess
  * AmazonRoute53FullAccess
  * AmazonS3FullAccess
  * IAMFullAccess
  * AmazonVPCFullAccess
  * See [this article](https://github.com/kubernetes/kops/blob/master/docs/aws.md#setup-iam-user) for how to do this programmatically
* An S3 bucket to store the kubernetes cluster state
  * See [the article](https://aws.amazon.com/blogs/compute/kubernetes-clusters-aws-kops/) for how to do this programmatically.
* A DNS (Route53 or other) domain or subdomain to assign to the new kubernetes cluster
  * See [this article](https://aws.amazon.com/blogs/compute/kubernetes-clusters-aws-kops/) for how to do this programmatically.

### Creating the cluster

Using Kops, perform the following:

1) Creating the cluster:
  
  `kops create cluster --state=s3://<THE BUCKET YOU CREATED PREVIOUSLY> --name=<THE DNS ZONE YOU WISH TO USE> --zones=<A COMMA SEPARATED LIST OF AWS AZs TO USE> --node-count=<NUMBER OF NODES TO PROVISION> --node-size=<EC2 INSTANCE SIZE TO PROVISION> --master-count=<NUMBER OF MASTER NODES PER AZ> --master-type=<EC2 INSTANCE SIZE TO PROVISION FOR MASTER NODES>`
  
2) Once complete, Kops will create the state files, to apply the configuration to AWS:
  
  `kops update cluster <WHATEVER YOU PROVISIONED IN --name PREVIOUSLY> --state=s3://<THE BUCKET YOU CREATED IN THE PREREQS> --yes`

3) Kops will then deploy the cluster

4) We then have to install Helm to make application installation much easier
  Follow the instructions [here](https://helm.sh/docs/using_helm/) to install the helm client
  
5) Install Helm to our cluster with the following:
  
  `helm init`
  
6) Grab the MongoDB production deployment Helm Chart template from [here](https://raw.githubusercontent.com/kubernetes/charts/master/stable/mongodb/values-production.yaml)
  Feel free to customize the template as necessary for your environment

7) Deploy the new MongoDB cluster with the following commands:
  
  `helm install --name <WHATEVER YOU WOULD LIKE TO CALL IT> -f <PATH TO THE values-production.yaml FILE> stable/mongodb`

8) Rinse and Repeat steps 6 & 7 to deploy multiple MongoDB clusters
  
### Increasing the amount of Kubernetes worker nodes

To increase the number of worker nodes, perform the following steps:

1) Edit the cluster spec file:
  
  `kops edit instancegroup nodes --name=<THE NAME OF YOUR CLUSTER> --state=s3://<YOUR STATE BUCKET>`
  
2) Find the `maxSize` and `minSize` entries under spec and modify them to you desired number

3) Update the cluster with the new configuration:
  `kops update cluster <THE NAME OF YOUR CLUSTER> --yes`

### Increase the amount of MongoDB replicas

Increasing the number MondoDB replicas running is accomplished with the following command:

`kubectl scale statefulset <THE NAME OF THE MONGODB CLUSTER TO SCALE> --replicas=<NUMBER TO SCALE TO>`