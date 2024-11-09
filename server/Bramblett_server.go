package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them

	for {

		// accepts connections and handles error
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handles connection
		termination := handleConnection(conn)

		// if the connection sent "end" it will terminate server

		if termination == true {
			break
		}

	}
}

func ReverseByteSlice(b []byte) {
	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

func handleConnection(conn net.Conn) bool {

	// Close the connection when we're done
	closure := false
	for {

		//makes buffer
		buf := make([]byte, 512)

		//connects and handles error
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return false
		}

		// checks if it recieved "end"
		if string(buf[:3]) == "end" {
			defer conn.Close()
			closure = true
		}

		// reverses buf
		ReverseByteSlice(buf)

		//sends buf and handles error
		_, err = conn.Write([]byte(buf))
		fmt.Printf("Sent: %s\n", buf)

		if err != nil {
			fmt.Println(err)
			return false
		}

		// closes if recieved "end"
		if closure {
			return true
		}

	}

}
