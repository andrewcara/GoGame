package physics

import (
	linalg "HeadSoccer/math/helper"
	dynamics "HeadSoccer/math/helper/dynamic_properties"
	"HeadSoccer/math/helper/point"
	"HeadSoccer/shapes"
	"testing"
)

func TestGJK(t *testing.T) {
	// Default gravity for test cases
	gravity := linalg.Vector{X: 0, Y: -9.8}

	tests := []struct {
		name     string
		circle1  shapes.Shape
		circle2  shapes.Shape
		expected bool
	}{
		{
			name: "Overlapping circles - Same position",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: true,
		},
		{
			name: "Overlapping circles - Partial overlap",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 8, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: true,
		},
		{
			name: "Non-overlapping circles - Just touching",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 10, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: true,
		},
		{
			name: "Non-overlapping circles - Far apart",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 20, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: false,
		},
		{
			name: "Different sized circles - Overlapping",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 10,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 5, Y: 5},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GJK(tt.circle1, tt.circle2)
			if result != tt.expected {
				t.Errorf("%s: GJK() = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

// TestGJKEdgeCases tests special cases that might cause issues
func TestGJKEdgeCases(t *testing.T) {
	gravity := linalg.Vector{X: 0, Y: -9.8}

	tests := []struct {
		name     string
		circle1  shapes.Shape
		circle2  shapes.Shape
		expected bool
	}{
		{
			name: "Very small distances",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 10.0001, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: false,
		},
		{
			name: "Moving circles",
			circle1: &shapes.Circle{
				Center: point.Point{X: 0, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0.00000006, Y: 0.00000006},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 8, Y: 0},
				Radius: 5,
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: -0.00000006, Y: 0.00000006},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GJK(tt.circle1, tt.circle2)
			if result != tt.expected {
				t.Errorf("%s: GJK() = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
