package main

import (
	"fmt"
	"os"
)

// func main() {
// 	//read and write in file package os
// 	file,err:=os.Create("index.txt")
// 	if err != nil {
// 		fmt.Println("err:" ,err)
// 		return
// 	}
// 	len,err:=file.WriteString("hi file me")
// 	if err != nil {
// 		fmt.Println("err:" ,err)
// 		return
// 	}
// 	fmt.Println(len,"character")
// 	file.Close()

// }
// func main() {
// 	//read and write in file package os
// 	file,err:=os.Create("index")
// 	if err != nil {
// 		fmt.Println("err:" ,err)
// 		return
// 	}
// 	by:=[]byte{0104 ,122, 164, 154, 200, 189, 58,175,99,58,65,66,101}
// 	len,err:=file.Write(by)
// 	if err != nil {
// 		fmt.Println("err:" ,err)
// 		file.Close()
// 		return
// 	}
// 	fmt.Println(len,"character")
// 	file.Close()

// }
//*****************************************************
func main() {
	//read and write in file package os
	file,err:=os.Create("test.txt")
	if err != nil {
		fmt.Println("err:" ,err)
		return
	}
	s:=[]string{"my name is abolfazl","I am programmer golang ", "I love you golang" ,"I am from iran ","I would like to travel to Canada"}
	for _,value:= range s {
		_,err:=fmt.Fprintln(file,value)
		if err != nil {
			fmt.Println("err:" ,err)
			return
		}
	}
	
	err=file.Close()
	if err != nil {
		fmt.Println("err:" ,err)
		return
	}

}