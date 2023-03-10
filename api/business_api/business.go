package business_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"

	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/mid"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Business struct {
	services *services.Services
}

func New(services *services.Services, app *echo.Echo) Business {
	bis := Business{
		services: services,
	}
	businessGrp := app.Group("/business", mid.Protected(services))
	businessGrp.GET("", bis.EchoGetAllBusinesses)
	businessGrp.POST("", bis.EchoCreateBusiness)
	businessGrp.DELETE("", bis.EchoCloseBusiness)
	businessGrp.PATCH("/update_staff_role", bis.EchoUpdateStaffRole)
	businessGrp.POST("/add_staff", bis.EchoAddStaff)
	businessGrp.DELETE("/delete_staff", bis.DeleteStaff)
	businessGrp.POST("/add_business_address", bis.AddBusinessAddress)
	businessGrp.PATCH("/update_business_name", bis.UpdateBusinessName)
	businessGrp.GET("/get_basic_info", bis.GetBasicInfo)
	businessGrp.PATCH("/update_pickup_address", bis.UpdateDefaultPickupAddress)
	businessGrp.POST("/address", bis.AddBusinessAddress)
	businessGrp.POST("/phone", bis.UpdateBusinessPhoneNumber)
	businessGrp.POST("/allow_booking", bis.AllowBooking)
	return bis
}
func (business Business) GetBusinessById(ctx context.Context, req string) (*v1.Business, error) {
	return nil, nil
}

func (business Business) responde(result interface{}, err error, ctx echo.Context) {
	if err != nil {
		fmt.Println(err.Error())
		errMessage := v1.Ok{}
		newErr, ok := err.(errs.ApiErr)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, "please try again latter")
			return
		}
		if newErr.Cod >= 500 {
			ctx.JSON(http.StatusInternalServerError, "please try again latter")
			return

		} else {
			errMessage.Message = newErr.Message()
			errMessage.Code = int32(newErr.Cod)
			errMessage.Success = false
			ctx.JSON(newErr.Cod, echo.Map{"data": errMessage})
			return
		}

	}
	res, ok := result.(*v1.Ok)
	if ok {
		ctx.JSON(int(res.Code), echo.Map{"data": res})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, "please try again latter")
		return
	}
}
