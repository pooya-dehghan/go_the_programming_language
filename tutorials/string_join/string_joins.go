package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// benchmark
func main() {
	now := time.Now()
	fmt.Println(strings.Join(os.Args[1:], ","))
	after := time.Since(now)
	fmt.Printf("Time taken to join strings: %s\n", after)
}
