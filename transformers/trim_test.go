package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestMultipleColumnTrimTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name    string
		columns []string
		input   map[string]interface{}
		output  map[string]interface{}
	}{
		{
			name:    "trim all columns",
			columns: []string{"col1", "col2"},
			input: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
		},
		{
			name:    "trim some columns",
			columns: []string{"col1"},
			input: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "  value2  ",
			},
		},
		{
			name:    "trim no columns",
			columns: []string{},
			input: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
			output: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.MultipleColumnTrimTransformer{
				Columns: tc.columns,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestSingleColumnTrimTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name   string
		column string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name:   "trim column",
			column: "col1",
			input: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "  value2  ",
			},
		},
		{
			name:   "trim no columns",
			column: "",
			input: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
			output: map[string]interface{}{
				"col1": "  value1  ",
				"col2": "  value2  ",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnTrimTransformer{
				Column: tc.column,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
