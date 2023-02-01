package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPdfGen(t *testing.T) {

	url := "https://selectpdf-selectpdf-html-to-pdf-v1.p.rapidapi.com/?key=e4e2a06e65mshfc6c51448317368p15d59ajsn8f332307530d&url=https%3A%2F%2Fwww.firstshipper.com%2Fbol%3FquoteId%3D23102%26businessId%3Dkandelsuren%40gmail.com%26auth%3DeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJyeWU1em9zZnh1aGMxbWoiLCJleHAiOjE2NzY0Mzg3ODcsImlkIjoiYndjand3aGRxenIwdGxsIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.pJvlyeVUCcwgZ-JaC7nPqA29GYFysy3G-Em5JQb-pZs%26bidId%3D23102-0"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "e4e2a06e65mshfc6c51448317368p15d59ajsn8f332307530d")
	req.Header.Add("X-RapidAPI-Host", "selectpdf-selectpdf-html-to-pdf-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	// err = ioutil.WriteFile("file.pdf", body, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(string(body))
}
