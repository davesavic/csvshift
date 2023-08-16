// Code generated from /Users/david.savic/Repositories/10-percent-days/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // CsvShiftGrammar

import "github.com/antlr4-go/antlr"

// CsvShiftGrammarListener is a complete listener for a parse tree produced by CsvShiftGrammarParser.
type CsvShiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterCsvTransform is called when entering the csvTransform production.
	EnterCsvTransform(c *CsvTransformContext)

	// EnterInputSection is called when entering the inputSection production.
	EnterInputSection(c *InputSectionContext)

	// EnterOutputSection is called when entering the outputSection production.
	EnterOutputSection(c *OutputSectionContext)

	// EnterColumnModifierSection is called when entering the columnModifierSection production.
	EnterColumnModifierSection(c *ColumnModifierSectionContext)

	// EnterSingleColumnModifierSection is called when entering the singleColumnModifierSection production.
	EnterSingleColumnModifierSection(c *SingleColumnModifierSectionContext)

	// EnterMultipleColumnModifierSection is called when entering the multipleColumnModifierSection production.
	EnterMultipleColumnModifierSection(c *MultipleColumnModifierSectionContext)

	// EnterSingleColumnTransformation is called when entering the singleColumnTransformation production.
	EnterSingleColumnTransformation(c *SingleColumnTransformationContext)

	// EnterMultipleColumnTransformation is called when entering the multipleColumnTransformation production.
	EnterMultipleColumnTransformation(c *MultipleColumnTransformationContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

	// ExitCsvTransform is called when exiting the csvTransform production.
	ExitCsvTransform(c *CsvTransformContext)

	// ExitInputSection is called when exiting the inputSection production.
	ExitInputSection(c *InputSectionContext)

	// ExitOutputSection is called when exiting the outputSection production.
	ExitOutputSection(c *OutputSectionContext)

	// ExitColumnModifierSection is called when exiting the columnModifierSection production.
	ExitColumnModifierSection(c *ColumnModifierSectionContext)

	// ExitSingleColumnModifierSection is called when exiting the singleColumnModifierSection production.
	ExitSingleColumnModifierSection(c *SingleColumnModifierSectionContext)

	// ExitMultipleColumnModifierSection is called when exiting the multipleColumnModifierSection production.
	ExitMultipleColumnModifierSection(c *MultipleColumnModifierSectionContext)

	// ExitSingleColumnTransformation is called when exiting the singleColumnTransformation production.
	ExitSingleColumnTransformation(c *SingleColumnTransformationContext)

	// ExitMultipleColumnTransformation is called when exiting the multipleColumnTransformation production.
	ExitMultipleColumnTransformation(c *MultipleColumnTransformationContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)
}
