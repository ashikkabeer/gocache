package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Hello, Welcome to GoCache!")

	// Start a TCP server on port 6379
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// start receiving connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	// create an infinite loop to handle requests and respond to them
	for {
		buf := make([]byte, 1024)

		// read messages from client
		_, err = conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}


}