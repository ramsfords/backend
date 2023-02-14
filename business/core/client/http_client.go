package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type HttpClient struct {
	*fasthttp.Client `json:"client"`
}

func New() *HttpClient {

	http.DefaultTransport.(*http.Transport).MaxIdleConns = 1024
	http.DefaultTransport.(*http.Transport).TLSHandshakeTimeout = time.Duration(0 * time.Second)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &HttpClient{
		Client: &fasthttp.Client{},
	}
}
