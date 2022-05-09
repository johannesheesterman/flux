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
		"", "'='", "'{'", "'}'", "';'", "'['", "','", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "ID",
		"COMMENT", "MODIFIER", "KEY", "NEWLINE", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "STRING", "NUMBER",
		"BOOLEAN", "INT", "ID", "COMMENT", "MODIFIER", "KEY", "NEWLINE", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 138, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 52, 8, 7, 10, 7, 12,
		7, 55, 9, 7, 1, 7, 1, 7, 1, 8, 3, 8, 60, 8, 8, 1, 8, 1, 8, 1, 8, 4, 8,
		65, 8, 8, 11, 8, 12, 8, 66, 3, 8, 69, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 80, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10,
		85, 8, 10, 10, 10, 12, 10, 88, 9, 10, 3, 10, 90, 8, 10, 1, 11, 1, 11, 4,
		11, 94, 8, 11, 11, 11, 12, 11, 95, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 102,
		8, 12, 10, 12, 12, 12, 105, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 120, 8, 13,
		1, 14, 4, 14, 123, 8, 14, 11, 14, 12, 14, 124, 1, 15, 3, 15, 128, 8, 15,
		1, 15, 1, 15, 1, 16, 4, 16, 133, 8, 16, 11, 16, 12, 16, 134, 1, 16, 1,
		16, 2, 53, 95, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 0, 23, 11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 1,
		0, 5, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 10, 10, 13, 13, 3, 0, 48, 57, 65,
		90, 97, 122, 2, 0, 9, 9, 32, 32, 149, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37,
		1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0, 11,
		45, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 59, 1, 0, 0,
		0, 19, 79, 1, 0, 0, 0, 21, 89, 1, 0, 0, 0, 23, 91, 1, 0, 0, 0, 25, 99,
		1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 122, 1, 0, 0, 0, 31, 127, 1, 0, 0,
		0, 33, 132, 1, 0, 0, 0, 35, 36, 5, 61, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38,
		5, 123, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 125, 0, 0, 40, 6, 1, 0, 0,
		0, 41, 42, 5, 59, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 91, 0, 0, 44, 10,
		1, 0, 0, 0, 45, 46, 5, 44, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 93, 0,
		0, 48, 14, 1, 0, 0, 0, 49, 53, 5, 34, 0, 0, 50, 52, 9, 0, 0, 0, 51, 50,
		1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0,
		54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 34, 0, 0, 57, 16, 1,
		0, 0, 0, 58, 60, 5, 45, 0, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60,
		61, 1, 0, 0, 0, 61, 68, 3, 21, 10, 0, 62, 64, 5, 46, 0, 0, 63, 65, 7, 0,
		0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 18, 1, 0, 0, 0, 70, 71, 5, 116, 0, 0, 71, 72, 5, 114, 0, 0, 72, 73,
		5, 117, 0, 0, 73, 80, 5, 101, 0, 0, 74, 75, 5, 102, 0, 0, 75, 76, 5, 97,
		0, 0, 76, 77, 5, 108, 0, 0, 77, 78, 5, 115, 0, 0, 78, 80, 5, 101, 0, 0,
		79, 70, 1, 0, 0, 0, 79, 74, 1, 0, 0, 0, 80, 20, 1, 0, 0, 0, 81, 90, 5,
		48, 0, 0, 82, 86, 7, 1, 0, 0, 83, 85, 7, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85,
		88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 90, 1, 0, 0,
		0, 88, 86, 1, 0, 0, 0, 89, 81, 1, 0, 0, 0, 89, 82, 1, 0, 0, 0, 90, 22,
		1, 0, 0, 0, 91, 93, 5, 91, 0, 0, 92, 94, 9, 0, 0, 0, 93, 92, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 1,
		0, 0, 0, 97, 98, 5, 93, 0, 0, 98, 24, 1, 0, 0, 0, 99, 103, 5, 35, 0, 0,
		100, 102, 8, 2, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 26, 1, 0, 0, 0, 105, 103, 1,
		0, 0, 0, 106, 107, 5, 112, 0, 0, 107, 108, 5, 117, 0, 0, 108, 109, 5, 98,
		0, 0, 109, 110, 5, 108, 0, 0, 110, 111, 5, 105, 0, 0, 111, 120, 5, 99,
		0, 0, 112, 113, 5, 112, 0, 0, 113, 114, 5, 114, 0, 0, 114, 115, 5, 105,
		0, 0, 115, 116, 5, 118, 0, 0, 116, 117, 5, 97, 0, 0, 117, 118, 5, 116,
		0, 0, 118, 120, 5, 101, 0, 0, 119, 106, 1, 0, 0, 0, 119, 112, 1, 0, 0,
		0, 120, 28, 1, 0, 0, 0, 121, 123, 7, 3, 0, 0, 122, 121, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 30, 1,
		0, 0, 0, 126, 128, 5, 13, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0,
		0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 10, 0, 0, 130, 32, 1, 0, 0, 0, 131,
		133, 7, 4, 0, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 6, 16,
		0, 0, 137, 34, 1, 0, 0, 0, 14, 0, 53, 59, 66, 68, 79, 86, 89, 95, 103,
		119, 124, 127, 134, 1, 6, 0, 0,
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
	LambdaLexerT__0     = 1
	LambdaLexerT__1     = 2
	LambdaLexerT__2     = 3
	LambdaLexerT__3     = 4
	LambdaLexerT__4     = 5
	LambdaLexerT__5     = 6
	LambdaLexerT__6     = 7
	LambdaLexerSTRING   = 8
	LambdaLexerNUMBER   = 9
	LambdaLexerBOOLEAN  = 10
	LambdaLexerID       = 11
	LambdaLexerCOMMENT  = 12
	LambdaLexerMODIFIER = 13
	LambdaLexerKEY      = 14
	LambdaLexerNEWLINE  = 15
	LambdaLexerWS       = 16
)
