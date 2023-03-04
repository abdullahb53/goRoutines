package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	port := "5010"

	listener, err := net.Listen("tcp4", "localhost:"+port)
	if err != nil {
		panic(err)
	}

	go func() {

		fmt.Println("Listen connections." + port)
		for {

			new_socket, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err:", err)
				break
				continue

			}

			handler(new_socket, listener)

		}
	}()

	// I want to see what will happen
	ch := make(chan bool)
	<-ch

}

func handler(new_socket net.Conn, listener net.Listener) {
	for {

		// Make a new buffer which is size 10 byte.
		buffer := make([]byte, 6)
		_, err := new_socket.Read(buffer[:])
		if err != nil {
			fmt.Println(err)
			new_socket.Close()
			break
		}

		fmt.Println(string(buffer))

		_, err = new_socket.Write(buffer)
		if err != nil {
			fmt.Println(err)
		}

		// if i send close text, the listener will no longer allow new connections.
		if strings.Contains(string(buffer), "close") {
			fmt.Println("close")
			listener.Close()
		}

	}
}
