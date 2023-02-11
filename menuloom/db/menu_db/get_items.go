package menu_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db MenuDb) GetItems(ctx context.Context, restaurantPk string) ([]*v1.Item, error) {
	res, err := db.Client.Query(ctx, &dynamodb.QueryInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		ExpressionAttributeNames: map[string]string{
			"#pk": "pk",
			"#sk": "sk",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":         &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
			":begin_with": &types.AttributeValueMemberS{Value: "items"},
		},
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :begin_with)"),
	})
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	if err != nil {
		fmt.Println("no items")
	}
	itemsCollection := []*v1.Item{}
	for _, items := range res.Items {
		datas, ok := items["items"]
		if ok {
			data := v1.Item{}
			v, ok := datas.(*types.AttributeValueMemberM)
			if ok {
				value := v.Value
				err = attributevalue.UnmarshalMap(value, &data)
				if err == nil {
					itemsCollection = append(itemsCollection, &data)
				}
				continue
			}
		} else {
			continue
		}
	}
	return itemsCollection, err
}
