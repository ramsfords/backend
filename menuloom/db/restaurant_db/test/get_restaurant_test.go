package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/menuloom/core/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "KeyConditionExpression": "#69240 = :69240",
// "ExpressionAttributeNames": {"#69240":"pk"},
// "ExpressionAttributeValues": {":69240": {"S":"pk#127.0.0.1:3000"}}

func TestGetRestaurant(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
			TableName:              aws.String(conf.GetMenuloomTableName()),
			KeyConditionExpression: aws.String("#pk = :pk"),
			ExpressionAttributeNames: map[string]string{
				"#pk": "pk",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: "pk#himalayen.menuloom.com"},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		categories := []v1.Category{}
		itemsCollection := []v1.Item{}
		users := []v1.User{}
		restaurants := []v1.CreateRestaurantData{}
		itemCounters := 0
		for _, items := range res.Items {
			categorykey := items["categories"]
			itemKey := items["items"]
			restaurantkey := items["restaurant"]
			usersKey := items["users"]
			if categorykey != nil {
				itemCounters += 1
				category := v1.Category{}
				values, ok := categorykey.(*types.AttributeValueMemberM)
				if ok {
					err = attributevalue.Unmarshal(values, &category)
					if err == nil {
						categories = append(categories, category)
					}
				}
				continue
			}
			if itemKey != nil {
				itemCounters += 1

				values, ok := itemKey.(*types.AttributeValueMemberM)
				if ok {
					itemTemp := v1.Item{}
					err = attributevalue.Unmarshal(values, &itemTemp)
					if err == nil {
						itemsCollection = append(itemsCollection, itemTemp)
					}
				}
				continue
			}
			if restaurantkey != nil {
				itemCounters += 1
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
				itemCounters += 1
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
			// fmt.Println(categorykey, itemKey, restaurantkey, usersKey)
			// for _, v := range item {
			// 	_, ok := v.(*types.AttributeValueMemberS)
			// 	if !ok {
			// 		itemCounters += 1
			// 		newV, ok := v.(*types.AttributeValueMemberM)
			// 		if ok {
			// 			newValue := newV.Value
			// 			for _, newvVs := range newValue {
			// 				fmt.Println(newvVs)
			// 			}
			// 		}
			// 		fmt.Println(v)
			// 	}
			// if ok && value.Value == "items" || value.Value == "categories" || value.Value == "restaurant" || value.Value == "users" {
			// 	switch value.Value {
			// 	case "categories":
			// 		categoriesNumber += 1
			// 		fmt.Println("categories")
			// 	case "items":
			// 		itemsNumber += 1
			// 		fmt.Println("items")
			// 	case "restaurant":
			// 		restaurantNumber += 1
			// 		fmt.Println("restaurant")
			// 	case "users":
			// 		usersNumber += 1
			// 		fmt.Println("users")
			// 	}
			// 	continue
			// } else {
			// 	continue
			// }

		}
		fmt.Println(
			"itemCounters: ", itemCounters,
		)
		restaurantData := &models.Restaurant{
			Categories: categories,
			Items:      itemsCollection,
			Users:      users,
			Restaurant: restaurants,
		}
		fmt.Print(restaurantData)
		// restaurantData := &v1.CreateRestaurantData{}
		// restaurant, ok := res.Item["restaurant"]
		// if !ok {
		// 	t.Fatal(err)
		// }
		// err = attributevalue.Unmarshal(restaurant, restaurantData)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// fmt.Println(res, err)
	})
}
