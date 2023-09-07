// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/OLC2_Proyecto1_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.9.2
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
		T__0=1, PARDER=2, PARIZQ=3, LLAVEIZQ=4, LLAVEDER=5, CORCHETEIZQ=6, CORCHETEDER=7, 
		DOSPT=8, PTCOMA=9, INTERROGACION=10, PUNTO=11, GUIONBAJO=12, DIR=13, MASIGUAL=14, 
		MENOSIGUAL=15, IGUAL=16, DIV=17, MOD=18, MENOS=19, MAS=20, POR=21, IGUALIGUAL=22, 
		DIFERENTE=23, MAYORIGUAL=24, MENORIGUAL=25, MAYOR=26, MENOR=27, AND=28, 
		OR=29, NOT=30, PRINT=31, VAR=32, LET=33, TRUE=34, FALSE=35, IF=36, ELSE=37, 
		SWITCH=38, CASE=39, DEFAULT=40, WHILE=41, FOR=42, IN=43, RANGO=44, GUARD=45, 
		CONTINUE=46, RETURN=47, BREAK=48, NIL=49, APPEND=50, REMOVELAST=51, REMOVE=52, 
		AT=53, ISEMPTY=54, COUNT=55, FUNC=56, REPEATING=57, STRUCT=58, MUTATING=59, 
		INOUT=60, ATOI=61, IOTA=62, INT=63, FLOAT=64, BOOL=65, CHARACTER=66, STRING=67, 
		ENBLANCO=68, ENTERO=69, DECIMAL=70, CARACTER=71, CADENA=72, ID=73, UL_C=74, 
		ML_C=75, ERROR=76;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "GUIONBAJO", "DIR", "MASIGUAL", 
			"MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", 
			"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", 
			"NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", 
			"CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", 
			"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", 
			"COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", "INOUT", "ATOI", 
			"IOTA", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "DIG", 
			"ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", 
			"'?'", "'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", 
			"'+'", "'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", 
			"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'", 
			"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'", 
			"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'", 
			"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'", 
			"'repeating'", "'struct'", "'mutating'", "'inout'", "'atoi'", "'iota'", 
			"'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", 
			"CORCHETEDER", "DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "GUIONBAJO", 
			"DIR", "MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", 
			"POR", "IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", 
			"IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", 
			"GUARD", "CONTINUE", "RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", 
			"REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", 
			"INOUT", "ATOI", "IOTA", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", 
			"ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", 
			"ML_C", "ERROR"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2N\u0218\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6"+
		"\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3"+
		"\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3"+
		"\24\3\25\3\25\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3"+
		"\32\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36\3\37\3"+
		"\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3$\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3"+
		"(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3"+
		",\3,\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3/\3/\3\60\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62"+
		"\3\62\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66"+
		"\3\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\38\39\39"+
		"\39\39\39\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3<\3<\3<"+
		"\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3@"+
		"\3@\3@\3@\3A\3A\3A\3A\3A\3A\3B\3B\3B\3B\3B\3C\3C\3C\3C\3C\3C\3C\3C\3C"+
		"\3C\3D\3D\3D\3D\3D\3D\3D\3E\6E\u01c2\nE\rE\16E\u01c3\3E\3E\3F\3F\3G\6"+
		"G\u01cb\nG\rG\16G\u01cc\3H\6H\u01d0\nH\rH\16H\u01d1\3H\3H\6H\u01d6\nH"+
		"\rH\16H\u01d7\3I\3I\3I\3I\5I\u01de\nI\3I\3I\3J\3J\3J\3J\7J\u01e6\nJ\f"+
		"J\16J\u01e9\13J\3J\3J\3K\3K\7K\u01ef\nK\fK\16K\u01f2\13K\3K\3K\6K\u01f6"+
		"\nK\rK\16K\u01f7\5K\u01fa\nK\3L\3L\3L\3L\7L\u0200\nL\fL\16L\u0203\13L"+
		"\3L\3L\3M\3M\3M\3M\7M\u020b\nM\fM\16M\u020e\13M\3M\3M\3M\3M\3M\3N\3N\3"+
		"N\3N\3\u020c\2O\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64"+
		"g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089"+
		"F\u008b\2\u008dG\u008fH\u0091I\u0093J\u0095K\u0097L\u0099M\u009bN\3\2"+
		"\t\5\2\13\f\17\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aa"+
		"c|\4\2\f\f\17\17\2\u0222\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2"+
		"\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3"+
		"\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2"+
		"\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2"+
		"Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3"+
		"\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2"+
		"\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2"+
		"w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2"+
		"\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008d"+
		"\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2"+
		"\2\2\u0097\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2\2\3\u009d\3\2\2\2\5\u009f"+
		"\3\2\2\2\7\u00a1\3\2\2\2\t\u00a3\3\2\2\2\13\u00a5\3\2\2\2\r\u00a7\3\2"+
		"\2\2\17\u00a9\3\2\2\2\21\u00ab\3\2\2\2\23\u00ad\3\2\2\2\25\u00af\3\2\2"+
		"\2\27\u00b1\3\2\2\2\31\u00b3\3\2\2\2\33\u00b5\3\2\2\2\35\u00b7\3\2\2\2"+
		"\37\u00ba\3\2\2\2!\u00bd\3\2\2\2#\u00bf\3\2\2\2%\u00c1\3\2\2\2\'\u00c3"+
		"\3\2\2\2)\u00c5\3\2\2\2+\u00c7\3\2\2\2-\u00c9\3\2\2\2/\u00cc\3\2\2\2\61"+
		"\u00cf\3\2\2\2\63\u00d2\3\2\2\2\65\u00d5\3\2\2\2\67\u00d7\3\2\2\29\u00d9"+
		"\3\2\2\2;\u00dc\3\2\2\2=\u00df\3\2\2\2?\u00e1\3\2\2\2A\u00e7\3\2\2\2C"+
		"\u00eb\3\2\2\2E\u00ef\3\2\2\2G\u00f4\3\2\2\2I\u00fa\3\2\2\2K\u00fd\3\2"+
		"\2\2M\u0102\3\2\2\2O\u0109\3\2\2\2Q\u010e\3\2\2\2S\u0116\3\2\2\2U\u011c"+
		"\3\2\2\2W\u0120\3\2\2\2Y\u0123\3\2\2\2[\u0127\3\2\2\2]\u012d\3\2\2\2_"+
		"\u0136\3\2\2\2a\u013d\3\2\2\2c\u0143\3\2\2\2e\u0147\3\2\2\2g\u014e\3\2"+
		"\2\2i\u0159\3\2\2\2k\u0160\3\2\2\2m\u0163\3\2\2\2o\u016b\3\2\2\2q\u0171"+
		"\3\2\2\2s\u0176\3\2\2\2u\u0180\3\2\2\2w\u0187\3\2\2\2y\u0190\3\2\2\2{"+
		"\u0196\3\2\2\2}\u019b\3\2\2\2\177\u01a0\3\2\2\2\u0081\u01a4\3\2\2\2\u0083"+
		"\u01aa\3\2\2\2\u0085\u01af\3\2\2\2\u0087\u01b9\3\2\2\2\u0089\u01c1\3\2"+
		"\2\2\u008b\u01c7\3\2\2\2\u008d\u01ca\3\2\2\2\u008f\u01cf\3\2\2\2\u0091"+
		"\u01d9\3\2\2\2\u0093\u01e1\3\2\2\2\u0095\u01f9\3\2\2\2\u0097\u01fb\3\2"+
		"\2\2\u0099\u0206\3\2\2\2\u009b\u0214\3\2\2\2\u009d\u009e\7.\2\2\u009e"+
		"\4\3\2\2\2\u009f\u00a0\7+\2\2\u00a0\6\3\2\2\2\u00a1\u00a2\7*\2\2\u00a2"+
		"\b\3\2\2\2\u00a3\u00a4\7}\2\2\u00a4\n\3\2\2\2\u00a5\u00a6\7\177\2\2\u00a6"+
		"\f\3\2\2\2\u00a7\u00a8\7]\2\2\u00a8\16\3\2\2\2\u00a9\u00aa\7_\2\2\u00aa"+
		"\20\3\2\2\2\u00ab\u00ac\7<\2\2\u00ac\22\3\2\2\2\u00ad\u00ae\7=\2\2\u00ae"+
		"\24\3\2\2\2\u00af\u00b0\7A\2\2\u00b0\26\3\2\2\2\u00b1\u00b2\7\60\2\2\u00b2"+
		"\30\3\2\2\2\u00b3\u00b4\7a\2\2\u00b4\32\3\2\2\2\u00b5\u00b6\7(\2\2\u00b6"+
		"\34\3\2\2\2\u00b7\u00b8\7-\2\2\u00b8\u00b9\7?\2\2\u00b9\36\3\2\2\2\u00ba"+
		"\u00bb\7/\2\2\u00bb\u00bc\7?\2\2\u00bc \3\2\2\2\u00bd\u00be\7?\2\2\u00be"+
		"\"\3\2\2\2\u00bf\u00c0\7\61\2\2\u00c0$\3\2\2\2\u00c1\u00c2\7\'\2\2\u00c2"+
		"&\3\2\2\2\u00c3\u00c4\7/\2\2\u00c4(\3\2\2\2\u00c5\u00c6\7-\2\2\u00c6*"+
		"\3\2\2\2\u00c7\u00c8\7,\2\2\u00c8,\3\2\2\2\u00c9\u00ca\7?\2\2\u00ca\u00cb"+
		"\7?\2\2\u00cb.\3\2\2\2\u00cc\u00cd\7#\2\2\u00cd\u00ce\7?\2\2\u00ce\60"+
		"\3\2\2\2\u00cf\u00d0\7@\2\2\u00d0\u00d1\7?\2\2\u00d1\62\3\2\2\2\u00d2"+
		"\u00d3\7>\2\2\u00d3\u00d4\7?\2\2\u00d4\64\3\2\2\2\u00d5\u00d6\7@\2\2\u00d6"+
		"\66\3\2\2\2\u00d7\u00d8\7>\2\2\u00d88\3\2\2\2\u00d9\u00da\7(\2\2\u00da"+
		"\u00db\7(\2\2\u00db:\3\2\2\2\u00dc\u00dd\7~\2\2\u00dd\u00de\7~\2\2\u00de"+
		"<\3\2\2\2\u00df\u00e0\7#\2\2\u00e0>\3\2\2\2\u00e1\u00e2\7r\2\2\u00e2\u00e3"+
		"\7t\2\2\u00e3\u00e4\7k\2\2\u00e4\u00e5\7p\2\2\u00e5\u00e6\7v\2\2\u00e6"+
		"@\3\2\2\2\u00e7\u00e8\7x\2\2\u00e8\u00e9\7c\2\2\u00e9\u00ea\7t\2\2\u00ea"+
		"B\3\2\2\2\u00eb\u00ec\7n\2\2\u00ec\u00ed\7g\2\2\u00ed\u00ee\7v\2\2\u00ee"+
		"D\3\2\2\2\u00ef\u00f0\7v\2\2\u00f0\u00f1\7t\2\2\u00f1\u00f2\7w\2\2\u00f2"+
		"\u00f3\7g\2\2\u00f3F\3\2\2\2\u00f4\u00f5\7h\2\2\u00f5\u00f6\7c\2\2\u00f6"+
		"\u00f7\7n\2\2\u00f7\u00f8\7u\2\2\u00f8\u00f9\7g\2\2\u00f9H\3\2\2\2\u00fa"+
		"\u00fb\7k\2\2\u00fb\u00fc\7h\2\2\u00fcJ\3\2\2\2\u00fd\u00fe\7g\2\2\u00fe"+
		"\u00ff\7n\2\2\u00ff\u0100\7u\2\2\u0100\u0101\7g\2\2\u0101L\3\2\2\2\u0102"+
		"\u0103\7u\2\2\u0103\u0104\7y\2\2\u0104\u0105\7k\2\2\u0105\u0106\7v\2\2"+
		"\u0106\u0107\7e\2\2\u0107\u0108\7j\2\2\u0108N\3\2\2\2\u0109\u010a\7e\2"+
		"\2\u010a\u010b\7c\2\2\u010b\u010c\7u\2\2\u010c\u010d\7g\2\2\u010dP\3\2"+
		"\2\2\u010e\u010f\7f\2\2\u010f\u0110\7g\2\2\u0110\u0111\7h\2\2\u0111\u0112"+
		"\7c\2\2\u0112\u0113\7w\2\2\u0113\u0114\7n\2\2\u0114\u0115\7v\2\2\u0115"+
		"R\3\2\2\2\u0116\u0117\7y\2\2\u0117\u0118\7j\2\2\u0118\u0119\7k\2\2\u0119"+
		"\u011a\7n\2\2\u011a\u011b\7g\2\2\u011bT\3\2\2\2\u011c\u011d\7h\2\2\u011d"+
		"\u011e\7q\2\2\u011e\u011f\7t\2\2\u011fV\3\2\2\2\u0120\u0121\7k\2\2\u0121"+
		"\u0122\7p\2\2\u0122X\3\2\2\2\u0123\u0124\7\60\2\2\u0124\u0125\7\60\2\2"+
		"\u0125\u0126\7\60\2\2\u0126Z\3\2\2\2\u0127\u0128\7i\2\2\u0128\u0129\7"+
		"w\2\2\u0129\u012a\7c\2\2\u012a\u012b\7t\2\2\u012b\u012c\7f\2\2\u012c\\"+
		"\3\2\2\2\u012d\u012e\7e\2\2\u012e\u012f\7q\2\2\u012f\u0130\7p\2\2\u0130"+
		"\u0131\7v\2\2\u0131\u0132\7k\2\2\u0132\u0133\7p\2\2\u0133\u0134\7w\2\2"+
		"\u0134\u0135\7g\2\2\u0135^\3\2\2\2\u0136\u0137\7t\2\2\u0137\u0138\7g\2"+
		"\2\u0138\u0139\7v\2\2\u0139\u013a\7w\2\2\u013a\u013b\7t\2\2\u013b\u013c"+
		"\7p\2\2\u013c`\3\2\2\2\u013d\u013e\7d\2\2\u013e\u013f\7t\2\2\u013f\u0140"+
		"\7g\2\2\u0140\u0141\7c\2\2\u0141\u0142\7m\2\2\u0142b\3\2\2\2\u0143\u0144"+
		"\7p\2\2\u0144\u0145\7k\2\2\u0145\u0146\7n\2\2\u0146d\3\2\2\2\u0147\u0148"+
		"\7c\2\2\u0148\u0149\7r\2\2\u0149\u014a\7r\2\2\u014a\u014b\7g\2\2\u014b"+
		"\u014c\7p\2\2\u014c\u014d\7f\2\2\u014df\3\2\2\2\u014e\u014f\7t\2\2\u014f"+
		"\u0150\7g\2\2\u0150\u0151\7o\2\2\u0151\u0152\7q\2\2\u0152\u0153\7x\2\2"+
		"\u0153\u0154\7g\2\2\u0154\u0155\7N\2\2\u0155\u0156\7c\2\2\u0156\u0157"+
		"\7u\2\2\u0157\u0158\7v\2\2\u0158h\3\2\2\2\u0159\u015a\7t\2\2\u015a\u015b"+
		"\7g\2\2\u015b\u015c\7o\2\2\u015c\u015d\7q\2\2\u015d\u015e\7x\2\2\u015e"+
		"\u015f\7g\2\2\u015fj\3\2\2\2\u0160\u0161\7c\2\2\u0161\u0162\7v\2\2\u0162"+
		"l\3\2\2\2\u0163\u0164\7K\2\2\u0164\u0165\7u\2\2\u0165\u0166\7G\2\2\u0166"+
		"\u0167\7o\2\2\u0167\u0168\7r\2\2\u0168\u0169\7v\2\2\u0169\u016a\7{\2\2"+
		"\u016an\3\2\2\2\u016b\u016c\7e\2\2\u016c\u016d\7q\2\2\u016d\u016e\7w\2"+
		"\2\u016e\u016f\7p\2\2\u016f\u0170\7v\2\2\u0170p\3\2\2\2\u0171\u0172\7"+
		"h\2\2\u0172\u0173\7w\2\2\u0173\u0174\7p\2\2\u0174\u0175\7e\2\2\u0175r"+
		"\3\2\2\2\u0176\u0177\7t\2\2\u0177\u0178\7g\2\2\u0178\u0179\7r\2\2\u0179"+
		"\u017a\7g\2\2\u017a\u017b\7c\2\2\u017b\u017c\7v\2\2\u017c\u017d\7k\2\2"+
		"\u017d\u017e\7p\2\2\u017e\u017f\7i\2\2\u017ft\3\2\2\2\u0180\u0181\7u\2"+
		"\2\u0181\u0182\7v\2\2\u0182\u0183\7t\2\2\u0183\u0184\7w\2\2\u0184\u0185"+
		"\7e\2\2\u0185\u0186\7v\2\2\u0186v\3\2\2\2\u0187\u0188\7o\2\2\u0188\u0189"+
		"\7w\2\2\u0189\u018a\7v\2\2\u018a\u018b\7c\2\2\u018b\u018c\7v\2\2\u018c"+
		"\u018d\7k\2\2\u018d\u018e\7p\2\2\u018e\u018f\7i\2\2\u018fx\3\2\2\2\u0190"+
		"\u0191\7k\2\2\u0191\u0192\7p\2\2\u0192\u0193\7q\2\2\u0193\u0194\7w\2\2"+
		"\u0194\u0195\7v\2\2\u0195z\3\2\2\2\u0196\u0197\7c\2\2\u0197\u0198\7v\2"+
		"\2\u0198\u0199\7q\2\2\u0199\u019a\7k\2\2\u019a|\3\2\2\2\u019b\u019c\7"+
		"k\2\2\u019c\u019d\7q\2\2\u019d\u019e\7v\2\2\u019e\u019f\7c\2\2\u019f~"+
		"\3\2\2\2\u01a0\u01a1\7K\2\2\u01a1\u01a2\7p\2\2\u01a2\u01a3\7v\2\2\u01a3"+
		"\u0080\3\2\2\2\u01a4\u01a5\7H\2\2\u01a5\u01a6\7n\2\2\u01a6\u01a7\7q\2"+
		"\2\u01a7\u01a8\7c\2\2\u01a8\u01a9\7v\2\2\u01a9\u0082\3\2\2\2\u01aa\u01ab"+
		"\7D\2\2\u01ab\u01ac\7q\2\2\u01ac\u01ad\7q\2\2\u01ad\u01ae\7n\2\2\u01ae"+
		"\u0084\3\2\2\2\u01af\u01b0\7E\2\2\u01b0\u01b1\7j\2\2\u01b1\u01b2\7c\2"+
		"\2\u01b2\u01b3\7t\2\2\u01b3\u01b4\7c\2\2\u01b4\u01b5\7e\2\2\u01b5\u01b6"+
		"\7v\2\2\u01b6\u01b7\7g\2\2\u01b7\u01b8\7t\2\2\u01b8\u0086\3\2\2\2\u01b9"+
		"\u01ba\7U\2\2\u01ba\u01bb\7v\2\2\u01bb\u01bc\7t\2\2\u01bc\u01bd\7k\2\2"+
		"\u01bd\u01be\7p\2\2\u01be\u01bf\7i\2\2\u01bf\u0088\3\2\2\2\u01c0\u01c2"+
		"\t\2\2\2\u01c1\u01c0\3\2\2\2\u01c2\u01c3\3\2\2\2\u01c3\u01c1\3\2\2\2\u01c3"+
		"\u01c4\3\2\2\2\u01c4\u01c5\3\2\2\2\u01c5\u01c6\bE\2\2\u01c6\u008a\3\2"+
		"\2\2\u01c7\u01c8\t\3\2\2\u01c8\u008c\3\2\2\2\u01c9\u01cb\5\u008bF\2\u01ca"+
		"\u01c9\3\2\2\2\u01cb\u01cc\3\2\2\2\u01cc\u01ca\3\2\2\2\u01cc\u01cd\3\2"+
		"\2\2\u01cd\u008e\3\2\2\2\u01ce\u01d0\5\u008bF\2\u01cf\u01ce\3\2\2\2\u01d0"+
		"\u01d1\3\2\2\2\u01d1\u01cf\3\2\2\2\u01d1\u01d2\3\2\2\2\u01d2\u01d3\3\2"+
		"\2\2\u01d3\u01d5\7\60\2\2\u01d4\u01d6\5\u008bF\2\u01d5\u01d4\3\2\2\2\u01d6"+
		"\u01d7\3\2\2\2\u01d7\u01d5\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8\u0090\3\2"+
		"\2\2\u01d9\u01dd\7$\2\2\u01da\u01de\n\4\2\2\u01db\u01dc\7^\2\2\u01dc\u01de"+
		"\13\2\2\2\u01dd\u01da\3\2\2\2\u01dd\u01db\3\2\2\2\u01de\u01df\3\2\2\2"+
		"\u01df\u01e0\7$\2\2\u01e0\u0092\3\2\2\2\u01e1\u01e7\7$\2\2\u01e2\u01e6"+
		"\n\5\2\2\u01e3\u01e4\7^\2\2\u01e4\u01e6\13\2\2\2\u01e5\u01e2\3\2\2\2\u01e5"+
		"\u01e3\3\2\2\2\u01e6\u01e9\3\2\2\2\u01e7\u01e5\3\2\2\2\u01e7\u01e8\3\2"+
		"\2\2\u01e8\u01ea\3\2\2\2\u01e9\u01e7\3\2\2\2\u01ea\u01eb\7$\2\2\u01eb"+
		"\u0094\3\2\2\2\u01ec\u01f0\t\6\2\2\u01ed\u01ef\t\7\2\2\u01ee\u01ed\3\2"+
		"\2\2\u01ef\u01f2\3\2\2\2\u01f0\u01ee\3\2\2\2\u01f0\u01f1\3\2\2\2\u01f1"+
		"\u01fa\3\2\2\2\u01f2\u01f0\3\2\2\2\u01f3\u01f5\7a\2\2\u01f4\u01f6\t\7"+
		"\2\2\u01f5\u01f4\3\2\2\2\u01f6\u01f7\3\2\2\2\u01f7\u01f5\3\2\2\2\u01f7"+
		"\u01f8\3\2\2\2\u01f8\u01fa\3\2\2\2\u01f9\u01ec\3\2\2\2\u01f9\u01f3\3\2"+
		"\2\2\u01fa\u0096\3\2\2\2\u01fb\u01fc\7\61\2\2\u01fc\u01fd\7\61\2\2\u01fd"+
		"\u0201\3\2\2\2\u01fe\u0200\n\b\2\2\u01ff\u01fe\3\2\2\2\u0200\u0203\3\2"+
		"\2\2\u0201\u01ff\3\2\2\2\u0201\u0202\3\2\2\2\u0202\u0204\3\2\2\2\u0203"+
		"\u0201\3\2\2\2\u0204\u0205\bL\3\2\u0205\u0098\3\2\2\2\u0206\u0207\7\61"+
		"\2\2\u0207\u0208\7,\2\2\u0208\u020c\3\2\2\2\u0209\u020b\13\2\2\2\u020a"+
		"\u0209\3\2\2\2\u020b\u020e\3\2\2\2\u020c\u020d\3\2\2\2\u020c\u020a\3\2"+
		"\2\2\u020d\u020f\3\2\2\2\u020e\u020c\3\2\2\2\u020f\u0210\7,\2\2\u0210"+
		"\u0211\7\61\2\2\u0211\u0212\3\2\2\2\u0212\u0213\bM\3\2\u0213\u009a\3\2"+
		"\2\2\u0214\u0215\13\2\2\2\u0215\u0216\3\2\2\2\u0216\u0217\bN\3\2\u0217"+
		"\u009c\3\2\2\2\17\2\u01c3\u01cc\u01d1\u01d7\u01dd\u01e5\u01e7\u01f0\u01f7"+
		"\u01f9\u0201\u020c\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}