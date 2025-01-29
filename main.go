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
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	physics.PhyicsWorld
	Collision      bool
	accumulator    float64
	lastUpdateTime time.Time
}

func NewGame() *Game {
	return &Game{
		Collision:      false,
		lastUpdateTime: time.Now(),
		accumulator:    0,
	}
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
	temp_ball := physics.PhysicsObject{Shape: &ball}
	temp_ball2 := physics.PhysicsObject{Shape: &ball2}
	coll := physics.CollsionHandler{Collider: temp_ball}

	// Set gravity and forces to 0 for now
	// Check collision before updating positions
	if coll.HitsOtherObject(&temp_ball2) {
		g.Collision = true

		// // Move balls apart slightly to prevent stickin
	} else {
		g.Collision = false

	}
	ball.UpdateKinematics(screenWidth, screenHeight, timeDelta)
	ball2.UpdateKinematics(screenWidth, screenHeight, timeDelta)
}

// Update positions

func (g *Game) Draw(screen *ebiten.Image) {

	vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{200, 150, 3, 255}, false)
	vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{200, 150, 3, 255}, false)
	//fmt.Println("collision!!!!", loop_iter)

	// vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{255, 20, 3, 255}, false)
	// vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{255, 20, 3, 255}, false)

	return
}

func main() {

<<<<<<< Updated upstream
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")
	if err := ebiten.RunGame(&Game{Collision: false}); err != nil {
		log.Fatal(err)
	}
=======
	// Create Square to test furthest point

	square := Polygon{Center: Point{2.5, 2.5}, Vertices: []Point{{0, 5}, {5, 5}, {5, 0}, {0, 0}}}

	//square2 := Polygon{Center: Point{-4, -4}, Vertices: []Point{{-5, -5}, {-5, -3}, {-3, -5}, {-3, -3}}}
	circle := Circle{Center: Point{10, 10}, Radius: 2}
	fmt.Println((GJK(&square, &circle)))
>>>>>>> Stashed changes
}
