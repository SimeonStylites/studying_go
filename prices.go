package main

import "fmt"

func main () {
	prices := make(map[string]int)
	prices["хлеб"] = 50
	prices["молоко"] = 100
	prices["масло"] = 200
	prices["колбаса"] = 500
	prices["соль"] = 20
	prices["огурцы"] = 200
	prices["сыр"] = 600
	prices["ветчина"] = 700
	prices["буженина"] = 900
	prices["помидоры"] = 250
	prices["рыба"] = 300
	prices["хамон"] = 1500
	fmt.Println(prices)
	
	//деликатесы
	for k, v := range prices {
		if v>500 {
			fmt.Println(k)
		}
	}
	
	//заказ
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	sum := 0
	for _, v := range order {
		sum += prices[v] 
	}
	fmt.Println(sum)
}