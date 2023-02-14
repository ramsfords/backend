package models

type TitleHeadData struct {
	AccessorialServices []interface{} `json:"accessorialServices" dynamodbav:"accessorialServices"`
	OriginZip           string        `json:"originZip" dynamodbav:"originZip"`
	DestinationZip      string        `json:"destinationZip" dynamodbav:"destinationZip"`
	Weight              int           `json:"weight" dynamodbav:"weight"`
	QuoteKey            string        `json:"quoteKey" dynamodbav:"quoteKey"`
	Classes             []string      `json:"classes" dynamodbav:"classes"`
	PickUpDate          string        `json:"pickUpDate" dynamodbav:"pickUpDate"`
	FormatedPickUpDate  FormatedDate  `json:"formatedPickUpDate" dynamodbav:"formatedPickUpDate"`
}
