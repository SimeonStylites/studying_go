package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    PrintAllFilesWithFilter(".", "ест")
}

func PrintAllFilesWithFilter(path string, filter string) {
    // получаем список всех элементов в папке (и файлов, и директорий)
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    //  проходим по списку
    for _, f := range files {
        // получаем имя элемента
        // filepath.Join — функция, которая собирает путь к элементу с разделителями
        filename := filepath.Join(path, f.Name())
        // печатаем имя элемента, если в нем есть filter
		if FilterInString(filename, filter) {
			fmt.Println(filename)
		}
        // если элемент — директория, то вызываем для него рекурсивно ту же функцию
        if f.IsDir() {
            PrintAllFilesWithFilter(filename, filter)
        }
    }
}

func FilterInString(str string, filter string) (ok bool) {
	ok = false
	// приводим к массиву рун, чтобы работать с символами utf
	str_rune := []rune(str)
	filter_rune := []rune(filter)
	for i, _ := range str_rune {
		for j, r_filt := range filter_rune {
			if i+j >= len(str_rune) || r_filt != str_rune[i+j] {
				break
			}
			if j == len(filter_rune)-1 {
				ok = true
				return
			}
		}
	}
	return
}