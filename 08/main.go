package main

import (
	"fmt"
	"time"
)

//	func main() {
//		//councurince  tred
//		go fmt.Println("world")
//		fmt.Println("hello")
//		time.Sleep(3 * time.Second)
//	}
func main() {
	//chanel
	ch := make(chan bool)
	go func() { //ano sing function
		fmt.Println("world2")
		time.Sleep(time.Second)
		ch <- true
	}()
	go TestChanel(ch)

	fmt.Println("hello")
	<-ch
}
func TestChanel(c chan bool) {
	fmt.Println("world")
	c <- true
}
