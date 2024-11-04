package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}

	reqRaw, _ := json.Marshal(req)
	fmt.Println(string(reqRaw))
}