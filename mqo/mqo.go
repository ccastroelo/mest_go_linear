package mqo

func CalcCoef(x []float64, y []float64) []float64 {
	mediaX := media(x)
	mediaY := media(y)
	sXY := 0.0
	sXX := 0.0

	var result []float64
	for i := 0; i < len(x); i++ {
		sXY = sXY + (x[i]-mediaX)*(y[i]-mediaY)
		sXX = sXX + (x[i]-mediaX)*(x[i]-mediaX)
	}
	b1 := sXY / sXX
	b0 := mediaY - b1*mediaX
	result = append(result, b0, b1)
	return result

}

func media(elem []float64) float64 {
	var total float64 = 0
	for _, e := range elem {
		total = total + e
	}
	return total / float64(len(elem))
}
