// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/parser/Tswift_Lexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Tswift_Lexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, LLAVEIZQ=3, LLAVEDER=4, DOSPT=5, INTERROGACION=6, 
		MASIGUAL=7, MENOSIGUAL=8, IGUAL=9, DIV=10, MOD=11, MENOS=12, MAS=13, POR=14, 
		IGUALIGUAL=15, DIFERENTE=16, MAYORIGUAL=17, MENORIGUAL=18, MAYOR=19, MENOR=20, 
		AND=21, OR=22, NOT=23, PRINT=24, VAR=25, LET=26, TRUE=27, FALSE=28, RETURN=29, 
		IF=30, ELSE=31, SWITCH=32, CASE=33, BREAK=34, DEFAULT=35, WHILE=36, FOR=37, 
		IN=38, RANGO=39, GUARD=40, CONTINUE=41, NIL=42, INT=43, FLOAT=44, BOOL=45, 
		CHARACTER=46, STRING=47, ENBLANCO=48, ENTERO=49, DECIMAL=50, BOOLEANO=51, 
		CARACTER=52, CADENA=53, ID=54, UL_C=55, ML_C=56;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "INTERROGACION", 
			"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", 
			"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", 
			"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "RETURN", 
			"IF", "ELSE", "SWITCH", "CASE", "BREAK", "DEFAULT", "WHILE", "FOR", "IN", 
			"RANGO", "GUARD", "CONTINUE", "NIL", "INT", "FLOAT", "BOOL", "CHARACTER", 
			"STRING", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", "BOOLEANO", "CARACTER", 
			"CADENA", "ID", "UL_C", "ML_C"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", "'{'", "'}'", "':'", "'?'", "'+='", "'-='", "'='", 
			"'/'", "'%'", "'-'", "'+'", "'*'", "'=='", "'!='", "'>='", "'<='", "'>'", 
			"'<'", "'&&'", "'||'", "'!'", "'print'", "'var'", "'let'", "'true'", 
			"'false'", "'return'", "'if'", "'else'", "'switch'", "'case'", "'break'", 
			"'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", "'continue'", 
			"'nil'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "INTERROGACION", 
			"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", 
			"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", 
			"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "RETURN", 
			"IF", "ELSE", "SWITCH", "CASE", "BREAK", "DEFAULT", "WHILE", "FOR", "IN", 
			"RANGO", "GUARD", "CONTINUE", "NIL", "INT", "FLOAT", "BOOL", "CHARACTER", 
			"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "BOOLEANO", "CARACTER", "CADENA", 
			"ID", "UL_C", "ML_C"
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


	public Tswift_Lexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Tswift_Lexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2:\u0189\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\3\2\3\2\3\3\3\3\3"+
		"\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\13\3"+
		"\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27"+
		"\3\27\3\27\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32"+
		"\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3"+
		" \3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$"+
		"\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3("+
		"\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3,"+
		"\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3/\3/\3/\3\60\3"+
		"\60\3\60\3\60\3\60\3\60\3\60\3\61\6\61\u0133\n\61\r\61\16\61\u0134\3\61"+
		"\3\61\3\62\3\62\3\63\6\63\u013c\n\63\r\63\16\63\u013d\3\64\6\64\u0141"+
		"\n\64\r\64\16\64\u0142\3\64\3\64\6\64\u0147\n\64\r\64\16\64\u0148\3\65"+
		"\3\65\5\65\u014d\n\65\3\66\3\66\3\66\3\66\5\66\u0153\n\66\3\66\3\66\3"+
		"\67\3\67\3\67\3\67\7\67\u015b\n\67\f\67\16\67\u015e\13\67\3\67\3\67\3"+
		"8\38\78\u0164\n8\f8\168\u0167\138\38\38\68\u016b\n8\r8\168\u016c\58\u016f"+
		"\n8\39\39\39\39\79\u0175\n9\f9\169\u0178\139\39\39\3:\3:\3:\3:\7:\u0180"+
		"\n:\f:\16:\u0183\13:\3:\3:\3:\3:\3:\3\u0181\2;\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\2e\63g\64i\65k\66m\67o8q9s:\3\2\t\5\2\13\f"+
		"\17\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17"+
		"\17\2\u0194\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2"+
		"\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2"+
		"\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\3u\3\2\2\2\5w\3\2\2\2\7y\3\2\2\2\t{"+
		"\3\2\2\2\13}\3\2\2\2\r\177\3\2\2\2\17\u0081\3\2\2\2\21\u0084\3\2\2\2\23"+
		"\u0087\3\2\2\2\25\u0089\3\2\2\2\27\u008b\3\2\2\2\31\u008d\3\2\2\2\33\u008f"+
		"\3\2\2\2\35\u0091\3\2\2\2\37\u0093\3\2\2\2!\u0096\3\2\2\2#\u0099\3\2\2"+
		"\2%\u009c\3\2\2\2\'\u009f\3\2\2\2)\u00a1\3\2\2\2+\u00a3\3\2\2\2-\u00a6"+
		"\3\2\2\2/\u00a9\3\2\2\2\61\u00ab\3\2\2\2\63\u00b1\3\2\2\2\65\u00b5\3\2"+
		"\2\2\67\u00b9\3\2\2\29\u00be\3\2\2\2;\u00c4\3\2\2\2=\u00cb\3\2\2\2?\u00ce"+
		"\3\2\2\2A\u00d3\3\2\2\2C\u00da\3\2\2\2E\u00df\3\2\2\2G\u00e5\3\2\2\2I"+
		"\u00ed\3\2\2\2K\u00f3\3\2\2\2M\u00f7\3\2\2\2O\u00fa\3\2\2\2Q\u00fe\3\2"+
		"\2\2S\u0104\3\2\2\2U\u010d\3\2\2\2W\u0111\3\2\2\2Y\u0115\3\2\2\2[\u011b"+
		"\3\2\2\2]\u0120\3\2\2\2_\u012a\3\2\2\2a\u0132\3\2\2\2c\u0138\3\2\2\2e"+
		"\u013b\3\2\2\2g\u0140\3\2\2\2i\u014c\3\2\2\2k\u014e\3\2\2\2m\u0156\3\2"+
		"\2\2o\u016e\3\2\2\2q\u0170\3\2\2\2s\u017b\3\2\2\2uv\7+\2\2v\4\3\2\2\2"+
		"wx\7*\2\2x\6\3\2\2\2yz\7}\2\2z\b\3\2\2\2{|\7\177\2\2|\n\3\2\2\2}~\7<\2"+
		"\2~\f\3\2\2\2\177\u0080\7A\2\2\u0080\16\3\2\2\2\u0081\u0082\7-\2\2\u0082"+
		"\u0083\7?\2\2\u0083\20\3\2\2\2\u0084\u0085\7/\2\2\u0085\u0086\7?\2\2\u0086"+
		"\22\3\2\2\2\u0087\u0088\7?\2\2\u0088\24\3\2\2\2\u0089\u008a\7\61\2\2\u008a"+
		"\26\3\2\2\2\u008b\u008c\7\'\2\2\u008c\30\3\2\2\2\u008d\u008e\7/\2\2\u008e"+
		"\32\3\2\2\2\u008f\u0090\7-\2\2\u0090\34\3\2\2\2\u0091\u0092\7,\2\2\u0092"+
		"\36\3\2\2\2\u0093\u0094\7?\2\2\u0094\u0095\7?\2\2\u0095 \3\2\2\2\u0096"+
		"\u0097\7#\2\2\u0097\u0098\7?\2\2\u0098\"\3\2\2\2\u0099\u009a\7@\2\2\u009a"+
		"\u009b\7?\2\2\u009b$\3\2\2\2\u009c\u009d\7>\2\2\u009d\u009e\7?\2\2\u009e"+
		"&\3\2\2\2\u009f\u00a0\7@\2\2\u00a0(\3\2\2\2\u00a1\u00a2\7>\2\2\u00a2*"+
		"\3\2\2\2\u00a3\u00a4\7(\2\2\u00a4\u00a5\7(\2\2\u00a5,\3\2\2\2\u00a6\u00a7"+
		"\7~\2\2\u00a7\u00a8\7~\2\2\u00a8.\3\2\2\2\u00a9\u00aa\7#\2\2\u00aa\60"+
		"\3\2\2\2\u00ab\u00ac\7r\2\2\u00ac\u00ad\7t\2\2\u00ad\u00ae\7k\2\2\u00ae"+
		"\u00af\7p\2\2\u00af\u00b0\7v\2\2\u00b0\62\3\2\2\2\u00b1\u00b2\7x\2\2\u00b2"+
		"\u00b3\7c\2\2\u00b3\u00b4\7t\2\2\u00b4\64\3\2\2\2\u00b5\u00b6\7n\2\2\u00b6"+
		"\u00b7\7g\2\2\u00b7\u00b8\7v\2\2\u00b8\66\3\2\2\2\u00b9\u00ba\7v\2\2\u00ba"+
		"\u00bb\7t\2\2\u00bb\u00bc\7w\2\2\u00bc\u00bd\7g\2\2\u00bd8\3\2\2\2\u00be"+
		"\u00bf\7h\2\2\u00bf\u00c0\7c\2\2\u00c0\u00c1\7n\2\2\u00c1\u00c2\7u\2\2"+
		"\u00c2\u00c3\7g\2\2\u00c3:\3\2\2\2\u00c4\u00c5\7t\2\2\u00c5\u00c6\7g\2"+
		"\2\u00c6\u00c7\7v\2\2\u00c7\u00c8\7w\2\2\u00c8\u00c9\7t\2\2\u00c9\u00ca"+
		"\7p\2\2\u00ca<\3\2\2\2\u00cb\u00cc\7k\2\2\u00cc\u00cd\7h\2\2\u00cd>\3"+
		"\2\2\2\u00ce\u00cf\7g\2\2\u00cf\u00d0\7n\2\2\u00d0\u00d1\7u\2\2\u00d1"+
		"\u00d2\7g\2\2\u00d2@\3\2\2\2\u00d3\u00d4\7u\2\2\u00d4\u00d5\7y\2\2\u00d5"+
		"\u00d6\7k\2\2\u00d6\u00d7\7v\2\2\u00d7\u00d8\7e\2\2\u00d8\u00d9\7j\2\2"+
		"\u00d9B\3\2\2\2\u00da\u00db\7e\2\2\u00db\u00dc\7c\2\2\u00dc\u00dd\7u\2"+
		"\2\u00dd\u00de\7g\2\2\u00deD\3\2\2\2\u00df\u00e0\7d\2\2\u00e0\u00e1\7"+
		"t\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4\7m\2\2\u00e4F"+
		"\3\2\2\2\u00e5\u00e6\7f\2\2\u00e6\u00e7\7g\2\2\u00e7\u00e8\7h\2\2\u00e8"+
		"\u00e9\7c\2\2\u00e9\u00ea\7w\2\2\u00ea\u00eb\7n\2\2\u00eb\u00ec\7v\2\2"+
		"\u00ecH\3\2\2\2\u00ed\u00ee\7y\2\2\u00ee\u00ef\7j\2\2\u00ef\u00f0\7k\2"+
		"\2\u00f0\u00f1\7n\2\2\u00f1\u00f2\7g\2\2\u00f2J\3\2\2\2\u00f3\u00f4\7"+
		"h\2\2\u00f4\u00f5\7q\2\2\u00f5\u00f6\7t\2\2\u00f6L\3\2\2\2\u00f7\u00f8"+
		"\7k\2\2\u00f8\u00f9\7p\2\2\u00f9N\3\2\2\2\u00fa\u00fb\7\60\2\2\u00fb\u00fc"+
		"\7\60\2\2\u00fc\u00fd\7\60\2\2\u00fdP\3\2\2\2\u00fe\u00ff\7i\2\2\u00ff"+
		"\u0100\7w\2\2\u0100\u0101\7c\2\2\u0101\u0102\7t\2\2\u0102\u0103\7f\2\2"+
		"\u0103R\3\2\2\2\u0104\u0105\7e\2\2\u0105\u0106\7q\2\2\u0106\u0107\7p\2"+
		"\2\u0107\u0108\7v\2\2\u0108\u0109\7k\2\2\u0109\u010a\7p\2\2\u010a\u010b"+
		"\7w\2\2\u010b\u010c\7g\2\2\u010cT\3\2\2\2\u010d\u010e\7p\2\2\u010e\u010f"+
		"\7k\2\2\u010f\u0110\7n\2\2\u0110V\3\2\2\2\u0111\u0112\7K\2\2\u0112\u0113"+
		"\7p\2\2\u0113\u0114\7v\2\2\u0114X\3\2\2\2\u0115\u0116\7H\2\2\u0116\u0117"+
		"\7n\2\2\u0117\u0118\7q\2\2\u0118\u0119\7c\2\2\u0119\u011a\7v\2\2\u011a"+
		"Z\3\2\2\2\u011b\u011c\7D\2\2\u011c\u011d\7q\2\2\u011d\u011e\7q\2\2\u011e"+
		"\u011f\7n\2\2\u011f\\\3\2\2\2\u0120\u0121\7E\2\2\u0121\u0122\7j\2\2\u0122"+
		"\u0123\7c\2\2\u0123\u0124\7t\2\2\u0124\u0125\7c\2\2\u0125\u0126\7e\2\2"+
		"\u0126\u0127\7v\2\2\u0127\u0128\7g\2\2\u0128\u0129\7t\2\2\u0129^\3\2\2"+
		"\2\u012a\u012b\7U\2\2\u012b\u012c\7v\2\2\u012c\u012d\7t\2\2\u012d\u012e"+
		"\7k\2\2\u012e\u012f\7p\2\2\u012f\u0130\7i\2\2\u0130`\3\2\2\2\u0131\u0133"+
		"\t\2\2\2\u0132\u0131\3\2\2\2\u0133\u0134\3\2\2\2\u0134\u0132\3\2\2\2\u0134"+
		"\u0135\3\2\2\2\u0135\u0136\3\2\2\2\u0136\u0137\b\61\2\2\u0137b\3\2\2\2"+
		"\u0138\u0139\t\3\2\2\u0139d\3\2\2\2\u013a\u013c\5c\62\2\u013b\u013a\3"+
		"\2\2\2\u013c\u013d\3\2\2\2\u013d\u013b\3\2\2\2\u013d\u013e\3\2\2\2\u013e"+
		"f\3\2\2\2\u013f\u0141\5c\62\2\u0140\u013f\3\2\2\2\u0141\u0142\3\2\2\2"+
		"\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143\u0144\3\2\2\2\u0144\u0146"+
		"\7\60\2\2\u0145\u0147\5c\62\2\u0146\u0145\3\2\2\2\u0147\u0148\3\2\2\2"+
		"\u0148\u0146\3\2\2\2\u0148\u0149\3\2\2\2\u0149h\3\2\2\2\u014a\u014d\5"+
		"\67\34\2\u014b\u014d\59\35\2\u014c\u014a\3\2\2\2\u014c\u014b\3\2\2\2\u014d"+
		"j\3\2\2\2\u014e\u0152\7)\2\2\u014f\u0153\n\4\2\2\u0150\u0151\7^\2\2\u0151"+
		"\u0153\13\2\2\2\u0152\u014f\3\2\2\2\u0152\u0150\3\2\2\2\u0153\u0154\3"+
		"\2\2\2\u0154\u0155\7)\2\2\u0155l\3\2\2\2\u0156\u015c\7$\2\2\u0157\u015b"+
		"\n\5\2\2\u0158\u0159\7^\2\2\u0159\u015b\13\2\2\2\u015a\u0157\3\2\2\2\u015a"+
		"\u0158\3\2\2\2\u015b\u015e\3\2\2\2\u015c\u015a\3\2\2\2\u015c\u015d\3\2"+
		"\2\2\u015d\u015f\3\2\2\2\u015e\u015c\3\2\2\2\u015f\u0160\7$\2\2\u0160"+
		"n\3\2\2\2\u0161\u0165\t\6\2\2\u0162\u0164\t\7\2\2\u0163\u0162\3\2\2\2"+
		"\u0164\u0167\3\2\2\2\u0165\u0163\3\2\2\2\u0165\u0166\3\2\2\2\u0166\u016f"+
		"\3\2\2\2\u0167\u0165\3\2\2\2\u0168\u016a\7a\2\2\u0169\u016b\t\7\2\2\u016a"+
		"\u0169\3\2\2\2\u016b\u016c\3\2\2\2\u016c\u016a\3\2\2\2\u016c\u016d\3\2"+
		"\2\2\u016d\u016f\3\2\2\2\u016e\u0161\3\2\2\2\u016e\u0168\3\2\2\2\u016f"+
		"p\3\2\2\2\u0170\u0171\7\61\2\2\u0171\u0172\7\61\2\2\u0172\u0176\3\2\2"+
		"\2\u0173\u0175\n\b\2\2\u0174\u0173\3\2\2\2\u0175\u0178\3\2\2\2\u0176\u0174"+
		"\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u0179\3\2\2\2\u0178\u0176\3\2\2\2\u0179"+
		"\u017a\b9\3\2\u017ar\3\2\2\2\u017b\u017c\7\61\2\2\u017c\u017d\7,\2\2\u017d"+
		"\u0181\3\2\2\2\u017e\u0180\13\2\2\2\u017f\u017e\3\2\2\2\u0180\u0183\3"+
		"\2\2\2\u0181\u0182\3\2\2\2\u0181\u017f\3\2\2\2\u0182\u0184\3\2\2\2\u0183"+
		"\u0181\3\2\2\2\u0184\u0185\7,\2\2\u0185\u0186\7\61\2\2\u0186\u0187\3\2"+
		"\2\2\u0187\u0188\b:\3\2\u0188t\3\2\2\2\20\2\u0134\u013d\u0142\u0148\u014c"+
		"\u0152\u015a\u015c\u0165\u016c\u016e\u0176\u0181\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}