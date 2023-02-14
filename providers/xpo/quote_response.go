package xpo

type TotAccessorialAmt struct {
	Amt        float64 `json:"amt"`
	CurrencyCd string  `json:"currencyCd"`
}
type TotCharge struct {
	Amt        float64 `json:"amt"`
	CurrencyCd string  `json:"currencyCd"`
}
type TotDiscountAmt struct {
	Amt        float64 `json:"amt"`
	CurrencyCd string  `json:"currencyCd"`
}
type ChargeAmt struct {
	Amt        float64 `json:"amt"`
	CurrencyCd string  `json:"currencyCd"`
}

type CalcMethod struct {
	PerUnitRate float64 `json:"perUnitRate"`
}
type Charge struct {
	ChargeAmt  `json:"chargeAmt"`
	CalcMethod `json:"calcMethod"`
}

type TotCommodityCharge struct {
	Amt        float64 `json:"amt"`
	CurrencyCd string  `json:"currencyCd"`
}
type TotCommodityWeight struct {
	Weight    float64 `json:"weight"`
	WeightUom string  `json:"weightUom"`
}

type TransitTime struct {
	DestPostalCd      string `json:"destPostalCd"`
	DestSicCd         string `json:"destSicCd"`
	EstdDlvrDate      int64  `json:"estdDlvrDate"`
	GarntInd          bool   `json:"garntInd"`
	OrigPostalCd      string `json:"origPostalCd"`
	OrigSicCd         string `json:"origSicCd"`
	RequestedPkupDate int64  `json:"requestedPkupDate"`
	TransitDays       int    `json:"transitDays"`
	IsPkupDateHoliday bool   `json:"isPkupDateHoliday"`
}
type Data struct {
	RateQuote   RateQuote     `json:"rateQuote"`
	Msgs        []interface{} `json:"msgs"`
	TransitTime TransitTime   `json:"transitTime"`
}
type RateQuote struct {
	ConfirmationNbr              string                       `json:"confirmationNbr"`
	SpotQuoteNbr                 string                       `json:"spotQuoteNbr"`
	QuoteCreateDate              int64                        `json:"quoteCreateDate"`
	QuoteCreatedByCd             string                       `json:"quoteCreatedByCd"`
	AccessorialTariffName        string                       `json:"accessorialTariffName"`
	ActlDiscountPct              float64                      `json:"actlDiscountPct"`
	AmcInd                       bool                         `json:"amcInd"`
	TotAccessorialAmt            TotAccessorialAmt            `json:"totAccessorialAmt"`
	TotCharge                    TotCharge                    `json:"totCharge"`
	TotDiscountAmt               TotDiscountAmt               `json:"totDiscountAmt"`
	XPOQuoteResponseShipmentInfo XPOQuoteResponseShipmentInfo `json:"shipmentInfo"`
	VspApplied                   bool                         `json:"vspApplied"`
	GuarantdEligible             bool                         `json:"guarantdEligible"`
	SqEligible                   bool                         `json:"sqEligible"`
	G12Eligible                  bool                         `json:"g12Eligible"`
	RrsEligible                  bool                         `json:"rrsEligible"`
	SpecialServiceCharges        []interface{}                `json:"specialServiceCharges"`
}
type XPOQuoteResponse struct {
	Code                 string `json:"code"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	Data                 Data   `json:"data"`
}
type XPOQuoteResponseShipmentInfo struct {
	Accessorials       []Accessorial      `json:"accessorials"`
	Commodity          []Commodity        `json:"commodity"`
	Shipper            Shipper            `json:"shipper"`
	Consignee          Consignee          `json:"consignee"`
	PaymentTermCd      string             `json:"paymentTermCd"`
	GarntInd           bool               `json:"garntInd"`
	G12Ind             bool               `json:"g12Ind"`
	RrsInd             bool               `json:"rrsInd"`
	LinealFt           int                `json:"linealFt"`
	ShipmentDate       int64              `json:"shipmentDate"`
	Comment            string             `json:"comment"`
	PalletCnt          int                `json:"palletCnt"`
	TotCommodityCharge TotCommodityCharge `json:"totCommodityCharge"`
	TotCommodityWeight TotCommodityWeight `json:"totCommodityWeight"`
}
type AutoGenerated struct {
	Code                 string `json:"code"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	Data                 Data   `json:"data"`
}
