package main

import (
	"log"

	"snake/controller"
	"snake/model"
	"snake/view"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	*model.Game
}

func (g *Game) Update() error {
	controller.Update(g.Game)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	view.Draw(g.Game, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return model.ScreenWidth, model.ScreenHeight
}

func main() {
	game := model.NewGame()
	g := &Game{Game: game}

	ebiten.SetWindowSize(model.ScreenWidth*2, model.ScreenHeight*2)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
