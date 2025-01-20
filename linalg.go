package main

import (
	"math"
)

// Vector represents a 2D vector with x and y components
type Vector struct {
	X, Y float64
}

type Vec3 struct {
	X, Y, Z float64
}

// NewVector creates a new vector with given x and y components
func NewVector(start, end Point) Vector {
	return Vector{
		X: end.X - start.X,
		Y: end.Y - start.Y,
	}
}

// Add returns the sum of two vectors
func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Subtract returns the difference between two vectors
func (v Vector) Subtract(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Scale multiplies the vector by a scalar value
func (v Vector) Scale(scalar float64) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

// DotProduct calculates the dot product of two vectors
func DotProduct(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// CrossProduct calculates the magnitude of the cross product of two 2D vectors
// Note: In 2D, cross product gives a scalar value representing the area of the parallelogram

// Magnitude calculates the length of the vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a unit vector in the same direction
func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	if mag == 0 {
		return v // Return original vector if magnitude is 0
	}
	return Vector{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

// Angle calculates the angle between two vectors in radians
func Angle(v1, v2 Vector) float64 {
	dot := DotProduct(v1, v2)
	mags := v1.Magnitude() * v2.Magnitude()

	// Handle potential floating-point errors
	if mags == 0 {
		return 0
	}

	cos := dot / mags
	// Ensure cos is within valid range for acos
	if cos > 1 {
		cos = 1
	} else if cos < -1 {
		cos = -1
	}

	return math.Acos(cos)
}

func CrossProduct2D(u, v Vector) float64 {
	return (u.X * v.Y) - (u.Y * v.X)
}
func CrossProduct3D(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: (a.Z*b.X - a.X*b.Z),
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// TripleCrossProduct2D calculates (a × b) × a for two 2D vectors
func TripleCrossProduct3D(a, b, c Vector) Vector {
	// Convert 2D vectors to 3D (set z=0).
	A := Vec3{a.X, a.Y, 0}
	B := Vec3{b.X, b.Y, 0}
	C := Vec3{c.X, c.Y, 0}

	// First cross product: A × B
	first := CrossProduct3D(A, B)

	// Second cross product: (A × B) × C
	second := CrossProduct3D(first, C)

	// Return the x and y components as a Vec2
	return Vector{second.X, second.Y}
}

// Additional useful methods

// Perpendicular returns a vector rotated 90 degrees counterclockwise
func (v Vector) Perpendicular() Vector {
	return Vector{
		X: -v.Y,
		Y: v.X,
	}
}

// Project returns the projection of this vector onto another vector
func (v Vector) Project(onto Vector) Vector {
	dot := DotProduct(v, onto)
	mag2 := DotProduct(onto, onto)
	if mag2 == 0 {
		return Vector{0, 0}
	}
	scale := dot / mag2
	return onto.Scale(scale)
}

// Reflect returns the reflection of this vector across another vector
func (v Vector) Reflect(across Vector) Vector {
	proj := v.Project(across)
	return proj.Scale(2).Subtract(v)
}
