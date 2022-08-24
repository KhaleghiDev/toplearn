package main

import (
	"fmt"
	"time"
)

func main() {
	//concarence -> chanelbuffer select
	// buffer := make(chan string, 2)
	// buffer <- "hello"
	// buffer <- "world"
	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)
	fmt.Println(<-WaitChanel(5, 2))

}
func WaitChanel(v, i int) chan int {
	chanel := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Millisecond)
		chanel <- v
	}()

	return chanel
}
