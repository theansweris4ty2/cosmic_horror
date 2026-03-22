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
	srcX := g.player.animationFrame % 4
	srcY := g.player.animationFrame / 7

	opts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.playerImg.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}
func (p *Player) checkCollisions() {
	p.veloX = 0
	p.veloY = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.animationCount < 5 {
			p.animationCount += 1
		} else {
			p.animationCount = 0
		}
		frames := []int{3, 7, 11, 15, 19, 23}
		p.animationFrame = p.walk(frames)
		p.veloX += 1
		if p.x+16 >= 360 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.animationCount < 5 {
			p.animationCount += 1
		} else {
			p.animationCount = 0
		}
		frames := []int{2, 6, 10, 14, 18, 22}
		p.animationFrame = p.walk(frames)
		p.veloX -= 1
		if p.x <= 0 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if p.animationCount < 5 {
			p.animationCount += 1
		} else {
			p.animationCount = 0
		}
		frames := []int{1, 5, 9, 13, 17, 21}
		p.animationFrame = p.walk(frames)
		p.veloY -= 1
		if p.y <= 0 {
			p.veloY *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if p.animationCount < 5 {
			p.animationCount += 1
		} else {
			p.animationCount = 0
		}
		frames := []int{0, 4, 8, 12, 16, 20}
		p.animationFrame = p.walk(frames)
		p.veloY += 1

		if p.y+16 >= 270 {
			p.veloY *= -1
		}

	}

	p.x += p.veloX
	p.y += p.veloY

}
func (p *Player) walk(frames []int) int {

	if p.animationFrame < 5 {
		p.animationFrame += 1
	} else {
		p.animationFrame = 0
	}
	return frames[p.animationFrame]

}
