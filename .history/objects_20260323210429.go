package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	playerImages   []*ebiten.Image
	x, y           float64
	veloX, veloY   float64
	animationFrame int
	action         int
	forward        bool
}

type Camera struct {
	x, y float64
}

type Game struct {
	player *Player
	camera *Camera
	timer  int
}
