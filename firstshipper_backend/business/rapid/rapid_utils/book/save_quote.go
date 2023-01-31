package book

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
)

func SaveQuoteStep3(quoteRequest *model.QuoteRequest) error {
	quoteRequest.RapidSaveQuote.QuoteErrors = []string{}
	quoteRequest.RapidSaveQuote.Step = 3
	err := makeQuoteDetails(quoteRequest)
	if err != nil {
		return err
	}
	err = newConfirmAndDispatch(quoteRequest)
	if err != nil {
		return err
	}
	return nil
}
