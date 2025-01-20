package main

type Status int

type Simplex struct {
	Values []Point
}

// GJK function detects if two shapes collide
var Origin = Point{0, 0}

func GJK(s1 Shape, s2 Shape) bool {

	simplex := Simplex{make([]Point, 0)}
	//direction
	direction_vector := NewVector(s1.GetCenter(), s2.GetCenter())
	support_point := Support(s1, s2, direction_vector)
	simplex.Values = append(simplex.Values, support_point)
	direction_vector = NewVector(simplex.Values[0], Origin)

	for {
		support_point = Support(s1, s2, direction_vector)
		result := simplex.EvolveSimples(s1, s2, &direction_vector)
		if DotProduct(NewVector(Origin, support_point), direction_vector) < 0 {
			return false
		}
		simplex.Values = append(simplex.Values, support_point)

		if result {
			return true
		}
	}
}

func (s *Simplex) EvolveSimples(s1 Shape, s2 Shape, d *Vector) bool {

	switch len(s.Values) {
	case 2:
		B := s.Values[0]
		A := s.Values[1]

		ab := NewVector(A, B)
		a0 := NewVector(A, Origin)
		*d = TripleCrossProduct3D(ab, a0, ab)
		return false

	case 3:
		A := s.Values[2] // Latest point added
		B := s.Values[1]
		C := s.Values[0]

		a0 := NewVector(A, Origin) // Vector from A to origin
		ab := NewVector(A, B)
		ac := NewVector(A, C)

		abPerp := TripleCrossProduct3D(ac, ab, ab)
		acPerp := TripleCrossProduct3D(ab, ac, ac)

		if DotProduct(abPerp, a0) > 0 {
			s.Values = append(s.Values[:0], s.Values[1:]...)
			*d = abPerp
			return false

		} else if DotProduct(acPerp, a0) > 0 {
			s.Values = append(s.Values[:1], s.Values[2:]...)
			*d = acPerp
			return false

		} else {
			// Origin is inside the triangle
			return true
		}
	}
	return false
}

func Support(s1, s2 Shape, d Vector) Point {
	s1_furth_p := s1.FurthestPoint(d)
	s2_furth_p := s2.FurthestPoint(d.Scale(-1))
	//fmt.Println(s1_furth_p, s2_furth_p)

	return Point{X: s1_furth_p.X - s2_furth_p.X, Y: s1_furth_p.Y - s2_furth_p.Y}
}
