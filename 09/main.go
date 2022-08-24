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
	// *****call other func chen ****
	// fmt.Println(<-WaitChanel(5, 2))
	// *****select  ****
	select {
	case ch1 := <-WaitChanel(6, 9):
		fmt.Println(ch1)
	case ch2 := <-WaitChanel(8, 2):
		fmt.Println(ch2)
		// default:
		// 	fmt.Println(" all slow chen!")

	}

	/*
	*tick
	*after to chanel
	 */
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After((500 * time.Millisecond))
	for {
		select {
		case <-tick:
			fmt.Println("tick ...")

		case <-boom:
			fmt.Println("booom!!!")
			return
		default:
			fmt.Println("'*|*'")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
func WaitChanel(v, i int) chan int {
	chanel := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Millisecond)
		chanel <- v
	}()

	return chanel
}
