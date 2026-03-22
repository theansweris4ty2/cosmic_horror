package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func (c *Camera) followTarget(t *Player) {
	c.x = -t.x
	c.y = -t.y
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
		g.player.veloY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.veloY += 1
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
