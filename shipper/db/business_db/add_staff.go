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

func (businessDb BusinessDb) SaveStaff(ctx context.Context, businessId string, staff v1.User) error {
	// marshalledStaffList, err := attributevalue.MarshalList([]v1.User{staff})
	// if err != nil {
	// 	return err
	// }
	marshalledStaff, err := attributevalue.MarshalMap(staff)
	if err != nil {
		return err
	}
	addStaffInput := &dynamodb.PutItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":      &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk":      &types.AttributeValueMemberS{Value: "user#" + staff.Email},
			"user_pk": &types.AttributeValueMemberS{Value: "user"},
			"user_sk": &types.AttributeValueMemberS{Value: staff.Email},
			"user":    &types.AttributeValueMemberM{Value: marshalledStaff},
		},
	}
	// updateTransaction := &dynamodb.TransactWriteItemsInput{
	// 	TransactItems: []types.TransactWriteItem{
	// 		{
	// 			Update: &types.Update{
	// 				TableName: aws.String(bis.tableName),
	// 				Key: map[string]types.AttributeValue{
	// 					"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 					"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 				},
	// 				UpdateExpression: aws.String("SET #business.#staffs = list_append(if_not_exists(#business.#staffs, :empty_list), :staffs)"),
	// 				ExpressionAttributeNames: map[string]string{
	// 					"#business": "business",
	// 					"#staffs":   "staffs",
	// 				},
	// 				ExpressionAttributeValues: map[string]types.AttributeValue{
	// 					":staffs": &types.AttributeValueMemberL{Value: marshalledStaffList},
	// 					":empty_list": &types.AttributeValueMemberL{
	// 						Value: []types.AttributeValue{},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Put: &types.Put{
	// 				TableName: aws.String(bis.tableName),
	// 				Item: map[string]types.AttributeValue{
	// 					"pk":      &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 					"sk":      &types.AttributeValueMemberS{Value: "user#" + staff.Email},
	// 					"user_pk": &types.AttributeValueMemberS{Value: "user"},
	// 					"user_sk": &types.AttributeValueMemberS{Value: staff.Email},
	// 					"user":    &types.AttributeValueMemberM{Value: marshalledStaff},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// _, err = bis.TransactWriteItems(ctx, updateTransaction)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	_, err = businessDb.Client.PutItem(ctx, addStaffInput)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
