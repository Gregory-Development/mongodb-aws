resource "aws_iam_access_key" "eksctl-access-key" {
  user = "${aws_iam_user.eksctl-user.name}"
}

resource "aws_iam_user" "eksctl-user" {
  name = "eksctl-service-account"
}

resource "local_file" "credentials_file" {
  filename = "./credentials"
  content = <<EOF
[eksctl]
region = ${var.region}
aws_access_key_id = ${aws_iam_access_key.eksctl-access-key.id}
aws_secret_access_key = ${aws_iam_access_key.eksctl-access-key.secret}
EOF
}