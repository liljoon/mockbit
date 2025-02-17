# General
variable "name_prefix" {
  type = string
}

variable "location" {
  type = string
}

variable "subscription_id" {
  type = string
}

# AKS
variable "node_count" {
  type = number
}

variable "vm_size" {
  type = string
}

# DB
variable "db_administrator_login" {
  type = string
}

variable "db_administrator_password" {
  type = string
}

variable "db_sku_name" {
  type = string
}