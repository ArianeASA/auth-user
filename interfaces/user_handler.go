package interfaces

import (
	"auth-user/adapters"
	"auth-user/communs/web"
	"auth-user/domain"
	"auth-user/domain/dto"
	"auth-user/usecases"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"strings"
)

func Router(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	httpRequest := req.RequestContext.HTTP
	//log.Printf("EVENT : %v", req)
	cases, response := initRouter()
	if response != nil {
		return *response, nil
	}

	if strings.HasSuffix(httpRequest.Path, "/users") && httpRequest.Method == "POST" {
		var user dto.User
		proxyResponse := getObject[dto.User](req, &user)
		if proxyResponse.StatusCode != 0 {
			return proxyResponse, nil
		}
		return cases.HandleCreateUser(domain.UserToDomain(user))
	}

	if strings.HasSuffix(httpRequest.Path, "/users/token") && httpRequest.Method == "POST" {
		var credential dto.Credentials
		proxyResponse := getObject[dto.Credentials](req, &credential)
		if proxyResponse.StatusCode != 0 {
			return proxyResponse, nil
		}
		return cases.HandleGetToken(domain.CredentialsToDomain(credential))
	}

	log.Println(fmt.Sprintf("endpoint não encontrado %s", httpRequest.Path))
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
		Headers:    web.Headers(),
	}, nil
}

func initRouter() (usecases.Usecases, *events.APIGatewayProxyResponse) {
	authClient, err := adapters.NewAuthClient()
	if err != nil {
		return nil, &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       web.GetJson(web.NewError(err.Error())),
			Headers:    web.Headers(),
		}
	}

	cases := usecases.NewUsecases(authClient)
	return cases, nil
}

func getObject[T any](req events.APIGatewayV2HTTPRequest, object *T) events.APIGatewayProxyResponse {
	err := json.Unmarshal([]byte(req.Body), object)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    web.Headers(),
			Body:       web.GetJson(web.NewError(err.Error())),
		}
	}
	return events.APIGatewayProxyResponse{}
}
