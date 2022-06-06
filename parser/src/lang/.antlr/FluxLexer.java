// Generated from /home/johannes/fun/lambda/parser/src/lang/Flux.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class FluxLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		STRING=10, NUMBER=11, BOOLEAN=12, ID=13, COMMENT=14, KEY=15, NEWLINE=16, 
		WS=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"STRING", "NUMBER", "BOOLEAN", "INT", "ID", "COMMENT", "KEY", "NEWLINE", 
			"WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "';'", "'('", "','", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "STRING", 
			"NUMBER", "BOOLEAN", "ID", "COMMENT", "KEY", "NEWLINE", "WS"
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


	public FluxLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Flux.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\23\u0083\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3"+
		"\b\3\t\3\t\3\n\3\n\3\13\3\13\7\13<\n\13\f\13\16\13?\13\13\3\13\3\13\3"+
		"\f\5\fD\n\f\3\f\3\f\3\f\6\fI\n\f\r\f\16\fJ\5\fM\n\f\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\5\rX\n\r\3\16\3\16\3\16\7\16]\n\16\f\16\16\16`\13\16"+
		"\5\16b\n\16\3\17\3\17\6\17f\n\17\r\17\16\17g\3\17\3\17\3\20\3\20\7\20"+
		"n\n\20\f\20\16\20q\13\20\3\21\6\21t\n\21\r\21\16\21u\3\22\5\22y\n\22\3"+
		"\22\3\22\3\23\6\23~\n\23\r\23\16\23\177\3\23\3\23\4=g\2\24\3\3\5\4\7\5"+
		"\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\2\35\17\37\20!\21#\22"+
		"%\23\3\2\7\3\2\62;\3\2\63;\4\2\f\f\17\17\5\2\62;C\\c|\4\2\13\13\"\"\2"+
		"\u008d\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2"+
		"\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3"+
		"\2\2\2\2\31\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\3\'\3\2\2\2\5)\3\2\2\2\7+\3\2\2\2\t-\3\2\2\2\13/\3\2\2\2\r\61"+
		"\3\2\2\2\17\63\3\2\2\2\21\65\3\2\2\2\23\67\3\2\2\2\259\3\2\2\2\27C\3\2"+
		"\2\2\31W\3\2\2\2\33a\3\2\2\2\35c\3\2\2\2\37k\3\2\2\2!s\3\2\2\2#x\3\2\2"+
		"\2%}\3\2\2\2\'(\7?\2\2(\4\3\2\2\2)*\7=\2\2*\6\3\2\2\2+,\7*\2\2,\b\3\2"+
		"\2\2-.\7.\2\2.\n\3\2\2\2/\60\7+\2\2\60\f\3\2\2\2\61\62\7}\2\2\62\16\3"+
		"\2\2\2\63\64\7\177\2\2\64\20\3\2\2\2\65\66\7]\2\2\66\22\3\2\2\2\678\7"+
		"_\2\28\24\3\2\2\29=\7$\2\2:<\13\2\2\2;:\3\2\2\2<?\3\2\2\2=>\3\2\2\2=;"+
		"\3\2\2\2>@\3\2\2\2?=\3\2\2\2@A\7$\2\2A\26\3\2\2\2BD\7/\2\2CB\3\2\2\2C"+
		"D\3\2\2\2DE\3\2\2\2EL\5\33\16\2FH\7\60\2\2GI\t\2\2\2HG\3\2\2\2IJ\3\2\2"+
		"\2JH\3\2\2\2JK\3\2\2\2KM\3\2\2\2LF\3\2\2\2LM\3\2\2\2M\30\3\2\2\2NO\7v"+
		"\2\2OP\7t\2\2PQ\7w\2\2QX\7g\2\2RS\7h\2\2ST\7c\2\2TU\7n\2\2UV\7u\2\2VX"+
		"\7g\2\2WN\3\2\2\2WR\3\2\2\2X\32\3\2\2\2Yb\7\62\2\2Z^\t\3\2\2[]\t\2\2\2"+
		"\\[\3\2\2\2]`\3\2\2\2^\\\3\2\2\2^_\3\2\2\2_b\3\2\2\2`^\3\2\2\2aY\3\2\2"+
		"\2aZ\3\2\2\2b\34\3\2\2\2ce\7]\2\2df\13\2\2\2ed\3\2\2\2fg\3\2\2\2gh\3\2"+
		"\2\2ge\3\2\2\2hi\3\2\2\2ij\7_\2\2j\36\3\2\2\2ko\7%\2\2ln\n\4\2\2ml\3\2"+
		"\2\2nq\3\2\2\2om\3\2\2\2op\3\2\2\2p \3\2\2\2qo\3\2\2\2rt\t\5\2\2sr\3\2"+
		"\2\2tu\3\2\2\2us\3\2\2\2uv\3\2\2\2v\"\3\2\2\2wy\7\17\2\2xw\3\2\2\2xy\3"+
		"\2\2\2yz\3\2\2\2z{\7\f\2\2{$\3\2\2\2|~\t\6\2\2}|\3\2\2\2~\177\3\2\2\2"+
		"\177}\3\2\2\2\177\u0080\3\2\2\2\u0080\u0081\3\2\2\2\u0081\u0082\b\23\2"+
		"\2\u0082&\3\2\2\2\17\2=CJLW^agoux\177\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}