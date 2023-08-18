provider "aws" {
  region = var.region
}

resource "random_string" "suffix" {
  length  = 6
  special = false
  upper   = false
}

module "s3_bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "3.14.1"

  bucket = "${var.s3_bucket}-${var.region}-${random_string.suffix.result}"
}

resource "aws_s3_object" "lambda_handler" {
  bucket = module.s3_bucket.s3_bucket_id
  key    = "todo-list-lambda-handler.zip"
  source = "todo-list-lambda-handler.zip"
}

module "lambda_function" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "5.3.0"

  function_name = var.lambda_function_name
  description   = var.lambda_function_description
  handler       = "bootstrap"
  runtime       = "go1.x"

  create_package = false
  s3_existing_package = {
    bucket = module.s3_bucket.s3_bucket_id
    key    = "todo-list-lambda-handler.zip"
  }

  create_unqualified_alias_allowed_triggers = true

  allowed_triggers = {
    APIGatewayAny = {
      service    = "apigateway"
      source_arn = "${module.api_gateway.apigatewayv2_api_execution_arn}/*/*/*"
    }
  }

  create_current_version_allowed_triggers = false

  environment_variables = {
    "NEW_RELIC_LICENSE_KEY" = var.NEW_RELIC_LICENSE_KEY
  }

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
    allow_headers = ["*"]
    allow_methods = ["*"]
    allow_origins = ["*"]
  }

  integrations = {
    "ANY /{proxy+}" = {
      lambda_arn             = module.lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds   = 10000
    }
  }

  create_api_domain_name = false

  tags = {
    Name = var.api_gateway_name
  }
}
