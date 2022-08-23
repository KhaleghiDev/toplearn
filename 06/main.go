package main

import "fmt"

//struct
type person struct{
	name string
	famile string
	age  int
}

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
	fmt.Println(ok, "value :", v)

	if v, ok := m["secent"]; ok {
		fmt.Println("secent: ", v)
	}
	//struct

	reza := person{age:50,famile: "razavi",name:   "reza"}
	fmt.Print(reza)

	people:=map[string]person{
		"ali":person{age:20,famile: "razavi",name: "ali"},
		"mohammad":person{age:50,famile: "razavi",name:   "mohammad"},
	}
	fmt.Print("peple : ",people)
	fmt.Print("ali : ",people["ali"])
	fmt.Print("name : ",people["ali"].name)
}
	
	


