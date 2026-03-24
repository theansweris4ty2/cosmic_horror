package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func (c *Camera) followTarget(t *Player) {
	c.x = -t.x + float64(screenWidth/10)
	c.y = -t.y + float64(screenHeight/10)
}

func (g *Game) Update() error {
	// g.camera.followTarget(g.player)
	g.player.actions()

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(g.camera.x, g.camera.y)
	// opts.GeoM.Reset()
	srcX := g.player.animationFrame % 10 * 128
	srcY := 0

	opts.GeoM.Translate(g.player.x, g.player.y)
	switch g.player.action {
	case "walk":
		g.player.actionImage = 0
	case "walk_left":
		g.player.actionImage = 1
	case "jump":
		g.player.actionImage = 2
	case "fight":
		g.player.actionImage = 3
	case "fight_left":
		g.player.actionImage = 4
	default:
		g.player.actionImage = 5
	}
	screen.DrawImage(g.player.playerImages[g.player.actionImage].SubImage(image.Rect(srcX, srcY, srcX+128, srcY+128)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}
func (p *Player) actions() {
	p.veloX = 0
	p.veloY = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.action = "walk"
		p.forward = true
		if p.animationFrame < 8 {
			p.animationFrame += 1
		} else {
			p.animationFrame = 0
		}
		p.veloX += 1
		if p.x+16 >= 360 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.action = "walk_left"
		p.forward = false
		if p.animationFrame > 0 {
			p.animationFrame -= 1
		} else {
			p.animationFrame = 8
		}

		p.veloX -= 1
		if p.x <= 0 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.action = "jump"
		if p.animationFrame < 8 {
			p.animationFrame += 1
		} else {
			p.animationFrame = 0
		}
		p.veloY -= 1
		if p.y <= 0 {
			p.veloY *= -1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if p.forward {
			p.action = "fight"
		} else {
			p.action = "fight_left"
		}
		if p.animationFrame < 5 {
			p.animationFrame += 1
		} else {
			p.animationFrame = 0
		}
	}

	p.x += p.veloX
	p.y += p.veloY

}
