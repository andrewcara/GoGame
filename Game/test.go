package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 300
	screenHeight = 300
	squareWidth  = 15
)

var (
	squarePositionX = float64(screenWidth) / 2
	squarePositionY = float64(screenHeight) / 2
	ballVelX        = float64(0.00000006)
	ballVelY        = float64(0.00000004)
	prevUpdateTime  = time.Now()
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) Update() error {

	//here we are taking into account the fact that the screen may update at different times depending on what the processor is dealing with
	//Ensure that the object moves smoothly regardless of time to update frame

	timeDelta := float64(time.Since(prevUpdateTime))
	prevUpdateTime = time.Now()

	squarePositionX += ballVelX * timeDelta
	squarePositionY += ballVelY * timeDelta

	const maxX = screenWidth - squareWidth
	const maxY = screenHeight - squareWidth
	const minX = squareWidth
	const minY = squareWidth

	if squarePositionX >= screenWidth-squareWidth || squarePositionX <= squareWidth {
		if squarePositionX > maxX {
			squarePositionX = maxX
		} else if squarePositionX < minX {
			squarePositionX = minX
		}
		ballVelX *= -1
	}
	if squarePositionY >= screenHeight-squareWidth || squarePositionY <= squareWidth {
		if squarePositionY > maxY {
			squarePositionY = maxY
		} else if squarePositionY < minY {
			squarePositionY = minY
		}

		ballVelY *= -1
	}

	return nil
}

func (g *Game) drawSquare(screen *ebiten.Image, x, y, radius int) {
	//Here we are denoting the center of the square by the x and y coordinate

	starting_x := x - (radius)
	starting_y := y - (radius)
	purpleCol := color.RGBA{255, 0, 255, 255}

	for x := starting_x; x < starting_x+(radius*2); x++ {
		for y := starting_y; y < starting_y+(radius*2); y++ {
			screen.Set(x, y, purpleCol)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	x := int(math.Round(squarePositionX))
	y := int(math.Round(squarePositionY))
	purpleCol := color.RGBA{255, 0, 255, 255}

	g.drawSquare(screen, x, y, squareWidth)
	screen.Set(x, y, purpleCol)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
