package application

import (
	parser "csvshift/gen"
	"csvshift/listeners"
	"encoding/csv"
	"github.com/antlr4-go/antlr/v4"
	"io"
	"log"
	"os"
)

func Run(source, destination, path string) {
	script := parseScript(path)

	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatalf("Error opening source file: %v", err)
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			log.Fatalf("Error closing source file: %v", err)
		}
	}(sourceFile)

	destinationFile, err := os.Create(destination)
	if err != nil {
		log.Fatalf("Error creating destination file: %v", err)
	}
	defer func(destinationFile *os.File) {
		err := destinationFile.Close()
		if err != nil {
			log.Fatalf("Error closing destination file: %v", err)
		}
	}(destinationFile)

	csvReader := csv.NewReader(sourceFile)
	csvWriter := csv.NewWriter(destinationFile)
	_ = csvWriter.Write(script.OutputColumns)

	headers, _ := csvReader.Read()
	headerSet := make(map[string]bool)
	for _, h := range headers {
		headerSet[h] = true
	}

	for _, header := range script.InputColumns {
		if _, exists := headerSet[header]; !exists {
			log.Fatalf("Required header '%s' not found in input file", header)
		}
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Error reading input file: %v", err)
		}

		row := make(map[string]interface{})
		for i, header := range headers {
			row[header] = record[i]
		}

		for _, transformer := range script.RowTransformers {
			for _, t := range transformer.Transformers {
				t.Apply(row)
			}
		}

		var output []string
		for _, header := range script.OutputColumns {
			rVal, ok := row[header].(string)
			if !ok {
				log.Fatalf("Error converting output column '%s' value to string: %v", header, err)
			}

			output = append(output, rVal)
		}

		err = csvWriter.Write(output)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	}

	csvWriter.Flush()
}

func parseScript(path string) *listeners.ParsedData {
	input, err := antlr.NewFileStream(path)
	if err != nil {
		log.Fatalf("Error opening script file: %v", err)
	}

	lexer := parser.NewCsvShiftGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCsvShiftGrammarParser(stream)
	p.BuildParseTrees = true
	tree := p.CsvTransform()
	listener := listeners.NewCsvShiftListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.Data
}
