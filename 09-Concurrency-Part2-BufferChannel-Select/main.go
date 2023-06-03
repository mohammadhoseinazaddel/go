package main

import (
	"fmt"
	"time"
)

func main() {
	// buffer := make(chan string, 2)
	// buffer <- "Hello"
	// buffer <- "World"

	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)
	//*********************************************
	// fmt.Println(<-WaitChannel(5, 2))
	//*********************************************

	// select {
	// case v1 := <-WaitChannel(5, 2):
	// 	fmt.Println(v1)
	// case v2 := <-WaitChannel(6, 2):
	// 	fmt.Println(v2)
	// 	// default:
	// 	// 	fmt.Println("All Channel Are Slow.")
	// }

	//*********************************************

	tick := time.Tick(100 * time.Millisecond)

	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOOM!!!")
			return
		default:
			fmt.Println("*|*")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

// func WaitChannel(v, i int) chan int {
// 	channel := make(chan int)

// 	go func() {
// 		time.Sleep(time.Duration(i) * time.Second)
// 		channel <- v
// 	}()

// 	return channel
// }
