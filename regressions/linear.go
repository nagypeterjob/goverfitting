package regressions

import (
	"math/rand"

	mat "gonum.org/v1/gonum/mat"
)

type LinearRegression struct {
	TrainX *mat.Dense
	Trainy *mat.Dense
}

// NewLinearRegression constructor
func NewLinearRegression(trainX, trainy *mat.Dense) *LinearRegression {
	return &LinearRegression{trainX, trainy}
}

// NormalEquation implementation
func (lr *LinearRegression) NormalEquation() *mat.Dense {
	var T = lr.TrainX.T()
	var TdotX mat.Dense
	var inv mat.Dense
	var b mat.Dense
	var thetaBest mat.Dense
	//theta = (X.T * X)^-1 * X.T * y
	TdotX.Mul(T, lr.TrainX)
	inv.Inverse(&TdotX)
	b.Mul(T, lr.Trainy)
	thetaBest.Mul(&inv, &b)
	return mat.DenseCopyOf(&thetaBest)
}

// BGD (Batch Gradient Descent) training implementation
func (lr *LinearRegression) BGD(eta float64, iterations int) *mat.Dense {
	m, c := lr.TrainX.Dims()

	rnd := make([]float64, c)
	for i := range rnd {
		rnd[i] = rand.NormFloat64()
	}
	theta := mat.NewDense(c, 1, rnd)

	var hypothesis mat.Dense
	var loss mat.Dense
	var gradient mat.Dense
	var transposed mat.Dense

	for i := 0; i < iterations; i++ {
		hypothesis.Mul(lr.TrainX, theta)
		loss.Sub(&hypothesis, lr.Trainy)
		transposed.Mul(lr.TrainX.T(), &loss)
		gradient.Scale(float64(2)/float64(m), &transposed)
		gradient.Scale(eta, &gradient)
		theta.Sub(theta, &gradient)
	}
	return theta
}

// SGD (Stohastic Gradient Descent) implementation
func (lr *LinearRegression) SGD(epochs int, t0, t1 float64) *mat.Dense {
	m, c := lr.TrainX.Dims()

	rnd := make([]float64, c)
	for i := range rnd {
		rnd[i] = rand.Float64()
	}
	theta := mat.NewDense(c, 1, rnd)
	var gradient mat.Dense
	var hypothesis mat.Dense
	var transposed mat.Dense
	var loss mat.Dense

	for i := 0; i < epochs; i++ {
		for j := 0; j < m; j++ {
			randomIndex := rand.Intn(m)
			Xi := lr.TrainX.RawRowView(randomIndex)
			Xy := lr.Trainy.RawRowView(randomIndex)
			X := mat.NewDense(1, c, Xi)
			y := mat.NewDense(1, 1, Xy)

			hypothesis.Mul(X, theta)
			loss.Sub(&hypothesis, y)
			transposed.Mul(X.T(), &loss)
			// gradients = 2 * X.T * (X * theta - y)
			gradient.Scale(float64(2), &transposed)
			eta := learningSchedule(float64(i*m+j), t0, t1)
			// theta = theta - eta * gradients
			gradient.Scale(eta, &gradient)
			theta.Sub(theta, &gradient)
		}
	}
	return theta
}

// MBGD is a Mini-batch gradient descent naive implementaiton
func (lr *LinearRegression) MBGD(iterations, minibatchSize int, t0, t1 float64) *mat.Dense {
	m, c := lr.TrainX.Dims()
	_, yc := lr.Trainy.Dims()
	rnd := make([]float64, c)
	for i := range rnd {
		rnd[i] = rand.NormFloat64()
	}
	theta := mat.NewDense(c, 1, rnd)

	var gradient mat.Dense
	var hypothesis mat.Dense
	var transposed mat.Dense
	var loss mat.Dense

	var t float64
	for i := 0; i < iterations; i++ {
		shuffledIndices := rand.Perm(m)
		X := shuffle(lr.TrainX, shuffledIndices)
		y := shuffle(lr.Trainy, shuffledIndices)
		for j := 0; j < m; j += minibatchSize {
			t++
			Xi := X.Slice(j, j+minibatchSize, 0, c)
			yi := y.Slice(j, j+minibatchSize, 0, yc)

			hypothesis.Mul(Xi, theta)
			loss.Sub(&hypothesis, yi)
			transposed.Mul(Xi.T(), &loss)
			// gradients = 2 * X.T * (X * theta - y)
			gradient.Scale(float64(2), &transposed)
			eta := learningSchedule(t, t0, t1)
			// theta = theta - eta * gradients
			gradient.Scale(eta, &gradient)
			theta.Sub(theta, &gradient)
		}
	}
	return theta
}

func shuffle(matrix *mat.Dense, permutations []int) *mat.Dense {
	r, c := matrix.Dims()
	dest := mat.NewDense(r, c, nil)

	for i, v := range permutations {
		dest.SetRow(v, matrix.RawRowView(i))
	}
	return dest
}

func learningSchedule(t, t0, t1 float64) float64 {
	return t0 / (t + t1)
}

func (lr *LinearRegression) Predict(theta, x *mat.Dense) mat.Dense {
	var result mat.Dense
	result.Mul(x, theta)
	return result
}
