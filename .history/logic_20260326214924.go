package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func (c *Camera) followTarget(t *Player) {
	c.x = -t.x + 50
	c.y = -t.y + 50
}

func (g *Game) Update() error {
	g.timer += 1
	var gravity float64
	g.player.movement(g)
	if g.player.y+128 < 228 {
		gravity = 2

		if g.player.forward == true {
			g.player.action = 2
		} else {
			g.player.action = 6
		}

	} else {
		gravity = 0
	}
	g.player.veloY = gravity
	g.player.x += g.player.veloX
	g.player.y += g.player.veloY
	g.camera.x = -g.player.x + 100
	g.camera.y = g.player.y - 100

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	srcX := g.player.animationFrame % 10 * 128
	srcY := 0

	for i := -1; i < 10; i++ {
		opts.GeoM.Translate((g.camera.x)+float64(i*384), 0)
		// opts.GeoM.Translate(g.camera.x, g.camera.y)
		screen.DrawImage(g.background, opts)
		opts.GeoM.Reset()
	}

	// opts2 := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.camera.x, g.camera.y)
	opts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.playerImages[g.player.action].SubImage(image.Rect(srcX, srcY, srcX+128, srcY+128)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
func (p *Player) movement(g *Game) {
	p.veloX = 0
	p.veloY = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.action = 0
		p.forward = true
		if p.animationFrame < 8 {
			if g.timer%3 == 0 {
				p.animationFrame += 1
			}

		} else {
			p.animationFrame = 0
		}
		p.veloX += 1
		if p.x+16 >= 360 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.action = 1
		p.forward = false
		if p.animationFrame > 0 {
			if g.timer%3 == 0 {
				p.animationFrame -= 1
			}

		} else {
			p.animationFrame = 8
		}

		p.veloX -= 1
		if p.x <= 0 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.y+128 >= 228 {
		if p.forward == true {
			p.action = 2
			p.animationFrame = 0
			p.animationFrame += 1
		} else {
			p.action = 6
			p.animationFrame = 8
			p.animationFrame -= 1
		}

		p.y += -40

	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if p.forward {
			p.action = 3
			if p.animationFrame < 5 {
				if g.timer%3 == 0 {
					p.animationFrame += 1
				}
			} else {
				p.animationFrame = 0
			}
		} else {
			p.action = 4
			if p.animationFrame > 0 {
				if g.timer%3 == 0 {
					p.animationFrame -= 1
				}
			} else {
				p.animationFrame = 5
			}
		}

	}

}
