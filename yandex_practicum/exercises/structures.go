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
		Code: int,
		Message: string
	}
}

func main() {
	fmt.Println("Hi")
	test1()
}

func test1() {
	p := Person{Name: "Алекс", Email: "alex@yandex.ru"}

	str, _ := json.Marshal(p)

	fmt.Println(string(str))
}
