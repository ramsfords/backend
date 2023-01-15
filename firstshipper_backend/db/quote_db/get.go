package quote_db

import (
	"context"

	v1 "github.com/ramsfords/types_gen/v1"
)

func (quotedb QuoteDb) GetQuote(ctx context.Context, QuoteId string) (v1.QuoteRequest, error) {
	return v1.QuoteRequest{}, nil
}
