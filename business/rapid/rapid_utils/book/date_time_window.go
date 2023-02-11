package book

import (
	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/business/rapid/models"
)

func pickupDateTimeWindow(quoteRequest *model.QuoteRequest) models.DateTimeWindow {
	startTime := "10:00:00 AM"
	endTime := "5:00:00 PM"
	return models.DateTimeWindow{
		StartTime: startTime,
		EndTime:   endTime,
		Date:      quoteRequest.QuoteRequest.PickupDate,
	}

}
func deliveryDateTimeWindow() models.DateTimeWindow {
	startTime := "10:00:00 AM"
	endTime := "4:30:00 PM"
	return models.DateTimeWindow{
		StartTime: startTime,
		EndTime:   endTime,
	}
}
