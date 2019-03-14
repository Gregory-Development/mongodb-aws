resource "aws_iam_access_key" "eksctl-access-key" {
  user = "${aws_iam_user.eksctl-user.name}"
}

resource "aws_iam_user" "eksctl-user" {
  name = "eksctl-service-account"
}

resource "local_file" "credentials_file_with_export" {
  count = "${var.update-env}"
  filename = "./credentials"
  content = <<EOF
[eksctl]
region = ${var.region}
aws_access_key_id = ${aws_iam_access_key.eksctl-access-key.id}
aws_secret_access_key = ${aws_iam_access_key.eksctl-access-key.secret}
EOF


  provisioner "local-exec" {
    command = "export AWS_ACCESS_KEY_ID=${aws_iam_access_key.eksctl-access-key.id} && export AWS_SECRET_ACCESS_KEY=${aws_iam_access_key.eksctl-access-key.secret} && export AWS_REGION=${var.region}"
  }
}

resource "local_file" "credentials_file_without_export" {
  count = "${1 - var.update-env}"
  filename = "./credentials"
  content = <<EOF
[eksctl]
region = ${var.region}
aws_access_key_id = ${aws_iam_access_key.eksctl-access-key.id}
aws_secret_access_key = ${aws_iam_access_key.eksctl-access-key.secret}
EOF
}