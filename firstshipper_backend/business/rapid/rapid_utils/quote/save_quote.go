package quote

import "github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"

func SaveQuoteStep2(quoteDetails *models.QuoteDetails, quoteRate *models.QuoteRate) *models.SaveQuote {
	saveQt := &models.SaveQuote{
		ConfirmAndDispatch:            nil,
		InfoMessage:                   nil,
		IsFromSavedQuote:              false,
		IsKeepAdminChangesModeEnabled: false,
		IsSystemAdmin:                 false,
		OrderID:                       nil,
		PickupDate:                    quoteDetails.OriginShippingDetails.PickupDate,
		QuoteDetails:                  quoteDetails,
		QuoteErrors:                   []string{},
		QuoteRate:                     quoteRate,
		Step:                          2,
	}
	return saveQt
}
