package restaurant_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db RestaurantDb) GetRestaurantData(ctx context.Context, restaurantPk string) (*v1.CreateRestaurantData, error) {
	res, err := db.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
			"sk": &types.AttributeValueMemberS{Value: "restaurant#" + restaurantPk},
		},
	})
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	data := &v1.CreateRestaurantData{}
	dbItem, ok := res.Item["restaurant"]
	if ok {
		v, ok := dbItem.(*types.AttributeValueMemberM)
		if ok {
			value := v.Value
			err = attributevalue.UnmarshalMap(value, data)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
	return data, nil
}
