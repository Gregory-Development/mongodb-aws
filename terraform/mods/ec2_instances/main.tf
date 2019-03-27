resource "aws_instance" "mongodb_one" {
  ami = "ami-02bcbb802e03574ba"
  instance_type = "t2.micro"

  root_block_device {
    volume_size = "${var.storage_size}"
    volume_type = "${var.storage_type}"
  }

  security_groups = [
    "${var.security_groups}"
  ]

  associate_public_ip_address = "${var.use_public_ip_address}"

  key_name = "mongo"

  tags {
    Name = "${var.name}"
  }
}