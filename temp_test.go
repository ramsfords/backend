package main

import (
	"fmt"
	"html/template"
	"testing"
)

func TestTemplate(t *testing.T) {
	temp, err := template.ParseFiles("bol.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)
}
