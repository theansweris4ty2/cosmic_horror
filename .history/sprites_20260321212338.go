package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TilemapLayerJSON struct {
	Data   []int `jason:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TilemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`
}

func loadPlayerImages() []*ebiten.Image {
	playerImg1, _, err := ebitenutil.NewImageFromFile("Walk.png")
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
	return []*ebiten.Image{playerImg1, playerImg2, playerImg3, playerImg4, playerImg5}
}
