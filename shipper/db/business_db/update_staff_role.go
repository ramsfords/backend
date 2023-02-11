package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) UpdateStaffRole(ctx context.Context, businessId string, staffEmail string, roles []v1.Role) error {
	newRole, err := attributevalue.MarshalListWithOptions(roles)
	if err != nil {
		return err
	}
	_, err = businessDb.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		UpdateExpression:    aws.String("SET #user.#roles = :roles"),
		ConditionExpression: aws.String(fmt.Sprintf("attribute_exists(%s)", "sk")),
		TableName:           aws.String(businessDb.GetFirstShipperTableName()),
		ExpressionAttributeNames: map[string]string{
			"#user":  "user",
			"#roles": "roles",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":roles": &types.AttributeValueMemberL{Value: newRole},
		},
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + staffEmail},
		},

		ReturnValues: types.ReturnValueUpdatedNew,
	})
	return err
}
