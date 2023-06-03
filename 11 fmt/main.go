package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Print("hello \n")
	fmt.Print("hello")
	fmt.Print("hello \n")
	st := struct {
		i int
		f float32
	}{i: 5, f: 5.6}
	fmt.Printf("hello %v \n %+v \n %#v \n %T \n", st, st, st, st)
	fmt.Printf("valu of i is %d and f is %9.2f", st.i, st.f)

	var s string
	var l int
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &l)
	fmt.Printf("\n %s %d", s, l)
}
