package bol_api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (bol Bol) EchoGetBOL(ctx echo.Context) error {
	quoteId := ctx.QueryParam("quoteId")
	// if !okay {
	// 	utils.Respond(bol.provider, ctx, utils.Ok{Success: false, StatusCode: errs.ErrInputDataNotValid.Cod}, errs.ErrInvalidInputs)
	// 	return
	// }
	if len(quoteId) < 1 || quoteId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	businessId := ctx.QueryParam("businessId")
	if len(businessId) < 1 || businessId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	ctxx := ctx.Request().Context()
	qtReq, err := bol.services.GetBooking(ctxx, quoteId)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	fmt.Print(qtReq)
	// clms, ok := ctx.Get("claims")
	// var claims auth.Claims
	// if ok {
	// 	claims, ok = clms.(auth.Claims)
	// 	if !ok {
	// 		utils.Respond(bol.provider, ctx, utils.Ok{Success: false, StatusCode: errs.ErrInvalidInputs.Cod}, errs.ErrInternal)
	// 		return
	// 	}
	// }
	// booking, err := bol.provider.BookingRepo().Get(ctx, id)
	// if err != nil {
	// 	utils.Respond(bol.provider, ctx, utils.Ok{Success: false, StatusCode: errs.ErrInvalidInputs.Cod}, errs.ErrInvalidInputs)
	// 	return
	// }
	// if claims.BusinessId != booking.ShipmentDetails.BusinessId {
	// 	utils.Respond(bol.provider, ctx, utils.Ok{Success: false, StatusCode: errs.ErrNotAllowed.Cod}, errs.ErrNotAllowed)
	// 	return
	// }
	return ctx.JSON(http.StatusOK, qtReq)
	// s3Input := &s3.GetObjectInput{
	// 	Bucket: aws.String("firstshipperbol"),
	// 	Key:    aws.String("bol" + quoteId + ".pdf"),
	// }
	// bolPdf, err := bol.services.S3Client.Client.GetObject(context.Background(), s3Input)
	// if err != nil {
	// 	return ctx.NoContent(http.StatusNotFound)
	// }
	// fmt.Println(bolPdf)
	// pdfBytes, err := io.ReadAll(bolPdf.Body)
	// if err != nil {
	// 	return ctx.NoContent(http.StatusNotFound)
	// }
	// ctx.Request().Header.Set("Cache-Control", "max-age=604800")
	// return ctx.Blob(http.StatusOK, "application/pdf", pdfBytes)

}
