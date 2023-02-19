package errs

import (
	"errors"
	"fmt"
)

var (
	ErrEmailNotVerficed = ApiErr{
		Cod: 400,
		Msg: "Email address is not verified. please confirm your email.",
	}
	ErrMissingData = ApiErr{
		Cod: 400,
		Msg: "Missing or invalid fields in input data.",
	}
	ErrConfirmEmail = ApiErr{
		Cod: 400,
		Msg: "Email confirmation did not work, please verify that you are signed up already.",
	}
	ErrTooManyRequest = ApiErr{
		Cod: 429,
		Msg: "Too Many Request. Please try again later.",
	}
	ErrConfirmEmailExpired = ApiErr{
		Cod: 400,
		Msg: "Email confirmation Code already Expired.",
	}
	ErrMissingFormData = ApiErr{
		Cod: 400,
		Msg: "Missing or invalid fields in input data.",
	}
	ErrForgotPasswordUserNotFound = ApiErr{
		Cod: 200,
		Msg: "if you have account with us. You will receive password reset email.",
	}
	ErrEmailNotValid = ApiErr{
		Cod: 400,
		Msg: "Missing or invalid Email Address.",
	}
	ErrInputDataNotValid = ApiErr{
		Cod: 400,
		Msg: "Input data validation failure",
	}
	ErrWeakPassword = ApiErr{
		Cod: 400,
		Msg: "Password is not Strong Enough.",
	}
	ErrNewPasswordSameAsOldPassword = ApiErr{
		Cod: 400,
		Msg: "can not use same old password to reset new password",
	}
	ErrPasswordDoesNotMatch = ApiErr{
		Cod: 400,
		Msg: "Password or User Does not match",
	}
	ErrUploadInvalidFile = ApiErr{
		Cod: 400,
		Msg: "Unable to read uploaded file",
	}
	ErrForgtPasswordTokenNotIssued = ApiErr{
		Cod: 400,
		Msg: "Please request forgot password token/click for forgot password",
	}
	ErrInvalidInputs = ApiErr{
		Cod: 400,
		Msg: "Invalid or Insufficient Inputs / Bad Request",
	}
	ErrNotFound = ApiErr{
		Cod: 404,
		Msg: "Not found",
	}
	ErrUserNotFound = ApiErr{
		Cod: 404,
		Msg: "User is not found. Please sign up",
	}
	ErrNotAllowed = ApiErr{
		Cod: 405,
		Msg: "Not Allowed / Bad Request",
	}
	ErrTimeout = ApiErr{
		Cod: 408,
		Msg: "Operation was timedout",
	}
	ErrConflict = ApiErr{
		Cod: 409,
		Msg: "Resource already exist",
	}
	ErrUserAlreadyExits = ApiErr{
		Cod: 409,
		Msg: "User Already Exits, Please log in",
	}
	ErrBusinessAlreadyExits = ApiErr{
		Cod: 409,
		Msg: "Business Already Exits, Please login with admin Email",
	}
	ErrAcccountAlreadyExists = ApiErr{
		Cod: 409,
		Msg: "Account already exists with email or user_name",
	}
	ErrSignUpNotEnabled = ApiErr{
		Cod: 500,
		Msg: "sign up is not enabled. We are working on opening our system for users ASAP.",
	}
	ErrCancel = ApiErr{
		Cod: 500,
		Msg: "Operation was cancelled",
	}
	ErrInternal = ApiErr{
		Cod: 500,
		Msg: "Internal error",
	}
	ErrCognitoPasswordResetFailed = ApiErr{
		Cod: 500,
		Msg: "Could not reset password at cognito for user",
	}
	ErrStoreInternal = ApiErr{
		Cod: 500,
		Msg: "Internal backend store error",
	}
	ErrUnknown = ApiErr{
		Cod: 500,
		Msg: "Internal server error",
	}
	ErrDownloadFailed = ApiErr{
		Cod: 500,
		Msg: "Download failed",
	}
	ErrMetadataUpdateFailed = ApiErr{
		Cod: 500,
		Msg: "Metadata update failed",
	}
	ErrGetMetadataFailed = ApiErr{
		Cod: 500,
		Msg: "Unable to get metadata information",
	}
	ErrUploadFailed = ApiErr{
		Cod: 500,
		Msg: "Upload failed",
	}
	ErrCannotStartDatabase = ApiErr{
		Cod: 500,
		Msg: "Can not connect to the database",
	}
	ErrCannotStartGrpcServer = ApiErr{
		Cod: 500,
		Msg: "Can not connect Grpc service",
	}
	ErrCannotStartCacheDatabase = ApiErr{
		Cod: 500,
		Msg: "Can not connect to cache database",
	}
	ErrCannotGetCacheData = ApiErr{
		Cod: 500,
		Msg: "Can not get data from cache database",
	}
	ErrCannotSetCacheData = ApiErr{
		Cod: 500,
		Msg: "Can set data to cache database",
	}
	ErrSignInRequstData = ApiErr{
		Cod: 500,
		Msg: "Sign in request most include username and password",
	}
	ErrStartingComponents = ApiErr{
		Cod: 500,
		Msg: "Starting application error",
	}
	ErrStartingThirdPartyEmailServer = ApiErr{
		Cod: 500,
		Msg: "Starting thridParty Email Service",
	}
	ErrLocationCreationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in creating location",
	}
	ErrLocationUpdationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in updating location",
	}
	ErrLocationDeletionFailed = ApiErr{
		Cod: 500,
		Msg: "Error in deleting location",
	}
	ErrBookingCreationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in creating booking",
	}
	ErrBookingFetchFailed = ApiErr{
		Cod: 500,
		Msg: "Error in getting booking",
	}
	ErrQuoteDeletionFailed = ApiErr{
		Cod: 500,
		Msg: "Error in deleting quote",
	}
	ErrQuoteFetchFailed = ApiErr{
		Cod: 500,
		Msg: "Error in getting quote",
	}
	ErrQuoteUpdationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in updatin quote",
	}
	ErrBusinessFetchFailed = ApiErr{
		Cod: 500,
		Msg: "Error in getting business",
	}
	ErrBusinessCreationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in updatin business",
	}
	ErrBusinessUpdationFailed = ApiErr{
		Cod: 500,
		Msg: "Error in updating business",
	}
	ErrBusinessClosureFailed = ApiErr{
		Cod: 500,
		Msg: "Error in closing business",
	}
	ErrInvalidPickupZipCode = ApiErr{
		Cod: 500,
		Msg: "pickup zipcode is not valid",
	}
	ErrInvalidPickupLocationServices = ApiErr{
		Cod: 500,
		Msg: "pickup location services required is not valid. you must choose if the location needs liftgate pick up or it has a loading doc",
	}
	ErrInvalidDeliveryLocationServices = ApiErr{
		Cod: 500,
		Msg: "delivery location services required is not valid. you must choose if the location needs liftgate delivery or it has a loading dock",
	}
	ErrInvalidDeliveryZipCode = ApiErr{
		Cod: 500,
		Msg: "delivery zipcode is not valid",
	}
	ErrInvalidShipmentCommodity = ApiErr{
		Cod: 500,
		Msg: "shipment commodities details is not valid",
	}
	InvalidDimensionUOM = ApiErr{
		Cod: 500,
		Msg: "only Inch measurement are accepted. Please provide dimension in inch",
	}
	InvalidWeightUOM = ApiErr{
		Cod: 500,
		Msg: "only LBS measurement are accepted. Please provide Weights in LBS",
	}
	InvalidPackageType = ApiErr{
		Cod: 500,
		Msg: "shipment items package type if not valid. Please provide what kind of package are you shipping such as pallets, box and so on",
	}
	InvalidShipmentDescription = ApiErr{
		Cod: 500,
		Msg: "shipment item shipping description is not valid. please mention what are you shipping",
	}
	InvalidDimension = ApiErr{
		Cod: 500,
		Msg: "shipping items dimensions are not valid",
	}
	InvalidWeight = ApiErr{
		Cod: 500,
		Msg: "shipping items weight is not valid",
	}
	InvalidPickupAddress = ApiErr{
		Cod: 500,
		Msg: "pickup address is not valid",
	}
	InvalidAddress = ApiErr{
		Cod: 500,
		Msg: "address is not valid",
	}
	InvalidContactInfo = ApiErr{
		Cod: 500,
		Msg: "contact is not valid",
	}
	InvalidCompanyName = ApiErr{
		Cod: 500,
		Msg: "company name is not valid",
	}
	Redirect = ApiErr{
		Cod: 302,
		Msg: "redirect to new path",
	}
	InvalidEmailAddress = ApiErr{
		Cod: 500,
		Msg: "email is not valid",
	}
	InvalidPhoneNumber = ApiErr{
		Cod: 500,
		Msg: "phone number is not valid",
	}
	InvalidCommodityServices = ApiErr{
		Cod: 500,
		Msg: "commodity services are not valid",
	}
	InvalidMismatchBookingAndQuoteCommodity = ApiErr{
		Cod: 500,
		Msg: "quote commodities and bookcommodities are different",
	}
	ErrDataNotFound = ApiErr{
		Cod: 404,
		Msg: "commodity services are not valid",
	}
	InvalidBookingData = ApiErr{
		Cod: 500,
		Msg: "booking data and quote data does not match",
	}
)

type Response interface {
	Code() int
	Message() string
	error
}
type ApiErr struct {
	Msg string `json:"msg,omitempty"`
	Cod int    `json:"Cod,omitempty"`
	Err error  `json:"err,omitempty"`
}

func (a ApiErr) Message() string {
	if a.Msg == "" && a.Err != nil {
		return a.Err.Error()
	}
	return a.Msg
}
func (a ApiErr) Error() string {
	var err error
	err = a.Err
	if err == nil {
		err = errors.New("")
	}
	if a.Cod >= 500 {
		return a.Msg
	}
	return fmt.Sprintf("Err: %s, Msg: %s", err.Error(), a.Message())
}
func (a ApiErr) Code() int {
	return a.Cod
}

// NewApiError helper function to contruct API error
func NewApiError(Cod int, Msg string, err error) ApiErr {
	return ApiErr{Cod: Cod, Msg: Msg, Err: err}
}

// custom error with api err
func NewCustomErr(err ApiErr, title string) ApiErr {
	newApiErr := ApiErr{
		Msg: title + " " + err.Msg,
		Cod: err.Cod,
	}
	return newApiErr

}
func WrapApiError(err ApiErr, rootErr error) Response {
	return ApiErr{
		Msg: err.Msg,
		Cod: err.Cod,
		Err: rootErr,
	}
}
