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
	testPanic()
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
func testPanic(){

	defer func ()  {//try cache handele error preventing crash 
		if err:=recover();err!=nil{
			fmt.Println("panic handeles :")
			fmt.Println(err)
		}
	}()
	fmt.Println("Hello")
	panic("hahahahaa i am panic")
	
}
