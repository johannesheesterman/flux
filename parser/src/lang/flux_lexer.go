// Code generated from Flux.g4 by ANTLR 4.10. DO NOT EDIT.

package lang

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

type FluxLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var fluxlexerLexerStaticData struct {
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

func fluxlexerLexerInit() {
	staticData := &fluxlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "';'", "'('", "','", "')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN",
		"ID", "COMMENT", "KEY", "NEWLINE", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"STRING", "NUMBER", "BOOLEAN", "INT", "ID", "COMMENT", "KEY", "NEWLINE",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 129, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 58, 8, 9, 10, 9, 12, 9, 61, 9, 9, 1, 9, 1, 9, 1, 10,
		3, 10, 66, 8, 10, 1, 10, 1, 10, 1, 10, 4, 10, 71, 8, 10, 11, 10, 12, 10,
		72, 3, 10, 75, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 86, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 91, 8, 12,
		10, 12, 12, 12, 94, 9, 12, 3, 12, 96, 8, 12, 1, 13, 1, 13, 4, 13, 100,
		8, 13, 11, 13, 12, 13, 101, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 108, 8,
		14, 10, 14, 12, 14, 111, 9, 14, 1, 15, 4, 15, 114, 8, 15, 11, 15, 12, 15,
		115, 1, 16, 3, 16, 119, 8, 16, 1, 16, 1, 16, 1, 17, 4, 17, 124, 8, 17,
		11, 17, 12, 17, 125, 1, 17, 1, 17, 2, 59, 101, 0, 18, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0,
		27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 1, 0, 5, 1, 0, 48, 57, 1, 0, 49,
		57, 2, 0, 10, 10, 13, 13, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 9, 9, 32,
		32, 139, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 39, 1, 0, 0,
		0, 5, 41, 1, 0, 0, 0, 7, 43, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 47, 1,
		0, 0, 0, 13, 49, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 53, 1, 0, 0, 0, 19,
		55, 1, 0, 0, 0, 21, 65, 1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 95, 1, 0, 0,
		0, 27, 97, 1, 0, 0, 0, 29, 105, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33, 118,
		1, 0, 0, 0, 35, 123, 1, 0, 0, 0, 37, 38, 5, 61, 0, 0, 38, 2, 1, 0, 0, 0,
		39, 40, 5, 59, 0, 0, 40, 4, 1, 0, 0, 0, 41, 42, 5, 40, 0, 0, 42, 6, 1,
		0, 0, 0, 43, 44, 5, 44, 0, 0, 44, 8, 1, 0, 0, 0, 45, 46, 5, 41, 0, 0, 46,
		10, 1, 0, 0, 0, 47, 48, 5, 123, 0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 5, 125,
		0, 0, 50, 14, 1, 0, 0, 0, 51, 52, 5, 91, 0, 0, 52, 16, 1, 0, 0, 0, 53,
		54, 5, 93, 0, 0, 54, 18, 1, 0, 0, 0, 55, 59, 5, 34, 0, 0, 56, 58, 9, 0,
		0, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5, 34, 0, 0,
		63, 20, 1, 0, 0, 0, 64, 66, 5, 45, 0, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 74, 3, 25, 12, 0, 68, 70, 5, 46, 0, 0,
		69, 71, 7, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 70, 1,
		0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0, 74,
		75, 1, 0, 0, 0, 75, 22, 1, 0, 0, 0, 76, 77, 5, 116, 0, 0, 77, 78, 5, 114,
		0, 0, 78, 79, 5, 117, 0, 0, 79, 86, 5, 101, 0, 0, 80, 81, 5, 102, 0, 0,
		81, 82, 5, 97, 0, 0, 82, 83, 5, 108, 0, 0, 83, 84, 5, 115, 0, 0, 84, 86,
		5, 101, 0, 0, 85, 76, 1, 0, 0, 0, 85, 80, 1, 0, 0, 0, 86, 24, 1, 0, 0,
		0, 87, 96, 5, 48, 0, 0, 88, 92, 7, 1, 0, 0, 89, 91, 7, 0, 0, 0, 90, 89,
		1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 95, 88, 1,
		0, 0, 0, 96, 26, 1, 0, 0, 0, 97, 99, 5, 91, 0, 0, 98, 100, 9, 0, 0, 0,
		99, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 101, 99,
		1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 5, 93, 0, 0, 104, 28, 1, 0,
		0, 0, 105, 109, 5, 35, 0, 0, 106, 108, 8, 2, 0, 0, 107, 106, 1, 0, 0, 0,
		108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		30, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 114, 7, 3, 0, 0, 113, 112, 1,
		0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 32, 1, 0, 0, 0, 117, 119, 5, 13, 0, 0, 118, 117, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 5, 10, 0, 0, 121, 34,
		1, 0, 0, 0, 122, 124, 7, 4, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0,
		0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0,
		127, 128, 6, 17, 0, 0, 128, 36, 1, 0, 0, 0, 13, 0, 59, 65, 72, 74, 85,
		92, 95, 101, 109, 115, 118, 125, 1, 6, 0, 0,
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

// FluxLexerInit initializes any static state used to implement FluxLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFluxLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FluxLexerInit() {
	staticData := &fluxlexerLexerStaticData
	staticData.once.Do(fluxlexerLexerInit)
}

// NewFluxLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFluxLexer(input antlr.CharStream) *FluxLexer {
	FluxLexerInit()
	l := new(FluxLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &fluxlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Flux.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FluxLexer tokens.
const (
	FluxLexerT__0    = 1
	FluxLexerT__1    = 2
	FluxLexerT__2    = 3
	FluxLexerT__3    = 4
	FluxLexerT__4    = 5
	FluxLexerT__5    = 6
	FluxLexerT__6    = 7
	FluxLexerT__7    = 8
	FluxLexerT__8    = 9
	FluxLexerSTRING  = 10
	FluxLexerNUMBER  = 11
	FluxLexerBOOLEAN = 12
	FluxLexerID      = 13
	FluxLexerCOMMENT = 14
	FluxLexerKEY     = 15
	FluxLexerNEWLINE = 16
	FluxLexerWS      = 17
)
