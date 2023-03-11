package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("listen error: %x", err)
	}

	for {
		connection_socket, err := ln.Accept()
		if err != nil {
			log.Fatalf("connection_socket err: %x", err)
		}

		go fs.readLoop(connection_socket)

	}
}

func (fs *FileServer) readLoop(conn_sock net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn_sock, binary.LittleEndian, &size)
		i, err := io.CopyN(buf, conn_sock, size)
		if err != nil {
			log.Fatalf("connection_socket err: %x", err)
		}

		// fmt.Println(buf.Bytes()) // show data from buffer
		fmt.Println("Received data over tcp,size:", i)
	}

}

func senderSimulation(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	socket, err := net.Dial("tcp4", ":3000")
	if err != nil {
		return err
	}
	binary.Write(socket, binary.LittleEndian, int64(size))
	n, err := io.CopyN(socket, bytes.NewReader(file), int64(size))

	if err != nil {
		return err
	}

	fmt.Println("data sent to server size:", n)

	return nil
}

func main() {

	go func() {
		time.Sleep(4 * time.Second)
		err := senderSimulation(3000)
		if err != nil {
			log.Println("sender", err)
		}
	}()

	newServer := &FileServer{}
	newServer.start()

}

// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// type FileServer struct {
// }

// func (fs *FileServer) start() {
// 	ln, err := net.Listen("tcp", ":3000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		go fs.readLoop(conn)
// 	}
// }

// func (fs *FileServer) readLoop(conn net.Conn) {
// 	buf := make([]byte, 2048)
// 	for {
// 		n, err := conn.Read(buf)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		file := buf[:n]
// 		fmt.Println(file)
// 		fmt.Printf("received %d bytes over the network \n", n)

// 	}

// }
// func sendFile(size int) error {
// 	file := make([]byte, size)
// 	_, err := io.ReadFull(rand.Reader, file)
// 	if err != nil {
// 		return err
// 	}

// 	conn, err := net.Dial("tcp4", ":3000")
// 	if err != nil {
// 		return err
// 	}

// 	n, err := conn.Write(file)
// 	if err != niListen()
// 		return err
// 	}
// 	fmt.Printf("written %d  bytes over the network\n", n)
// 	return nil
// }

// func main() {
// 	go func() {
// 		time.Sleep(4. * time.Second)
// 		sendFile(1000)
// 	}()
// 	server := &FileServer{}
// 	server.start()

// }
