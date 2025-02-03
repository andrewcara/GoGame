package physics

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"math"
)

type PhysicsBody struct {
	Shape               Shape
	Dynamic             dynamics.DynamicProperties
	CoefficientFriction float64
	Restitution         float64
	IsStatic            bool
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
	// Air resistance coefficients
	const airResistanceCoeffX = 0.1
	const airResistanceCoeffY = 0.1

	// Friction coefficient (ground friction)

	// Store original center for restoration if needed
	originalCenter := p.Shape.GetCenter()

	// Apply air resistance (proportional to velocity)
	airResistanceX := -airResistanceCoeffX * p.Dynamic.Velocity.X
	airResistanceY := -airResistanceCoeffY * p.Dynamic.Velocity.Y

	// Update velocity due to gravity and air resistance
	newVelocityY := p.Dynamic.Velocity.Y + (gravity.Y+airResistanceY)*timeDelta
	newVelocityX := p.Dynamic.Velocity.X + airResistanceX*timeDelta

	// Apply average velocity to position
	averageVelocityY := (p.Dynamic.Velocity.Y + newVelocityY) / 2
	averageVelocityX := (p.Dynamic.Velocity.X + newVelocityX) / 2

	newCenterX := originalCenter.X + averageVelocityX*timeDelta
	newCenterY := originalCenter.Y + averageVelocityY*timeDelta

	// Update velocity
	p.Dynamic.Velocity.Y = newVelocityY
	p.Dynamic.Velocity.X = newVelocityX

	// Update position
	p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY})

	// Check boundaries
	boundary := p.Shape.GetBoundaryPoints()
	collisionOccurred := false

	// Bottom boundary (ground) - apply friction
	if boundary.MaxY >= float64(screenHeight) {
		overlap := boundary.MaxY - float64(screenHeight)
		p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY - overlap})

		// Apply friction when on ground
		p.Dynamic.Velocity.X *= math.Pow(1-p.CoefficientFriction, timeDelta)
		p.Dynamic.Velocity.Y = -math.Abs(p.Dynamic.Velocity.Y) * p.Restitution
		collisionOccurred = true
	}

	// Top boundary
	if boundary.MinY <= 0 {
		p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY - boundary.MinY})
		p.Dynamic.Velocity.Y = math.Abs(p.Dynamic.Velocity.Y) * p.Restitution // Bounce with energy loss
		collisionOccurred = true
	}

	// Right boundary
	if boundary.MaxX >= float64(screenWidth) {
		overlap := boundary.MaxX - float64(screenWidth)
		p.Shape.SetCenter(Point{X: newCenterX - overlap, Y: newCenterY})
		p.Dynamic.Velocity.X *= -p.Restitution // Bounce with energy loss
		collisionOccurred = true
	}

	// Left boundary
	if boundary.MinX <= 0 {
		p.Shape.SetCenter(Point{X: newCenterX - boundary.MinX, Y: newCenterY})
		p.Dynamic.Velocity.X *= -p.Restitution // Bounce with energy loss
		collisionOccurred = true
	}

	// If no collision occurred, keep the new position
	if !collisionOccurred {
		p.Shape.SetCenter(Point{X: newCenterX, Y: newCenterY})
	}
}
