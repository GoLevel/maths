package maths

import (
	"fmt"
	"math"
)

// Hex represents a hexagonal cell in axial coordinates (q,r)
type Hex[T interface {
	int64 | float64
}] struct {
	Q, R T
}

// NewHex creates a new Hex with the given coordinates
func NewHex[T interface {
	int64 | float64
}](q, r T) Hex[T] {
	return Hex[T]{Q: q, R: r}
}

// String returns the coordinates as `q:r`
func (h Hex[T]) String() string {
	return fmt.Sprintf("%v:%v", h.Q, h.R)
}

// Add returns the sum of two hexes
func (h Hex[T]) Add(other Hex[T]) Hex[T] {
	return Hex[T]{
		Q: h.Q + other.Q,
		R: h.R + other.R,
	}
}

// Subtract returns the difference between two hexes
func (h Hex[T]) Subtract(other Hex[T]) Hex[T] {
	return Hex[T]{
		Q: h.Q - other.Q,
		R: h.R - other.R,
	}
}

// Multiply returns the hex multiplied by a scalar
func (h Hex[T]) Multiply(scalar T) Hex[T] {
	return Hex[T]{
		Q: h.Q * scalar,
		R: h.R * scalar,
	}
}

// Divide returns the hex divided by a scalar
func (h Hex[T]) Divide(scalar T) Hex[T] {
	return Hex[T]{
		Q: h.Q / scalar,
		R: h.R / scalar,
	}
}

// Clone returns a copy of the hex
func (h Hex[T]) Clone() Hex[T] {
	return Hex[T]{
		Q: h.Q,
		R: h.R,
	}
}

// Distance returns the distance between two hexes
func (h Hex[T]) Distance(other Hex[T]) float64 {
	// Convert to float64 for calculations
	q1, r1 := float64(h.Q), float64(h.R)
	q2, r2 := float64(other.Q), float64(other.R)

	// Calculate s coordinates (s = -q-r)
	s1 := -q1 - r1
	s2 := -q2 - r2

	// Use manhattan distance formula for hex grids
	return math.Abs(q1-q2) + math.Abs(r1-r2) + math.Abs(s1-s2)/2
}

// directions represents the six directions in a hexagonal grid
var directions = []Hex[int64]{
	{Q: 1, R: 0}, {Q: 0, R: 1},
	{Q: -1, R: 1}, {Q: -1, R: 0},
	{Q: 0, R: -1}, {Q: +1, R: -1},
}

// Neighbours returns all adjacent hexes
func (h Hex[T]) Neighbours() []Hex[T] {
	result := make([]Hex[T], 6)
	for i, dir := range directions {
		result[i] = Hex[T]{
			Q: h.Q + T(dir.Q),
			R: h.R + T(dir.R),
		}
	}
	return result
}

// Circle returns all hexes at the given radius
func (h Hex[T]) Circle(radius int) []Hex[T] {
	var results []Hex[T]
	if radius < 0 {
		return results
	}

	// Start at the top-right corner
	hex := h.Add(Hex[T]{Q: T(radius), R: T(-radius)})

	// For each of the six sides
	for i := 0; i < 6; i++ {
		// Move radius times in each direction
		for j := 0; j < radius; j++ {
			results = append(results, hex)
			hex = hex.Add(Hex[T]{Q: T(directions[i].Q), R: T(directions[i].R)})
		}
	}

	return results
}

// LineTo returns a line of hexes from this hex to another
func (h Hex[T]) LineTo(other Hex[T]) []Hex[T] {
	distance := int(h.Distance(other))
	results := make([]Hex[T], distance+1)

	// Use linear interpolation
	for i := 0; i <= distance; i++ {
		t := float64(i) / float64(distance)
		results[i] = Hex[T]{
			Q: T(float64(h.Q)*(1-t) + float64(other.Q)*t),
			R: T(float64(h.R)*(1-t) + float64(other.R)*t),
		}
	}

	return results
}

// SpiralRing returns a single ring of hexes at the given radius
func (h Hex[T]) SpiralRing(radius int) []Hex[T] {
	var results []Hex[T]
	if radius < 0 {
		return results
	}
	if radius == 0 {
		return []Hex[T]{h}
	}

	directionScale := directions[0].Multiply(int64(radius))
	hex := h.Add(Hex[T]{Q: T(directionScale.Q), R: T(-directionScale.R)})
	for i := 0; i < 6; i++ {
		d := (i + 2) % 6

		for j := 0; j < radius; j++ {
			hex = hex.Add(Hex[T]{Q: T(directions[d].Q), R: T(directions[d].R)})
			results = append(results, hex)
		}
	}

	return results
}

// Spiral returns all hexes in a spiral pattern up to the given radius
func (h Hex[T]) Spiral(radius int) []Hex[T] {
	results := []Hex[T]{h} // Start with center

	// For each radius
	for r := 1; r <= radius; r++ {
		results = append(results, h.SpiralRing(r)...)
	}

	return results
}

// Round returns the hex with rounded components
func (h Hex[T]) Round() Hex[T] {
	return Hex[T]{
		Q: T(math.Round(float64(h.Q))),
		R: T(math.Round(float64(h.R))),
	}
}

// ToInt converts the hex to a Hex with int64 components
func (h Hex[T]) ToInt() Hex[int64] {
	return Hex[int64]{
		Q: int64(h.Q),
		R: int64(h.R),
	}
}

// ToFloat converts the hex to a Hex with float64 components
func (h Hex[T]) ToFloat() Hex[float64] {
	return Hex[float64]{
		Q: float64(h.Q),
		R: float64(h.R),
	}
}
