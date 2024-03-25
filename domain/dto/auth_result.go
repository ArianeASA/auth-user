package dto

type AuthResult struct {
	AccessToken *string `json:"access_token" type:"string" sensitive:"true"`
	ExpiresIn   *int64  `json:"expires_in" type:"integer"`
	TokenType   *string `json:"token_type" type:"string"`
}
