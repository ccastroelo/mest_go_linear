package avaliacao_validacao

import (
	"linear/matrizes"
	"math"
)

func AvaliaValidaModelo(coefs []float64, varIndep [][]float64, varDep []float64, dims []int) (float64, float64, float64, float64, []float64) {
	aMAE := 0.0
	aMSE := 0.0
	mediaY := matrizes.Media(varDep)
	SomaDifMedia := 0.0
	//	RA := 0.0
	estimadoY := make([]float64, len(varDep))
	for i, linha := range varIndep {
		yEst := 0.0
		for j, dim := range dims {
			if j == 0 {
				yEst = yEst + coefs[j]
			}
			yEst = yEst + coefs[j+1]*linha[dim]

		}
		estimadoY[i] = yEst

		aMAE = aMAE + math.Abs(varDep[i]-yEst)
		aMSE = aMSE + math.Pow((varDep[i]-yEst), 2)
		SomaDifMedia = SomaDifMedia + math.Pow((varDep[i]-mediaY), 2)
	}
	R2 := 1 - (aMSE / SomaDifMedia)
	MSE := aMSE / float64(len(varDep))
	RMSE := math.Sqrt(MSE)
	MAE := aMAE / float64(len(varDep))

	return MAE, MSE, RMSE, R2, estimadoY
}
