// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/OLC2_Proyecto1_202111478/parser/Tswift_Lexer.g4 by ANTLR 4.9.2
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
		PARDER=1, PARIZQ=2, LLAVEIZQ=3, LLAVEDER=4, CORCHETEIZQ=5, CORCHETEDER=6, 
		DOSPT=7, PTCOMA=8, INTERROGACION=9, PUNTO=10, MASIGUAL=11, MENOSIGUAL=12, 
		IGUAL=13, DIV=14, MOD=15, MENOS=16, MAS=17, POR=18, IGUALIGUAL=19, DIFERENTE=20, 
		MAYORIGUAL=21, MENORIGUAL=22, MAYOR=23, MENOR=24, AND=25, OR=26, NOT=27, 
		PRINT=28, VAR=29, LET=30, TRUE=31, FALSE=32, IF=33, ELSE=34, SWITCH=35, 
		CASE=36, DEFAULT=37, WHILE=38, FOR=39, IN=40, RANGO=41, GUARD=42, CONTINUE=43, 
		RETURN=44, BREAK=45, NIL=46, APPEND=47, REMOVELAST=48, REMOVE=49, AT=50, 
		ISEMPTY=51, COUNT=52, INT=53, FLOAT=54, BOOL=55, CHARACTER=56, STRING=57, 
		ENBLANCO=58, ENTERO=59, DECIMAL=60, CARACTER=61, CADENA=62, ID=63, UL_C=64, 
		ML_C=65;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "MASIGUAL", "MENOSIGUAL", 
			"IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", "DIFERENTE", 
			"MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", "NOT", "PRINT", 
			"VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "INT", 
			"FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "DIG", "ENTERO", 
			"DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", 
			"'.'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'", "'*'", "'=='", 
			"'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", "'!'", "'print'", 
			"'var'", "'let'", "'true'", "'false'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", 
			"'continue'", "'return'", "'break'", "'nil'", "'append'", "'removeLast'", 
			"'remove'", "'at'", "'IsEmpty'", "'count'", "'Int'", "'Float'", "'Bool'", 
			"'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "MASIGUAL", "MENOSIGUAL", 
			"IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", "DIFERENTE", 
			"MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", "NOT", "PRINT", 
			"VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "INT", 
			"FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO", "DECIMAL", 
			"CARACTER", "CADENA", "ID", "UL_C", "ML_C"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2C\u01c9\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3"+
		"\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3"+
		"\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3"+
		"\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\31\3"+
		"\31\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3"+
		"!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3"+
		"%\3&\3&\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3"+
		")\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3,\3-\3-\3-\3"+
		"-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\60"+
		"\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62"+
		"\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\38\38\38\38\38\39\39\39\39\39\39\39\39\39\39\3:\3"+
		":\3:\3:\3:\3:\3:\3;\6;\u0177\n;\r;\16;\u0178\3;\3;\3<\3<\3=\6=\u0180\n"+
		"=\r=\16=\u0181\3>\6>\u0185\n>\r>\16>\u0186\3>\3>\6>\u018b\n>\r>\16>\u018c"+
		"\3?\3?\3?\3?\5?\u0193\n?\3?\3?\3@\3@\3@\3@\7@\u019b\n@\f@\16@\u019e\13"+
		"@\3@\3@\3A\3A\7A\u01a4\nA\fA\16A\u01a7\13A\3A\3A\6A\u01ab\nA\rA\16A\u01ac"+
		"\5A\u01af\nA\3B\3B\3B\3B\7B\u01b5\nB\fB\16B\u01b8\13B\3B\3B\3C\3C\3C\3"+
		"C\7C\u01c0\nC\fC\16C\u01c3\13C\3C\3C\3C\3C\3C\3\u01c1\2D\3\3\5\4\7\5\t"+
		"\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w\2y="+
		"{>}?\177@\u0081A\u0083B\u0085C\3\2\t\5\2\13\f\17\17\"\"\3\2\62;\4\2))"+
		"^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17\2\u01d3\2\3\3\2\2\2"+
		"\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2"+
		"\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2"+
		"\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2"+
		"\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2"+
		"\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2"+
		"\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2"+
		"\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W"+
		"\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2"+
		"\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2"+
		"\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177"+
		"\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\3\u0087\3\2\2"+
		"\2\5\u0089\3\2\2\2\7\u008b\3\2\2\2\t\u008d\3\2\2\2\13\u008f\3\2\2\2\r"+
		"\u0091\3\2\2\2\17\u0093\3\2\2\2\21\u0095\3\2\2\2\23\u0097\3\2\2\2\25\u0099"+
		"\3\2\2\2\27\u009b\3\2\2\2\31\u009e\3\2\2\2\33\u00a1\3\2\2\2\35\u00a3\3"+
		"\2\2\2\37\u00a5\3\2\2\2!\u00a7\3\2\2\2#\u00a9\3\2\2\2%\u00ab\3\2\2\2\'"+
		"\u00ad\3\2\2\2)\u00b0\3\2\2\2+\u00b3\3\2\2\2-\u00b6\3\2\2\2/\u00b9\3\2"+
		"\2\2\61\u00bb\3\2\2\2\63\u00bd\3\2\2\2\65\u00c0\3\2\2\2\67\u00c3\3\2\2"+
		"\29\u00c5\3\2\2\2;\u00cb\3\2\2\2=\u00cf\3\2\2\2?\u00d3\3\2\2\2A\u00d8"+
		"\3\2\2\2C\u00de\3\2\2\2E\u00e1\3\2\2\2G\u00e6\3\2\2\2I\u00ed\3\2\2\2K"+
		"\u00f2\3\2\2\2M\u00fa\3\2\2\2O\u0100\3\2\2\2Q\u0104\3\2\2\2S\u0107\3\2"+
		"\2\2U\u010b\3\2\2\2W\u0111\3\2\2\2Y\u011a\3\2\2\2[\u0121\3\2\2\2]\u0127"+
		"\3\2\2\2_\u012b\3\2\2\2a\u0132\3\2\2\2c\u013d\3\2\2\2e\u0144\3\2\2\2g"+
		"\u0147\3\2\2\2i\u014f\3\2\2\2k\u0155\3\2\2\2m\u0159\3\2\2\2o\u015f\3\2"+
		"\2\2q\u0164\3\2\2\2s\u016e\3\2\2\2u\u0176\3\2\2\2w\u017c\3\2\2\2y\u017f"+
		"\3\2\2\2{\u0184\3\2\2\2}\u018e\3\2\2\2\177\u0196\3\2\2\2\u0081\u01ae\3"+
		"\2\2\2\u0083\u01b0\3\2\2\2\u0085\u01bb\3\2\2\2\u0087\u0088\7+\2\2\u0088"+
		"\4\3\2\2\2\u0089\u008a\7*\2\2\u008a\6\3\2\2\2\u008b\u008c\7}\2\2\u008c"+
		"\b\3\2\2\2\u008d\u008e\7\177\2\2\u008e\n\3\2\2\2\u008f\u0090\7]\2\2\u0090"+
		"\f\3\2\2\2\u0091\u0092\7_\2\2\u0092\16\3\2\2\2\u0093\u0094\7<\2\2\u0094"+
		"\20\3\2\2\2\u0095\u0096\7=\2\2\u0096\22\3\2\2\2\u0097\u0098\7A\2\2\u0098"+
		"\24\3\2\2\2\u0099\u009a\7\60\2\2\u009a\26\3\2\2\2\u009b\u009c\7-\2\2\u009c"+
		"\u009d\7?\2\2\u009d\30\3\2\2\2\u009e\u009f\7/\2\2\u009f\u00a0\7?\2\2\u00a0"+
		"\32\3\2\2\2\u00a1\u00a2\7?\2\2\u00a2\34\3\2\2\2\u00a3\u00a4\7\61\2\2\u00a4"+
		"\36\3\2\2\2\u00a5\u00a6\7\'\2\2\u00a6 \3\2\2\2\u00a7\u00a8\7/\2\2\u00a8"+
		"\"\3\2\2\2\u00a9\u00aa\7-\2\2\u00aa$\3\2\2\2\u00ab\u00ac\7,\2\2\u00ac"+
		"&\3\2\2\2\u00ad\u00ae\7?\2\2\u00ae\u00af\7?\2\2\u00af(\3\2\2\2\u00b0\u00b1"+
		"\7#\2\2\u00b1\u00b2\7?\2\2\u00b2*\3\2\2\2\u00b3\u00b4\7@\2\2\u00b4\u00b5"+
		"\7?\2\2\u00b5,\3\2\2\2\u00b6\u00b7\7>\2\2\u00b7\u00b8\7?\2\2\u00b8.\3"+
		"\2\2\2\u00b9\u00ba\7@\2\2\u00ba\60\3\2\2\2\u00bb\u00bc\7>\2\2\u00bc\62"+
		"\3\2\2\2\u00bd\u00be\7(\2\2\u00be\u00bf\7(\2\2\u00bf\64\3\2\2\2\u00c0"+
		"\u00c1\7~\2\2\u00c1\u00c2\7~\2\2\u00c2\66\3\2\2\2\u00c3\u00c4\7#\2\2\u00c4"+
		"8\3\2\2\2\u00c5\u00c6\7r\2\2\u00c6\u00c7\7t\2\2\u00c7\u00c8\7k\2\2\u00c8"+
		"\u00c9\7p\2\2\u00c9\u00ca\7v\2\2\u00ca:\3\2\2\2\u00cb\u00cc\7x\2\2\u00cc"+
		"\u00cd\7c\2\2\u00cd\u00ce\7t\2\2\u00ce<\3\2\2\2\u00cf\u00d0\7n\2\2\u00d0"+
		"\u00d1\7g\2\2\u00d1\u00d2\7v\2\2\u00d2>\3\2\2\2\u00d3\u00d4\7v\2\2\u00d4"+
		"\u00d5\7t\2\2\u00d5\u00d6\7w\2\2\u00d6\u00d7\7g\2\2\u00d7@\3\2\2\2\u00d8"+
		"\u00d9\7h\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7n\2\2\u00db\u00dc\7u\2\2"+
		"\u00dc\u00dd\7g\2\2\u00ddB\3\2\2\2\u00de\u00df\7k\2\2\u00df\u00e0\7h\2"+
		"\2\u00e0D\3\2\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7n\2\2\u00e3\u00e4\7"+
		"u\2\2\u00e4\u00e5\7g\2\2\u00e5F\3\2\2\2\u00e6\u00e7\7u\2\2\u00e7\u00e8"+
		"\7y\2\2\u00e8\u00e9\7k\2\2\u00e9\u00ea\7v\2\2\u00ea\u00eb\7e\2\2\u00eb"+
		"\u00ec\7j\2\2\u00ecH\3\2\2\2\u00ed\u00ee\7e\2\2\u00ee\u00ef\7c\2\2\u00ef"+
		"\u00f0\7u\2\2\u00f0\u00f1\7g\2\2\u00f1J\3\2\2\2\u00f2\u00f3\7f\2\2\u00f3"+
		"\u00f4\7g\2\2\u00f4\u00f5\7h\2\2\u00f5\u00f6\7c\2\2\u00f6\u00f7\7w\2\2"+
		"\u00f7\u00f8\7n\2\2\u00f8\u00f9\7v\2\2\u00f9L\3\2\2\2\u00fa\u00fb\7y\2"+
		"\2\u00fb\u00fc\7j\2\2\u00fc\u00fd\7k\2\2\u00fd\u00fe\7n\2\2\u00fe\u00ff"+
		"\7g\2\2\u00ffN\3\2\2\2\u0100\u0101\7h\2\2\u0101\u0102\7q\2\2\u0102\u0103"+
		"\7t\2\2\u0103P\3\2\2\2\u0104\u0105\7k\2\2\u0105\u0106\7p\2\2\u0106R\3"+
		"\2\2\2\u0107\u0108\7\60\2\2\u0108\u0109\7\60\2\2\u0109\u010a\7\60\2\2"+
		"\u010aT\3\2\2\2\u010b\u010c\7i\2\2\u010c\u010d\7w\2\2\u010d\u010e\7c\2"+
		"\2\u010e\u010f\7t\2\2\u010f\u0110\7f\2\2\u0110V\3\2\2\2\u0111\u0112\7"+
		"e\2\2\u0112\u0113\7q\2\2\u0113\u0114\7p\2\2\u0114\u0115\7v\2\2\u0115\u0116"+
		"\7k\2\2\u0116\u0117\7p\2\2\u0117\u0118\7w\2\2\u0118\u0119\7g\2\2\u0119"+
		"X\3\2\2\2\u011a\u011b\7t\2\2\u011b\u011c\7g\2\2\u011c\u011d\7v\2\2\u011d"+
		"\u011e\7w\2\2\u011e\u011f\7t\2\2\u011f\u0120\7p\2\2\u0120Z\3\2\2\2\u0121"+
		"\u0122\7d\2\2\u0122\u0123\7t\2\2\u0123\u0124\7g\2\2\u0124\u0125\7c\2\2"+
		"\u0125\u0126\7m\2\2\u0126\\\3\2\2\2\u0127\u0128\7p\2\2\u0128\u0129\7k"+
		"\2\2\u0129\u012a\7n\2\2\u012a^\3\2\2\2\u012b\u012c\7c\2\2\u012c\u012d"+
		"\7r\2\2\u012d\u012e\7r\2\2\u012e\u012f\7g\2\2\u012f\u0130\7p\2\2\u0130"+
		"\u0131\7f\2\2\u0131`\3\2\2\2\u0132\u0133\7t\2\2\u0133\u0134\7g\2\2\u0134"+
		"\u0135\7o\2\2\u0135\u0136\7q\2\2\u0136\u0137\7x\2\2\u0137\u0138\7g\2\2"+
		"\u0138\u0139\7N\2\2\u0139\u013a\7c\2\2\u013a\u013b\7u\2\2\u013b\u013c"+
		"\7v\2\2\u013cb\3\2\2\2\u013d\u013e\7t\2\2\u013e\u013f\7g\2\2\u013f\u0140"+
		"\7o\2\2\u0140\u0141\7q\2\2\u0141\u0142\7x\2\2\u0142\u0143\7g\2\2\u0143"+
		"d\3\2\2\2\u0144\u0145\7c\2\2\u0145\u0146\7v\2\2\u0146f\3\2\2\2\u0147\u0148"+
		"\7K\2\2\u0148\u0149\7u\2\2\u0149\u014a\7G\2\2\u014a\u014b\7o\2\2\u014b"+
		"\u014c\7r\2\2\u014c\u014d\7v\2\2\u014d\u014e\7{\2\2\u014eh\3\2\2\2\u014f"+
		"\u0150\7e\2\2\u0150\u0151\7q\2\2\u0151\u0152\7w\2\2\u0152\u0153\7p\2\2"+
		"\u0153\u0154\7v\2\2\u0154j\3\2\2\2\u0155\u0156\7K\2\2\u0156\u0157\7p\2"+
		"\2\u0157\u0158\7v\2\2\u0158l\3\2\2\2\u0159\u015a\7H\2\2\u015a\u015b\7"+
		"n\2\2\u015b\u015c\7q\2\2\u015c\u015d\7c\2\2\u015d\u015e\7v\2\2\u015en"+
		"\3\2\2\2\u015f\u0160\7D\2\2\u0160\u0161\7q\2\2\u0161\u0162\7q\2\2\u0162"+
		"\u0163\7n\2\2\u0163p\3\2\2\2\u0164\u0165\7E\2\2\u0165\u0166\7j\2\2\u0166"+
		"\u0167\7c\2\2\u0167\u0168\7t\2\2\u0168\u0169\7c\2\2\u0169\u016a\7e\2\2"+
		"\u016a\u016b\7v\2\2\u016b\u016c\7g\2\2\u016c\u016d\7t\2\2\u016dr\3\2\2"+
		"\2\u016e\u016f\7U\2\2\u016f\u0170\7v\2\2\u0170\u0171\7t\2\2\u0171\u0172"+
		"\7k\2\2\u0172\u0173\7p\2\2\u0173\u0174\7i\2\2\u0174t\3\2\2\2\u0175\u0177"+
		"\t\2\2\2\u0176\u0175\3\2\2\2\u0177\u0178\3\2\2\2\u0178\u0176\3\2\2\2\u0178"+
		"\u0179\3\2\2\2\u0179\u017a\3\2\2\2\u017a\u017b\b;\2\2\u017bv\3\2\2\2\u017c"+
		"\u017d\t\3\2\2\u017dx\3\2\2\2\u017e\u0180\5w<\2\u017f\u017e\3\2\2\2\u0180"+
		"\u0181\3\2\2\2\u0181\u017f\3\2\2\2\u0181\u0182\3\2\2\2\u0182z\3\2\2\2"+
		"\u0183\u0185\5w<\2\u0184\u0183\3\2\2\2\u0185\u0186\3\2\2\2\u0186\u0184"+
		"\3\2\2\2\u0186\u0187\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u018a\7\60\2\2"+
		"\u0189\u018b\5w<\2\u018a\u0189\3\2\2\2\u018b\u018c\3\2\2\2\u018c\u018a"+
		"\3\2\2\2\u018c\u018d\3\2\2\2\u018d|\3\2\2\2\u018e\u0192\7$\2\2\u018f\u0193"+
		"\n\4\2\2\u0190\u0191\7^\2\2\u0191\u0193\13\2\2\2\u0192\u018f\3\2\2\2\u0192"+
		"\u0190\3\2\2\2\u0193\u0194\3\2\2\2\u0194\u0195\7$\2\2\u0195~\3\2\2\2\u0196"+
		"\u019c\7$\2\2\u0197\u019b\n\5\2\2\u0198\u0199\7^\2\2\u0199\u019b\13\2"+
		"\2\2\u019a\u0197\3\2\2\2\u019a\u0198\3\2\2\2\u019b\u019e\3\2\2\2\u019c"+
		"\u019a\3\2\2\2\u019c\u019d\3\2\2\2\u019d\u019f\3\2\2\2\u019e\u019c\3\2"+
		"\2\2\u019f\u01a0\7$\2\2\u01a0\u0080\3\2\2\2\u01a1\u01a5\t\6\2\2\u01a2"+
		"\u01a4\t\7\2\2\u01a3\u01a2\3\2\2\2\u01a4\u01a7\3\2\2\2\u01a5\u01a3\3\2"+
		"\2\2\u01a5\u01a6\3\2\2\2\u01a6\u01af\3\2\2\2\u01a7\u01a5\3\2\2\2\u01a8"+
		"\u01aa\7a\2\2\u01a9\u01ab\t\7\2\2\u01aa\u01a9\3\2\2\2\u01ab\u01ac\3\2"+
		"\2\2\u01ac\u01aa\3\2\2\2\u01ac\u01ad\3\2\2\2\u01ad\u01af\3\2\2\2\u01ae"+
		"\u01a1\3\2\2\2\u01ae\u01a8\3\2\2\2\u01af\u0082\3\2\2\2\u01b0\u01b1\7\61"+
		"\2\2\u01b1\u01b2\7\61\2\2\u01b2\u01b6\3\2\2\2\u01b3\u01b5\n\b\2\2\u01b4"+
		"\u01b3\3\2\2\2\u01b5\u01b8\3\2\2\2\u01b6\u01b4\3\2\2\2\u01b6\u01b7\3\2"+
		"\2\2\u01b7\u01b9\3\2\2\2\u01b8\u01b6\3\2\2\2\u01b9\u01ba\bB\3\2\u01ba"+
		"\u0084\3\2\2\2\u01bb\u01bc\7\61\2\2\u01bc\u01bd\7,\2\2\u01bd\u01c1\3\2"+
		"\2\2\u01be\u01c0\13\2\2\2\u01bf\u01be\3\2\2\2\u01c0\u01c3\3\2\2\2\u01c1"+
		"\u01c2\3\2\2\2\u01c1\u01bf\3\2\2\2\u01c2\u01c4\3\2\2\2\u01c3\u01c1\3\2"+
		"\2\2\u01c4\u01c5\7,\2\2\u01c5\u01c6\7\61\2\2\u01c6\u01c7\3\2\2\2\u01c7"+
		"\u01c8\bC\3\2\u01c8\u0086\3\2\2\2\17\2\u0178\u0181\u0186\u018c\u0192\u019a"+
		"\u019c\u01a5\u01ac\u01ae\u01b6\u01c1\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}