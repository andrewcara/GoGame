package shapes

import (
	"HeadSoccer/Sprites"
	linalg "HeadSoccer/math/helper"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	coefficient_friction  = 0.8
	collision_restitution = 0.8
)

type Circle struct {
	Center Point
	Radius float64
	Image  *ebiten.Image
}

func (c *Circle) FurthestPoint(direction_vector linalg.Vector) Point {
	normalized_vec := direction_vector.Normalize()
	scaled_vec := normalized_vec.Scale(c.Radius)
	return Point{X: c.Center.X + scaled_vec.X, Y: c.Center.Y + scaled_vec.Y}
}

func (c *Circle) GetCenter() Point {
	return c.Center
}

func (c *Circle) SetCenter(point Point) {
	c.Center = point
}
func (c *Circle) GetSurfacePoint(direction_vector linalg.Vector) Point {
	normalized_vec := direction_vector.Normalize()
	scaled_vec := normalized_vec.Scale(c.Radius)
	return Point{X: c.Center.X + scaled_vec.X, Y: c.Center.Y + scaled_vec.Y}
}
func (c *Circle) GetBoundaryPoints() BoundaryPoints {
	minX := c.Center.Subtract(Point{X: c.Radius, Y: 0}).X
	maxX := c.Center.Add(Point{X: c.Radius, Y: 0}).X
	minY := c.Center.Subtract(Point{X: 0, Y: c.Radius}).Y
	maxY := c.Center.Add(Point{X: 0, Y: c.Radius}).Y
	return BoundaryPoints{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
	}
}

func (c *Circle) SetImage(image_path string) {
	c.Image = Sprites.CreateImage(image_path)
}

func (c *Circle) DrawShape(screen *ebiten.Image) {

	origWidth, origHeight := float64(c.Image.Bounds().Dx()), float64(c.Image.Bounds().Dy())

	// Calculate scale factors
	scaleX := c.Radius * 2 / origWidth
	scaleY := c.Radius * 2 / origHeight

	// Create draw options
	op := &ebiten.DrawImageOptions{}
	// Set scale
	op.GeoM.Scale(scaleX, scaleY)
	op.Filter = ebiten.FilterLinear
	// Set position (after scaling)
	op.GeoM.Translate(float64(c.Center.X-c.Radius), float64(c.Center.Y-c.Radius))
	screen.DrawImage(c.Image, op)
}
