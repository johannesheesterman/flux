// Generated from /home/johannes/fun/lambda/src/parser/Lambda.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class LambdaLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, STRING=8, NUMBER=9, 
		BOOLEAN=10, ID=11, COMMENT=12, MODIFIER=13, KEY=14, NEWLINE=15, WS=16;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "STRING", "NUMBER", 
			"BOOLEAN", "INT", "ID", "COMMENT", "MODIFIER", "KEY", "NEWLINE", "WS"
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


	public LambdaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lambda.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\22\u008c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t"+
		"\7\t\66\n\t\f\t\16\t9\13\t\3\t\3\t\3\n\5\n>\n\n\3\n\3\n\3\n\6\nC\n\n\r"+
		"\n\16\nD\5\nG\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13R\n"+
		"\13\3\f\3\f\3\f\7\fW\n\f\f\f\16\fZ\13\f\5\f\\\n\f\3\r\3\r\6\r`\n\r\r\r"+
		"\16\ra\3\r\3\r\3\16\3\16\7\16h\n\16\f\16\16\16k\13\16\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17z\n\17\3\20\6\20"+
		"}\n\20\r\20\16\20~\3\21\5\21\u0082\n\21\3\21\3\21\3\22\6\22\u0087\n\22"+
		"\r\22\16\22\u0088\3\22\3\22\4\67a\2\23\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\2\31\r\33\16\35\17\37\20!\21#\22\3\2\7\3\2\62;\3\2\63"+
		";\4\2\f\f\17\17\5\2\62;C\\c|\4\2\13\13\"\"\2\u0097\2\3\3\2\2\2\2\5\3\2"+
		"\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2"+
		"\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\3%\3\2\2\2\5\'\3\2\2\2\7)\3\2"+
		"\2\2\t+\3\2\2\2\13-\3\2\2\2\r/\3\2\2\2\17\61\3\2\2\2\21\63\3\2\2\2\23"+
		"=\3\2\2\2\25Q\3\2\2\2\27[\3\2\2\2\31]\3\2\2\2\33e\3\2\2\2\35y\3\2\2\2"+
		"\37|\3\2\2\2!\u0081\3\2\2\2#\u0086\3\2\2\2%&\7?\2\2&\4\3\2\2\2\'(\7}\2"+
		"\2(\6\3\2\2\2)*\7\177\2\2*\b\3\2\2\2+,\7=\2\2,\n\3\2\2\2-.\7]\2\2.\f\3"+
		"\2\2\2/\60\7.\2\2\60\16\3\2\2\2\61\62\7_\2\2\62\20\3\2\2\2\63\67\7$\2"+
		"\2\64\66\13\2\2\2\65\64\3\2\2\2\669\3\2\2\2\678\3\2\2\2\67\65\3\2\2\2"+
		"8:\3\2\2\29\67\3\2\2\2:;\7$\2\2;\22\3\2\2\2<>\7/\2\2=<\3\2\2\2=>\3\2\2"+
		"\2>?\3\2\2\2?F\5\27\f\2@B\7\60\2\2AC\t\2\2\2BA\3\2\2\2CD\3\2\2\2DB\3\2"+
		"\2\2DE\3\2\2\2EG\3\2\2\2F@\3\2\2\2FG\3\2\2\2G\24\3\2\2\2HI\7v\2\2IJ\7"+
		"t\2\2JK\7w\2\2KR\7g\2\2LM\7h\2\2MN\7c\2\2NO\7n\2\2OP\7u\2\2PR\7g\2\2Q"+
		"H\3\2\2\2QL\3\2\2\2R\26\3\2\2\2S\\\7\62\2\2TX\t\3\2\2UW\t\2\2\2VU\3\2"+
		"\2\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\\\3\2\2\2ZX\3\2\2\2[S\3\2\2\2[T\3"+
		"\2\2\2\\\30\3\2\2\2]_\7]\2\2^`\13\2\2\2_^\3\2\2\2`a\3\2\2\2ab\3\2\2\2"+
		"a_\3\2\2\2bc\3\2\2\2cd\7_\2\2d\32\3\2\2\2ei\7%\2\2fh\n\4\2\2gf\3\2\2\2"+
		"hk\3\2\2\2ig\3\2\2\2ij\3\2\2\2j\34\3\2\2\2ki\3\2\2\2lm\7r\2\2mn\7w\2\2"+
		"no\7d\2\2op\7n\2\2pq\7k\2\2qz\7e\2\2rs\7r\2\2st\7t\2\2tu\7k\2\2uv\7x\2"+
		"\2vw\7c\2\2wx\7v\2\2xz\7g\2\2yl\3\2\2\2yr\3\2\2\2z\36\3\2\2\2{}\t\5\2"+
		"\2|{\3\2\2\2}~\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177 \3\2\2\2\u0080\u0082"+
		"\7\17\2\2\u0081\u0080\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0083\3\2\2\2"+
		"\u0083\u0084\7\f\2\2\u0084\"\3\2\2\2\u0085\u0087\t\6\2\2\u0086\u0085\3"+
		"\2\2\2\u0087\u0088\3\2\2\2\u0088\u0086\3\2\2\2\u0088\u0089\3\2\2\2\u0089"+
		"\u008a\3\2\2\2\u008a\u008b\b\22\2\2\u008b$\3\2\2\2\20\2\67=DFQX[aiy~\u0081"+
		"\u0088\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}