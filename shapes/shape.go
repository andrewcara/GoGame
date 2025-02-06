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
	SetCenter(Point)
	GetSurfacePoint(direction_vector linalg.Vector) Point
	GetBoundaryPoints() BoundaryPoints
	//Where we are passing in the dimensions of the screen into the update function
}
