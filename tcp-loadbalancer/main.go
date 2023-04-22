package main

import (
	"io"
	"log"
	"net"
)

var (
	count  = 0
	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func chooseBackend() string {
	s := server[count%len(server)]
	count++
	return s
}

func proxy(backend string, c net.Conn) error {
	bc, err := net.Dial("tcp4", backend)
	if err != nil {
		log.Printf("failed to connect to backend %s: %v", backend, err)
	}

	// c -> bc
	go io.Copy(bc, c)

	// bc -> c
	go io.Copy(c, bc)

	return nil
}

func main() {

	ls, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	log.Printf("INFO: tcp is listening at: %s", ls.Addr().String())
	defer ls.Close()

	for {
		socket, err := ls.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
		}

		backend := chooseBackend()

		go func() {
			err := proxy(backend, socket)
			if err != nil {
				log.Printf("WARNING: proxying failed: %v", err)
			}
		}()

	}

}
