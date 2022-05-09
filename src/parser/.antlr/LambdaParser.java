// Generated from /home/johannes/fun/lambda/src/parser/Lambda.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class LambdaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, STRING=8, NUMBER=9, 
		BOOLEAN=10, ID=11, COMMENT=12, MODIFIER=13, KEY=14, NEWLINE=15, WS=16;
	public static final int
		RULE_prog = 0, RULE_stat = 1, RULE_obj = 2, RULE_pair = 3, RULE_array = 4, 
		RULE_value = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stat", "obj", "pair", "array", "value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'{'", "'}'", "';'", "'['", "','", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, "STRING", "NUMBER", "BOOLEAN", 
			"ID", "COMMENT", "MODIFIER", "KEY", "NEWLINE", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Lambda.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public LambdaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgContext extends ParserRuleContext {
		public List<StatContext> stat() {
			return getRuleContexts(StatContext.class);
		}
		public StatContext stat(int i) {
			return getRuleContext(StatContext.class,i);
		}
		public List<TerminalNode> EOF() { return getTokens(LambdaParser.EOF); }
		public TerminalNode EOF(int i) {
			return getToken(LambdaParser.EOF, i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(LambdaParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(LambdaParser.NEWLINE, i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(12);
				stat();
				setState(20);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(16);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NEWLINE) {
						{
						{
						setState(13);
						match(NEWLINE);
						}
						}
						setState(18);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
					break;
				case 2:
					{
					setState(19);
					match(EOF);
					}
					break;
				}
				}
				}
				setState(24); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID || _la==COMMENT );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatContext extends ParserRuleContext {
		public StatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stat; }
	 
		public StatContext() { }
		public void copyFrom(StatContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class CommentStatementContext extends StatContext {
		public TerminalNode COMMENT() { return getToken(LambdaParser.COMMENT, 0); }
		public CommentStatementContext(StatContext ctx) { copyFrom(ctx); }
	}
	public static class AssignmentStatementContext extends StatContext {
		public TerminalNode ID() { return getToken(LambdaParser.ID, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public AssignmentStatementContext(StatContext ctx) { copyFrom(ctx); }
	}
	public static class VariableDeclarationStatementContext extends StatContext {
		public TerminalNode ID() { return getToken(LambdaParser.ID, 0); }
		public ObjContext obj() {
			return getRuleContext(ObjContext.class,0);
		}
		public VariableDeclarationStatementContext(StatContext ctx) { copyFrom(ctx); }
	}

	public final StatContext stat() throws RecognitionException {
		StatContext _localctx = new StatContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stat);
		int _la;
		try {
			setState(36);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new AssignmentStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				match(ID);
				setState(27);
				match(T__0);
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__4) | (1L << STRING) | (1L << NUMBER) | (1L << BOOLEAN))) != 0)) {
					{
					setState(28);
					value();
					}
				}

				}
				break;
			case 2:
				_localctx = new VariableDeclarationStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				match(ID);
				setState(33);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__1) {
					{
					setState(32);
					obj();
					}
				}

				}
				break;
			case 3:
				_localctx = new CommentStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				match(COMMENT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjContext extends ParserRuleContext {
		public List<PairContext> pair() {
			return getRuleContexts(PairContext.class);
		}
		public PairContext pair(int i) {
			return getRuleContext(PairContext.class,i);
		}
		public List<TerminalNode> NEWLINE() { return getTokens(LambdaParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(LambdaParser.NEWLINE, i);
		}
		public ObjContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_obj; }
	}

	public final ObjContext obj() throws RecognitionException {
		ObjContext _localctx = new ObjContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_obj);
		int _la;
		try {
			setState(68);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				match(T__1);
				setState(42);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(39);
					match(NEWLINE);
					}
					}
					setState(44);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(45);
				pair();
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MODIFIER) | (1L << KEY) | (1L << NEWLINE))) != 0)) {
					{
					{
					setState(49);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NEWLINE) {
						{
						{
						setState(46);
						match(NEWLINE);
						}
						}
						setState(51);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(52);
					pair();
					}
					}
					setState(57);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(58);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				match(T__1);
				{
				setState(64);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(61);
					match(NEWLINE);
					}
					}
					setState(66);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				setState(67);
				match(T__2);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PairContext extends ParserRuleContext {
		public TerminalNode KEY() { return getToken(LambdaParser.KEY, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public TerminalNode MODIFIER() { return getToken(LambdaParser.MODIFIER, 0); }
		public List<TerminalNode> NEWLINE() { return getTokens(LambdaParser.NEWLINE); }
		public TerminalNode NEWLINE(int i) {
			return getToken(LambdaParser.NEWLINE, i);
		}
		public PairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pair; }
	}

	public final PairContext pair() throws RecognitionException {
		PairContext _localctx = new PairContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_pair);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==MODIFIER) {
				{
				setState(70);
				match(MODIFIER);
				}
			}

			setState(73);
			match(KEY);
			setState(74);
			match(T__0);
			setState(75);
			value();
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(76);
				match(T__3);
				}
			}

			setState(85);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(82);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(79);
						match(NEWLINE);
						}
						} 
					}
					setState(84);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
				}
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array; }
	}

	public final ArrayContext array() throws RecognitionException {
		ArrayContext _localctx = new ArrayContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_array);
		int _la;
		try {
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				match(T__4);
				setState(88);
				value();
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__5) {
					{
					{
					setState(89);
					match(T__5);
					setState(90);
					value();
					}
					}
					setState(95);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(96);
				match(T__6);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(98);
				match(T__4);
				setState(99);
				match(T__6);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(LambdaParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(LambdaParser.NUMBER, 0); }
		public TerminalNode BOOLEAN() { return getToken(LambdaParser.BOOLEAN, 0); }
		public ObjContext obj() {
			return getRuleContext(ObjContext.class,0);
		}
		public ArrayContext array() {
			return getRuleContext(ArrayContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_value);
		try {
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(NUMBER);
				}
				break;
			case BOOLEAN:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(BOOLEAN);
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 4);
				{
				setState(105);
				obj();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 5);
				{
				setState(106);
				array();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\22p\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\7\2\21\n\2\f\2\16\2\24\13\2"+
		"\3\2\5\2\27\n\2\6\2\31\n\2\r\2\16\2\32\3\3\3\3\3\3\5\3 \n\3\3\3\3\3\5"+
		"\3$\n\3\3\3\5\3\'\n\3\3\4\3\4\7\4+\n\4\f\4\16\4.\13\4\3\4\3\4\7\4\62\n"+
		"\4\f\4\16\4\65\13\4\3\4\7\48\n\4\f\4\16\4;\13\4\3\4\3\4\3\4\3\4\7\4A\n"+
		"\4\f\4\16\4D\13\4\3\4\5\4G\n\4\3\5\5\5J\n\5\3\5\3\5\3\5\3\5\5\5P\n\5\3"+
		"\5\7\5S\n\5\f\5\16\5V\13\5\5\5X\n\5\3\6\3\6\3\6\3\6\7\6^\n\6\f\6\16\6"+
		"a\13\6\3\6\3\6\3\6\3\6\5\6g\n\6\3\7\3\7\3\7\3\7\3\7\5\7n\n\7\3\7\2\2\b"+
		"\2\4\6\b\n\f\2\2\2\177\2\30\3\2\2\2\4&\3\2\2\2\6F\3\2\2\2\bI\3\2\2\2\n"+
		"f\3\2\2\2\fm\3\2\2\2\16\26\5\4\3\2\17\21\7\21\2\2\20\17\3\2\2\2\21\24"+
		"\3\2\2\2\22\20\3\2\2\2\22\23\3\2\2\2\23\27\3\2\2\2\24\22\3\2\2\2\25\27"+
		"\7\2\2\3\26\22\3\2\2\2\26\25\3\2\2\2\27\31\3\2\2\2\30\16\3\2\2\2\31\32"+
		"\3\2\2\2\32\30\3\2\2\2\32\33\3\2\2\2\33\3\3\2\2\2\34\35\7\r\2\2\35\37"+
		"\7\3\2\2\36 \5\f\7\2\37\36\3\2\2\2\37 \3\2\2\2 \'\3\2\2\2!#\7\r\2\2\""+
		"$\5\6\4\2#\"\3\2\2\2#$\3\2\2\2$\'\3\2\2\2%\'\7\16\2\2&\34\3\2\2\2&!\3"+
		"\2\2\2&%\3\2\2\2\'\5\3\2\2\2(,\7\4\2\2)+\7\21\2\2*)\3\2\2\2+.\3\2\2\2"+
		",*\3\2\2\2,-\3\2\2\2-/\3\2\2\2.,\3\2\2\2/9\5\b\5\2\60\62\7\21\2\2\61\60"+
		"\3\2\2\2\62\65\3\2\2\2\63\61\3\2\2\2\63\64\3\2\2\2\64\66\3\2\2\2\65\63"+
		"\3\2\2\2\668\5\b\5\2\67\63\3\2\2\28;\3\2\2\29\67\3\2\2\29:\3\2\2\2:<\3"+
		"\2\2\2;9\3\2\2\2<=\7\5\2\2=G\3\2\2\2>B\7\4\2\2?A\7\21\2\2@?\3\2\2\2AD"+
		"\3\2\2\2B@\3\2\2\2BC\3\2\2\2CE\3\2\2\2DB\3\2\2\2EG\7\5\2\2F(\3\2\2\2F"+
		">\3\2\2\2G\7\3\2\2\2HJ\7\17\2\2IH\3\2\2\2IJ\3\2\2\2JK\3\2\2\2KL\7\20\2"+
		"\2LM\7\3\2\2MO\5\f\7\2NP\7\6\2\2ON\3\2\2\2OP\3\2\2\2PW\3\2\2\2QS\7\21"+
		"\2\2RQ\3\2\2\2SV\3\2\2\2TR\3\2\2\2TU\3\2\2\2UX\3\2\2\2VT\3\2\2\2WT\3\2"+
		"\2\2WX\3\2\2\2X\t\3\2\2\2YZ\7\7\2\2Z_\5\f\7\2[\\\7\b\2\2\\^\5\f\7\2]["+
		"\3\2\2\2^a\3\2\2\2_]\3\2\2\2_`\3\2\2\2`b\3\2\2\2a_\3\2\2\2bc\7\t\2\2c"+
		"g\3\2\2\2de\7\7\2\2eg\7\t\2\2fY\3\2\2\2fd\3\2\2\2g\13\3\2\2\2hn\7\n\2"+
		"\2in\7\13\2\2jn\7\f\2\2kn\5\6\4\2ln\5\n\6\2mh\3\2\2\2mi\3\2\2\2mj\3\2"+
		"\2\2mk\3\2\2\2ml\3\2\2\2n\r\3\2\2\2\24\22\26\32\37#&,\639BFIOTW_fm";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}