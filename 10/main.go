package main

import "fmt"

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
func newTStract() testInter{
	return new(testruct)
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

func main() {
	var test testInter
	test= newTStract()   //&testruct{}
	test.sayHello()
	test.say("hi how are you !")
	test.increment()
	test.increment()
	test.increment()
	fmt.Println(test.GetValuePrint())

}
