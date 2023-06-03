package main

import "fmt"

type testInter interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
}

type tstruct struct {
	i int
}

func NewTStruct() testInter {
	return new(tstruct)
}

func NewTStructWithV(v int) testInter {
	return &tstruct{i: v}
}

func (ts *tstruct) SayHello() {
	fmt.Println("Hello!")
}

func (ts *tstruct) Say(s string) {
	fmt.Println(s)
}

func (ts *tstruct) Increment() {
	ts.i++
}

func (ts *tstruct) GetInternalValue() int {
	return ts.i
}

type embeddingStruct struct {
	*tstruct
}

func testEmpty(v interface{}) {

	// if i, ok := v.(int); ok { //Type Assertion
	// 	fmt.Println("I am int my value is :", i)
	// } else {
	// 	fmt.Println("I am not int")
	// }

	switch val := v.(type) { //Type Switch
	case int:
		fmt.Println("I am int my value is :", val)
	case string:
		fmt.Println("I am string my value is :", val)
	default:
		fmt.Println("I am not int and string my value is :", val)
	}
}

func main() {
	var test testInter
	test = NewTStructWithV(5) //NewTStruct() //&tstruct{} //new(tstruct)

	test.SayHello()
	test.Say("Hello Again!!!")

	test.Increment()
	test.Increment()
	test.Increment()

	fmt.Println(test.GetInternalValue())

	te := embeddingStruct{tstruct: &tstruct{i: 80}}
	te.Say("hahaha ....")
	te.Increment()
	te.Increment()
	te.Increment()
	te.Increment()
	fmt.Println(te.GetInternalValue())

	testEmpty(5)
	testEmpty("String")
	testEmpty(te)
}
