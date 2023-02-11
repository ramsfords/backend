package book

import (
	"github.com/ramsfords/backend/shipper/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func DispatchModel(quoteRequest *model.QuoteRequest, bid *v1.Bid) error {
	quoteRequest.RapidSaveQuote.QuoteErrors = []string{}
	quoteRequest.RapidSaveQuote.Step = 3
	err := newConfirmAndDispatch(quoteRequest, bid)
	if err != nil {
		return err
	}
	return nil
}
