package physics

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"math"
)

type PhysicsBody struct {
	Shape   Shape
	Dynamic dynamics.DynamicProperties
}

func (p *PhysicsBody) GetVelocity() linalg.Vector {
	return p.Dynamic.Velocity
}

func (p *PhysicsBody) SetVelocity(new_velocity linalg.Vector) {
	p.Dynamic.Velocity = new_velocity
}

func (p *PhysicsBody) GetMass() float64 {
	return p.Dynamic.Mass
}

func (p *PhysicsBody) UpdateKinematics(screenWidth, screenHeight int, timeDelta float64, gravity linalg.Vector) {
	// Apply velocity to position
	// Apply gravity
	p.Dynamic.Velocity.Y += (gravity.Y) * timeDelta
	newCenterX := p.Shape.GetCenter().X + p.Dynamic.Velocity.X*timeDelta
	newCenterY := p.Shape.GetCenter().Y + p.Dynamic.Velocity.Y*timeDelta + (0.5 * gravity.Y * math.Pow(timeDelta, 2))

	// Temporarily update center to check boundary collisions

	p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY})
	// Check and handle boundary collisions
	boundary := p.Shape.GetBoundaryPoints()
	collisionOccurred := false
	center := p.Shape.GetCenter()
	// Right boundary
	if boundary.MaxX >= float64(screenWidth) {
		overlap := boundary.MaxX - float64(screenWidth)
		p.Shape.SetCenter(Point{X: center.X - overlap, Y: center.Y})
		p.Dynamic.Velocity.X = -math.Abs(p.Dynamic.Velocity.X)
		collisionOccurred = true
	}

	// Left boundary
	if boundary.MinX <= 0 {
		p.Shape.SetCenter(Point{X: center.X - boundary.MinX, Y: center.Y})
		p.Dynamic.Velocity.X = math.Abs(p.Dynamic.Velocity.X)
		collisionOccurred = true
	}

	// Bottom boundary
	if boundary.MaxY >= float64(screenHeight) {
		overlap := boundary.MaxY - float64(screenHeight)

		p.Shape.SetCenter(Point{X: center.X, Y: center.Y - overlap})
		p.Dynamic.Velocity.Y = -math.Abs(p.Dynamic.Velocity.Y)
		collisionOccurred = true
	}

	// Top boundary
	if boundary.MinY <= 0 {
		p.Shape.SetCenter(Point{X: center.X, Y: center.Y - boundary.MinY})
		p.Dynamic.Velocity.Y = math.Abs(p.Dynamic.Velocity.Y)
		collisionOccurred = true
	}

	// If no collision occurred, keep the new position
	if !collisionOccurred {
		p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY})
	}
	p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY})

}
