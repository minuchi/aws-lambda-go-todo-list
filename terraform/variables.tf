variable "region" {
  description = "AWS region"
  default     = "ap-northeast-2"
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
