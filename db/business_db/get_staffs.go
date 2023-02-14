package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "firstshipper-dev",
//     "ScanIndexForward": true,
//     "ConsistentRead": false,
//     "KeyConditionExpression": "#87671 = :87671 And begins_with(#87672, :87672)",
//     "ProjectionExpression": "#87670",
//     "ExpressionAttributeValues": {
//       ":87671": {
//         "S": "pk#1cc284"
//       },
//       ":87672": {
//         "S": "user"
//       }
//     },
//     "ExpressionAttributeNames": {
//       "#87670": "users",
//       "#87671": "pk",
//       "#87672": "sk"
//     }

func (businessDb BusinessDb) GetStaffsForABusiness(ctx context.Context, businessId string) ([]*v1.User, error) {
	res, err := businessDb.Client.Query(ctx, &dynamodb.QueryInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		ExpressionAttributeNames: map[string]string{
			"#pk":    "pk",
			"#sk":    "sk",
			"#users": "users",
		},
		ConsistentRead:   aws.Bool(false),
		ScanIndexForward: aws.Bool(false),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "user"},
		},
		ProjectionExpression:   aws.String("#users"),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
	})
	if err != nil {
		return []*v1.User{}, err
	}
	if len(res.Items) == 0 {
		return []*v1.User{}, nil
	}
	usrs := []*v1.User{}
	for _, item := range res.Items {
		data, ok := item["users"]
		if !ok {
			continue
		}
		usr := &v1.User{}
		err = attributevalue.Unmarshal(data, &usr)
		if err != nil {
			return []*v1.User{}, err
		}
		usrs = append(usrs, usr)
	}
	return usrs, nil
}
