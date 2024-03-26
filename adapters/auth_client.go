package adapters

import (
	"auth-user/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
	"os"
)

type authExternalClient struct {
	Client *cognitoidentityprovider.CognitoIdentityProvider
}

func NewAuthClient() (domain.Auth, error) {
	sess, err := NewAwsClient()
	if err != nil {
		return nil, err
	}

	cognitoClient := cognitoidentityprovider.New(sess)
	return &authExternalClient{Client: cognitoClient}, nil
}

func (auth *authExternalClient) NewUser(user domain.User) error {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv(clientId)),
		Password: aws.String(user.Password),
		Username: aws.String(user.UserName),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			}, {
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			}, {
				Name:  aws.String("registration_number"),
				Value: aws.String(user.RegistrationNumber),
			},
		},
	}

	result, err := auth.Client.SignUp(input)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result.String())
	return nil

}

func (auth *authExternalClient) NewToken(cred domain.Credentials) (domain.AuthResult, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(cred.UserName),
			"PASSWORD": aws.String(cred.Password),
		},

		ClientId: aws.String(os.Getenv("CLIENT_ID")),
	}

	result, err := auth.Client.InitiateAuth(input)
	if err != nil {
		log.Println(err)
		return domain.AuthResult{}, err
	}

	log.Println(result.AuthenticationResult.AccessToken)
	return domain.AuthResult{
		AccessToken: result.AuthenticationResult.AccessToken,
		ExpiresIn:   result.AuthenticationResult.ExpiresIn,
		TokenType:   result.AuthenticationResult.TokenType,
	}, nil
}
