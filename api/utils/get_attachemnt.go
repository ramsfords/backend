package utils

import (
	"fmt"
	"os"
)

func GetAttachment() []byte {
	data, err := os.ReadFile("/Users/surenl/projects/carriers/api/utils/bol0.pdf")
	fmt.Println(err)
	return data
}
