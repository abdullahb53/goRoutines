package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Player struct {
	mu     sync.RWMutex
	health int32
}

func (p *Player) takeDamage(value int) {
	health := p.getHealth()
	atomic.StoreInt32(&p.health, int32(health-value))
	// p.mu.TryLock()
	// p.health -= int32(value)
	// p.mu.Unlock()

}

func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}

func NewPlayer() *Player {
	return &Player{
		health: 300,
	}
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println("Player health:", p.getHealth())
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.takeDamage(rand.Intn(40))
		if p.getHealth() <= 0 {
			fmt.Println("GAME OVER")
			break
		}
		<-ticker.C
	}
}

func main() {

	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)

}
