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
	////read file os and make 
	file, err := os.Open("index.txt")
	errCheck(err)
	byt := make([]byte, 3)
	for {
		number, err := file.Read(byt)
		errCheck(err)
		if err != nil {
			break
		}
		fmt.Printf("%d byte ,content %s \n", number, byt)
	}
	

}
