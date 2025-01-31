package main

import (
	player "HeadSoccer/input"
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Gravity is positve since "down" on the screen is positive and up is negative
var gravity = linalg.Vector{X: 0, Y: 98.81}

const (
	physicsTickRate   = 1.0 / 100
	screenWidth       = 600
	screenHeight      = 300
	squareWidth       = 15
	maxSteps          = 3
	moveInputVelocity = 10
)

type Game struct {
	world          physics.PhysicsWorld
	Collision      bool
	accumulator    float64
	lastUpdateTime time.Time
	pressedKeys    []ebiten.Key
}

func NewGame() *Game {
	// Ball initialization - centered more visibly
	ball := physics.PhysicsBody{
		Shape: &shapes.Circle{
			Center: linalg.Point{X: screenWidth / 2, Y: screenHeight / 2},
			Radius: 15,
		},
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         5.0,
			Acceleration: gravity,
		},
	}

	// Player 1 - Left side of screen
	player1Center := linalg.Point{X: 20, Y: 20}
	// Define vertices relative to center
	vertices1 := []linalg.Point{
		{X: 10, Y: 10}, // top left
		{X: 10, Y: 30}, // top right
		{X: 30, Y: 30}, // bottom left
		{X: 30, Y: 10}, // bottom right
	}
	var polygon1 shapes.Polygon
	polygon1.Initialize(player1Center, vertices1)

	player1Body := physics.PhysicsBody{
		Shape: &polygon1,
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         5.0,
			Acceleration: gravity,
		},
	}

	// Player 2 - Right side of screen

	// Initialize physics world
	world := physics.PhysicsWorld{
		Objects: make([]*physics.PhysicsBody, 0),
		Gravity: gravity,
	}

	game := &Game{
		world:          world,
		Collision:      false,
		lastUpdateTime: time.Now(),
		accumulator:    0,
	}

	// Add Physics Objects to Game
	game.world.Objects = append(game.world.Objects,
		&player1Body,
		//&player2Body,
		&ball,
	)

	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) HanlePlayerInput(player_moves player.PlayerMoves) {

	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	player1 := &g.world.Objects[player_moves.PlayerID]
	for _, key := range g.pressedKeys {
		switch key.String() {
		case player_moves.Down:
			move := linalg.Vector{X: 0, Y: moveInputVelocity}
			(*player1).SetVelocity(move.Add((*player1).GetVelocity()))
		case player_moves.Up:
			move := linalg.Vector{X: 0, Y: -moveInputVelocity}
			(*player1).SetVelocity(move.Add((*player1).GetVelocity()))
		case player_moves.Right:
			move := linalg.Vector{X: moveInputVelocity, Y: 0}
			(*player1).SetVelocity(move.Add((*player1).GetVelocity()))
		case player_moves.Left:
			move := linalg.Vector{X: -moveInputVelocity, Y: 0}
			(*player1).SetVelocity(move.Add((*player1).GetVelocity()))
		}
	}
}

func (g *Game) HandleUserInput() {
	player1_moves := player.PlayerMoves{Up: "ArrowUp", Down: "ArrowDown", Left: "ArrowLeft", Right: "ArrowRight", PlayerID: 0}
	g.HanlePlayerInput(player1_moves)

}

func (g *Game) Update() error {

	g.HandleUserInput()
	currentTime := time.Now()
	frameTime := currentTime.Sub(g.lastUpdateTime).Seconds()
	g.lastUpdateTime = currentTime

	if frameTime > 0.25 {
		frameTime = 0.25
	}

	g.accumulator += frameTime
	steps := 0

	for g.accumulator >= physicsTickRate && steps < maxSteps {
		g.UpdatePhysics(physicsTickRate)
		g.accumulator -= physicsTickRate
		steps++
	}

	return nil
}

//Current implementation is that there must be two objects

func (g *Game) UpdatePhysics(timeDelta float64) {
	for i := 0; i < len(g.world.Objects); i++ {
		for j := i + 1; j < len(g.world.Objects); j++ {
			obj1 := g.world.Objects[i]
			obj2 := g.world.Objects[j]

			if physics.CollisionOccurs(obj1, obj2) {
				g.Collision = true
			} else {
				g.Collision = false
			}
			obj1.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
			obj2.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
		}
	}
}

// Update positions

func (g *Game) Draw(screen *ebiten.Image) {
	color := color.RGBA{200, 150, 3, 255}
	for _, obj := range g.world.Objects {
		obj.Shape.DrawShape(screen, color)
	}
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
