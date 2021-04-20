package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Result struct {
	headers Headers
	url string
}

type Headers struct {
	Accept string
}

func main() {
	client := resty.New()
	resp, err := client.R().SetResult(&Result{}).Get("https://httpbin.org/get")
	if err == nil {
		fmt.Println("", resp.Result())
	}
}