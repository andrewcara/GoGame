package shapes

import (
	linalg "HeadSoccer/math/helper"
	"HeadSoccer/math/helper/point"
)

type Point = point.Point

// Every Shape has a center and a furthest point given some directional vector
// We need the furthest point given a vector of a shape to Calculate GJK
type Shape interface {
	FurthestPoint(linalg.Vector) Point
	GetCenter() Point
	GetVelocity() linalg.Vector
	SetVelocity(linalg.Vector)
	GetMass() float64

	//Where we are passing in the dimensions of the screen into the update function

	//TO-DO on implementation
	UpdateKinematics(int, int, float64)
}
