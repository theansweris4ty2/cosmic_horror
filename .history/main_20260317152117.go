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
	playerImg *ebiten.Image
	x, y      float64
	// veloX, veloY int
}

type Game struct {
	player *Player
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.x += 1
	}

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
	p := Player{playerImg, 100, 100}
	g := Game{&p}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
