// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LambdaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lambdalexerLexerStaticData struct {
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

func lambdalexerLexerInit() {
	staticData := &lambdalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "';'", "'{'", "'}'", "'['", "','", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "COMMENT",
		"KEY", "NEWLINE", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "STRING", "NUMBER",
		"BOOLEAN", "INT", "COMMENT", "KEY", "NEWLINE", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 111, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 5, 7, 48, 8, 7, 10, 7, 12, 7, 51, 9, 7, 1, 7, 1, 7, 1,
		8, 3, 8, 56, 8, 8, 1, 8, 1, 8, 1, 8, 4, 8, 61, 8, 8, 11, 8, 12, 8, 62,
		3, 8, 65, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 76, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10, 81, 8, 10, 10, 10, 12, 10, 84,
		9, 10, 3, 10, 86, 8, 10, 1, 11, 1, 11, 5, 11, 90, 8, 11, 10, 11, 12, 11,
		93, 9, 11, 1, 12, 4, 12, 96, 8, 12, 11, 12, 12, 12, 97, 1, 13, 3, 13, 101,
		8, 13, 1, 13, 1, 13, 1, 14, 4, 14, 106, 8, 14, 11, 14, 12, 14, 107, 1,
		14, 1, 14, 1, 49, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 0, 23, 11, 25, 12, 27, 13, 29, 14, 1, 0, 5, 1, 0,
		48, 57, 1, 0, 49, 57, 2, 0, 10, 10, 13, 13, 3, 0, 48, 57, 65, 90, 97, 122,
		2, 0, 9, 9, 32, 32, 120, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1,
		31, 1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 37, 1, 0, 0, 0,
		9, 39, 1, 0, 0, 0, 11, 41, 1, 0, 0, 0, 13, 43, 1, 0, 0, 0, 15, 45, 1, 0,
		0, 0, 17, 55, 1, 0, 0, 0, 19, 75, 1, 0, 0, 0, 21, 85, 1, 0, 0, 0, 23, 87,
		1, 0, 0, 0, 25, 95, 1, 0, 0, 0, 27, 100, 1, 0, 0, 0, 29, 105, 1, 0, 0,
		0, 31, 32, 5, 61, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 5, 59, 0, 0, 34, 4,
		1, 0, 0, 0, 35, 36, 5, 123, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 5, 125, 0,
		0, 38, 8, 1, 0, 0, 0, 39, 40, 5, 91, 0, 0, 40, 10, 1, 0, 0, 0, 41, 42,
		5, 44, 0, 0, 42, 12, 1, 0, 0, 0, 43, 44, 5, 93, 0, 0, 44, 14, 1, 0, 0,
		0, 45, 49, 5, 34, 0, 0, 46, 48, 9, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 52, 53, 5, 34, 0, 0, 53, 16, 1, 0, 0, 0, 54, 56, 5,
		45, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57,
		64, 3, 21, 10, 0, 58, 60, 5, 46, 0, 0, 59, 61, 7, 0, 0, 0, 60, 59, 1, 0,
		0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65,
		1, 0, 0, 0, 64, 58, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 18, 1, 0, 0, 0,
		66, 67, 5, 116, 0, 0, 67, 68, 5, 114, 0, 0, 68, 69, 5, 117, 0, 0, 69, 76,
		5, 101, 0, 0, 70, 71, 5, 102, 0, 0, 71, 72, 5, 97, 0, 0, 72, 73, 5, 108,
		0, 0, 73, 74, 5, 115, 0, 0, 74, 76, 5, 101, 0, 0, 75, 66, 1, 0, 0, 0, 75,
		70, 1, 0, 0, 0, 76, 20, 1, 0, 0, 0, 77, 86, 5, 48, 0, 0, 78, 82, 7, 1,
		0, 0, 79, 81, 7, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0,
		85, 77, 1, 0, 0, 0, 85, 78, 1, 0, 0, 0, 86, 22, 1, 0, 0, 0, 87, 91, 5,
		35, 0, 0, 88, 90, 8, 2, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91,
		89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 24, 1, 0, 0, 0, 93, 91, 1, 0, 0,
		0, 94, 96, 7, 3, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 26, 1, 0, 0, 0, 99, 101, 5, 13, 0,
		0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102,
		103, 5, 10, 0, 0, 103, 28, 1, 0, 0, 0, 104, 106, 7, 4, 0, 0, 105, 104,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 6, 14, 0, 0, 110, 30, 1, 0, 0, 0,
		12, 0, 49, 55, 62, 64, 75, 82, 85, 91, 97, 100, 107, 1, 6, 0, 0,
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

// LambdaLexerInit initializes any static state used to implement LambdaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLambdaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LambdaLexerInit() {
	staticData := &lambdalexerLexerStaticData
	staticData.once.Do(lambdalexerLexerInit)
}

// NewLambdaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLambdaLexer(input antlr.CharStream) *LambdaLexer {
	LambdaLexerInit()
	l := new(LambdaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &lambdalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Lambda.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LambdaLexer tokens.
const (
	LambdaLexerT__0    = 1
	LambdaLexerT__1    = 2
	LambdaLexerT__2    = 3
	LambdaLexerT__3    = 4
	LambdaLexerT__4    = 5
	LambdaLexerT__5    = 6
	LambdaLexerT__6    = 7
	LambdaLexerSTRING  = 8
	LambdaLexerNUMBER  = 9
	LambdaLexerBOOLEAN = 10
	LambdaLexerCOMMENT = 11
	LambdaLexerKEY     = 12
	LambdaLexerNEWLINE = 13
	LambdaLexerWS      = 14
)
