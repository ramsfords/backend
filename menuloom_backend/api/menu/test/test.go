package test

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

func getGetClient() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8090/restaurant", nil)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJxNnM1a24wbHNpeHU1a2IiLCJleHAiOjE2NzM3NDcyOTgsImlkIjoicXQzZ2g1ZW84ZDBlMjhqIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.b5YE7-QJSVEmd7p4GSZcz-SAM1ehfnX3YHnuMH6Ieag")
	return req
}
func getPostClient(data any) *http.Request {
	datas, _ := json.Marshal(data)
	buf := bytes.NewReader(datas)
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:8090/restaurant", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJxNnM1a24wbHNpeHU1a2IiLCJleHAiOjE2NzM3NDcyOTgsImlkIjoicXQzZ2g1ZW84ZDBlMjhqIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.b5YE7-QJSVEmd7p4GSZcz-SAM1ehfnX3YHnuMH6Ieag")
	return req
}
func getPatchClient(data any) *http.Request {
	datas, _ := json.Marshal(data)
	buf := bytes.NewReader(datas)
	req, _ := http.NewRequest(http.MethodPatch, "http://127.0.0.1:8090/restaurant", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJxNnM1a24wbHNpeHU1a2IiLCJleHAiOjE2NzM3NDcyOTgsImlkIjoicXQzZ2g1ZW84ZDBlMjhqIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.b5YE7-QJSVEmd7p4GSZcz-SAM1ehfnX3YHnuMH6Ieag")
	return req
}
