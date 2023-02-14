package S3

// func (S3 S3Client) UploadSampleImage(webSiteUrl string) (*manager.UploadOutput, error) {
// 	file, err := os.Open("order.jpg")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fileName := strings.Split(file.Name(), ".")
// 	key := webSiteUrl + "/images/" + fileName[0]
// 	res, err := S3.Upload(context.Background(), &s3.PutObjectInput{
// 		Bucket:             aws.String(S3.Config.S3AssetsBucket),
// 		Key:                &key,
// 		ContentType:        aws.String("image/jpg"),
// 		Body:               file,
// 		ContentDisposition: aws.String("inline"),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(res)
// 	return res, nil
// }
