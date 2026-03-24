package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "image/color"
)

// TODO Resize screenwidth to 960 and screen height to 672 so map will be 60 x 42 tiles
const (
	screenWidth  = 960
	screenHeight = 672
)

func main() {
	playerImages := loadPlayerImages()
	p := new(Player{playerImages, 100, 100, 0, 0, 0, 5, true})
	c := Camera{0, 0}
	g := Game{p, &c}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
