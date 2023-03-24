package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	websocket "golang.org/x/net/websocket"
)

var engine *actor.Engine

type socketMsg struct {
	pid *actor.PID
	ws  *websocket.Conn
}

type storageSocket struct {
	storage map[*websocket.Conn]*actor.PID
}

func newStorageSocket() actor.Receiver {
	return &storageSocket{
		storage: make(map[*websocket.Conn]*actor.PID),
	}
}

func (f *storageSocket) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("[WEBSOCKET] storageSocket has started")
	case *socketMsg:

		fmt.Println("storageSocket has received", msg.ws)
		f.storage[msg.ws] = msg.pid
	}
}

type setWebsocketVal struct {
	pid *actor.PID
	ws  *websocket.Conn
}

type foo struct {
	ws *websocket.Conn
}

func newFoo() actor.Receiver {
	return &foo{
		ws: &websocket.Conn{},
	}
}

func (f *foo) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("[WEBSOCKET] foo has started")

	case *setWebsocketVal:
		fmt.Println("foo has received", msg.pid)
		f.ws = msg.ws
		engine.Send(storeWSocketPIDs, msg.ws)
		orws := msg.ws
		go func(orws *websocket.Conn) {

			buf := make([]byte, 1024)
			for {
				n, err := orws.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println("read error: ", err)
					continue
				}
				msg := buf[:n]
				fmt.Println("messageee:", msg)

			}

		}(orws)

	}
}

type pureHandleWS func(ws *websocket.Conn) *actor.PID

func HandleFuncImitation(f pureHandleWS) websocket.Handler {
	return func(c *websocket.Conn) {
		pid := f(c)
		fmt.Println("new websocket spawned: ", pid.ID)
	}
}

func handlerWS(ws *websocket.Conn) *actor.PID {
	pid := engine.Spawn(newFoo, ws.RemoteAddr().String())
	defer engine.Send(pid, &setWebsocketVal{
		pid: pid,
		ws:  ws,
	})
	return pid
}

var storeWSocketPIDs *actor.PID

func main() {

	engine = actor.NewEngine()
	storeWSocketPIDs = engine.Spawn(newStorageSocket, "storer")

	http.Handle("/ws", websocket.Handler(HandleFuncImitation(handlerWS)))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
