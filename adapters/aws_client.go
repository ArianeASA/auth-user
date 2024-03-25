package adapters

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func NewAwsClient() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionAws)},
	)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	return sess, nil

}
