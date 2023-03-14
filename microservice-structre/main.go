package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	ApiServer := NewApiServer(svc)

	log.Fatal(ApiServer.Start(":3000"))
}
