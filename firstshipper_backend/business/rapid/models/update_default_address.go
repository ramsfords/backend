package models

type UpdateDefaultsAddress struct {
	IsCompanyNameSelected      bool `json:"isCompanyNameSelected" dynamodbav:"isCompanyNameSelected"`
	IsStreetLine1Selected      bool `json:"isStreetLine1Selected" dynamodbav:"isStreetLine1Selected"`
	IsStreetLine2Selected      bool `json:"isStreetLine2Selected" dynamodbav:"isStreetLine2Selected"`
	IsDeliveryFromTimeSelected bool `json:"isDeliveryFromTimeSelected" dynamodbav:"isDeliveryFromTimeSelected"`
	IsDeliveryToTimeSelected   bool `json:"isDeliveryToTimeSelected" dynamodbav:"isDeliveryToTimeSelected"`
	IsLastNameSelected         bool `json:"isLastNameSelected" dynamodbav:"isLastNameSelected"`
	IsFirstNameSelected        bool `json:"isFirstNameSelected" dynamodbav:"isFirstNameSelected"`
	IsEmailSelected            bool `json:"isEmailSelected" dynamodbav:"isEmailSelected"`
	IsPhoneSelected            bool `json:"isPhoneSelected" dynamodbav:"isPhoneSelected"`
	IsExtSelected              bool `json:"isExtSelected" dynamodbav:"isExtSelected"`
	IsStartTimeSelected        bool `json:"isStartTimeSelected" dynamodbav:"isStartTimeSelected"`
	IsEndTimeSelected          bool `json:"isEndTimeSelected" dynamodbav:"isEndTimeSelected"`
}
