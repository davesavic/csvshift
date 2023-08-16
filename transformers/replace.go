package transformers

import "strings"

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
