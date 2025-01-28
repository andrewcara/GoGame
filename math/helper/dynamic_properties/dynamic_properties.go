package dynamics

import linalg "HeadSoccer/math/helper"

// Values associated with any Physics Objects that govern how they are treated in Main
type DynamicProperties struct {
	Velocity   linalg.Vector
	Force      linalg.Vector
	Mass       float64
	Accelation linalg.Vector
}
