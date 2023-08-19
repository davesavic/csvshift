package transformers

import (
	parser "csvshift/gen"
	"strings"
)

type singleColumnSplitTransformer struct {
	Column      string
	Separator   string
	IntoColumns []string
}

type SingleColumnSplitTransformerFactory struct{}

func (t *singleColumnSplitTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		parts := strings.Split(val, t.Separator)
		for i, column := range t.IntoColumns {
			if i < len(parts) {
				row[column] = parts[i]
			} else {
				row[column] = ""
			}
		}
	}
}

func (t *SingleColumnSplitTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.SPLIT() != nil
}

func (t *SingleColumnSplitTransformerFactory) Create(column string, modifier parser.ISingleColumnTransformationContext) Transformer {
	intoCols := modifier.Columns().AllIDENTIFIER()
	var intoColNames []string
	for _, intoCol := range intoCols {
		intoColNames = append(intoColNames, intoCol.GetText())
	}

	return &singleColumnSplitTransformer{
		Column:      column,
		Separator:   ExtractStringContent(modifier.STRING(0).GetText()),
		IntoColumns: intoColNames,
	}
}
