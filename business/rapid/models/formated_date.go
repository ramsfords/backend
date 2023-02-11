package models

type FormatedDate struct {
	FormatedDay     string `json:"formatedDay" dynamodbav:"formatedDay"`
	FormatedWeekDay string `json:"formatedWeekDay" dynamodbav:"formatedWeekDay"`
	FormatedMonth   string `json:"formatedMonth" dynamodbav:"formatedMonth"`
}
