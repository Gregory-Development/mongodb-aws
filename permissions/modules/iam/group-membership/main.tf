resource "aws_iam_group_membership" "eksctl-service-account-group-memebership" {
  group = "${var.group}"
  name = "eksctl-service-account-group-memebership"
  users = ["${var.users}"]
}