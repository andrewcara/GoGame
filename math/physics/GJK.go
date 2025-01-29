package physics

import (
	linalg "HeadSoccer/math/helper"
	"HeadSoccer/math/helper/point"
	"HeadSoccer/shapes"
)

type Point = point.Point
type Shape = shapes.Shape

type Simplex struct {
	Values []Point
}

var Origin = Point{X: 0, Y: 0}

// Support computes the Minkowski difference of the furthest points
func Support(s1, s2 Shape, d linalg.Vector) Point {
	// Normalize the direction vector for consistent scaling
	d = d.Normalize()

	s1_furth_p := s1.FurthestPoint(d)
	s2_furth_p := s2.FurthestPoint(d.Scale(-1))

	return Point{X: s1_furth_p.X - s2_furth_p.X, Y: s1_furth_p.Y - s2_furth_p.Y}
}

func GJK(s1 Shape, s2 Shape) bool {
	simplex := Simplex{make([]Point, 0)}

	// Initial direction using centers
	direction_vector := linalg.NewVector(s1.GetCenter(), s2.GetCenter())
	if direction_vector.Magnitude() == 0 {
		// If centers overlap, use any non-zero vector
		direction_vector = linalg.Vector{X: 1, Y: 0}
	}

	// Get first support point
	support_point := Support(s1, s2, direction_vector)
	simplex.Values = append(simplex.Values, support_point)

	// Set new direction towards origin
	direction_vector = linalg.NewVector(simplex.Values[0], Origin)

	maxIterations := 20 // Prevent infinite loops
	iterations := 0

	for iterations < maxIterations {
		support_point = Support(s1, s2, direction_vector)

		// Check if we've made progress towards the origin
		if linalg.DotProduct(linalg.NewVector(Origin, support_point), direction_vector) < 0 {
			return false // No collision
		}

		simplex.Values = append(simplex.Values, support_point)

		if simplex.EvolveSimples(&direction_vector) {
			return true // Collision detected
		}

		iterations++
	}

	return false // No conclusive result after max iterations
}

func (s *Simplex) EvolveSimples(d *linalg.Vector) bool {
	switch len(s.Values) {
	case 2:
		return s.handleLine(d)
	case 3:
		return s.handleTriangle(d)
	}
	return false
}

func (s *Simplex) handleLine(d *linalg.Vector) bool {
	B := s.Values[0]
	A := s.Values[1]

	ab := linalg.NewVector(A, B)
	a0 := linalg.NewVector(A, Origin)

	// Compute perpendicular direction towards origin
	*d = linalg.TripleCrossProduct3D(ab, a0, ab)

	// If the perpendicular direction is zero, check if origin is on the line
	if d.Magnitude() < 1e-10 {
		// Check if origin is between A and B
		ba := linalg.NewVector(B, A)
		if linalg.DotProduct(linalg.NewVector(B, Origin), ba) > 0 &&
			linalg.DotProduct(a0, ab) > 0 {
			return true
		}
		*d = linalg.Vector{
			X: -ab.Y,
			Y: ab.X,
		}
	}

	return false
}

func (s *Simplex) handleTriangle(d *linalg.Vector) bool {
	A := s.Values[2] // Latest point added
	B := s.Values[1]
	C := s.Values[0]

	ab := linalg.NewVector(A, B)
	ac := linalg.NewVector(A, C)
	a0 := linalg.NewVector(A, Origin)

	// Compute perpendicular vectors
	abPerp := linalg.TripleCrossProduct3D(ac, ab, ab)
	acPerp := linalg.TripleCrossProduct3D(ab, ac, ac)

	if linalg.DotProduct(abPerp, a0) > 0 {
		// Origin is outside AB edge
		s.Values = append(s.Values[:0], s.Values[1:]...)
		*d = abPerp
		return false
	} else if linalg.DotProduct(acPerp, a0) > 0 {
		// Origin is outside AC edge
		s.Values = append(s.Values[:1], s.Values[2:]...)
		*d = acPerp
		return false
	}

	// Origin is inside the triangle
	return true
}
