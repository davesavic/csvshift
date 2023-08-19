package transformers_test

import (
	"csvshift/transformers"
	"reflect"
	"testing"
)

func TestSingleColumnRegexReplaceTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name        string
		column      string
		pattern     string
		replacement string
		input       map[string]interface{}
		output      map[string]interface{}
	}{
		{
			name:        "replace using groups",
			column:      "date",
			pattern:     `(\d{4})-(\d{2})-(\d{2})`,
			replacement: "$3/$2/$1",
			input: map[string]interface{}{
				"date": "2023-08-17",
			},
			output: map[string]interface{}{
				"date": "17/08/2023",
			},
		},
		{
			name:        "replace with empty string",
			column:      "date",
			pattern:     `\d{4}-`,
			replacement: "",
			input: map[string]interface{}{
				"date": "2023-08-17",
			},
			output: map[string]interface{}{
				"date": "08-17",
			},
		},
		{
			name:        "no match to replace",
			column:      "date",
			pattern:     `\d{4}-`,
			replacement: "2023-",
			input: map[string]interface{}{
				"date": "08-17",
			},
			output: map[string]interface{}{
				"date": "08-17",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.SingleColumnRegexReplaceTransformer{
				Column:      tc.column,
				Pattern:     tc.pattern,
				Replacement: tc.replacement,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}

func TestMultipleColumnRegexReplaceTransformer_Apply(t *testing.T) {
	testCases := []struct {
		name        string
		columns     []string
		pattern     string
		replacement string
		input       map[string]interface{}
		output      map[string]interface{}
	}{
		{
			name:        "replace using groups",
			columns:     []string{"start_date", "end_date"},
			pattern:     `(\d{4})-(\d{2})-(\d{2})`,
			replacement: "$3/$2/$1",
			input: map[string]interface{}{
				"start_date": "2023-08-17",
				"end_date":   "2023-09-17",
			},
			output: map[string]interface{}{
				"start_date": "17/08/2023",
				"end_date":   "17/09/2023",
			},
		},
		{
			name:        "replace with empty string",
			columns:     []string{"start_date", "end_date"},
			pattern:     `\d{4}-`,
			replacement: "",
			input: map[string]interface{}{
				"start_date": "2023-08-17",
				"end_date":   "2023-09-17",
			},
			output: map[string]interface{}{
				"start_date": "08-17",
				"end_date":   "09-17",
			},
		},
		{
			name:        "no match to replace",
			columns:     []string{"start_date", "end_date"},
			pattern:     `\d{4}-`,
			replacement: "2023-",
			input: map[string]interface{}{
				"start_date": "08-17",
				"end_date":   "09-17",
			},
			output: map[string]interface{}{
				"start_date": "08-17",
				"end_date":   "09-17",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transformer := transformers.MultipleColumnRegexReplaceTransformer{
				Columns:     tc.columns,
				Pattern:     tc.pattern,
				Replacement: tc.replacement,
			}
			transformer.Apply(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("expected %v, got %v", tc.output, tc.input)
			}
		})
	}
}
