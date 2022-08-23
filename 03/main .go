package main

import "fmt"

// func comp(num1 int,num2 int)(bool,int){
// if num1>num2 {
// 	return false,num1-num2
// }else if num2>num1{
// 	return false,num2-num1
// }
// return true ,0
// }
func comp(num1 int,num2 int)(isEqual bool,different int){
	if num1>num2 {
		isEqual=false
		different= num1-num2
		
	}else if num2>num1{
		isEqual=false
		different= num2-num1
	}else{
		isEqual=true
		different= 0
	}
	return
	}
func main() {
	fmt.Println("Hello world!")
	fmt.Println(comp(1,3))
	fmt.Println(comp(5,1))
	fmt.Println(comp(1,1))


}