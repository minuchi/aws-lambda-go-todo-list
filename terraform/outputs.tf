output "api_gateway_domain_name" {
  value = module.api_gateway.apigatewayv2_api_api_endpoint
}

output "s3_bucket_id" {
  value = module.s3_bucket.s3_bucket_id
}
