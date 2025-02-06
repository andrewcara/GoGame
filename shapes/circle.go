package shapes

import (
	linalg "HeadSoccer/math/helper"
)

const (
	coefficient_friction  = 0.8
	collision_restitution = 0.8
)

type Circle struct {
	Center Point
	Radius float64
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
