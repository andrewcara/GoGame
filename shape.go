package main

type Point struct {
	X, Y float64
}

func (p *Point) Add(p2 Point) Point {
	return (Point{X: p.X + p2.X, Y: p.Y + p2.Y})
}
func (p *Point) Subtract(p2 Point) Point {
	return (Point{X: p.X - p2.X, Y: p.Y - p2.Y})
}
func (p *Point) Multiply(p2 Point) Point {
	return (Point{X: p.X * p2.X, Y: p.Y * p2.Y})
}

// Every Shape has a center and a furthest point given some directional vector
type Shape interface {
	FurthestPoint(Vector) Point
	GetCenter() Point
	//Where we are passing in the dimensions of the screen into the update function
	UpdateKinematics(int, int, float64)
}

type Polygon struct {
	Shape
	Center   Point
	Vertices []Point
	VelX     float64
	VelY     float64
}

// Find the Point that is the furthest from the Center given a directional vector
func (p *Polygon) FurthestPoint(direction_vector Vector) Point {
	maximum_dot_product := DotProduct(NewVector(p.Center, p.Vertices[0]), direction_vector)
	maximum_point := p.Vertices[0]
	//Iterate through every point in the list of Points to Find the furthest
	//The furthest point will be the maximum dot product

	for _, point := range p.Vertices[1:len(p.Vertices)] {
		local_dot_proudct := DotProduct(direction_vector, NewVector(p.Center, point))

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

type Circle struct {
	Shape
	Center Point
	Radius float64
	VelX   float64
	VelY   float64
}

func (c *Circle) FurthestPoint(direction_vector Vector) Point {
	normalized_vec := direction_vector.Normalize()
	scaled_vec := normalized_vec.Scale(c.Radius)
	return Point{X: c.Center.X + scaled_vec.X, Y: c.Center.Y + scaled_vec.Y}
}

func (c *Circle) GetCenter() Point {
	return c.Center
}

func (c *Circle) UpdateKinematics(screenWidth, screenHeight int, timeDelta float64) {

	c.Center.X += c.VelX * timeDelta
	c.Center.Y += c.VelY * timeDelta
	maxX := float64(screenWidth) - c.Radius
	maxY := float64(screenHeight) - c.Radius

	if (c.Center.X) >= float64(screenWidth)-(c.Radius) || c.Center.X <= c.Radius {
		if c.Center.X > maxX {
			c.Center.X = maxX
		} else if c.Center.X < c.Radius {
			c.Center.X = c.Radius
		}
		c.VelX *= -1
	}
	if (c.Center.Y) >= float64(screenWidth)-(c.Radius) || c.Center.Y <= c.Radius {
		if c.Center.Y > maxY {
			c.Center.Y = maxY
		} else if c.Center.Y < c.Radius {
			c.Center.Y = c.Radius
		}
		c.VelY *= -1
	}

}
