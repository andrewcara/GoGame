package main

import (
	player "HeadSoccer/input"
	linalg "HeadSoccer/math/helper"
	"sync"

	"HeadSoccer/initialization"
	"HeadSoccer/math/physics"
	"time"

	"github.com/google/uuid"
)

// Gravity is positve since "down" on the screen is positive and up is negative
var gravity = linalg.Vector{X: 0, Y: 100.81}
var id = uuid.New()

const (
	physicsTickRate   = 1.0 / 100
	screenWidth       = 600
	screenHeight      = 300
	squareWidth       = 15
	maxSteps          = 3
	moveInputVelocity = 15
	moveForce         = 500.0 // Base movement force
	maxSpeed          = 200.0 // Maximum speed cap

)

type Game struct {
	world          physics.PhysicsWorld
	Collision      bool
	accumulator    float64
	lastUpdateTime time.Time
	mu             sync.Mutex
}

func NewGame() *Game {

	//Initalize crossbars

	world := initialization.Setup(screenWidth, screenHeight, gravity)

	game := &Game{
		world:          world,
		Collision:      false,
		lastUpdateTime: time.Now(),
		accumulator:    0,
	}
	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) HanlePlayerInput(player_moves player.PlayerMoves, move string) {
	player := &g.world.Objects[player_moves.PlayerID]
	boundary := (*player).Shape.GetBoundaryPoints()
	onGround := boundary.MaxY >= float64(screenHeight)-1
	switch move {
	case player_moves.Down:
		move := linalg.Vector{X: 0, Y: moveInputVelocity}
		(*player).SetVelocity(move.Add((*player).GetVelocity()))
	case player_moves.Up:
		if onGround {
			move := linalg.Vector{X: 0, Y: -100}
			(*player).SetVelocity(move.Add((*player).GetVelocity()))
		}
	case player_moves.Right:
		if onGround {
			move := linalg.Vector{X: moveInputVelocity, Y: 0}
			(*player).SetVelocity(move.Add((*player).GetVelocity()))
		}
	case player_moves.Left:
		if onGround {
			move := linalg.Vector{X: -moveInputVelocity, Y: 0}
			(*player).SetVelocity(move.Add((*player).GetVelocity()))
		}
	}
}

func (g *Game) HandleUserInput(ID int, move string) {
	player_moves := player.PlayerMoves{Up: "ArrowUp", Down: "ArrowDown", Left: "ArrowLeft", Right: "ArrowRight", PlayerID: ID}
	g.HanlePlayerInput(player_moves, move)

}

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
			if !obj1.IsStatic {
				obj1.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
			}
			if !obj2.IsStatic {
				obj2.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)

			}
		}
	}
}
