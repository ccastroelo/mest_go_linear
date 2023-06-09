package matrizes

import "fmt"

func Media(elem []float64) float64 {
	var total float64 = 0
	for _, e := range elem {
		total = total + e
	}
	return total / float64(len(elem))
}

func ImprimeMatriz(A [][]float64) {
	for i := range A {
		for j := range A[i] {
			fmt.Printf("%12.4f ", A[i][j])
		}
		fmt.Println()
	}
}
