package entities

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/talvor/go-rpg/camera"
)

type margin struct {
	Top    int
	Left   int
	Right  int
	Bottom int
}
type Sprite struct {
	X, Y, DX, DY float64
	Tilesize     int
	Collider     *image.Rectangle
	Margin       *margin
}

func NewSprite(x, y float64, tilesize int) *Sprite {
	m := &margin{
		Top:    0,
		Left:   0,
		Right:  0,
		Bottom: 0,
	}

	return &Sprite{
		X:        x,
		Y:        y,
		Tilesize: tilesize,
		Margin:   m,
	}
}

func (s *Sprite) SetCollider(collider *image.Rectangle) {
	s.Collider = collider
	s.Margin.Top = collider.Min.Y
	s.Margin.Left = collider.Min.X
	s.Margin.Right = s.Tilesize - collider.Max.X
	s.Margin.Bottom = s.Tilesize - collider.Max.Y
}

func (s *Sprite) GetColliderRect() image.Rectangle {
	colliderRect := image.Rect(
		int(s.X)+s.Collider.Min.X,
		int(s.Y)+s.Collider.Min.Y,
		int(s.X)+s.Collider.Max.X,
		int(s.Y)+s.Collider.Max.Y,
	)
	return colliderRect
}

func (s *Sprite) DrawColliderRect(dst *ebiten.Image, cam *camera.Camera) {
	cr := s.GetColliderRect()
	vector.StrokeRect(
		dst,
		float32(cr.Min.X)+float32(cam.X),
		float32(cr.Min.Y)+float32(cam.Y),
		float32(cr.Dx()),
		float32(cr.Dy()),
		0.5,
		color.RGBA{255, 0, 0, 255},
		true,
	)
}

func (s *Sprite) DrawSpriteRect(dst *ebiten.Image, cam *camera.Camera) {
	vector.StrokeRect(
		dst,
		float32(s.X)+float32(cam.X),
		float32(s.Y)+float32(cam.Y),
		float32(s.Tilesize),
		float32(s.Tilesize),
		0.5,
		color.RGBA{0, 255, 0, 255},
		true,
	)
}

func (s *Sprite) CheckCollisionHorizontal(colliders []image.Rectangle) {
	colliderRect := s.GetColliderRect()

	for _, collider := range colliders {
		if collider.Overlaps(
			colliderRect,
		) {
			if s.DX < 0.0 {
				// Moving left
				s.X = float64(collider.Max.X - s.Margin.Left + 1)
			} else if s.DX > 0.0 {
				// Moving right
				s.X = float64(collider.Min.X - s.Tilesize + s.Margin.Right - 1)
			}
		}
	}
}

func (s *Sprite) CheckCollisionVertical(colliders []image.Rectangle) {
	colliderRect := s.GetColliderRect()
	for _, collider := range colliders {
		if collider.Overlaps(
			colliderRect,
		) {
			if s.DY < 0.0 {
				// Moving up
				s.Y = float64(collider.Max.Y - s.Margin.Top + 1)
			} else if s.DY > 0.0 {
				// Moving down
				s.Y = float64(collider.Min.Y - s.Tilesize + s.Margin.Bottom - 1)
			}
		}
	}
}

func (s *Sprite) printSprite() {
	jSprite, _ := json.Marshal(s)
	fmt.Println(string(jSprite))
}
