package S3

// func (S3 S3Client) UploadSampleFile(webSiteUrl string) (*manager.UploadOutput, error) {
// 	file, err := os.Open("sample.txt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fileName := strings.Split(file.Name(), ".")
// 	key := webSiteUrl + "/" + fileName[0]
// 	res, err := S3.Upload(context.Background(), &s3.PutObjectInput{
// 		Bucket:      aws.String(S3.Configs.S3AssetsBucket),
// 		Key:         &key,
// 		ContentType: aws.String("text/plain"),
// 		Body:        file,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(res)
// 	return res, nil
// }
