provider "aws" {}

module "ec2_1" {
  source = "./mods/ec2_instances"
  name = "MONGODB_1"
  security_groups = "${module.sg.name}"
}

module "ec2_2" {
  source = "./mods/ec2_instances"
  name = "MONGODB_2"
  security_groups = "${module.sg.name}"
}

module "ec2_3" {
  source = "./mods/ec2_instances"
  name = "MONGODB_3"
  security_groups = "${module.sg.name}"
}

module "ec2_4" {
  source = "./mods/ec2_instances"
  name = "MONGODB_4"
  security_groups = "${module.sg.name}"
}

module "ec2_5" {
  source = "./mods/ec2_instances"
  name = "MONGODB_5"
  security_groups = "${module.sg.name}"
}

module "sg" {
  source = "./mods/security_groups"
  name = "MONGO_DB_SECURITY_GROUP"
}

resource "local_file" "file" {
  filename = "../ansible/hosts"
  content = <<EOF
[mongo-hosts]
${module.ec2_1.pub_ip_address} ansible_ssh_private_key_file=~/.ssh/mongo.pem
${module.ec2_2.pub_ip_address} ansible_ssh_private_key_file=~/.ssh/mongo.pem
${module.ec2_3.pub_ip_address} ansible_ssh_private_key_file=~/.ssh/mongo.pem
${module.ec2_4.pub_ip_address} ansible_ssh_private_key_file=~/.ssh/mongo.pem
${module.ec2_5.pub_ip_address} ansible_ssh_private_key_file=~/.ssh/mongo.pem
EOF
}