// Code generated from Predikit.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type PredikitLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PredikitLexerLexerStaticData struct {
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

func predikitlexerLexerInit() {
	staticData := &PredikitLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'all'", "'any'", "'none'", "'test'", "'not'", "'tool'", "'@'",
		"':'", "'$'", "'{'", "'}'", "'('", "')'", "'true'", "'false'", "'='",
		"'!='", "'=~'", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "PK_ALL", "PK_ANY", "PK_NONE", "PK_TEST", "PK_NOT", "PK_TOOL", "PK_RETRYING",
		"PK_COLON", "PK_DOLLAR", "PK_LCURLY", "PK_RCURLY", "PK_LPAREN", "PK_RPAREN",
		"PK_TRUE", "PK_FALSE", "PK_CMP_EQ", "PK_CMP_NEQ", "PK_CMP_RE", "PK_CMP_GT",
		"PK_CMP_GTE", "PK_CMP_LT", "PK_CMP_LTE", "PK_FN_LIT", "PK_STRING_LIT",
		"PK_ID", "PK_TYPENAME", "PK_INT", "LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"PK_ALL", "PK_ANY", "PK_NONE", "PK_TEST", "PK_NOT", "PK_TOOL", "PK_RETRYING",
		"PK_COLON", "PK_DOLLAR", "PK_LCURLY", "PK_RCURLY", "PK_LPAREN", "PK_RPAREN",
		"PK_TRUE", "PK_FALSE", "PK_CMP_EQ", "PK_CMP_NEQ", "PK_CMP_RE", "PK_CMP_GT",
		"PK_CMP_GTE", "PK_CMP_LT", "PK_CMP_LTE", "PK_FN_LIT", "PK_STRING_LIT",
		"ESC", "PK_ID", "PK_TYPENAME", "PK_INT", "DIGIT", "LOWER", "UPPER",
		"LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 229, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		5, 22, 142, 8, 22, 10, 22, 12, 22, 145, 9, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 5, 23, 152, 8, 23, 10, 23, 12, 23, 155, 9, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 163, 8, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 5, 25, 170, 8, 25, 10, 25, 12, 25, 173, 9, 25, 1, 26, 1, 26,
		1, 26, 4, 26, 178, 8, 26, 11, 26, 12, 26, 179, 1, 27, 4, 27, 183, 8, 27,
		11, 27, 12, 27, 184, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 5, 31, 197, 8, 31, 10, 31, 12, 31, 200, 9, 31, 1, 31,
		3, 31, 203, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 5, 32, 213, 8, 32, 10, 32, 12, 32, 216, 9, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 33, 4, 33, 224, 8, 33, 11, 33, 12, 33, 225, 1, 33, 1,
		33, 4, 143, 153, 198, 214, 0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		0, 51, 25, 53, 26, 55, 27, 57, 0, 59, 0, 61, 0, 63, 28, 65, 29, 67, 30,
		1, 0, 2, 4, 0, 33, 33, 46, 46, 63, 63, 95, 95, 3, 0, 9, 10, 13, 13, 32,
		32, 239, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1, 69,
		1, 0, 0, 0, 3, 73, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 82, 1, 0, 0, 0, 9,
		87, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 96, 1, 0, 0, 0, 15, 98, 1, 0, 0,
		0, 17, 100, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 104, 1, 0, 0, 0, 23, 106,
		1, 0, 0, 0, 25, 108, 1, 0, 0, 0, 27, 110, 1, 0, 0, 0, 29, 115, 1, 0, 0,
		0, 31, 121, 1, 0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 126, 1, 0, 0, 0, 37, 129,
		1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 134, 1, 0, 0, 0, 43, 136, 1, 0, 0,
		0, 45, 139, 1, 0, 0, 0, 47, 148, 1, 0, 0, 0, 49, 162, 1, 0, 0, 0, 51, 164,
		1, 0, 0, 0, 53, 174, 1, 0, 0, 0, 55, 182, 1, 0, 0, 0, 57, 186, 1, 0, 0,
		0, 59, 188, 1, 0, 0, 0, 61, 190, 1, 0, 0, 0, 63, 192, 1, 0, 0, 0, 65, 208,
		1, 0, 0, 0, 67, 223, 1, 0, 0, 0, 69, 70, 5, 97, 0, 0, 70, 71, 5, 108, 0,
		0, 71, 72, 5, 108, 0, 0, 72, 2, 1, 0, 0, 0, 73, 74, 5, 97, 0, 0, 74, 75,
		5, 110, 0, 0, 75, 76, 5, 121, 0, 0, 76, 4, 1, 0, 0, 0, 77, 78, 5, 110,
		0, 0, 78, 79, 5, 111, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81, 5, 101, 0, 0,
		81, 6, 1, 0, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5, 101, 0, 0, 84, 85,
		5, 115, 0, 0, 85, 86, 5, 116, 0, 0, 86, 8, 1, 0, 0, 0, 87, 88, 5, 110,
		0, 0, 88, 89, 5, 111, 0, 0, 89, 90, 5, 116, 0, 0, 90, 10, 1, 0, 0, 0, 91,
		92, 5, 116, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5,
		108, 0, 0, 95, 12, 1, 0, 0, 0, 96, 97, 5, 64, 0, 0, 97, 14, 1, 0, 0, 0,
		98, 99, 5, 58, 0, 0, 99, 16, 1, 0, 0, 0, 100, 101, 5, 36, 0, 0, 101, 18,
		1, 0, 0, 0, 102, 103, 5, 123, 0, 0, 103, 20, 1, 0, 0, 0, 104, 105, 5, 125,
		0, 0, 105, 22, 1, 0, 0, 0, 106, 107, 5, 40, 0, 0, 107, 24, 1, 0, 0, 0,
		108, 109, 5, 41, 0, 0, 109, 26, 1, 0, 0, 0, 110, 111, 5, 116, 0, 0, 111,
		112, 5, 114, 0, 0, 112, 113, 5, 117, 0, 0, 113, 114, 5, 101, 0, 0, 114,
		28, 1, 0, 0, 0, 115, 116, 5, 102, 0, 0, 116, 117, 5, 97, 0, 0, 117, 118,
		5, 108, 0, 0, 118, 119, 5, 115, 0, 0, 119, 120, 5, 101, 0, 0, 120, 30,
		1, 0, 0, 0, 121, 122, 5, 61, 0, 0, 122, 32, 1, 0, 0, 0, 123, 124, 5, 33,
		0, 0, 124, 125, 5, 61, 0, 0, 125, 34, 1, 0, 0, 0, 126, 127, 5, 61, 0, 0,
		127, 128, 5, 126, 0, 0, 128, 36, 1, 0, 0, 0, 129, 130, 5, 62, 0, 0, 130,
		38, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132, 133, 5, 61, 0, 0, 133, 40,
		1, 0, 0, 0, 134, 135, 5, 60, 0, 0, 135, 42, 1, 0, 0, 0, 136, 137, 5, 60,
		0, 0, 137, 138, 5, 61, 0, 0, 138, 44, 1, 0, 0, 0, 139, 143, 5, 40, 0, 0,
		140, 142, 9, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 143,
		1, 0, 0, 0, 146, 147, 5, 41, 0, 0, 147, 46, 1, 0, 0, 0, 148, 153, 5, 34,
		0, 0, 149, 152, 3, 49, 24, 0, 150, 152, 9, 0, 0, 0, 151, 149, 1, 0, 0,
		0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 153,
		151, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157,
		5, 34, 0, 0, 157, 48, 1, 0, 0, 0, 158, 159, 5, 92, 0, 0, 159, 163, 5, 34,
		0, 0, 160, 161, 5, 92, 0, 0, 161, 163, 5, 92, 0, 0, 162, 158, 1, 0, 0,
		0, 162, 160, 1, 0, 0, 0, 163, 50, 1, 0, 0, 0, 164, 171, 3, 59, 29, 0, 165,
		170, 3, 61, 30, 0, 166, 170, 3, 59, 29, 0, 167, 170, 3, 57, 28, 0, 168,
		170, 7, 0, 0, 0, 169, 165, 1, 0, 0, 0, 169, 166, 1, 0, 0, 0, 169, 167,
		1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 52, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0,
		174, 177, 3, 61, 30, 0, 175, 178, 3, 59, 29, 0, 176, 178, 3, 61, 30, 0,
		177, 175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 54, 1, 0, 0, 0, 181, 183, 3,
		57, 28, 0, 182, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 184, 185, 1, 0, 0, 0, 185, 56, 1, 0, 0, 0, 186, 187, 2, 48, 57, 0,
		187, 58, 1, 0, 0, 0, 188, 189, 2, 97, 122, 0, 189, 60, 1, 0, 0, 0, 190,
		191, 2, 65, 90, 0, 191, 62, 1, 0, 0, 0, 192, 193, 5, 47, 0, 0, 193, 194,
		5, 47, 0, 0, 194, 198, 1, 0, 0, 0, 195, 197, 9, 0, 0, 0, 196, 195, 1, 0,
		0, 0, 197, 200, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 203, 5, 13, 0, 0, 202,
		201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205,
		5, 10, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 6, 31, 0, 0, 207, 64, 1, 0,
		0, 0, 208, 209, 5, 47, 0, 0, 209, 210, 5, 42, 0, 0, 210, 214, 1, 0, 0,
		0, 211, 213, 9, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		215, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 217, 218, 5, 42, 0, 0, 218, 219, 5, 47, 0, 0, 219, 220, 1,
		0, 0, 0, 220, 221, 6, 32, 0, 0, 221, 66, 1, 0, 0, 0, 222, 224, 7, 1, 0,
		0, 223, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225,
		226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 6, 33, 0, 0, 228, 68,
		1, 0, 0, 0, 14, 0, 143, 151, 153, 162, 169, 171, 177, 179, 184, 198, 202,
		214, 225, 1, 6, 0, 0,
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

// PredikitLexerInit initializes any static state used to implement PredikitLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPredikitLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PredikitLexerInit() {
	staticData := &PredikitLexerLexerStaticData
	staticData.once.Do(predikitlexerLexerInit)
}

// NewPredikitLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPredikitLexer(input antlr.CharStream) *PredikitLexer {
	PredikitLexerInit()
	l := new(PredikitLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PredikitLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Predikit.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PredikitLexer tokens.
const (
	PredikitLexerPK_ALL        = 1
	PredikitLexerPK_ANY        = 2
	PredikitLexerPK_NONE       = 3
	PredikitLexerPK_TEST       = 4
	PredikitLexerPK_NOT        = 5
	PredikitLexerPK_TOOL       = 6
	PredikitLexerPK_RETRYING   = 7
	PredikitLexerPK_COLON      = 8
	PredikitLexerPK_DOLLAR     = 9
	PredikitLexerPK_LCURLY     = 10
	PredikitLexerPK_RCURLY     = 11
	PredikitLexerPK_LPAREN     = 12
	PredikitLexerPK_RPAREN     = 13
	PredikitLexerPK_TRUE       = 14
	PredikitLexerPK_FALSE      = 15
	PredikitLexerPK_CMP_EQ     = 16
	PredikitLexerPK_CMP_NEQ    = 17
	PredikitLexerPK_CMP_RE     = 18
	PredikitLexerPK_CMP_GT     = 19
	PredikitLexerPK_CMP_GTE    = 20
	PredikitLexerPK_CMP_LT     = 21
	PredikitLexerPK_CMP_LTE    = 22
	PredikitLexerPK_FN_LIT     = 23
	PredikitLexerPK_STRING_LIT = 24
	PredikitLexerPK_ID         = 25
	PredikitLexerPK_TYPENAME   = 26
	PredikitLexerPK_INT        = 27
	PredikitLexerLINE_COMMENT  = 28
	PredikitLexerCOMMENT       = 29
	PredikitLexerWS            = 30
)
