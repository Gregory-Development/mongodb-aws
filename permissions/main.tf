provider "aws" {
  region = "us-west-1"
}

module "eksctl-user-group-policy" {
  source = "./modules/iam"
  update-env = ""
}