package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/websocket"
)

func TestWebsocket(t *testing.T) {
	server := CreateServer()
	s := httptest.NewServer(websocket.Handler((server.handlerWS)))
	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, err := websocket.Dial(u, "", u)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	for i := 0; i < 100; i++ {
		if _, err := ws.Write([]byte("Some text.")); err != nil {
			t.Fatalf("%v", err)
		}
		buf := make([]byte, 1024)
		_, err := ws.Read(buf)
		if err != nil {
			t.Fatalf("%v", err)
		}
		println(string(buf))
	}

}
