package main

import (
	"fmt"

	loadfile "linear/loadfile"
	showgraph "linear/showgraph"
)

func main() {
	testSlicePerc := 0.25

	hVarDep, varDep, varDepTest, hVarIndep, varIndep, varIndepTest, err := loadfile.LoadInput("./assets/consumo_cerveja.csv", testSlicePerc)
	if err != nil {
		fmt.Println(err)
	}
	_ = hVarDep
	_ = varDepTest
	_ = varIndepTest

	//	showgraph.Showgraph(varIndep, varDep, hVarIndep, 0)
	//	showgraph.Showgraph(varIndep, varDep, hVarIndep, 1)
	showgraph.Showgraph(varIndep, varDep, hVarIndep, 2)
	//	showgraph.Showgraph(varIndep, varDep, hVarIndep, 3)
	//	showgraph.Showgraph(varIndep, varDep, hVarIndep, 4)

}
