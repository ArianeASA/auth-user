package usecases

import (
	"auth-user/communs/web"
	"auth-user/domain"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func (use *usecases) HandleCreateUser(user domain.User) (events.APIGatewayProxyResponse, error) {

	err := use.auth.NewUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       web.GetJson(web.NewError(err.Error())),
			Headers:    web.Headers(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
		Headers:    web.Headers(),
	}, nil
}
