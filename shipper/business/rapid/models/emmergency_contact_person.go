package models

type EmergencyContactPerson struct {
	Name  string `json:"name,omitempty" dynamodbav:"name"`
	Phone string `json:"phone,omitempty" dynamodbav:"phone"`
}
