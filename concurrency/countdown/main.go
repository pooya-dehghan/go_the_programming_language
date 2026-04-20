package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Printf("counting down , starting now\n")
	tick := time.Tick(200 * time.Millisecond)

	for i := 10; i > 0; i-- {
		select {
		case <-tick:
			fmt.Printf("tick %v\n", i)
		case <-abort:
			fmt.Printf("aborting \n")
			return
		}
	}

	fmt.Printf("aborted, good by dear your \n")
}
