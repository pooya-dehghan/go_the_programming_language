package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	_, err := io.WriteString(conn, "hi\n")

	if err != nil {
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("err: ", err)
		}
		handleConn(conn)
	}
}
