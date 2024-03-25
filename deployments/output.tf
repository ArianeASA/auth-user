output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}

output "function_name" {
  description = "Name of the Lambda function."

  value = aws_lambda_function.auth_user_lambda.function_name
}


output "base_url" {
  description = "Base URL for API Gateway stage "

  value = "https://${aws_apigatewayv2_api.auth_user_api.id}.execute-api.${var.aws_region}.amazonaws.com/dev"
}


#output "uri_name_invoke" {
#  description = "Uri Lambda function."
#  value = aws_lambda_function.auth_user_lambda.invoke_arn
#}
#
#output "resor_arn" {
#  description = "Permission"
#  value = aws_lambda_permission.apigw_lambda_token.source_arn
#}
#
#output "cog_client_id" {
#  description = "Client ID do Cog"
#  value = aws_cognito_user_pool_client.cognito_user_pool_client.id
#}