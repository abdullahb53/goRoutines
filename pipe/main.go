package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		defer pipeWriter.Close()
		_, err := pipeWriter.Write([]byte("123456789"))
		if err != nil {
			log.Fatalf("%v \n", err)
		}

	}()

	time.Sleep(time.Second * 1)
	buffer := make([]byte, 10)
	_, err := pipeReader.Read(buffer)
	if err != nil {
		log.Fatalf("%v \n", err)
	}
	fmt.Print(buffer)

	_, err = pipeReader.Read(buffer)
	if err != nil {
		log.Fatalf("%v \n", err)
	}
	fmt.Print(buffer)

}
