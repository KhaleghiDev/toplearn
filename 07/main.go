package main

import "fmt"

func main() {
	//pointer

	var PI *int //memory address of int
	var I int = 3
	PI = &I
	fmt.Println(PI)
	incriment(PI)
	fmt.Println(*PI )
	// incriment(&I)
	fmt.Println(&I)

}
func incriment(p *int){
	*p++
}