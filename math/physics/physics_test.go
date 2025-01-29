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
			name: "Polygon-Polygon Overlap",
			circle1: &shapes.Polygon{
				Vertices: []point.Point{
					{X: 0, Y: 0},
					{X: 5, Y: 0},
					{X: 5, Y: 5},
					{X: 0, Y: 5},
				},
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Polygon{
				Vertices: []point.Point{
					{X: 4, Y: 4},
					{X: 9, Y: 4},
					{X: 9, Y: 9},
					{X: 4, Y: 9},
				},
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
			name: "Polygon-Polygon No Overlap",
			circle1: &shapes.Polygon{
				Vertices: []point.Point{
					{X: 0, Y: 0},
					{X: 5, Y: 0},
					{X: 5, Y: 5},
					{X: 0, Y: 5},
				},
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Polygon{
				Vertices: []point.Point{
					{X: 10, Y: 10},
					{X: 15, Y: 10},
					{X: 15, Y: 15},
					{X: 10, Y: 15},
				},
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
			name: "Polygon-Circle Overlap",
			circle1: &shapes.Polygon{
				Center: point.Point{X: 2.5, Y: 2.5},
				Vertices: []point.Point{
					{X: 0, Y: 0},
					{X: 5, Y: 0},
					{X: 5, Y: 5},
					{X: 0, Y: 5},
				},
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 4, Y: 4},
				Radius: 3,
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
			name: "Polygon-Circle No Overlap",
			circle1: &shapes.Polygon{
				Vertices: []point.Point{
					{X: 0, Y: 0},
					{X: 5, Y: 0},
					{X: 5, Y: 5},
					{X: 0, Y: 5},
				},
				Dynamic: dynamics.DynamicProperties{
					Velocity:   linalg.Vector{X: 0, Y: 0},
					Force:      linalg.Vector{X: 0, Y: -9.8},
					Mass:       1.0,
					Accelation: gravity,
				},
			},
			circle2: &shapes.Circle{
				Center: point.Point{X: 10, Y: 10},
				Radius: 3,
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
