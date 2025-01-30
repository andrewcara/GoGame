package dynamics

import linalg "HeadSoccer/math/helper"

// Values associated with any Physics Objects that govern how they are treated in Main
type DynamicProperties struct {
	Velocity   linalg.Vector
	Force      linalg.Vector
	Mass       float64
	Accelation linalg.Vector
}

//Implement basic case where there is only fricitonal force on the horizontal ground therefore no need to caluclate Y component of force

func (d *DynamicProperties) friction_force(coefficient_friction, gravity, timeDelta float64) {
	frictional_force := coefficient_friction * d.Mass * gravity

	if d.Velocity.X >= 0 {

		d.Velocity.X -= ((frictional_force / d.Mass) * timeDelta)
	} else {

		d.Velocity.X += ((frictional_force / d.Mass) * timeDelta)
	}

}
