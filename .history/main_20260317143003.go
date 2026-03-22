package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
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

	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(screenWidth/2, screenHeight/2)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	return screenWidth, screenHeight

}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")

}
