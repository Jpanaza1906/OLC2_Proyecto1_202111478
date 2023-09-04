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
		T__0=1, PARDER=2, PARIZQ=3, LLAVEIZQ=4, LLAVEDER=5, DOSPT=6, PTCOMA=7, 
		INTERROGACION=8, MASIGUAL=9, MENOSIGUAL=10, IGUAL=11, DIV=12, MOD=13, 
		MENOS=14, MAS=15, POR=16, IGUALIGUAL=17, DIFERENTE=18, MAYORIGUAL=19, 
		MENORIGUAL=20, MAYOR=21, MENOR=22, AND=23, OR=24, NOT=25, PRINT=26, VAR=27, 
		LET=28, TRUE=29, FALSE=30, IF=31, ELSE=32, SWITCH=33, CASE=34, DEFAULT=35, 
		WHILE=36, FOR=37, IN=38, RANGO=39, GUARD=40, CONTINUE=41, RETURN=42, BREAK=43, 
		NIL=44, INT=45, FLOAT=46, BOOL=47, CHARACTER=48, STRING=49, ENBLANCO=50, 
		ENTERO=51, DECIMAL=52, CARACTER=53, CADENA=54, ID=55, UL_C=56, ML_C=57;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "PTCOMA", 
			"INTERROGACION", "MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", 
			"MAS", "POR", "IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
			"MAYOR", "MENOR", "AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", 
			"FALSE", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", 
			"RANGO", "GUARD", "CONTINUE", "RETURN", "BREAK", "NIL", "INT", "FLOAT", 
			"BOOL", "CHARACTER", "STRING", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", 
			"CARACTER", "CADENA", "ID", "UL_C", "ML_C"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "')'", "'('", "'{'", "'}'", "':'", "';'", "'?'", "'+='", 
			"'-='", "'='", "'/'", "'%'", "'-'", "'+'", "'*'", "'=='", "'!='", "'>='", 
			"'<='", "'>'", "'<'", "'&&'", "'||'", "'!'", "'print'", "'var'", "'let'", 
			"'true'", "'false'", "'if'", "'else'", "'switch'", "'case'", "'default'", 
			"'while'", "'for'", "'in'", "'...'", "'guard'", "'continue'", "'return'", 
			"'break'", "'nil'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "PTCOMA", 
			"INTERROGACION", "MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", 
			"MAS", "POR", "IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
			"MAYOR", "MENOR", "AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", 
			"FALSE", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", 
			"RANGO", "GUARD", "CONTINUE", "RETURN", "BREAK", "NIL", "INT", "FLOAT", 
			"BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", 
			"CADENA", "ID", "UL_C", "ML_C"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2;\u018b\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3\3"+
		"\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22"+
		"\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\27"+
		"\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%"+
		"\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*"+
		"\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3-\3-\3-\3-\3.\3."+
		"\3.\3.\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\63"+
		"\6\63\u0139\n\63\r\63\16\63\u013a\3\63\3\63\3\64\3\64\3\65\6\65\u0142"+
		"\n\65\r\65\16\65\u0143\3\66\6\66\u0147\n\66\r\66\16\66\u0148\3\66\3\66"+
		"\6\66\u014d\n\66\r\66\16\66\u014e\3\67\3\67\3\67\3\67\5\67\u0155\n\67"+
		"\3\67\3\67\38\38\38\38\78\u015d\n8\f8\168\u0160\138\38\38\39\39\79\u0166"+
		"\n9\f9\169\u0169\139\39\39\69\u016d\n9\r9\169\u016e\59\u0171\n9\3:\3:"+
		"\3:\3:\7:\u0177\n:\f:\16:\u017a\13:\3:\3:\3;\3;\3;\3;\7;\u0182\n;\f;\16"+
		";\u0185\13;\3;\3;\3;\3;\3;\3\u0183\2<\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.["+
		"/]\60_\61a\62c\63e\64g\2i\65k\66m\67o8q9s:u;\3\2\t\5\2\13\f\17\17\"\""+
		"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17\2\u0195"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2"+
		"\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\3w\3\2\2\2\5y\3\2\2\2\7{\3\2\2\2\t"+
		"}\3\2\2\2\13\177\3\2\2\2\r\u0081\3\2\2\2\17\u0083\3\2\2\2\21\u0085\3\2"+
		"\2\2\23\u0087\3\2\2\2\25\u008a\3\2\2\2\27\u008d\3\2\2\2\31\u008f\3\2\2"+
		"\2\33\u0091\3\2\2\2\35\u0093\3\2\2\2\37\u0095\3\2\2\2!\u0097\3\2\2\2#"+
		"\u0099\3\2\2\2%\u009c\3\2\2\2\'\u009f\3\2\2\2)\u00a2\3\2\2\2+\u00a5\3"+
		"\2\2\2-\u00a7\3\2\2\2/\u00a9\3\2\2\2\61\u00ac\3\2\2\2\63\u00af\3\2\2\2"+
		"\65\u00b1\3\2\2\2\67\u00b7\3\2\2\29\u00bb\3\2\2\2;\u00bf\3\2\2\2=\u00c4"+
		"\3\2\2\2?\u00ca\3\2\2\2A\u00cd\3\2\2\2C\u00d2\3\2\2\2E\u00d9\3\2\2\2G"+
		"\u00de\3\2\2\2I\u00e6\3\2\2\2K\u00ec\3\2\2\2M\u00f0\3\2\2\2O\u00f3\3\2"+
		"\2\2Q\u00f7\3\2\2\2S\u00fd\3\2\2\2U\u0106\3\2\2\2W\u010d\3\2\2\2Y\u0113"+
		"\3\2\2\2[\u0117\3\2\2\2]\u011b\3\2\2\2_\u0121\3\2\2\2a\u0126\3\2\2\2c"+
		"\u0130\3\2\2\2e\u0138\3\2\2\2g\u013e\3\2\2\2i\u0141\3\2\2\2k\u0146\3\2"+
		"\2\2m\u0150\3\2\2\2o\u0158\3\2\2\2q\u0170\3\2\2\2s\u0172\3\2\2\2u\u017d"+
		"\3\2\2\2wx\7.\2\2x\4\3\2\2\2yz\7+\2\2z\6\3\2\2\2{|\7*\2\2|\b\3\2\2\2}"+
		"~\7}\2\2~\n\3\2\2\2\177\u0080\7\177\2\2\u0080\f\3\2\2\2\u0081\u0082\7"+
		"<\2\2\u0082\16\3\2\2\2\u0083\u0084\7=\2\2\u0084\20\3\2\2\2\u0085\u0086"+
		"\7A\2\2\u0086\22\3\2\2\2\u0087\u0088\7-\2\2\u0088\u0089\7?\2\2\u0089\24"+
		"\3\2\2\2\u008a\u008b\7/\2\2\u008b\u008c\7?\2\2\u008c\26\3\2\2\2\u008d"+
		"\u008e\7?\2\2\u008e\30\3\2\2\2\u008f\u0090\7\61\2\2\u0090\32\3\2\2\2\u0091"+
		"\u0092\7\'\2\2\u0092\34\3\2\2\2\u0093\u0094\7/\2\2\u0094\36\3\2\2\2\u0095"+
		"\u0096\7-\2\2\u0096 \3\2\2\2\u0097\u0098\7,\2\2\u0098\"\3\2\2\2\u0099"+
		"\u009a\7?\2\2\u009a\u009b\7?\2\2\u009b$\3\2\2\2\u009c\u009d\7#\2\2\u009d"+
		"\u009e\7?\2\2\u009e&\3\2\2\2\u009f\u00a0\7@\2\2\u00a0\u00a1\7?\2\2\u00a1"+
		"(\3\2\2\2\u00a2\u00a3\7>\2\2\u00a3\u00a4\7?\2\2\u00a4*\3\2\2\2\u00a5\u00a6"+
		"\7@\2\2\u00a6,\3\2\2\2\u00a7\u00a8\7>\2\2\u00a8.\3\2\2\2\u00a9\u00aa\7"+
		"(\2\2\u00aa\u00ab\7(\2\2\u00ab\60\3\2\2\2\u00ac\u00ad\7~\2\2\u00ad\u00ae"+
		"\7~\2\2\u00ae\62\3\2\2\2\u00af\u00b0\7#\2\2\u00b0\64\3\2\2\2\u00b1\u00b2"+
		"\7r\2\2\u00b2\u00b3\7t\2\2\u00b3\u00b4\7k\2\2\u00b4\u00b5\7p\2\2\u00b5"+
		"\u00b6\7v\2\2\u00b6\66\3\2\2\2\u00b7\u00b8\7x\2\2\u00b8\u00b9\7c\2\2\u00b9"+
		"\u00ba\7t\2\2\u00ba8\3\2\2\2\u00bb\u00bc\7n\2\2\u00bc\u00bd\7g\2\2\u00bd"+
		"\u00be\7v\2\2\u00be:\3\2\2\2\u00bf\u00c0\7v\2\2\u00c0\u00c1\7t\2\2\u00c1"+
		"\u00c2\7w\2\2\u00c2\u00c3\7g\2\2\u00c3<\3\2\2\2\u00c4\u00c5\7h\2\2\u00c5"+
		"\u00c6\7c\2\2\u00c6\u00c7\7n\2\2\u00c7\u00c8\7u\2\2\u00c8\u00c9\7g\2\2"+
		"\u00c9>\3\2\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc\7h\2\2\u00cc@\3\2\2\2\u00cd"+
		"\u00ce\7g\2\2\u00ce\u00cf\7n\2\2\u00cf\u00d0\7u\2\2\u00d0\u00d1\7g\2\2"+
		"\u00d1B\3\2\2\2\u00d2\u00d3\7u\2\2\u00d3\u00d4\7y\2\2\u00d4\u00d5\7k\2"+
		"\2\u00d5\u00d6\7v\2\2\u00d6\u00d7\7e\2\2\u00d7\u00d8\7j\2\2\u00d8D\3\2"+
		"\2\2\u00d9\u00da\7e\2\2\u00da\u00db\7c\2\2\u00db\u00dc\7u\2\2\u00dc\u00dd"+
		"\7g\2\2\u00ddF\3\2\2\2\u00de\u00df\7f\2\2\u00df\u00e0\7g\2\2\u00e0\u00e1"+
		"\7h\2\2\u00e1\u00e2\7c\2\2\u00e2\u00e3\7w\2\2\u00e3\u00e4\7n\2\2\u00e4"+
		"\u00e5\7v\2\2\u00e5H\3\2\2\2\u00e6\u00e7\7y\2\2\u00e7\u00e8\7j\2\2\u00e8"+
		"\u00e9\7k\2\2\u00e9\u00ea\7n\2\2\u00ea\u00eb\7g\2\2\u00ebJ\3\2\2\2\u00ec"+
		"\u00ed\7h\2\2\u00ed\u00ee\7q\2\2\u00ee\u00ef\7t\2\2\u00efL\3\2\2\2\u00f0"+
		"\u00f1\7k\2\2\u00f1\u00f2\7p\2\2\u00f2N\3\2\2\2\u00f3\u00f4\7\60\2\2\u00f4"+
		"\u00f5\7\60\2\2\u00f5\u00f6\7\60\2\2\u00f6P\3\2\2\2\u00f7\u00f8\7i\2\2"+
		"\u00f8\u00f9\7w\2\2\u00f9\u00fa\7c\2\2\u00fa\u00fb\7t\2\2\u00fb\u00fc"+
		"\7f\2\2\u00fcR\3\2\2\2\u00fd\u00fe\7e\2\2\u00fe\u00ff\7q\2\2\u00ff\u0100"+
		"\7p\2\2\u0100\u0101\7v\2\2\u0101\u0102\7k\2\2\u0102\u0103\7p\2\2\u0103"+
		"\u0104\7w\2\2\u0104\u0105\7g\2\2\u0105T\3\2\2\2\u0106\u0107\7t\2\2\u0107"+
		"\u0108\7g\2\2\u0108\u0109\7v\2\2\u0109\u010a\7w\2\2\u010a\u010b\7t\2\2"+
		"\u010b\u010c\7p\2\2\u010cV\3\2\2\2\u010d\u010e\7d\2\2\u010e\u010f\7t\2"+
		"\2\u010f\u0110\7g\2\2\u0110\u0111\7c\2\2\u0111\u0112\7m\2\2\u0112X\3\2"+
		"\2\2\u0113\u0114\7p\2\2\u0114\u0115\7k\2\2\u0115\u0116\7n\2\2\u0116Z\3"+
		"\2\2\2\u0117\u0118\7K\2\2\u0118\u0119\7p\2\2\u0119\u011a\7v\2\2\u011a"+
		"\\\3\2\2\2\u011b\u011c\7H\2\2\u011c\u011d\7n\2\2\u011d\u011e\7q\2\2\u011e"+
		"\u011f\7c\2\2\u011f\u0120\7v\2\2\u0120^\3\2\2\2\u0121\u0122\7D\2\2\u0122"+
		"\u0123\7q\2\2\u0123\u0124\7q\2\2\u0124\u0125\7n\2\2\u0125`\3\2\2\2\u0126"+
		"\u0127\7E\2\2\u0127\u0128\7j\2\2\u0128\u0129\7c\2\2\u0129\u012a\7t\2\2"+
		"\u012a\u012b\7c\2\2\u012b\u012c\7e\2\2\u012c\u012d\7v\2\2\u012d\u012e"+
		"\7g\2\2\u012e\u012f\7t\2\2\u012fb\3\2\2\2\u0130\u0131\7U\2\2\u0131\u0132"+
		"\7v\2\2\u0132\u0133\7t\2\2\u0133\u0134\7k\2\2\u0134\u0135\7p\2\2\u0135"+
		"\u0136\7i\2\2\u0136d\3\2\2\2\u0137\u0139\t\2\2\2\u0138\u0137\3\2\2\2\u0139"+
		"\u013a\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013c\3\2"+
		"\2\2\u013c\u013d\b\63\2\2\u013df\3\2\2\2\u013e\u013f\t\3\2\2\u013fh\3"+
		"\2\2\2\u0140\u0142\5g\64\2\u0141\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		"\u0141\3\2\2\2\u0143\u0144\3\2\2\2\u0144j\3\2\2\2\u0145\u0147\5g\64\2"+
		"\u0146\u0145\3\2\2\2\u0147\u0148\3\2\2\2\u0148\u0146\3\2\2\2\u0148\u0149"+
		"\3\2\2\2\u0149\u014a\3\2\2\2\u014a\u014c\7\60\2\2\u014b\u014d\5g\64\2"+
		"\u014c\u014b\3\2\2\2\u014d\u014e\3\2\2\2\u014e\u014c\3\2\2\2\u014e\u014f"+
		"\3\2\2\2\u014fl\3\2\2\2\u0150\u0154\7$\2\2\u0151\u0155\n\4\2\2\u0152\u0153"+
		"\7^\2\2\u0153\u0155\13\2\2\2\u0154\u0151\3\2\2\2\u0154\u0152\3\2\2\2\u0155"+
		"\u0156\3\2\2\2\u0156\u0157\7$\2\2\u0157n\3\2\2\2\u0158\u015e\7$\2\2\u0159"+
		"\u015d\n\5\2\2\u015a\u015b\7^\2\2\u015b\u015d\13\2\2\2\u015c\u0159\3\2"+
		"\2\2\u015c\u015a\3\2\2\2\u015d\u0160\3\2\2\2\u015e\u015c\3\2\2\2\u015e"+
		"\u015f\3\2\2\2\u015f\u0161\3\2\2\2\u0160\u015e\3\2\2\2\u0161\u0162\7$"+
		"\2\2\u0162p\3\2\2\2\u0163\u0167\t\6\2\2\u0164\u0166\t\7\2\2\u0165\u0164"+
		"\3\2\2\2\u0166\u0169\3\2\2\2\u0167\u0165\3\2\2\2\u0167\u0168\3\2\2\2\u0168"+
		"\u0171\3\2\2\2\u0169\u0167\3\2\2\2\u016a\u016c\7a\2\2\u016b\u016d\t\7"+
		"\2\2\u016c\u016b\3\2\2\2\u016d\u016e\3\2\2\2\u016e\u016c\3\2\2\2\u016e"+
		"\u016f\3\2\2\2\u016f\u0171\3\2\2\2\u0170\u0163\3\2\2\2\u0170\u016a\3\2"+
		"\2\2\u0171r\3\2\2\2\u0172\u0173\7\61\2\2\u0173\u0174\7\61\2\2\u0174\u0178"+
		"\3\2\2\2\u0175\u0177\n\b\2\2\u0176\u0175\3\2\2\2\u0177\u017a\3\2\2\2\u0178"+
		"\u0176\3\2\2\2\u0178\u0179\3\2\2\2\u0179\u017b\3\2\2\2\u017a\u0178\3\2"+
		"\2\2\u017b\u017c\b:\3\2\u017ct\3\2\2\2\u017d\u017e\7\61\2\2\u017e\u017f"+
		"\7,\2\2\u017f\u0183\3\2\2\2\u0180\u0182\13\2\2\2\u0181\u0180\3\2\2\2\u0182"+
		"\u0185\3\2\2\2\u0183\u0184\3\2\2\2\u0183\u0181\3\2\2\2\u0184\u0186\3\2"+
		"\2\2\u0185\u0183\3\2\2\2\u0186\u0187\7,\2\2\u0187\u0188\7\61\2\2\u0188"+
		"\u0189\3\2\2\2\u0189\u018a\b;\3\2\u018av\3\2\2\2\17\2\u013a\u0143\u0148"+
		"\u014e\u0154\u015c\u015e\u0167\u016e\u0170\u0178\u0183\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}