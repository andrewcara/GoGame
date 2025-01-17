package main

import "math"

// Point represents a 2D point with x and y coordinates

// Vector represents a 2D vector defined by two points
type Vector struct {
	Start Point
	End   Point
}

// NewVector creates a new vector from two points
func NewVector(start, end Point) Vector {
	return Vector{
		Start: start,
		End:   end,
	}
}

// ToComponents converts a vector to its x and y components
func (v Vector) ToComponents() (float64, float64) {
	return v.End.x - v.Start.x, v.End.y - v.Start.y
}

// DotProduct calculates the dot product of two vectors
func DotProduct(v1, v2 Vector) float64 {
	// Convert vectors to component form
	x1, y1 := v1.ToComponents()
	x2, y2 := v2.ToComponents()

	// Dot product = x1*x2 + y1*y2
	return x1*x2 + y1*y2
}

// CrossProduct calculates the magnitude of the cross product of two 2D vectors
// Note: In 2D, cross product gives a scalar value representing the area of the parallelogram
// formed by the two vectors
func CrossProduct(v1, v2 Vector) float64 {
	// Convert vectors to component form
	x1, y1 := v1.ToComponents()
	x2, y2 := v2.ToComponents()

	// Cross product in 2D = x1*y2 - y1*x2
	return x1*y2 - y1*x2
}

// Magnitude calculates the length of the vector
func (v Vector) Magnitude() float64 {
	x, y := v.ToComponents()
	return math.Sqrt(x*x + y*y)
}

func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	if mag == 0 {
		return v // Return original vector if magnitude is 0
	}

	x, y := v.ToComponents()
	return Vector{
		Start: v.Start,
		End: Point{
			x: v.Start.x + (x / mag),
			y: v.Start.y + (y / mag),
		},
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
