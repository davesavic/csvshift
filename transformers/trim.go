package transformers

import (
	parser "csvshift/gen"
	"strings"
)

type MultipleColumnTrimTransformerFactory struct{}

type MultipleColumnTrimTransformer struct {
	Columns []string
}

type SingleColumnTrimTransformerFactory struct{}

type SingleColumnTrimTransformer struct {
	Column string
}

func (t *MultipleColumnTrimTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.TrimSpace(val)
		}
	}
}

func (t *MultipleColumnTrimTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.TRIM() != nil
}

func (t *MultipleColumnTrimTransformerFactory) Create(columns []string, _ parser.IMultipleColumnTransformationContext) Transformer {
	return &MultipleColumnTrimTransformer{
		Columns: columns,
	}
}

func (t *SingleColumnTrimTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		row[t.Column] = strings.TrimSpace(val)
	}
}

func (t *SingleColumnTrimTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.TRIM() != nil
}

func (t *SingleColumnTrimTransformerFactory) Create(column string, _ parser.ISingleColumnTransformationContext) Transformer {
	return &SingleColumnTrimTransformer{
		Column: column,
	}
}
