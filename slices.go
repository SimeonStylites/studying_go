package main

import "fmt"

func main () {
	var mySlice []int
	fmt.Println(cap(mySlice))
	
	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
    workDaysSlice := weekTempArr[:5]
    weekendSlice := weekTempArr[5:]
    fromTuesdayToThursDaySlice := weekTempArr[1:4] 
    weekTempSlice := weekTempArr[:]
    
    fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7

	workDaysSlice[0] = 0
	fmt.Println(workDaysSlice) // [0 2 3 4 5]
	
	//strange mechanics
	a := []int{1, 2, 3, 4}
    b := a[2:3]   // b = [3]
    b = append(b, 7)
    fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
    fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
    b = append(b, 8, 9, 10)
    b[0] = 11
    fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
    fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6
}