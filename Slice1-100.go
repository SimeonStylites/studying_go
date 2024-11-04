package main

import "fmt"

func main () {
	s := make([]int, 100, 100)
	for i := range s {
		s[i] = i+1
	}
	s = append(s[:10], s[90:]...)
	fmt.Println(s)
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
	
	//строка
	input := "The quick brown 狐 jumped over the lazy 犬"
    // Get Unicode code points. 
    n := 0
    // создаём слайс рун 
    runes := make([]rune, len(input))
    // добавляем руны в слайс
    for _, r := range input {
        runes[n] = r
        n++
    }
    // убираем лишние нулевые руны
    runes = runes[0:n]
    // разворачиваем
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    // преобразуем в строку UTF-8. 
    output := string(runes)
    fmt.Println(output)
}