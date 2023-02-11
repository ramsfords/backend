package S3

// func (S3 S3Client) UploadImageFromUrl(menuloomStaticImageFolderUrl string, items []string) (interface{}, error) {
// 	var errs error
// 	// get the image from the url
// 	for index, resource := range items {
// 		validUrl, err := url.ParseRequestURI(resource)
// 		if err != nil {
// 			errs = err
// 			continue
// 		}
// 		fmt.Println(validUrl)
// 		// filename := path.Base(URL)
// 		response, err := http.Get(resource)
// 		if err != nil {
// 			errs = err
// 			continue
// 		}

// 		if response.StatusCode != 200 {
// 			fmt.Println("Received non 200 response code")
// 		}
// 		// // Create a empty file
// 		// file, err := os.Create(filename)
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// }
// 		// defer file.Close()

// 		// Write the bytes to the fiel
// 		//readData, err := io.ReadAll(response.Request.Body)
// 		_, err = S3.Upload(context.Background(), &s3.PutObjectInput{
// 			Bucket:             S3.Config.S3Buck,
// 			Key:                aws.String(menuloomStaticImageFolderUrl + "/" + item.Name),
// 			ContentType:        aws.String(response.Header.Get("Content-Type")),
// 			ContentDisposition: aws.String("inline"),
// 			Body:               response.Body,
// 		})
// 		if err != nil {
// 			errs = err
// 			continue
// 		}
// 		item.Images = append(item.Images, &proto_gen.Image{})
// 		item.Images[index] = &proto_gen.Image{}
// 	}
// 	if errs != nil {
// 		return item, errs
// 	}
// 	return item, nil

// }
