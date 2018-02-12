package matrix

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
		},
	)
	expected := NewMatrix(
		[][]float64{
			{6, 6, 6, 6, 6},
			{6, 6, 6, 6, 6},
			{6, 6, 6, 6, 6},
		},
	)
	result := m.Add(m2)

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("Add is %+v, want %+v\n", result, expected)
	}
}

func TestMul(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
		},
	)
	expected := NewMatrix(
		[][]float64{
			{5, 8, 9, 8, 5},
			{5, 8, 9, 8, 5},
			{5, 8, 9, 8, 5},
		},
	)
	result := m.Mul(m2)

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("Mul is %+v, want %+v\n", result, expected)
	}
}

func TestMulNM(t *testing.T) {
	tcs := []struct {
		n *Matrix
		m *Matrix
		e *Matrix
	}{
		{
			NewMatrix(
				[][]float64{
					{1, 2, 3},
					{4, 5, 6},
				},
			),
			NewMatrix(
				[][]float64{
					{1, 4},
					{2, 5},
					{3, 6},
				},
			),
			NewMatrix(
				[][]float64{
					{14, 32},
					{32, 77},
				},
			),
		},
		{
			NewMatrix(
				[][]float64{
					{44, 1, 45, 56},
					{4, 5, 6, 9},
					{56, 63, 5, 6},
				},
			),
			NewMatrix(
				[][]float64{
					{1, 4, 64},
					{2, 5, 41},
					{3, 6, 0},
					{45, 99, 1},
				},
			),
			NewMatrix(
				[][]float64{
					{2701, 5995, 2913},
					{437, 968, 470},
					{467, 1163, 6173},
				},
			),
		},
	}

	for i, tc := range tcs {
		res := tc.n.Mul(tc.m)
		success := reflect.DeepEqual(res, tc.e)
		if !success {
			t.Errorf("[%d] MulNM is %+v, want %+v\n", i, res, tc.e)
		}
	}
}

func TestSub(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
			{5, 4, 3, 2, 1},
		},
	)
	expected := NewMatrix(
		[][]float64{
			{-4, -2, 0, 2, 4},
			{-4, -2, 0, 2, 4},
			{-4, -2, 0, 2, 4},
		},
	)
	result := m.Sub(m2)

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("Sub is %+v, want %+v\n", result, expected)
	}
}

func TestOperationChain(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3},
			{1, 2, 3},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{1, 1, 1},
			{1, 1, 1},
		},
	)
	m3 := NewMatrix(
		[][]float64{
			{2, 2},
			{2, 2},
			{25, 15},
		},
	)
	m4 := NewMatrix(
		[][]float64{
			{0, 5},
			{5, 0},
		},
	)

	expected := NewMatrix(
		[][]float64{
			{350, 550},
			{350, 550},
		},
	)
	result := m.Add(m2).Mul(m3).Mul(m4)

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("OperationChaining is %+v, want %+v\n", result, expected)
	}
}

func TestT(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 5},
		},
	)
	transp := m.T()
	expected := NewMatrix(
		[][]float64{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
			{4, 4, 5},
		},
	)
	success := reflect.DeepEqual(transp, expected)
	if !success {
		t.Errorf("T is %+v, want %+v\n", transp, expected)
	}
}

func TestShape(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
	)
	er := 3
	ec := 5
	rr, cc := m.Shape()

	rowSuccess := reflect.DeepEqual(er, rr)
	colSuccess := reflect.DeepEqual(ec, cc)
	if !rowSuccess {
		t.Errorf("Shape row is %+v, want %+v\n", er, rr)
	}
	if !colSuccess {
		t.Errorf("Shape col is %+v, want %+v\n", ec, cc)
	}
}

func TestDet(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{-2, 2, -3},
			{-1, 1, 3},
			{2, 0, -1},
		},
	)

	m2 := NewMatrix(
		[][]float64{
			{1, 2, 3, 4, 5},
			{5, 2, 3, 4, 5},
			{4, 2, 3, 4, 5},
			{3, 2, 3, 4, 5},
			{100, 2, 3, 4, 5},
		},
	)

	var expected float64 = 18
	var expectedBig float64
	det := m.Det(-1)
	detBig := m2.Det(-1)

	success := reflect.DeepEqual(det, expected)
	if !success {
		t.Errorf("Det is %+v, want %+v\n", det, expected)
	}
	successBig := reflect.DeepEqual(detBig, expectedBig)
	if !successBig {
		t.Errorf("Det is %+v, want %+v\n", detBig, expectedBig)
	}
}

func TestInverse(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{5, -2, 2, 7},
			{1, 0, 0, 3},
			{-3, 1, 5, 0},
			{3, -1, -9, 4},
		},
	)

	inv := m.Inverse()

	for i := range inv.N {
		for j := range inv.N[0] {
			inv.N[i][j] = float64(int(inv.N[i][j]*100+0.5)) / 100
		}
	}

	expected := NewMatrix(
		[][]float64{
			{-0.13, 0.86, -0.67, -0.4},
			{-0.63, 2.36, -0.92, -0.65},
			{0.05, 0.05, -0.01, -0.1},
			{0.05, 0.05, 0.23, 0.14},
		},
	)

	success := reflect.DeepEqual(inv, expected)
	if !success {
		t.Errorf("Inverse is %+v, want %+v\n", inv, expected)
	}
}

func TestInverseExample(t *testing.T) {

	m := NewMatrix(
		[][]float64{
			{3, 0, 2},
			{2, 0, -2},
			{0, 1, 1},
		},
	)

	inv := m.Inverse()
	expected := NewMatrix(
		[][]float64{
			{0.2, 0.2, 0},
			{-0.2, 0.3, 1},
			{0.2, -0.3, 0},
		},
	)

	for i := range inv.N {
		for j := range inv.N[0] {
			inv.N[i][j] = float64(int(inv.N[i][j]*100)) / 100
		}
	}

	success := reflect.DeepEqual(inv, expected)
	if !success {
		t.Errorf("Inverse is %+v, want %+v\n", inv, expected)
	}
}

func TestEmpty(t *testing.T) {
	m := NewMatrix(
		[][]float64{},
	)
	expected := true
	result := m.empty()

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("empty is %+v, want %+v\n", result, expected)
	}
}

func TestZero(t *testing.T) {
	m := NewMatrix(
		nil,
	)
	expected := NewMatrix(
		[][]float64{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	)
	result := m.Zero(3, 5)

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("Zero is %+v, want %+v\n", result, expected)
	}
}

func TestAt(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{0, 0, 0, 0, 0},
			{0, 0, 9999, 0, 0},
			{0, 0, 0, 0, 0},
		},
	)
	result := m.At(1, 2)
	var expected float64 = 9999

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("At is %+v, want %+v\n", result, expected)
	}
}

func TestSameShape(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{0, 0, 0},
			{0, 0, 0},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{3, 65, 103},
			{4, 44, 66},
		},
	)
	result := m.sameShape(m2)
	expected := true

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("sameShape is %+v, want %+v\n", result, expected)
	}
}

func TestNM(t *testing.T) {
	m := NewMatrix(
		[][]float64{
			{0, 0, 0},
			{0, 0, 0},
		},
	)
	m2 := NewMatrix(
		[][]float64{
			{3, 65},
			{4, 44},
			{103, 61},
		},
	)
	result := m.nm(m2)
	expected := true

	success := reflect.DeepEqual(result, expected)
	if !success {
		t.Errorf("nm is %+v, want %+v\n", result, expected)
	}
}
