package shapes

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	coefficient_friction  = 0.8
	collision_restitution = 0.8
)

type Circle struct {
	Center  Point
	Radius  float64
	Dynamic dynamics.DynamicProperties
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

func (c *Circle) GetVelocity() linalg.Vector {
	return c.Dynamic.Velocity
}

func (c *Circle) SetVelocity(new_velocity linalg.Vector) {
	c.Dynamic.Velocity = new_velocity
}

func (c *Circle) GetMass() float64 {
	return c.Dynamic.Mass
}
func (c *Circle) GetSurfacePoint(direction_vector linalg.Vector) Point {
	normalized_vec := direction_vector.Normalize()
	scaled_vec := normalized_vec.Scale(c.Radius)
	return Point{X: c.Center.X + scaled_vec.X, Y: c.Center.Y + scaled_vec.Y}
}

func (c *Circle) UpdateKinematics(screenWidth, screenHeight int, timeDelta float64, gravity linalg.Vector) {

	//Adjust velocity based on gravity
	c.Dynamic.Velocity.Y += (gravity.Y) * timeDelta

	c.Center.X += c.Dynamic.Velocity.X * timeDelta
	c.Center.Y += c.Dynamic.Velocity.Y * timeDelta
	maxX := float64(screenWidth) - c.Radius
	maxY := float64(screenHeight) - c.Radius

	//collision with the sides
	if (c.Center.X) >= float64(screenWidth)-(c.Radius) || c.Center.X <= c.Radius {
		if c.Center.X > maxX {
			c.Center.X = maxX
		} else if c.Center.X < c.Radius {
			c.Center.X = c.Radius
		}
		c.Dynamic.Velocity.X *= -1 * collision_restitution
	}

	//collisions with the ground here maxY is actuall the bottom of the screen
	if (c.Center.Y) >= float64(screenWidth)-(c.Radius) || c.Center.Y <= c.Radius {

		if c.Center.Y > maxY {
			c.Center.Y = maxY
		} else if c.Center.Y <= c.Radius {
			c.Center.Y = c.Radius
		}
		c.Dynamic.Velocity.Y *= -1 * collision_restitution
	}

}

func (c *Circle) DrawShape(screen *ebiten.Image, color color.RGBA) {
	vector.DrawFilledCircle(screen, float32(c.Center.X), float32(c.Center.Y), float32(c.Radius), color, false)

}
