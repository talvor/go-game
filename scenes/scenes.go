package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneID uint

const (
	StartSceneID SceneID = iota
	GameSceneID
	PauseSceneID
	ExitSceneID
)

type Scene interface {
	Update() SceneID
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
	IsLoaded() bool
}
