package cognito

import (
	"context"
	"testing"

	"github.com/ramsfords/backend/configs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestCongito(t *testing.T) {
	config := configs.GetConfig()
	cognito, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}
	user := &v1.User{
		Email:       "kandelsuren@gmail.com",
		Password:    "Ferina@1234",
		Name:        "Suren Kandel",
		PhoneNumber: "+9779841000000",
	}
	outPut, err := cognito.CreateUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(outPut)
}
