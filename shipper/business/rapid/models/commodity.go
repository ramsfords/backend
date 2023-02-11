package models

type HazmatDetailInfo struct {
	UnCode             *string     `json:"unCode"  dynamodbav:"unCode"`
	UnNumber           *string     `json:"unNumber"  dynamodbav:"unNumber"`
	ProperShippingName *string     `json:"properShippingName"  dynamodbav:"properShippingName"`
	HazardClass        interface{} `json:"hazardClass"  dynamodbav:"hazardClass"`
	PackingGroup       interface{} `json:"packingGroup"  dynamodbav:"packingGroup"`
	PackingType        *string     `json:"packingType"  dynamodbav:"packingType"`
	ClassCode          interface{} `json:"classCode"  dynamodbav:"packingType"`
	PackingGroupCode   interface{} `json:"packingGroupCode"  dynamodbav:"packingGroupCode"`
}
type Commodity struct {
	Description         *string          `json:"description"  dynamodbav:"description"`
	Nmfc                interface{}      `json:"nmfc"  dynamodbav:"nmfc"`
	NmfcSub             interface{}      `json:"nmfcSub"  dynamodbav:"nmfcSub"`
	CommodityClass      *string          `json:"commodityClass"  dynamodbav:"commodityClass"`
	Pieces              int              `json:"pieces"  dynamodbav:"pieces"`
	TotalWeight         *string          `json:"totalWeight"  dynamodbav:"totalWeight"`
	CurrentTotalWeight  *string          `json:"currentTotalWeight"  dynamodbav:"currentTotalWeight"`
	IsHazardous         bool             `json:"isHazardous"  dynamodbav:"isHazardous"`
	HazmatDetailInfo    HazmatDetailInfo `json:"hazmatDetailInfo"  dynamodbav:"hazmatDetailInfo"`
	CustomerOrderNumber interface{}      `json:"customerOrderNumber"  dynamodbav:"customerOrderNumber"`
	AdditionalInfo      interface{}      `json:"additionalInfo"  dynamodbav:"additionalInfo"`
}
