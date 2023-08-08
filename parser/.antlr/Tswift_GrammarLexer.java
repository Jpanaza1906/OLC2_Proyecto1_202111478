// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Tswift_GrammarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, ASIG=3, DIV=4, INT=5, CONSOLE=6, ENBLANCO=7, ENTERO=8, 
		DECIMAL=9, ID=10, RETURN=11, UL_C=12;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "ENBLANCO", "DIG", 
			"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", "'='", "'/'", "'int'", "'Console.out'", null, null, 
			null, null, "'return'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "ASIG", "DIV", "INT", "CONSOLE", "ENBLANCO", 
			"ENTERO", "DECIMAL", "ID", "RETURN", "UL_C"
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


	public Tswift_GrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Tswift_Grammar.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\16o\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6"+
		"\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\6\b\67\n"+
		"\b\r\b\16\b8\3\b\3\b\3\t\3\t\3\n\6\n@\n\n\r\n\16\nA\3\13\6\13E\n\13\r"+
		"\13\16\13F\3\13\3\13\6\13K\n\13\r\13\16\13L\3\f\3\f\7\fQ\n\f\f\f\16\f"+
		"T\13\f\3\f\3\f\6\fX\n\f\r\f\16\fY\5\f\\\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\16\3\16\3\16\3\16\7\16i\n\16\f\16\16\16l\13\16\3\16\3\16\2\2\17\3"+
		"\3\5\4\7\5\t\6\13\7\r\b\17\t\21\2\23\n\25\13\27\f\31\r\33\16\3\2\7\5\2"+
		"\13\f\17\17\"\"\3\2\62;\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17\2u\2\3"+
		"\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2"+
		"\2\17\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\3\35\3\2\2\2\5\37\3\2\2\2\7!\3\2\2\2\t#\3\2\2\2\13%\3\2\2\2\r"+
		")\3\2\2\2\17\66\3\2\2\2\21<\3\2\2\2\23?\3\2\2\2\25D\3\2\2\2\27[\3\2\2"+
		"\2\31]\3\2\2\2\33d\3\2\2\2\35\36\7+\2\2\36\4\3\2\2\2\37 \7*\2\2 \6\3\2"+
		"\2\2!\"\7?\2\2\"\b\3\2\2\2#$\7\61\2\2$\n\3\2\2\2%&\7k\2\2&\'\7p\2\2\'"+
		"(\7v\2\2(\f\3\2\2\2)*\7E\2\2*+\7q\2\2+,\7p\2\2,-\7u\2\2-.\7q\2\2./\7n"+
		"\2\2/\60\7g\2\2\60\61\7\60\2\2\61\62\7q\2\2\62\63\7w\2\2\63\64\7v\2\2"+
		"\64\16\3\2\2\2\65\67\t\2\2\2\66\65\3\2\2\2\678\3\2\2\28\66\3\2\2\289\3"+
		"\2\2\29:\3\2\2\2:;\b\b\2\2;\20\3\2\2\2<=\t\3\2\2=\22\3\2\2\2>@\5\21\t"+
		"\2?>\3\2\2\2@A\3\2\2\2A?\3\2\2\2AB\3\2\2\2B\24\3\2\2\2CE\5\21\t\2DC\3"+
		"\2\2\2EF\3\2\2\2FD\3\2\2\2FG\3\2\2\2GH\3\2\2\2HJ\7\60\2\2IK\5\21\t\2J"+
		"I\3\2\2\2KL\3\2\2\2LJ\3\2\2\2LM\3\2\2\2M\26\3\2\2\2NR\t\4\2\2OQ\t\5\2"+
		"\2PO\3\2\2\2QT\3\2\2\2RP\3\2\2\2RS\3\2\2\2S\\\3\2\2\2TR\3\2\2\2UW\7a\2"+
		"\2VX\t\5\2\2WV\3\2\2\2XY\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z\\\3\2\2\2[N\3\2"+
		"\2\2[U\3\2\2\2\\\30\3\2\2\2]^\7t\2\2^_\7g\2\2_`\7v\2\2`a\7w\2\2ab\7t\2"+
		"\2bc\7p\2\2c\32\3\2\2\2de\7\61\2\2ef\7\61\2\2fj\3\2\2\2gi\n\6\2\2hg\3"+
		"\2\2\2il\3\2\2\2jh\3\2\2\2jk\3\2\2\2km\3\2\2\2lj\3\2\2\2mn\b\16\3\2n\34"+
		"\3\2\2\2\13\28AFLRY[j\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}