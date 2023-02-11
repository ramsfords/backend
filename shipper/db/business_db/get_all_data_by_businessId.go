package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/shipper/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) GetAllDataByBusinessId(ctx context.Context, businessId string) (*model.BusinessData, error) {
	bisInput := &dynamodb.QueryInput{
		TableName:              aws.String(businessDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("pk = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
		},
	}
	res, err := businessDb.Client.Query(ctx, bisInput)
	if err != nil {
		return nil, err
	}
	if len(res.Items) == 0 {
		return nil, err
	}
	data := &model.BusinessData{}
	for _, item := range res.Items {
		fmt.Println(item)
		businessData, ok := item["business"]
		if ok {
			err = attributevalue.Unmarshal(businessData, &data.Business)
			if err != nil {
				fmt.Println(businessData)
			}
			fmt.Println(businessData)
		}
		userData, ok := item["users"]
		if ok {
			user := &v1.User{}
			err = attributevalue.Unmarshal(userData, user)
			if err != nil {
				fmt.Println(businessData)
			}
			data.Users = append(data.Users, user)
			fmt.Println(businessData)
		}
		quotesData, ok := item["quotes"]
		if ok {
			quote := &model.QuoteRequest{}
			err = attributevalue.Unmarshal(quotesData, quote)
			if err != nil {
				fmt.Println(businessData)
			}
			data.QuoteRequests = append(data.QuoteRequests, quote)
			fmt.Println(businessData)
		}
		fmt.Println(userData)
	}

	return data, nil
}
