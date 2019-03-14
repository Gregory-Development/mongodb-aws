data "aws_region" "current" {}

module "iam-user" {
  source = "./user"
  region = "${data.aws_region.current.name}"
  update-env = "${var.update-env}"
}

module "group" {
  source = "./group"
}

module "group-policy" {
  source = "./group-policy"
  group = "${module.group.id}"
}

module "group-membership" {
  source = "./group-membership"
  group = "${module.group.name}"
  users = "${module.iam-user.user-name}"
}

locals {}