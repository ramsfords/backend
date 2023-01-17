package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestGetMenu(t *testing.T) {

	res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(conf.GetMenuloomTableName()),
		KeyConditionExpression: aws.String("#pk = :pk"),
		ExpressionAttributeNames: map[string]string{
			"#pk": "pk",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayantasteusa.com"},
		},
	})
	categories := []*v1.Category{}
	itemsCollection := []*v1.Item{}
	for _, items := range res.Items {
		categorykey := items["categories"]
		itemKey := items["items"]
		if categorykey != nil {
			category := v1.Category{}
			values, ok := categorykey.(*types.AttributeValueMemberM)
			if ok {
				err = attributevalue.Unmarshal(values, &category)
				if err == nil {
					categories = append(categories, &category)
				}
			}
			continue
		}
		if itemKey != nil {
			item := v1.Item{}
			values, ok := itemKey.(*types.AttributeValueMemberM)
			if ok {
				err = attributevalue.Unmarshal(values, &item)
				if err == nil {
					itemsCollection = append(itemsCollection, &item)
				}
			}
			continue
		}
		// if restaurantkey != nil {
		// 	restaurant := proto_gen.CreateRestaurantData{}
		// 	values, ok := restaurantkey.(*types.AttributeValueMemberM)
		// 	if ok {
		// 		err = attributevalue.Unmarshal(values, &restaurant)
		// 		if err == nil {
		// 			restaurants = append(restaurants, &restaurant)
		// 		}
		// 	}
		// 	continue
		// }
	}
	for i, category := range categories {
		for _, item := range itemsCollection {
			if itemHasCategory(item, category.Name) {
				categories[i].Items = append(categories[i].Items, item)
			}
		}
	}
	if err != nil {
		t.Error(err)
	}
	db := GetProdDb()
	for _, category := range categories {
		marshalledMap, err := attributevalue.Marshal(category)
		if err != nil {
			t.Error(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String("menuloom-prod"),
			Item: map[string]types.AttributeValue{
				"pk":         &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayantasteusa.com"},
				"sk":         &types.AttributeValueMemberS{Value: "categories#" + category.Name},
				"categories": marshalledMap,
			}})
		if err != nil {
			t.Error(err)
		}
		fmt.Println(res)
	}
	for _, item := range itemsCollection {
		marshalledMap, err := attributevalue.Marshal(item)
		if err != nil {
			t.Error(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String("menuloom-prod"),
			Item: map[string]types.AttributeValue{
				"pk":         &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayantasteusa.com"},
				"sk":         &types.AttributeValueMemberS{Value: "items#" + item.Name},
				"categories": marshalledMap,
			}})
		if err != nil {
			t.Error(err)
		}
		fmt.Println(res)
	}
	fmt.Println(db)
	fmt.Println(res)
}
func itemHasCategory(item *v1.Item, category string) bool {
	for _, itemCategory := range item.Categories {
		if itemCategory == category {
			return true
		}
	}
	return false
}
