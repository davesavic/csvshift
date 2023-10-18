// Code generated from /home/sav/Development/Go/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CsvShiftGrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CsvShiftGrammarParser struct {
	*antlr.BaseParser
}

var CsvShiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csvshiftgrammarParserInit() {
	staticData := &CsvShiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Input Columns'", "'Output Columns'", "'Column'", "'Columns'",
		"'-> Trim'", "'-> Replace'", "'-> Join with'", "'-> ToLower'", "'-> ToUpper'",
		"'-> Split on'", "'-> RegexReplace'", "'-> RegexExtract'", "'with'",
		"'as'", "','", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "JOIN",
		"LOWER", "UPPER", "SPLIT", "REGEXREPLACE", "REGEXEXTRACT", "WITH", "AS",
		"COMMA", "ARROW", "IDENTIFIER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"csvTransform", "inputSection", "outputSection", "columnModifierSection",
		"singleColumnModifierSection", "multipleColumnModifierSection", "singleColumnTransformation",
		"multipleColumnTransformation", "columns",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 5, 0, 21,
		8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 4, 1, 4, 1, 4, 4, 4, 42, 8,
		4, 11, 4, 12, 4, 43, 1, 5, 1, 5, 1, 5, 4, 5, 49, 8, 5, 11, 5, 12, 5, 50,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 72, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 89, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 94, 8, 8, 10, 8, 12, 8, 97,
		9, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 105, 0, 18, 1,
		0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 36, 1, 0, 0, 0, 8, 38,
		1, 0, 0, 0, 10, 45, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 88, 1, 0, 0, 0,
		16, 90, 1, 0, 0, 0, 18, 22, 3, 2, 1, 0, 19, 21, 3, 6, 3, 0, 20, 19, 1,
		0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23,
		25, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25, 26, 3, 4, 2, 0, 26, 27, 5, 0, 0,
		1, 27, 1, 1, 0, 0, 0, 28, 29, 5, 1, 0, 0, 29, 30, 3, 16, 8, 0, 30, 3, 1,
		0, 0, 0, 31, 32, 5, 2, 0, 0, 32, 33, 3, 16, 8, 0, 33, 5, 1, 0, 0, 0, 34,
		37, 3, 8, 4, 0, 35, 37, 3, 10, 5, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0,
		0, 0, 37, 7, 1, 0, 0, 0, 38, 39, 5, 3, 0, 0, 39, 41, 5, 17, 0, 0, 40, 42,
		3, 12, 6, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 9, 1, 0, 0, 0, 45, 46, 5, 4, 0, 0, 46, 48, 3, 16,
		8, 0, 47, 49, 3, 14, 7, 0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50,
		48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 11, 1, 0, 0, 0, 52, 72, 5, 5, 0,
		0, 53, 54, 5, 6, 0, 0, 54, 55, 5, 18, 0, 0, 55, 56, 5, 13, 0, 0, 56, 72,
		5, 18, 0, 0, 57, 72, 5, 8, 0, 0, 58, 72, 5, 9, 0, 0, 59, 60, 5, 10, 0,
		0, 60, 61, 5, 18, 0, 0, 61, 62, 5, 14, 0, 0, 62, 72, 3, 16, 8, 0, 63, 64,
		5, 11, 0, 0, 64, 65, 5, 18, 0, 0, 65, 66, 5, 13, 0, 0, 66, 72, 5, 18, 0,
		0, 67, 68, 5, 12, 0, 0, 68, 69, 5, 18, 0, 0, 69, 70, 5, 14, 0, 0, 70, 72,
		3, 16, 8, 0, 71, 52, 1, 0, 0, 0, 71, 53, 1, 0, 0, 0, 71, 57, 1, 0, 0, 0,
		71, 58, 1, 0, 0, 0, 71, 59, 1, 0, 0, 0, 71, 63, 1, 0, 0, 0, 71, 67, 1,
		0, 0, 0, 72, 13, 1, 0, 0, 0, 73, 89, 5, 5, 0, 0, 74, 75, 5, 6, 0, 0, 75,
		76, 5, 18, 0, 0, 76, 77, 5, 13, 0, 0, 77, 89, 5, 18, 0, 0, 78, 79, 5, 7,
		0, 0, 79, 80, 5, 18, 0, 0, 80, 81, 5, 14, 0, 0, 81, 89, 5, 17, 0, 0, 82,
		89, 5, 8, 0, 0, 83, 89, 5, 9, 0, 0, 84, 85, 5, 11, 0, 0, 85, 86, 5, 18,
		0, 0, 86, 87, 5, 13, 0, 0, 87, 89, 5, 18, 0, 0, 88, 73, 1, 0, 0, 0, 88,
		74, 1, 0, 0, 0, 88, 78, 1, 0, 0, 0, 88, 82, 1, 0, 0, 0, 88, 83, 1, 0, 0,
		0, 88, 84, 1, 0, 0, 0, 89, 15, 1, 0, 0, 0, 90, 95, 5, 17, 0, 0, 91, 92,
		5, 15, 0, 0, 92, 94, 5, 17, 0, 0, 93, 91, 1, 0, 0, 0, 94, 97, 1, 0, 0,
		0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 17, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 7, 22, 36, 43, 50, 71, 88, 95,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CsvShiftGrammarParserInit initializes any static state used to implement CsvShiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCsvShiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CsvShiftGrammarParserInit() {
	staticData := &CsvShiftGrammarParserStaticData
	staticData.once.Do(csvshiftgrammarParserInit)
}

// NewCsvShiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCsvShiftGrammarParser(input antlr.TokenStream) *CsvShiftGrammarParser {
	CsvShiftGrammarParserInit()
	this := new(CsvShiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CsvShiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CsvShiftGrammar.g4"

	return this
}

// CsvShiftGrammarParser tokens.
const (
	CsvShiftGrammarParserEOF          = antlr.TokenEOF
	CsvShiftGrammarParserINPUT        = 1
	CsvShiftGrammarParserOUTPUT       = 2
	CsvShiftGrammarParserCOLUMN       = 3
	CsvShiftGrammarParserCOLUMNS      = 4
	CsvShiftGrammarParserTRIM         = 5
	CsvShiftGrammarParserREPLACE      = 6
	CsvShiftGrammarParserJOIN         = 7
	CsvShiftGrammarParserLOWER        = 8
	CsvShiftGrammarParserUPPER        = 9
	CsvShiftGrammarParserSPLIT        = 10
	CsvShiftGrammarParserREGEXREPLACE = 11
	CsvShiftGrammarParserREGEXEXTRACT = 12
	CsvShiftGrammarParserWITH         = 13
	CsvShiftGrammarParserAS           = 14
	CsvShiftGrammarParserCOMMA        = 15
	CsvShiftGrammarParserARROW        = 16
	CsvShiftGrammarParserIDENTIFIER   = 17
	CsvShiftGrammarParserSTRING       = 18
	CsvShiftGrammarParserWS           = 19
)

// CsvShiftGrammarParser rules.
const (
	CsvShiftGrammarParserRULE_csvTransform                  = 0
	CsvShiftGrammarParserRULE_inputSection                  = 1
	CsvShiftGrammarParserRULE_outputSection                 = 2
	CsvShiftGrammarParserRULE_columnModifierSection         = 3
	CsvShiftGrammarParserRULE_singleColumnModifierSection   = 4
	CsvShiftGrammarParserRULE_multipleColumnModifierSection = 5
	CsvShiftGrammarParserRULE_singleColumnTransformation    = 6
	CsvShiftGrammarParserRULE_multipleColumnTransformation  = 7
	CsvShiftGrammarParserRULE_columns                       = 8
)

// ICsvTransformContext is an interface to support dynamic dispatch.
type ICsvTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputSection() IInputSectionContext
	OutputSection() IOutputSectionContext
	EOF() antlr.TerminalNode
	AllColumnModifierSection() []IColumnModifierSectionContext
	ColumnModifierSection(i int) IColumnModifierSectionContext

	// IsCsvTransformContext differentiates from other interfaces.
	IsCsvTransformContext()
}

type CsvTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsvTransformContext() *CsvTransformContext {
	var p = new(CsvTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_csvTransform
	return p
}

func InitEmptyCsvTransformContext(p *CsvTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_csvTransform
}

func (*CsvTransformContext) IsCsvTransformContext() {}

func NewCsvTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsvTransformContext {
	var p = new(CsvTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_csvTransform

	return p
}

func (s *CsvTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *CsvTransformContext) InputSection() IInputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSectionContext)
}

func (s *CsvTransformContext) OutputSection() IOutputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputSectionContext)
}

func (s *CsvTransformContext) EOF() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserEOF, 0)
}

func (s *CsvTransformContext) AllColumnModifierSection() []IColumnModifierSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnModifierSectionContext); ok {
			len++
		}
	}

	tst := make([]IColumnModifierSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnModifierSectionContext); ok {
			tst[i] = t.(IColumnModifierSectionContext)
			i++
		}
	}

	return tst
}

func (s *CsvTransformContext) ColumnModifierSection(i int) IColumnModifierSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnModifierSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnModifierSectionContext)
}

func (s *CsvTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsvTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CsvTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterCsvTransform(s)
	}
}

func (s *CsvTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitCsvTransform(s)
	}
}

func (p *CsvShiftGrammarParser) CsvTransform() (localctx ICsvTransformContext) {
	localctx = NewCsvTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CsvShiftGrammarParserRULE_csvTransform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.InputSection()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsvShiftGrammarParserCOLUMN || _la == CsvShiftGrammarParserCOLUMNS {
		{
			p.SetState(19)
			p.ColumnModifierSection()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.OutputSection()
	}
	{
		p.SetState(26)
		p.Match(CsvShiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputSectionContext is an interface to support dynamic dispatch.
type IInputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
	Columns() IColumnsContext

	// IsInputSectionContext differentiates from other interfaces.
	IsInputSectionContext()
}

type InputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSectionContext() *InputSectionContext {
	var p = new(InputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_inputSection
	return p
}

func InitEmptyInputSectionContext(p *InputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_inputSection
}

func (*InputSectionContext) IsInputSectionContext() {}

func NewInputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSectionContext {
	var p = new(InputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_inputSection

	return p
}

func (s *InputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSectionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserINPUT, 0)
}

func (s *InputSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *InputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterInputSection(s)
	}
}

func (s *InputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitInputSection(s)
	}
}

func (p *CsvShiftGrammarParser) InputSection() (localctx IInputSectionContext) {
	localctx = NewInputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CsvShiftGrammarParserRULE_inputSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(CsvShiftGrammarParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Columns()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutputSectionContext is an interface to support dynamic dispatch.
type IOutputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTPUT() antlr.TerminalNode
	Columns() IColumnsContext

	// IsOutputSectionContext differentiates from other interfaces.
	IsOutputSectionContext()
}

type OutputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputSectionContext() *OutputSectionContext {
	var p = new(OutputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_outputSection
	return p
}

func InitEmptyOutputSectionContext(p *OutputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_outputSection
}

func (*OutputSectionContext) IsOutputSectionContext() {}

func NewOutputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputSectionContext {
	var p = new(OutputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_outputSection

	return p
}

func (s *OutputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputSectionContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserOUTPUT, 0)
}

func (s *OutputSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *OutputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterOutputSection(s)
	}
}

func (s *OutputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitOutputSection(s)
	}
}

func (p *CsvShiftGrammarParser) OutputSection() (localctx IOutputSectionContext) {
	localctx = NewOutputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CsvShiftGrammarParserRULE_outputSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(CsvShiftGrammarParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Columns()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnModifierSectionContext is an interface to support dynamic dispatch.
type IColumnModifierSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleColumnModifierSection() ISingleColumnModifierSectionContext
	MultipleColumnModifierSection() IMultipleColumnModifierSectionContext

	// IsColumnModifierSectionContext differentiates from other interfaces.
	IsColumnModifierSectionContext()
}

type ColumnModifierSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnModifierSectionContext() *ColumnModifierSectionContext {
	var p = new(ColumnModifierSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columnModifierSection
	return p
}

func InitEmptyColumnModifierSectionContext(p *ColumnModifierSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columnModifierSection
}

func (*ColumnModifierSectionContext) IsColumnModifierSectionContext() {}

func NewColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnModifierSectionContext {
	var p = new(ColumnModifierSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_columnModifierSection

	return p
}

func (s *ColumnModifierSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnModifierSectionContext) SingleColumnModifierSection() ISingleColumnModifierSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleColumnModifierSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleColumnModifierSectionContext)
}

func (s *ColumnModifierSectionContext) MultipleColumnModifierSection() IMultipleColumnModifierSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleColumnModifierSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleColumnModifierSectionContext)
}

func (s *ColumnModifierSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnModifierSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnModifierSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterColumnModifierSection(s)
	}
}

func (s *ColumnModifierSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitColumnModifierSection(s)
	}
}

func (p *CsvShiftGrammarParser) ColumnModifierSection() (localctx IColumnModifierSectionContext) {
	localctx = NewColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CsvShiftGrammarParserRULE_columnModifierSection)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsvShiftGrammarParserCOLUMN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.SingleColumnModifierSection()
		}

	case CsvShiftGrammarParserCOLUMNS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.MultipleColumnModifierSection()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleColumnModifierSectionContext is an interface to support dynamic dispatch.
type ISingleColumnModifierSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLUMN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllSingleColumnTransformation() []ISingleColumnTransformationContext
	SingleColumnTransformation(i int) ISingleColumnTransformationContext

	// IsSingleColumnModifierSectionContext differentiates from other interfaces.
	IsSingleColumnModifierSectionContext()
}

type SingleColumnModifierSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleColumnModifierSectionContext() *SingleColumnModifierSectionContext {
	var p = new(SingleColumnModifierSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnModifierSection
	return p
}

func InitEmptySingleColumnModifierSectionContext(p *SingleColumnModifierSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnModifierSection
}

func (*SingleColumnModifierSectionContext) IsSingleColumnModifierSectionContext() {}

func NewSingleColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleColumnModifierSectionContext {
	var p = new(SingleColumnModifierSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnModifierSection

	return p
}

func (s *SingleColumnModifierSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleColumnModifierSectionContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserCOLUMN, 0)
}

func (s *SingleColumnModifierSectionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserIDENTIFIER, 0)
}

func (s *SingleColumnModifierSectionContext) AllSingleColumnTransformation() []ISingleColumnTransformationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleColumnTransformationContext); ok {
			len++
		}
	}

	tst := make([]ISingleColumnTransformationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleColumnTransformationContext); ok {
			tst[i] = t.(ISingleColumnTransformationContext)
			i++
		}
	}

	return tst
}

func (s *SingleColumnModifierSectionContext) SingleColumnTransformation(i int) ISingleColumnTransformationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleColumnTransformationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleColumnTransformationContext)
}

func (s *SingleColumnModifierSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleColumnModifierSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleColumnModifierSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterSingleColumnModifierSection(s)
	}
}

func (s *SingleColumnModifierSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitSingleColumnModifierSection(s)
	}
}

func (p *CsvShiftGrammarParser) SingleColumnModifierSection() (localctx ISingleColumnModifierSectionContext) {
	localctx = NewSingleColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CsvShiftGrammarParserRULE_singleColumnModifierSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CsvShiftGrammarParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(CsvShiftGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8032) != 0) {
		{
			p.SetState(40)
			p.SingleColumnTransformation()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultipleColumnModifierSectionContext is an interface to support dynamic dispatch.
type IMultipleColumnModifierSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLUMNS() antlr.TerminalNode
	Columns() IColumnsContext
	AllMultipleColumnTransformation() []IMultipleColumnTransformationContext
	MultipleColumnTransformation(i int) IMultipleColumnTransformationContext

	// IsMultipleColumnModifierSectionContext differentiates from other interfaces.
	IsMultipleColumnModifierSectionContext()
}

type MultipleColumnModifierSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleColumnModifierSectionContext() *MultipleColumnModifierSectionContext {
	var p = new(MultipleColumnModifierSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnModifierSection
	return p
}

func InitEmptyMultipleColumnModifierSectionContext(p *MultipleColumnModifierSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnModifierSection
}

func (*MultipleColumnModifierSectionContext) IsMultipleColumnModifierSectionContext() {}

func NewMultipleColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleColumnModifierSectionContext {
	var p = new(MultipleColumnModifierSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnModifierSection

	return p
}

func (s *MultipleColumnModifierSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleColumnModifierSectionContext) COLUMNS() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserCOLUMNS, 0)
}

func (s *MultipleColumnModifierSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *MultipleColumnModifierSectionContext) AllMultipleColumnTransformation() []IMultipleColumnTransformationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultipleColumnTransformationContext); ok {
			len++
		}
	}

	tst := make([]IMultipleColumnTransformationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultipleColumnTransformationContext); ok {
			tst[i] = t.(IMultipleColumnTransformationContext)
			i++
		}
	}

	return tst
}

func (s *MultipleColumnModifierSectionContext) MultipleColumnTransformation(i int) IMultipleColumnTransformationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleColumnTransformationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleColumnTransformationContext)
}

func (s *MultipleColumnModifierSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleColumnModifierSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipleColumnModifierSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterMultipleColumnModifierSection(s)
	}
}

func (s *MultipleColumnModifierSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitMultipleColumnModifierSection(s)
	}
}

func (p *CsvShiftGrammarParser) MultipleColumnModifierSection() (localctx IMultipleColumnModifierSectionContext) {
	localctx = NewMultipleColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CsvShiftGrammarParserRULE_multipleColumnModifierSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CsvShiftGrammarParserCOLUMNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Columns()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3040) != 0) {
		{
			p.SetState(47)
			p.MultipleColumnTransformation()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleColumnTransformationContext is an interface to support dynamic dispatch.
type ISingleColumnTransformationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRIM() antlr.TerminalNode
	REPLACE() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	WITH() antlr.TerminalNode
	LOWER() antlr.TerminalNode
	UPPER() antlr.TerminalNode
	SPLIT() antlr.TerminalNode
	AS() antlr.TerminalNode
	Columns() IColumnsContext
	REGEXREPLACE() antlr.TerminalNode
	REGEXEXTRACT() antlr.TerminalNode

	// IsSingleColumnTransformationContext differentiates from other interfaces.
	IsSingleColumnTransformationContext()
}

type SingleColumnTransformationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleColumnTransformationContext() *SingleColumnTransformationContext {
	var p = new(SingleColumnTransformationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnTransformation
	return p
}

func InitEmptySingleColumnTransformationContext(p *SingleColumnTransformationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnTransformation
}

func (*SingleColumnTransformationContext) IsSingleColumnTransformationContext() {}

func NewSingleColumnTransformationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleColumnTransformationContext {
	var p = new(SingleColumnTransformationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnTransformation

	return p
}

func (s *SingleColumnTransformationContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleColumnTransformationContext) TRIM() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserTRIM, 0)
}

func (s *SingleColumnTransformationContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserREPLACE, 0)
}

func (s *SingleColumnTransformationContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CsvShiftGrammarParserSTRING)
}

func (s *SingleColumnTransformationContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserSTRING, i)
}

func (s *SingleColumnTransformationContext) WITH() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserWITH, 0)
}

func (s *SingleColumnTransformationContext) LOWER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserLOWER, 0)
}

func (s *SingleColumnTransformationContext) UPPER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserUPPER, 0)
}

func (s *SingleColumnTransformationContext) SPLIT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserSPLIT, 0)
}

func (s *SingleColumnTransformationContext) AS() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserAS, 0)
}

func (s *SingleColumnTransformationContext) Columns() IColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *SingleColumnTransformationContext) REGEXREPLACE() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserREGEXREPLACE, 0)
}

func (s *SingleColumnTransformationContext) REGEXEXTRACT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserREGEXEXTRACT, 0)
}

func (s *SingleColumnTransformationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleColumnTransformationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleColumnTransformationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterSingleColumnTransformation(s)
	}
}

func (s *SingleColumnTransformationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitSingleColumnTransformation(s)
	}
}

func (p *CsvShiftGrammarParser) SingleColumnTransformation() (localctx ISingleColumnTransformationContext) {
	localctx = NewSingleColumnTransformationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CsvShiftGrammarParserRULE_singleColumnTransformation)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsvShiftGrammarParserTRIM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(CsvShiftGrammarParserTRIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserREPLACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(CsvShiftGrammarParserREPLACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(CsvShiftGrammarParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserLOWER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Match(CsvShiftGrammarParserLOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserUPPER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Match(CsvShiftGrammarParserUPPER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserSPLIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Match(CsvShiftGrammarParserSPLIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(CsvShiftGrammarParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Columns()
		}

	case CsvShiftGrammarParserREGEXREPLACE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.Match(CsvShiftGrammarParserREGEXREPLACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(CsvShiftGrammarParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserREGEXEXTRACT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(67)
			p.Match(CsvShiftGrammarParserREGEXEXTRACT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(CsvShiftGrammarParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Columns()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultipleColumnTransformationContext is an interface to support dynamic dispatch.
type IMultipleColumnTransformationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRIM() antlr.TerminalNode
	REPLACE() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	WITH() antlr.TerminalNode
	JOIN() antlr.TerminalNode
	AS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LOWER() antlr.TerminalNode
	UPPER() antlr.TerminalNode
	REGEXREPLACE() antlr.TerminalNode

	// IsMultipleColumnTransformationContext differentiates from other interfaces.
	IsMultipleColumnTransformationContext()
}

type MultipleColumnTransformationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleColumnTransformationContext() *MultipleColumnTransformationContext {
	var p = new(MultipleColumnTransformationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnTransformation
	return p
}

func InitEmptyMultipleColumnTransformationContext(p *MultipleColumnTransformationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnTransformation
}

func (*MultipleColumnTransformationContext) IsMultipleColumnTransformationContext() {}

func NewMultipleColumnTransformationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleColumnTransformationContext {
	var p = new(MultipleColumnTransformationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnTransformation

	return p
}

func (s *MultipleColumnTransformationContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleColumnTransformationContext) TRIM() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserTRIM, 0)
}

func (s *MultipleColumnTransformationContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserREPLACE, 0)
}

func (s *MultipleColumnTransformationContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(CsvShiftGrammarParserSTRING)
}

func (s *MultipleColumnTransformationContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserSTRING, i)
}

func (s *MultipleColumnTransformationContext) WITH() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserWITH, 0)
}

func (s *MultipleColumnTransformationContext) JOIN() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserJOIN, 0)
}

func (s *MultipleColumnTransformationContext) AS() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserAS, 0)
}

func (s *MultipleColumnTransformationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserIDENTIFIER, 0)
}

func (s *MultipleColumnTransformationContext) LOWER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserLOWER, 0)
}

func (s *MultipleColumnTransformationContext) UPPER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserUPPER, 0)
}

func (s *MultipleColumnTransformationContext) REGEXREPLACE() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserREGEXREPLACE, 0)
}

func (s *MultipleColumnTransformationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleColumnTransformationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipleColumnTransformationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterMultipleColumnTransformation(s)
	}
}

func (s *MultipleColumnTransformationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitMultipleColumnTransformation(s)
	}
}

func (p *CsvShiftGrammarParser) MultipleColumnTransformation() (localctx IMultipleColumnTransformationContext) {
	localctx = NewMultipleColumnTransformationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CsvShiftGrammarParserRULE_multipleColumnTransformation)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsvShiftGrammarParserTRIM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(CsvShiftGrammarParserTRIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserREPLACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(CsvShiftGrammarParserREPLACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(CsvShiftGrammarParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserJOIN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(CsvShiftGrammarParserJOIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(CsvShiftGrammarParserAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(CsvShiftGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserLOWER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.Match(CsvShiftGrammarParserLOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserUPPER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)
			p.Match(CsvShiftGrammarParserUPPER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsvShiftGrammarParserREGEXREPLACE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(84)
			p.Match(CsvShiftGrammarParserREGEXREPLACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(CsvShiftGrammarParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(CsvShiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columns
	return p
}

func InitEmptyColumnsContext(p *ColumnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columns
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(CsvShiftGrammarParserIDENTIFIER)
}

func (s *ColumnsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserIDENTIFIER, i)
}

func (s *ColumnsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CsvShiftGrammarParserCOMMA)
}

func (s *ColumnsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserCOMMA, i)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CsvShiftGrammarListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *CsvShiftGrammarParser) Columns() (localctx IColumnsContext) {
	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CsvShiftGrammarParserRULE_columns)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(CsvShiftGrammarParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsvShiftGrammarParserCOMMA {
		{
			p.SetState(91)
			p.Match(CsvShiftGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(CsvShiftGrammarParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
