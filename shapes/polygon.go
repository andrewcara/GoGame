package shapes

import (
	linalg "HeadSoccer/math/helper"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Polygon struct {
	Center   Point
	Vertices []Point
	Offsets  []Point
	//Velocity to be implemented as a vector
}

// Find the Point that is the furthest from the Center given a directional vector
func (p *Polygon) FurthestPoint(direction_vector linalg.Vector) Point {
	// Ensure vertices are up to date
	maximum_dot_product := linalg.DotProduct(linalg.NewVector(p.Center, p.Vertices[0]), direction_vector)
	maximum_point := p.Vertices[0]

	for _, point := range p.Vertices[1:] {
		local_dot_product := linalg.DotProduct(direction_vector, linalg.NewVector(p.Center, point))
		if local_dot_product > maximum_dot_product {
			maximum_point = point
			maximum_dot_product = local_dot_product
		}
	}
	return maximum_point
}

func (p *Polygon) updateVertices() {
	p.Vertices = make([]Point, len(p.Offsets))
	for i, offset := range p.Offsets {
		p.Vertices[i] = Point{
			X: p.Center.X + offset.X,
			Y: p.Center.Y + offset.Y,
		}
	}
}

func (p *Polygon) Initialize(center Point, vertices []Point) {
	p.Center = center
	p.Vertices = vertices
	p.calculateOffsets() // Calculate the offsets based on the initial vertices// Set dynamic properties
}

func (p *Polygon) calculateOffsets() {
	p.Offsets = make([]Point, len(p.Vertices)) // Exclude the center
	for i, vertex := range p.Vertices {
		// Subtract the center from each vertex to calculate the offset
		p.Offsets[i] = Point{X: vertex.X - p.Center.X, Y: vertex.Y - p.Center.Y}
	}
}

func (p *Polygon) GetSurfacePoint(direction_vector linalg.Vector) Point {
	// First get furthest point in this direction as a starting point
	maxPoint := p.FurthestPoint(direction_vector)
	maxDist := linalg.DotProduct(direction_vector, linalg.NewVector(p.Center, maxPoint))

	// Check each edge to see if collision point lies between vertices
	for i := 0; i < len(p.Vertices); i++ {
		v1 := p.Vertices[i]
		v2 := p.Vertices[(i+1)%len(p.Vertices)]

		// Get edge vector
		edge := linalg.NewVector(v1, v2)

		// If edge is perpendicular to direction vector, it might contain contact point
		edgeNormal := linalg.Vector{X: -edge.Y, Y: edge.X}.Normalize()
		if linalg.DotProduct(edgeNormal, direction_vector) > 0 {
			// Project direction onto edge to find potential contact point
			t := linalg.DotProduct(direction_vector, edge) / linalg.DotProduct(edge, edge)

			// If projection lies between vertices (0 <= t <= 1)
			if t >= 0 && t <= 1 {
				contactPoint := Point{
					X: v1.X + edge.X*t,
					Y: v1.Y + edge.Y*t,
				}

				// Check if this point is further in the direction vector
				dist := linalg.DotProduct(direction_vector, linalg.NewVector(p.Center, contactPoint))
				if dist > maxDist {
					maxDist = dist
					maxPoint = contactPoint
				}
			}
		}
	}

	return maxPoint
}

func (p *Polygon) GetCenter() Point {
	return p.Center
}

func (p *Polygon) SetCenter(point Point) {
	p.Center = point
	p.updateVertices()
}

func (p *Polygon) DrawShape(screen *ebiten.Image, color color.RGBA) {
	for i := 0; i < len(p.Vertices); i++ {
		v1 := p.Vertices[i]
		v2 := p.Vertices[(i+1)%len(p.Vertices)]
		drawLine(screen, v1, v2, color)
	}
}

func drawLine(screen *ebiten.Image, p1, p2 Point, clr color.Color) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	steps := math.Max(math.Abs(dx), math.Abs(dy))

	if steps == 0 {
		screen.Set(int(p1.X), int(p1.Y), clr)
		return
	}

	xIncrement := dx / steps
	yIncrement := dy / steps

	x := p1.X
	y := p1.Y
	for i := float64(0); i <= steps; i++ {
		screen.Set(int(math.Round(x)), int(math.Round(y)), clr)
		x += xIncrement
		y += yIncrement
	}
}

// GetBoundaryPoints returns the extreme points of the polygon
func (p *Polygon) GetBoundaryPoints() BoundaryPoints {
	if len(p.Vertices) == 0 {
		return BoundaryPoints{}
	}

	minX := p.Vertices[0].X
	maxX := p.Vertices[0].X
	minY := p.Vertices[0].Y
	maxY := p.Vertices[0].Y

	for _, vertex := range p.Vertices[1:] {
		minX = math.Min(minX, vertex.X)
		maxX = math.Max(maxX, vertex.X)
		minY = math.Min(minY, vertex.Y)
		maxY = math.Max(maxY, vertex.Y)
	}

	return BoundaryPoints{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
	}
}
