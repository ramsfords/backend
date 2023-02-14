package book

import (
	"github.com/ramsfords/backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func SaveQuoteStep3(quoteRequest *model.QuoteRequest, bid *v1.Bid) error {
	quoteRequest.RapidSaveQuote.QuoteErrors = []string{}
	quoteRequest.RapidSaveQuote.Step = 3
	err := makeQuoteDetails(quoteRequest, bid)
	if err != nil {
		return err
	}
	err = newConfirmAndDispatch(quoteRequest, bid)
	if err != nil {
		return err
	}
	return nil
}
