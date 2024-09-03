package controller

import (
	"snake/model"

	"github.com/hajimehoshi/ebiten/v2"
)

func Update(game *model.Game) {
	if game.GameOver {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && game.Direction.Y == 0 {
		game.Direction = model.Point{X: 0, Y: -model.SnakeSize}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && game.Direction.Y == 0 {
		game.Direction = model.Point{X: 0, Y: model.SnakeSize}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && game.Direction.X == 0 {
		game.Direction = model.Point{X: -model.SnakeSize, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && game.Direction.X == 0 {
		game.Direction = model.Point{X: model.SnakeSize, Y: 0}
	}

	game.Update()
}
