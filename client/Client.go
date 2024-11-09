package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// dials server
	conn, err := net.Dial("tcp", "localhost:7777")

	//handles errors
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewReader(os.Stdin)

	for {
		var send string

		fmt.Print("Enter Something: ")
		send, err := scanner.ReadString('\n')

		_, err = conn.Write([]byte(send))

		if err != nil {
			fmt.Println(err)
			return
		}

		buff := make([]byte, 512)
		conn.Read(buff)

		fmt.Printf("\nRecived back: %s\n", buff)

		if string(buff[509:]) == "dne" {
			conn.Close()
			break
		}

	}

}
