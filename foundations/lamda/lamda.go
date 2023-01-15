package lamda

import (
	"context"
	"log"

	awsconf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/pkg/errors"
)

func New() *lambda.Client {

	awsConf, err := awsconf.LoadDefaultConfig(context.Background(), awsconf.WithRegion("us-west-1"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not start lamda service"))
	}
	return lambda.NewFromConfig(awsConf)
}
