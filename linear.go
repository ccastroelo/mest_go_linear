package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	loadfile "linear/loadfile"
	"linear/mqo"
	"linear/showgraph"
)

func main() {
	var hVarDep string
	var hVarIndep []string
	var varDepTest, varDep []float64
	var varIndep, varIndepTest [][]float64
	sair := false
	for !sair {
		fmt.Println("Opções: ")
		fmt.Println("0 > Sair ")
		fmt.Println("1 > Carregar arquivo e particionar treino e teste ")
		fmt.Println("2 > Gerar gráficos de dispersão dos dados de treino")
		fmt.Println("3 > Calcula o modelo - Linear Simples")

		opcao := bufio.NewScanner(os.Stdin)
		opcao.Scan()

		switch {
		case opcao.Text() == "0":
			sair = true
		case opcao.Text() == "1":
			fmt.Println(">> Carregar Arquivos")
			fmt.Println("Qual o percentual utilizado para teste? Formato: 0.00")
			opcao := bufio.NewScanner(os.Stdin)
			opcao.Scan()
			testPerc, err := strconv.ParseFloat(opcao.Text(), 64)
			_hVarDep, _varDep, _varDepTest, _hVarIndep, _varIndep, _varIndepTest, err := loadfile.LoadInput("./assets/WHR2023.csv", testPerc)
			if err != nil {
				fmt.Println(err)
			}
			hVarDep = _hVarDep
			varDep = _varDep
			varDepTest = _varDepTest
			hVarIndep = _hVarIndep
			varIndep = _varIndep
			varIndepTest = _varIndepTest

			_ = hVarDep
			_ = varDepTest
			_ = varIndepTest

			fmt.Println("Arquivos particionados")
			fmt.Println("Arquivo de treino: ", len(varIndep), " linhas")
			fmt.Println("Arquivo de teste: ", len(varIndepTest), "linhas")
			//fmt.Println(hVarIndep[0])
			//fmt.Println(varIndep[0])
		case opcao.Text() == "2":
			fmt.Println("Gerando gráficos de dispersão dos dados de treino")
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 0, false)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 1, false)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 2, false)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 3, false)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 4, false)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, 5, false)
		case opcao.Text() == "3":
			fmt.Println(">> Calcular regressao linear simples")
			fmt.Println("Qual a dimensão (informe o número):")
			for i, h := range hVarIndep {
				fmt.Println(i, ")", h)
			}
			fmt.Println(">")
			opcao := bufio.NewScanner(os.Stdin)
			opcao.Scan()
			dim, err := strconv.Atoi(opcao.Text())
			if err != nil {
				fmt.Println(err)
			}
			v := mqo.CalcCoef(varIndep[dim], varDep)
			fmt.Println("*******")
			fmt.Println(v)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, dim, true)
		}
	}

}
