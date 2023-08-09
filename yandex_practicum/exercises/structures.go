package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json: "Имя"`
	Email       string    `json: "Почта"`
	DateOfBirth time.Time `json:"-"`
}

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
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

func main() {
	fmt.Println("Hi")
	test2()
}

func test1() {
	p := Person{Name: "Алекс", Email: "alex@yandex.ru"}

	str, _ := json.Marshal(p)

	fmt.Println(string(str))
}


func test2() {
	resp := Response{}
		
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}