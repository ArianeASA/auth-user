
resource "aws_cognito_user_pool" "cognito_user_pool" {
  name = "auth-user-pool"
  auto_verified_attributes = ["email"]
  email_verification_message = "Bem-vindo ao nosso serviço User AUTH. Para verificar sua conta, clique no link abaixo.\n {####}"
  email_verification_subject = "Verifique sua conta USER AUTH"

  verification_message_template {
    default_email_option  = "CONFIRM_WITH_LINK"
  }

  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

}
resource "aws_cognito_user_pool_domain" "user_pool_domain" {
  domain      = "auth-user-x-domain"
  user_pool_id = aws_cognito_user_pool.cognito_user_pool.id
}

resource "aws_cognito_user_pool_client" "cognito_user_pool_client" {
  name         = "auth-user-x-pool-client"
  user_pool_id = aws_cognito_user_pool.cognito_user_pool.id

  explicit_auth_flows = ["ALLOW_REFRESH_TOKEN_AUTH", "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_USER_SRP_AUTH"]
}

