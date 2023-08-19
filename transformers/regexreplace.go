package transformers

import (
	parser "csvshift/gen"
	"regexp"
)

type SingleColumnRegexReplaceTransformer struct {
	Column      string
	Pattern     string
	Replacement string
}

type SingleColumnRegexReplaceTransformerFactory struct{}

type MultipleColumnRegexReplaceTransformer struct {
	Columns     []string
	Pattern     string
	Replacement string
}

type MultipleColumnRegexReplaceTransformerFactory struct{}

func (t *MultipleColumnRegexReplaceTransformer) Apply(row map[string]interface{}) {
	for _, column := range t.Columns {
		val, ok := row[column].(string)
		if ok {
			re := regexp.MustCompile(t.Pattern)
			row[column] = re.ReplaceAllString(val, t.Replacement)
		}
	}
}

func (t *MultipleColumnRegexReplaceTransformerFactory) IsMatch(modifier parser.IMultipleColumnTransformationContext) bool {
	return modifier.REGEXREPLACE() != nil
}

func (t *MultipleColumnRegexReplaceTransformerFactory) Create(columns []string, modifier parser.IMultipleColumnTransformationContext) Transformer {
	return &MultipleColumnRegexReplaceTransformer{
		Columns:     columns,
		Pattern:     ExtractStringContent(modifier.STRING(0).GetText()),
		Replacement: ExtractStringContent(modifier.STRING(1).GetText()),
	}
}

func (t *SingleColumnRegexReplaceTransformer) Apply(row map[string]interface{}) {
	val, ok := row[t.Column].(string)
	if ok {
		re := regexp.MustCompile(t.Pattern)
		row[t.Column] = re.ReplaceAllString(val, t.Replacement)
	}
}

func (t *SingleColumnRegexReplaceTransformerFactory) IsMatch(modifier parser.ISingleColumnTransformationContext) bool {
	return modifier.REGEXREPLACE() != nil
}

func (t *SingleColumnRegexReplaceTransformerFactory) Create(column string, modifier parser.ISingleColumnTransformationContext) Transformer {
	return &SingleColumnRegexReplaceTransformer{
		Column:      column,
		Pattern:     ExtractStringContent(modifier.STRING(0).GetText()),
		Replacement: ExtractStringContent(modifier.STRING(1).GetText()),
	}
}
