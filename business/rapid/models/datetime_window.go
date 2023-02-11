package models

type DateTimeWindow struct {
	StartTime string `json:"startTime" dynamodbav:"startTime"`
	EndTime   string `json:"endTime" dynamodbav:"endTime"`
	Date      string `json:"date" dynamodbav:"date"`
}
