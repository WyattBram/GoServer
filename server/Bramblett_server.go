package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		termination := handleConnection(conn)

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

	closure := false
	for {

		buf := make([]byte, 512)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if string(buf[:3]) == "end" {
			defer conn.Close()
			closure = true
		}

		ReverseByteSlice(buf)

		_, err = conn.Write([]byte(buf))
		fmt.Printf("Sent: %s\n", buf)

		if err != nil {
			fmt.Println(err)
			return false
		}

		if closure {
			return true
		}

	}

}
