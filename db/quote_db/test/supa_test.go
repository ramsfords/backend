package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ramsfords/backend/configs"
	v1 "github.com/ramsfords/types_gen/v1"
	supa "github.com/surendrakandel/supa-go"
)

type Quote struct {
	Id             int64            `json:"id"`
	QuoteId        string           `json:"quoteId"`
	Quote          *v1.QuoteRequest `json:"quote"`
	OrganizationId string           `json:"organizationId"`
}
type Organization struct {
	Id             int64       `json:"id"`
	OrganizationId string      `json:"organizationId"`
	Organization   v1.Business `json:"organization"`
}

func TestSupa(t *testing.T) {
	conf := configs.GetConfig()
	supaClient := supa.CreateClient(conf.GetSupaConfig().Url, conf.GetSupaConfig().AnonKey)
	for i := 100; i < 1010; i++ {

		// org := Organization{
		// 	Id:             123,
		// 	OrganizationId: "123",
		// 	Organization:   v1.Business{},
		// }
		// var resultOrg Organization
		// err := supaClient.DB.From("organizations").Insert(org).Execute(resultOrg)
		// fmt.Println("err", err)
		qt := Quote{
			Id:             123,
			QuoteId:        "12355" + fmt.Sprintf("%d", i),
			Quote:          &v1.QuoteRequest{},
			OrganizationId: "123",
		}
		var result []Quote
		err := supaClient.DB.From("quotes").Insert(qt).Execute(&result)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("done")
		time.Sleep(5 * time.Second)
	}
}
