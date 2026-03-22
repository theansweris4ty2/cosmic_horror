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

type Game struct {
}

func (g *Game) Update() error {

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

	playerImg, _, err := ebitenutil.NewImageFromFile("ninja.png")

	if err != nil {
		fmt.Println(err)
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(playerImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}

func main() {
	g := Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")
	if err := ebiten.RunGame(&g); err != nil {
		fmt.Println(err)
	}

}
