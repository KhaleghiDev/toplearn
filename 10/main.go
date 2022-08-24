package main

import (
	"fmt"
)

// method and interface
type testInter interface {
	sayHello()
	say(s string)
	increment()
	GetValuePrint() int
}
type testruct struct {
	i int
}

func newTStract() testInter {
	return new(testruct)
}
func newTStractWithv(v int) testInter {
	return &testruct{i: v}
}
func (ts *testruct) sayHello() {
	fmt.Print("Hello \n")
}
func (ts *testruct) say(input string) {
	fmt.Println(input)
}
func (ts *testruct) increment() {
	ts.i++
}
func (ts *testruct) GetValuePrint() int {
	return ts.i
}

type emmbddingStract struct {
	*testruct
}

func testEmpaty(v interface{}) {
	// fmt.Println(v)
	// if i,ok:=v.(int);ok {
	// 	fmt.Println("I Am int my value is :",i)
	// }else{
	// 	fmt.Println("I Am not int :")
	// }
	switch val := v.(type) {
	case int:
		fmt.Println("int", val)
	case float32:
		fmt.Println("flout", val)
	case string:
		fmt.Println("string", val)
	case bool:
		fmt.Println("bool", val)
	default:
		fmt.Println("not int,float32,bool , string", val)
	}
}

func main() {
	var test testInter
	test = newTStract() //&testruct{}
	test = newTStractWithv(5)
	test.sayHello()
	test.say("hi how are you !")
	test.increment()
	test.increment()
	test.increment()
	fmt.Println(test.GetValuePrint())

	te := emmbddingStract{testruct: &testruct{i: 800}}
	te.say("hahahaaaaa...")
	te.increment()
	te.increment()
	te.increment()
	te.increment()
	te.increment()
	te.increment()
	fmt.Println(te.GetValuePrint())
	testEmpaty(5)
	testEmpaty("string")
	testEmpaty(te)
	testEmpaty(&te)

}
