package book

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
)

func makeQuoteDetails(quoteRequest *model.QuoteRequest) error {
	err := originShippingDetails(quoteRequest)
	if err != nil {
		return err
	}
	err = destinationShippingDetails(quoteRequest)
	if err != nil {
		return err
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.FreightCharge = 1

	bol := "BOL" + quoteRequest.QuoteRequest.QuoteId
	poNumber := "PO" + quoteRequest.QuoteRequest.QuoteId
	quoteRequest.RapidSaveQuote.QuoteDetails.ReferenceNumberInfo = &models.ReferenceNumberInfo{
		CustomerBOL:     &bol,
		ReferenceNumber: &quoteRequest.Bids[0].QuoteId,
		PoNumber:        &poNumber,
		PickupNumber:    &poNumber,
	}
	return nil
}
