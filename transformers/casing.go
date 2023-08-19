package transformers

import (
	parser "csvshift/gen"
	"strings"
)

type singleColumnLowerTransformer struct {
	Column string
}

type SingleColumnLowerTransformerFactory struct{}

type singleColumnUpperTransformer struct {
	Column string
}

type SingleColumnUpperTransformerFactory struct{}

type multipleColumnLowerTransformer struct {
	Columns []string
}

type MultipleColumnLowerTransformerFactory struct{}

type multipleColumnUpperTransformer struct {
	Columns []string
}

type MultipleColumnUpperTransformerFactory struct{}

func (t *multipleColumnLowerTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ToLower(val)
		}
	}
}

func (t *MultipleColumnLowerTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.LOWER() != nil
}

func (t *MultipleColumnLowerTransformerFactory) Create(columns []string, _ parser.IMultipleColumnTransformationContext) Transformer {
	return &multipleColumnLowerTransformer{
		Columns: columns,
	}
}

func (t *multipleColumnUpperTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ToUpper(val)
		}
	}
}

func (t *MultipleColumnUpperTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.UPPER() != nil
}

func (t *MultipleColumnUpperTransformerFactory) Create(columns []string, _ parser.IMultipleColumnTransformationContext) Transformer {
	return &multipleColumnUpperTransformer{
		Columns: columns,
	}
}

func (t *singleColumnLowerTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		row[t.Column] = strings.ToLower(val)
	}
}

func (t *SingleColumnLowerTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.LOWER() != nil
}

func (t *SingleColumnLowerTransformerFactory) Create(column string, _ parser.ISingleColumnTransformationContext) Transformer {
	return &singleColumnLowerTransformer{
		Column: column,
	}
}

func (t *singleColumnUpperTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		row[t.Column] = strings.ToUpper(val)
	}
}

func (t *SingleColumnUpperTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.UPPER() != nil
}

func (t *SingleColumnUpperTransformerFactory) Create(column string, _ parser.ISingleColumnTransformationContext) Transformer {
	return &singleColumnUpperTransformer{
		Column: column,
	}
}
