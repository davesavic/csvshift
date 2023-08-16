package transformers

import "strings"

type LowerTransformer struct {
	Columns []string
}

func (t *LowerTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ToLower(val)
		}
	}
}

type UpperTransformer struct {
	Columns []string
}

func (t *UpperTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			row[column] = strings.ToUpper(val)
		}
	}
}
