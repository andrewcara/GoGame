package initialization

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
)

func Setup(screenWidth, screenHeight float64, gravity linalg.Vector) physics.PhysicsWorld {
	crossbar1_center := linalg.Point{X: 25, Y: screenHeight - 77.5}
	// Define vertices relative to center
	crossbar1_vertices := []linalg.Point{
		{X: 0, Y: screenHeight - 75},  // top left
		{X: 50, Y: screenHeight - 75}, // top right
		{X: 50, Y: screenHeight - 80}, // bottom left
		{X: 0, Y: screenHeight - 80},  // bottom right
	}
	var crossbar1 shapes.Polygon
	crossbar1.Initialize(crossbar1_center, crossbar1_vertices)

	crossbarBody := physics.PhysicsBody{
		Shape: &crossbar1,
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         2500,
			Acceleration: gravity,
		},
		CoefficientFriction: 0.5,
		Restitution:         1,
		IsStatic:            true,
	}

	crossbar2_center := linalg.Point{X: screenWidth - 25.5, Y: screenHeight - 77.5}
	// Define vertices relative to center
	crossbar2_vertices := []linalg.Point{
		{X: screenWidth - 1, Y: screenHeight - 75},  // top left
		{X: screenWidth - 51, Y: screenHeight - 75}, // top right
		{X: screenWidth - 51, Y: screenHeight - 80}, // bottom left
		{X: screenWidth - 1, Y: screenHeight - 80},  // bottom right
	}
	var crossbar2 shapes.Polygon
	crossbar2.Initialize(crossbar2_center, crossbar2_vertices)

	crossbarBody2 := physics.PhysicsBody{
		Shape: &crossbar2,
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         2500,
			Acceleration: gravity,
		},
		CoefficientFriction: 0.5,
		Restitution:         1,
		IsStatic:            true,
	}

	// Ball initialization - centered more visibly
	ball := physics.PhysicsBody{
		Shape: &shapes.Circle{
			Center: linalg.Point{X: screenWidth / 2, Y: screenHeight / 2},
			Radius: 15,
		},
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         5.0,
			Acceleration: gravity,
		},
		CoefficientFriction: 0.4,
		Restitution:         0.8,
		IsStatic:            false,
	}

	ball.Shape.SetImage("soccer_ball.png")

	// Player 1 - Left side of screen
	player1Center := linalg.Point{X: 400, Y: 400}
	// Define vertices relative to center
	vertices1 := []linalg.Point{
		{X: 390, Y: 390}, // top left
		{X: 390, Y: 410}, // top right
		{X: 410, Y: 410}, // bottom left
		{X: 410, Y: 390}, // bottom right
	}
	var polygon1 shapes.Polygon
	polygon1.Initialize(player1Center, vertices1)
	polygon1.SetImage("messi.png")

	player1Body := physics.PhysicsBody{
		Shape: &polygon1,
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         25.0,
			Acceleration: gravity,
		},
		CoefficientFriction: 0.8,
		Restitution:         0.1,
		IsStatic:            false,
	}

	player2Center := linalg.Point{X: 120, Y: 120}
	// Define vertices relative to center
	vertices2 := []linalg.Point{
		{X: 110, Y: 110}, // top left
		{X: 110, Y: 130}, // top right
		{X: 130, Y: 130}, // bottom left
		{X: 130, Y: 110}, // bottom right
	}
	var polygon2 shapes.Polygon
	polygon2.Initialize(player2Center, vertices2)

	player2Body := physics.PhysicsBody{
		Shape: &polygon2,
		Dynamic: dynamics.DynamicProperties{
			Velocity:     linalg.Vector{X: 0, Y: 0},
			Force:        linalg.Vector{X: 0, Y: -9.8},
			Mass:         25.0,
			Acceleration: gravity,
		},
		CoefficientFriction: 0.8,
		Restitution:         0.1,
		IsStatic:            false,
	}

	// Initialize physics world
	world := physics.PhysicsWorld{
		Objects: make([]*physics.PhysicsBody, 0),
		Gravity: gravity,
	}

	// Add Physics Objects to World
	world.Objects = append(world.Objects,
		&player1Body,
		&player2Body,
		&ball,
		&crossbarBody,
		&crossbarBody2,
	)

	return world
}
