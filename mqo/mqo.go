package mqo

import (
	"linear/matrizes"
)

func MqoBasico(varIndep [][]float64, varDep []float64, dim int) []float64 {

	/* ***********
	* Formula:
	* B1 = Somatório((xi - mediaX) * (yi - mediaY)) / Somatório de (xi - mediax)**2
	* B0 = mediaY - B1*mediaX
	 */

	// gera um vetor com a coluna da dimensao escolhida para facilitar
	x := make([]float64, len(varIndep))
	for i := 0; i < len(varIndep); i++ {
		x[i] = varIndep[i][dim]
	}

	// calcula médias
	mediaX := matrizes.Media(x)
	mediaY := matrizes.Media(varDep)

	// calcula somatorio de (xi - mediaX) * (yi - mediaY)
	// sobre   somatório de (xi - mediax)**2
	// para b1
	sXY := 0.0
	sXX := 0.0
	for i := 0; i < len(x); i++ {
		sXY = sXY + (x[i]-mediaX)*(varDep[i]-mediaY)
		sXX = sXX + (x[i]-mediaX)*(x[i]-mediaX)
	}
	b1 := sXY / sXX
	b0 := mediaY - b1*mediaX
	var result []float64
	result = append(result, b0, b1)
	return result

}

func MqoMatriz(varIndep [][]float64, varDep []float64, dims []int) []float64 {

	/*************
	* Formula:
	* Bi = (XT*X)**-1 * XT * Y  => resulta em uma matriz de dimensão Nx1 de coeficientes (B0, B1, ...)
	 */

	// Tranforma o vetor de dimensão N em matriz de dimensão Nx1
	my := make([][]float64, len(varDep))
	for i := 0; i < len(varDep); i++ {
		my[i] = make([]float64, 1)
		my[i][0] = varDep[i]
	}

	// adiciona 1 à matriz de variáveis independentes (para multiplicar o interceptor)
	mx := make([][]float64, len(varIndep))
	for i := 0; i < len(varIndep); i++ {
		mx[i] = make([]float64, len(dims)+1)
		for j := 0; j < len(dims); j++ {
			if j == 0 {
				mx[i][j] = 1
			}
			mx[i][j+1] = varIndep[i][dims[j]]

		}
	}

	//	(XT*X)**-1 * XT * Y

	m1 := matrizTransposta(mx)
	m2 := multiplicaMatriz(m1, mx)
	m3 := matrizInversa(m2)
	m4 := multiplicaMatriz(m3, m1)
	m5 := multiplicaMatriz(m4, my)

	// é feita a transposta para facilitar transformar em um vetor de coeficientes B0, B1, ... BN
	m6 := matrizTransposta(m5)

	return m6[0]
}

func matrizTransposta(varIndep [][]float64) [][]float64 {

	// cria matriz transposta
	mT := make([][]float64, len(varIndep[0]))

	// inverte linha por coluna
	for i := 0; i < len(varIndep[0]); i++ {
		mT[i] = make([]float64, len(varIndep))
		for j := 0; j < len(varIndep); j++ {
			mT[i][j] = varIndep[j][i]
		}
	}
	return mT
}

func multiplicaMatriz(a [][]float64, b [][]float64) [][]float64 {
	linhasA := len(a)
	colsA := len(a[0])
	linhasB := len(b)
	colsB := len(b[0])

	if colsA != linhasB {
		panic("Dimensoes incompativeis")
	}

	// cria a matriz resultante
	mR := make([][]float64, linhasA)
	for i := range mR {
		mR[i] = make([]float64, colsB)
	}

	// multiplica coluna k por linha k e soma
	for i := 0; i < linhasA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				mR[i][j] = mR[i][j] + a[i][k]*b[k][j]
			}
		}
	}

	return mR
}

func matrizInversa(m [][]float64) [][]float64 {

	mAux := clone(m)
	c := len(mAux)

	// cria e preenche identidade
	mI := make([][]float64, c)
	for i := range mI {
		mI[i] = make([]float64, c)
		mI[i][i] = 1
	}

	// calcula a identidade por gauss-charles
	for k := 0; k < c; k++ { // linhas da matriz
		for i := 0; i < c; i++ {
			if i != k {
				elem := mAux[i][k] / mAux[k][k]
				for j := 0; j < c; j++ {
					mAux[i][j] = mAux[i][j] - elem*mAux[k][j]
					mI[i][j] = mI[i][j] - elem*mI[k][j]
				}
			}
		}

		aux := 1 / mAux[k][k]
		for j := 0; j < c; j++ {
			mAux[k][j] = mAux[k][j] * aux
			mI[k][j] = mI[k][j] * aux
		}
	}

	return mI
}

// matriz/slices são passados por referencia,
// então tem q clonar (deep copy) para nao baguncar
func clone(m [][]float64) [][]float64 {
	x := make([][]float64, len(m))
	for i := range m {
		x[i] = make([]float64, len(m[0]))
		for j := range m[i] {
			x[i][j] = m[i][j]
		}
	}
	return x
}
