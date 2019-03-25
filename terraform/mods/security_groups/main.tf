resource "aws_security_group" "mongodb" {
  name = "${var.name}"
  description = "security group for mongodb"

  tags {
    Name = "${var.name}"
  }
}

resource "aws_security_group_rule" "mongodb_allow_all" {
  from_port = 0
  protocol = "-1"
  security_group_id = "${aws_security_group.mongodb.id}"
  to_port = 0
  type = "egress"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "mognodb_ssh" {
  from_port = 22
  protocol = "tcp"
  security_group_id = "${aws_security_group.mongodb.id}"
  to_port = 22
  type = "ingress"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "mongodb_mongodb" {
  from_port = 27017
  protocol = "tcp"
  security_group_id = "${aws_security_group.mongodb.id}"
  to_port = 27017
  type = "ingress"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "mongodb_replication" {
  from_port = 27019
  protocol = "tcp"
  security_group_id = "${aws_security_group.mongodb.id}"
  to_port = 27019
  type = "ingress"
  cidr_blocks = ["0.0.0.0/0"]
}