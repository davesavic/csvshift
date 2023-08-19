package transformers

import (
	parser "csvshift/gen"
	"strings"
)

type multipleColumnReplaceTransformer struct {
	Columns []string
	From    string
	To      string
}

type MultipleColumnReplaceTransformerFactory struct{}

type singleColumnReplaceTransformer struct {
	Column string
	From   string
	To     string
}

type SingleColumnReplaceTransformerFactory struct{}

func (t *multipleColumnReplaceTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ReplaceAll(val, t.From, t.To)
		}
	}
}

func (t *MultipleColumnReplaceTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.REPLACE() != nil
}

func (t *MultipleColumnReplaceTransformerFactory) Create(columns []string, modifier parser.IMultipleColumnTransformationContext) Transformer {
	return &multipleColumnReplaceTransformer{
		Columns: columns,
		From:    ExtractStringContent(modifier.STRING(0).GetText()),
		To:      ExtractStringContent(modifier.STRING(1).GetText()),
	}
}

func (t *singleColumnReplaceTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		row[t.Column] = strings.ReplaceAll(val, t.From, t.To)
	}
}

func (t *SingleColumnReplaceTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.REPLACE() != nil
}

func (t *SingleColumnReplaceTransformerFactory) Create(column string, modifier parser.ISingleColumnTransformationContext) Transformer {
	return &singleColumnReplaceTransformer{
		Column: column,
		From:   ExtractStringContent(modifier.STRING(0).GetText()),
		To:     ExtractStringContent(modifier.STRING(1).GetText()),
	}
}
