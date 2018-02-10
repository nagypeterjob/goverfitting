package vector

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{3, 2, 1})
	v3 := v1.Add(v2)
	expected := NewVector([]float64{4, 4, 4})

	for i := range expected.N {
		success := reflect.DeepEqual(v3.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Add is %+v, want %+v\n", i, v3.N[i], expected.N[i])
		}
	}
}

func TestMul(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := NewVector([]float64{3, 2, 1})
	v3 := v1.Mul(v2)
	expected := NewVector([]float64{3, 4, 3})

	for i := range expected.N {
		success := reflect.DeepEqual(v3.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Mul is %+v, want %+v\n", i, v3.N[i], expected.N[i])
		}
	}
}

func TestNeg(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3})
	v2 := v1.Neg()
	expected := NewVector([]float64{-1, -2, -3})

	for i := range expected.N {
		success := reflect.DeepEqual(v2.N[i], expected.N[i])
		if !success {
			t.Errorf("[%d] Neg is %+v, want %+v\n", i, v2.N[i], expected.N[i])
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
