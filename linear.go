package main

import (
	"bufio"
	"fmt"
	"linear/avaliacao_validacao"
	"linear/carrega_csv"
	"linear/gradiente"

	"linear/mqo"
	"linear/showgraph"
	"os"
	"strconv"
)

func main() {
	var hVarDep string
	var hVarIndep []string
	var varDepTest, varDep, coefs []float64
	var varIndep, varIndepTest [][]float64
	var dims []int
	sair := false

	// debug
	/*
		_hVarDep, _varDep, _varDepTest, _hVarIndep, _varIndep, _varIndepTest, err := carrega_csv.CarregaCSV("./assets/teste.csv", 0)
		if err != nil {
			fmt.Println(err)
		}
		hVarDep = _hVarDep
		varDep = _varDep
		varDepTest = _varDepTest
		hVarIndep = _hVarIndep
		varIndep = _varIndep
		varIndepTest = _varIndepTest
		dims = []int{0, 1}
		passos := 1000
		taxa := 0.00003
		_coefs := gradiente.CalcGradient(varIndep, varDep, dims, passos, taxa)
		coefs = _coefs
		fmt.Println("")
		imprimeFormula(hVarDep, hVarIndep, dims, coefs)
		fmt.Println("")
		sair = true

	*/
	// fim debug

	for !sair {
		fmt.Println("*******************************************************************************************")
		fmt.Println("Opções: ")
		fmt.Println("0 > Sair ")
		fmt.Println("1 > Carregar arquivo e particionar treino e teste ")
		fmt.Println("2 > Gerar gráficos de dispersão dos dados de treino")
		fmt.Println("3 > Calcula o modelo - Linear Simples")
		fmt.Println("4 > Calcula o modelo - Linear Multiplo")
		//		fmt.Println("5 > Calcula o modelo - Linear Multiplo com Gradiente Descendente")
		fmt.Println("5 > Avaliação do modelo")
		fmt.Println("6 > Validação do modelo (+ gráfico se linear simples)")

		opcao := bufio.NewScanner(os.Stdin)
		opcao.Scan()
		fmt.Println("")
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println("")

		switch {
		case opcao.Text() == "0":
			sair = true
		case opcao.Text() == "1":
			fmt.Println("")
			fmt.Println(">> Carregar Arquivo")
			fmt.Println("")
			fmt.Println("Qual o nome do arquivo .csv?")
			arq := bufio.NewScanner(os.Stdin)
			arq.Scan()
			fmt.Println("")
			fmt.Println("Arquivo de teste separado? (s/n)")
			opcaoTeste := bufio.NewScanner(os.Stdin)
			opcaoTeste.Scan()
			if opcaoTeste.Text() == "S" || opcaoTeste.Text() == "s" {
				fmt.Println("Qual o nome do arquivo de teste .csv?")
				arqTeste := bufio.NewScanner(os.Stdin)
				arqTeste.Scan()
				_hVarDep, _varDep, _, _hVarIndep, _varIndep, _, err := carrega_csv.CarregaCSV("./assets/"+arq.Text()+".csv", 0)
				if err != nil {
					fmt.Println(err)
				}
				_, _varDepTest, _, _, _varIndepTest, _, err := carrega_csv.CarregaCSV("./assets/"+arqTeste.Text()+".csv", 0)
				if err != nil {
					fmt.Println(err)
				}
				hVarDep = _hVarDep
				varDep = _varDep
				varDepTest = _varDepTest
				hVarIndep = _hVarIndep
				varIndep = _varIndep
				varIndepTest = _varIndepTest
			} else {
				fmt.Println("")
				fmt.Println("Qual o percentual utilizado para teste? Formato: 99")
				opcao := bufio.NewScanner(os.Stdin)
				opcao.Scan()
				testPerc, err := strconv.ParseFloat(opcao.Text(), 64)

				_hVarDep, _varDep, _varDepTest, _hVarIndep, _varIndep, _varIndepTest, err := carrega_csv.CarregaCSV("./assets/"+arq.Text()+".csv", testPerc/100)
				if err != nil {
					fmt.Println(err)
				}
				hVarDep = _hVarDep
				varDep = _varDep
				varDepTest = _varDepTest
				hVarIndep = _hVarIndep
				varIndep = _varIndep
				varIndepTest = _varIndepTest

			}
			coefs = []float64{}
			fmt.Println("")
			fmt.Println("Arquivos particionados")
			fmt.Println("Arquivo de treino: ", len(varIndep), " linhas")
			fmt.Println("Arquivo de teste: ", len(varIndepTest), "linhas")
			fmt.Println("Qtd Variáveis independentes: ", len(varIndep[0]), " colunas")
			listaDimensoes(hVarIndep)
		case opcao.Text() == "2":
			if len(hVarIndep) == 0 {
				fmt.Println("")
				fmt.Println(">> Arquivo não carregado")
				break
			}
			fmt.Println("Gerando gráficos de dispersão dos dados de treino")
			for i := 0; i < len(varIndep[0]); i++ {
				showgraph.Showgraph(varIndep, varDep, hVarIndep, i, 0, 0, false)
			}
			fmt.Println("Gráficos gerados")
			fmt.Println("")
		case opcao.Text() == "3":
			if len(hVarIndep) == 0 {
				fmt.Println("")
				fmt.Println(">> Arquivo não carregado")
				break
			}
			fmt.Println("")
			fmt.Println(">> Calcular regressao linear simples")
			fmt.Println("")
			fmt.Println("Qual a dimensão (informe o número):")
			listaDimensoes(hVarIndep)
			fmt.Println(">")
			opcao := bufio.NewScanner(os.Stdin)
			opcao.Scan()
			dim, err := strconv.Atoi(opcao.Text())
			if err != nil {
				fmt.Println(err)
			}
			_coefs := mqo.MqoBasico(varIndep, varDep, dim)
			coefs = _coefs
			dims = []int{dim}
			imprimeFormula(hVarDep, hVarIndep, dims, coefs)
			showgraph.Showgraph(varIndep, varDep, hVarIndep, dim, coefs[0], coefs[1], false)
			fmt.Println("")
			fmt.Println("Gráfico atualizado")
			fmt.Println("")
		case opcao.Text() == "4":
			if len(hVarIndep) == 0 {
				fmt.Println("")
				fmt.Println(">> Arquivo não carregado")
				break
			}
			fmt.Println("")
			fmt.Println(">> Calcula regressão linear multipla")
			fmt.Println("")
			fmt.Println("Informe as dimensões (informe o número <enter> para finalizar):")
			listaDimensoes(hVarIndep)
			fmt.Println(">")
			_dims := make([]int, 0, len(hVarIndep))
			sair := false
			for !sair {
				opcao := bufio.NewScanner(os.Stdin)
				opcao.Scan()
				if opcao.Text() == "" {
					sair = true
					continue
				}
				dim, err := strconv.Atoi(opcao.Text())
				if err != nil {
					fmt.Println(err)
				}
				_dims = append(_dims, dim)
			}
			dims = _dims
			_coefs := mqo.MqoMatriz(varIndep, varDep, dims)
			coefs = _coefs
			fmt.Println("")
			imprimeFormula(hVarDep, hVarIndep, dims, coefs)
			fmt.Println("")
		case opcao.Text() == "desativado do menu":
			fmt.Println("")
			fmt.Println(">> Calcula regressão linear multipla com gradiente descendente")
			fmt.Println("")
			fmt.Println("Informe as dimensões (informe o número <enter> para finalizar):")
			listaDimensoes(hVarIndep)
			fmt.Println(">")
			_dims := make([]int, 0, len(hVarIndep))
			sair := false
			for !sair {
				opcao := bufio.NewScanner(os.Stdin)
				opcao.Scan()
				if opcao.Text() == "" {
					sair = true
					continue
				}
				dim, err := strconv.Atoi(opcao.Text())
				if err != nil {
					fmt.Println(err)
				}
				_dims = append(_dims, dim)
			}
			dims = _dims
			fmt.Println("")
			fmt.Println("Qual a quantidade de iteracoes?")
			_passos := bufio.NewScanner(os.Stdin)
			_passos.Scan()
			passos, err := strconv.Atoi(_passos.Text())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("")
			fmt.Println("Qual a taxa de aprendizagem?")
			_taxa := bufio.NewScanner(os.Stdin)
			_taxa.Scan()
			taxa, err := strconv.ParseFloat(_taxa.Text(), 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_coefs := gradiente.CalcGradient(varIndep, varDep, dims, passos, taxa)
			coefs = _coefs
			fmt.Println("")
			imprimeFormula(hVarDep, hVarIndep, dims, coefs)
			fmt.Println("")
		case opcao.Text() == "5":
			fmt.Println("")
			fmt.Println(">> Avaliação do modelo")
			if len(coefs) == 0 {
				fmt.Println("")
				fmt.Println(">> Modelo não calculado")
				break
			}
			MAE, MSE, RMSE, R2, _ := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndep, varDep, dims)
			fmt.Println("")
			fmt.Println("MAE = ", MAE)
			fmt.Println("MSE = ", MSE)
			fmt.Println("RMSE = ", RMSE)
			fmt.Println("R2 = ", R2)
			fmt.Println("")
		case opcao.Text() == "6":
			fmt.Println("")
			fmt.Println(">> Validação do modelo")
			if len(coefs) == 0 {
				fmt.Println("")
				fmt.Println(">> Modelo não calculado")
				break
			}
			MAE, MSE, RMSE, R2, ysEst := avaliacao_validacao.AvaliaValidaModelo(coefs, varIndepTest, varDepTest, dims)
			fmt.Println("MAE = ", MAE)
			fmt.Println("MSE = ", MSE)
			fmt.Println("RMSE = ", RMSE)
			fmt.Println("R2 = ", R2)
			fmt.Println("")
			fmt.Println("Observado => Estimado")
			for i, yEst := range ysEst {
				fmt.Println(varDepTest[i], " => ", yEst)
			}
			fmt.Println("")
			if len(dims) == 1 {
				fmt.Println(coefs[0])
				fmt.Println(coefs[1])
				imprimeFormula(hVarDep, hVarIndep, dims, coefs)
				showgraph.Showgraph(varIndepTest, varDepTest, hVarIndep, dims[0], coefs[0], coefs[1], true)
			}

		}
	}

}

func listaDimensoes(hVarIndep []string) {
	for i, h := range hVarIndep {
		fmt.Println(i, ")", h)
	}
}

func imprimeFormula(hVarDep string, hVarIndep []string, dims []int, coefs []float64) {
	formula := hVarDep + " = " + strconv.FormatFloat(coefs[0], 'f', -1, 32)
	for i := 1; i < len(coefs); i++ {
		if (coefs[i]) > 0 {
			formula = formula + " +"
		}
		formula = formula + " " + strconv.FormatFloat(coefs[i], 'f', -1, 32) + " * " + hVarIndep[dims[i-1]]
	}
	fmt.Println("")
	fmt.Println("Formula:")
	fmt.Println(formula)
	fmt.Println("")
}
