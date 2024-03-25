resource "aws_apigatewayv2_api" "auth_user_api" {
  name        = "auth-user-api"
  description = "Auth User API Gateway"
  body = templatefile("./openapi.json",
    {
      quote_receiver =  aws_lambda_function.auth_user_lambda.invoke_arn
    }
  )

  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.auth_user_api.id

  name        = "dev"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.fiap_food_log_group.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
    }
    )
  }
}

resource "aws_apigatewayv2_integration"  "lambda_integration" {
  api_id = aws_apigatewayv2_api.auth_user_api.id

  integration_uri    = aws_lambda_function.auth_user_lambda.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

