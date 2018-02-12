package matrix

import (
	"math"

	utils "github.ibm.com/Content-Delivery-Org/goverfitting/math"
)

// Matrix type
type Matrix struct {
	N [][]float64
}

// NewMatrix creates new Matrix
func NewMatrix(n [][]float64) *Matrix {
	return &Matrix{N: n}
}

// Add adds two matrices
func (m *Matrix) Add(mat *Matrix) *Matrix {
	return m.operation(mat, utils.Addition)
}

// Mul multiplies two matrices
func (m *Matrix) Mul(mat *Matrix) *Matrix {
	// N*M multiplication
	if m.nm(mat) {
		return m.mulNM(mat, utils.Multiplication)
	}
	// N*N multiplication
	return m.operation(mat, utils.Multiplication)
}

// Sub subtracts two matrices
func (m *Matrix) Sub(mat *Matrix) *Matrix {
	return m.operation(mat, utils.Subtraction)
}

// MulNM multiplies N*M and M*N matrices
func (m *Matrix) mulNM(mat *Matrix, op utils.Operand) *Matrix {
	if m.empty() || mat.empty() {
		panic("failed Mul - empty matrix")
	}
	r, _ := m.Shape()

	var (
		n = NewMatrix(nil).Zero(r, r)
	)
	for i := range m.N {
		for k := 0; k < r; k++ {
			for j := range m.N[0] {
				n.N[k][i] += op(m.At(k, j), mat.At(j, i))
			}
		}
	}
	m = nil
	m = n
	return m
}

//MulScalar multiplies matrix values with scalar
func (m *Matrix) MulScalar(scalar float64) *Matrix {
	return m.scalarOperation(scalar, utils.Multiplication)
}

// Generic method that takes an operation, and performs it on two matrices
func (m *Matrix) operation(mat *Matrix, op utils.Operand) *Matrix {
	if m.empty() || mat.empty() {
		panic("failed Mul - empty matrix")
	}
	if !m.sameShape(mat) {
		panic("failed Mul - matrix shape mismatch")
	}

	for i := range m.N {
		for j := range m.N[0] {
			m.N[i][j] = op(m.At(i, j), mat.At(i, j))
		}
	}
	return m
}

func (m *Matrix) scalarOperation(scalar float64, op utils.Operand) *Matrix {
	if m.empty() {
		panic("failed scalar mul - empty matrix")
	}
	for i := range m.N {
		for j := range m.N[0] {
			m.N[i][j] = op(m.At(i, j), scalar)
		}
	}
	return m
}

// Det returns the determinant of the matrix
func (m *Matrix) Det(n int) float64 {

	if len(m.N) == 0 {
		return 1
	}
	row, _ := m.Shape()

	if row == 2 {
		return m.N[0][0]*m.N[1][1] - m.N[1][0]*m.N[0][1]
	}
	if n == -1 {
		n = row
	}

	var det float64
	var r, c int
	b := NewMatrix(nil).Zero(n-1, n-1)
	for k := 0; k < n; k++ {
		r = 0
		for i := 1; i < n; i++ {
			c = 0
			for j := 0; j < n; j++ {
				if j == k {
					continue
				}
				b.N[r][c] = m.N[i][j]
				c++
			}
			r++
		}
		det += math.Pow(-1, float64(k)) * m.N[0][k] * b.Det(len(b.N))
	}
	return det
}

//Inverse returns the inverse of the matrix
func (m *Matrix) Inverse() *Matrix {
	if m.empty() {
		panic("failed Inverse - empty matrix")
	}
	var det = m.Det(-1)
	if det != 0 {
		return m.adjoint().MulScalar(1 / det)
	}
	panic("There is no inverse for matrix")
}

//T transposes matrix
func (m *Matrix) T() *Matrix {
	if m.empty() {
		panic("failed T - empty matrix")
	}
	r, c := m.Shape()
	n := NewMatrix(nil).Zero(c, r)

	for i := range n.N {
		for j := range n.N[0] {
			n.N[i][j] = m.N[j][i]
		}
	}
	return n
}

func (m *Matrix) cofactor(tmp *Matrix, p, q int) {
	if m.empty() {
		panic("failed cofactor - empty matrix")
	}
	var i = 0
	var j = 0
	for row := range m.N {
		for col := range m.N[0] {

			if row != p && col != q {
				tmp.N[i][j] = m.N[row][col]
				j++
				if j == len(m.N)-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func (m *Matrix) adjoint() *Matrix {
	if m.empty() {
		panic("failed adjoint - empty matrix")
	}
	r, c := m.Shape()
	tmp := NewMatrix(nil).Zero(r, c)
	adj := NewMatrix(nil).Zero(r, c)
	if r == 1 {
		adj.N[0][0] = 1
		return adj
	}
	var sign = 1
	for i := range m.N {
		for j := range m.N[0] {
			m.cofactor(tmp, i, j)
			if (i+j)%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}
			adj.N[j][i] = float64(sign) * (tmp.Det(r - 1))
		}
	}
	return adj
}

// Shape returns matrix shape (r, c)
func (m *Matrix) Shape() (r, c int) {
	if m.empty() {
		panic("failed shape - empty matrix")
	}
	rows := len(m.N)
	cols := len(m.N[0])
	return rows, cols
}

func (m *Matrix) sameShape(mat *Matrix) bool {
	r, c := m.Shape()
	rMat, cMat := mat.Shape()
	return r == rMat && c == cMat
}

func (m *Matrix) nm(mat *Matrix) bool {
	r, c := m.Shape()
	rMat, cMat := mat.Shape()
	return r == cMat && c == rMat
}

func (m *Matrix) empty() bool {
	return len(m.N) == 0
}

// Zero creates an r * c matrix with zero values
func (m *Matrix) Zero(r, c int) *Matrix {
	n := make([][]float64, r)
	for i := 0; i < r; i++ {
		n[i] = make([]float64, c)
	}
	m.N = n
	return m
}

// At returns value for given row & col | indexed from 0
func (m *Matrix) At(r, c int) float64 {
	if m.empty() {
		panic("failed At - empty matrix")
	}
	return m.N[r][c]
}
