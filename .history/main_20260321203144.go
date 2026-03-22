package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "image/color"
)

// TODO Resize screenwidth to 960 and screen height to 672 so map will be 60 x 42 tiles
const (
	screenWidth  = 900
	screenHeight = 630
)

type Player struct {
	playerImages   []*ebiten.Image
	x, y           float64
	veloX, veloY   float64
	animationFrame int
	action         string
	forward        bool
}

type Camera struct {
	x, y float64
}

type Game struct {
	player *Player
	camera *Camera
}

func main() {

	p := new(Player{[]*ebiten.Image{playerImg, playerImg2, playerImg3, playerImg4, playerImg5}, 100, 100, 0, 0, 0, "walk", true})
	c := Camera{0, 0}
	g := Game{p, &c}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
