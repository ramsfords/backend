package zoho

import (
	"fmt"
	"testing"

	"github.com/ramsfords/backend/configs"
)

func TestAuth(t *testing.T) {
	conf := configs.GetConfig()
	res, err := New(conf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)

}
