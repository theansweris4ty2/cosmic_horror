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
	playerImg      *ebiten.Image
	x, y           float64
	veloX, veloY   float64
	animationCount int
}

type Camera struct {
	x, y float64
}

type Game struct {
	player *Player
	camera *Camera
}

func main() {
	playerImg, _, err := ebitenutil.NewImageFromFile("ninja.png")
	if err != nil {
		fmt.Println(err)
	}
	p := new(Player{playerImg, 100, 100, 0, 0, 0})
	c := Camera{0, 0}
	g := Game{p, &c}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
