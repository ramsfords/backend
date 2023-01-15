package file

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type UploadAssets interface {
	GetImagesUrlList() []string
	GetUrl(index int) string
	GetContainerName() string
	SetUrl(index int, url string)
}

func UploadImageFromUrl(menuloomStaticImageFolderUrl string, item UploadAssets, downloadTo string) (interface{}, error) {
	var errs error
	// get the image from the url
	for index, _ := range item.GetImagesUrlList() {
		urlStr := item.GetUrl(index)
		validUrl, err := url.ParseRequestURI(urlStr)
		if err != nil {
			errs = err
			continue
		}
		fmt.Println(validUrl)
		// filename := path.Base(URL)
		response, err := http.Get(urlStr)
		if err != nil {
			errs = err
			continue
		}

		if response.StatusCode != 200 {
			fmt.Println("Received non 200 response code")
		}
		// // Create a empty file
		// file, err := os.Create(filename)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer file.Close()

		// Write the bytes to the fiel
		//readData, err := io.ReadAll(response.Request.Body)
		//Create a empty file
		fileName := downloadTo + item.GetContainerName() + getExtensionFromUrl(response.Header.Get("Content-Type"))
		_, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
		item.SetUrl(index, fileName)
	}
	if errs != nil {
		return item, errs
	}
	return item, nil

}
func getExtensionFromUrl(contentType string) string {
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	case "image/tiff":
		return ".tiff"
	case "image/bmp":
		return ".bmp"
	case "image/avif":
		return ".avif"
	case "image/jpg":
		return ".jpg"
	default:
		return ""
	}
}
