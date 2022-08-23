package main

import "fmt"

func main() {
	//pointer

	var PI *int //memory address of int
	var I int = 3
	PI = &I
	fmt.Println(PI)
	incriment(PI)
	fmt.Println(*PI)
	// incriment(&I)
	fmt.Println(&I)
	testDefer()
}
func incriment(p *int) {
	*p++
}
func testDefer() {
	defer fmt.Println("wold")
	defer fmt.Println("wold2")
	fmt.Println("Hello")
}
