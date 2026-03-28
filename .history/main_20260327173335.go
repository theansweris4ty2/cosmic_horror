package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "image/color"
)

const (
	screenWidth  = 900
	screenHeight = 630
)

func main() {
	playerImages, background, middleground, monsterImages := loadPlayerImages()
	p := new(Player{playerImages, 100, 115, 0, 0, 0, 5, true})
	m := new(Monster{[]*ebiten.Image{monsterImages}, 90, 600, 0, 0, 9, 5, true})
	c := Camera{0, 0}
	g := Game{p, m, &c, 0, background, middleground}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cosmic Horror")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
