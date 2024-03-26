package usecases

import (
	"auth-user/domain"
	"github.com/aws/aws-lambda-go/events"
)

type usecases struct {
	auth domain.Auth
}

type Usecases interface {
	HandleGetToken(credentials domain.Credentials) (events.APIGatewayProxyResponse, error)
	HandleCreateUser(user domain.User) (events.APIGatewayProxyResponse, error)
}

func NewUsecases(auth domain.Auth) Usecases {
	return &usecases{auth: auth}
}
