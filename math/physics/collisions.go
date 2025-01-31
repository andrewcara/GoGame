package physics

import (
	linalg "HeadSoccer/math/helper"
)

//We Need to detect collisions with the ground, the walls and the top of the screen
//We also need to detect collision between players and the ball and both players

func CollisionOccurs(object1, object2 *PhysicsBody) bool {
	//if objects collide
	//calculate normal to collision and alter the shapes velocity
	if GJK(*&object1.Shape, *&object2.Shape) {

		//Since there is a time delta between each update their may be a case where the collision is detected with the shapes overlapping
		//We need to offset the distances between whatever overlap does occur so that we can properlt draw the ball

		SetNewDistances(&object1.Shape, &object2.Shape)
		collision_normal := linalg.NewVector((*object1).Shape.GetCenter(), (*object2).Shape.GetCenter())
		collision_normal = collision_normal.Normalize()

		v1_init := (*object1).GetVelocity()
		v2_init := (*object2).GetVelocity()

		v1Normal := collision_normal.Scale(linalg.DotProduct(collision_normal, (v1_init)))
		v2Normal := collision_normal.Scale(linalg.DotProduct(collision_normal, (v2_init)))

		v1Tangential := v1_init.Subtract(v1Normal)
		v2Tangential := v2_init.Subtract(v2Normal)

		// Elastic collision formula for 1D along the collision normal
		mass1 := (*object1).GetMass()
		mass2 := (*object2).GetMass()

		//Can add a coefficient of restitution here if we want to make the collision inelestic
		newV1Normal := v1Normal.Scale((mass1 - mass2) / (mass1 + mass2)).Add(v2Normal.Scale(2 * mass2 / (mass1 + mass2)))
		newV2Normal := v2Normal.Scale((mass2 - mass1) / (mass1 + mass2)).Add(v1Normal.Scale(2 * mass1 / (mass1 + mass2)))

		// Update velocities

		(*object1).SetVelocity((v1Tangential.Add(newV1Normal)))
		(*object2).SetVelocity(v2Tangential.Add(newV2Normal))

		return true
	}
	return false

}

func SetNewDistances(object1, object2 *Shape) {
	collision_normal := linalg.NewVector((*object1).GetCenter(), (*object2).GetCenter())

	// Get the "radius" of each shape along the collision normal
	radius1Vector := linalg.NewVector((*object1).FurthestPoint(collision_normal), (*object1).GetCenter())
	radius1 := radius1Vector.Magnitude()

	radius2Vector := linalg.NewVector((*object2).FurthestPoint(collision_normal), (*object2).GetCenter())
	radius2 := radius2Vector.Magnitude()

	// Calculate centers and distance
	dx := (*object2).GetCenter().X - (*object1).GetCenter().X
	dy := (*object2).GetCenter().Y - (*object1).GetCenter().Y
	dist := collision_normal.Magnitude()

	// Calculate overlap and movement
	overlap := (radius1 + radius2 - dist) / 2

	// If there is overlap, move objects apart
	if overlap > 0 {
		moveX := (dx / dist) * overlap
		moveY := (dy / dist) * overlap

		// Update positions
		(*object1).SetCenter(Point{
			X: (*object1).GetCenter().X - moveX,
			Y: (*object1).GetCenter().Y - moveY,
		})

		(*object2).SetCenter(Point{
			X: (*object2).GetCenter().X + moveX,
			Y: (*object2).GetCenter().Y + moveY,
		})
	}
}
