package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	// byt, err := ioutil.ReadFile("test.txt")
	// errCheck(err)
	// fmt.Println(string(byt))

	// file, err := os.Open("test2.txt")
	// errCheck(err)
	// byt2 := make([]byte, 1)
	// for {
	// 	number, err := file.Read(byt2)
	// 	errCheck(err)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("%d byte , Content : %s \n", number, string(byt2))
	// }

	// file, err := os.Open("test.txt")
	// errCheck(err)
	// reader := bufio.NewReader(file)

	// content, err := reader.Peek(3)
	// errCheck(err)
	// fmt.Printf(" Content : %s \n", string(content))

	// file, err := os.Open("test.txt")
	// errCheck(err)

	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	file, err := os.Open("test.txt")
	errCheck(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("///// ", scanner.Text())
	}

}
