package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Result struct {
	Headers Headers `json:"headers"`
	Url string `json:"url"`
}

type Headers struct {
	Accept string `json:"Accept"`
}

type UinResp struct {
	Code int     `json:"code"`
	Data UinData `json:"data"`
	Msg  string  `json:"msg"`
}

type UinData struct {
	CreateAt         int64  `json:"createAt"`
	Id               int    `json:"id"`
	ObsPlanProductId int    `json:"obsPlanProductId"`
	ObsProductId     int    `json:"obsProductId"`
	ProjectId        int    `json:"projectId"`
	UpdateAt         int64  `json:"updateAt"`
	YjnTiUni         string `json:"yjnTiUni"`
}

func main() {
	client := resty.New()
	result := UinResp{}
	resp, err := client.R().SetResult(&result).Get("http://127.0.0.1:1333/api/v2/project/obs/33/getProjectObsBind")
		fmt.Println(result)
	if err == nil {
		fmt.Println(resp)
	}
		//fmt.Println(resp)
}