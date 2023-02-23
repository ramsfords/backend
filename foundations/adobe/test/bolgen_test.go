package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/adobe"
	v1 "github.com/ramsfords/types_gen/v1"
)

var adobes *adobe.Adobe

func init() {
	conf := configs.GetConfig()
	seClient := S3.New(conf)
	adobes = adobe.NewAdobe(seClient)
	adobes.GetToken()
}
func TestBolGen(t *testing.T) {
	// adobe.UrlToPdf("30007-0")
	adobes.UploadBOlTOS3("https://dcplatformstorageservice-prod-us-east-1.s3-accelerate.amazonaws.com/1a0995378e964e85a260ce3e98f5e6b7_EEDA264963D959D90A495EEC%40techacct.adobe.com/d9e785d9-4262-4fd8-8bb7-500d5975531e?X-Amz-Security-Token=FwoGZXIvYXdzEHgaDEg1e4HP7%2BeBX2ncGiLUAeX%2BYGT%2BES%2FRjdO%2FhHMLXROcuYsUnNbF9WrjMFC6OV8HRJsu3wnufm%2BvUHndsetuMwrr1vsCmJUqPMXCLjrodDCYiOXCEK01jc8IrW7SLVRo8aLS7W0NUv15tLcfADAiaX17Ccp0SjtjbZjFsksLKnvnVKxIdYr8sG5X6DAv7rodByYjh74HY9G9p5vnaoBWwE4Wfe4K%2F9y9oJ1H1LbxnwxxYIuvhS671fk%2FllRGVoyT9Z5iuHrzkMQewP45Rl5L2pYx7cTTSYXbb4oYKygmMxrmQxq8KJru8J4GMi0yaKGSvDt6SEx%2B5uIIsi%2BuKcSdi9p3tR%2FP5wOuGv97bz7i%2BkIynbTFbhBNprA%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20230202T222505Z&X-Amz-SignedHeaders=host&X-Amz-Expires=3599&X-Amz-Credential=ASIAWD2N7EVPIVYDQ7NM%2F20230202%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=385e674d8b66448be5877aed387a8adda359688cca6d6181e793d18eaea4fa45", &v1.Bid{}, "")
}

func TestTokenGen(t *testing.T) {
	urls := "https://ims-na1.adobelogin.com/ims/exchange/jwt/"
	form := url.Values{}
	form.Add("client_id", "1a0995378e964e85a260ce3e98f5e6b7")
	form.Add("client_secret", "p8e-PG3RgfPMq_b3t_KT9ZhLbPJJ1LASyPMA")
	form.Add("jwt_token", "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE2NzU0NjE4NzMsImlzcyI6IjhGNzk2MUFFNjJGQTcyNEMwQTQ5NUM5M0BBZG9iZU9yZyIsInN1YiI6IkVFREEyNjQ5NjNEOTU5RDkwQTQ5NUVFQ0B0ZWNoYWNjdC5hZG9iZS5jb20iLCJodHRwczovL2ltcy1uYTEuYWRvYmVsb2dpbi5jb20vcy9lbnRfZG9jdW1lbnRjbG91ZF9zZGsiOnRydWUsImF1ZCI6Imh0dHBzOi8vaW1zLW5hMS5hZG9iZWxvZ2luLmNvbS9jLzFhMDk5NTM3OGU5NjRlODVhMjYwY2UzZTk4ZjVlNmI3In0.wKaTWPW7iZsNgzEsNdWDZuKaJBIfMPJrmqStxJM8CDQe1nZOEK8IWQycDrcTfd3KNAEuj2qvMUcCEY19L8cSUWBdzhLMTnH1wKoef_6SomgPzmUTgTw9_-0twPqaTzbkt_Ckzs-ajjEt-zHEEBQgI_sz1cNgvf9QawGuLEhKnV0l98fh3SHB5EirxgYbqZ4A_wBAxioDCR7mi86RN9kmGwGXdYeDWZFnXFPbDa6k5WmGQGPXL6cfmvrph4NkmYgWEmg3hwQgNVbpl9Ef4fgbUubV3CQ496-cHGCLGawbOP4qfWdV_rv5zN0VuKlrTPEK7UNIFOhe0NY6VDLuwe8LtA")

	res, err := http.PostForm(urls, form)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	token := &TokenRes{}
	err = json.Unmarshal(body, token)
	if err != nil {
		fmt.Println(err)
	}
	token.AccessToken = "Bearer " + token.AccessToken
	fmt.Println(res)
	// err = ioutil.WriteFile("file.pdf", body, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(string(body))

}

type TokenRes struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
