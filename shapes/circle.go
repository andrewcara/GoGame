package shapes

import (
	linalg "HeadSoccer/math/helper"
)

type Circle struct {
	Shape
	Center   Point
	Radius   float64
	Velocity linalg.Vector
}

func (c *Circle) FurthestPoint(direction_vector linalg.Vector) Point {
	normalized_vec := direction_vector.Normalize()
	scaled_vec := normalized_vec.Scale(c.Radius)
	return Point{X: c.Center.X + scaled_vec.X, Y: c.Center.Y + scaled_vec.Y}
}

func (c *Circle) GetCenter() Point {
	return c.Center
}

func (c *Circle) UpdateKinematics(screenWidth, screenHeight int, timeDelta float64) {

	c.Center.X += c.Velocity.X * timeDelta
	c.Center.Y += c.Velocity.Y * timeDelta
	maxX := float64(screenWidth) - c.Radius
	maxY := float64(screenHeight) - c.Radius

	if (c.Center.X) >= float64(screenWidth)-(c.Radius) || c.Center.X <= c.Radius {
		if c.Center.X > maxX {
			c.Center.X = maxX
		} else if c.Center.X < c.Radius {
			c.Center.X = c.Radius
		}
		c.Velocity.X *= -1
	}
	if (c.Center.Y) >= float64(screenWidth)-(c.Radius) || c.Center.Y <= c.Radius {
		if c.Center.Y > maxY {
			c.Center.Y = maxY
		} else if c.Center.Y < c.Radius {
			c.Center.Y = c.Radius
		}
		c.Velocity.Y *= -1
	}

}
