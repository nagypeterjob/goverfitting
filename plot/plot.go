package plot

import (
	"fmt"
	"image/color"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Plot struct {
	Title  string
	LabelX string
	LabelY string
	P      *plot.Plot
}

func NewPlot(title, labelx, labely string) *Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	return &Plot{title, labelx, labely, p}
}

func (p *Plot) Init(grid bool, xmin, xmax, ymin, ymax float64) {
	p.P.Title.Text = p.Title
	p.P.X.Label.Text = p.LabelX
	p.P.Y.Label.Text = p.LabelY
	p.P.X.Min = xmin
	p.P.X.Max = xmax
	p.P.Y.Min = ymin
	p.P.Y.Max = ymax
	if grid == true {
		p.P.Add(plotter.NewGrid())
	}
}

func (p *Plot) AddScatter(data plotter.XYs) {
	s, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.P.Add(s)
	p.P.Legend.Add("Points", s)
}

func (p *Plot) AddLine(data plotter.XYs) {
	l, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}
	p.P.Add(l)
	p.P.Legend.Add("regression", l)
}

func (p *Plot) Save(filename string) {
	if err := p.P.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprint("images/", filename)); err != nil {
		panic(err)
	}
}

func (p *Plot) CreateLongLine(m float64, intercept float64, legend string) {
	fun := plotter.NewFunction(func(x float64) float64 { return m*x + intercept })
	fun.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	fun.Width = vg.Points(2)
	fun.Color = color.RGBA{B: 255, A: 255}
	p.P.Add(fun)
	p.P.Legend.Add(legend, fun)
}

func (p *Plot) ToPlotterPoint(X, y *mat.Dense) plotter.XYs {
	r, _ := X.Dims()
	pts := make(plotter.XYs, r)
	for i := range pts {
		pts[i].X = X.At(i, 1)
		pts[i].Y = y.At(i, 0)
	}
	return pts
}
