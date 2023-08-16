variable "region" {
  description = "AWS region"
  default     = "ap-northeast-2"
}

variable "s3_bucket" {
  description = "S3 bucket name to store lambda code"
  default     = "lambda-builds"
}

variable "lambda_function_name" {
  description = "Lambda function name"
  default     = "todo-list"
}

variable "lambda_function_description" {
  description = "Lambda function description"
  default     = "A To-Do List"
}

variable "api_gateway_name" {
  description = "Lambda function name"
  default     = "todo-list"
}

variable "api_gateway_description" {
  description = "Lambda function description"
  default     = "A To-Do List"
}

variable "NEW_RELIC_LICENSE_KEY" {
  description = "New Relic license key"
  default     = ""
}
