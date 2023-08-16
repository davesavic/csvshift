// Code generated from /home/sav/Development/Go/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
	modeNames    []string
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
		"'-> Trim'", "'-> Replace'", "'-> Join with'", "'-> ToLower'", "'-> ToUpper'",
		"'-> Split on'", "'with'", "'as'", "','", "'\"'", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "JOIN",
		"LOWER", "UPPER", "SPLIT", "WITH", "AS", "COMMA", "QUOTE", "ARROW",
		"IDENTIFIER", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "JOIN", "LOWER",
		"UPPER", "SPLIT", "WITH", "AS", "COMMA", "QUOTE", "ARROW", "IDENTIFIER",
		"STRING", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 185, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 5, 15, 165, 8, 15, 10, 15, 12, 15, 168, 9, 15, 1,
		16, 1, 16, 5, 16, 172, 8, 16, 10, 16, 12, 16, 175, 9, 16, 1, 16, 1, 16,
		1, 17, 4, 17, 180, 8, 17, 11, 17, 12, 17, 181, 1, 17, 1, 17, 0, 0, 18,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 1, 0, 4, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 10,
		10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 187, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 66,
		1, 0, 0, 0, 7, 73, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0, 11, 89, 1, 0, 0, 0, 13,
		100, 1, 0, 0, 0, 15, 113, 1, 0, 0, 0, 17, 124, 1, 0, 0, 0, 19, 135, 1,
		0, 0, 0, 21, 147, 1, 0, 0, 0, 23, 152, 1, 0, 0, 0, 25, 155, 1, 0, 0, 0,
		27, 157, 1, 0, 0, 0, 29, 159, 1, 0, 0, 0, 31, 162, 1, 0, 0, 0, 33, 169,
		1, 0, 0, 0, 35, 179, 1, 0, 0, 0, 37, 38, 5, 73, 0, 0, 38, 39, 5, 110, 0,
		0, 39, 40, 5, 112, 0, 0, 40, 41, 5, 117, 0, 0, 41, 42, 5, 116, 0, 0, 42,
		43, 5, 32, 0, 0, 43, 44, 5, 67, 0, 0, 44, 45, 5, 111, 0, 0, 45, 46, 5,
		108, 0, 0, 46, 47, 5, 117, 0, 0, 47, 48, 5, 109, 0, 0, 48, 49, 5, 110,
		0, 0, 49, 50, 5, 115, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 79, 0, 0, 52,
		53, 5, 117, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 112, 0, 0, 55, 56, 5,
		117, 0, 0, 56, 57, 5, 116, 0, 0, 57, 58, 5, 32, 0, 0, 58, 59, 5, 67, 0,
		0, 59, 60, 5, 111, 0, 0, 60, 61, 5, 108, 0, 0, 61, 62, 5, 117, 0, 0, 62,
		63, 5, 109, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65, 5, 115, 0, 0, 65, 4, 1,
		0, 0, 0, 66, 67, 5, 67, 0, 0, 67, 68, 5, 111, 0, 0, 68, 69, 5, 108, 0,
		0, 69, 70, 5, 117, 0, 0, 70, 71, 5, 109, 0, 0, 71, 72, 5, 110, 0, 0, 72,
		6, 1, 0, 0, 0, 73, 74, 5, 67, 0, 0, 74, 75, 5, 111, 0, 0, 75, 76, 5, 108,
		0, 0, 76, 77, 5, 117, 0, 0, 77, 78, 5, 109, 0, 0, 78, 79, 5, 110, 0, 0,
		79, 80, 5, 115, 0, 0, 80, 8, 1, 0, 0, 0, 81, 82, 5, 45, 0, 0, 82, 83, 5,
		62, 0, 0, 83, 84, 5, 32, 0, 0, 84, 85, 5, 84, 0, 0, 85, 86, 5, 114, 0,
		0, 86, 87, 5, 105, 0, 0, 87, 88, 5, 109, 0, 0, 88, 10, 1, 0, 0, 0, 89,
		90, 5, 45, 0, 0, 90, 91, 5, 62, 0, 0, 91, 92, 5, 32, 0, 0, 92, 93, 5, 82,
		0, 0, 93, 94, 5, 101, 0, 0, 94, 95, 5, 112, 0, 0, 95, 96, 5, 108, 0, 0,
		96, 97, 5, 97, 0, 0, 97, 98, 5, 99, 0, 0, 98, 99, 5, 101, 0, 0, 99, 12,
		1, 0, 0, 0, 100, 101, 5, 45, 0, 0, 101, 102, 5, 62, 0, 0, 102, 103, 5,
		32, 0, 0, 103, 104, 5, 74, 0, 0, 104, 105, 5, 111, 0, 0, 105, 106, 5, 105,
		0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5, 32, 0, 0, 108, 109, 5, 119,
		0, 0, 109, 110, 5, 105, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 104,
		0, 0, 112, 14, 1, 0, 0, 0, 113, 114, 5, 45, 0, 0, 114, 115, 5, 62, 0, 0,
		115, 116, 5, 32, 0, 0, 116, 117, 5, 84, 0, 0, 117, 118, 5, 111, 0, 0, 118,
		119, 5, 76, 0, 0, 119, 120, 5, 111, 0, 0, 120, 121, 5, 119, 0, 0, 121,
		122, 5, 101, 0, 0, 122, 123, 5, 114, 0, 0, 123, 16, 1, 0, 0, 0, 124, 125,
		5, 45, 0, 0, 125, 126, 5, 62, 0, 0, 126, 127, 5, 32, 0, 0, 127, 128, 5,
		84, 0, 0, 128, 129, 5, 111, 0, 0, 129, 130, 5, 85, 0, 0, 130, 131, 5, 112,
		0, 0, 131, 132, 5, 112, 0, 0, 132, 133, 5, 101, 0, 0, 133, 134, 5, 114,
		0, 0, 134, 18, 1, 0, 0, 0, 135, 136, 5, 45, 0, 0, 136, 137, 5, 62, 0, 0,
		137, 138, 5, 32, 0, 0, 138, 139, 5, 83, 0, 0, 139, 140, 5, 112, 0, 0, 140,
		141, 5, 108, 0, 0, 141, 142, 5, 105, 0, 0, 142, 143, 5, 116, 0, 0, 143,
		144, 5, 32, 0, 0, 144, 145, 5, 111, 0, 0, 145, 146, 5, 110, 0, 0, 146,
		20, 1, 0, 0, 0, 147, 148, 5, 119, 0, 0, 148, 149, 5, 105, 0, 0, 149, 150,
		5, 116, 0, 0, 150, 151, 5, 104, 0, 0, 151, 22, 1, 0, 0, 0, 152, 153, 5,
		97, 0, 0, 153, 154, 5, 115, 0, 0, 154, 24, 1, 0, 0, 0, 155, 156, 5, 44,
		0, 0, 156, 26, 1, 0, 0, 0, 157, 158, 5, 34, 0, 0, 158, 28, 1, 0, 0, 0,
		159, 160, 5, 45, 0, 0, 160, 161, 5, 62, 0, 0, 161, 30, 1, 0, 0, 0, 162,
		166, 7, 0, 0, 0, 163, 165, 7, 1, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168,
		1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 32, 1, 0,
		0, 0, 168, 166, 1, 0, 0, 0, 169, 173, 3, 27, 13, 0, 170, 172, 8, 2, 0,
		0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 177,
		3, 27, 13, 0, 177, 34, 1, 0, 0, 0, 178, 180, 7, 3, 0, 0, 179, 178, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0,
		182, 183, 1, 0, 0, 0, 183, 184, 6, 17, 0, 0, 184, 36, 1, 0, 0, 0, 4, 0,
		166, 173, 181, 1, 6, 0, 0,
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
	CsvShiftGrammarLexerINPUT      = 1
	CsvShiftGrammarLexerOUTPUT     = 2
	CsvShiftGrammarLexerCOLUMN     = 3
	CsvShiftGrammarLexerCOLUMNS    = 4
	CsvShiftGrammarLexerTRIM       = 5
	CsvShiftGrammarLexerREPLACE    = 6
	CsvShiftGrammarLexerJOIN       = 7
	CsvShiftGrammarLexerLOWER      = 8
	CsvShiftGrammarLexerUPPER      = 9
	CsvShiftGrammarLexerSPLIT      = 10
	CsvShiftGrammarLexerWITH       = 11
	CsvShiftGrammarLexerAS         = 12
	CsvShiftGrammarLexerCOMMA      = 13
	CsvShiftGrammarLexerQUOTE      = 14
	CsvShiftGrammarLexerARROW      = 15
	CsvShiftGrammarLexerIDENTIFIER = 16
	CsvShiftGrammarLexerSTRING     = 17
	CsvShiftGrammarLexerWS         = 18
)
