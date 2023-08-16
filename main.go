package main

import (
	parser "csvshift/gen"
	"encoding/csv"
	"fmt"
	"github.com/antlr4-go/antlr"
	"io"
	"log"
	"os"
	"strings"
)

type Transformer interface {
	Apply(row map[string]interface{})
}

type ReplaceTransformer struct {
	Columns []string
	From    string
	To      string
}

func (t *ReplaceTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ReplaceAll(val, t.From, t.To)
		}
	}
}

type TrimTransformer struct {
	Columns []string
}

func (t *TrimTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.TrimSpace(val)
		}
	}
}

type CombineTransformer struct {
	Columns []string
	With    string
	To      string
}

func (t *CombineTransformer) Apply(row map[string]interface{}) {
	values := make([]string, 0, len(t.Columns))
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if !ok {
			continue
		}

		values = append(values, val)
	}

	row[t.To] = strings.Join(values, t.With)
}

type RowTransformer struct {
	Transformers []Transformer
}

type ParsedData struct {
	InputColumns    []string
	RowTransformers []RowTransformer
	OutputColumns   []string
}

type CsvShiftListener struct {
	*parser.BaseCsvShiftGrammarListener
	Data *ParsedData
}

func NewCsvShiftListener() *CsvShiftListener {
	return &CsvShiftListener{
		Data: &ParsedData{},
	}
}

func (s *CsvShiftListener) AddRowTransformer(transformer RowTransformer) {
	s.Data.RowTransformers = append(s.Data.RowTransformers, transformer)
}

func (s *CsvShiftListener) ExitInputSection(ctx *parser.InputSectionContext) {
	columnsCtx := ctx.Columns()
	if columnsCtx == nil {
		return
	}

	for i := 0; i < len(columnsCtx.AllIDENTIFIER()); i++ {
		identifier := columnsCtx.IDENTIFIER(i)
		s.Data.InputColumns = append(s.Data.InputColumns, identifier.GetText())
	}
}

func (s *CsvShiftListener) ExitSingleColumnModifierSection(ctx *parser.SingleColumnModifierSectionContext) {
	column := ctx.IDENTIFIER().GetText()
	modifiers := ctx.AllSingleColumnTransformation()
	var transformers []Transformer

	for _, modifier := range modifiers {
		switch true {
		case modifier.REPLACE() != nil:
			transformers = append(transformers, &ReplaceTransformer{
				Columns: []string{column},
				From:    extractStringContent(modifier.STRING(0).GetText()),
				To:      extractStringContent(modifier.STRING(1).GetText()),
			})
		case modifier.TRIM() != nil:
			transformers = append(transformers, &TrimTransformer{
				Columns: []string{column},
			})
		}
	}

	if len(transformers) == 0 {
		return
	}

	s.AddRowTransformer(RowTransformer{
		Transformers: transformers,
	})
}

func (s *CsvShiftListener) ExitMultipleColumnModifierSection(ctx *parser.MultipleColumnModifierSectionContext) {
	columnsCtx := ctx.Columns()
	if columnsCtx == nil {
		return
	}

	var columns []string
	for i := 0; i < len(columnsCtx.AllIDENTIFIER()); i++ {
		columns = append(columns, columnsCtx.IDENTIFIER(i).GetText())
	}

	modifiers := ctx.AllMultipleColumnTransformation()
	var transformers []Transformer

	for _, modifier := range modifiers {
		switch true {
		case modifier.TRIM() != nil:
			transformers = append(transformers, &TrimTransformer{
				Columns: columns,
			})
		case modifier.REPLACE() != nil:
			transformers = append(transformers, &ReplaceTransformer{
				Columns: columns,
				From:    extractStringContent(modifier.STRING(0).GetText()),
				To:      extractStringContent(modifier.STRING(1).GetText()),
			})
		case modifier.COMBINE() != nil:
			transformers = append(transformers, &CombineTransformer{
				Columns: columns,
				With:    extractStringContent(modifier.STRING(0).GetText()),
				To:      modifier.IDENTIFIER().GetText(),
			})
		}
	}

	if len(transformers) == 0 {
		return
	}

	s.AddRowTransformer(RowTransformer{
		Transformers: transformers,
	})
}

func (s *CsvShiftListener) ExitOutputSection(ctx *parser.OutputSectionContext) {
	columnsCtx := ctx.Columns()
	if columnsCtx == nil {
		return
	}

	for i := 0; i < len(columnsCtx.AllIDENTIFIER()); i++ {
		identifier := columnsCtx.IDENTIFIER(i)
		s.Data.OutputColumns = append(s.Data.OutputColumns, identifier.GetText())
	}
}

func extractStringContent(str string) string {
	return str[1 : len(str)-1]
}

func main() {
	input, _ := antlr.NewFileStream("script.csvshift")

	lexer := parser.NewCsvShiftGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCsvShiftGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CsvTransform()
	listener := NewCsvShiftListener()
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
