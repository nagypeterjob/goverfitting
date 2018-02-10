package vector

type Vector struct {
	N []float64
}

func NewVector(n []float64) *Vector {
	return &Vector{N: n}
}

// Add adds two vectors
func (v *Vector) Add(vect *Vector) *Vector {
	if len(v.N) != len(vect.N) {
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
	if len(v.N) != len(vect.N) {
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
	if len(v.N) != len(vect.N) {
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
