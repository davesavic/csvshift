// Code generated from /home/sav/Development/Go/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // CsvShiftGrammar

import "github.com/antlr4-go/antlr"

// BaseCsvShiftGrammarListener is a complete listener for a parse tree produced by CsvShiftGrammarParser.
type BaseCsvShiftGrammarListener struct{}

var _ CsvShiftGrammarListener = &BaseCsvShiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCsvShiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCsvShiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCsvShiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCsvShiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCsvTransform is called when production csvTransform is entered.
func (s *BaseCsvShiftGrammarListener) EnterCsvTransform(ctx *CsvTransformContext) {}

// ExitCsvTransform is called when production csvTransform is exited.
func (s *BaseCsvShiftGrammarListener) ExitCsvTransform(ctx *CsvTransformContext) {}

// EnterInputSection is called when production inputSection is entered.
func (s *BaseCsvShiftGrammarListener) EnterInputSection(ctx *InputSectionContext) {}

// ExitInputSection is called when production inputSection is exited.
func (s *BaseCsvShiftGrammarListener) ExitInputSection(ctx *InputSectionContext) {}

// EnterOutputSection is called when production outputSection is entered.
func (s *BaseCsvShiftGrammarListener) EnterOutputSection(ctx *OutputSectionContext) {}

// ExitOutputSection is called when production outputSection is exited.
func (s *BaseCsvShiftGrammarListener) ExitOutputSection(ctx *OutputSectionContext) {}

// EnterColumnModifierSection is called when production columnModifierSection is entered.
func (s *BaseCsvShiftGrammarListener) EnterColumnModifierSection(ctx *ColumnModifierSectionContext) {}

// ExitColumnModifierSection is called when production columnModifierSection is exited.
func (s *BaseCsvShiftGrammarListener) ExitColumnModifierSection(ctx *ColumnModifierSectionContext) {}

// EnterSingleColumnModifierSection is called when production singleColumnModifierSection is entered.
func (s *BaseCsvShiftGrammarListener) EnterSingleColumnModifierSection(ctx *SingleColumnModifierSectionContext) {
}

// ExitSingleColumnModifierSection is called when production singleColumnModifierSection is exited.
func (s *BaseCsvShiftGrammarListener) ExitSingleColumnModifierSection(ctx *SingleColumnModifierSectionContext) {
}

// EnterMultipleColumnModifierSection is called when production multipleColumnModifierSection is entered.
func (s *BaseCsvShiftGrammarListener) EnterMultipleColumnModifierSection(ctx *MultipleColumnModifierSectionContext) {
}

// ExitMultipleColumnModifierSection is called when production multipleColumnModifierSection is exited.
func (s *BaseCsvShiftGrammarListener) ExitMultipleColumnModifierSection(ctx *MultipleColumnModifierSectionContext) {
}

// EnterSingleColumnTransformation is called when production singleColumnTransformation is entered.
func (s *BaseCsvShiftGrammarListener) EnterSingleColumnTransformation(ctx *SingleColumnTransformationContext) {
}

// ExitSingleColumnTransformation is called when production singleColumnTransformation is exited.
func (s *BaseCsvShiftGrammarListener) ExitSingleColumnTransformation(ctx *SingleColumnTransformationContext) {
}

// EnterMultipleColumnTransformation is called when production multipleColumnTransformation is entered.
func (s *BaseCsvShiftGrammarListener) EnterMultipleColumnTransformation(ctx *MultipleColumnTransformationContext) {
}

// ExitMultipleColumnTransformation is called when production multipleColumnTransformation is exited.
func (s *BaseCsvShiftGrammarListener) ExitMultipleColumnTransformation(ctx *MultipleColumnTransformationContext) {
}

// EnterColumns is called when production columns is entered.
func (s *BaseCsvShiftGrammarListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseCsvShiftGrammarListener) ExitColumns(ctx *ColumnsContext) {}
