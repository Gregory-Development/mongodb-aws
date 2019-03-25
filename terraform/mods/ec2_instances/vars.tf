variable "storage_type" {
  description = "the type of storage to use"
  default = "gp2"
}

variable "storage_size" {
  description = "the size of storage to allocate"
  default = 100
}

variable "use_public_ip_address" {
  description = "whether or not to assign a public ip address"
  default = true
}

variable "security_groups" {
  description = "the security groups to assign to the instance"
}

variable "name" {
  description = "the name to assign to the instance"
}

variable "key_name" {
  description = "the name of the key to use"
  default = "mongo"
}