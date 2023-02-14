package utils

var (
	// CarrierContactType is the type of the carrier contact
	XPO = CarrierContact{
		CarrierName: "XPO LTL",
		Email:       "",
		Phone:       "213-744-0664",
	}
	RoadRunner = CarrierContact{
		CarrierName: "Roadrunner Freight",
		Email:       "LTLCustomerService@roadrunnerLTL.com",
		Phone:       "323-780-3488",
	}
	ClearLane = CarrierContact{
		CarrierName: "Clear Lane Freight Systems",
		Email:       "customerservice@clearlanefreight.com",
		Phone:       "866-491-9255",
	}
	RnL = CarrierContact{
		CarrierName: "R&L",
		Email:       "",
		Phone:       "800-535-1984",
	}
)

type CarrierContact struct {
	CarrierName string `json:"carrierName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

func GetCarrierContact(carrierName string) CarrierContact {
	switch carrierName {
	case "XPO LTL":
		return XPO
	case "Roadrunner Freight":
		return RoadRunner
	case "Clear Lane Freight Systems":
		return ClearLane
	case "R&L":
		return RnL
	default:
		return CarrierContact{}
	}
}
