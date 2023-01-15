package rapid_utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func NewPickupDateTimeWindow(quoteReq *v1.QuoteRequest) models.DateTimeWindow {
	pickupByTime := quoteReq.Pickup.ShipperPickupReadyBy
	hour := strings.Split(pickupByTime, ":")[0]
	if len(hour) == 0 {
		startTime := "9:00:00 AM"
		endTime := "5:00:00 PM"
		now, err := time.Parse(time.RFC3339, quoteReq.ShipmentDetails.PickupDate)
		if err != nil {
			fmt.Println(err)
		}
		pickupDate := now.Format("01/02/2006")
		return models.DateTimeWindow{
			StartTime: startTime,
			EndTime:   endTime,
			Date:      pickupDate,
		}
	}
	var isAm bool
	if len(hour) > 1 {
		isZero := string(hour[0])
		if isZero == "0" {
			hour = hour[1:]
			hourInt, err := strconv.Atoi(hour)
			if err != nil {
				fmt.Println(err)
			}
			if hourInt < 12 {
				isAm = true
			}
		} else {
			hourInt, err := strconv.Atoi(hour)
			if err != nil {
				fmt.Println(err)
			}
			if hourInt < 12 {
				isAm = false
			}
		}
	}
	minute := strings.Split(pickupByTime, ":")[1]
	postFix := ""
	if isAm {
		postFix = "AM"
	} else {
		postFix = "PM"
	}
	startTime := hour + ":" + minute + ":" + "00 " + postFix
	endTime := "5:00:00 PM"
	now, err := time.Parse(time.RFC3339, quoteReq.ShipmentDetails.DisplayDate)
	if err != nil {
		fmt.Println(err)
	}
	pickupDate := now.Format("01/02/2006")
	return models.DateTimeWindow{
		StartTime: startTime,
		EndTime:   endTime,
		Date:      pickupDate,
	}
}
func NewDeliveryDateTimeWindow(booking *v1.QuoteRequest) models.DateTimeWindow {
	startTime := "9:00:00 AM"
	endTime := "5:00:00 PM"
	return models.DateTimeWindow{
		StartTime: startTime,
		EndTime:   endTime,
	}
}
