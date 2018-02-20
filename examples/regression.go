package examples

import (
	"math/rand"

	"github.ibm.com/Content-Delivery-Org/goverfitting/plot"
	"github.ibm.com/Content-Delivery-Org/goverfitting/regressions"
	"gonum.org/v1/gonum/mat"
)

func RunLinearRegression() {

	//Generating samle data - X and y
	data := make([]float64, 100)
	ydata := make([]float64, 50)

	for i := range data {
		if i%2 != 0 {
			data[i] = 4 * rand.Float64()
		} else {
			data[i] = 1
		}
	}
	X := mat.NewDense(50, 2, data)

	for i := range ydata {
		ydata[i] = 4 + 3*X.At(i, 0) + rand.NormFloat64()
	}
	y := mat.NewDense(50, 1, ydata)

	linear := regressions.NewLinearRegression(X, y)

	//Normal Equation - var theta = linear.NormalEquation()
	//Batch Gradient Descent - var theta = linear.BGD(0.1, 1000)
	//Stohastic Gradient Descent - var theta = linear.SGD(50, 5, 50)
	var theta = linear.MBGD(50, 10, 10, 1000) //Mini-Batch Gradient Descent

	testX := mat.NewDense(2, 2, []float64{
		1, 0,
		1, 2,
	})
	prediction := linear.Predict(theta, testX)

	// Plotting
	p := plot.NewPlot("Linear Regression", "x", "y")
	p.Init(true, 0, 4, 0, 20)

	// Drawing data points on plot
	scatter := p.ToPlotterPoint(X, y)

	pred := p.ToPlotterPoint(testX, &prediction)

	//Calculating slope and intercept to plot a "long line"
	slope := (pred[1].Y - pred[0].Y) / (pred[1].X - pred[0].X)
	intercept := pred[1].Y - slope

	p.AddScatter(scatter)
	p.CreateLongLine(slope, intercept, "Regression")
	p.Save("linreg.png")
}
