package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestSingleColumnLowerTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name   string
		column string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name:   "convert to lowercase in specified column",
			column: "col1",
			input: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "VALUE2",
			},
		},
		{
			name:   "already lowercase",
			column: "col1",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "VALUE2",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "VALUE2",
			},
		},
		{
			name:   "convert non-existent column",
			column: "col3",
			input: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
			},
			output: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnLowerTransformer{
				Column: tc.column,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestSingleColumnUpperTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name   string
		column string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name:   "convert to uppercase in specified column",
			column: "col1",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "value2",
			},
		},
		{
			name:   "already uppercase",
			column: "col1",
			input: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "value2",
			},
		},
		{
			name:   "convert non-existent column",
			column: "col3",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnUpperTransformer{
				Column: tc.column,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestMultipleColumnUpperTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name    string
		columns []string
		input   map[string]interface{}
		output  map[string]interface{}
	}{
		{
			name:    "convert to uppercase in specified columns",
			columns: []string{"col1", "col2"},
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "value3",
			},
			output: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
				"col3": "value3",
			},
		},
		{
			name:    "already uppercase",
			columns: []string{"col1", "col2"},
			input: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
				"col3": "value3",
			},
			output: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
				"col3": "value3",
			},
		},
		{
			name:    "convert non-existent columns",
			columns: []string{"col4", "col5"},
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "value3",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "value3",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.MultipleColumnUpperTransformer{
				Columns: tc.columns,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestMultipleColumnLowerTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name    string
		columns []string
		input   map[string]interface{}
		output  map[string]interface{}
	}{
		{
			name:    "convert to lowercase in specified columns",
			columns: []string{"col1", "col2"},
			input: map[string]interface{}{
				"col1": "VALUE1",
				"col2": "VALUE2",
				"col3": "VALUE3",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "VALUE3",
			},
		},
		{
			name:    "already lowercase",
			columns: []string{"col1", "col2"},
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "VALUE3",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "VALUE3",
			},
		},
		{
			name:    "convert non-existent columns",
			columns: []string{"col4", "col5"},
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "VALUE3",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
				"col3": "VALUE3",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.MultipleColumnLowerTransformer{
				Columns: tc.columns,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
