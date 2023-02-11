package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) AddContact(ctx context.Context, businessId string, contact *v1.Location) error {
	marshalledContact, err := attributevalue.Marshal(contact)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(businessDb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":      &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk":      &types.AttributeValueMemberS{Value: "contact#" + contact.Contact.EmailAddress},
			"contact": marshalledContact,
		},
		ConditionExpression: aws.String("attribute_not_exists(sk)"),
		ReturnValues:        types.ReturnValueAllNew,
	}
	_, err = businessDb.Client.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
