package models

import (
	v1 "github.com/ramsfords/types_gen/v1"
)

type Restaurant struct {
	Categories []v1.Category             `json:"categories"`
	Items      []v1.Item                 `json:"items"`
	Users      []v1.User                 `json:"users"`
	Restaurant []v1.CreateRestaurantData `json:"restaurant"`
}
type Menu struct {
	Menu []*v1.Category `json:"categories"`
}
