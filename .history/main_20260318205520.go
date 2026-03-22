package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	// "image/color"
)

const (
	screenWidth  = 900
	screenHeight = 630
)

type Player struct {
	playerImg    *ebiten.Image
	x, y         float64
	veloX, veloY float64
}

type Game struct {
	player *Player
}

func (g *Game) Update() error {

	g.player.veloX = 0
	g.player.veloY = 0
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.veloX += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.veloX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.veloY += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.veloY -= 1
	}
	g.player.x += g.player.veloX
	g.player.y += g.player.veloY
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.playerImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}

func main() {
	playerImg, _, err := ebitenutil.NewImageFromFile("ninja.png")
	if err != nil {
		fmt.Println(err)
	}
	p := Player{playerImg, 100, 100, 0, 0}
	g := Game{&p}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
