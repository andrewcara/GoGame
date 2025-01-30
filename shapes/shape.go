package shapes

import (
	linalg "HeadSoccer/math/helper"
	"HeadSoccer/math/helper/point"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point = point.Point

// Every Shape has a center and a furthest point given some directional vector
// We need the furthest point given a vector of a shape to Calculate GJK
type Shape interface {
	FurthestPoint(linalg.Vector) Point
	GetCenter() Point
	SetCenter(Point)
	GetVelocity() linalg.Vector
	SetVelocity(linalg.Vector)
	GetMass() float64
	GetSurfacePoint(direction_vector linalg.Vector) Point
	DrawShape(*ebiten.Image, color.RGBA)
	//Where we are passing in the dimensions of the screen into the update function
	UpdateKinematics(int, int, float64, linalg.Vector)
}
