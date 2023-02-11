package test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"testing"
)

func TestDownloadImageFromURL(t *testing.T) {
	// Get the response bytes from the url

	URL := "https://en.wikipedia.org/wiki/Biryani#/media/File:Biryani_of_Lahore.jpg"
	validUrl, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validUrl)
	filename := path.Base(URL)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Received non 200 response code")
	}
	// Create a empty file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
