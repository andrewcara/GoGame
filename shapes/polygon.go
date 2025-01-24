package shapes

import (
	linalg "HeadSoccer/math/helper"
)

type Polygon struct {
	Shape
	Center   Point
	Vertices []Point
	//Velocity to be implemented as a vector
	VelX float64
	VelY float64
}

// Find the Point that is the furthest from the Center given a directional vector
func (p *Polygon) FurthestPoint(direction_vector linalg.Vector) Point {

	maximum_dot_product := linalg.DotProduct(linalg.NewVector(p.Center, p.Vertices[0]), direction_vector)
	maximum_point := p.Vertices[0]
	//Iterate through every point in the list of Points to Find the furthest
	//The furthest point will be the maximum dot product

	for _, point := range p.Vertices[1:len(p.Vertices)] {
		local_dot_proudct := linalg.DotProduct(direction_vector, linalg.NewVector(p.Center, point))

		if local_dot_proudct > maximum_dot_product {
			maximum_point = point
			maximum_dot_product = local_dot_proudct

		}
	}
	return maximum_point
}

func (p *Polygon) GetCenter() Point {
	return p.Center
}
