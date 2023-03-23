package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anthdm/hollywood/actor"
	websocket "golang.org/x/net/websocket"
)

var engine *actor.Engine

type broadcastMsg struct {
	pid  *actor.PID
	data string
}

type foo struct {
}

func newFoo() actor.Receiver {
	return &foo{}
}

func (f *foo) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("foo has started")
	case *broadcastMsg:
		fmt.Println("foo has received", msg.data)
		engine.Send(msg.pid, msg)

	}
}

func handlerWS(ws *websocket.Conn) {
	wsConn := ws.RemoteAddr().String()
	fmt.Println("new incoming connection from client: ", wsConn)

	pid := engine.Spawn(newFoo, wsConn)
	engine.Send(pid, &broadcastMsg{
		pid:  pid,
		data: "I'm broadcast message.",
	})

}

func main() {

	engine = actor.NewEngine()
	http.Handle("/ws", websocket.Handler(handlerWS))

	time.Sleep(time.Second * 1000)

}
