// Code generated from /Users/david.savic/Repositories/10-percent-days/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr4-go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type CsvShiftGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var csvshiftgrammarlexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func csvshiftgrammarlexerLexerInit() {
  staticData := &csvshiftgrammarlexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
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
    "INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "COMBINE", 
    "WITH", "AS", "COMMA", "QUOTE", "ARROW", "IDENTIFIER", "STRING", "WS",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 15, 148, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 
	1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 
	1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 
	5, 12, 128, 8, 12, 10, 12, 12, 12, 131, 9, 12, 1, 13, 1, 13, 5, 13, 135, 
	8, 13, 10, 13, 12, 13, 138, 9, 13, 1, 13, 1, 13, 1, 14, 4, 14, 143, 8, 
	14, 11, 14, 12, 14, 144, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 
	9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 
	29, 15, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 
	95, 97, 122, 3, 0, 10, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 
	150, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 
	0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 
	0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 
	1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 
	31, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 60, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 
	9, 75, 1, 0, 0, 0, 11, 83, 1, 0, 0, 0, 13, 94, 1, 0, 0, 0, 15, 110, 1, 
	0, 0, 0, 17, 115, 1, 0, 0, 0, 19, 118, 1, 0, 0, 0, 21, 120, 1, 0, 0, 0, 
	23, 122, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 132, 1, 0, 0, 0, 29, 142, 
	1, 0, 0, 0, 31, 32, 5, 73, 0, 0, 32, 33, 5, 110, 0, 0, 33, 34, 5, 112, 
	0, 0, 34, 35, 5, 117, 0, 0, 35, 36, 5, 116, 0, 0, 36, 37, 5, 32, 0, 0, 
	37, 38, 5, 67, 0, 0, 38, 39, 5, 111, 0, 0, 39, 40, 5, 108, 0, 0, 40, 41, 
	5, 117, 0, 0, 41, 42, 5, 109, 0, 0, 42, 43, 5, 110, 0, 0, 43, 44, 5, 115, 
	0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5, 79, 0, 0, 46, 47, 5, 117, 0, 0, 47, 
	48, 5, 116, 0, 0, 48, 49, 5, 112, 0, 0, 49, 50, 5, 117, 0, 0, 50, 51, 5, 
	116, 0, 0, 51, 52, 5, 32, 0, 0, 52, 53, 5, 67, 0, 0, 53, 54, 5, 111, 0, 
	0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 117, 0, 0, 56, 57, 5, 109, 0, 0, 57, 
	58, 5, 110, 0, 0, 58, 59, 5, 115, 0, 0, 59, 4, 1, 0, 0, 0, 60, 61, 5, 67, 
	0, 0, 61, 62, 5, 111, 0, 0, 62, 63, 5, 108, 0, 0, 63, 64, 5, 117, 0, 0, 
	64, 65, 5, 109, 0, 0, 65, 66, 5, 110, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 
	5, 67, 0, 0, 68, 69, 5, 111, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5, 117, 
	0, 0, 71, 72, 5, 109, 0, 0, 72, 73, 5, 110, 0, 0, 73, 74, 5, 115, 0, 0, 
	74, 8, 1, 0, 0, 0, 75, 76, 5, 45, 0, 0, 76, 77, 5, 62, 0, 0, 77, 78, 5, 
	32, 0, 0, 78, 79, 5, 84, 0, 0, 79, 80, 5, 114, 0, 0, 80, 81, 5, 105, 0, 
	0, 81, 82, 5, 109, 0, 0, 82, 10, 1, 0, 0, 0, 83, 84, 5, 45, 0, 0, 84, 85, 
	5, 62, 0, 0, 85, 86, 5, 32, 0, 0, 86, 87, 5, 82, 0, 0, 87, 88, 5, 101, 
	0, 0, 88, 89, 5, 112, 0, 0, 89, 90, 5, 108, 0, 0, 90, 91, 5, 97, 0, 0, 
	91, 92, 5, 99, 0, 0, 92, 93, 5, 101, 0, 0, 93, 12, 1, 0, 0, 0, 94, 95, 
	5, 45, 0, 0, 95, 96, 5, 62, 0, 0, 96, 97, 5, 32, 0, 0, 97, 98, 5, 67, 0, 
	0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 109, 0, 0, 100, 101, 5, 98, 0, 0, 
	101, 102, 5, 105, 0, 0, 102, 103, 5, 110, 0, 0, 103, 104, 5, 101, 0, 0, 
	104, 105, 5, 32, 0, 0, 105, 106, 5, 119, 0, 0, 106, 107, 5, 105, 0, 0, 
	107, 108, 5, 116, 0, 0, 108, 109, 5, 104, 0, 0, 109, 14, 1, 0, 0, 0, 110, 
	111, 5, 119, 0, 0, 111, 112, 5, 105, 0, 0, 112, 113, 5, 116, 0, 0, 113, 
	114, 5, 104, 0, 0, 114, 16, 1, 0, 0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 
	5, 115, 0, 0, 117, 18, 1, 0, 0, 0, 118, 119, 5, 44, 0, 0, 119, 20, 1, 0, 
	0, 0, 120, 121, 5, 34, 0, 0, 121, 22, 1, 0, 0, 0, 122, 123, 5, 45, 0, 0, 
	123, 124, 5, 62, 0, 0, 124, 24, 1, 0, 0, 0, 125, 129, 7, 0, 0, 0, 126, 
	128, 7, 1, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 
	1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 26, 1, 0, 0, 0, 131, 129, 1, 0, 
	0, 0, 132, 136, 3, 21, 10, 0, 133, 135, 8, 2, 0, 0, 134, 133, 1, 0, 0, 
	0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 
	139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 3, 21, 10, 0, 140, 28, 
	1, 0, 0, 0, 141, 143, 7, 3, 0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 
	0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 
	146, 147, 6, 14, 0, 0, 147, 30, 1, 0, 0, 0, 4, 0, 129, 136, 144, 1, 6, 
	0, 0,
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

// CsvShiftGrammarLexerInit initializes any static state used to implement CsvShiftGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCsvShiftGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CsvShiftGrammarLexerInit() {
  staticData := &csvshiftgrammarlexerLexerStaticData
  staticData.once.Do(csvshiftgrammarlexerLexerInit)
}

// NewCsvShiftGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCsvShiftGrammarLexer(input antlr.CharStream) *CsvShiftGrammarLexer {
  CsvShiftGrammarLexerInit()
	l := new(CsvShiftGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &csvshiftgrammarlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CsvShiftGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CsvShiftGrammarLexer tokens.
const (
	CsvShiftGrammarLexerINPUT = 1
	CsvShiftGrammarLexerOUTPUT = 2
	CsvShiftGrammarLexerCOLUMN = 3
	CsvShiftGrammarLexerCOLUMNS = 4
	CsvShiftGrammarLexerTRIM = 5
	CsvShiftGrammarLexerREPLACE = 6
	CsvShiftGrammarLexerCOMBINE = 7
	CsvShiftGrammarLexerWITH = 8
	CsvShiftGrammarLexerAS = 9
	CsvShiftGrammarLexerCOMMA = 10
	CsvShiftGrammarLexerQUOTE = 11
	CsvShiftGrammarLexerARROW = 12
	CsvShiftGrammarLexerIDENTIFIER = 13
	CsvShiftGrammarLexerSTRING = 14
	CsvShiftGrammarLexerWS = 15
)

