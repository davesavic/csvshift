package transformers

import (
	parser "csvshift/gen"
	"strings"
)

type MultipleColumnJoinTransformerFactory struct{}

type multipleColumnJoinTransformer struct {
	Columns []string
	With    string
	To      string
}

func (t *multipleColumnJoinTransformer) Apply(row map[string]interface{}) {
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

func (t *MultipleColumnJoinTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.JOIN() != nil
}

func (t *MultipleColumnJoinTransformerFactory) Create(columns []string, modifier parser.IMultipleColumnTransformationContext) Transformer {
	return &multipleColumnJoinTransformer{
		Columns: columns,
		With:    ExtractStringContent(modifier.STRING(0).GetText()),
		To:      modifier.IDENTIFIER().GetText(),
	}
}
