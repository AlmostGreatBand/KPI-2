package main

import (
	"flag"
	lab2 "github.com/AlmostGreatBand/KPI-2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	defaultValue = ""
	inputExpression = flag.String("e", defaultValue, "Expression to compute")
	fileExpression = flag.String("f", defaultValue, "File with expression to compute")
	outputFile = flag.String("o", defaultValue, "Save expression to file")
)

func main() {
	flag.Parse()
	expr := *inputExpression
	fromFileName := *fileExpression
	toFileName := *outputFile

	var from io.Reader
	if expr != defaultValue {
		from = strings.NewReader(expr)
	} else if fromFileName != defaultValue {
		data, err := ioutil.ReadFile(fromFileName)
		if err != nil {
			os.Stderr.WriteString("provided file hasn't been found")
			return
		}
		from = strings.NewReader(string(data))
	} else {
		os.Stderr.WriteString("can't find expression")
		return
	}

	var (
		to io.Writer
		err error
	)
	if toFileName != defaultValue {
		if to, err = os.Create(toFileName); err != nil {
			os.Stderr.WriteString("an error occupied while creating file")
		}
	} else {
		to = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input: from,
		Output: to,
	}
	if err := handler.Compute(); err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
