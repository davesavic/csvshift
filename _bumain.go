package _bumain

import (
	parser "csvshift/gen"
	"csvshift/listeners"
	"encoding/csv"
	"fmt"
	"github.com/antlr4-go/antlr"
	"io"
	"log"
	"os"
)

func _bumain() {
	input, _ := antlr.NewFileStream("script.csvshift")

	lexer := parser.NewCsvShiftGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCsvShiftGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CsvTransform()
	listener := listeners.NewCsvShiftListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Printf("Input Columns: %v\n", listener.Data.InputColumns)
	fmt.Printf("Output Columns: %v\n", listener.Data.OutputColumns)
	fmt.Printf("Row Transformers: %v\n", len(listener.Data.RowTransformers))

	file, err := os.Open("large.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	dstFile, err := os.Create("destination.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dstFile)

	reader := csv.NewReader(file)
	writer := csv.NewWriter(dstFile)
	_ = writer.Write(listener.Data.OutputColumns)

	headers, _ := reader.Read()
	headerSet := make(map[string]bool)
	for _, h := range headers {
		headerSet[h] = true
	}

	for _, header := range listener.Data.InputColumns {
		if _, exists := headerSet[header]; !exists {
			log.Fatalf("Required header '%s' not found in input file", header)
		}
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		row := make(map[string]interface{})
		for i, header := range headers {
			row[header] = record[i]
		}

		for _, transformer := range listener.Data.RowTransformers {
			for _, t := range transformer.Transformers {
				fmt.Printf("In: %v\n", row)
				t.Apply(row)
				fmt.Printf("Out: %v\n", row)
			}
		}

		var output []string
		for _, header := range listener.Data.OutputColumns {
			output = append(output, row[header].(string))
		}

		_ = writer.Write(output)
	}

	writer.Flush()
}
