package main

import (
	"github.com/Mitu217/go-analysis-practice/passes/practice"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(practice.Analyzer) }
