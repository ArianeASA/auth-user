package usecases

import (
	"auth-user/communs/web"
	"auth-user/domain"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func (use *usecases) HandleGetToken(credentials domain.Credentials) (events.APIGatewayProxyResponse, error) {
	token, err := use.auth.NewToken(credentials)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       web.GetJson(web.NewError(err.Error())),
			Headers:    web.Headers(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       web.GetJson(token.DomainToResponse()),
		Headers:    web.Headers(),
	}, nil

}
