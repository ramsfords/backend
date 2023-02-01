package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
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
	SaveBooking(ctx context.Context, bookingRes *model.QuoteRequest) error
	UpdateBooking(ctx context.Context, booking *v1.BookingResponse, businessId string) error
	DeleteBooking(ctx context.Context, BookingId string, businessId string) error
	GetBooking(ctx context.Context, bookingId string, businessId string) (*v1.BookingResponse, error)
	GetAllBookingsByBusinessId(ctx context.Context, businessId string) ([]*v1.BookingResponse, error)
	GetAllBookings(ctx context.Context) ([]*v1.BookingResponse, error)

	//LTL Location DB
	SaveLocation(ctx context.Context, businessId string, location *v1.Location) error
	DeleteLocation(ctx context.Context, locationId string, businessId string) error
	GetLocations(ctx context.Context, businessId string) ([]*v1.Location, error)
	GetAllLocations(ctx context.Context) ([]*v1.Location, error)
	UpdateLocation(ctx context.Context, businessId string, location *v1.Address) error
	AddLocationAddress(ctx context.Context, businessId string, address *v1.Address) (*v1.Address, error)
	GetLocation(ctx context.Context, businessId string, locationId string) (*v1.Location, error)

	// Quote DB
	SaveQuote(ctx context.Context, qtReq *model.QuoteRequest) error
	GetQuoteByQuoteId(ctx context.Context, quoteId string, businessId string) (*model.QuoteRequest, error)
	DeleteQuote(ctx context.Context, quoteId string) error
	DeleteAllQuoteByBusinessId(ctx context.Context, buisnessId string) error
	DeleteQuotesByQuoteIds(ctx context.Context, businessId string, quoteId []string) error
	GetAllQuotesByBusinessId(ctx context.Context, businessId string) ([]*v1.QuoteRequest, error)
	GetBidsByQuoteId(ctx context.Context, businessId string, quoteId string) ([]*v1.Bid, error)
	GetBidsWithQuoteByQuoteId(ctx context.Context, businessId string, quoteId string) (*model.QuoteRequest, error)
	GetBidByBidID(ctx context.Context, businessId string, quoteId string, bidId string) (*v1.Bid, error)
	UpdateQuote(ctx context.Context, quoteReq *v1.QuoteRequest) error
	GetBidByQuoteId(ctx context.Context, businessId string, quoteId string, bidId string) (*model.QuoteRequest, error)

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
	GetStaffsForABusiness(ctx context.Context, businessId string) ([]*v1.User, error)
	GetBusiness(ctx context.Context, businessId string) (*v1.Business, error)
	SaveDefaultPickup(ctx context.Context, businessId string, address v1.Location) error
	SaveBusiness(ctx context.Context, business v1.Business, businessId string) error
	UpdateStaffRole(ctx context.Context, businessId string, staffEmail string, roles []v1.Role) error
	UpdateBusiness(ctx context.Context, businessId string, business v1.Business) error
	AddPhoneNumber(ctx context.Context, businessId string, phoneNumber *v1.PhoneNumber) (*v1.PhoneNumber, error)
	UpdateBusinessName(ctx context.Context, businessId string, businessName string) error
	GetAllDataByBusinessId(ctx context.Context, businessId string) (*model.BusinessData, error)
	// Incraase quote Count
	IncreateQuoteCount()
	// Get Quote Count
	GetQuoteCount() int64
}

type Repository struct {
	sync.Mutex
	QuoteCount int64
	Config     *configs.Config
	DB         dynamo.DB
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
	getCountInput := &dynamodb.GetItemInput{
		TableName: aws.String(configs.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: "quoteCount",
			},
			"sk": &types.AttributeValueMemberS{
				Value: "quoteCount",
			},
		},
	}
	qtCount, err := db.Client.GetItem(context.Background(), getCountInput)
	if err != nil {
		log.Println("Error getting quote count", err)
	}
	countValue := qtCount.Item["quoteCount"].(*types.AttributeValueMemberN).Value
	fmt.Println(qtCount)
	var repo DB = &Repository{}
	var count int64
	if countValue != "0" {
		if _, ok := repo.(*Repository); ok {
			quoteIntCount, err := strconv.ParseInt(countValue, 10, 64)
			if err != nil {
				log.Println("Error converting quote count to int", err)
			}
			count = quoteIntCount
		}
	}
	repo = &Repository{
		QuoteCount: count,
		Config:     configs,
		DB:         db,
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
	repo.IncreateQuoteCount()
	return repo
}
func (repo *Repository) IncreateQuoteCount() {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()
	repo.QuoteCount++
	res, err := repo.DB.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(repo.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{
				Value: "quoteCount",
			},
			"sk": &types.AttributeValueMemberS{
				Value: "quoteCount",
			},
		},
		UpdateExpression: aws.String("set quoteCount =  :val"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberN{
				Value: fmt.Sprint(repo.QuoteCount),
			},
		},
		ReturnValues: "UPDATED_NEW",
	})
	if err != nil {
		log.Println("Error updating quote count", err)
	}
	fmt.Println(res)
}
func (repo *Repository) GetQuoteCount() int64 {
	return repo.QuoteCount
}
