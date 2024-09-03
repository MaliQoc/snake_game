package model

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	SnakeSize    = 10
	InitialSpeed = 10
	NumObstacles = 5
	Margin       = 20
)

type Point struct {
	X, Y int
}

type Game struct {
	Snake        []Point
	Direction    Point
	Food         Point
	Obstacles    []Point
	GameOver     bool
	Score        int
	Speed        int
	FrameCounter int
}

func (g *Game) Update() {
	if g.GameOver {
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.Direction.Y == 0 {
		g.Direction = Point{X: 0, Y: -SnakeSize}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.Direction.Y == 0 {
		g.Direction = Point{X: 0, Y: SnakeSize}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.Direction.X == 0 {
		g.Direction = Point{X: -SnakeSize, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.Direction.X == 0 {
		g.Direction = Point{X: SnakeSize, Y: 0}
	}

	g.FrameCounter++
	if g.FrameCounter < g.Speed {
		return
	}
	g.FrameCounter = 0

	head := g.Snake[0]
	newHead := Point{X: head.X + g.Direction.X, Y: head.Y + g.Direction.Y}

	if newHead.X < 0 || newHead.X >= ScreenWidth || newHead.Y < 0 || newHead.Y >= ScreenHeight {
		g.GameOver = true
	}

	for _, segment := range g.Snake[1:] {
		if segment == newHead {
			g.GameOver = true
		}
	}

	for _, obstacle := range g.Obstacles {
		if newHead == obstacle {
			g.GameOver = true
		}
	}

	if newHead == g.Food {
		g.Snake = append([]Point{newHead}, g.Snake...)
		g.Score++
		g.placeFood()
		g.placeObstacles()

		if g.Speed > 2 {
			g.Speed--
		}
	} else {
		g.Snake = append([]Point{newHead}, g.Snake[:len(g.Snake)-1]...)
	}
}

func (g *Game) placeFood() {
	for {
		g.Food = Point{
			X: rand.Intn((ScreenWidth-Margin)/SnakeSize) * SnakeSize,
			Y: rand.Intn((ScreenHeight-Margin)/SnakeSize) * SnakeSize,
		}

		occupied := false
		for _, obstacle := range g.Obstacles {
			if g.Food == obstacle {
				occupied = true
				break
			}
		}
		for _, segment := range g.Snake {
			if g.Food == segment {
				occupied = true
				break
			}
		}
		if !occupied {
			break
		}
	}
}

func (g *Game) placeObstacles() {
	g.Obstacles = nil
	for i := 0; i < NumObstacles; i++ {
		var obstacle Point
		for {
			obstacle = Point{
				X: rand.Intn((ScreenWidth-Margin)/SnakeSize) * SnakeSize,
				Y: rand.Intn((ScreenHeight-Margin)/SnakeSize) * SnakeSize,
			}
			occupied := false
			for _, segment := range g.Snake {
				if obstacle == segment {
					occupied = true
					break
				}
			}
			if obstacle == g.Food {
				occupied = true
			}
			if !occupied {
				break
			}
		}
		g.Obstacles = append(g.Obstacles, obstacle)
	}
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		Snake:     []Point{{X: ScreenWidth / 2, Y: ScreenHeight / 2}},
		Direction: Point{X: SnakeSize, Y: 0},
		Speed:     InitialSpeed,
		GameOver:  false,
	}
	game.placeFood()
	game.placeObstacles()
	return game
}
