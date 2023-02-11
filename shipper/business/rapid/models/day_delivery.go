package models

type DayDelivery struct {
	FormatedGuaranteedDate      string       `json:"formatedGuaranteedDate" dynamodbav:"formatedGuaranteedDate"`
	GuarantedDate               string       `json:"guarantedDate" dynamodbav:"guarantedDate"`
	IsHoliday                   bool         `json:"isHoliday" dynamodbav:"isHoliday"`
	IsWeekend                   bool         `json:"isWeekend" dynamodbav:"isWeekend"`
	HolidayName                 string       `json:"holidayName" dynamodbav:"holidayName"`
	IsUnspecifiedTransitCarrier bool         `json:"isUnspecifiedTransitCarrier" dynamodbav:"isUnspecifiedTransitCarrier"`
	DayDeliveryFormated         FormatedDate `json:"dayDeliveryFormated" dynamodbav:"dayDeliveryFormated"`
	TransitDays                 int          `json:"transitDays" dynamodbav:"transitDays"`
	Standart                    []Standard   `json:"standart" dynamodbav:"standart"`
}
