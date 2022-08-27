package main

import (
	"fmt"
	"io/ioutil"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println("Error :", e)
		return
	}
}
func main() {
	//read file ioutil
	by, err := ioutil.ReadFile("test.txt") //reade content file
	errCheck(err)
	fmt.Println(string(by))

}
