package main

type Point struct {
	x, y float64
}

type Shape struct {
	Center    Point
	Perimeter []Point
}

func (s *Shape) FurthestPoint(direction_vector Vector) Point {
	maximum_dot_product := DotProduct(NewVector(s.Center, s.Perimeter[0]), direction_vector)
	maximum_point := s.Perimeter[0]

	for _, point := range s.Perimeter[1:len(s.Perimeter)] {
		local_dot_proudct := DotProduct(direction_vector, NewVector(s.Center, s.Perimeter[0]))

		if local_dot_proudct > maximum_dot_product {
			maximum_point = point
			maximum_dot_product = local_dot_proudct

		}
	}
	return maximum_point
}
