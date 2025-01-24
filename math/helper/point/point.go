package point

type Point struct {
	X, Y float64
}

func (p *Point) Add(p2 Point) Point {
	return (Point{X: p.X + p2.X, Y: p.Y + p2.Y})
}
func (p *Point) Subtract(p2 Point) Point {
	return (Point{X: p.X - p2.X, Y: p.Y - p2.Y})
}
func (p *Point) Multiply(p2 Point) Point {
	return (Point{X: p.X * p2.X, Y: p.Y * p2.Y})
}
