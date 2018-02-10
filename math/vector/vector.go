package vector

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
	n := make([]float64, len(v.N))
	for i := range vect.N {
		n[i] = v.N[i] + vect.N[i]
	}
	return NewVector(n)
}

// Mul multiplicates two vectors
func (v *Vector) Mul(vect *Vector) *Vector {
	if v.Size() != vect.Size() {
		panic("failed vector multiplication - len mismatch")
	}
	n := make([]float64, len(v.N))
	for i := range vect.N {
		n[i] = v.N[i] * vect.N[i]
	}
	return NewVector(n)
}

// Neg negates vector value signs
func (v *Vector) Neg() *Vector {

	n := make([]float64, len(v.N))
	for i := range v.N {
		n[i] = -v.N[i]
	}

	return NewVector(n)
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
	if v.Size() == 0 || scalar == 1 {
		return v
	}
	n := make([]float64, len(v.N))
	for i := range v.N {
		n[i] = v.N[i] + scalar
	}
	return NewVector(n)
}

// MulScalar multiplies scalar value to each vector element
func (v *Vector) MulScalar(scalar float64) *Vector {
	if v.Size() == 0 || scalar == 1 {
		return v
	}
	n := make([]float64, len(v.N))
	for i := range v.N {
		n[i] = v.N[i] * scalar
	}
	return NewVector(n)
}
