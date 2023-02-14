package cloudinery

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ramsfords/backend/configs"
)

type Cloudinery struct {
	*cloudinary.Cloudinary
}

func New(conf configs.CloudinaryConfig) *Cloudinery {
	cld, err := cloudinary.NewFromParams(conf.CloudName, conf.ApiKey, conf.ApiSecret)
	if err != nil {
		log.Fatal("could not create cloudinary instance")
	}
	cld.Config.URL.Secure = true

	return &Cloudinery{cld}

}
func (cloudinery Cloudinery) UploadImageFromUrl(file interface{}, itemName string, folderName string) (*uploader.UploadResult, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cloudinery.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       itemName,
		UniqueFilename: api.Bool(false),
		Folder:         folderName,
		Overwrite:      api.Bool(true)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
