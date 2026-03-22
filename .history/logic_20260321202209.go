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

	frame := 0
	opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(g.camera.x, g.camera.y)
	// screen.DrawImage(g.player.playerImg.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), opts)
	// opts.GeoM.Reset()
	srcX := g.player.animationFrame % 10 * 128
	srcY := 0

	opts.GeoM.Translate(g.player.x, g.player.y)
	switch g.player.action {
	case "walk":
		frame = 0
	case "walk_left":
		frame = 1
	case "jump":
		frame = 2
	case "fight":
		frame = 3
	case "fight_left":
		frame = 4
	default:
		frame = 0
	}
	screen.DrawImage(g.player.playerImages[frame].SubImage(image.Rect(srcX, srcY, srcX+128, srcY+128)).(*ebiten.Image), opts)
	opts.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 360, 270
}
func (p *Player) checkCollisions() {
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		frames := []int{0, 4, 8, 12, 16, 20}
		p.animationFrame = p.walk(frames)
		p.veloY += 1

		if p.y+16 >= 270 {
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
func (p *Player) walk(frames []int) int {

	if p.animationCount < 5 {
		p.animationCount += 1
	} else {
		p.animationCount = 0
	}
	return frames[p.animationCount]

}
