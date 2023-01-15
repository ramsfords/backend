package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

func NewSaveQuoteStep2(quoteDetails *models.QuoteDetails, quoteRate *models.QuoteRate) *models.SaveQuote {
	saveQt := &models.SaveQuote{
		QuoteDetails: quoteDetails,
		QuoteRate:    quoteRate,
		Step:         2,
		QuoteErrors:  []string{},
		PickupDate:   quoteDetails.OriginShippingDetails.PickupDate,
	}
	return saveQt
}
func NewSaveQuoteStep3(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote, serv services.Services) (models.SaveQuote, error) {
	saveQuote = UpdateQuoteDetails(quoteReq, saveQuote)
	saveQuote, err := NewConfirmAndDispatchStep3(quoteReq, saveQuote)
	if err != nil {
		serv.Logger.Error("Error in NewConfirmAndDispatchStep3", err)
		return *saveQuote, err
	}
	saveQuote.Step = 3
	return *saveQuote, nil
}
