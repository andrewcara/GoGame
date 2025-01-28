package physics

import (
	linalg "HeadSoccer/math/helper"
	"HeadSoccer/shapes"
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

		//Since there is a time delta between each update their may be a case where the collision is detected with the shapes overlapping
		//We need to offset the distances between whatever overlap does occur so that we can properlt draw the ball

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

		c.SetNewDistances(other_object)
		return true
	}
	return false

}

func (c *CollsionHandler) SetNewDistances(other_object *PhysicsObject) {
	collision_normal := linalg.NewVector(c.Collider.Shape.GetCenter(), other_object.Shape.GetCenter())
	//Interface methods used

	//TO-DO Implement Distance calculation for polygons as well by making this function universal

	ballRadiusVector := linalg.NewVector(c.Collider.Shape.FurthestPoint(collision_normal), c.Collider.Shape.GetCenter())
	ballRadius := ballRadiusVector.Magnitude()

	ball2RadiusVector := linalg.NewVector(other_object.Shape.FurthestPoint(collision_normal), other_object.Shape.GetCenter())
	ball2Radius := ball2RadiusVector.Magnitude()

	ball := shapes.Circle{Center: c.Collider.Shape.GetCenter(), Radius: ballRadius}
	ball2 := shapes.Circle{Center: other_object.Shape.GetCenter(), Radius: ball2Radius}

	dx := ball2.Center.X - ball.Center.X
	dy := ball2.Center.Y - ball.Center.Y

	dist := collision_normal.Magnitude()
	overlap := (ball.Radius + ball2.Radius - dist) / 2
	moveX := (dx / dist) * overlap
	moveY := (dy / dist) * overlap

	c.Collider.Shape.SetCenter(Point{X: ball.Center.X - moveX, Y: ball.Center.Y - moveY})
	other_object.Shape.SetCenter(Point{X: ball2.Center.X + moveX, Y: ball2.Center.Y + moveY})
}
