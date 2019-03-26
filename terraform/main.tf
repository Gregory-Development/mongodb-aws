provider "aws" {
  region = "us-east-2"
}

module "ec2_1" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_1"
  security_groups = "${module.sg.name}"
}

module "ec2_2" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_2"
  security_groups = "${module.sg.name}"
}

module "ec2_3" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_3"
  security_groups = "${module.sg.name}"
}

module "ec2_4" {
  source = "./mods/ec2_instances"
  name = "MONGODB_QUERY_ROUTER_1"
  security_groups = "${module.sg.name}"
}

module "ec2_5" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_1"
  security_groups = "${module.sg.name}"
}

module "ec2_6" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_2"
  security_groups = "${module.sg.name}"
}

module "sg" {
  source = "./mods/security_groups"
  name = "MONGO_DB_SECURITY_GROUP"
}

resource "local_file" "file" {
  filename = "../ansible/hosts"
  content = <<EOF
[mongo-config-hosts]
${module.ec2_1.pub_ip_address}
${module.ec2_2.pub_ip_address}
${module.ec2_3.pub_ip_address}

[mongo-query-router-hosts]
${module.ec2_4.pub_ip_address}

[mongo-shard-hosts]
${module.ec2_5.pub_ip_address}
${module.ec2_6.pub_ip_address}

[mongo:children]
mongo-config-hosts
mongo-query-router-hosts
mongo-shards

[mongo:vars]
ansible_ssh_private_key_file=~/.ssh/mongo.pem
ansible_user=ec2-user
EOF
}