package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/menuloom_backend/core/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db MenuDb) GetMenu(ctx context.Context, id string) (*models.Menu, error) {
	res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(db.GetMenuloomTableName()),
		KeyConditionExpression: aws.String("#pk = :pk"),
		ExpressionAttributeNames: map[string]string{
			"#pk": "pk",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + id},
		},
	})
	if err != nil {
		return nil, err
	}
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
	restaurantData := &models.Menu{
		Menu: categories,
	}
	return restaurantData, nil
}
func itemHasCategory(item *v1.Item, category string) bool {
	for _, itemCategory := range item.Categories {
		if itemCategory == category {
			return true
		}
	}
	return false
}
