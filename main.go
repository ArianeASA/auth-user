package main

import (
	"auth-user/interfaces"
	"github.com/aws/aws-lambda-go/lambda"
)
// Compare this
func main() {
	lambda.Start(interfaces.Router)
}
