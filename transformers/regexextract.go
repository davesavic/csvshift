package transformers

import (
	parser "csvshift/gen"
	"regexp"
)

type SingleColumnRegexExtractTransformer struct {
	Column      string
	Pattern     string
	IntoColumns []string
}

type SingleColumnRegexExtractTransformerFactory struct{}

func (t *SingleColumnRegexExtractTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if !ok {
		return
	}

	re := regexp.MustCompile(t.Pattern)
	matches := re.FindStringSubmatch(val)
	// Start from 1 because matches[0] is the full match, and matches[1], matches[2], ... are the capturing groups.
	for i, column := range t.IntoColumns {
		if i+1 < len(matches) {
			row[column] = matches[i+1]
		} else {
			row[column] = ""
		}
	}
}

func (t *SingleColumnRegexExtractTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.REGEXEXTRACT() != nil
}

func (t *SingleColumnRegexExtractTransformerFactory) Create(column string, modifier parser.ISingleColumnTransformationContext) Transformer {
	intoCols := modifier.Columns().AllIDENTIFIER()
	var intoColNames []string
	for _, intoCol := range intoCols {
		intoColNames = append(intoColNames, intoCol.GetText())
	}

	return &SingleColumnRegexExtractTransformer{
		Column:      column,
		Pattern:     ExtractStringContent(modifier.STRING(0).GetText()),
		IntoColumns: intoColNames,
	}
}
