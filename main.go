package main

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
	"image/color"
	"log"
	"time"

	//"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2"
)

// Gravity is positve since "down" on the screen is positive and up is negative
var gravity = linalg.Vector{X: 0, Y: 98.81}

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

	ball := shapes.Circle{Center: linalg.Point{X: 100, Y: 30}, Radius: 15,
		Dynamic: dynamics.DynamicProperties{
			Velocity:   linalg.Vector{X: float64(150), Y: float64(100)}, // Example velocity
			Force:      linalg.Vector{X: 0, Y: -9.8},                    // Gravity force
			Mass:       1.0,
			Accelation: gravity, // Example mass
		}}
	ball2 := shapes.Circle{Center: linalg.Point{X: 20, Y: 20}, Radius: 20,
		Dynamic: dynamics.DynamicProperties{
			Velocity:   linalg.Vector{X: float64(260), Y: float64(-100)}, // Example velocity
			Force:      linalg.Vector{X: 0, Y: -9.8},                     // Gravity force
			Mass:       1.0,
			Accelation: gravity, // Example mass
		}}

	_ = ball
	initialCenter := linalg.Point{X: 100, Y: 100}
	vertices := []linalg.Point{
		{X: 130, Y: 130},
		{X: 150, Y: 100},
		{X: 140, Y: 70},
		{X: 110, Y: 50},
		{X: 80, Y: 60},
		{X: 60, Y: 90},
		{X: 70, Y: 120},
		{X: 100, Y: 140},
		{X: 120, Y: 150},
	}
	dynamicProps := dynamics.DynamicProperties{
		Velocity: linalg.Vector{X: 20, Y: 15},
		Force:    linalg.Vector{X: 0, Y: -9.8}, // Example gravity force
		Mass:     10.0,
	}

	var polygon shapes.Polygon
	polygon.Initialize(initialCenter, vertices, dynamicProps)

	initialCenter2 := linalg.Point{X: 50, Y: 50}
	vertices2 := []linalg.Point{
		{X: 75, Y: 75},
		{X: 75, Y: 25},
		{X: 25, Y: 25},
		{X: 25, Y: 75},
	}

	dynamicProps2 := dynamics.DynamicProperties{
		Velocity: linalg.Vector{X: 50, Y: 30},
		Force:    linalg.Vector{X: 0, Y: -9.8}, // Example gravity force
		Mass:     5.0,
	}

	var polygon2 shapes.Polygon
	polygon2.Initialize(initialCenter2, vertices2, dynamicProps2)

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

	// Add Physic Objects to Game
	game.world.Objects = append(game.world.Objects,
		&ball,
		&ball2,
		&polygon,
		&polygon2,
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

//Current implementation is that there must be two objects

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
			obj1.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
			obj2.UpdateKinematics(screenWidth, screenHeight, timeDelta, gravity)
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
