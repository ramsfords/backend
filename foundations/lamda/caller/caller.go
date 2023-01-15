package caller

import (
	"github.com/go-resty/resty/v2"
)

func New() *resty.Client {
	client := resty.New()
	return client

}
