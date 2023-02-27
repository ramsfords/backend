package authapi

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (auth AuthApi) EchoResetPassword(ctx echo.Context) error {
	data := &v1.ResetPassword{}
	err := ctx.Bind(data)
	if err != nil || data.NewPassword != data.ConfirmPassword || !utils.IsEmailValid(data.Email) || len(data.NewPassword) < 6 || data.Token == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// resetTokenEmail, err := auth.services.Crypto.Decrypt(data.Token)
	// if err != nil || string(resetTokenEmail) != data.Email {
	// 	return ctx.NoContent(http.StatusBadRequest)
	// }

	// auth.services.CognitoClient.ResetPassword(ctx.Request().Context(), data.Email, data.NewPassword)
	passwordChangeData := map[string]interface{}{
		"password": data.NewPassword}
	_, err = auth.services.SupaClient.Auth.UpdateUser(ctx.Request().Context(), data.Token, passwordChangeData)
	if err != nil {
		auth.services.Logger.Error("Error updating user password", map[string]interface{}{"err": err})
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}

// access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNjc3NTM4ODkzLCJzdWIiOiI0NWIxZGRhOC03NDk0LTQyYWUtYmQ4Yy1hNjcxYTkyYzBiOWQiLCJlbWFpbCI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInBob25lIjoiIiwiYXBwX21ldGFkYXRhIjp7InByb3ZpZGVyIjoiZW1haWwiLCJwcm92aWRlcnMiOlsiZW1haWwiXX0sInVzZXJfbWV0YWRhdGEiOnsiYnVzaW5lc3MiOnsiYWNjb3VudGluZ0VtYWlsIjoia2FuZGVsc3VyZW5AZ21haWwuY29tIiwiYWRtaW5FbWFpbCI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsImFkbWluVXNlciI6eyJlbWFpbCI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsIm5hbWUiOiJzdXJlbmRyYSBrYW5kZWwiLCJvcmdhbml6YXRpb25JZCI6IjEyMzU4OTkiLCJ0eXBlIjoidXNlciIsInVzZXJOYW1lIjoia2FuZGVsc3VyZW5AZ21haWwuY29tIn0sImJ1c2luZXNzRW1haWwiOiJrYW5kZWxzdXJlbkBnbWFpbC5jb20iLCJidXNpbmVzc0lkIjoiMTIzNTkwMCIsImN1c3RvbWVyU2VydmljZUVtYWlsIjoia2FuZGVsc3VyZW5AZ21haWwuY29tIiwiaGlnaFByaW9yaXR5RW1haWwiOiJrYW5kZWxzdXJlbkBnbWFpbC5jb20iLCJuZWVkc0FkZHJlc3NVcGRhdGUiOnRydWUsIm5lZWRzRGVmYXVsdERlbGl2ZXJ5QWRkcmVzc1VwZGF0ZSI6dHJ1ZSwibmVlZHNEZWZhdWx0UGlja3VwQWRkcmVzc1VwZGF0ZSI6dHJ1ZSwicmVmZXJyZWRCeSI6InJlZmVycmFsIiwidHlwZSI6ImJ1c2luZXNzIn0sImNvbmZpcm1QYXNzd29yZCI6IkZlcmluYUAxMjM0IiwiZW1haWwiOiJrYW5kZWxzdXJlbkBnbWFpbC5jb20iLCJuYW1lIjoic3VyZW5kcmEga2FuZGVsIiwib3JnYW5pemF0aW9uSWQiOiIxMjM1OTAwIiwicGFzc3dvcmQiOiJGZXJpbmFAMTIzNCIsInVzZXJOYW1lIjoia2FuZGVsc3VyZW5AZ21haWwuY29tIn0sInJvbGUiOiJhdXRoZW50aWNhdGVkIiwiYWFsIjoiYWFsMSIsImFtciI6W3sibWV0aG9kIjoib3RwIiwidGltZXN0YW1wIjoxNjc3NTM1MjkzfV0sInNlc3Npb25faWQiOiJlNzJkMmJjNy1lYjY4LTRjOTEtYTUyZC0zYjVlMjY3ZDZkMDcifQ.NFB1thmFurzrDL2RWW4CnRTsgX51Fh9PDOdd3wWRM-c&expires_in=3600&refresh_token=-H7MpjN1yCJG9P5gk6jlXQ&token_type=bearer&type=recovery
