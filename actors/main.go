package main

import (
	"log"
	"net"

	"github.com/anthdm/hollywood/actor"
)

type session struct {
	conn net.Conn
}
type connRem struct {
	pid *actor.PID
}

type connAdd struct {
	pid  *actor.PID
	conn net.Conn
}

func newSession(conn net.Conn) actor.Producer {
	return func() actor.Receiver {
		return &session{
			conn: conn,
		}
	}

}

func (s *session) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Initialized:

	case actor.Started:
		log.Println("new connection:", s.conn.RemoteAddr())
		go s.readLoop(c)
	case actor.Stopped:
	case []byte:
		s.conn.Write(msg)
	}
}

func (s *session) readLoop(c *actor.Context) {
	buf := make([]byte, 1024)
	for {
		n, err := s.conn.Read(buf)
		if err != nil {
			log.Println("conn read error", err)
			break
		}
		msg := buf[:n]
		c.Send(c.PID(), msg)
	}

	// Loop is done
	c.Send(c.Parent(), &connRem{pid: c.PID()})

}

type server struct {
	listenAddr string
	ln         net.Listener
	sessions   map[*actor.PID]net.Conn
}

func newServer(listenAddr string) actor.Producer {
	return func() actor.Receiver {
		return &server{
			listenAddr: listenAddr,
			sessions:   make(map[*actor.PID]net.Conn),
		}
	}
}

func (s *server) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Initialized:
		ln, err := net.Listen("tcp", s.listenAddr)
		if err != nil {
			panic(err)
		}
		s.ln = ln
	case actor.Started:
		log.Println("server started", s.listenAddr)
		go s.acceptLoop(c)
	case actor.Stopped:
		_ = msg
	case *connAdd:
		log.Println("added new conn to my map", msg.conn.RemoteAddr(), "pid:", msg.pid)
		s.sessions[msg.pid] = msg.conn
	case *connRem:
		log.Println("Removed connection from my map.", msg.pid)
		delete(s.sessions, msg.pid)
	}
}

func (s *server) acceptLoop(c *actor.Context) {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println("accept error", err)
			break
		}
		pid := c.SpawnChild(newSession(conn), "session", actor.WithTags(conn.RemoteAddr().String()))
		c.Send(c.PID(), &connAdd{
			pid:  pid,
			conn: conn,
		})
	}
}

func main() {
	e := actor.NewEngine()
	e.Spawn(newServer(":6000"), "server")

	<-make(chan struct{})

}
