package physics

import linalg "HeadSoccer/math/helper"

// Values associated with any Physics Objects that govern how they are treated in Main
type DynamicProperties struct {
	Position linalg.Vector
	Velocity linalg.Vector
	Force    linalg.Vector
	Mass     float64
}
