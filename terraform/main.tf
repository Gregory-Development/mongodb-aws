provider "aws" {
  region = "us-east-2"
}

module "ec2_1" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_0_0"
  security_groups = "${module.sg.name}"
}

module "ec2_2" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_0_1"
  security_groups = "${module.sg.name}"
}

module "ec2_3" {
  source = "./mods/ec2_instances"
  name = "MONGODB_CONFIG_0_2"
  security_groups = "${module.sg.name}"
}

module "ec2_4" {
  source = "./mods/ec2_instances"
  name = "MONGODB_QUERY_ROUTER_0_0"
  security_groups = "${module.sg.name}"
}

module "ec2_5" {
  source = "./mods/ec2_instances"
  name = "MONGODB_QUERY_ROUTER_0_1"
  security_groups = "${module.sg.name}"
}

module "ec2_6" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_0_0"
  security_groups = "${module.sg.name}"
}

module "ec2_7" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_0_1"
  security_groups = "${module.sg.name}"
}

module "ec2_8" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_1_0"
  security_groups = "${module.sg.name}"
}

module "ec2_9" {
  source = "./mods/ec2_instances"
  name = "MONGODB_SHARD_1_1"
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
${module.ec2_5.pub_ip_address}

[mongo-shard-0-hosts]
${module.ec2_6.pub_ip_address}
${module.ec2_7.pub_ip_address}

[mongo-shard-1-hosts]
${module.ec2_8.pub_ip_address}
${module.ec2_9.pub_ip_address}

[mongo-shard-hosts:children]
mongo-shard-0-hosts
mongo-shard-1-hosts

[mongo:children]
mongo-config-hosts
mongo-query-router-hosts
mongo-shard-hosts

[mongo:vars]
ansible_ssh_private_key_file=~/.ssh/mongo.pem
ansible_user=ec2-user
EOF
}

resource "local_file" "mogno_keyfile" {
  filename = "../ansible/roles/install_mongo/templates/mongodb-keyfile.j2"
  content = <<EOF
g/aGilRpAUSx1L8Um5ydG1OXAzHWaCQZL6tRr5axl41GaNZgSIaLd/1uQfUt7cNM
/v8EVbAWE7+thYwKAo97j/u5AceypIbgmiBIT/+I/yXeE4cim+lrENpD1aCBkmCS
Ao98rasCJlKSoCH1Ai90/aMvJTW+6oAgSpRJRE+/XjvG5AV/6xyKshyj7zDlyIkF
sG7FqVbFJQJHvH6Zw4PCUPWGDgrmD+9GxuifdSs+i5A7Maqiq4aM1xq4gdDXBWD4
sA8zI9oguDgUzz9XZmirQKD9sfWy83YFa+TvPuHbAeIbUdLCzzFaujk/OjEoObr7
CAa9Z8OEnHJ03AReYAtIZJs63nt9oTYAzmnOWc4wkNLTapG52BXuzwautHgj/QR3
USYLxa27tSMvXeE0sLBwMqnEHy2mkfMS8j5ATG572iLWRtUeiwVLaMO6J9zpZ6+O
2L2ES2uorFWGj2MMDzYHOUd5jCfDGKYcIDCNfdAr1L1fuTneTisknZ0OJnopSdNd
3FK+EIk9FvZ03fYlaLYzhItHn2w9tcgHtWxRs11/2DZfqz3M0URax0gaapy+e2dv
lAVCYsFdqipcCc7Jof429XF3p0Duwy5TKFWFE2D/bbmKqS5FufDTP6dizOKewyfF
9cGy23/kjfzyhPZH76Rw41eywUJXllS7x6N6nqcEhnGD3FUw9hD1n9dj+7TmIm9N
Fl0PZTrMnfp3AoimiB4b8a/QbOACWGQ5gHcjddldLrTXOMAwe5gsWz7U+zqfdJKQ
vxahHwt/O163cx/GheO7+RfvdJAzUB8VW6v2Cp/QqQNR7mSc371kkInpgf0vSpE6
pRJPapPSMQshkSEYlCSx07t2ibqyJrvfBTgTxnmbJl2SEPCga3yxAlzd/M7Vydqz
eTwRTtR41DPKTWUQEJqWIgVN4354LFn46qiML2V3I7RUVquX1WZ3ssAYCFyRxUWt
ips9an9Aygh25VpaMn9Sl6sHOcmwKibS95JQkkCmyJ26JW+k
EOF
}