package main

import (
	"auth-user/interfaces"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(interfaces.Router)
}
