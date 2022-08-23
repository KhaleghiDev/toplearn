package main

import (
	"fmt"
)

//	func comp(num1 int, num2 int) (isEqual bool, diff int) {
//		switch {
//		case num1 > num2:
//			isEqual = false
//			diff = num1 - num2
//		case num2 > num1:
//			isEqual = false
//			diff = num2 - num1
//			fallthrough
//		default:
//			isEqual = true
//			diff = num2 - num1
//		}
//		return
//	}
func main() {
	// 	var x int = 33
	// 	y := "ali"
	// 	fmt.Println("add ==>", x, y)
	// 	fmt.Println(comp(20, 11))
	// 	fmt.Println(comp(9, 11))
	// 	fmt.Println(comp(20, 20))
	// switch x{
	// case 30:
	// 	fmt.Println("30")
	// case 32:
	// 	fmt.Println(32)
	// case 33:
	// 	fmt.Println(33)
	// default:
	// 	fmt.Println(" end switch")
	// }

	var i int
	fmt.Println("enter namber conter:")
	fmt.Scan(&i)
	human(i)
}

func human(naumber int) {
	for i := 1; i <= naumber; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("|*_*|")
		}
		fmt.Printf("\n")
	}
}
