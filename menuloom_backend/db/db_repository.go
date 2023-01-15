package db

import (
	"context"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
	"github.com/ramsfords/backend/menuloom_backend/core/models"
	"github.com/ramsfords/backend/menuloom_backend/db/menu_db"
	"github.com/ramsfords/backend/menuloom_backend/db/restaurant_db"
	"github.com/ramsfords/backend/menuloom_backend/db/user_db"
	validatedb "github.com/ramsfords/backend/menuloom_backend/db/validate_db"
	v1 "github.com/ramsfords/types_gen/v1"
)

type IDBContract interface {
	// Categories
	// Restaurant
	CreateRestaurant(ctx context.Context, data *v1.CreateRestaurantData) error
	GetRestaurant(ctx context.Context, id string) (*models.Restaurant, error)
	GetRestaurantData(ctx context.Context, restaurantPk string) (*v1.CreateRestaurantData, error)
	UpdateEmail(ctx context.Context, email string) error
	UpdateS3ProdUrl(ctx context.Context, url string) error
	UpdateAddress(ctx context.Context, data *v1.RestaurantAddress) error
	UpdateGoogleViewportUrl(ctx context.Context, url string) error
	UpdateOwnerId(ctx context.Context, id string) error
	UpdatePhoneNumber(ctx context.Context, phoneNumber string) error
	UpdateRestaurantName(ctx context.Context, name string) error
	UpdateRestaurantOpenHours(ctx context.Context, hours map[string]*v1.Hours) error
	UpdateRestaurantWebUrl(ctx context.Context, url string) error

	// MenuItems
	AddItem(ctx context.Context, data *v1.Item, restaurantUrl string) error
	GetMenu(ctx context.Context, id string) (*models.Menu, error)
	CreateCategories(ctx context.Context, restaurantPk string, data []*v1.Category) error
	CreateCategory(ctx context.Context, data *v1.Category, restaurantUrl string) error
	CreateItems(ctx context.Context, datas []*v1.Item, restaurantUrl string) error
	CreateItem(ctx context.Context, data *v1.Item, restaurantUrl string) error
	GetItems(ctx context.Context, restaurantPk string) ([]*v1.Item, error)
	RemoveItem(ctx context.Context, itemName string, restaurantPk string) error
	GetCategories(ctx context.Context, id string) ([]*v1.Category, error)
	UpdateCategories(ctx context.Context, data []*v1.Category, restaurantPk string) error
	UpdateCategory(ctx context.Context, data *v1.Category, restaurantPk string) error
	UpdateItem(ctx context.Context, data *v1.Item, restaurantPk string) error
	UpdateItems(ctx context.Context, data []*v1.Item, restaurantPk string) error

	// Users
	CreateUser(ctx context.Context, input *v1.User) error
	AddUser(ctx context.Context, data *v1.User) error
	Getuser(ctx context.Context, email string) (*v1.User, error)
	RemoveUser(ctx context.Context, email string) error
	UpdateUserRole(ctx context.Context, email string, data []*v1.Role) error

	// restaurant data validate
	UpdateValidate(ctx context.Context, id string, isValid bool) error
	GetValidate(ctx context.Context, id string) (bool, error)
}

type Repository struct {
	dynamo.DB
	menu_db.MenuDb
	restaurant_db.RestaurantDb
	user_db.UserDb
	validatedb.ValidateDb
}

func New(configs *configs.Config, db dynamo.DB) Repository {
	return Repository{
		MenuDb: menu_db.MenuDb{
			DB:     db,
			Config: configs,
		},
		RestaurantDb: restaurant_db.RestaurantDb{
			DB:     db,
			Config: configs,
		},
		ValidateDb: validatedb.ValidateDb{
			DB: db,
		},
	}
}
