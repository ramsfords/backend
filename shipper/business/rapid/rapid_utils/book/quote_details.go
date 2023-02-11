package book

import (
	"github.com/ramsfords/backend/shipper/business/core/model"
	"github.com/ramsfords/backend/shipper/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func makeQuoteDetails(quoteRequest *model.QuoteRequest, bid *v1.Bid) error {
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
		ReferenceNumber: &bid.BidId,
		PoNumber:        &poNumber,
		PickupNumber:    &poNumber,
	}
	return nil
}
