provider "aws" {
  region = var.region
}

module "lambda_function" {
  source = "terraform-aws-modules/lambda/aws"
  version = "5.3.0"

  function_name = var.lambda_function_name
  description   = var.lambda_function_description
  handler       = "bootstrap"
  runtime       = "go1.x"

  create_package  = false
  local_existing_package = "this_is_for_terraform_apply.zip"

  tags = {
    Name = var.lambda_function_name
  }
}
