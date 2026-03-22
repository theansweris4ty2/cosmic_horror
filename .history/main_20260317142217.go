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

func main() {
	ebiten.SetScreenSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello World")

}
