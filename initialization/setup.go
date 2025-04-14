package initialization

import (
	headsoccer_constants "HeadSoccer/constants"
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"HeadSoccer/math/helper/point"

	"HeadSoccer/math/physics"
	"HeadSoccer/shapes"
)

func Setup(screenWidth, screenHeight float64, gravity linalg.Vector) physics.PhysicsWorld {
	crossbar1_center := linalg.Point{X: 25, Y: screenHeight - 77.5 - headsoccer_constants.GroundHeight}
	// Define vertices relative to center
	crossbar1_vertices := []linalg.Point{
		{X: 0, Y: screenHeight - 75},  // top left
		{X: 50, Y: screenHeight - 75}, // top right
		{X: 50, Y: screenHeight - 80}, // bottom left
		{X: 0, Y: screenHeight - 80},  // bottom right
	}
	var crossbar1 shapes.Polygon
	crossbar1.Initialize(crossbar1_center, crossbar1_vertices)
	crossbar1.SetImage("left_net.png")
	crossbar1.SetImageDimensions(80, 50, &point.Point{X: screenWidth - (50 / 2), Y: screenHeight - (80 / 2) - headsoccer_constants.GroundHeight})
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

	crossbar2_center := linalg.Point{X: screenWidth - 25.5, Y: screenHeight - 77.5 - headsoccer_constants.GroundHeight}
	// Define vertices relative to center
	crossbar2_vertices := []linalg.Point{
		{X: screenWidth - 1, Y: screenHeight - 75},  // top left
		{X: screenWidth - 51, Y: screenHeight - 75}, // top right
		{X: screenWidth - 51, Y: screenHeight - 80}, // bottom left
		{X: screenWidth - 1, Y: screenHeight - 80},  // bottom right
	}
	var crossbar2 shapes.Polygon
	crossbar2.Initialize(crossbar2_center, crossbar2_vertices)
	crossbar2.SetImage("right_net.png")
	crossbar2.SetImageDimensions(80, 50, &point.Point{X: (50 / 2), Y: screenHeight - (80 / 2) - headsoccer_constants.GroundHeight})

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
			Center: linalg.Point{X: headsoccer_constants.SoccerBallStartingX, Y: headsoccer_constants.SoccerBallStartingY},
			Radius: headsoccer_constants.BallRadius,
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
	player1Center := linalg.Point{X: headsoccer_constants.Player1StartingX, Y: headsoccer_constants.Player1StartingY}
	// Define vertices relative to center
	var vertices1 = []linalg.Point{
		{X: headsoccer_constants.Player1StartingX - 10, Y: headsoccer_constants.Player1StartingY - 10}, // top left
		{X: headsoccer_constants.Player1StartingX - 10, Y: headsoccer_constants.Player1StartingY + 10}, // top right
		{X: headsoccer_constants.Player1StartingX + 10, Y: headsoccer_constants.Player1StartingY + 10}, // bottom left
		{X: headsoccer_constants.Player1StartingX + 10, Y: headsoccer_constants.Player1StartingY - 10}, // bottom right
	}
	var polygon1 shapes.Polygon
	polygon1.Initialize(player1Center, vertices1)
	polygon1.SetImage("messi.png")
	polygon1.SetImageDimensions(20, 20, &polygon1.Center)

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

	player2Center := linalg.Point{X: headsoccer_constants.Player2StartingX, Y: headsoccer_constants.Player2StartingY}
	// Define vertices relative to center
	var vertices2 = []linalg.Point{
		{X: headsoccer_constants.Player2StartingX - 10, Y: headsoccer_constants.Player2StartingY - 10}, // top left
		{X: headsoccer_constants.Player2StartingX - 10, Y: headsoccer_constants.Player2StartingY + 10}, // top right
		{X: headsoccer_constants.Player2StartingX + 10, Y: headsoccer_constants.Player2StartingY + 10}, // bottom left
		{X: headsoccer_constants.Player2StartingX + 10, Y: headsoccer_constants.Player2StartingY - 10}, // bottom right
	}
	var polygon2 shapes.Polygon
	polygon2.Initialize(player2Center, vertices2)
	polygon2.SetImage("ronaldo.png")
	polygon2.SetImageDimensions(20, 20, &polygon2.Center)

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
