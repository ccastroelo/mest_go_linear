package showgraph

import (
	"log"
	"os"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}
func Showgraph(varIndep [][]float64, varDep []float64, hVarIndep []string, varIndepIdx int, linha bool) {
	size := len(varDep)
	data := xy{
		x: make([]float64, size),
		y: make([]float64, size),
	}
	for i := 0; i < size; i++ {
		data.y[i] = varDep[i]
		data.x[i] = varIndep[i][varIndepIdx]
	}
	var line *plotter.Function
	if linha {

		b, a := stat.LinearRegression(data.x, data.y, nil, false)
		log.Printf("%v*x+%v", a, b)
		_line := plotter.NewFunction(func(x float64) float64 { return a*x + b })
		line = _line
	}

	p := plot.New()

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Panic(err)
	}
	if linha {

		p.Add(scatter, line)
	} else {
		p.Add(scatter)
	}

	w, err := p.WriterTo(300, 300, "png")
	if err != nil {
		log.Panic(err)
	}

	g, err := os.Create(hVarIndep[varIndepIdx] + "_graph.png")
	if err != nil {
		log.Fatalf("Não foi possível criar o arquivo de gráfico")
	}
	defer g.Close()

	_, err = w.WriteTo(g)
	if err != nil {
		log.Panic(err)
	}

}
