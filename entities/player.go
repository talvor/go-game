package entities

import (
	"encoding/json"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/talvor/go-rpg/camera"
	"github.com/talvor/tiled/animation/renderer"
	"github.com/talvor/tiled/common"
)

type Player struct {
	*Sprite
	Health    uint
	MaxHealth uint
	renderer  *renderer.Renderer
	Direction string
	Action    string
}

func NewPlayer(x, y float64, renderer *renderer.Renderer) *Player {
	sprite := NewSprite(x, y, 32)
	return &Player{
		Sprite:    sprite,
		Health:    100,
		MaxHealth: 100,
		renderer:  renderer,
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

func (p *Player) Draw(screen *ebiten.Image, cam *camera.Camera) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.X, p.Y)
	op.GeoM.Translate(cam.X, cam.Y)

	opts := &common.DrawOptions{
		Screen: screen,
		Op:     op,
	}

	err := p.renderer.Draw("Player", fmt.Sprintf("%s_%s", p.Action, p.Direction), opts)
	if err != nil {
		log.Println("Error drawing animation:", err)
	}
}

func (p *Player) Update(colliders []image.Rectangle) {
	p.DX = 0.0
	p.DY = 0.0

	p.Action = "idle"

	// Move horizontally
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.DX = -2
		p.Direction = "left"
		p.Action = "walk"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.DX = 2
		p.Direction = "right"
		p.Action = "walk"
	}
	p.determineCollider()
	p.X += p.DX
	p.CheckCollisionHorizontal(colliders)

	// Move vertically
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.DY = -2
		p.Direction = "up"
		p.Action = "walk"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.DY = 2
		p.Direction = "down"
		p.Action = "walk"
	}
	p.determineCollider()
	p.Y += p.DY
	p.CheckCollisionVertical(colliders)
}

func (p *Player) determineCollider() {
	collider := p.renderer.GetCollider(
		"Player",
		fmt.Sprintf("%s_%s", p.Action, p.Direction),
		"collider",
	)
	p.SetCollider(collider)
}

func (p *Player) printPlayer() {
	jPlayer, _ := json.Marshal(p)
	fmt.Println(string(jPlayer))
}
