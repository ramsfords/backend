package models

type Dimensions struct {
	Length        string `json:"length,omitempty" dynamodbav:"length"`
	Width         string `json:"width,omitempty" dynamodbav:"width"`
	Height        string `json:"height,omitempty" dynamodbav:"height"`
	CurrentLength string `json:"currentLength,omitempty" dynamodbav:"currentLength"`
	CurrentWidth  string `json:"currentWidth,omitempty" dynamodbav:"currentWidth"`
	CurrentHeight string `json:"currentHeight,omitempty" dynamodbav:"currentHeight"`
}
