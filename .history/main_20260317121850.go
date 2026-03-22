package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
)
// type Camera struct {
// 	x, y float64
// }

// func (c *Camera) followTarget(t *Player, screenWidth, screenHeight int) {
// 	c.x = -t.x + float64(screenWidth/200)
// 	c.y = -t.y + float64(screenHeight/200)
// }

// type SpriteSheet struct {
// 	WidthInTiles  int
// 	HeightInTiles int
// 	TileWidth     int
// 	TileHeight    int
// }
// type TilemapLayerJSON struct {
// 	Data   []int `jason:"data"`
// 	Width  int   `json:"width"`
// 	Height int   `json:"height"`
// }

// type TilemapJSON struct {
// 	Layers []TilemapLayerJSON `json:"layers"`
// }
// type Player struct {
// 	Image *ebiten.Image

// 	x, y             float64
// 	x_speed, y_speed float64
// 	animationIndex   int
// 	spriteCounter    int
// 	currentAnimation int
// }

type Game struct {
	// player            *Player
	// playerSpriteSheet *SpriteSheet
	// tilemapJSON       *TilemapJSON
	// tilemapImg        *ebiten.Image
	// camera            *Camera
}

const (
	screenWidth  = 900
	screenHeight = 630
)

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	playerImg, _, err := ebitenutil.NewImageFromFile("./images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}
	tilemapImg, _, err := ebitenutil.NewImageFromFile("./images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	playerSpriteSheet := NewSpriteSheet(4, 7, 16, 16)

	tilemapJSON, err := newTilemapJSON("maps/tilesets/spawn.json")

	if err != nil {
		log.Fatal(err)
	}
	player := Player{playerImg, 100, 100, 0, 0, 0, 0, 0}
	camera := Camera{0, 0}
	game := Game{&player, playerSpriteSheet, tilemapJSON, tilemapImg, &camera}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
func (g *Game) Update() error {
	g.camera.followTarget(g.player, screenWidth, screenHeight)
	g.player.checkPosition()

	return nil
}