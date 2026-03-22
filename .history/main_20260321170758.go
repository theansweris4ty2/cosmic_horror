package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "image/color"
)

const (
	screenWidth  = 900
	screenHeight = 630
)

type Player struct {
	playerImages   []*ebiten.Image
	x, y           float64
	veloX, veloY   float64
	animationCount int
	animationFrame int
	action         string
}

type Camera struct {
	x, y float64
}

type Game struct {
	player *Player
	camera *Camera
}

func main() {
	playerImg, _, err := ebitenutil.NewImageFromFile("Walk.png")
	playerImg2, _, err := ebitenutil.NewImageFromFile("Walk_left.png")
	if err != nil {
		fmt.Println(err)
	}
	p := new(Player{[]*ebiten.Image{playerImg, playerImg2}, 100, 100, 0, 0, 0, 0, "walk"})
	c := Camera{0, 0}
	g := Game{p, &c}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
