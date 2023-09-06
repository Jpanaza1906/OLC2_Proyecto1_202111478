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
		ML_C=65, ERROR=66;
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
			"DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR"
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
			"CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2D\u01cf\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\3\2\3\2\3\3\3\3\3\4\3\4\3\5"+
		"\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\r"+
		"\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3"+
		"\31\3\31\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!"+
		"\3!\3!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%"+
		"\3%\3%\3&\3&\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)"+
		"\3)\3)\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3,\3-\3-"+
		"\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60"+
		"\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62"+
		"\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\39\39\39\39\39\39\39\39\39\3"+
		"9\3:\3:\3:\3:\3:\3:\3:\3;\6;\u0179\n;\r;\16;\u017a\3;\3;\3<\3<\3=\6=\u0182"+
		"\n=\r=\16=\u0183\3>\6>\u0187\n>\r>\16>\u0188\3>\3>\6>\u018d\n>\r>\16>"+
		"\u018e\3?\3?\3?\3?\5?\u0195\n?\3?\3?\3@\3@\3@\3@\7@\u019d\n@\f@\16@\u01a0"+
		"\13@\3@\3@\3A\3A\7A\u01a6\nA\fA\16A\u01a9\13A\3A\3A\6A\u01ad\nA\rA\16"+
		"A\u01ae\5A\u01b1\nA\3B\3B\3B\3B\7B\u01b7\nB\fB\16B\u01ba\13B\3B\3B\3C"+
		"\3C\3C\3C\7C\u01c2\nC\fC\16C\u01c5\13C\3C\3C\3C\3C\3C\3D\3D\3D\3D\3\u01c3"+
		"\2E\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w\2y={>}?\177@\u0081A\u0083B\u0085C\u0087D\3\2\t\5\2\13\f\17"+
		"\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17"+
		"\2\u01d9\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2"+
		"\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2y\3\2\2\2\2{"+
		"\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085"+
		"\3\2\2\2\2\u0087\3\2\2\2\3\u0089\3\2\2\2\5\u008b\3\2\2\2\7\u008d\3\2\2"+
		"\2\t\u008f\3\2\2\2\13\u0091\3\2\2\2\r\u0093\3\2\2\2\17\u0095\3\2\2\2\21"+
		"\u0097\3\2\2\2\23\u0099\3\2\2\2\25\u009b\3\2\2\2\27\u009d\3\2\2\2\31\u00a0"+
		"\3\2\2\2\33\u00a3\3\2\2\2\35\u00a5\3\2\2\2\37\u00a7\3\2\2\2!\u00a9\3\2"+
		"\2\2#\u00ab\3\2\2\2%\u00ad\3\2\2\2\'\u00af\3\2\2\2)\u00b2\3\2\2\2+\u00b5"+
		"\3\2\2\2-\u00b8\3\2\2\2/\u00bb\3\2\2\2\61\u00bd\3\2\2\2\63\u00bf\3\2\2"+
		"\2\65\u00c2\3\2\2\2\67\u00c5\3\2\2\29\u00c7\3\2\2\2;\u00cd\3\2\2\2=\u00d1"+
		"\3\2\2\2?\u00d5\3\2\2\2A\u00da\3\2\2\2C\u00e0\3\2\2\2E\u00e3\3\2\2\2G"+
		"\u00e8\3\2\2\2I\u00ef\3\2\2\2K\u00f4\3\2\2\2M\u00fc\3\2\2\2O\u0102\3\2"+
		"\2\2Q\u0106\3\2\2\2S\u0109\3\2\2\2U\u010d\3\2\2\2W\u0113\3\2\2\2Y\u011c"+
		"\3\2\2\2[\u0123\3\2\2\2]\u0129\3\2\2\2_\u012d\3\2\2\2a\u0134\3\2\2\2c"+
		"\u013f\3\2\2\2e\u0146\3\2\2\2g\u0149\3\2\2\2i\u0151\3\2\2\2k\u0157\3\2"+
		"\2\2m\u015b\3\2\2\2o\u0161\3\2\2\2q\u0166\3\2\2\2s\u0170\3\2\2\2u\u0178"+
		"\3\2\2\2w\u017e\3\2\2\2y\u0181\3\2\2\2{\u0186\3\2\2\2}\u0190\3\2\2\2\177"+
		"\u0198\3\2\2\2\u0081\u01b0\3\2\2\2\u0083\u01b2\3\2\2\2\u0085\u01bd\3\2"+
		"\2\2\u0087\u01cb\3\2\2\2\u0089\u008a\7+\2\2\u008a\4\3\2\2\2\u008b\u008c"+
		"\7*\2\2\u008c\6\3\2\2\2\u008d\u008e\7}\2\2\u008e\b\3\2\2\2\u008f\u0090"+
		"\7\177\2\2\u0090\n\3\2\2\2\u0091\u0092\7]\2\2\u0092\f\3\2\2\2\u0093\u0094"+
		"\7_\2\2\u0094\16\3\2\2\2\u0095\u0096\7<\2\2\u0096\20\3\2\2\2\u0097\u0098"+
		"\7=\2\2\u0098\22\3\2\2\2\u0099\u009a\7A\2\2\u009a\24\3\2\2\2\u009b\u009c"+
		"\7\60\2\2\u009c\26\3\2\2\2\u009d\u009e\7-\2\2\u009e\u009f\7?\2\2\u009f"+
		"\30\3\2\2\2\u00a0\u00a1\7/\2\2\u00a1\u00a2\7?\2\2\u00a2\32\3\2\2\2\u00a3"+
		"\u00a4\7?\2\2\u00a4\34\3\2\2\2\u00a5\u00a6\7\61\2\2\u00a6\36\3\2\2\2\u00a7"+
		"\u00a8\7\'\2\2\u00a8 \3\2\2\2\u00a9\u00aa\7/\2\2\u00aa\"\3\2\2\2\u00ab"+
		"\u00ac\7-\2\2\u00ac$\3\2\2\2\u00ad\u00ae\7,\2\2\u00ae&\3\2\2\2\u00af\u00b0"+
		"\7?\2\2\u00b0\u00b1\7?\2\2\u00b1(\3\2\2\2\u00b2\u00b3\7#\2\2\u00b3\u00b4"+
		"\7?\2\2\u00b4*\3\2\2\2\u00b5\u00b6\7@\2\2\u00b6\u00b7\7?\2\2\u00b7,\3"+
		"\2\2\2\u00b8\u00b9\7>\2\2\u00b9\u00ba\7?\2\2\u00ba.\3\2\2\2\u00bb\u00bc"+
		"\7@\2\2\u00bc\60\3\2\2\2\u00bd\u00be\7>\2\2\u00be\62\3\2\2\2\u00bf\u00c0"+
		"\7(\2\2\u00c0\u00c1\7(\2\2\u00c1\64\3\2\2\2\u00c2\u00c3\7~\2\2\u00c3\u00c4"+
		"\7~\2\2\u00c4\66\3\2\2\2\u00c5\u00c6\7#\2\2\u00c68\3\2\2\2\u00c7\u00c8"+
		"\7r\2\2\u00c8\u00c9\7t\2\2\u00c9\u00ca\7k\2\2\u00ca\u00cb\7p\2\2\u00cb"+
		"\u00cc\7v\2\2\u00cc:\3\2\2\2\u00cd\u00ce\7x\2\2\u00ce\u00cf\7c\2\2\u00cf"+
		"\u00d0\7t\2\2\u00d0<\3\2\2\2\u00d1\u00d2\7n\2\2\u00d2\u00d3\7g\2\2\u00d3"+
		"\u00d4\7v\2\2\u00d4>\3\2\2\2\u00d5\u00d6\7v\2\2\u00d6\u00d7\7t\2\2\u00d7"+
		"\u00d8\7w\2\2\u00d8\u00d9\7g\2\2\u00d9@\3\2\2\2\u00da\u00db\7h\2\2\u00db"+
		"\u00dc\7c\2\2\u00dc\u00dd\7n\2\2\u00dd\u00de\7u\2\2\u00de\u00df\7g\2\2"+
		"\u00dfB\3\2\2\2\u00e0\u00e1\7k\2\2\u00e1\u00e2\7h\2\2\u00e2D\3\2\2\2\u00e3"+
		"\u00e4\7g\2\2\u00e4\u00e5\7n\2\2\u00e5\u00e6\7u\2\2\u00e6\u00e7\7g\2\2"+
		"\u00e7F\3\2\2\2\u00e8\u00e9\7u\2\2\u00e9\u00ea\7y\2\2\u00ea\u00eb\7k\2"+
		"\2\u00eb\u00ec\7v\2\2\u00ec\u00ed\7e\2\2\u00ed\u00ee\7j\2\2\u00eeH\3\2"+
		"\2\2\u00ef\u00f0\7e\2\2\u00f0\u00f1\7c\2\2\u00f1\u00f2\7u\2\2\u00f2\u00f3"+
		"\7g\2\2\u00f3J\3\2\2\2\u00f4\u00f5\7f\2\2\u00f5\u00f6\7g\2\2\u00f6\u00f7"+
		"\7h\2\2\u00f7\u00f8\7c\2\2\u00f8\u00f9\7w\2\2\u00f9\u00fa\7n\2\2\u00fa"+
		"\u00fb\7v\2\2\u00fbL\3\2\2\2\u00fc\u00fd\7y\2\2\u00fd\u00fe\7j\2\2\u00fe"+
		"\u00ff\7k\2\2\u00ff\u0100\7n\2\2\u0100\u0101\7g\2\2\u0101N\3\2\2\2\u0102"+
		"\u0103\7h\2\2\u0103\u0104\7q\2\2\u0104\u0105\7t\2\2\u0105P\3\2\2\2\u0106"+
		"\u0107\7k\2\2\u0107\u0108\7p\2\2\u0108R\3\2\2\2\u0109\u010a\7\60\2\2\u010a"+
		"\u010b\7\60\2\2\u010b\u010c\7\60\2\2\u010cT\3\2\2\2\u010d\u010e\7i\2\2"+
		"\u010e\u010f\7w\2\2\u010f\u0110\7c\2\2\u0110\u0111\7t\2\2\u0111\u0112"+
		"\7f\2\2\u0112V\3\2\2\2\u0113\u0114\7e\2\2\u0114\u0115\7q\2\2\u0115\u0116"+
		"\7p\2\2\u0116\u0117\7v\2\2\u0117\u0118\7k\2\2\u0118\u0119\7p\2\2\u0119"+
		"\u011a\7w\2\2\u011a\u011b\7g\2\2\u011bX\3\2\2\2\u011c\u011d\7t\2\2\u011d"+
		"\u011e\7g\2\2\u011e\u011f\7v\2\2\u011f\u0120\7w\2\2\u0120\u0121\7t\2\2"+
		"\u0121\u0122\7p\2\2\u0122Z\3\2\2\2\u0123\u0124\7d\2\2\u0124\u0125\7t\2"+
		"\2\u0125\u0126\7g\2\2\u0126\u0127\7c\2\2\u0127\u0128\7m\2\2\u0128\\\3"+
		"\2\2\2\u0129\u012a\7p\2\2\u012a\u012b\7k\2\2\u012b\u012c\7n\2\2\u012c"+
		"^\3\2\2\2\u012d\u012e\7c\2\2\u012e\u012f\7r\2\2\u012f\u0130\7r\2\2\u0130"+
		"\u0131\7g\2\2\u0131\u0132\7p\2\2\u0132\u0133\7f\2\2\u0133`\3\2\2\2\u0134"+
		"\u0135\7t\2\2\u0135\u0136\7g\2\2\u0136\u0137\7o\2\2\u0137\u0138\7q\2\2"+
		"\u0138\u0139\7x\2\2\u0139\u013a\7g\2\2\u013a\u013b\7N\2\2\u013b\u013c"+
		"\7c\2\2\u013c\u013d\7u\2\2\u013d\u013e\7v\2\2\u013eb\3\2\2\2\u013f\u0140"+
		"\7t\2\2\u0140\u0141\7g\2\2\u0141\u0142\7o\2\2\u0142\u0143\7q\2\2\u0143"+
		"\u0144\7x\2\2\u0144\u0145\7g\2\2\u0145d\3\2\2\2\u0146\u0147\7c\2\2\u0147"+
		"\u0148\7v\2\2\u0148f\3\2\2\2\u0149\u014a\7K\2\2\u014a\u014b\7u\2\2\u014b"+
		"\u014c\7G\2\2\u014c\u014d\7o\2\2\u014d\u014e\7r\2\2\u014e\u014f\7v\2\2"+
		"\u014f\u0150\7{\2\2\u0150h\3\2\2\2\u0151\u0152\7e\2\2\u0152\u0153\7q\2"+
		"\2\u0153\u0154\7w\2\2\u0154\u0155\7p\2\2\u0155\u0156\7v\2\2\u0156j\3\2"+
		"\2\2\u0157\u0158\7K\2\2\u0158\u0159\7p\2\2\u0159\u015a\7v\2\2\u015al\3"+
		"\2\2\2\u015b\u015c\7H\2\2\u015c\u015d\7n\2\2\u015d\u015e\7q\2\2\u015e"+
		"\u015f\7c\2\2\u015f\u0160\7v\2\2\u0160n\3\2\2\2\u0161\u0162\7D\2\2\u0162"+
		"\u0163\7q\2\2\u0163\u0164\7q\2\2\u0164\u0165\7n\2\2\u0165p\3\2\2\2\u0166"+
		"\u0167\7E\2\2\u0167\u0168\7j\2\2\u0168\u0169\7c\2\2\u0169\u016a\7t\2\2"+
		"\u016a\u016b\7c\2\2\u016b\u016c\7e\2\2\u016c\u016d\7v\2\2\u016d\u016e"+
		"\7g\2\2\u016e\u016f\7t\2\2\u016fr\3\2\2\2\u0170\u0171\7U\2\2\u0171\u0172"+
		"\7v\2\2\u0172\u0173\7t\2\2\u0173\u0174\7k\2\2\u0174\u0175\7p\2\2\u0175"+
		"\u0176\7i\2\2\u0176t\3\2\2\2\u0177\u0179\t\2\2\2\u0178\u0177\3\2\2\2\u0179"+
		"\u017a\3\2\2\2\u017a\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017b\u017c\3\2"+
		"\2\2\u017c\u017d\b;\2\2\u017dv\3\2\2\2\u017e\u017f\t\3\2\2\u017fx\3\2"+
		"\2\2\u0180\u0182\5w<\2\u0181\u0180\3\2\2\2\u0182\u0183\3\2\2\2\u0183\u0181"+
		"\3\2\2\2\u0183\u0184\3\2\2\2\u0184z\3\2\2\2\u0185\u0187\5w<\2\u0186\u0185"+
		"\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189"+
		"\u018a\3\2\2\2\u018a\u018c\7\60\2\2\u018b\u018d\5w<\2\u018c\u018b\3\2"+
		"\2\2\u018d\u018e\3\2\2\2\u018e\u018c\3\2\2\2\u018e\u018f\3\2\2\2\u018f"+
		"|\3\2\2\2\u0190\u0194\7$\2\2\u0191\u0195\n\4\2\2\u0192\u0193\7^\2\2\u0193"+
		"\u0195\13\2\2\2\u0194\u0191\3\2\2\2\u0194\u0192\3\2\2\2\u0195\u0196\3"+
		"\2\2\2\u0196\u0197\7$\2\2\u0197~\3\2\2\2\u0198\u019e\7$\2\2\u0199\u019d"+
		"\n\5\2\2\u019a\u019b\7^\2\2\u019b\u019d\13\2\2\2\u019c\u0199\3\2\2\2\u019c"+
		"\u019a\3\2\2\2\u019d\u01a0\3\2\2\2\u019e\u019c\3\2\2\2\u019e\u019f\3\2"+
		"\2\2\u019f\u01a1\3\2\2\2\u01a0\u019e\3\2\2\2\u01a1\u01a2\7$\2\2\u01a2"+
		"\u0080\3\2\2\2\u01a3\u01a7\t\6\2\2\u01a4\u01a6\t\7\2\2\u01a5\u01a4\3\2"+
		"\2\2\u01a6\u01a9\3\2\2\2\u01a7\u01a5\3\2\2\2\u01a7\u01a8\3\2\2\2\u01a8"+
		"\u01b1\3\2\2\2\u01a9\u01a7\3\2\2\2\u01aa\u01ac\7a\2\2\u01ab\u01ad\t\7"+
		"\2\2\u01ac\u01ab\3\2\2\2\u01ad\u01ae\3\2\2\2\u01ae\u01ac\3\2\2\2\u01ae"+
		"\u01af\3\2\2\2\u01af\u01b1\3\2\2\2\u01b0\u01a3\3\2\2\2\u01b0\u01aa\3\2"+
		"\2\2\u01b1\u0082\3\2\2\2\u01b2\u01b3\7\61\2\2\u01b3\u01b4\7\61\2\2\u01b4"+
		"\u01b8\3\2\2\2\u01b5\u01b7\n\b\2\2\u01b6\u01b5\3\2\2\2\u01b7\u01ba\3\2"+
		"\2\2\u01b8\u01b6\3\2\2\2\u01b8\u01b9\3\2\2\2\u01b9\u01bb\3\2\2\2\u01ba"+
		"\u01b8\3\2\2\2\u01bb\u01bc\bB\3\2\u01bc\u0084\3\2\2\2\u01bd\u01be\7\61"+
		"\2\2\u01be\u01bf\7,\2\2\u01bf\u01c3\3\2\2\2\u01c0\u01c2\13\2\2\2\u01c1"+
		"\u01c0\3\2\2\2\u01c2\u01c5\3\2\2\2\u01c3\u01c4\3\2\2\2\u01c3\u01c1\3\2"+
		"\2\2\u01c4\u01c6\3\2\2\2\u01c5\u01c3\3\2\2\2\u01c6\u01c7\7,\2\2\u01c7"+
		"\u01c8\7\61\2\2\u01c8\u01c9\3\2\2\2\u01c9\u01ca\bC\3\2\u01ca\u0086\3\2"+
		"\2\2\u01cb\u01cc\13\2\2\2\u01cc\u01cd\3\2\2\2\u01cd\u01ce\bD\3\2\u01ce"+
		"\u0088\3\2\2\2\17\2\u017a\u0183\u0188\u018e\u0194\u019c\u019e\u01a7\u01ae"+
		"\u01b0\u01b8\u01c3\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}