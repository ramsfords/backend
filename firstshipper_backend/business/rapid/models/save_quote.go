package models

type SaveQuote struct {
	QuoteDetails                  *QuoteDetails       `json:"quoteDetails" dynamodbav:"quoteDetails"`
	ConfirmAndDispatch            *ConfirmAndDispatch `json:"confirmAndDispatch" dynamodbav:"confirmAndDispatch"`
	QuoteRate                     *QuoteRate          `json:"quoteRate" dynamodbav:"quoteRate"`
	SavedQuoteID                  string              `json:"savedQuoteId" dynamodbav:"savedQuoteId"`
	ReferenceID                   string              `json:"referenceId" dynamodbav:"referenceId"`
	IsFromSavedQuote              bool                `json:"isFromSavedQuote" dynamodbav:"isFromSavedQuote"`
	PickupDate                    string              `json:"pickupDate" dynamodbav:"pickupDate"`
	InfoMessage                   *string             `json:"infoMessage" dynamodbav:"infoMessage"`
	IsFavorite                    bool                `json:"isFavorite" dynamodbav:"isFavorite"`
	IsSystemAdmin                 bool                `json:"isSystemAdmin" dynamodbav:"isSystemAdmin"`
	IsKeepAdminChangesModeEnabled bool                `json:"isKeepAdminChangesModeEnabled" dynamodbav:"isKeepAdminChangesModeEnabled"`
	OrderID                       *string             `json:"orderId" dynamodbav:"orderId"`
	QuoteErrors                   []string            `json:"quoteErrors" dynamodbav:"quoteErrors"`
	Step                          int                 `json:"step" dynamodbav:"step"`
	GfpTotals                     interface{}         `json:"gfpTotals"`
	GfpPackageType                []interface{}       `json:"gfpPackageType"`
}
