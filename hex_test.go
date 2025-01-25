package maths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex1     Hex[int64]
		hex2     Hex[int64]
		expected Hex[int64]
	}{
		{
			name:     "Add positive int64",
			hex1:     Hex[int64]{Q: 2, R: 3},
			hex2:     Hex[int64]{Q: 1, R: 2},
			expected: Hex[int64]{Q: 3, R: 5},
		},
		{
			name:     "Add negative int64",
			hex1:     Hex[int64]{Q: -2, R: -3},
			hex2:     Hex[int64]{Q: -1, R: -2},
			expected: Hex[int64]{Q: -3, R: -5},
		},
		{
			name:     "Add mixed int64",
			hex1:     Hex[int64]{Q: -2, R: 3},
			hex2:     Hex[int64]{Q: 1, R: -2},
			expected: Hex[int64]{Q: -1, R: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex1.Add(tt.hex2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexAddFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex1     Hex[float64]
		hex2     Hex[float64]
		expected Hex[float64]
	}{
		{
			name:     "Add positive float64",
			hex1:     Hex[float64]{Q: 2.5, R: 3.5},
			hex2:     Hex[float64]{Q: 1.5, R: 2.5},
			expected: Hex[float64]{Q: 4.0, R: 6.0},
		},
		{
			name:     "Add negative float64",
			hex1:     Hex[float64]{Q: -2.5, R: -3.5},
			hex2:     Hex[float64]{Q: -1.5, R: -2.5},
			expected: Hex[float64]{Q: -4.0, R: -6.0},
		},
		{
			name:     "Add mixed float64",
			hex1:     Hex[float64]{Q: -2.5, R: 3.5},
			hex2:     Hex[float64]{Q: 1.5, R: -2.5},
			expected: Hex[float64]{Q: -1.0, R: 1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex1.Add(tt.hex2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexSubtract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex1     Hex[int64]
		hex2     Hex[int64]
		expected Hex[int64]
	}{
		{
			name:     "Subtract positive int64",
			hex1:     Hex[int64]{Q: 4, R: 6},
			hex2:     Hex[int64]{Q: 1, R: 2},
			expected: Hex[int64]{Q: 3, R: 4},
		},
		{
			name:     "Subtract negative int64",
			hex1:     Hex[int64]{Q: -4, R: -6},
			hex2:     Hex[int64]{Q: -1, R: -2},
			expected: Hex[int64]{Q: -3, R: -4},
		},
		{
			name:     "Subtract mixed int64",
			hex1:     Hex[int64]{Q: -2, R: 2},
			hex2:     Hex[int64]{Q: 1, R: -2},
			expected: Hex[int64]{Q: -3, R: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex1.Subtract(tt.hex2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexSubtractFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex1     Hex[float64]
		hex2     Hex[float64]
		expected Hex[float64]
	}{
		{
			name:     "Subtract positive float64",
			hex1:     Hex[float64]{Q: 5.0, R: 7.0},
			hex2:     Hex[float64]{Q: 1.5, R: 2.5},
			expected: Hex[float64]{Q: 3.5, R: 4.5},
		},
		{
			name:     "Subtract negative float64",
			hex1:     Hex[float64]{Q: -5.0, R: -7.0},
			hex2:     Hex[float64]{Q: -1.5, R: -2.5},
			expected: Hex[float64]{Q: -3.5, R: -4.5},
		},
		{
			name:     "Subtract mixed float64",
			hex1:     Hex[float64]{Q: -2.0, R: 2.0},
			hex2:     Hex[float64]{Q: 1.5, R: -2.5},
			expected: Hex[float64]{Q: -3.5, R: 4.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex1.Subtract(tt.hex2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexMultiply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex      Hex[int64]
		scalar   int64
		expected Hex[int64]
	}{
		{
			name:     "Multiply positive int64",
			hex:      Hex[int64]{Q: 2, R: 3},
			scalar:   2,
			expected: Hex[int64]{Q: 4, R: 6},
		},
		{
			name:     "Multiply negative int64",
			hex:      Hex[int64]{Q: -2, R: -3},
			scalar:   2,
			expected: Hex[int64]{Q: -4, R: -6},
		},
		{
			name:     "Multiply mixed int64",
			hex:      Hex[int64]{Q: 2, R: -3},
			scalar:   -2,
			expected: Hex[int64]{Q: -4, R: 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex.Multiply(tt.scalar)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexMultiplyFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex      Hex[float64]
		scalar   float64
		expected Hex[float64]
	}{
		{
			name:     "Multiply positive float64",
			hex:      Hex[float64]{Q: 2.5, R: 3.5},
			scalar:   2.0,
			expected: Hex[float64]{Q: 5.0, R: 7.0},
		},
		{
			name:     "Multiply negative float64",
			hex:      Hex[float64]{Q: -2.5, R: -3.5},
			scalar:   2.0,
			expected: Hex[float64]{Q: -5.0, R: -7.0},
		},
		{
			name:     "Multiply mixed float64",
			hex:      Hex[float64]{Q: 2.5, R: -3.5},
			scalar:   -2.0,
			expected: Hex[float64]{Q: -5.0, R: 7.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex.Multiply(tt.scalar)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexDivide(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex      Hex[int64]
		scalar   int64
		expected Hex[int64]
	}{
		{
			name:     "Divide positive int64",
			hex:      Hex[int64]{Q: 4, R: 6},
			scalar:   2,
			expected: Hex[int64]{Q: 2, R: 3},
		},
		{
			name:     "Divide negative int64",
			hex:      Hex[int64]{Q: -4, R: -6},
			scalar:   2,
			expected: Hex[int64]{Q: -2, R: -3},
		},
		{
			name:     "Divide mixed int64",
			hex:      Hex[int64]{Q: -4, R: 6},
			scalar:   -2,
			expected: Hex[int64]{Q: 2, R: -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex.Divide(tt.scalar)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexDivideFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hex      Hex[float64]
		scalar   float64
		expected Hex[float64]
	}{
		{
			name:     "Divide positive float64",
			hex:      Hex[float64]{Q: 5.0, R: 7.0},
			scalar:   2.0,
			expected: Hex[float64]{Q: 2.5, R: 3.5},
		},
		{
			name:     "Divide negative float64",
			hex:      Hex[float64]{Q: -5.0, R: -7.0},
			scalar:   2.0,
			expected: Hex[float64]{Q: -2.5, R: -3.5},
		},
		{
			name:     "Divide mixed float64",
			hex:      Hex[float64]{Q: -5.0, R: 7.0},
			scalar:   -2.0,
			expected: Hex[float64]{Q: 2.5, R: -3.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.hex.Divide(tt.scalar)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHexNeighbours(t *testing.T) {
	t.Parallel()

	test := NewHex(int64(0), int64(0))
	expected := []Hex[int64]{
		{Q: 1, R: 0}, {Q: 0, R: 1}, {Q: -1, R: 1},
		{Q: -1, R: 0}, {Q: 0, R: -1}, {Q: 1, R: -1},
	}

	result := test.Neighbours()
	assert.Len(t, result, 6)
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], r)
		}
	}
}

func TestHexSpiralRing(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		radius   int
		center   Hex[int64]
		expected []Hex[int64]
	}{
		{
			radius:   0,
			center:   Hex[int64]{Q: 0, R: 0},
			expected: []Hex[int64]{{Q: 0, R: 0}},
		},
		{
			radius: 1,
			center: Hex[int64]{Q: 0, R: 0},
			expected: []Hex[int64]{
				{Q: 0, R: 1}, {Q: -1, R: 1}, {Q: -1, R: 0},
				{Q: 0, R: -1}, {Q: +1, R: -1}, {Q: 1, R: 0},
			},
		},
		{
			radius: 2,
			center: Hex[int64]{Q: 0, R: 0},
			expected: []Hex[int64]{
				{Q: 1, R: 1}, {Q: 0, R: 2}, {Q: -1, R: 2},
				{Q: -2, R: 2}, {Q: -2, R: 1}, {Q: -2, R: 0},
				{Q: -1, R: -1}, {Q: 0, R: -2}, {Q: 1, R: -2},
				{Q: 2, R: -2}, {Q: 2, R: -1}, {Q: 2, R: 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("radius %d", testCase.radius), func(t *testing.T) {
			t.Parallel()
			result := testCase.center.SpiralRing(testCase.radius)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
