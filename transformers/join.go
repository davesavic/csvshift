package transformers

import "strings"

type JoinTransformer struct {
	Columns []string
	With    string
	To      string
}

func (t *JoinTransformer) Apply(row map[string]interface{}) {
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
