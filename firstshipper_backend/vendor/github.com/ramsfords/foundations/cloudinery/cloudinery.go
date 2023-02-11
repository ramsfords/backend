package cloudinery

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ramsfords/configs"
)

type Cloudinary struct {
	*cloudinary.Cloudinary
	conf configs.CloudinaryConfig
}

func New(conf configs.Config) Cloudinary {
	base := Cloudinary{
		conf: conf.CloudinaryConfig,
	}
	cld, err := cloudinary.NewFromParams(base.conf.CloudName, base.conf.ApiKey, base.conf.ApiSecret)
	if err != nil {
		log.Fatal("could not create cloudinary instance")
	}
	cld.Config.URL.Secure = true
	base.Cloudinary = cld

	return base

}
func (cloud Cloudinary) UploadImageFromUrl(file interface{}, itemName string, folderName string) (*uploader.UploadResult, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cloud.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       itemName,
		UniqueFilename: api.Bool(false),
		Folder:         folderName,
		Overwrite:      api.Bool(true)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
