package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestSingleColumnSplitTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name        string
		column      string
		separator   string
		intoColumns []string
		input       map[string]interface{}
		output      map[string]interface{}
	}{
		{
			name:        "split column by comma into two columns",
			column:      "col1",
			separator:   ",",
			intoColumns: []string{"col2", "col3"},
			input: map[string]interface{}{
				"col1": "value1,value2",
			},
			output: map[string]interface{}{
				"col1": "value1,value2",
				"col2": "value1",
				"col3": "value2",
			},
		},
		{
			name:        "split column by space into three columns, with missing value",
			column:      "col1",
			separator:   " ",
			intoColumns: []string{"col2", "col3", "col4"},
			input: map[string]interface{}{
				"col1": "value1 value2",
			},
			output: map[string]interface{}{
				"col1": "value1 value2",
				"col2": "value1",
				"col3": "value2",
				"col4": "",
			},
		},
		{
			name:        "split column by dash, no match",
			column:      "col1",
			separator:   "-",
			intoColumns: []string{"col2", "col3"},
			input: map[string]interface{}{
				"col1": "value1 value2",
			},
			output: map[string]interface{}{
				"col1": "value1 value2",
				"col2": "",
				"col3": "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnSplitTransformer{
				Column:      tc.column,
				Separator:   tc.separator,
				IntoColumns: tc.intoColumns,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
