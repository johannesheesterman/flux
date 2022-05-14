// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Lambda

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LambdaParser struct {
	*antlr.BaseParser
}

var lambdaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lambdaParserInit() {
	staticData := &lambdaParserStaticData
	staticData.literalNames = []string{
		"", "'='", "';'", "'('", "','", "')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN",
		"ID", "COMMENT", "KEY", "NEWLINE", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "stat", "func", "obj", "pair", "array", "value",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 123, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12, 0, 20,
		9, 0, 1, 0, 3, 0, 23, 8, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 33, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12, 2, 50, 9,
		2, 3, 2, 52, 8, 2, 1, 2, 1, 2, 3, 2, 56, 8, 2, 1, 3, 1, 3, 5, 3, 60, 8,
		3, 10, 3, 12, 3, 63, 9, 3, 1, 3, 1, 3, 5, 3, 67, 8, 3, 10, 3, 12, 3, 70,
		9, 3, 1, 3, 5, 3, 73, 8, 3, 10, 3, 12, 3, 76, 9, 3, 1, 3, 5, 3, 79, 8,
		3, 10, 3, 12, 3, 82, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 88, 8, 3, 10,
		3, 12, 3, 91, 9, 3, 1, 3, 3, 3, 94, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		100, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 106, 8, 5, 10, 5, 12, 5, 109,
		9, 5, 3, 5, 111, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 121, 8, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 139, 0, 24,
		1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 41, 1, 0, 0, 0, 6, 93, 1, 0, 0, 0, 8,
		95, 1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 22, 3, 2,
		1, 0, 15, 17, 5, 16, 0, 0, 16, 15, 1, 0, 0, 0, 17, 20, 1, 0, 0, 0, 18,
		16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 23, 1, 0, 0, 0, 20, 18, 1, 0, 0,
		0, 21, 23, 5, 0, 0, 1, 22, 18, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 25,
		1, 0, 0, 0, 24, 14, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0,
		26, 27, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28, 29, 5, 13, 0, 0, 29, 30, 5,
		1, 0, 0, 30, 32, 3, 12, 6, 0, 31, 33, 5, 2, 0, 0, 32, 31, 1, 0, 0, 0, 32,
		33, 1, 0, 0, 0, 33, 40, 1, 0, 0, 0, 34, 35, 5, 13, 0, 0, 35, 36, 5, 1,
		0, 0, 36, 40, 3, 4, 2, 0, 37, 40, 3, 4, 2, 0, 38, 40, 5, 14, 0, 0, 39,
		28, 1, 0, 0, 0, 39, 34, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0,
		0, 40, 3, 1, 0, 0, 0, 41, 42, 5, 15, 0, 0, 42, 51, 5, 3, 0, 0, 43, 48,
		3, 12, 6, 0, 44, 45, 5, 4, 0, 0, 45, 47, 3, 12, 6, 0, 46, 44, 1, 0, 0,
		0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 52,
		1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 43, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0,
		52, 53, 1, 0, 0, 0, 53, 55, 5, 5, 0, 0, 54, 56, 5, 2, 0, 0, 55, 54, 1,
		0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 5, 1, 0, 0, 0, 57, 61, 5, 6, 0, 0, 58,
		60, 5, 16, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0,
		0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 74,
		3, 8, 4, 0, 65, 67, 5, 16, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0,
		68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1,
		0, 0, 0, 71, 73, 3, 8, 4, 0, 72, 68, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 80, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 77, 79, 5, 16, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0,
		83, 84, 5, 7, 0, 0, 84, 94, 1, 0, 0, 0, 85, 89, 5, 6, 0, 0, 86, 88, 5,
		16, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89,
		90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 94, 5, 7, 0,
		0, 93, 57, 1, 0, 0, 0, 93, 85, 1, 0, 0, 0, 94, 7, 1, 0, 0, 0, 95, 96, 5,
		15, 0, 0, 96, 97, 5, 1, 0, 0, 97, 99, 3, 12, 6, 0, 98, 100, 5, 2, 0, 0,
		99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 9, 1, 0, 0, 0, 101, 110,
		5, 8, 0, 0, 102, 107, 3, 12, 6, 0, 103, 104, 5, 4, 0, 0, 104, 106, 3, 12,
		6, 0, 105, 103, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0,
		107, 108, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110,
		102, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113,
		5, 9, 0, 0, 113, 11, 1, 0, 0, 0, 114, 121, 5, 10, 0, 0, 115, 121, 5, 11,
		0, 0, 116, 121, 5, 12, 0, 0, 117, 121, 5, 13, 0, 0, 118, 121, 3, 6, 3,
		0, 119, 121, 3, 10, 5, 0, 120, 114, 1, 0, 0, 0, 120, 115, 1, 0, 0, 0, 120,
		116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 119,
		1, 0, 0, 0, 121, 13, 1, 0, 0, 0, 18, 18, 22, 26, 32, 39, 48, 51, 55, 61,
		68, 74, 80, 89, 93, 99, 107, 110, 120,
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

// LambdaParserInit initializes any static state used to implement LambdaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLambdaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LambdaParserInit() {
	staticData := &lambdaParserStaticData
	staticData.once.Do(lambdaParserInit)
}

// NewLambdaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLambdaParser(input antlr.TokenStream) *LambdaParser {
	LambdaParserInit()
	this := new(LambdaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lambdaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Lambda.g4"

	return this
}

// LambdaParser tokens.
const (
	LambdaParserEOF     = antlr.TokenEOF
	LambdaParserT__0    = 1
	LambdaParserT__1    = 2
	LambdaParserT__2    = 3
	LambdaParserT__3    = 4
	LambdaParserT__4    = 5
	LambdaParserT__5    = 6
	LambdaParserT__6    = 7
	LambdaParserT__7    = 8
	LambdaParserT__8    = 9
	LambdaParserSTRING  = 10
	LambdaParserNUMBER  = 11
	LambdaParserBOOLEAN = 12
	LambdaParserID      = 13
	LambdaParserCOMMENT = 14
	LambdaParserKEY     = 15
	LambdaParserNEWLINE = 16
	LambdaParserWS      = 17
)

// LambdaParser rules.
const (
	LambdaParserRULE_prog  = 0
	LambdaParserRULE_stat  = 1
	LambdaParserRULE_func  = 2
	LambdaParserRULE_obj   = 3
	LambdaParserRULE_pair  = 4
	LambdaParserRULE_array = 5
	LambdaParserRULE_value = 6
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
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

	return t.(IStatContext)
}

func (s *ProgContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserEOF)
}

func (s *ProgContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserEOF, i)
}

func (s *ProgContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserNEWLINE)
}

func (s *ProgContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserNEWLINE, i)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LambdaParserRULE_prog)
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserID)|(1<<LambdaParserCOMMENT)|(1<<LambdaParserKEY))) != 0) {
		{
			p.SetState(14)
			p.Stat()
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			p.SetState(18)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == LambdaParserNEWLINE {
				{
					p.SetState(15)
					p.Match(LambdaParserNEWLINE)
				}

				p.SetState(20)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case 2:
			{
				p.SetState(21)
				p.Match(LambdaParserEOF)
			}

		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallAssignmentStatementContext struct {
	*StatContext
}

func NewFunctionCallAssignmentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallAssignmentStatementContext {
	var p = new(FunctionCallAssignmentStatementContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *FunctionCallAssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallAssignmentStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(LambdaParserID, 0)
}

func (s *FunctionCallAssignmentStatementContext) Func() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FunctionCallAssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterFunctionCallAssignmentStatement(s)
	}
}

func (s *FunctionCallAssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitFunctionCallAssignmentStatement(s)
	}
}

func (s *FunctionCallAssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitFunctionCallAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommentStatementContext struct {
	*StatContext
}

func NewCommentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentStatementContext {
	var p = new(CommentStatementContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *CommentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentStatementContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(LambdaParserCOMMENT, 0)
}

func (s *CommentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterCommentStatement(s)
	}
}

func (s *CommentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitCommentStatement(s)
	}
}

func (s *CommentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitCommentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementContext struct {
	*StatContext
}

func NewAssignmentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(LambdaParserID, 0)
}

func (s *AssignmentStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallStatementContext struct {
	*StatContext
}

func NewFunctionCallStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *FunctionCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallStatementContext) Func() IFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *FunctionCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitFunctionCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LambdaParserRULE_stat)
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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(LambdaParserID)
		}
		{
			p.SetState(29)
			p.Match(LambdaParserT__0)
		}
		{
			p.SetState(30)
			p.Value()
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LambdaParserT__1 {
			{
				p.SetState(31)
				p.Match(LambdaParserT__1)
			}

		}

	case 2:
		localctx = NewFunctionCallAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(LambdaParserID)
		}
		{
			p.SetState(35)
			p.Match(LambdaParserT__0)
		}
		{
			p.SetState(36)
			p.Func()
		}

	case 3:
		localctx = NewFunctionCallStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.Func()
		}

	case 4:
		localctx = NewCommentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.Match(LambdaParserCOMMENT)
		}

	}

	return localctx
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_func
	return p
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) KEY() antlr.TerminalNode {
	return s.GetToken(LambdaParserKEY, 0)
}

func (s *FuncContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *FuncContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Func() (localctx IFuncContext) {
	this := p
	_ = this

	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LambdaParserRULE_func)
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
		p.SetState(41)
		p.Match(LambdaParserKEY)
	}
	{
		p.SetState(42)
		p.Match(LambdaParserT__2)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__5)|(1<<LambdaParserT__7)|(1<<LambdaParserSTRING)|(1<<LambdaParserNUMBER)|(1<<LambdaParserBOOLEAN)|(1<<LambdaParserID))) != 0 {
		{
			p.SetState(43)
			p.Value()
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserT__3 {
			{
				p.SetState(44)
				p.Match(LambdaParserT__3)
			}
			{
				p.SetState(45)
				p.Value()
			}

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(53)
		p.Match(LambdaParserT__4)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LambdaParserT__1 {
		{
			p.SetState(54)
			p.Match(LambdaParserT__1)
		}

	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *ObjContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserNEWLINE)
}

func (s *ObjContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserNEWLINE, i)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitObj(s)
	}
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LambdaParserRULE_obj)
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

	var _alt int

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(LambdaParserT__5)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserNEWLINE {
			{
				p.SetState(58)
				p.Match(LambdaParserNEWLINE)
			}

			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(64)
			p.Pair()
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(68)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == LambdaParserNEWLINE {
					{
						p.SetState(65)
						p.Match(LambdaParserNEWLINE)
					}

					p.SetState(70)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(71)
					p.Pair()
				}

			}
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserNEWLINE {
			{
				p.SetState(77)
				p.Match(LambdaParserNEWLINE)
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(83)
			p.Match(LambdaParserT__6)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(LambdaParserT__5)
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserNEWLINE {
			{
				p.SetState(86)
				p.Match(LambdaParserNEWLINE)
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(92)
			p.Match(LambdaParserT__6)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) KEY() antlr.TerminalNode {
	return s.GetToken(LambdaParserKEY, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LambdaParserRULE_pair)
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
		p.SetState(95)
		p.Match(LambdaParserKEY)
	}
	{
		p.SetState(96)
		p.Match(LambdaParserT__0)
	}
	{
		p.SetState(97)
		p.Value()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LambdaParserT__1 {
		{
			p.SetState(98)
			p.Match(LambdaParserT__1)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LambdaParserRULE_array)
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
		p.SetState(101)
		p.Match(LambdaParserT__7)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__5)|(1<<LambdaParserT__7)|(1<<LambdaParserSTRING)|(1<<LambdaParserNUMBER)|(1<<LambdaParserBOOLEAN)|(1<<LambdaParserID))) != 0 {
		{
			p.SetState(102)
			p.Value()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserT__3 {
			{
				p.SetState(103)
				p.Match(LambdaParserT__3)
			}
			{
				p.SetState(104)
				p.Value()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(112)
		p.Match(LambdaParserT__8)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(LambdaParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LambdaParserNUMBER, 0)
}

func (s *ValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(LambdaParserBOOLEAN, 0)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(LambdaParserID, 0)
}

func (s *ValueContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LambdaParserRULE_value)

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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(LambdaParserSTRING)
		}

	case LambdaParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(LambdaParserNUMBER)
		}

	case LambdaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Match(LambdaParserBOOLEAN)
		}

	case LambdaParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Match(LambdaParserID)
		}

	case LambdaParserT__5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.Obj()
		}

	case LambdaParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
