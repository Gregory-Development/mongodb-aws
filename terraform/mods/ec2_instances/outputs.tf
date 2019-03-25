output "pub_ip_address" {
  value = "${aws_instance.mongodb_one.public_ip}"
}