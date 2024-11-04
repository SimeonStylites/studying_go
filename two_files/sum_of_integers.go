package main

import ("fmt"; "math/rand")

func main () {
	//array A with numbers
	A := make(map[int]int, 100)
	for k := 0; k<100; k++ {
		A[k] = rand.Intn(100)
	}
	fmt.Println(A)
	//Number k = m+n
	k := 69
	//Finding indexes
	s := []int{}
	out:
	for i := 0; i < len(A); i++ {
		for j := i+1; j < len(A); j++ {
			if A[i]+A[j]==k {
				s = append(s,i)
				s = append(s,j)
				break out
			}
		}
	}
	fmt.Println(s)
	
	
	B := []int{}
	for k := 0; k<100; k++ {
		B = append(B,rand.Intn(100))
	}
	fmt.Println(B)
	fmt.Println(Find(B, 69))
}