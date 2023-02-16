package auth

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/ramsfords/backend/configs"
)

func TestLogin(t *testing.T) {
	code := "e2b11864-a652-45e5-97e8-e913a22ad7a4"
	config := configs.GetConfig()
	awsConf := config.GetAwsConfig()
	baseAuth := base64.RawStdEncoding.EncodeToString([]byte(awsConf.CognitoClientID + ":" + awsConf.CognitoClientSecret))
	auth := Auth{
		cognitoRegion:       awsConf.CognitoRegion,
		cognitoUserPoolID:   awsConf.CognitoUserPoolID,
		baseAuth:            baseAuth,
		cognitoClientSecret: awsConf.CognitoClientSecret,
		cognitoClientID:     awsConf.CognitoClientID,
		redirectUrl:         awsConf.CognitoRedirectUri,
		authUrl:             awsConf.CognitoUrl,
	}
	client := &http.Client{}
	encodedBody := fmt.Sprintf("code=%s&grant_type=authorization_code&redirect_uri=%s", code, auth.redirectUrl)
	var data = strings.NewReader(encodedBody)
	req, err := http.NewRequest("POST", "https://firstshipper-dev.auth.us-west-1.amazoncognito.com/oauth2/token", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Basic "+auth.baseAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
func TestRefreshtoken(t *testing.T) {
	refreshToken := "eyJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiUlNBLU9BRVAifQ.bI9Ge_EqpQ2Q-ZlIWn0eRDWIlg_abJOzzGaD7kKNrvae7owlnbHLVTANpZ8QUxXw1fM9smQuqk-G1h329TEmfk_o7qBB3nT9uGQez_rktvYEgm3usF6b9-1oYaCqIo7bgbdO6Bx0slgQzp4oqsUXKItomhyGRNIO-BgI7ZvwdWeCakLISXXmTMe0OPCFSuJgFg1nRK_5ax2LXyO35UiEsKEOefDxopFOzq5vSD3RXuhwJgFTHuiakhlS0Y7m6grswmKotW3AIcrLIMgVK36SoIZBfjBU45ijryPcIndOUTIqG4VdirPTLDNh7AHpDqDSh1WGUMMeMgvkro9sap5Yog.QKEwdt_6-G1KH8H-.Y3BSvi80IN4XvxPWqAQj0FT5--tsb5QVgX59yvcpeqUhFwY5oi0pj85Less3jf-p3lMmLiDRZj4C9ll2PL9lkkUcDJyYsctlqwkc7nQy0mCXXiartFvcOlMBKxiQFPM6caov6lCOlrj_wNDlDYbjOIoRE3q6a0a1N1SglVxD2WexoFCuE21_khe6DDFMeMmOvT1CEhm26DvbDHY0giq9hEpg-4CYmXSykbF9-i7mUy6K86k_k6QjbcrbWUKT8KraCL6zZ1kQ2zGGGTy0MhMHGEQypWsE9ld4Ac0f0t10jYJ237M-DtIAhHCKP8xrYv-yQQSen2z972WTXhy52iunYonbU7RKoVnIY-4-HNOJeD59HRiOjwnYsfOD7PbB9kMj9zmOfl8uewrvbeKKkNer5lyN7To9YP7MbqfiGY3gzh3rmAh01dgIqMFR_j9NNOBy_x5MXu94J0xHb3tpHK7Exi87aIuJ2WC4_T0I1Cv5Rw_bSom2GOl27cMktrSGIKVV-VvuzvUviZ9DBkTT-vl7fcyaaoL2TFBHMl8pDKtaMynnfsweixW9KFrO9c88UDIZjZBpPYqlaqZXLU-hs_IiLmfEObv5B9Aqj62NuKA2HUjIcmZjCeuIdjYzuE6WanKLXTdpBS9MKblXHgCXaVoWq8PCMK1SLZpjCbG7w0zuSi94GZeyPeJfygka-jngsdBOTFyZPUlHKoc50feW2Q06HMURQ5A8WbcOiFYKUc7XuzZqeIXG6cPitoqr78cUSC8IckS3E1RJam7cY1mKRsS3JfUn5zUeJT_vr2MlwT6cv-oAFFjdhjvSh7FADUuFmrWDPAusHVqsVY_FLM5Rxx2jLCKgt96VqKJHlpujc3L8KweS1sKQEOUKqfOK0gv0l5046Vf-Ps71Xio-3OHiodqgGLgU0m939hM-sKzDwKHBCK_VWzALVYEYZ9pDwY34jfaj52XwA3WyIz5BX919dYeE6ZtNkeO6aMfc0fF_a3Ev3rKNx3jeiqM_v63iu9NcxleOXOTjXt1qPbUJ41q2eT9yvuU2lX8ek05DDWm74JRJzGNaA3pnF7YnpgWWU0nZdiVbdMSr0_tMHoZvJeSxrLQmcAvhH8aLiTZLHDI0z4_o5YU2H2YZVrlLMsfGaxChi74hfYyw21_wqRSl0yKjdtCIf0fSWYXo4qAdTTNlPnBV62kSmcUuFuvy7Mb-yNt3EF1L9xKX2M6NytYfSx9mK_n3Zic9VPkJIvjMdYM.a-_yElBD42D8sH5rO1ft8g"
	config := configs.GetConfig()
	awsConf := config.GetAwsConfig()
	baseAuth := base64.RawStdEncoding.EncodeToString([]byte(awsConf.CognitoClientID + ":" + awsConf.CognitoClientSecret))
	auth := Auth{
		cognitoRegion:       awsConf.CognitoRegion,
		cognitoUserPoolID:   awsConf.CognitoUserPoolID,
		baseAuth:            baseAuth,
		cognitoClientSecret: awsConf.CognitoClientSecret,
		cognitoClientID:     awsConf.CognitoClientID,
		redirectUrl:         awsConf.CognitoRedirectUri,
		authUrl:             awsConf.CognitoUrl,
	}
	client := &http.Client{}
	encodedBody := fmt.Sprintf("refresh_token=%s&grant_type=refresh_token&client_id=%s", refreshToken, auth.cognitoClientID)
	var data = strings.NewReader(encodedBody)
	req, err := http.NewRequest("POST", "https://firstshipper-dev.auth.us-west-1.amazoncognito.com/oauth2/token", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Basic "+auth.baseAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
