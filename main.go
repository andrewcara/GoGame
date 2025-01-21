package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	//"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	ball      = Circle{Center: Point{10, 10}, Radius: 15, VelX: float64(0.00000006), VelY: float64(0.00000004)}
	ball2     = Circle{Center: Point{20, 20}, Radius: 20, VelX: float64(0.00000004), VelY: float64(0.00000002)}
	loop_iter = 0

	prevUpdateTime = time.Now()
)

const (
	screenWidth  = 300
	screenHeight = 300
	squareWidth  = 15
)

type Game struct {
	Collision bool
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) Update() error {
	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	if GJK(&ball, &ball2) {
		g.Collision = true
	} else {
		g.Collision = false
	}

	ball.UpdateKinematics(screenWidth, screenHeight, timeDelta)
	ball2.UpdateKinematics(screenWidth, screenHeight, timeDelta)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.Collision {
		vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{100, 20, 3, 255}, false)
		vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{100, 20, 3, 255}, false)
		fmt.Println("collision!!!!", loop_iter)
	} else {

		vector.DrawFilledCircle(screen, float32(ball.Center.X), float32(ball.Center.Y), float32(ball.Radius), color.RGBA{255, 20, 3, 255}, false)
		vector.DrawFilledCircle(screen, float32(ball2.Center.X), float32(ball2.Center.Y), float32(ball2.Radius), color.RGBA{255, 20, 3, 255}, false)
	}
	loop_iter++
	return
}

func main() {

	// Create Square to test furthest point
	//square1 := Polygon{Center: Point{0, 0}, Vertices: []Point{{2, 2}, {-2, -2}, {-2, 2}, {2, -2}}}
	// square1 := Polygon{Center: Point{-2, -2}, Vertices: []Point{{-2.9, -2.9}, {-3, -1}, {-1, -1}, {-1, -3}}}

	// square2 := Polygon{Center: Point{-4, -4}, Vertices: []Point{{-5, -5}, {-5, -3}, {-3, -5}, {-3, -3}}}
	// //circle := Circle{Center: Point{4, 4}, Radius: 2}
	// fmt.Println((GJK(&square1, &square2)))

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")
	if err := ebiten.RunGame(&Game{Collision: false}); err != nil {
		log.Fatal(err)
	}
}
