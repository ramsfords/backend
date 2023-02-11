package test

import (
	"io/ioutil"
	"net/http"
	"testing"

	restaurant "github.com/ramsfords/types_gen/v1"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestGetRestaurants(t *testing.T) {
	// create httpexpect instance
	// req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8090/restaurant", nil)

	// rec := httptest.NewRecorder()
	// res := rec.Result()
	req := getGetClient()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}
func TestCreateRestaurants(t *testing.T) {
	data := restaurant.CreateRestaurantData{
		RestaurantName:            "Himalayen",
		RestaurantWebUrl:          "https://www.himalayen.com",
		RestaurantS3DevUrl:        "",
		RestaurantS3StaticProdUrl: "",
		Address:                   &v1.RestaurantAddress{Street1: "Himalayen", City: "Himalayen", State: "Himalayen", Country: "Himalayen", ZipCode: "77081"},
		PhoneNumber:               "",
		Email:                     "",
		OwnerId:                   "1",
		OpenHours: map[string]*v1.Hours{
			"monday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"tuesday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"wednesday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"thrusday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"friday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"saturday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"sunday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "15:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
		},
	}
	req := getPostClient(&data)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	datas, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(datas) != "ABC" {
		t.Errorf("expected ABC got %v", string(datas))
	}

}

func TestAddStaff(t *testing.T) {
	req := getPostClient("hello")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}
func TestRemoveStaff(t *testing.T) {
	req := getPostClient("hello")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}

func TestGetStaffs(t *testing.T) {
	req := getPostClient("hello")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}
func TestUpdateRestaurant(t *testing.T) {
	req := getPostClient("hello")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}

}
