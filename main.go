package main

import (
	player "HeadSoccer/input"
	linalg "HeadSoccer/math/helper"

	"HeadSoccer/initialization"
	"HeadSoccer/math/physics"
	"image/color"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	pressedKeys    []ebiten.Key
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

	// Add Physics Objects to Game
	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) HanlePlayerInput(player_moves player.PlayerMoves) {
	player := &g.world.Objects[player_moves.PlayerID]
	boundary := (*player).Shape.GetBoundaryPoints()
	onGround := boundary.MaxY >= float64(screenHeight)-1
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, key := range g.pressedKeys {
		switch key.String() {
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
}

func (g *Game) HandleUserInput() {
	player1_moves := player.PlayerMoves{Up: "ArrowUp", Down: "ArrowDown", Left: "ArrowLeft", Right: "ArrowRight", PlayerID: 0}
	player2_moves := player.PlayerMoves{Up: "W", Down: "S", Left: "A", Right: "D", PlayerID: 1}
	g.HanlePlayerInput(player1_moves)
	g.HanlePlayerInput(player2_moves)

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
			if !obj1.IsStatic {
				obj1.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
			}
			if !obj2.IsStatic {
				obj2.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)

			}
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
