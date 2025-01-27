package physics

import (
	linalg "HeadSoccer/math/helper"
)

type CollsionHandler struct {
	Collider PhysicsObject
}

//We Need to detect collisions with the ground, the walls and the top of the screen
//We also need to detect collision between players and the ball and both players

func (c *CollsionHandler) HitsOtherObject(other_object *PhysicsObject) bool {
	//if objects collide
	//calculate normal to collision and alter the shapes velocity
	if GJK(c.Collider.Shape, other_object.Shape) {
		collision_normal := linalg.NewVector(c.Collider.Shape.GetCenter(), other_object.Shape.GetCenter())
		collision_normal = collision_normal.Normalize()

		v1_init := c.Collider.Shape.GetVelocity()
		v2_init := other_object.Shape.GetVelocity()

		v1Normal := collision_normal.Scale(linalg.DotProduct(collision_normal, (v1_init)))
		v2Normal := collision_normal.Scale(linalg.DotProduct(collision_normal, (v2_init)))

		v1Tangential := v1_init.Subtract(v1Normal)
		v2Tangential := v2_init.Subtract(v2Normal)

		// Elastic collision formula for 1D along the collision normal
		mass1 := c.Collider.Shape.GetMass()
		mass2 := other_object.Shape.GetMass()

		newV1Normal := v1Normal.Scale((mass1 - mass2) / (mass1 + mass2)).Add(v2Normal.Scale(2 * mass2 / (mass1 + mass2)))
		newV2Normal := v2Normal.Scale((mass2 - mass1) / (mass1 + mass2)).Add(v1Normal.Scale(2 * mass1 / (mass1 + mass2)))

		// Update velocities
		c.Collider.Shape.SetVelocity(v1Tangential.Add(newV1Normal))
		other_object.Shape.SetVelocity(v2Tangential.Add(newV2Normal))

		return true
	}
	return false

}
