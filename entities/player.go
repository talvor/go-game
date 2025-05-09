package entities

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/talvor/tiled/animation/renderer"
	"github.com/talvor/tiled/common"
)

type Player struct {
	*Sprite
	Health    uint
	MaxHealth uint
	Renderer  *renderer.Renderer
	Direction string
	Action    string
}

func NewPlayer(x, y float64, renderer *renderer.Renderer) *Player {
	return &Player{
		Sprite:    &Sprite{X: x, Y: y},
		Health:    100,
		MaxHealth: 100,
		Renderer:  renderer,
		Direction: "down",
		Action:    "idle",
	}
}

func (p *Player) TakeDamage(hits uint) {
	p.Health -= hits
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *Player) Heal(hits uint) {
	p.Health += hits
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.X, p.Y)

	opts := &common.DrawOptions{
		Screen: screen,
		Op:     op,
	}

	err := p.Renderer.Draw("Player", fmt.Sprintf("%s_%s", p.Action, p.Direction), opts)
	if err != nil {
		log.Println("Error drawing animation:", err)
	}
}

func (p *Player) Update() {
	p.DX = 0.0
	p.DY = 0.0

	p.Action = "walk"
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.DX = -2
		p.Direction = "left"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.DX = 2
		p.Direction = "right"
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.DY = -2
		p.Direction = "up"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.DY = 2
		p.Direction = "down"
	} else {
		p.Action = "idle"
	}

	p.X += p.DX
	p.Y += p.DY
}
