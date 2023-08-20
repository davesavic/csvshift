package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestSingleColumnRegexExtractTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name        string
		column      string
		pattern     string
		intoColumns []string
		input       map[string]interface{}
		output      map[string]interface{}
	}{
		{
			name:        "extract key and value from column by colon into two columns",
			column:      "col1",
			pattern:     "(\\w+):(\\w+)",
			intoColumns: []string{"col2", "col3"},
			input: map[string]interface{}{
				"col1": "key1:value1",
			},
			output: map[string]interface{}{
				"col1": "key1:value1",
				"col2": "key1",
				"col3": "value1",
			},
		},
		{
			name:        "extract key and value from column by colon into two columns, with optional value",
			column:      "col1",
			pattern:     "(\\w+):?(\\w*)",
			intoColumns: []string{"col2", "col3"},
			input: map[string]interface{}{
				"col1": "key1:",
			},
			output: map[string]interface{}{
				"col1": "key1:",
				"col2": "key1",
				"col3": "",
			},
		},
		{
			name:        "non-matching pattern does not extract anything",
			column:      "col1",
			pattern:     "(\\w+):(\\w+)",
			intoColumns: []string{"col2", "col3"},
			input: map[string]interface{}{
				"col1": "key1value1",
			},
			output: map[string]interface{}{
				"col1": "key1value1",
				"col2": "",
				"col3": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnRegexExtractTransformer{
				Column:      tc.column,
				Pattern:     tc.pattern,
				IntoColumns: tc.intoColumns,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
