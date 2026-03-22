package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init(g *Game) []*ebiten.Image {
	playerImg, _, err := ebitenutil.NewImageFromFile("Walk.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg2, _, err := ebitenutil.NewImageFromFile("Walk_left.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg3, _, err := ebitenutil.NewImageFromFile("Jump.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg4, _, err := ebitenutil.NewImageFromFile("Attack_1.png")
	if err != nil {
		fmt.Println(err)
	}
	playerImg5, _, err := ebitenutil.NewImageFromFile("Attack_1_back.png")
	if err != nil {
		fmt.Println(err)
	}
	return new([]*ebiten.Image{playerImg1, playerImg2, playerImg3, playerImg4, playerImg5})
}
