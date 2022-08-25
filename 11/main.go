package main

import "fmt"

func main() {
	st := struct {
		i int
		f float32
	}{i: 5, f: 8.5}
	name := "ahmad"
	old := 20
	// cheach fmt
	fmt.Println("hello .")
	fmt.Print("...\n ", st.i)
	fmt.Printf("\n %s of is %d  old \n", name, old)
	fmt.Printf("\n%v ,\n %+v,\n %#v  ,\n %T ,%t \n", st, st, st, st, st)
	fmt.Printf("hi of i is %d and %2.f", st.i, st.f)
	fmt.Println("\n enter namber plese : ")
	var input int
	i, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Println("input :", i)
}
