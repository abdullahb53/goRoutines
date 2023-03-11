package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	websocket "golang.org/x/net/websocket"
)

type Server struct {
	connection_sockets map[*websocket.Conn]bool
}

func CreateServer() *Server {
	return &Server{
		connection_sockets: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handlerWS(ws *websocket.Conn) {
	s.connection_sockets[ws] = true
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr().String())
	s.readLoop(ws)
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client to orderbook feed:", ws.RemoteAddr().String())

	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error: ", err)
			continue
		}
		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.connection_sockets {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}

func main() {
	server := CreateServer()
	http.Handle("/ws", websocket.Handler(server.handlerWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(":3000", nil)
}
