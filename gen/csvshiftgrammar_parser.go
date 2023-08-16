// Code generated from /Users/david.savic/Repositories/10-percent-days/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // CsvShiftGrammar

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr4-go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type CsvShiftGrammarParser struct {
	*antlr.BaseParser
}

var csvshiftgrammarParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func csvshiftgrammarParserInit() {
  staticData := &csvshiftgrammarParserStaticData
  staticData.literalNames = []string{
    "", "'Input Columns'", "'Output Columns'", "'Column'", "'Columns'", 
    "'-> Trim'", "'-> Replace'", "'-> Combine with'", "'with'", "'as'", 
    "','", "'\"'", "'->'",
  }
  staticData.symbolicNames = []string{
    "", "INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "COMBINE", 
    "WITH", "AS", "COMMA", "QUOTE", "ARROW", "IDENTIFIER", "STRING", "WS",
  }
  staticData.ruleNames = []string{
    "csvTransform", "inputSection", "outputSection", "columnModifierSection", 
    "singleColumnModifierSection", "multipleColumnModifierSection", "singleColumnTransformation", 
    "multipleColumnTransformation", "columns",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 15, 79, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 5, 0, 21, 
	8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 
	1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 4, 1, 4, 1, 4, 4, 4, 42, 8, 
	4, 11, 4, 12, 4, 43, 1, 5, 1, 5, 1, 5, 4, 5, 49, 8, 5, 11, 5, 12, 5, 50, 
	1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 58, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 
	7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 69, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 74, 
	8, 8, 10, 8, 12, 8, 77, 9, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 
	16, 0, 0, 77, 0, 18, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 
	6, 36, 1, 0, 0, 0, 8, 38, 1, 0, 0, 0, 10, 45, 1, 0, 0, 0, 12, 57, 1, 0, 
	0, 0, 14, 68, 1, 0, 0, 0, 16, 70, 1, 0, 0, 0, 18, 22, 3, 2, 1, 0, 19, 21, 
	3, 6, 3, 0, 20, 19, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 
	22, 23, 1, 0, 0, 0, 23, 25, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25, 26, 3, 
	4, 2, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 29, 5, 1, 0, 0, 29, 
	30, 3, 16, 8, 0, 30, 3, 1, 0, 0, 0, 31, 32, 5, 2, 0, 0, 32, 33, 3, 16, 
	8, 0, 33, 5, 1, 0, 0, 0, 34, 37, 3, 8, 4, 0, 35, 37, 3, 10, 5, 0, 36, 34, 
	1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 7, 1, 0, 0, 0, 38, 39, 5, 3, 0, 0, 
	39, 41, 5, 13, 0, 0, 40, 42, 3, 12, 6, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 
	0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 9, 1, 0, 0, 0, 45, 
	46, 5, 4, 0, 0, 46, 48, 3, 16, 8, 0, 47, 49, 3, 14, 7, 0, 48, 47, 1, 0, 
	0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 11, 
	1, 0, 0, 0, 52, 58, 5, 5, 0, 0, 53, 54, 5, 6, 0, 0, 54, 55, 5, 14, 0, 0, 
	55, 56, 5, 8, 0, 0, 56, 58, 5, 14, 0, 0, 57, 52, 1, 0, 0, 0, 57, 53, 1, 
	0, 0, 0, 58, 13, 1, 0, 0, 0, 59, 69, 5, 5, 0, 0, 60, 61, 5, 6, 0, 0, 61, 
	62, 5, 14, 0, 0, 62, 63, 5, 8, 0, 0, 63, 69, 5, 14, 0, 0, 64, 65, 5, 7, 
	0, 0, 65, 66, 5, 14, 0, 0, 66, 67, 5, 9, 0, 0, 67, 69, 5, 13, 0, 0, 68, 
	59, 1, 0, 0, 0, 68, 60, 1, 0, 0, 0, 68, 64, 1, 0, 0, 0, 69, 15, 1, 0, 0, 
	0, 70, 75, 5, 13, 0, 0, 71, 72, 5, 10, 0, 0, 72, 74, 5, 13, 0, 0, 73, 71, 
	1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 
	76, 17, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 7, 22, 36, 43, 50, 57, 68, 75,
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
  staticData := &csvshiftgrammarParserStaticData
  staticData.once.Do(csvshiftgrammarParserInit)
}

// NewCsvShiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCsvShiftGrammarParser(input antlr.TokenStream) *CsvShiftGrammarParser {
	CsvShiftGrammarParserInit()
	this := new(CsvShiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &csvshiftgrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CsvShiftGrammar.g4"

	return this
}


// CsvShiftGrammarParser tokens.
const (
	CsvShiftGrammarParserEOF = antlr.TokenEOF
	CsvShiftGrammarParserINPUT = 1
	CsvShiftGrammarParserOUTPUT = 2
	CsvShiftGrammarParserCOLUMN = 3
	CsvShiftGrammarParserCOLUMNS = 4
	CsvShiftGrammarParserTRIM = 5
	CsvShiftGrammarParserREPLACE = 6
	CsvShiftGrammarParserCOMBINE = 7
	CsvShiftGrammarParserWITH = 8
	CsvShiftGrammarParserAS = 9
	CsvShiftGrammarParserCOMMA = 10
	CsvShiftGrammarParserQUOTE = 11
	CsvShiftGrammarParserARROW = 12
	CsvShiftGrammarParserIDENTIFIER = 13
	CsvShiftGrammarParserSTRING = 14
	CsvShiftGrammarParserWS = 15
)

// CsvShiftGrammarParser rules.
const (
	CsvShiftGrammarParserRULE_csvTransform = 0
	CsvShiftGrammarParserRULE_inputSection = 1
	CsvShiftGrammarParserRULE_outputSection = 2
	CsvShiftGrammarParserRULE_columnModifierSection = 3
	CsvShiftGrammarParserRULE_singleColumnModifierSection = 4
	CsvShiftGrammarParserRULE_multipleColumnModifierSection = 5
	CsvShiftGrammarParserRULE_singleColumnTransformation = 6
	CsvShiftGrammarParserRULE_multipleColumnTransformation = 7
	CsvShiftGrammarParserRULE_columns = 8
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsvTransformContext() *CsvTransformContext {
	var p = new(CsvTransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_csvTransform
	return p
}

func (*CsvTransformContext) IsCsvTransformContext() {}

func NewCsvTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsvTransformContext {
	var p = new(CsvTransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_csvTransform

	return p
}

func (s *CsvTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *CsvTransformContext) InputSection() IInputSectionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSectionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSectionContext)
}

func (s *CsvTransformContext) OutputSection() IOutputSectionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputSectionContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnModifierSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewCsvTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CsvShiftGrammarParserRULE_csvTransform)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.InputSection()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == CsvShiftGrammarParserCOLUMN || _la == CsvShiftGrammarParserCOLUMNS {
		{
			p.SetState(19)
			p.ColumnModifierSection()
		}


		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.OutputSection()
	}
	{
		p.SetState(26)
		p.Match(CsvShiftGrammarParserEOF)
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSectionContext() *InputSectionContext {
	var p = new(InputSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_inputSection
	return p
}

func (*InputSectionContext) IsInputSectionContext() {}

func NewInputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSectionContext {
	var p = new(InputSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_inputSection

	return p
}

func (s *InputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSectionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserINPUT, 0)
}

func (s *InputSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewInputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CsvShiftGrammarParserRULE_inputSection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(CsvShiftGrammarParserINPUT)
	}
	{
		p.SetState(29)
		p.Columns()
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputSectionContext() *OutputSectionContext {
	var p = new(OutputSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_outputSection
	return p
}

func (*OutputSectionContext) IsOutputSectionContext() {}

func NewOutputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputSectionContext {
	var p = new(OutputSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_outputSection

	return p
}

func (s *OutputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputSectionContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserOUTPUT, 0)
}

func (s *OutputSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewOutputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CsvShiftGrammarParserRULE_outputSection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(CsvShiftGrammarParserOUTPUT)
	}
	{
		p.SetState(32)
		p.Columns()
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnModifierSectionContext() *ColumnModifierSectionContext {
	var p = new(ColumnModifierSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columnModifierSection
	return p
}

func (*ColumnModifierSectionContext) IsColumnModifierSectionContext() {}

func NewColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnModifierSectionContext {
	var p = new(ColumnModifierSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_columnModifierSection

	return p
}

func (s *ColumnModifierSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnModifierSectionContext) SingleColumnModifierSection() ISingleColumnModifierSectionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleColumnModifierSectionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleColumnModifierSectionContext)
}

func (s *ColumnModifierSectionContext) MultipleColumnModifierSection() IMultipleColumnModifierSectionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleColumnModifierSectionContext); ok {
			t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CsvShiftGrammarParserRULE_columnModifierSection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleColumnModifierSectionContext() *SingleColumnModifierSectionContext {
	var p = new(SingleColumnModifierSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnModifierSection
	return p
}

func (*SingleColumnModifierSectionContext) IsSingleColumnModifierSectionContext() {}

func NewSingleColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleColumnModifierSectionContext {
	var p = new(SingleColumnModifierSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleColumnTransformationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewSingleColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CsvShiftGrammarParserRULE_singleColumnModifierSection)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(CsvShiftGrammarParserCOLUMN)
	}
	{
		p.SetState(39)
		p.Match(CsvShiftGrammarParserIDENTIFIER)
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == CsvShiftGrammarParserTRIM || _la == CsvShiftGrammarParserREPLACE {
		{
			p.SetState(40)
			p.SingleColumnTransformation()
		}


		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleColumnModifierSectionContext() *MultipleColumnModifierSectionContext {
	var p = new(MultipleColumnModifierSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnModifierSection
	return p
}

func (*MultipleColumnModifierSectionContext) IsMultipleColumnModifierSectionContext() {}

func NewMultipleColumnModifierSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleColumnModifierSectionContext {
	var p = new(MultipleColumnModifierSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnModifierSection

	return p
}

func (s *MultipleColumnModifierSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleColumnModifierSectionContext) COLUMNS() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserCOLUMNS, 0)
}

func (s *MultipleColumnModifierSectionContext) Columns() IColumnsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleColumnTransformationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	this := p
	_ = this

	localctx = NewMultipleColumnModifierSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CsvShiftGrammarParserRULE_multipleColumnModifierSection)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CsvShiftGrammarParserCOLUMNS)
	}
	{
		p.SetState(46)
		p.Columns()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 224) != 0) {
		{
			p.SetState(47)
			p.MultipleColumnTransformation()
		}


		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
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

	// IsSingleColumnTransformationContext differentiates from other interfaces.
	IsSingleColumnTransformationContext()
}

type SingleColumnTransformationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleColumnTransformationContext() *SingleColumnTransformationContext {
	var p = new(SingleColumnTransformationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_singleColumnTransformation
	return p
}

func (*SingleColumnTransformationContext) IsSingleColumnTransformationContext() {}

func NewSingleColumnTransformationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleColumnTransformationContext {
	var p = new(SingleColumnTransformationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSingleColumnTransformationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CsvShiftGrammarParserRULE_singleColumnTransformation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsvShiftGrammarParserTRIM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(CsvShiftGrammarParserTRIM)
		}


	case CsvShiftGrammarParserREPLACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(CsvShiftGrammarParserREPLACE)
		}
		{
			p.SetState(54)
			p.Match(CsvShiftGrammarParserSTRING)
		}
		{
			p.SetState(55)
			p.Match(CsvShiftGrammarParserWITH)
		}
		{
			p.SetState(56)
			p.Match(CsvShiftGrammarParserSTRING)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
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
	COMBINE() antlr.TerminalNode
	AS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsMultipleColumnTransformationContext differentiates from other interfaces.
	IsMultipleColumnTransformationContext()
}

type MultipleColumnTransformationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleColumnTransformationContext() *MultipleColumnTransformationContext {
	var p = new(MultipleColumnTransformationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_multipleColumnTransformation
	return p
}

func (*MultipleColumnTransformationContext) IsMultipleColumnTransformationContext() {}

func NewMultipleColumnTransformationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleColumnTransformationContext {
	var p = new(MultipleColumnTransformationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

func (s *MultipleColumnTransformationContext) COMBINE() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserCOMBINE, 0)
}

func (s *MultipleColumnTransformationContext) AS() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserAS, 0)
}

func (s *MultipleColumnTransformationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsvShiftGrammarParserIDENTIFIER, 0)
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
	this := p
	_ = this

	localctx = NewMultipleColumnTransformationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CsvShiftGrammarParserRULE_multipleColumnTransformation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsvShiftGrammarParserTRIM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(CsvShiftGrammarParserTRIM)
		}


	case CsvShiftGrammarParserREPLACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(CsvShiftGrammarParserREPLACE)
		}
		{
			p.SetState(61)
			p.Match(CsvShiftGrammarParserSTRING)
		}
		{
			p.SetState(62)
			p.Match(CsvShiftGrammarParserWITH)
		}
		{
			p.SetState(63)
			p.Match(CsvShiftGrammarParserSTRING)
		}


	case CsvShiftGrammarParserCOMBINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(CsvShiftGrammarParserCOMBINE)
		}
		{
			p.SetState(65)
			p.Match(CsvShiftGrammarParserSTRING)
		}
		{
			p.SetState(66)
			p.Match(CsvShiftGrammarParserAS)
		}
		{
			p.SetState(67)
			p.Match(CsvShiftGrammarParserIDENTIFIER)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsvShiftGrammarParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CsvShiftGrammarParserRULE_columns)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(CsvShiftGrammarParserIDENTIFIER)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == CsvShiftGrammarParserCOMMA {
		{
			p.SetState(71)
			p.Match(CsvShiftGrammarParserCOMMA)
		}
		{
			p.SetState(72)
			p.Match(CsvShiftGrammarParserIDENTIFIER)
		}


		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


