package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func (c *Camera) followTarget(t *Player) {
	c.x = -t.x + float64(screenWidth/10)
	c.y = -t.y + float64(screenHeight/10)
}

func (g *Game) Update() error {
	// g.camera.followTarget(g.player)
	g.player.checkCollisions()

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(g.camera.x, g.camera.y)
	// screen.DrawImage(g.player.playerImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)
	// opts.GeoM.Reset()

	opts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.playerImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}
func (p *Player) checkCollisions() {
	p.veloX = 0
	p.veloY = 0
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.veloX += 1
		if p.x+16 >= 360 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.veloX -= 1
		if p.x <= 0 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.veloY -= 1
		if p.y <= 0 {
			p.veloY *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.veloY += 1
		if p.y+16 >= 270 {
			p.veloY *= -1
		}
	}

	p.x += p.veloX
	p.y += p.veloY

}
