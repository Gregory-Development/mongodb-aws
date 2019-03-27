# Deploy Highly Available MongoDB Cluster to AWS via Terraform/Ansible

## Be Patient ... docs are coming

### Terraform How To

This does assume you have terraform installed, if you don't, see [their documentation](https://www.terraform.io/downloads.html) for installing.

To create the base infrastructure, run the following commands (after cloning this repo and "cd"ing to it of course)

* `cd ./terraform`
* if you are using AWS Profiles (like i am), do this: `AWS_PROFILE=<your profile name> terraform init`
  * if you are NOT using profiles, make sure your AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and AWS_REGION environmental variables are set and then run `terraform init`

if you need to get rid of the infrastructure for any reason, just run:
* `AWS_PROFILE terraform destroy` or `terraform destroy` if you are doing the 3 AWS var method

### Ansible

This does assume you have ansible installed, if you don't, see [their documentation](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html) for installing.

To set up the cluster, run the following commands (after setting up the terraform infrastructure)

* `cd ./ansible/roles` or `cd ../ansible/roles` if you are still in the terraform directory
* `ansible-playbook -i hosts playbook.yml`