package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (businessDb BusinessDb) UpdateAllowBooking(ctx context.Context, businessId string, allow bool) (bool, error) {
	qryInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#business":     "business",
			"#validBooking": "allowBooking",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":allowBooking": &types.AttributeValueMemberBOOL{Value: allow},
		},
		UpdateExpression: aws.String("SET #business.#validBooking = :allowBooking"),
		ReturnValues:     types.ReturnValueAllNew,
	}
	_, err := businessDb.Client.UpdateItem(ctx, qryInput)
	if err != nil {
		fmt.Println(err.Error())
		return allow, err
	}
	return allow, nil
}
