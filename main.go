package main

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
	"fmt"
	"image/color"
	"log"
	"time"

	//"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	ball = shapes.Circle{Center: linalg.Point{X: 100, Y: 30}, Radius: 15,
		Dynamic: dynamics.DynamicProperties{
			Velocity: linalg.Vector{X: float64(0.00000006), Y: float64(0.00000006)}, // Example velocity
			Force:    linalg.Vector{X: 0, Y: -9.8},                                  // Gravity force
			Mass:     1.0,                                                           // Example mass
		}}
	ball2 = shapes.Circle{Center: linalg.Point{X: 20, Y: 20}, Radius: 20,
		Dynamic: dynamics.DynamicProperties{
			Velocity: linalg.Vector{X: float64(0.00000006), Y: float64(0.00000006)}, // Example velocity
			Force:    linalg.Vector{X: 0, Y: -9.8},                                  // Gravity force
			Mass:     1.0,                                                           // Example mass
		}}
	loop_iter = 0

	prevUpdateTime = time.Now()
)

const (
	screenWidth  = 300
	screenHeight = 300
	squareWidth  = 15
)

type Game struct {
	physics.PhyicsWorld
	Collision bool
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) Update() error {
	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	temp_ball := physics.PhysicsObject{Shape: &ball}
	temp_ball2 := physics.PhysicsObject{Shape: &ball2}

	coll := physics.CollsionHandler{Collider: temp_ball}

	ball.UpdateKinematics(screenWidth, screenHeight, timeDelta)

	if coll.HitsOtherObject(&temp_ball2) {
		println("collision")
	} else {
		ball2.UpdateKinematics(screenWidth, screenHeight, timeDelta)
		g.Collision = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.Collision {
		vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{200, 150, 3, 255}, false)
		vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{200, 150, 3, 255}, false)
		fmt.Println("collision!!!!", loop_iter)
	} else {

		vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{255, 20, 3, 255}, false)
		vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{255, 20, 3, 255}, false)
	}
	loop_iter++
	return
}

func main() {

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")
	if err := ebiten.RunGame(&Game{Collision: false}); err != nil {
		log.Fatal(err)
	}
}
