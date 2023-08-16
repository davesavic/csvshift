package transformers

import "strings"

type SplitTransformer struct {
	Column      string
	Separator   string
	IntoColumns []string
}

func (t *SplitTransformer) Apply(row map[string]interface{}) {
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
