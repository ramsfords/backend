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

func TestCreateCategories(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		var errs error
		data := v1.Categories{
			Categories: []*v1.Category{
				{
					Name:         "breakfast",
					LocalName:    "local breakfast",
					Description:  "our breakfast is the best breakfast in the world",
					ServingTime:  v1.ServingTime_BREAKFAST,
					RestaurantId: "himalayen.menuloom.com",
					Type:         "category",
					Pk:           "pk#",
					Sk:           "category#",
					Images:       []*v1.Image{},
				},
				{
					Name:         "lunch",
					LocalName:    "local lunch",
					Description:  "our lunch is the best breakfast in the world",
					ServingTime:  v1.ServingTime_BREAKFAST,
					RestaurantId: "himalayen.menuloom.com",
					Type:         "category",
					Pk:           "pk#",
					Sk:           "category#",
					Images:       []*v1.Image{},
				},
			},
		}
		for _, category := range data.Categories {
			data, err := attributevalue.Marshal(&category)
			if err != nil {
				errs = err
			}
			_, err = db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
				TableName: aws.String(conf.GetMenuloomTableName()),
				Item: map[string]types.AttributeValue{
					"pk":         &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayen.menuloom.com"},
					"sk":         &types.AttributeValueMemberS{Value: "category#" + category.Name},
					"categories": data,
				},
				ConditionExpression: aws.String("attribute_not_exists(sk)"),
			})
			if err != nil {
				errs = err
			}

		}
		if errs != nil {
			fmt.Println(errs.Error())
		}

	})

}
