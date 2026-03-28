package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// TODO Create function to play sounds and add sound files

func (g *Game) Update() error {
	g.timer += 1

	var gravity float64
	g.player.movement(g)
	if g.player.y+128 <= 240 {
		gravity = 2

		if g.player.forward == true {
			g.player.action = 2
		} else {
			g.player.action = 6
		}

	} else {
		gravity = 0
	}
	g.monster.x += g.monster.veloX
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
	msrcX := g.monster.animationFrame % 10 * 128
	msrcY := 0
	for i := -1; i < 3; i++ {
		opts.GeoM.Translate((g.camera.x)+float64(i*384), 0)
		screen.DrawImage(g.background, opts)
		g.background.DrawImage(g.middleground, opts)
		opts.GeoM.Reset()
	}

	opts.GeoM.Translate(g.camera.x, g.camera.y)
	opts.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(g.player.playerImages[g.player.action].SubImage(image.Rect(srcX, srcY, srcX+128, srcY+128)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

	opts.GeoM.Translate(g.monster.x, g.monster.y)
	opts.GeoM.Translate(g.camera.x, g.camera.y)
	screen.DrawImage(g.monster.monsterImages[0].SubImage(image.Rect(msrcX, msrcY, msrcX+128, msrcY+128)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}
func (p *Player) movement(g *Game) {
	p.veloX = 0
	p.veloY = 0
	g.monster.veloX = 0
	g.monster.veloY = 0

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
		if p.x+16 >= 900 {
			p.veloX *= -1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && ebiten.IsKeyPressed(ebiten.KeyR) {
		p.action = 7
		p.forward = true
		if p.animationFrame < 8 {
			if g.timer%8 == 0 {
				p.animationFrame += 1
			}

		} else {
			p.animationFrame = 0
		}
		p.veloX += 3
		if p.x+16 >= 900 {
			p.veloX *= -3
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && ebiten.IsKeyPressed(ebiten.KeyR) {
		p.action = 8
		p.forward = false
		if p.animationFrame > 0 {
			if g.timer%8 == 0 {
				p.animationFrame -= 1
			}

		} else {
			p.animationFrame = 8
		}

		p.veloX -= 3
		if p.x <= 0 {
			p.veloX *= -3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.y+128 >= 240 {
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
			if p.animationFrame < 9 {
				if g.timer%3 == 0 {
					p.animationFrame += 1
				}
			} else {
				p.animationFrame = 5
			}
		}

	}
	g.monster.veloX = -1
	if g.timer%20 == 0 && g.monster.animationFrame > 0 {
		g.monster.animationFrame -= 1
	} else {
		g.monster.animationFrame = 9
	}

}
