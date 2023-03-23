package main

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

type server struct {
}

func newServer() actor.Receiver {
	return &server{}
}

func (f *server) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("server has started")
	case *actor.PID:
		fmt.Println("server has received:", msg)
	}
}

func main() {

	e := actor.NewEngine()
	r := remote.New(e, remote.Config{ListenAddr: "127.0.0.1:4000"})

	e.WithRemote(r)
	e.Spawn(newServer, "server")

	time.Sleep(time.Second * 1000)

}

// func createListener(port int) (net.Listener, error) {
// 	// Create a TCP listener on the specified port
// 	addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: port}
// 	tcpListener, err := net.ListenTCP("tcp", addr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get the underlying socket handle
// 	conn, err := tcpListener.SyscallConn()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var socketHandle syscall.Handle
// 	err = conn.Control(func(fd uintptr) {
// 		socketHandle = syscall.Handle(fd)
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Enable SO_REUSEADDR socket option
// 	var reuseAddr int = 1
// 	optval := *(*[4]byte)(unsafe.Pointer(&reuseAddr))
// 	err = syscall.Setsockopt(socketHandle, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, &optval[0], int32(len(optval)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create a new TCP listener from the TCP listener object
// 	return tcpListener, nil
// }

// NEW reuseaddr

// func createListener(port int) (net.Listener, error) {
//     // Create a TCP listener on the specified port
//     addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: port}
//     tcpListener, err := net.ListenTCP("tcp", addr)
//     if err != nil {
//         return nil, err
//     }

//     // Get the underlying socket handle
//     conn, err := tcpListener.SyscallConn()
//     if err != nil {
//         return nil, err
//     }
//     var socketHandle syscall.Handle
//     err = conn.Control(func(fd uintptr) {
//         socketHandle = syscall.Handle(fd)
//     })
//     if err != nil {
//         return nil, err
//     }

//     // Enable SO_REUSEADDR socket option
//     err = syscall.SetsockoptInt(socketHandle, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
//     if err != nil {
//         return nil, err
//     }

//     // Create a new TCP listener from the TCP listener object
//     return tcpListener, nil
// }
