package listeners

import (
	parser "csvshift/gen"
	"csvshift/transformers"
)

type ParsedData struct {
	InputColumns    []string
	RowTransformers []transformers.RowTransformer
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

func (s *CsvShiftListener) AddRowTransformer(transformer transformers.RowTransformer) {
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
	var trs []transformers.Transformer

	for _, modifier := range modifiers {
		switch true {
		case modifier.REPLACE() != nil:
			trs = append(trs, &transformers.ReplaceTransformer{
				Columns: []string{column},
				From:    extractStringContent(modifier.STRING(0).GetText()),
				To:      extractStringContent(modifier.STRING(1).GetText()),
			})
		case modifier.TRIM() != nil:
			trs = append(trs, &transformers.TrimTransformer{
				Columns: []string{column},
			})
		}
	}

	if len(trs) == 0 {
		return
	}

	s.AddRowTransformer(transformers.RowTransformer{
		Transformers: trs,
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
	var trs []transformers.Transformer

	for _, modifier := range modifiers {
		switch true {
		case modifier.TRIM() != nil:
			trs = append(trs, &transformers.TrimTransformer{
				Columns: columns,
			})
		case modifier.REPLACE() != nil:
			trs = append(trs, &transformers.ReplaceTransformer{
				Columns: columns,
				From:    extractStringContent(modifier.STRING(0).GetText()),
				To:      extractStringContent(modifier.STRING(1).GetText()),
			})
		case modifier.COMBINE() != nil:
			trs = append(trs, &transformers.CombineTransformer{
				Columns: columns,
				With:    extractStringContent(modifier.STRING(0).GetText()),
				To:      modifier.IDENTIFIER().GetText(),
			})
		}
	}

	if len(trs) == 0 {
		return
	}

	s.AddRowTransformer(transformers.RowTransformer{
		Transformers: trs,
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
