package view

import (
	"image/color"
	"strconv"

	"snake/model"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw(game *model.Game, screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	for _, segment := range game.Snake {
		ebitenutil.DrawRect(screen, float64(segment.X), float64(segment.Y), model.SnakeSize, model.SnakeSize, color.RGBA{0, 255, 0, 255})
	}

	ebitenutil.DrawRect(screen, float64(game.Food.X), float64(game.Food.Y), model.SnakeSize, model.SnakeSize, color.RGBA{255, 0, 0, 255})

	for _, obstacle := range game.Obstacles {
		ebitenutil.DrawRect(screen, float64(obstacle.X), float64(obstacle.Y), model.SnakeSize, model.SnakeSize, color.RGBA{255, 255, 0, 255})
	}

	if game.GameOver {
		ebitenutil.DebugPrint(screen, "Game Over!\nScore: "+strconv.Itoa(game.Score))
	} else {
		ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(game.Score))
	}
}
