package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/talvor/go-rpg/entities"
	"github.com/talvor/go-rpg/scenes"
	"github.com/talvor/go-rpg/utils"
	"github.com/talvor/tiled"
)

type Game struct {
	sceneMap      map[scenes.SceneID]scenes.Scene
	activeSceneId scenes.SceneID
}

func NewGame() *Game {
	animationRenderer := tiled.NewAnimationRenderer([]string{
		utils.GetAssetsDirectory("animations"),
	}, []string{
		utils.GetAssetsDirectory("tilesets", "Player"),
	})

	mapRenderer := tiled.NewMapRenderer([]string{
		utils.GetAssetsDirectory("maps"),
	}, []string{
		utils.GetAssetsDirectory("tilesets", "Tiles"),
	})

	player := entities.NewPlayer(16, 16, animationRenderer)
	sceneMap := map[scenes.SceneID]scenes.Scene{
		scenes.GameSceneID:  scenes.NewGameScene(animationRenderer, mapRenderer, player),
		scenes.StartSceneID: scenes.NewStartScene(),
		scenes.PauseSceneID: scenes.NewPauseScene(),
	}
	activeSceneId := scenes.StartSceneID
	sceneMap[activeSceneId].FirstLoad()
	return &Game{
		sceneMap:      sceneMap,
		activeSceneId: activeSceneId,
	}
}

func (g *Game) Update() error {
	nextSceneID := g.sceneMap[g.activeSceneId].Update()
	// switched scenes
	if nextSceneID == scenes.ExitSceneID {
		g.sceneMap[g.activeSceneId].OnExit()
		return ebiten.Termination
	}
	if nextSceneID != g.activeSceneId {
		nextScene := g.sceneMap[nextSceneID]
		// if not loaded? then load in
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad()
		}
		nextScene.OnEnter()
		g.sceneMap[g.activeSceneId].OnExit()
	}
	g.activeSceneId = nextSceneID
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneMap[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
