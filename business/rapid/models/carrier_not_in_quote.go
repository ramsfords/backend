package models

type CarrierNotInQuote struct {
	Icon              string   `json:"icon,omitempty"  dynamodbav:"icon"`
	Name              string   `json:"name,omitempty"  dynamodbav:"name"`
	Code              string   `json:"code,omitempty"  dynamodbav:"code"`
	Messages          []string `json:"messages,omitempty"  dynamodbav:"messages"`
	DisplayAPIWarning bool     `json:"displayAPIWarning,omitempty"  dynamodbav:"displayAPIWarning"`
	APIOutageMessage  string   `json:"apiOutageMessage,omitempty"  dynamodbav:"apiOutageMessage"`
}
