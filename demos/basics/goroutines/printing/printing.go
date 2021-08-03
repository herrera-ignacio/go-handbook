package main

import (
	"fmt"
	"time"
)

// Function to execute as a goroutine
func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count() // Starts goroutine
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Second * 1)
	fmt.Println("Done")
}
