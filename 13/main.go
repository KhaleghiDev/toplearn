package main

import (
	"fmt"
	"os"
)

func main() {
	//read and write in file package os
	file,err:=os.Create("index.txt")
	if err != nil {
		fmt.Println("err:" ,err)
		return
	}
	len,err:=file.WriteString("hi file me")
	if err != nil {
		fmt.Println("err:" ,err)
		return
	}
	fmt.Println(len,"character")
	file.Close()

}