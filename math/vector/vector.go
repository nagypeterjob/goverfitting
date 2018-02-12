package vector

import (
	utils "github.ibm.com/Content-Delivery-Org/goverfitting/math"
)

// Vector type
type Vector struct {
	N []float64
}

// NewVector creates new Vector
func NewVector(n []float64) *Vector {
	return &Vector{N: n}
}

// Add adds two vectors
func (v *Vector) Add(vect *Vector) *Vector {
	if v.Size() != vect.Size() {
		panic("failed vector addition - len mismatch")
	}
	return v.operation(vect, utils.Addition)
}

// Mul multiplicates two vectors
func (v *Vector) Mul(vect *Vector) *Vector {
	if v.Size() != vect.Size() {
		panic("failed vector multiplication - len mismatch")
	}
	return v.operation(vect, utils.Multiplication)
}

// Sub subtracts two vectors
func (v *Vector) Sub(vect *Vector) *Vector {
	if v.Size() != vect.Size() {
		panic("failed vector multiplication - len mismatch")
	}
	return v.operation(vect, utils.Subtraction)
}

func (v *Vector) operation(vect *Vector, op utils.Operand) *Vector {
	for i := range vect.N {
		v.N[i] = op(v.N[i], vect.N[i])
	}
	return v
}

// Neg negates vector value signs
func (v *Vector) Neg() *Vector {
	for i := range v.N {
		v.N[i] = -v.N[i]
	}
	return v
}

// Dot returns the dot product of two vectors
func (v *Vector) Dot(vect *Vector) float64 {
	if v.Size() != vect.Size() {
		panic("failed vector dot product - len mismatch")
	}
	var sum float64
	for i := range vect.N {
		sum += v.N[i] * vect.N[i]
	}
	return sum
}

// Size returns vector size
func (v *Vector) Size() int {
	return len(v.N)
}

// AddScalar adds scalar value to each vector element
func (v *Vector) AddScalar(scalar float64) *Vector {
	return v.scalarOperation(scalar, utils.Addition)
}

// MulScalar multiplies scalar value to each vector element
func (v *Vector) MulScalar(scalar float64) *Vector {
	return v.scalarOperation(scalar, utils.Multiplication)
}

func (v *Vector) scalarOperation(scalar float64, op utils.Operand) *Vector {
	if v.Size() == 0 || scalar == 1 {
		return v
	}
	for i := range v.N {
		v.N[i] = op(v.N[i], scalar)
	}
	return v
}
