package main

import (
	"fmt"
	"time"
	"encoding/json"
)

type Person struct {
    Name        string `json:"Имя"`
    Email       string	`json:"Почта"`
    DateOfBirth time.Time `json:"-"`
}

func main() {
	time_1 := time.Date(1996, 6, 21, 12, 12, 0, 0, time.UTC)
	p := Person{Name: "Иван", Email: "ivan@gmail.com", DateOfBirth: time_1}
	fmt.Println(p)
	j, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(j))
	}
	var out Person
	err = json.Unmarshal(j, &out)
	fmt.Println(out)
}