package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/menuloom_backend/core/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db RestaurantDb) GetRestaurant(ctx context.Context, id string) (*models.Restaurant, error) {
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

	categories := []v1.Category{}
	itemsCollection := []v1.Item{}
	users := []v1.User{}
	restaurants := []v1.CreateRestaurantData{}
	for _, items := range res.Items {
		categorykey := items["categories"]
		itemKey := items["items"]
		restaurantkey := items["restaurant"]
		usersKey := items["users"]
		if categorykey != nil {
			values, ok := categorykey.(*types.AttributeValueMemberM)
			if ok {
				category := v1.Category{}
				err = attributevalue.Unmarshal(values, &category)
				if err == nil {
					categories = append(categories, category)
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
					itemsCollection = append(itemsCollection, item)
				}
			}
			continue
		}
		if restaurantkey != nil {
			restaurant := v1.CreateRestaurantData{}
			values, ok := restaurantkey.(*types.AttributeValueMemberM)
			if ok {
				err = attributevalue.Unmarshal(values, &restaurant)
				if err == nil {
					restaurants = append(restaurants, restaurant)
				}
			}
			continue
		}
		if usersKey != nil {
			values, ok := usersKey.(*types.AttributeValueMemberM)
			if ok {
				user := v1.User{}
				err = attributevalue.Unmarshal(values, &user)
				if err == nil {
					users = append(users, user)
				}
			}
			continue
		}
	}

	restaurantData := &models.Restaurant{
		Categories: categories,
		Items:      itemsCollection,
		Users:      users,
		Restaurant: restaurants,
	}
	return restaurantData, nil
}
