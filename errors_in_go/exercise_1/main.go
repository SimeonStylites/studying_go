package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

type Myerror struct {
	numbers string
	long    string
	spaces  string
}

func (me Myerror) Error() string {
	return me.numbers + me.long + me.spaces
}

func MyCheck(s string) error {
	//длина
	me := Myerror{}
	if len(s) > 19 {
		me.long = "символов не меньше 20\n"
	}
	//цифры
out:
	for i := 0; i < 10; i++ {
		n := strconv.Itoa(i)
		for _, l := range s {
			if l == []rune(n)[0] {
				me.numbers = "в строке есть цифры\n"
				break out
			}
		}
	}
	//пробелы
	spaces := 0
	for _, l := range s {
		if l == rune(byte(' ')) {
			spaces++
		}
	}
	if spaces < 2 {
		me.spaces = "пробелов меньше 2\n"
	}
	return me
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
