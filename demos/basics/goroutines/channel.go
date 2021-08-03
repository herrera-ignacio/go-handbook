package main

import (
	"fmt"
	"time"
)

// An int type channel passed in
func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c // Waits for value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) // A channel is created
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c) // Starts the goroutine

	// Passes ints into channel
	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
