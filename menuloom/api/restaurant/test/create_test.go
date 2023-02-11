package test

import (
	"encoding/json"
	"fmt"
	"testing"

	v1 "github.com/ramsfords/types_gen/v1"
)

func TestCreate(t *testing.T) {
	data := v1.CreateRestaurantData{
		Name: "Test",
		OpenHours: map[string]*v1.Hours{
			"monday": &v1.Hours{
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "20:00",
					},
					{
						From: "10:00",
						To:   "20:00",
					},
				},
			},
			"tuesday": &v1.Hours{
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "20:00",
					},
					{
						From: "10:00",
						To:   "20:00",
					},
				},
			},
			"wednesday": &v1.Hours{
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "20:00",
					},
					{
						From: "10:00",
						To:   "20:00",
					},
				},
			},
			"trushday": &v1.Hours{
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "20:00",
					},
					{
						From: "10:00",
						To:   "20:00",
					},
				},
			},
		},
	}
	mData, err := json.Marshal(&data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(mData))
}
