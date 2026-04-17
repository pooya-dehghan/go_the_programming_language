package main

import (
	"crypto/sha256"
	"fmt"
)

func countHash(hashOne, hashTwo [32]byte) int {
	counter := 0
	for i := range hashOne {
		if hashOne[i] == hashTwo[i] {
			fmt.Printf("i %v :", i)
			fmt.Printf("hashOne[i] %v ", hashOne[i])
			fmt.Printf("hashTwo[i] %v ", hashTwo[i])
			fmt.Printf("\n=================================\n")
			counter++
		}
	}

	return counter
}

func main() {
	sh1 := sha256.Sum256([]byte("x"))
	sh2 := sha256.Sum256([]byte("x"))

	fmt.Printf("sh1 %x\n", sh1)
	fmt.Printf("sh2 %x\n", sh2)
	fmt.Printf("length %v\n", len(sh2))

	fmt.Printf("counter : %v\n", countHash(sh1, sh2))
}
