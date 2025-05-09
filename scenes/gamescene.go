package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/talvor/go-rpg/entities"
	anir "github.com/talvor/tiled/animation/renderer"
	"github.com/talvor/tiled/common"
	tmxr "github.com/talvor/tiled/tmx/renderer"
)

type GameScene struct {
	AnimationRenderer *anir.Renderer
	MapRenderer       *tmxr.Renderer
	loaded            bool
	direction         string
	player            *entities.Player
}

func NewGameScene(animationRenderer *anir.Renderer, mapRenderer *tmxr.Renderer, player *entities.Player) *GameScene {
	return &GameScene{
		AnimationRenderer: animationRenderer,
		MapRenderer:       mapRenderer,
		loaded:            false,
		direction:         "right",
		player:            player,
	}
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := &ebiten.DrawImageOptions{}
	op := &common.DrawOptions{Screen: screen, Op: opts}

	s.MapRenderer.DrawMapLayer("GameScene", "background", op)

	s.player.Draw(screen)
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

	s.player.Update()

	return GameSceneID
}

var _ Scene = (*GameScene)(nil)
