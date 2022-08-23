package main

import "fmt"

func main() {
	//map ans stuct
	var m = make(map[string]int)
	maps := make(map[string]int)
	m["first"] = 1
	m["secent"] = 2
	maps["one"] = 1
	maps["tow"] = 2

	fmt.Println(m, maps["tow "])
	// _, ok := m["first"]// output =>true 
	v, ok := m["first"]
	fmt.Println(ok,"value :", v)

	if v,ok:=m["secent"];ok {
		fmt.Println("secent: " ,v)
	}

}
