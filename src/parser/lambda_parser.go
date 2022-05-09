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
		"", "'='", "';'", "'{'", "'}'", "'['", "','", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "STRING", "NUMBER", "BOOLEAN", "COMMENT",
		"KEY", "NEWLINE", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "stat", "obj", "pair", "array", "value", "id",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 112, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12, 0, 20,
		9, 0, 1, 0, 3, 0, 23, 8, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 33, 8, 1, 1, 1, 3, 1, 36, 8, 1, 1, 2, 1, 2, 5, 2,
		40, 8, 2, 10, 2, 12, 2, 43, 9, 2, 1, 2, 1, 2, 5, 2, 47, 8, 2, 10, 2, 12,
		2, 50, 9, 2, 1, 2, 5, 2, 53, 8, 2, 10, 2, 12, 2, 56, 9, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 62, 8, 2, 10, 2, 12, 2, 65, 9, 2, 1, 2, 3, 2, 68, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 74, 8, 3, 1, 3, 5, 3, 77, 8, 3, 10, 3,
		12, 3, 80, 9, 3, 3, 3, 82, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 88, 8, 4,
		10, 4, 12, 4, 91, 9, 4, 3, 4, 93, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 102, 8, 5, 1, 6, 1, 6, 4, 6, 106, 8, 6, 11, 6, 12, 6, 107,
		1, 6, 1, 6, 1, 6, 1, 107, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 124, 0, 24,
		1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8,
		83, 1, 0, 0, 0, 10, 101, 1, 0, 0, 0, 12, 103, 1, 0, 0, 0, 14, 22, 3, 2,
		1, 0, 15, 17, 5, 13, 0, 0, 16, 15, 1, 0, 0, 0, 17, 20, 1, 0, 0, 0, 18,
		16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 23, 1, 0, 0, 0, 20, 18, 1, 0, 0,
		0, 21, 23, 5, 0, 0, 1, 22, 18, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 25,
		1, 0, 0, 0, 24, 14, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0,
		26, 27, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28, 29, 3, 12, 6, 0, 29, 30, 5,
		1, 0, 0, 30, 32, 3, 10, 5, 0, 31, 33, 5, 2, 0, 0, 32, 31, 1, 0, 0, 0, 32,
		33, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 36, 5, 11, 0, 0, 35, 28, 1, 0,
		0, 0, 35, 34, 1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 41, 5, 3, 0, 0, 38, 40,
		5, 13, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 54, 3,
		6, 3, 0, 45, 47, 5, 13, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48,
		46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1, 0, 0,
		0, 51, 53, 3, 6, 3, 0, 52, 48, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		57, 58, 5, 4, 0, 0, 58, 68, 1, 0, 0, 0, 59, 63, 5, 3, 0, 0, 60, 62, 5,
		13, 0, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63,
		64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 68, 5, 4, 0,
		0, 67, 37, 1, 0, 0, 0, 67, 59, 1, 0, 0, 0, 68, 5, 1, 0, 0, 0, 69, 70, 5,
		12, 0, 0, 70, 71, 5, 1, 0, 0, 71, 73, 3, 10, 5, 0, 72, 74, 5, 2, 0, 0,
		73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 81, 1, 0, 0, 0, 75, 77, 5,
		13, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78,
		79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 78, 1, 0, 0,
		0, 81, 82, 1, 0, 0, 0, 82, 7, 1, 0, 0, 0, 83, 92, 5, 5, 0, 0, 84, 89, 3,
		10, 5, 0, 85, 86, 5, 6, 0, 0, 86, 88, 3, 10, 5, 0, 87, 85, 1, 0, 0, 0,
		88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		94, 1, 0, 0, 0, 94, 95, 5, 7, 0, 0, 95, 9, 1, 0, 0, 0, 96, 102, 5, 8, 0,
		0, 97, 102, 5, 9, 0, 0, 98, 102, 5, 10, 0, 0, 99, 102, 3, 4, 2, 0, 100,
		102, 3, 8, 4, 0, 101, 96, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 98, 1,
		0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102, 11, 1, 0, 0, 0,
		103, 105, 5, 5, 0, 0, 104, 106, 9, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 110, 5, 7, 0, 0, 110, 13, 1, 0, 0, 0, 17, 18, 22, 26,
		32, 35, 41, 48, 54, 63, 67, 73, 78, 81, 89, 92, 101, 107,
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
	LambdaParserSTRING  = 8
	LambdaParserNUMBER  = 9
	LambdaParserBOOLEAN = 10
	LambdaParserCOMMENT = 11
	LambdaParserKEY     = 12
	LambdaParserNEWLINE = 13
	LambdaParserWS      = 14
)

// LambdaParser rules.
const (
	LambdaParserRULE_prog  = 0
	LambdaParserRULE_stat  = 1
	LambdaParserRULE_obj   = 2
	LambdaParserRULE_pair  = 3
	LambdaParserRULE_array = 4
	LambdaParserRULE_value = 5
	LambdaParserRULE_id    = 6
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

	for ok := true; ok; ok = _la == LambdaParserT__4 || _la == LambdaParserCOMMENT {
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

func (s *AssignmentStatementContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserT__4:
		localctx = NewAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Id()
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

	case LambdaParserCOMMENT:
		localctx = NewCommentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(LambdaParserCOMMENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 4, LambdaParserRULE_obj)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(LambdaParserT__2)
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserNEWLINE {
			{
				p.SetState(38)
				p.Match(LambdaParserNEWLINE)
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(44)
			p.Pair()
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserKEY || _la == LambdaParserNEWLINE {
			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == LambdaParserNEWLINE {
				{
					p.SetState(45)
					p.Match(LambdaParserNEWLINE)
				}

				p.SetState(50)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(51)
				p.Pair()
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(57)
			p.Match(LambdaParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(LambdaParserT__2)
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserNEWLINE {
			{
				p.SetState(60)
				p.Match(LambdaParserNEWLINE)
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(66)
			p.Match(LambdaParserT__3)
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

func (s *PairContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(LambdaParserNEWLINE)
}

func (s *PairContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(LambdaParserNEWLINE, i)
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
	p.EnterRule(localctx, 6, LambdaParserRULE_pair)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(LambdaParserKEY)
	}
	{
		p.SetState(70)
		p.Match(LambdaParserT__0)
	}
	{
		p.SetState(71)
		p.Value()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LambdaParserT__1 {
		{
			p.SetState(72)
			p.Match(LambdaParserT__1)
		}

	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(75)
					p.Match(LambdaParserNEWLINE)
				}

			}
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 8, LambdaParserRULE_array)
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
		p.SetState(83)
		p.Match(LambdaParserT__4)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LambdaParserT__2)|(1<<LambdaParserT__4)|(1<<LambdaParserSTRING)|(1<<LambdaParserNUMBER)|(1<<LambdaParserBOOLEAN))) != 0 {
		{
			p.SetState(84)
			p.Value()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LambdaParserT__5 {
			{
				p.SetState(85)
				p.Match(LambdaParserT__5)
			}
			{
				p.SetState(86)
				p.Value()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(94)
		p.Match(LambdaParserT__6)
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
	p.EnterRule(localctx, 10, LambdaParserRULE_value)

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LambdaParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(LambdaParserSTRING)
		}

	case LambdaParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(LambdaParserNUMBER)
		}

	case LambdaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(LambdaParserBOOLEAN)
		}

	case LambdaParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.Obj()
		}

	case LambdaParserT__4:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LambdaParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LambdaParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }
func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LambdaListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LambdaVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LambdaParser) Id() (localctx IIdContext) {
	this := p
	_ = this

	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LambdaParserRULE_id)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(LambdaParserT__4)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(104)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	{
		p.SetState(109)
		p.Match(LambdaParserT__6)
	}

	return localctx
}
