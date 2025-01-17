package main

// GJK function detects if two shapes collide
func GJK(s1 Shape, s2 Shape) bool {
	//direction

	vector_centres := NewVector(s2.Center, s1.Center)
	vector_centres = vector_centres.Normalize()

	return true
}
