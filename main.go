package main

import "fmt"

// const (
// 	screenWidth  = 300
// 	screenHeight = 300
// 	squareWidth  = 15
// )

// type Game struct{}

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }
// func (g *Game) Update() error {

// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	return
// }

func main() {

	// Create Square to test furthest point

	square := Polygon{Center: Point{0, 0}, Vertices: []Point{{2, 2}, {-2, -2}, {-2, 2}, {2, -2}}}

	square2 := Polygon{Center: Point{-4, -4}, Vertices: []Point{{-5, -5}, {-5, -3}, {-3, -5}, {-3, -3}}}

	fmt.Println((GJK(&square, &square2)))
}
