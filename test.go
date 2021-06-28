package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Rest of the program...

	ch := make(chan int)

	go func() {
		fmt.Scanf("\n")
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("Exiting.")
		os.Exit(0)
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out, exiting.")
	}
}
