package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {

	port := "5010"

	new_socket, err := net.Dial("tcp4", "localhost:"+port)
	if err != nil {
		panic(err)
	}

	go func() {

		for {

			buffer := make([]byte, 6)
			_, err := new_socket.Read(buffer[:])
			if err != nil {
				fmt.Println(err)
				new_socket.Close()
				break
			}
			fmt.Println("Message from server:", string(buffer))
		}

	}()
	i := 0
	for {
		i++
		time.Sleep(1 * time.Second)

		_, err := new_socket.Write([]byte("mssg" + strconv.Itoa(i)))
		if err != nil {
			fmt.Println("Write error. new_socket is gonna close: remote:", new_socket.RemoteAddr(), "local:", new_socket.LocalAddr(), "addr:", new_socket.LocalAddr().Network())
			new_socket.Close()
			break
		}
	}

	// I want to see what will happen
	ch := make(chan bool)
	<-ch

}
