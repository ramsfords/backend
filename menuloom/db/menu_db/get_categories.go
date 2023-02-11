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

func (db MenuDb) GetCategories(ctx context.Context, id string) ([]*v1.Category, error) {
	res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		ExpressionAttributeNames: map[string]string{
			"#pk": "pk",
			"#sk": "sk",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":          &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + id},
			":begins_with": &types.AttributeValueMemberS{Value: "categories"},
		},
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :begins_with)"),
	})
	if err != nil {
		return nil, err
	}
	categories := []*v1.Category{}
	for k, v := range res.Items {
		fmt.Println(k, v)
		cat, ok := v["categories"]
		if ok {
			fmt.Println(k, cat)
			data := &v1.Category{}
			err = attributevalue.Unmarshal(cat, data)
			if err != nil {
				fmt.Println(err)
			}
			categories = append(categories, data)
		}

		// cat, ok := v.(*types.AttributeValueMemberM)
		// if ok {
		// 	fmt.Println(k, cat)
		// 	data := &v1.Category{}
		// 	err = attributevalue.Unmarshal(cat, data)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	categories = append(categories, data)
		// }
	}
	return categories, nil
}
