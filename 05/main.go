package main

import "fmt"

func main() {
	// var a[5]int=[5]int{1,2,3,4,5}//arry
	a := [5]int{1, 2, 3, 4, 5} //arry

	MySlice := []int{6, 7, 8, 9, 10} //slice
	MySlice = append(MySlice, 11, 12, 13, 14, 15)
	fmt.Println(a, MySlice)
	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5
	fmt.Println(s)
	s1 := s[1:4]
	fmt.Println(s1) //len(s1) , cap(s1)
	fmt.Println("len :  ", len(s1),"cap :",cap(s1))
	s2:=make([]int,2)
	copy(s2,s[2:4])
	fmt.Println(s2)
	fmt.Println("len :  ", len(s2),"cap :",cap(s2))
}
