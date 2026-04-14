package main

import (
	"fmt"
	"os"
	"time"
)

// Echo program
func main() {
	now := time.Now()
	var holder, seperator string

	for arg := range os.Args {
		fmt.Printf("arg is : %v \n", arg)

		fmt.Printf("argurment value is %s \n", os.Args[arg])

		holder += os.Args[arg] + seperator
		seperator = " "
	}
	fmt.Printf("============================")

	fmt.Printf("holder value is %s \n", holder)

	after := time.Since(now)

	fmt.Printf("Time taken for operation : %s\n", after)
}
