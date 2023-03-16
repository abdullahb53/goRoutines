package main

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y int) {

}

type Position struct {
	x int
	y int
}

func (p *Position) Move() {

}

func Teleport(p *Position) {

}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {

	Player := NewPlayer()
	Player.Move()

	RaidBoss := NewEnemy()
	RaidBoss.Move()
	RaidBoss.MoveSpecial(1, 3)

}
