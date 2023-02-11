package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (businessDb BusinessDb) DeleteStaff(ctx context.Context, businessId string, email string) error {
	// staffs, err := bis.GetStaffs(ctx, businessId)
	// if err != nil {
	// 	return err
	// }
	// newStaffs := []v1.User{}
	// for _, i := range staffs {
	// 	if i.Email != email {
	// 		newStaffs = append(newStaffs, *i)
	// 	}
	// }
	// marshalledMembersList, err := attributevalue.MarshalList(newStaffs)
	// if err != nil {
	// 	return err
	// }
	// updateTransaction := &dynamodb.TransactWriteItemsInput{
	// 	TransactItems: []types.TransactWriteItem{
	// 		{
	// 			Delete: &types.Delete{
	// 				TableName: aws.String(bis.tableName),
	// 				Key: map[string]types.AttributeValue{
	// 					"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 					"sk": &types.AttributeValueMemberS{Value: "user#" + email},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Update: &types.Update{
	// 				TableName: aws.String(bis.tableName),
	// 				Key: map[string]types.AttributeValue{
	// 					"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 					"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
	// 				},
	// 				UpdateExpression: aws.String("SET #business.#staffs = :staffs)"),
	// 				ExpressionAttributeNames: map[string]string{
	// 					"#business": "business",
	// 					"#staffs":   "staffs",
	// 				},
	// 				ExpressionAttributeValues: map[string]types.AttributeValue{
	// 					":staffs": &types.AttributeValueMemberL{Value: marshalledMembersList},
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
	_, err := businessDb.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + email},
		},
		TableName:    aws.String(businessDb.GetFirstShipperTableName()),
		ReturnValues: "ALL_OLD",
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
