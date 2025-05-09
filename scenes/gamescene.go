package scenes

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	anir "github.com/talvor/tiled/animation/renderer"
	"github.com/talvor/tiled/common"
)

type GameScene struct {
	AnimationRenderer *anir.Renderer
	loaded            bool
	direction         string
}

func NewGameScene(animationRenderer *anir.Renderer) *GameScene {
	return &GameScene{
		AnimationRenderer: animationRenderer,
		loaded:            false,
		direction:         "right",
	}
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(16, 16)

	opts := &common.DrawOptions{
		Screen: screen,
		Op:     op,
	}

	err := s.AnimationRenderer.Draw("Player", fmt.Sprintf("walk_%s", s.direction), opts)
	if err != nil {
		log.Println("Error drawing animation:", err)
	}
}

func (s *GameScene) FirstLoad() {
	s.loaded = true
}

func (s *GameScene) IsLoaded() bool {
	return s.loaded
}

func (s *GameScene) OnEnter() {
}

func (s *GameScene) OnExit() {
}

func (s *GameScene) Update() SceneID {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneID
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneID
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s.direction = "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s.direction = "right"
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s.direction = "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s.direction = "down"
	}

	return GameSceneID
}

var _ Scene = (*GameScene)(nil)
