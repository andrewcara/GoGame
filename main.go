package main

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"fmt"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
	"image/color"
	"log"
	"time"

	//"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2"
)

var gravity = linalg.Vector{X: 0, Y: -9.81}

var (
	ball = shapes.Circle{Center: linalg.Point{X: 100, Y: 30}, Radius: 15,
		Dynamic: dynamics.DynamicProperties{
			Velocity:   linalg.Vector{X: float64(150), Y: float64(0.00000006)}, // Example velocity
			Force:      linalg.Vector{X: 0, Y: -9.8},                           // Gravity force
			Mass:       1.0,
			Accelation: gravity, // Example mass
		}}
	ball2 = shapes.Circle{Center: linalg.Point{X: 20, Y: 20}, Radius: 20,
		Dynamic: dynamics.DynamicProperties{
			Velocity:   linalg.Vector{X: float64(160), Y: float64(0.00000006)}, // Example velocity
			Force:      linalg.Vector{X: 0, Y: -9.8},                           // Gravity force
			Mass:       1.0,
			Accelation: gravity, // Example mass
		}}

	loop_iter = 0

	prevUpdateTime = time.Now()
)

const (
	physicsTickRate = 1.0 / 100
	screenWidth     = 300
	screenHeight    = 300
	squareWidth     = 15
	maxSteps        = 3
)

type Game struct {
	world          physics.PhysicsWorld
	Collision      bool
	accumulator    float64
	lastUpdateTime time.Time
}

func NewGame() *Game {

	initialCenter := linalg.Point{X: 125, Y: 125}
	vertices := []linalg.Point{
		{X: 150, Y: 100},
		{X: 150, Y: 150},
		{X: 100, Y: 150},
		{X: 100, Y: 100},
	}

	dynamicProps := dynamics.DynamicProperties{
		Velocity: linalg.Vector{X: 50, Y: 30},
		Force:    linalg.Vector{X: 0, Y: -9.8}, // Example gravity force
		Mass:     5.0,
	}

	var polygon shapes.Polygon
	polygon.Initialize(initialCenter, vertices, dynamicProps)

	world := physics.PhysicsWorld{
		Objects: make([]shapes.Shape, 0),
		Gravity: gravity,
	}

	game := &Game{
		world:          world,
		Collision:      false,
		lastUpdateTime: time.Now(),
		accumulator:    0,
	}

	// Wrap shapes in PhysicsObjects
	game.world.Objects = append(game.world.Objects,
		&ball,
		&ball2,
		&polygon,
	)

	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) Update() error {
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

func (g *Game) UpdatePhysics(timeDelta float64) {
	for i := 0; i < len(g.world.Objects); i++ {
		for j := i + 1; j < len(g.world.Objects); j++ {
			obj1 := g.world.Objects[i]
			obj2 := g.world.Objects[j]

			if physics.HitsOtherObject(&obj1, &obj2) {
				g.Collision = true
			} else {
				g.Collision = false
			}
			fmt.Println(obj1.GetMass(), obj2.GetMass())
			obj1.UpdateKinematics(screenWidth, screenHeight, timeDelta)
			obj2.UpdateKinematics(screenWidth, screenHeight, timeDelta)
		}
	}
}

// Update positions

func (g *Game) Draw(screen *ebiten.Image) {
	color := color.RGBA{200, 150, 3, 255}
	for _, obj := range g.world.Objects {
		obj.DrawShape(screen, color)
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
