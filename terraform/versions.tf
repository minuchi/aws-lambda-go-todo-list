terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "minuchi"

    workspaces {
      name = "aws-lambda-go-todo-list"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.11.0"
    }
  }

  required_version = ">= 1.4.5"
}
