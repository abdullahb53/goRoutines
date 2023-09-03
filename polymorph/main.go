package main

import "fmt"

type Tile struct{}

type TileWalker interface {
	WalkTile(Tile)
}

type Updater interface {
	Update()
}

type Transform struct {
	position int
}

type Enemy struct {
	Transform
	tileWalker TileWalker
}

func (e *Enemy) checkTilesCollided() {
	// Some stuff.
	// fmt.Println("Enemy is walking on tile", e.position)
	e.tileWalker.WalkTile(Tile{})
}

func (e *Enemy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

type FireEnemy struct {
	*Enemy
}

func (e *FireEnemy) WalkTile(tile Tile) {
	fmt.Println("Fire enemy is walking on tile")
}

type WaterEnemy struct {
	*Enemy
}

func (e *WaterEnemy) WalkTile(tile Tile) {
	fmt.Println("Water enemy is walking on tile, wet wet wet")
}

func main() {
	e := &WaterEnemy{}
	e.Enemy = &Enemy{
		tileWalker: e,
	}

	for i := 0; i < 100; i++ {
		Update(e)
	}
}

func Update(u Updater) {
	u.Update()
}
