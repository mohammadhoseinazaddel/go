package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go fmt.Println("World.")
// 	fmt.Println("hello")
// 	time.Sleep(3 * time.Second)
// }
//***********************************
// func main() {
// 	ch := make(chan bool)
// 	//go TestChanel(ch)
// 	go func() {
// 		fmt.Println("World.")
// 		time.Sleep(time.Second)
// 		ch <- true
// 	}()
// 	<-ch
// 	fmt.Println("hello")

// }

// func TestChanel(c chan bool) {
// 	fmt.Println("World.")
// 	time.Sleep(time.Second)
// 	c <- true
// }
//***********************************
func main() {
	ch := make(chan int)
	go testChan(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
func testChan(c chan int) {
	i := 0
	for i <= 10 {
		c <- i
		i++
		time.Sleep(time.Second)
	}
	close(c)
}
