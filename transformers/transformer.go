package transformers

import parser "csvshift/gen"

type Transformer interface {
	Apply(row map[string]interface{})
}

type MultipleColumnTransformerFactory interface {
	IsMatch(modifier parser.IMultipleColumnTransformationContext) bool
	Create(columns []string, modifier parser.IMultipleColumnTransformationContext) Transformer
}

type SingleColumnTransformerFactory interface {
	IsMatch(modifier parser.ISingleColumnTransformationContext) bool
	Create(column string, modifier parser.ISingleColumnTransformationContext) Transformer
}

type RowTransformer struct {
	Transformers []Transformer
}

func ExtractStringContent(str string) string {
	return str[1 : len(str)-1]
}
