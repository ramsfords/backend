package bol_api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v5"
)

func (bol Bol) GinGetBOL(ctx echo.Context) error {
	id := ctx.PathParam("id")
	// if !okay {
	// 	utils.Respond(bol.provider, ctx, utils.Ok{Success: false, StatusCode: errs.ErrInputDataNotValid.Cod}, errs.ErrInvalidInputs)
	// 	return
	// }
	if len(id) < 1 || id == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}

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
	s3Input := &s3.GetObjectInput{
		Bucket: aws.String("firstshipperbol"),
		Key:    aws.String("bol" + id + ".pdf"),
	}
	bolPdf, err := bol.services.S3Client.Client.GetObject(context.Background(), s3Input)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	fmt.Println(bolPdf)
	pdfBytes, err := io.ReadAll(bolPdf.Body)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	ctx.Request().Header.Set("Cache-Control", "max-age=604800")
	return ctx.Blob(http.StatusOK, "application/pdf", pdfBytes)

}
