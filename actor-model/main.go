package main

type Player struct {
	Name string
}

type GameState struct {
	// lock    sync.RWMutex
	players []*Player
	msgch   chan any
}

func (g *GameState) Receive(msg any) {
	g.msgch <- msg
}

func (g *GameState) loop() {
	for msg := range g.msgch {
		g.MustHandleMessage(msg)

	}
}

// ROUTER
func (g *GameState) MustHandleMessage(msg any) {
	switch v := msg.(type) {
	case *Player:
		g.addPlayer(v)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) addPlayer(player *Player) {
	// g.lock.Lock()
	g.players = append(g.players, player)
	// g.lock.Unlock()
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch:   make(chan any, 10),
	}
	go g.loop()

	return g

}

type Server struct {
	gamestate *GameState
}

func NewServer() *Server {
	return &Server{
		gamestate: NewGameState(),
	}
}

func (s *Server) handleNewPlayer(player *Player) error {
	s.gamestate.Receive(player)
	return nil
}

func main() {

}
