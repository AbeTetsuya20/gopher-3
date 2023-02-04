package main

import (
	"fmt"
	"myanalyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	fmt.Println("Start")
	unitchecker.Main(myanalyzer.Analyzer)
}
