package auth

import (
	"fmt"
	"testing"

	"github.com/ramsfords/backend/configs"
)

func TestAuth(t *testing.T) {
	conf := configs.GetConfig()
	res := New(conf)
	fmt.Println(res)

}
