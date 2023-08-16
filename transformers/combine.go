package transformers

import "strings"

type CombineTransformer struct {
	Columns []string
	With    string
	To      string
}

func (t *CombineTransformer) Apply(row map[string]interface{}) {
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
