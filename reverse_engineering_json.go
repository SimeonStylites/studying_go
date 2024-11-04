package main

import (
	"fmt"
	"encoding/json"
)

type Response struct {
    Header		Headers `json:"header"`
    Data		[]Datas `json:"data"`
}

type Headers struct {
	Code		int `json:"code"`	
	Message		string `json:"message"`
}

type Datas struct {
	Type		string `json:"type"`
	Id			int `json:"id"`
	Attributes 	Attributes_struct `json:"attributes"`
}

type Attributes_struct struct {
	Email		string `json:"email"`
	Article_ids	[]int
}

func ReadResponse(rawResp string) (Response, error) {
	slice_bytes := []byte(rawResp)
	var r Response
	err := json.Unmarshal(slice_bytes, &r)
	fmt.Println(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return r, nil
}

func main() {
	jstr := `{
		"header": {
			"code": 0,
			"message": ""
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10,11,12]
			}
		}]
	}`
	json, _ := ReadResponse(jstr)
	fmt.Println(json)
}