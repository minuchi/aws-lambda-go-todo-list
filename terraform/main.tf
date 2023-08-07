provider "aws" {
  region = var.region
}

module "lambda_function" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "5.3.0"

  function_name = var.lambda_function_name
  description   = var.lambda_function_description
  handler       = "bootstrap"
  runtime       = "go1.x"

  create_package         = false
  local_existing_package = "this_is_for_terraform_apply.zip"

  tags = {
    Name = var.lambda_function_name
  }
}

module "api_gateway" {
  source = "terraform-aws-modules/apigateway-v2/aws"

  name          = var.api_gateway_name
  description   = var.api_gateway_description
  protocol_type = "HTTP"

  cors_configuration = {
    allow_headers = ["content-type", "x-amz-date", "authorization", "x-api-key", "x-amz-security-token", "x-amz-user-agent"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }

  integrations = {
    "$default" = {
      lambda_arn = module.lambda_function.lambda_function_arn
    }
  }

  create_api_domain_name = false

  tags = {
    Name = var.api_gateway_name
  }
}
