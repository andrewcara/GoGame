package physics

import linalg "HeadSoccer/math/helper"

type PhysicsWorld struct {
	Objects []*PhysicsBody // Changed from []Shape to []*PhysicsObject
	Gravity linalg.Vector
}
