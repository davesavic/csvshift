package transformers

import "strings"

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
