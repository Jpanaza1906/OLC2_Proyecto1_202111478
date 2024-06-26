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
		INOUT=60, ATOI=61, IOTA=62, SELF=63, INT=64, FLOAT=65, BOOL=66, CHARACTER=67, 
		STRING=68, ENBLANCO=69, ENTERO=70, DECIMAL=71, CARACTER=72, CADENA=73, 
		ID=74, UL_C=75, ML_C=76, ERROR=77;
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
			"IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", 
			"DIG", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", 
			"ERROR"
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
			"'self'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
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
			"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHARACTER", 
			"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", 
			"UL_C", "ML_C", "ERROR"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2O\u021f\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3"+
		"\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3"+
		"\31\3\32\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36\3"+
		"\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#"+
		"\3$\3$\3$\3$\3$\3$\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'"+
		"\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+"+
		"\3,\3,\3,\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3/\3/\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62"+
		"\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\66"+
		"\3\66\3\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\38\3"+
		"9\39\39\39\39\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3<\3"+
		"<\3<\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3?\3?\3?\3?\3"+
		"?\3@\3@\3@\3@\3@\3A\3A\3A\3A\3B\3B\3B\3B\3B\3B\3C\3C\3C\3C\3C\3D\3D\3"+
		"D\3D\3D\3D\3D\3D\3D\3D\3E\3E\3E\3E\3E\3E\3E\3F\6F\u01c9\nF\rF\16F\u01ca"+
		"\3F\3F\3G\3G\3H\6H\u01d2\nH\rH\16H\u01d3\3I\6I\u01d7\nI\rI\16I\u01d8\3"+
		"I\3I\6I\u01dd\nI\rI\16I\u01de\3J\3J\3J\3J\5J\u01e5\nJ\3J\3J\3K\3K\3K\3"+
		"K\7K\u01ed\nK\fK\16K\u01f0\13K\3K\3K\3L\3L\7L\u01f6\nL\fL\16L\u01f9\13"+
		"L\3L\3L\6L\u01fd\nL\rL\16L\u01fe\5L\u0201\nL\3M\3M\3M\3M\7M\u0207\nM\f"+
		"M\16M\u020a\13M\3M\3M\3N\3N\3N\3N\7N\u0212\nN\fN\16N\u0215\13N\3N\3N\3"+
		"N\3N\3N\3O\3O\3O\3O\3\u0213\2P\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087"+
		"E\u0089F\u008bG\u008d\2\u008fH\u0091I\u0093J\u0095K\u0097L\u0099M\u009b"+
		"N\u009dO\3\2\t\5\2\13\f\17\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6"+
		"\2\62;C\\aac|\4\2\f\f\17\17\2\u0229\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2"+
		"\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23"+
		"\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2"+
		"\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2"+
		"\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3"+
		"\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2"+
		"\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2"+
		"\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2["+
		"\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2"+
		"\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2"+
		"\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2"+
		"\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089"+
		"\3\2\2\2\2\u008b\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2"+
		"\2\2\u0095\3\2\2\2\2\u0097\3\2\2\2\2\u0099\3\2\2\2\2\u009b\3\2\2\2\2\u009d"+
		"\3\2\2\2\3\u009f\3\2\2\2\5\u00a1\3\2\2\2\7\u00a3\3\2\2\2\t\u00a5\3\2\2"+
		"\2\13\u00a7\3\2\2\2\r\u00a9\3\2\2\2\17\u00ab\3\2\2\2\21\u00ad\3\2\2\2"+
		"\23\u00af\3\2\2\2\25\u00b1\3\2\2\2\27\u00b3\3\2\2\2\31\u00b5\3\2\2\2\33"+
		"\u00b7\3\2\2\2\35\u00b9\3\2\2\2\37\u00bc\3\2\2\2!\u00bf\3\2\2\2#\u00c1"+
		"\3\2\2\2%\u00c3\3\2\2\2\'\u00c5\3\2\2\2)\u00c7\3\2\2\2+\u00c9\3\2\2\2"+
		"-\u00cb\3\2\2\2/\u00ce\3\2\2\2\61\u00d1\3\2\2\2\63\u00d4\3\2\2\2\65\u00d7"+
		"\3\2\2\2\67\u00d9\3\2\2\29\u00db\3\2\2\2;\u00de\3\2\2\2=\u00e1\3\2\2\2"+
		"?\u00e3\3\2\2\2A\u00e9\3\2\2\2C\u00ed\3\2\2\2E\u00f1\3\2\2\2G\u00f6\3"+
		"\2\2\2I\u00fc\3\2\2\2K\u00ff\3\2\2\2M\u0104\3\2\2\2O\u010b\3\2\2\2Q\u0110"+
		"\3\2\2\2S\u0118\3\2\2\2U\u011e\3\2\2\2W\u0122\3\2\2\2Y\u0125\3\2\2\2["+
		"\u0129\3\2\2\2]\u012f\3\2\2\2_\u0138\3\2\2\2a\u013f\3\2\2\2c\u0145\3\2"+
		"\2\2e\u0149\3\2\2\2g\u0150\3\2\2\2i\u015b\3\2\2\2k\u0162\3\2\2\2m\u0165"+
		"\3\2\2\2o\u016d\3\2\2\2q\u0173\3\2\2\2s\u0178\3\2\2\2u\u0182\3\2\2\2w"+
		"\u0189\3\2\2\2y\u0192\3\2\2\2{\u0198\3\2\2\2}\u019d\3\2\2\2\177\u01a2"+
		"\3\2\2\2\u0081\u01a7\3\2\2\2\u0083\u01ab\3\2\2\2\u0085\u01b1\3\2\2\2\u0087"+
		"\u01b6\3\2\2\2\u0089\u01c0\3\2\2\2\u008b\u01c8\3\2\2\2\u008d\u01ce\3\2"+
		"\2\2\u008f\u01d1\3\2\2\2\u0091\u01d6\3\2\2\2\u0093\u01e0\3\2\2\2\u0095"+
		"\u01e8\3\2\2\2\u0097\u0200\3\2\2\2\u0099\u0202\3\2\2\2\u009b\u020d\3\2"+
		"\2\2\u009d\u021b\3\2\2\2\u009f\u00a0\7.\2\2\u00a0\4\3\2\2\2\u00a1\u00a2"+
		"\7+\2\2\u00a2\6\3\2\2\2\u00a3\u00a4\7*\2\2\u00a4\b\3\2\2\2\u00a5\u00a6"+
		"\7}\2\2\u00a6\n\3\2\2\2\u00a7\u00a8\7\177\2\2\u00a8\f\3\2\2\2\u00a9\u00aa"+
		"\7]\2\2\u00aa\16\3\2\2\2\u00ab\u00ac\7_\2\2\u00ac\20\3\2\2\2\u00ad\u00ae"+
		"\7<\2\2\u00ae\22\3\2\2\2\u00af\u00b0\7=\2\2\u00b0\24\3\2\2\2\u00b1\u00b2"+
		"\7A\2\2\u00b2\26\3\2\2\2\u00b3\u00b4\7\60\2\2\u00b4\30\3\2\2\2\u00b5\u00b6"+
		"\7a\2\2\u00b6\32\3\2\2\2\u00b7\u00b8\7(\2\2\u00b8\34\3\2\2\2\u00b9\u00ba"+
		"\7-\2\2\u00ba\u00bb\7?\2\2\u00bb\36\3\2\2\2\u00bc\u00bd\7/\2\2\u00bd\u00be"+
		"\7?\2\2\u00be \3\2\2\2\u00bf\u00c0\7?\2\2\u00c0\"\3\2\2\2\u00c1\u00c2"+
		"\7\61\2\2\u00c2$\3\2\2\2\u00c3\u00c4\7\'\2\2\u00c4&\3\2\2\2\u00c5\u00c6"+
		"\7/\2\2\u00c6(\3\2\2\2\u00c7\u00c8\7-\2\2\u00c8*\3\2\2\2\u00c9\u00ca\7"+
		",\2\2\u00ca,\3\2\2\2\u00cb\u00cc\7?\2\2\u00cc\u00cd\7?\2\2\u00cd.\3\2"+
		"\2\2\u00ce\u00cf\7#\2\2\u00cf\u00d0\7?\2\2\u00d0\60\3\2\2\2\u00d1\u00d2"+
		"\7@\2\2\u00d2\u00d3\7?\2\2\u00d3\62\3\2\2\2\u00d4\u00d5\7>\2\2\u00d5\u00d6"+
		"\7?\2\2\u00d6\64\3\2\2\2\u00d7\u00d8\7@\2\2\u00d8\66\3\2\2\2\u00d9\u00da"+
		"\7>\2\2\u00da8\3\2\2\2\u00db\u00dc\7(\2\2\u00dc\u00dd\7(\2\2\u00dd:\3"+
		"\2\2\2\u00de\u00df\7~\2\2\u00df\u00e0\7~\2\2\u00e0<\3\2\2\2\u00e1\u00e2"+
		"\7#\2\2\u00e2>\3\2\2\2\u00e3\u00e4\7r\2\2\u00e4\u00e5\7t\2\2\u00e5\u00e6"+
		"\7k\2\2\u00e6\u00e7\7p\2\2\u00e7\u00e8\7v\2\2\u00e8@\3\2\2\2\u00e9\u00ea"+
		"\7x\2\2\u00ea\u00eb\7c\2\2\u00eb\u00ec\7t\2\2\u00ecB\3\2\2\2\u00ed\u00ee"+
		"\7n\2\2\u00ee\u00ef\7g\2\2\u00ef\u00f0\7v\2\2\u00f0D\3\2\2\2\u00f1\u00f2"+
		"\7v\2\2\u00f2\u00f3\7t\2\2\u00f3\u00f4\7w\2\2\u00f4\u00f5\7g\2\2\u00f5"+
		"F\3\2\2\2\u00f6\u00f7\7h\2\2\u00f7\u00f8\7c\2\2\u00f8\u00f9\7n\2\2\u00f9"+
		"\u00fa\7u\2\2\u00fa\u00fb\7g\2\2\u00fbH\3\2\2\2\u00fc\u00fd\7k\2\2\u00fd"+
		"\u00fe\7h\2\2\u00feJ\3\2\2\2\u00ff\u0100\7g\2\2\u0100\u0101\7n\2\2\u0101"+
		"\u0102\7u\2\2\u0102\u0103\7g\2\2\u0103L\3\2\2\2\u0104\u0105\7u\2\2\u0105"+
		"\u0106\7y\2\2\u0106\u0107\7k\2\2\u0107\u0108\7v\2\2\u0108\u0109\7e\2\2"+
		"\u0109\u010a\7j\2\2\u010aN\3\2\2\2\u010b\u010c\7e\2\2\u010c\u010d\7c\2"+
		"\2\u010d\u010e\7u\2\2\u010e\u010f\7g\2\2\u010fP\3\2\2\2\u0110\u0111\7"+
		"f\2\2\u0111\u0112\7g\2\2\u0112\u0113\7h\2\2\u0113\u0114\7c\2\2\u0114\u0115"+
		"\7w\2\2\u0115\u0116\7n\2\2\u0116\u0117\7v\2\2\u0117R\3\2\2\2\u0118\u0119"+
		"\7y\2\2\u0119\u011a\7j\2\2\u011a\u011b\7k\2\2\u011b\u011c\7n\2\2\u011c"+
		"\u011d\7g\2\2\u011dT\3\2\2\2\u011e\u011f\7h\2\2\u011f\u0120\7q\2\2\u0120"+
		"\u0121\7t\2\2\u0121V\3\2\2\2\u0122\u0123\7k\2\2\u0123\u0124\7p\2\2\u0124"+
		"X\3\2\2\2\u0125\u0126\7\60\2\2\u0126\u0127\7\60\2\2\u0127\u0128\7\60\2"+
		"\2\u0128Z\3\2\2\2\u0129\u012a\7i\2\2\u012a\u012b\7w\2\2\u012b\u012c\7"+
		"c\2\2\u012c\u012d\7t\2\2\u012d\u012e\7f\2\2\u012e\\\3\2\2\2\u012f\u0130"+
		"\7e\2\2\u0130\u0131\7q\2\2\u0131\u0132\7p\2\2\u0132\u0133\7v\2\2\u0133"+
		"\u0134\7k\2\2\u0134\u0135\7p\2\2\u0135\u0136\7w\2\2\u0136\u0137\7g\2\2"+
		"\u0137^\3\2\2\2\u0138\u0139\7t\2\2\u0139\u013a\7g\2\2\u013a\u013b\7v\2"+
		"\2\u013b\u013c\7w\2\2\u013c\u013d\7t\2\2\u013d\u013e\7p\2\2\u013e`\3\2"+
		"\2\2\u013f\u0140\7d\2\2\u0140\u0141\7t\2\2\u0141\u0142\7g\2\2\u0142\u0143"+
		"\7c\2\2\u0143\u0144\7m\2\2\u0144b\3\2\2\2\u0145\u0146\7p\2\2\u0146\u0147"+
		"\7k\2\2\u0147\u0148\7n\2\2\u0148d\3\2\2\2\u0149\u014a\7c\2\2\u014a\u014b"+
		"\7r\2\2\u014b\u014c\7r\2\2\u014c\u014d\7g\2\2\u014d\u014e\7p\2\2\u014e"+
		"\u014f\7f\2\2\u014ff\3\2\2\2\u0150\u0151\7t\2\2\u0151\u0152\7g\2\2\u0152"+
		"\u0153\7o\2\2\u0153\u0154\7q\2\2\u0154\u0155\7x\2\2\u0155\u0156\7g\2\2"+
		"\u0156\u0157\7N\2\2\u0157\u0158\7c\2\2\u0158\u0159\7u\2\2\u0159\u015a"+
		"\7v\2\2\u015ah\3\2\2\2\u015b\u015c\7t\2\2\u015c\u015d\7g\2\2\u015d\u015e"+
		"\7o\2\2\u015e\u015f\7q\2\2\u015f\u0160\7x\2\2\u0160\u0161\7g\2\2\u0161"+
		"j\3\2\2\2\u0162\u0163\7c\2\2\u0163\u0164\7v\2\2\u0164l\3\2\2\2\u0165\u0166"+
		"\7K\2\2\u0166\u0167\7u\2\2\u0167\u0168\7G\2\2\u0168\u0169\7o\2\2\u0169"+
		"\u016a\7r\2\2\u016a\u016b\7v\2\2\u016b\u016c\7{\2\2\u016cn\3\2\2\2\u016d"+
		"\u016e\7e\2\2\u016e\u016f\7q\2\2\u016f\u0170\7w\2\2\u0170\u0171\7p\2\2"+
		"\u0171\u0172\7v\2\2\u0172p\3\2\2\2\u0173\u0174\7h\2\2\u0174\u0175\7w\2"+
		"\2\u0175\u0176\7p\2\2\u0176\u0177\7e\2\2\u0177r\3\2\2\2\u0178\u0179\7"+
		"t\2\2\u0179\u017a\7g\2\2\u017a\u017b\7r\2\2\u017b\u017c\7g\2\2\u017c\u017d"+
		"\7c\2\2\u017d\u017e\7v\2\2\u017e\u017f\7k\2\2\u017f\u0180\7p\2\2\u0180"+
		"\u0181\7i\2\2\u0181t\3\2\2\2\u0182\u0183\7u\2\2\u0183\u0184\7v\2\2\u0184"+
		"\u0185\7t\2\2\u0185\u0186\7w\2\2\u0186\u0187\7e\2\2\u0187\u0188\7v\2\2"+
		"\u0188v\3\2\2\2\u0189\u018a\7o\2\2\u018a\u018b\7w\2\2\u018b\u018c\7v\2"+
		"\2\u018c\u018d\7c\2\2\u018d\u018e\7v\2\2\u018e\u018f\7k\2\2\u018f\u0190"+
		"\7p\2\2\u0190\u0191\7i\2\2\u0191x\3\2\2\2\u0192\u0193\7k\2\2\u0193\u0194"+
		"\7p\2\2\u0194\u0195\7q\2\2\u0195\u0196\7w\2\2\u0196\u0197\7v\2\2\u0197"+
		"z\3\2\2\2\u0198\u0199\7c\2\2\u0199\u019a\7v\2\2\u019a\u019b\7q\2\2\u019b"+
		"\u019c\7k\2\2\u019c|\3\2\2\2\u019d\u019e\7k\2\2\u019e\u019f\7q\2\2\u019f"+
		"\u01a0\7v\2\2\u01a0\u01a1\7c\2\2\u01a1~\3\2\2\2\u01a2\u01a3\7u\2\2\u01a3"+
		"\u01a4\7g\2\2\u01a4\u01a5\7n\2\2\u01a5\u01a6\7h\2\2\u01a6\u0080\3\2\2"+
		"\2\u01a7\u01a8\7K\2\2\u01a8\u01a9\7p\2\2\u01a9\u01aa\7v\2\2\u01aa\u0082"+
		"\3\2\2\2\u01ab\u01ac\7H\2\2\u01ac\u01ad\7n\2\2\u01ad\u01ae\7q\2\2\u01ae"+
		"\u01af\7c\2\2\u01af\u01b0\7v\2\2\u01b0\u0084\3\2\2\2\u01b1\u01b2\7D\2"+
		"\2\u01b2\u01b3\7q\2\2\u01b3\u01b4\7q\2\2\u01b4\u01b5\7n\2\2\u01b5\u0086"+
		"\3\2\2\2\u01b6\u01b7\7E\2\2\u01b7\u01b8\7j\2\2\u01b8\u01b9\7c\2\2\u01b9"+
		"\u01ba\7t\2\2\u01ba\u01bb\7c\2\2\u01bb\u01bc\7e\2\2\u01bc\u01bd\7v\2\2"+
		"\u01bd\u01be\7g\2\2\u01be\u01bf\7t\2\2\u01bf\u0088\3\2\2\2\u01c0\u01c1"+
		"\7U\2\2\u01c1\u01c2\7v\2\2\u01c2\u01c3\7t\2\2\u01c3\u01c4\7k\2\2\u01c4"+
		"\u01c5\7p\2\2\u01c5\u01c6\7i\2\2\u01c6\u008a\3\2\2\2\u01c7\u01c9\t\2\2"+
		"\2\u01c8\u01c7\3\2\2\2\u01c9\u01ca\3\2\2\2\u01ca\u01c8\3\2\2\2\u01ca\u01cb"+
		"\3\2\2\2\u01cb\u01cc\3\2\2\2\u01cc\u01cd\bF\2\2\u01cd\u008c\3\2\2\2\u01ce"+
		"\u01cf\t\3\2\2\u01cf\u008e\3\2\2\2\u01d0\u01d2\5\u008dG\2\u01d1\u01d0"+
		"\3\2\2\2\u01d2\u01d3\3\2\2\2\u01d3\u01d1\3\2\2\2\u01d3\u01d4\3\2\2\2\u01d4"+
		"\u0090\3\2\2\2\u01d5\u01d7\5\u008dG\2\u01d6\u01d5\3\2\2\2\u01d7\u01d8"+
		"\3\2\2\2\u01d8\u01d6\3\2\2\2\u01d8\u01d9\3\2\2\2\u01d9\u01da\3\2\2\2\u01da"+
		"\u01dc\7\60\2\2\u01db\u01dd\5\u008dG\2\u01dc\u01db\3\2\2\2\u01dd\u01de"+
		"\3\2\2\2\u01de\u01dc\3\2\2\2\u01de\u01df\3\2\2\2\u01df\u0092\3\2\2\2\u01e0"+
		"\u01e4\7$\2\2\u01e1\u01e5\n\4\2\2\u01e2\u01e3\7^\2\2\u01e3\u01e5\13\2"+
		"\2\2\u01e4\u01e1\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e5\u01e6\3\2\2\2\u01e6"+
		"\u01e7\7$\2\2\u01e7\u0094\3\2\2\2\u01e8\u01ee\7$\2\2\u01e9\u01ed\n\5\2"+
		"\2\u01ea\u01eb\7^\2\2\u01eb\u01ed\13\2\2\2\u01ec\u01e9\3\2\2\2\u01ec\u01ea"+
		"\3\2\2\2\u01ed\u01f0\3\2\2\2\u01ee\u01ec\3\2\2\2\u01ee\u01ef\3\2\2\2\u01ef"+
		"\u01f1\3\2\2\2\u01f0\u01ee\3\2\2\2\u01f1\u01f2\7$\2\2\u01f2\u0096\3\2"+
		"\2\2\u01f3\u01f7\t\6\2\2\u01f4\u01f6\t\7\2\2\u01f5\u01f4\3\2\2\2\u01f6"+
		"\u01f9\3\2\2\2\u01f7\u01f5\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8\u0201\3\2"+
		"\2\2\u01f9\u01f7\3\2\2\2\u01fa\u01fc\7a\2\2\u01fb\u01fd\t\7\2\2\u01fc"+
		"\u01fb\3\2\2\2\u01fd\u01fe\3\2\2\2\u01fe\u01fc\3\2\2\2\u01fe\u01ff\3\2"+
		"\2\2\u01ff\u0201\3\2\2\2\u0200\u01f3\3\2\2\2\u0200\u01fa\3\2\2\2\u0201"+
		"\u0098\3\2\2\2\u0202\u0203\7\61\2\2\u0203\u0204\7\61\2\2\u0204\u0208\3"+
		"\2\2\2\u0205\u0207\n\b\2\2\u0206\u0205\3\2\2\2\u0207\u020a\3\2\2\2\u0208"+
		"\u0206\3\2\2\2\u0208\u0209\3\2\2\2\u0209\u020b\3\2\2\2\u020a\u0208\3\2"+
		"\2\2\u020b\u020c\bM\3\2\u020c\u009a\3\2\2\2\u020d\u020e\7\61\2\2\u020e"+
		"\u020f\7,\2\2\u020f\u0213\3\2\2\2\u0210\u0212\13\2\2\2\u0211\u0210\3\2"+
		"\2\2\u0212\u0215\3\2\2\2\u0213\u0214\3\2\2\2\u0213\u0211\3\2\2\2\u0214"+
		"\u0216\3\2\2\2\u0215\u0213\3\2\2\2\u0216\u0217\7,\2\2\u0217\u0218\7\61"+
		"\2\2\u0218\u0219\3\2\2\2\u0219\u021a\bN\3\2\u021a\u009c\3\2\2\2\u021b"+
		"\u021c\13\2\2\2\u021c\u021d\3\2\2\2\u021d\u021e\bO\3\2\u021e\u009e\3\2"+
		"\2\2\17\2\u01ca\u01d3\u01d8\u01de\u01e4\u01ec\u01ee\u01f7\u01fe\u0200"+
		"\u0208\u0213\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}