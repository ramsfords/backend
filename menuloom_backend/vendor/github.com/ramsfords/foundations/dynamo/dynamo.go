package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ramsfords/configs"
)

// type Database struct {
// 	*dynamodb.Client
// }

//	type DbKeys struct {
//		GlobleSkName  string
//		GloblePkName  string
//		PkName        string
//		SkName        string
//		GlobleSKValue string
//		GloblePKValue string
//		PKValue       string
//		SKValue       string
//		IndexName     string
//	}
type DB struct {
	Client dynamodb.Client
}

func New(conf configs.Config) DB {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "us-west-1" && conf.Env == "aws-deployment" {
			fmt.Println("hello")
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "https://dynamodb.us-west-1.amazonaws.com",
				SigningRegion: "us-west-1",
			}, nil
		} else if conf.Env == "dev" {
			return aws.Endpoint{
				URL: "http://127.0.0.1:8000",
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	confs := aws.Config{
		Region:                      "us-west-1",
		Credentials:                 &conf,
		EndpointResolverWithOptions: customResolver,
		RetryMaxAttempts:            10,
	}
	client := *dynamodb.NewFromConfig(confs)
	return DB{Client: client}

}

// func (db Database) Getdb() Database {
// 	return db
// }
// func (db Database) GetDynamoClient() *dynamodb.Client {
// 	return db.Client
// }
// func (db Database) StatusCheck(ctx context.Context) bool {
// 	_, err := db.Client.ListTables(ctx, &dynamodb.ListTablesInput{})
// 	return err == nil
// }

// func (rc Database) HashUserId(userName string) string {
// 	md5Hash := sha256.Sum256([]byte(userName))
// 	// table name will always try to keep data in one row for a user with same hasheduserid
// 	nodeName := base64.URLEncoding.EncodeToString(md5Hash[:])
// 	return nodeName
// }
// func (db Database) DbClient() *dynamodb.Client {
// 	return db.Client
// }

// func NewDynamoClient(env string, cred aws.CredentialsProvider) *dynamodb.Client {
// 	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 		if service == dynamodb.ServiceID && region == "us-west-1" && env == "aws-deployment" {
// 			fmt.Println("hello")
// 			return aws.Endpoint{
// 				PartitionID:   "aws",
// 				URL:           "https://dynamodb.us-west-1.amazonaws.com",
// 				SigningRegion: "us-west-1",
// 			}, nil
// 		} else if env == "dev" {
// 			return aws.Endpoint{
// 				URL: "http://0.0.0.0:8000",
// 			}, nil
// 		}
// 		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
// 		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
// 	})
// 	conf := aws.Config{
// 		Region:                      "us-west-1",
// 		Credentials:                 cred,
// 		EndpointResolverWithOptions: customResolver,
// 		RetryMaxAttempts:            10,
// 	}
// 	client := dynamodb.NewFromConfig(conf)

// 	return client

// }
