package main

import (
	"bufio"
	"fmt"
)

type ByteCounter struct {
	words int
	lines int
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	ad, to, err := bufio.ScanLines(p, true)
	if err != nil {
		return 0, err
	}
	c.lines = ad

	fmt.Printf("token is %v\n", to)

	adv, _, err := bufio.ScanWords(p, true)

	if err != nil {
		return 0, err
	}

	c.words = adv

	return len(p), nil
}

func main() {
	var c ByteCounter
	_, err := c.Write([]byte("Hello pooya dehghan!"))
	if err != nil {
		fmt.Errorf("err", err)
	}

	fmt.Fprintf(&c, "Bytes written: %s\n", "forgot you")

	fmt.Printf("the size %+v\n", c)
}
