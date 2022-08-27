package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
	//read file bufio  reade and scan
	files, err := os.Open("index.txt")
	errCheck(err)
	reader := bufio.NewReader(files)

	content, err := reader.Peek(6)
	errCheck(err)
	fmt.Printf("content : %s \n", content)
	//scaner bufio and split
	f, err := os.Open("test.txt") 
	errCheck(err)
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
