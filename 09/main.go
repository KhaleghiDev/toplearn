package main

import "fmt"

func main() {
	//concarence -> chanelbuffer select
	buffer := make(chan string, 2)
	buffer <- "hello"
	buffer <- "world"
	fmt.Println(<- buffer)
	fmt.Println(<- buffer)

}