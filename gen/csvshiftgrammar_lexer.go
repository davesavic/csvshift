// Code generated from /home/sav/Development/Go/csvshift/grammar/CsvShiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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

var CsvShiftGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csvshiftgrammarlexerLexerInit() {
	staticData := &CsvShiftGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Input Columns'", "'Output Columns'", "'Column'", "'Columns'",
		"'-> Trim'", "'-> Replace'", "'-> Join with'", "'-> ToLower'", "'-> ToUpper'",
		"'-> Split on'", "'-> RegexReplace'", "'-> RegexExtract'", "'with'",
		"'as'", "','", "'->'", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "JOIN",
		"LOWER", "UPPER", "SPLIT", "REGEXREPLACE", "REGEXEXTRACT", "WITH", "AS",
		"COMMA", "ARROW", "QUOTE", "ESCAPE_SEQUENCE", "STRING", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"INPUT", "OUTPUT", "COLUMN", "COLUMNS", "TRIM", "REPLACE", "JOIN", "LOWER",
		"UPPER", "SPLIT", "REGEXREPLACE", "REGEXEXTRACT", "WITH", "AS", "COMMA",
		"ARROW", "QUOTE", "ESCAPE_SEQUENCE", "STRING", "IDENTIFIER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 227, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 5, 18, 207, 8, 18, 10, 18, 12, 18, 210, 9, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 5, 19, 216, 8, 19, 10, 19, 12, 19, 219, 9, 19,
		1, 20, 4, 20, 222, 8, 20, 11, 20, 12, 20, 223, 1, 20, 1, 20, 0, 0, 21,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 1, 0, 4, 2, 0, 34, 34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 230, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 72,
		1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 87, 1, 0, 0, 0, 11, 95, 1, 0, 0, 0, 13,
		106, 1, 0, 0, 0, 15, 119, 1, 0, 0, 0, 17, 130, 1, 0, 0, 0, 19, 141, 1,
		0, 0, 0, 21, 153, 1, 0, 0, 0, 23, 169, 1, 0, 0, 0, 25, 185, 1, 0, 0, 0,
		27, 190, 1, 0, 0, 0, 29, 193, 1, 0, 0, 0, 31, 195, 1, 0, 0, 0, 33, 198,
		1, 0, 0, 0, 35, 200, 1, 0, 0, 0, 37, 203, 1, 0, 0, 0, 39, 213, 1, 0, 0,
		0, 41, 221, 1, 0, 0, 0, 43, 44, 5, 73, 0, 0, 44, 45, 5, 110, 0, 0, 45,
		46, 5, 112, 0, 0, 46, 47, 5, 117, 0, 0, 47, 48, 5, 116, 0, 0, 48, 49, 5,
		32, 0, 0, 49, 50, 5, 67, 0, 0, 50, 51, 5, 111, 0, 0, 51, 52, 5, 108, 0,
		0, 52, 53, 5, 117, 0, 0, 53, 54, 5, 109, 0, 0, 54, 55, 5, 110, 0, 0, 55,
		56, 5, 115, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 79, 0, 0, 58, 59, 5, 117,
		0, 0, 59, 60, 5, 116, 0, 0, 60, 61, 5, 112, 0, 0, 61, 62, 5, 117, 0, 0,
		62, 63, 5, 116, 0, 0, 63, 64, 5, 32, 0, 0, 64, 65, 5, 67, 0, 0, 65, 66,
		5, 111, 0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 117, 0, 0, 68, 69, 5, 109,
		0, 0, 69, 70, 5, 110, 0, 0, 70, 71, 5, 115, 0, 0, 71, 4, 1, 0, 0, 0, 72,
		73, 5, 67, 0, 0, 73, 74, 5, 111, 0, 0, 74, 75, 5, 108, 0, 0, 75, 76, 5,
		117, 0, 0, 76, 77, 5, 109, 0, 0, 77, 78, 5, 110, 0, 0, 78, 6, 1, 0, 0,
		0, 79, 80, 5, 67, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 108, 0, 0, 82,
		83, 5, 117, 0, 0, 83, 84, 5, 109, 0, 0, 84, 85, 5, 110, 0, 0, 85, 86, 5,
		115, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88, 5, 45, 0, 0, 88, 89, 5, 62, 0, 0,
		89, 90, 5, 32, 0, 0, 90, 91, 5, 84, 0, 0, 91, 92, 5, 114, 0, 0, 92, 93,
		5, 105, 0, 0, 93, 94, 5, 109, 0, 0, 94, 10, 1, 0, 0, 0, 95, 96, 5, 45,
		0, 0, 96, 97, 5, 62, 0, 0, 97, 98, 5, 32, 0, 0, 98, 99, 5, 82, 0, 0, 99,
		100, 5, 101, 0, 0, 100, 101, 5, 112, 0, 0, 101, 102, 5, 108, 0, 0, 102,
		103, 5, 97, 0, 0, 103, 104, 5, 99, 0, 0, 104, 105, 5, 101, 0, 0, 105, 12,
		1, 0, 0, 0, 106, 107, 5, 45, 0, 0, 107, 108, 5, 62, 0, 0, 108, 109, 5,
		32, 0, 0, 109, 110, 5, 74, 0, 0, 110, 111, 5, 111, 0, 0, 111, 112, 5, 105,
		0, 0, 112, 113, 5, 110, 0, 0, 113, 114, 5, 32, 0, 0, 114, 115, 5, 119,
		0, 0, 115, 116, 5, 105, 0, 0, 116, 117, 5, 116, 0, 0, 117, 118, 5, 104,
		0, 0, 118, 14, 1, 0, 0, 0, 119, 120, 5, 45, 0, 0, 120, 121, 5, 62, 0, 0,
		121, 122, 5, 32, 0, 0, 122, 123, 5, 84, 0, 0, 123, 124, 5, 111, 0, 0, 124,
		125, 5, 76, 0, 0, 125, 126, 5, 111, 0, 0, 126, 127, 5, 119, 0, 0, 127,
		128, 5, 101, 0, 0, 128, 129, 5, 114, 0, 0, 129, 16, 1, 0, 0, 0, 130, 131,
		5, 45, 0, 0, 131, 132, 5, 62, 0, 0, 132, 133, 5, 32, 0, 0, 133, 134, 5,
		84, 0, 0, 134, 135, 5, 111, 0, 0, 135, 136, 5, 85, 0, 0, 136, 137, 5, 112,
		0, 0, 137, 138, 5, 112, 0, 0, 138, 139, 5, 101, 0, 0, 139, 140, 5, 114,
		0, 0, 140, 18, 1, 0, 0, 0, 141, 142, 5, 45, 0, 0, 142, 143, 5, 62, 0, 0,
		143, 144, 5, 32, 0, 0, 144, 145, 5, 83, 0, 0, 145, 146, 5, 112, 0, 0, 146,
		147, 5, 108, 0, 0, 147, 148, 5, 105, 0, 0, 148, 149, 5, 116, 0, 0, 149,
		150, 5, 32, 0, 0, 150, 151, 5, 111, 0, 0, 151, 152, 5, 110, 0, 0, 152,
		20, 1, 0, 0, 0, 153, 154, 5, 45, 0, 0, 154, 155, 5, 62, 0, 0, 155, 156,
		5, 32, 0, 0, 156, 157, 5, 82, 0, 0, 157, 158, 5, 101, 0, 0, 158, 159, 5,
		103, 0, 0, 159, 160, 5, 101, 0, 0, 160, 161, 5, 120, 0, 0, 161, 162, 5,
		82, 0, 0, 162, 163, 5, 101, 0, 0, 163, 164, 5, 112, 0, 0, 164, 165, 5,
		108, 0, 0, 165, 166, 5, 97, 0, 0, 166, 167, 5, 99, 0, 0, 167, 168, 5, 101,
		0, 0, 168, 22, 1, 0, 0, 0, 169, 170, 5, 45, 0, 0, 170, 171, 5, 62, 0, 0,
		171, 172, 5, 32, 0, 0, 172, 173, 5, 82, 0, 0, 173, 174, 5, 101, 0, 0, 174,
		175, 5, 103, 0, 0, 175, 176, 5, 101, 0, 0, 176, 177, 5, 120, 0, 0, 177,
		178, 5, 69, 0, 0, 178, 179, 5, 120, 0, 0, 179, 180, 5, 116, 0, 0, 180,
		181, 5, 114, 0, 0, 181, 182, 5, 97, 0, 0, 182, 183, 5, 99, 0, 0, 183, 184,
		5, 116, 0, 0, 184, 24, 1, 0, 0, 0, 185, 186, 5, 119, 0, 0, 186, 187, 5,
		105, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 104, 0, 0, 189, 26, 1,
		0, 0, 0, 190, 191, 5, 97, 0, 0, 191, 192, 5, 115, 0, 0, 192, 28, 1, 0,
		0, 0, 193, 194, 5, 44, 0, 0, 194, 30, 1, 0, 0, 0, 195, 196, 5, 45, 0, 0,
		196, 197, 5, 62, 0, 0, 197, 32, 1, 0, 0, 0, 198, 199, 5, 34, 0, 0, 199,
		34, 1, 0, 0, 0, 200, 201, 5, 92, 0, 0, 201, 202, 9, 0, 0, 0, 202, 36, 1,
		0, 0, 0, 203, 208, 3, 33, 16, 0, 204, 207, 3, 35, 17, 0, 205, 207, 8, 0,
		0, 0, 206, 204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0,
		208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210,
		208, 1, 0, 0, 0, 211, 212, 3, 33, 16, 0, 212, 38, 1, 0, 0, 0, 213, 217,
		7, 1, 0, 0, 214, 216, 7, 2, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0,
		0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 40, 1, 0, 0, 0,
		219, 217, 1, 0, 0, 0, 220, 222, 7, 3, 0, 0, 221, 220, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 226, 6, 20, 0, 0, 226, 42, 1, 0, 0, 0, 5, 0, 206, 208,
		217, 223, 1, 6, 0, 0,
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
	staticData := &CsvShiftGrammarLexerLexerStaticData
	staticData.once.Do(csvshiftgrammarlexerLexerInit)
}

// NewCsvShiftGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCsvShiftGrammarLexer(input antlr.CharStream) *CsvShiftGrammarLexer {
	CsvShiftGrammarLexerInit()
	l := new(CsvShiftGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CsvShiftGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CsvShiftGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CsvShiftGrammarLexer tokens.
const (
	CsvShiftGrammarLexerINPUT           = 1
	CsvShiftGrammarLexerOUTPUT          = 2
	CsvShiftGrammarLexerCOLUMN          = 3
	CsvShiftGrammarLexerCOLUMNS         = 4
	CsvShiftGrammarLexerTRIM            = 5
	CsvShiftGrammarLexerREPLACE         = 6
	CsvShiftGrammarLexerJOIN            = 7
	CsvShiftGrammarLexerLOWER           = 8
	CsvShiftGrammarLexerUPPER           = 9
	CsvShiftGrammarLexerSPLIT           = 10
	CsvShiftGrammarLexerREGEXREPLACE    = 11
	CsvShiftGrammarLexerREGEXEXTRACT    = 12
	CsvShiftGrammarLexerWITH            = 13
	CsvShiftGrammarLexerAS              = 14
	CsvShiftGrammarLexerCOMMA           = 15
	CsvShiftGrammarLexerARROW           = 16
	CsvShiftGrammarLexerQUOTE           = 17
	CsvShiftGrammarLexerESCAPE_SEQUENCE = 18
	CsvShiftGrammarLexerSTRING          = 19
	CsvShiftGrammarLexerIDENTIFIER      = 20
	CsvShiftGrammarLexerWS              = 21
)
