package rapid_db

import (
	"context"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
)

func (rapiddb RapidDb) GetRapidQuote(ctx context.Context, quoteId string) (models.QuoteRate, error) {
	return models.QuoteRate{}, nil
}
