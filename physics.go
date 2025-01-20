package main

import "fmt"

type Status int

const (
	NO_INTERSCTION Status = iota + 1
	FOUND
	EVOLVING
)

type Simplex struct {
	Values []Point
	Status
}

// GJK function detects if two shapes collide
var Origin = Point{0, 0}

func GJK(s1 Shape, s2 Shape) bool {

	simplex := Simplex{make([]Point, 0), EVOLVING}
	//direction

	for simplex.Status == EVOLVING {
		fmt.Println(simplex.Values)
		simplex.EvolveSimples(s1, s2)
	}

	if simplex.Status == FOUND {
		return true
	}
	return false
}

func (s *Simplex) EvolveSimples(s1 Shape, s2 Shape) Status {
	var direction_vector Vector
	var support_point Point
	switch len(s.Values) {
	case 0:

		direction_vector = NewVector(s1.GetCenter(), s2.GetCenter())
		//fmt.Println(direction_vector)
		support_point = Support(s1, s2, direction_vector)
		//fmt.Println(support_point, direction_vector)
		s.Values = append(s.Values, support_point)
	case 1:
		direction_vector = NewVector(s2.GetCenter(), s1.GetCenter())
		support_point = Support(s1, s2, direction_vector)
		s.Values = append(s.Values, support_point)
	case 2:
		C := s.Values[0]
		B := s.Values[1]

		cb := NewVector(B, C)
		c0 := NewVector(B, Origin)

		direction_vector = TripleCrossProduct3D(cb, c0, cb)
		support_point = Support(s1, s2, direction_vector)
		s.Values = append(s.Values, support_point)

	case 3:
		A := s.Values[2]
		B := s.Values[1]
		C := s.Values[0]

		a0 := NewVector(A, Origin)
		ab := NewVector(A, B)
		ac := NewVector(A, C)

		//fmt.Print(C, B, A)

		abPerp := TripleCrossProduct3D(ac, ab, ab)
		acPerp := TripleCrossProduct3D(ab, ac, ac)
		//fmt.Println(abPerp, acPerp)
		fmt.Println(DotProduct(abPerp, a0), DotProduct(acPerp, a0))

		if DotProduct(abPerp, a0) > 0 {
			// the origin is outside line ab
			// get rid of c and add a new support in the direction of abPerp

			//Essentially we are removing the first index, and appending new value at end of value
			s.Values = append(s.Values[:0], s.Values[1:]...)
			direction_vector = abPerp
			support_point = Support(s1, s2, direction_vector)
			s.Values = append(s.Values, support_point)

		} else if DotProduct(acPerp, a0) > 0 {
			// the origin is outside line ac
			s.Values = append(s.Values[:1], s.Values[2:]...)
			direction_vector = abPerp
			support_point = Support(s1, s2, direction_vector)
			s.Values = append(s.Values, support_point)
		} else {
			// the origin is inside both ab and ac,
			// so it must be inside the triangle!
			s.Status = FOUND
			return s.Status
		}

	}

	fmt.Println(direction_vector, support_point)
	return PassesOrigin(NewVector(Origin, support_point), direction_vector)
}

func Support(s1, s2 Shape, d Vector) Point {
	s1_furth_p := s1.FurthestPoint(d)
	s2_furth_p := s2.FurthestPoint(d.Scale(-1))
	//fmt.Println(s1_furth_p, s2_furth_p)

	return Point{X: s1_furth_p.X - s2_furth_p.X, Y: s1_furth_p.Y - s2_furth_p.Y}
}

func PassesOrigin(v1, v2 Vector) Status {
	if DotProduct(v1, v2) >= 0 {
		return EVOLVING
	}
	return NO_INTERSCTION
}
