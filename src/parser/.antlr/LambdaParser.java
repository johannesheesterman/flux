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
		BOOLEAN=10, COMMENT=11, KEY=12, NEWLINE=13, WS=14;
	public static final int
		RULE_prog = 0, RULE_stat = 1, RULE_obj = 2, RULE_pair = 3, RULE_array = 4, 
		RULE_value = 5, RULE_id = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stat", "obj", "pair", "array", "value", "id"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "';'", "'{'", "'}'", "'['", "','", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, "STRING", "NUMBER", "BOOLEAN", 
			"COMMENT", "KEY", "NEWLINE", "WS"
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
			setState(24); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(14);
				stat();
				setState(22);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(18);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NEWLINE) {
						{
						{
						setState(15);
						match(NEWLINE);
						}
						}
						setState(20);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
					break;
				case 2:
					{
					setState(21);
					match(EOF);
					}
					break;
				}
				}
				}
				setState(26); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__4 || _la==COMMENT );
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
		public IdContext id() {
			return getRuleContext(IdContext.class,0);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public AssignmentStatementContext(StatContext ctx) { copyFrom(ctx); }
	}

	public final StatContext stat() throws RecognitionException {
		StatContext _localctx = new StatContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stat);
		int _la;
		try {
			setState(35);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
				_localctx = new AssignmentStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(28);
				id();
				setState(29);
				match(T__0);
				setState(30);
				value();
				setState(32);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__1) {
					{
					setState(31);
					match(T__1);
					}
				}

				}
				break;
			case COMMENT:
				_localctx = new CommentStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				match(COMMENT);
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
			setState(67);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(37);
				match(T__2);
				setState(41);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(38);
					match(NEWLINE);
					}
					}
					setState(43);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(44);
				pair();
				setState(54);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==KEY || _la==NEWLINE) {
					{
					{
					setState(48);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==NEWLINE) {
						{
						{
						setState(45);
						match(NEWLINE);
						}
						}
						setState(50);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					setState(51);
					pair();
					}
					}
					setState(56);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(57);
				match(T__3);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(59);
				match(T__2);
				{
				setState(63);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==NEWLINE) {
					{
					{
					setState(60);
					match(NEWLINE);
					}
					}
					setState(65);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				setState(66);
				match(T__3);
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
			setState(69);
			match(KEY);
			setState(70);
			match(T__0);
			setState(71);
			value();
			setState(73);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__1) {
				{
				setState(72);
				match(T__1);
				}
			}

			setState(81);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(78);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(75);
						match(NEWLINE);
						}
						} 
					}
					setState(80);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			match(T__4);
			setState(92);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__4) | (1L << STRING) | (1L << NUMBER) | (1L << BOOLEAN))) != 0)) {
				{
				setState(84);
				value();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__5) {
					{
					{
					setState(85);
					match(T__5);
					setState(86);
					value();
					}
					}
					setState(91);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(94);
			match(T__6);
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
			setState(101);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(96);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(97);
				match(NUMBER);
				}
				break;
			case BOOLEAN:
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				match(BOOLEAN);
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 4);
				{
				setState(99);
				obj();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 5);
				{
				setState(100);
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

	public static class IdContext extends ParserRuleContext {
		public IdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id; }
	}

	public final IdContext id() throws RecognitionException {
		IdContext _localctx = new IdContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_id);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			match(T__4);
			setState(105); 
			_errHandler.sync(this);
			_alt = 1+1;
			do {
				switch (_alt) {
				case 1+1:
					{
					{
					setState(104);
					matchWildcard();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(107); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			} while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(109);
			match(T__6);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\20r\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\7\2\23\n\2\f\2\16\2"+
		"\26\13\2\3\2\5\2\31\n\2\6\2\33\n\2\r\2\16\2\34\3\3\3\3\3\3\3\3\5\3#\n"+
		"\3\3\3\5\3&\n\3\3\4\3\4\7\4*\n\4\f\4\16\4-\13\4\3\4\3\4\7\4\61\n\4\f\4"+
		"\16\4\64\13\4\3\4\7\4\67\n\4\f\4\16\4:\13\4\3\4\3\4\3\4\3\4\7\4@\n\4\f"+
		"\4\16\4C\13\4\3\4\5\4F\n\4\3\5\3\5\3\5\3\5\5\5L\n\5\3\5\7\5O\n\5\f\5\16"+
		"\5R\13\5\5\5T\n\5\3\6\3\6\3\6\3\6\7\6Z\n\6\f\6\16\6]\13\6\5\6_\n\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\7\5\7h\n\7\3\b\3\b\6\bl\n\b\r\b\16\bm\3\b\3\b\3"+
		"\b\3m\2\t\2\4\6\b\n\f\16\2\2\2~\2\32\3\2\2\2\4%\3\2\2\2\6E\3\2\2\2\bG"+
		"\3\2\2\2\nU\3\2\2\2\fg\3\2\2\2\16i\3\2\2\2\20\30\5\4\3\2\21\23\7\17\2"+
		"\2\22\21\3\2\2\2\23\26\3\2\2\2\24\22\3\2\2\2\24\25\3\2\2\2\25\31\3\2\2"+
		"\2\26\24\3\2\2\2\27\31\7\2\2\3\30\24\3\2\2\2\30\27\3\2\2\2\31\33\3\2\2"+
		"\2\32\20\3\2\2\2\33\34\3\2\2\2\34\32\3\2\2\2\34\35\3\2\2\2\35\3\3\2\2"+
		"\2\36\37\5\16\b\2\37 \7\3\2\2 \"\5\f\7\2!#\7\4\2\2\"!\3\2\2\2\"#\3\2\2"+
		"\2#&\3\2\2\2$&\7\r\2\2%\36\3\2\2\2%$\3\2\2\2&\5\3\2\2\2\'+\7\5\2\2(*\7"+
		"\17\2\2)(\3\2\2\2*-\3\2\2\2+)\3\2\2\2+,\3\2\2\2,.\3\2\2\2-+\3\2\2\2.8"+
		"\5\b\5\2/\61\7\17\2\2\60/\3\2\2\2\61\64\3\2\2\2\62\60\3\2\2\2\62\63\3"+
		"\2\2\2\63\65\3\2\2\2\64\62\3\2\2\2\65\67\5\b\5\2\66\62\3\2\2\2\67:\3\2"+
		"\2\28\66\3\2\2\289\3\2\2\29;\3\2\2\2:8\3\2\2\2;<\7\6\2\2<F\3\2\2\2=A\7"+
		"\5\2\2>@\7\17\2\2?>\3\2\2\2@C\3\2\2\2A?\3\2\2\2AB\3\2\2\2BD\3\2\2\2CA"+
		"\3\2\2\2DF\7\6\2\2E\'\3\2\2\2E=\3\2\2\2F\7\3\2\2\2GH\7\16\2\2HI\7\3\2"+
		"\2IK\5\f\7\2JL\7\4\2\2KJ\3\2\2\2KL\3\2\2\2LS\3\2\2\2MO\7\17\2\2NM\3\2"+
		"\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2QT\3\2\2\2RP\3\2\2\2SP\3\2\2\2ST\3\2"+
		"\2\2T\t\3\2\2\2U^\7\7\2\2V[\5\f\7\2WX\7\b\2\2XZ\5\f\7\2YW\3\2\2\2Z]\3"+
		"\2\2\2[Y\3\2\2\2[\\\3\2\2\2\\_\3\2\2\2][\3\2\2\2^V\3\2\2\2^_\3\2\2\2_"+
		"`\3\2\2\2`a\7\t\2\2a\13\3\2\2\2bh\7\n\2\2ch\7\13\2\2dh\7\f\2\2eh\5\6\4"+
		"\2fh\5\n\6\2gb\3\2\2\2gc\3\2\2\2gd\3\2\2\2ge\3\2\2\2gf\3\2\2\2h\r\3\2"+
		"\2\2ik\7\7\2\2jl\13\2\2\2kj\3\2\2\2lm\3\2\2\2mn\3\2\2\2mk\3\2\2\2no\3"+
		"\2\2\2op\7\t\2\2p\17\3\2\2\2\23\24\30\34\"%+\628AEKPS[^gm";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}