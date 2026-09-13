package main

import (
	"encoding/json"
	"fmt"
)

const RawResp string = `
{
    "header": {
        "code": 0,
        "message": "OK"
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
} 
`

type Response struct {
	Header Header `json:"header"`
	Data   []Item `json:"data,omitempty"`
}

type Header struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Item struct {
	Type       string     `json:"type"`
	Id         int        `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Email      string `json:"email"`
	ArticleIds [3]int `json:"article_ids"`
}

func ReadResponse(rawResp string) (Response, error) {
	var resp Response

	err := json.Unmarshal([]byte(rawResp), &resp)

	return resp, err
}

func main() {
	result, error := ReadResponse(RawResp)
	if error != nil {
		fmt.Println("Error unmarshaling JSON:", error)
	}

	fmt.Printf("DTO %v\n", result.Header.Message)

	for _, item := range result.Data {
		fmt.Printf("DTO %v\n", item.Attributes.Email)
	}

}
