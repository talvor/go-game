package scenes

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/talvor/go-rpg/camera"
	"github.com/talvor/go-rpg/entities"
	anir "github.com/talvor/tiled/animation/renderer"
	"github.com/talvor/tiled/common"
	tmx "github.com/talvor/tiled/tmx"
	tmxr "github.com/talvor/tiled/tmx/renderer"
)

type GameScene struct {
	AnimationRenderer *anir.Renderer
	MapRenderer       *tmxr.Renderer
	loaded            bool
	direction         string
	player            *entities.Player
	cam               *camera.Camera
	gameMap           *tmx.Map
	colliders         []image.Rectangle
}

func NewGameScene(animationRenderer *anir.Renderer, mapRenderer *tmxr.Renderer, player *entities.Player) *GameScene {
	return &GameScene{
		AnimationRenderer: animationRenderer,
		MapRenderer:       mapRenderer,
		loaded:            false,
		direction:         "right",
		player:            player,
		cam:               nil,
		gameMap:           nil,
		colliders:         make([]image.Rectangle, 0),
	}
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := &ebiten.DrawImageOptions{}
	op := &common.DrawOptions{Screen: screen, Op: opts, OffsetX: s.cam.X, OffsetY: s.cam.Y}

	s.MapRenderer.DrawMapLayer("GameScene", "background", op)

	s.player.Draw(screen, s.cam)

	// for _, collider := range s.colliders {
	// 	vector.StrokeRect(
	// 		screen,
	// 		float32(collider.Min.X)+float32(s.cam.X),
	// 		float32(collider.Min.Y)+float32(s.cam.Y),
	// 		float32(collider.Dx()),
	// 		float32(collider.Dy()),
	// 		1.0,
	// 		color.RGBA{255, 0, 0, 255},
	// 		true,
	// 	)
	// }
}

func (s *GameScene) FirstLoad() {
	s.cam = camera.NewCamera(0, 0)
	gameMap, err := s.MapRenderer.MapManager.GetMapByName("GameScene")
	if err != nil {
		panic(err)
	}
	s.gameMap = gameMap

	// Register map edges as colliders
	mapWidth := s.gameMap.Width * s.gameMap.TileWidth
	mapHeight := s.gameMap.Height * s.gameMap.TileHeight
	mapEdges := []image.Rectangle{
		image.Rect(-1, -1, 1, mapHeight+1),                   // left
		image.Rect(-1, -1, mapWidth+1, 1),                    // top
		image.Rect(mapWidth-1, -1, mapWidth+1, mapHeight+1),  // right
		image.Rect(-1, mapHeight-1, mapWidth+1, mapHeight+1), // bottom
	}
	s.colliders = append(s.colliders, mapEdges...)

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

	s.player.Update(s.colliders)

	s.cam.FollowTarget(s.player.X+8, s.player.Y+8, 640, 480)
	s.cam.Constrain(
		float64(s.gameMap.Width*s.gameMap.TileWidth),
		float64(s.gameMap.Height*s.gameMap.TileHeight),
		640,
		480,
	)

	return GameSceneID
}

var _ Scene = (*GameScene)(nil)
