package vector

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{3, 2, 1})
	v1.Add(v2)
	expected := NewVector([]float64{4, 4, 4})

	for i := range expected.N {
		success := reflect.DeepEqual(v1.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Add is %+v, want %+v\n", i, v1.N[i], expected.N[i])
		}
	}
}

func TestMul(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{3, 2, 1})
	v1.Mul(v2)
	expected := NewVector([]float64{3, 4, 3})

	for i := range expected.N {
		success := reflect.DeepEqual(v1.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Mul is %+v, want %+v\n", i, v1.N[i], expected.N[i])
		}
	}
}

func TestOperationChain(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{3, 2, 1})
	v3 := NewVector([]float64{5, 5, 5})
	var scalar float64 = 15

	v1.Mul(v2).Add(v3).AddScalar(scalar)

	expected := NewVector([]float64{23, 24, 23})

	for i := range expected.N {
		success := reflect.DeepEqual(v1.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] OperationChaining is %+v, want %+v\n", i, v1.N[i], expected.N[i])
		}
	}
}

func TestNeg(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v1.Neg()
	expected := NewVector([]float64{-1, -2, -3})

	for i := range expected.N {
		success := reflect.DeepEqual(v1.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Neg is %+v, want %+v\n", i, v1.N[i], expected.N[i])
		}
	}
}

func TestDot(t *testing.T) {
	v1 := NewVector([]float64{1, 3, 5})
	v2 := NewVector([]float64{4, -2, 1})
	dot := v1.Dot(v2)
	var expected float64 = 3

	success := reflect.DeepEqual(dot, expected)
	if !success {
		t.Errorf("Dot is %+v, want %+v\n", dot, expected)
	}
}

func TestSize(t *testing.T) {
	v1 := NewVector([]float64{1, 3, 5})
	size := v1.Size()
	expected := len(v1.N)

	success := reflect.DeepEqual(size, expected)
	if !success {
		t.Errorf("Size is %+v, want %+v\n", size, expected)
	}
}

func TestAddScalar(t *testing.T) {
	tcs := []struct {
		n        *Vector
		scalar   float64
		expected *Vector
	}{
		{
			NewVector([]float64{1, 3, 5}),
			3,
			NewVector([]float64{4, 6, 8}),
		},
		{
			NewVector(nil),
			3,
			NewVector(nil),
		},
		{
			NewVector([]float64{3, 3, 3}),
			1,
			NewVector([]float64{3, 3, 3}),
		},
	}

	for i, tc := range tcs {
		tc.n.AddScalar(tc.scalar)
		success := reflect.DeepEqual(tc.n, tc.expected)
		if !success {
			t.Errorf("[%d] AddScalar is %+v, want %+v\n", i, tc.n, tc.expected)
		}
	}
}

func TestMulScalar(t *testing.T) {
	tcs := []struct {
		n        *Vector
		scalar   float64
		expected *Vector
	}{
		{
			NewVector([]float64{1, 3, 5}),
			3,
			NewVector([]float64{3, 9, 15}),
		},
		{
			NewVector(nil),
			3,
			NewVector(nil),
		},
		{
			NewVector([]float64{3, 3, 3}),
			1,
			NewVector([]float64{3, 3, 3}),
		},
		{
			NewVector([]float64{1, 1, 1}),
			0,
			NewVector([]float64{0, 0, 0}),
		},
	}

	for i, tc := range tcs {
		tc.n.MulScalar(tc.scalar)
		success := reflect.DeepEqual(tc.n, tc.expected)
		if !success {
			t.Errorf("[%d] MulScalar is %+v, want %+v\n", i, tc.n, tc.expected)
		}
	}
}
