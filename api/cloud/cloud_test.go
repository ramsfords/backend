package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	vision "cloud.google.com/go/vision/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func TestOcr(t *testing.T) {
	client, ctx := createClient()
	defer client.Close()
	ocr("../../test.webp", client, ctx)
	fmt.Println(ctx)
}
func createClient() (*vision.ImageAnnotatorClient, context.Context) {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, ctx
}
func ocr(imagePath string, client *vision.ImageAnnotatorClient, ctx context.Context) {
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	image := &pb.Image{
		Content: imageData,
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to perform OCR: %v", err)
	}

	if len(annotations) == 0 {
		fmt.Println("No text found in the image.")
		return
	}

	fmt.Println("Text found in the image:")
	for _, annotation := range annotations {
		fmt.Println(annotation.Description)
	}
	fmt.Println(annotations)
}

// curl -H "Content-Type: multipart/form-data" \
//      -H "Accept: application/json" \
//      -H "CLIENT-ID: vrfyDrZB5UzwuOiTNGvryjrozoca6VNe6k0kPkK" \
//      -H "AUTHORIZATION: apikey rakeshneupane2045:0d822a1c96a05f99d722b13c91cd4b83" \
//      -X POST \
//      -F 'file=@/Users/surendrakandel/work/backend/invoice.jpeg' \
//      -F 'file_name=test.webp' \
//      https://api.veryfi.com/api/v8/partner/documents/
