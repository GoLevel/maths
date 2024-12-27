package maths

import "math"

// Vector2D is a generic implementation of a 2D vector
type Vector2D[T interface {
	int64 | float64
}] struct {
	X, Y T
}

// New creates a new Vector2D
func NewVector2D[T interface {
	int64 | float64
}](x, y T) Vector2D[T] {
	return Vector2D[T]{X: x, Y: y}
}

// Add adds another vector to the current vector
func (v Vector2D[T]) Add(other Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Subtract subtracts another vector from the current vector
func (v Vector2D[T]) Subtract(other Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Multiply multiplies the vector by a scalar
func (v Vector2D[T]) Multiply(scalar T) Vector2D[T] {
	return Vector2D[T]{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

// Divide divides the vector by a scalar
func (v Vector2D[T]) Divide(scalar T) Vector2D[T] {
	return Vector2D[T]{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

// Clone creates a copy of the vector
func (v Vector2D[T]) Clone() Vector2D[T] {
	return Vector2D[T]{
		X: v.X,
		Y: v.Y,
	}
}

// Distance calculates the Euclidean distance between two vectors
func (v Vector2D[T]) Distance(other Vector2D[T]) float64 {
	dx := float64(v.X - other.X)
	dy := float64(v.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// ToInt converts the vector to a Vector2D with int64 components
func (v Vector2D[T]) ToInt() Vector2D[int64] {
	return Vector2D[int64]{
		X: int64(v.X),
		Y: int64(v.Y),
	}
}

// ToFloat converts the vector to a Vector2D with float64 components
func (v Vector2D[T]) ToFloat() Vector2D[float64] {
	return Vector2D[float64]{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}
