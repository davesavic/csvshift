package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestMultipleColumnJoinTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name    string
		columns []string
		with    string
		to      string
		input   map[string]interface{}
		output  map[string]interface{}
	}{
		{
			name:    "join all columns with comma",
			columns: []string{"col1", "col2", "col3"},
			with:    ",",
			to:      "result",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "value3",
			},
			output: map[string]interface{}{
				"col1":   "value1",
				"col2":   "value2",
				"col3":   "value3",
				"result": "value1,value2,value3",
			},
		},
		{
			name:    "join two columns with space",
			columns: []string{"col1", "col2"},
			with:    " ",
			to:      "result",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1":   "value1",
				"col2":   "value2",
				"result": "value1 value2",
			},
		},
		{
			name:    "join no columns",
			columns: []string{},
			with:    "-",
			to:      "result",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1":   "value1",
				"col2":   "value2",
				"result": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.MultipleColumnJoinTransformer{
				Columns: tc.columns,
				With:    tc.with,
				To:      tc.to,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
