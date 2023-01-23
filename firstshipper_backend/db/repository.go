package db

import (
	"context"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/db/booking_db"
	"github.com/ramsfords/backend/firstshipper_backend/db/business_db"
	"github.com/ramsfords/backend/firstshipper_backend/db/location_db"
	"github.com/ramsfords/backend/firstshipper_backend/db/quote_db"
	"github.com/ramsfords/backend/firstshipper_backend/db/rapid_db"
	"github.com/ramsfords/backend/firstshipper_backend/db/user_db"
	"github.com/ramsfords/backend/foundations/dynamo"
	v1 "github.com/ramsfords/types_gen/v1"
)

type DB interface {

	// Users
	SaveUser(ctx context.Context, usr v1.User, businessId string) error
	UpdateUser(ctx context.Context, businessId string, user v1.User) error
	DeleteUser(ctx context.Context, userId string, businessId string) error
	Getuser(ctx context.Context, email string) (*v1.User, error)
	UpdateUserPassword(ctx context.Context, usr v1.User, businessId string) error
	UpdateUserConfirmEmail(ctx context.Context, businessId, email string) (bool, error)

	//LTL BOOKING DB
	SaveBooking(ctx context.Context, booking *v1.BookingResponse) error
	UpdateBooking(ctx context.Context, booking *v1.BookingResponse, businessId string) error
	DeleteBooking(ctx context.Context, BookingId string, businessId string) error
	GetBooking(ctx context.Context, BookingId string) (*v1.BookingResponse, error)
	GetAllBookingsByBusinessId(ctx context.Context, businessId string) ([]*v1.BookingResponse, error)
	GetAllBookings(ctx context.Context) ([]*v1.BookingResponse, error)

	//LTL Location DB
	SaveLocation(ctx context.Context, businessId string, location *v1.Location) error
	DeleteLocation(ctx context.Context, locationId string, businessId string) error
	GetLocations(ctx context.Context, businessId string) ([]*v1.Location, error)
	GetAllLocations(ctx context.Context) ([]*v1.Location, error)
	UpdateLocation(ctx context.Context, businessId string, location *v1.Location) error
	AddLocationAddress(ctx context.Context, businessId string, address *v1.Address) (*v1.Address, error)
	GetLocation(ctx context.Context, businessId string, locationId string) (*v1.Location, error)

	// Quote DB
	SaveQuote(ctx context.Context, qtReq *model.QuoteRequest) error
	GetQuoteByQuoteId(ctx context.Context, quoteId string, businessId string) (*v1.QuoteRequest, error)
	DeleteQuote(ctx context.Context, quoteId string) error
	DeleteAllQuoteByBusinessId(ctx context.Context, buisnessId string) error
	DeleteQuotesByQuoteIds(ctx context.Context, businessId string, quoteId []string) error
	GetAllQuotesByBusinessId(ctx context.Context, businessId string) ([]*v1.QuoteRequest, error)
	GetBidsByQuoteId(ctx context.Context, businessId string, quoteId string) ([]*v1.Bid, error)
	GetBidByBidID(ctx context.Context, businessId string, quoteId string, bidId string) (*v1.Bid, error)
	UpdateQuote(ctx context.Context, quoteReq *v1.QuoteRequest) error
	GetBidsWithQuoteByQuoteId(ctx context.Context, businessId string, quoteId string) (*model.BidsWithQuote, error)
	GetBidWithQuoteByQuoteId(ctx context.Context, businessId string, quoteId string, bidId string) (*model.BidWithQuote, error)

	// Rapid Quote DB
	SaveRapidQuote(ctx context.Context, quote models.QuoteRate, quoteReq v1.QuoteRequest) error
	GetRapidQuote(ctx context.Context, quoteId string) (models.QuoteRate, error)
	UpdateRapidQuote(ctx context.Context, quoteRate models.QuoteRate, quoteReq v1.QuoteRequest) error
	DeleteRapidQuote(ctx context.Context, quoteId string) error

	// Business DB
	SaveStaff(ctx context.Context, businessId string, staff v1.User) error
	UpdateBusinessAddressUpdateNeeded(ctx context.Context, businessId string) error
	DeleteStaff(ctx context.Context, businessId string, email string) error
	DeleteBusiness(ctx context.Context, businessId string) error
	GetAllBusinesses(ctx context.Context, businessId string) ([]v1.Business, error)
	GetStaffsForABusiness(ctx context.Context, businessId string) ([]*v1.FrontEndUser, error)
	GetBusiness(ctx context.Context, businessId string) (*v1.Business, error)
	SaveDefaultPickup(ctx context.Context, businessId string, address v1.Location) error
	SaveBusiness(ctx context.Context, business v1.Business) error
	UpdateStaffRole(ctx context.Context, businessId string, staffEmail string, roles []v1.Role) error
	UpdateBusiness(ctx context.Context, businessId string, business v1.Business) error
	AddPhoneNumber(ctx context.Context, businessId string, phoneNumber *v1.PhoneNumber) (*v1.PhoneNumber, error)
}

type Repository struct {
	dynamo.DB
	user_db.UserDb
	booking_db.BookingDb
	location_db.LocationDb
	quote_db.QuoteDb
	rapid_db.RapidDb
	business_db.BusinessDb
}

// type Repository struct {
// 	dynamo.DB
// 	menu_db.MenuDb
// 	restaurant_db.RestaurantDb
// 	user_db.UserDb
// 	validatedb.ValidateDb
// }

func New(db dynamo.DB, configs *configs.Config) DB {
	var repo DB = Repository{}
	repo = Repository{
		UserDb: user_db.UserDb{
			Config: configs,
			DB:     db,
		},
		BookingDb: booking_db.BookingDb{
			Config: configs,
			DB:     db,
		},
		LocationDb: location_db.LocationDb{
			Config: configs,
			DB:     db,
		},
		QuoteDb: quote_db.QuoteDb{
			Config: configs,
			DB:     db,
		},
		RapidDb: rapid_db.RapidDb{
			Config: configs,
			DB:     db,
		},
		BusinessDb: business_db.BusinessDb{
			Config: configs,
			DB:     db,
		},
	}
	return repo
}
