package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "image/color"
)

const (
	screenWidth  = 1000
	screenHeight = 800
)

func main() {
	playerImages, background := loadPlayerImages()
	p := new(Player{playerImages, 100, 100, 0, 0, 0, 5, true})
	c := Camera{0, 0}
	g := Game{p, &c, 0, background}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
