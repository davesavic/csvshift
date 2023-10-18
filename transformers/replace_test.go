package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestMultipleColumnReplaceTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name    string
		columns []string
		from    string
		to      string
		input   map[string]interface{}
		output  map[string]interface{}
	}{
		{
			name:    "replace in all columns",
			columns: []string{"col1", "col2"},
			from:    "old",
			to:      "new",
			input: map[string]interface{}{
				"col1": "value1 old",
				"col2": "old value2",
			},
			output: map[string]interface{}{
				"col1": "value1 new",
				"col2": "new value2",
			},
		},
		{
			name:    "replace in some columns",
			columns: []string{"col1"},
			from:    "old",
			to:      "new",
			input: map[string]interface{}{
				"col1": "value1 old",
				"col2": "old value2",
			},
			output: map[string]interface{}{
				"col1": "value1 new",
				"col2": "old value2",
			},
		},
		{
			name:    "replace no occurrence",
			columns: []string{"col1", "col2"},
			from:    "nonexistent",
			to:      "new",
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
			transformer := transformers.MultipleColumnReplaceTransformer{
				Columns: tc.columns,
				From:    tc.from,
				To:      tc.to,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestSingleColumnReplaceTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name   string
		column string
		from   string
		to     string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name:   "replace in specified column",
			column: "col1",
			from:   "old",
			to:     "new",
			input: map[string]interface{}{
				"col1": "value1 old",
				"col2": "old value2",
			},
			output: map[string]interface{}{
				"col1": "value1 new",
				"col2": "old value2",
			},
		},
		{
			name:   "replace no occurrence",
			column: "col1",
			from:   "nonexistent",
			to:     "new",
			input: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
			output: map[string]interface{}{
				"col1": "value1",
				"col2": "value2",
			},
		},
		{
			name:   "replace in non-existent column",
			column: "col3",
			from:   "old",
			to:     "new",
			input: map[string]interface{}{
				"col1": "value1 old",
				"col2": "old value2",
			},
			output: map[string]interface{}{
				"col1": "value1 old",
				"col2": "old value2",
			},
		},
		{
			name:   "replace from ' to \" in column",
			column: "col1",
			from:   "'",
			to:     "\"",
			input: map[string]interface{}{
				"col1": "value1 '",
				"col2": "' value2",
			},
			output: map[string]interface{}{
				"col1": "value1 \"",
				"col2": "' value2",
			},
		},
		{
			name:   "replace from \" to ' in column",
			column: "col1",
			from:   "\"",
			to:     "'",
			input: map[string]interface{}{
				"col1": "value1 \"",
				"col2": "\" value2",
			},
			output: map[string]interface{}{
				"col1": "value1 '",
				"col2": "\" value2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnReplaceTransformer{
				Column: tc.column,
				From:   tc.from,
				To:     tc.to,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
