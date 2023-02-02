package test

import (
	"testing"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/adobe"
)

func TestBolGen(t *testing.T) {
	conf := configs.GetConfig()
	seClient := S3.New(conf)
	adobe := adobe.NewAdobe(seClient)
	adobe.UrlToPdf("23105")
}
