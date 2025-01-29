package physics

import linalg "HeadSoccer/math/helper"

type PhysicsWorld struct {
	Objects []Shape // Changed from []Shape to []*PhysicsObject
	Gravity linalg.Vector
}
