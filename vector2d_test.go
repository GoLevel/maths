package maths

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v1       Vector2D[float64]
		v2       Vector2D[float64]
		expected float64
	}{
		{Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}, 0},
		{Vector2D[float64]{X: 1, Y: 1}, Vector2D[float64]{X: 4, Y: 5}, 5},
		{Vector2D[float64]{X: -1, Y: -1}, Vector2D[float64]{X: -4, Y: -5}, 5},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, Vector2D[float64]{X: 4.5, Y: 6.5}, 5},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v1.Distance(tt.v2)
			if result != tt.expected {
				t.Errorf("Distance(%v, %v) = %v; expected %v", tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v1       Vector2D[float64]
		v2       Vector2D[float64]
		expected Vector2D[float64]
	}{
		{Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 1, Y: 1}, Vector2D[float64]{X: 2, Y: 2}, Vector2D[float64]{X: 3, Y: 3}},
		{Vector2D[float64]{X: -1, Y: -1}, Vector2D[float64]{X: 1, Y: 1}, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, Vector2D[float64]{X: 3.5, Y: 4.5}, Vector2D[float64]{X: 5, Y: 7}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v1.Add(tt.v2)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v; expected %v", tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v1       Vector2D[float64]
		v2       Vector2D[float64]
		expected Vector2D[float64]
	}{
		{Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 5, Y: 5}, Vector2D[float64]{X: 3, Y: 3}, Vector2D[float64]{X: 2, Y: 2}},
		{Vector2D[float64]{X: -1, Y: -1}, Vector2D[float64]{X: 1, Y: 1}, Vector2D[float64]{X: -2, Y: -2}},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, Vector2D[float64]{X: 0.5, Y: 1.5}, Vector2D[float64]{X: 1, Y: 1}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v1.Subtract(tt.v2)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; expected %v", tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v1       Vector2D[float64]
		scalar   float64
		expected Vector2D[float64]
	}{
		{Vector2D[float64]{X: 0, Y: 0}, 2, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 1, Y: 1}, 2, Vector2D[float64]{X: 2, Y: 2}},
		{Vector2D[float64]{X: -1, Y: -1}, 2, Vector2D[float64]{X: -2, Y: -2}},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, 2, Vector2D[float64]{X: 3, Y: 5}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v1.Multiply(tt.scalar)
			if result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; expected %v", tt.v1, tt.scalar, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v1       Vector2D[float64]
		scalar   float64
		expected Vector2D[float64]
	}{
		{Vector2D[float64]{X: 0, Y: 0}, 2, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 4, Y: 4}, 2, Vector2D[float64]{X: 2, Y: 2}},
		{Vector2D[float64]{X: -4, Y: -4}, 2, Vector2D[float64]{X: -2, Y: -2}},
		{Vector2D[float64]{X: 3, Y: 6}, 3, Vector2D[float64]{X: 1, Y: 2}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v1.Divide(tt.scalar)
			if result != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; expected %v", tt.v1, tt.scalar, result, tt.expected)
			}
		})
	}
}

func TestClone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v        Vector2D[float64]
		expected Vector2D[float64]
	}{
		{Vector2D[float64]{X: 0, Y: 0}, Vector2D[float64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 1, Y: 1}, Vector2D[float64]{X: 1, Y: 1}},
		{Vector2D[float64]{X: -1, Y: -1}, Vector2D[float64]{X: -1, Y: -1}},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, Vector2D[float64]{X: 1.5, Y: 2.5}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v.Clone()
			if result != tt.expected {
				t.Errorf("Clone() = %v; expected %v", result, tt.expected)
			}
		})
	}
}
func TestToInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v        Vector2D[float64]
		expected Vector2D[int64]
	}{
		{Vector2D[float64]{X: 0.0, Y: 0.0}, Vector2D[int64]{X: 0, Y: 0}},
		{Vector2D[float64]{X: 1.5, Y: 2.5}, Vector2D[int64]{X: 1, Y: 2}},
		{Vector2D[float64]{X: -1.5, Y: -2.5}, Vector2D[int64]{X: -1, Y: -2}},
		{Vector2D[float64]{X: 1.9, Y: 2.9}, Vector2D[int64]{X: 1, Y: 2}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v.ToInt()
			if result != tt.expected {
				t.Errorf("ToInt() = %v; expected %v", result, tt.expected)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		v        Vector2D[int64]
		expected Vector2D[float64]
	}{
		{Vector2D[int64]{X: 0, Y: 0}, Vector2D[float64]{X: 0.0, Y: 0.0}},
		{Vector2D[int64]{X: 1, Y: 2}, Vector2D[float64]{X: 1.0, Y: 2.0}},
		{Vector2D[int64]{X: -1, Y: -2}, Vector2D[float64]{X: -1.0, Y: -2.0}},
		{Vector2D[int64]{X: 123456789, Y: 987654321}, Vector2D[float64]{X: 123456789.0, Y: 987654321.0}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			result := tt.v.ToFloat()
			if result != tt.expected {
				t.Errorf("ToFloat() = %v; expected %v", result, tt.expected)
			}
		})
	}
}
