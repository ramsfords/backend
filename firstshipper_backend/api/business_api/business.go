package business_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"

	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Business struct {
	services *services.Services
	rapid    *rapid.Rapid
}

func New(services *services.Services, rapid *rapid.Rapid, app *echo.Group) Business {
	bis := Business{
		services: services,
		rapid:    rapid,
	}
	businessGrp := app.Group("/business")
	businessGrp.GET("", bis.GinGetAllBusinesses)
	businessGrp.GET("/:id", bis.GinGetAllBusiness)
	businessGrp.POST("", bis.GinCreateBusiness)
	businessGrp.DELETE("", bis.GinCloseBusiness)
	businessGrp.PATCH("", bis.GinUpdateBusiness)
	businessGrp.PATCH("/update_staff_role", bis.GinUpdateStaffRole)
	businessGrp.POST("/add_staff", bis.GinAddStaff)
	businessGrp.POST("/delete_staff", bis.DeleteStaff)
	businessGrp.POST("/add_business_address", bis.AddBusinessAddress)
	businessGrp.GET("/get_basic_info/:businessId", bis.GetBasicInfo)
	businessGrp.PATCH("/update_pickup_address", bis.UpdateDefaultPickupAddress)
	businessGrp.POST("/address/:businessId", bis.AddBusinessAddress)
	businessGrp.POST("/phone/:businessId", bis.UpdateBusinessPhoneNumber)

	return bis
}
func (business Business) GetBusinessById(ctx context.Context, req *v1.Id) (*v1.Business, error) {
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
