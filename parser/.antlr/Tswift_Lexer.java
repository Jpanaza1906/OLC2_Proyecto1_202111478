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
		DOSPT=7, PTCOMA=8, INTERROGACION=9, PUNTO=10, GUIONBAJO=11, DIR=12, MASIGUAL=13, 
		MENOSIGUAL=14, IGUAL=15, DIV=16, MOD=17, MENOS=18, MAS=19, POR=20, IGUALIGUAL=21, 
		DIFERENTE=22, MAYORIGUAL=23, MENORIGUAL=24, MAYOR=25, MENOR=26, AND=27, 
		OR=28, NOT=29, PRINT=30, VAR=31, LET=32, TRUE=33, FALSE=34, IF=35, ELSE=36, 
		SWITCH=37, CASE=38, DEFAULT=39, WHILE=40, FOR=41, IN=42, RANGO=43, GUARD=44, 
		CONTINUE=45, RETURN=46, BREAK=47, NIL=48, APPEND=49, REMOVELAST=50, REMOVE=51, 
		AT=52, ISEMPTY=53, COUNT=54, FUNC=55, REPEATING=56, STRUCT=57, MUTATING=58, 
		INOUT=59, INT=60, FLOAT=61, BOOL=62, CHARACTER=63, STRING=64, ENBLANCO=65, 
		ENTERO=66, DECIMAL=67, CARACTER=68, CADENA=69, ID=70, UL_C=71, ML_C=72, 
		ERROR=73;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "GUIONBAJO", "DIR", "MASIGUAL", 
			"MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", 
			"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", 
			"NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", 
			"CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", 
			"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", 
			"COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", "INOUT", "INT", "FLOAT", 
			"BOOL", "CHARACTER", "STRING", "ENBLANCO", "DIG", "ENTERO", "DECIMAL", 
			"CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", 
			"'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'", 
			"'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", 
			"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'", 
			"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'", 
			"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'", 
			"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'", 
			"'repeating'", "'struct'", "'mutating'", "'inout'", "'Int'", "'Float'", 
			"'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "GUIONBAJO", "DIR", "MASIGUAL", 
			"MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", 
			"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", 
			"NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", 
			"CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", 
			"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", 
			"COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", "INOUT", "INT", "FLOAT", 
			"BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", 
			"CADENA", "ID", "UL_C", "ML_C", "ERROR"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2K\u0206\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3"+
		"\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33"+
		"\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3"+
		"$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3("+
		"\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3+\3+\3+\3,\3,\3,\3,"+
		"\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62"+
		"\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66"+
		"\3\66\3\66\3\66\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\39\39\39"+
		"\39\39\39\39\39\39\39\3:\3:\3:\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3;\3;"+
		"\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3@\3@"+
		"\3@\3@\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3A\3A\3A\3B\6B\u01b0\nB\rB\16B\u01b1"+
		"\3B\3B\3C\3C\3D\6D\u01b9\nD\rD\16D\u01ba\3E\6E\u01be\nE\rE\16E\u01bf\3"+
		"E\3E\6E\u01c4\nE\rE\16E\u01c5\3F\3F\3F\3F\5F\u01cc\nF\3F\3F\3G\3G\3G\3"+
		"G\7G\u01d4\nG\fG\16G\u01d7\13G\3G\3G\3H\3H\7H\u01dd\nH\fH\16H\u01e0\13"+
		"H\3H\3H\6H\u01e4\nH\rH\16H\u01e5\5H\u01e8\nH\3I\3I\3I\3I\7I\u01ee\nI\f"+
		"I\16I\u01f1\13I\3I\3I\3J\3J\3J\3J\7J\u01f9\nJ\fJ\16J\u01fc\13J\3J\3J\3"+
		"J\3J\3J\3K\3K\3K\3K\3\u01fa\2L\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085\2"+
		"\u0087D\u0089E\u008bF\u008dG\u008fH\u0091I\u0093J\u0095K\3\2\t\5\2\13"+
		"\f\17\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f"+
		"\17\17\2\u0210\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3"+
		"\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2"+
		"\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2"+
		"_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3"+
		"\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2"+
		"\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083"+
		"\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2"+
		"\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\3\u0097"+
		"\3\2\2\2\5\u0099\3\2\2\2\7\u009b\3\2\2\2\t\u009d\3\2\2\2\13\u009f\3\2"+
		"\2\2\r\u00a1\3\2\2\2\17\u00a3\3\2\2\2\21\u00a5\3\2\2\2\23\u00a7\3\2\2"+
		"\2\25\u00a9\3\2\2\2\27\u00ab\3\2\2\2\31\u00ad\3\2\2\2\33\u00af\3\2\2\2"+
		"\35\u00b2\3\2\2\2\37\u00b5\3\2\2\2!\u00b7\3\2\2\2#\u00b9\3\2\2\2%\u00bb"+
		"\3\2\2\2\'\u00bd\3\2\2\2)\u00bf\3\2\2\2+\u00c1\3\2\2\2-\u00c4\3\2\2\2"+
		"/\u00c7\3\2\2\2\61\u00ca\3\2\2\2\63\u00cd\3\2\2\2\65\u00cf\3\2\2\2\67"+
		"\u00d1\3\2\2\29\u00d4\3\2\2\2;\u00d7\3\2\2\2=\u00d9\3\2\2\2?\u00df\3\2"+
		"\2\2A\u00e3\3\2\2\2C\u00e7\3\2\2\2E\u00ec\3\2\2\2G\u00f2\3\2\2\2I\u00f5"+
		"\3\2\2\2K\u00fa\3\2\2\2M\u0101\3\2\2\2O\u0106\3\2\2\2Q\u010e\3\2\2\2S"+
		"\u0114\3\2\2\2U\u0118\3\2\2\2W\u011b\3\2\2\2Y\u011f\3\2\2\2[\u0125\3\2"+
		"\2\2]\u012e\3\2\2\2_\u0135\3\2\2\2a\u013b\3\2\2\2c\u013f\3\2\2\2e\u0146"+
		"\3\2\2\2g\u0151\3\2\2\2i\u0158\3\2\2\2k\u015b\3\2\2\2m\u0163\3\2\2\2o"+
		"\u0169\3\2\2\2q\u016e\3\2\2\2s\u0178\3\2\2\2u\u017f\3\2\2\2w\u0188\3\2"+
		"\2\2y\u018e\3\2\2\2{\u0192\3\2\2\2}\u0198\3\2\2\2\177\u019d\3\2\2\2\u0081"+
		"\u01a7\3\2\2\2\u0083\u01af\3\2\2\2\u0085\u01b5\3\2\2\2\u0087\u01b8\3\2"+
		"\2\2\u0089\u01bd\3\2\2\2\u008b\u01c7\3\2\2\2\u008d\u01cf\3\2\2\2\u008f"+
		"\u01e7\3\2\2\2\u0091\u01e9\3\2\2\2\u0093\u01f4\3\2\2\2\u0095\u0202\3\2"+
		"\2\2\u0097\u0098\7+\2\2\u0098\4\3\2\2\2\u0099\u009a\7*\2\2\u009a\6\3\2"+
		"\2\2\u009b\u009c\7}\2\2\u009c\b\3\2\2\2\u009d\u009e\7\177\2\2\u009e\n"+
		"\3\2\2\2\u009f\u00a0\7]\2\2\u00a0\f\3\2\2\2\u00a1\u00a2\7_\2\2\u00a2\16"+
		"\3\2\2\2\u00a3\u00a4\7<\2\2\u00a4\20\3\2\2\2\u00a5\u00a6\7=\2\2\u00a6"+
		"\22\3\2\2\2\u00a7\u00a8\7A\2\2\u00a8\24\3\2\2\2\u00a9\u00aa\7\60\2\2\u00aa"+
		"\26\3\2\2\2\u00ab\u00ac\7a\2\2\u00ac\30\3\2\2\2\u00ad\u00ae\7(\2\2\u00ae"+
		"\32\3\2\2\2\u00af\u00b0\7-\2\2\u00b0\u00b1\7?\2\2\u00b1\34\3\2\2\2\u00b2"+
		"\u00b3\7/\2\2\u00b3\u00b4\7?\2\2\u00b4\36\3\2\2\2\u00b5\u00b6\7?\2\2\u00b6"+
		" \3\2\2\2\u00b7\u00b8\7\61\2\2\u00b8\"\3\2\2\2\u00b9\u00ba\7\'\2\2\u00ba"+
		"$\3\2\2\2\u00bb\u00bc\7/\2\2\u00bc&\3\2\2\2\u00bd\u00be\7-\2\2\u00be("+
		"\3\2\2\2\u00bf\u00c0\7,\2\2\u00c0*\3\2\2\2\u00c1\u00c2\7?\2\2\u00c2\u00c3"+
		"\7?\2\2\u00c3,\3\2\2\2\u00c4\u00c5\7#\2\2\u00c5\u00c6\7?\2\2\u00c6.\3"+
		"\2\2\2\u00c7\u00c8\7@\2\2\u00c8\u00c9\7?\2\2\u00c9\60\3\2\2\2\u00ca\u00cb"+
		"\7>\2\2\u00cb\u00cc\7?\2\2\u00cc\62\3\2\2\2\u00cd\u00ce\7@\2\2\u00ce\64"+
		"\3\2\2\2\u00cf\u00d0\7>\2\2\u00d0\66\3\2\2\2\u00d1\u00d2\7(\2\2\u00d2"+
		"\u00d3\7(\2\2\u00d38\3\2\2\2\u00d4\u00d5\7~\2\2\u00d5\u00d6\7~\2\2\u00d6"+
		":\3\2\2\2\u00d7\u00d8\7#\2\2\u00d8<\3\2\2\2\u00d9\u00da\7r\2\2\u00da\u00db"+
		"\7t\2\2\u00db\u00dc\7k\2\2\u00dc\u00dd\7p\2\2\u00dd\u00de\7v\2\2\u00de"+
		">\3\2\2\2\u00df\u00e0\7x\2\2\u00e0\u00e1\7c\2\2\u00e1\u00e2\7t\2\2\u00e2"+
		"@\3\2\2\2\u00e3\u00e4\7n\2\2\u00e4\u00e5\7g\2\2\u00e5\u00e6\7v\2\2\u00e6"+
		"B\3\2\2\2\u00e7\u00e8\7v\2\2\u00e8\u00e9\7t\2\2\u00e9\u00ea\7w\2\2\u00ea"+
		"\u00eb\7g\2\2\u00ebD\3\2\2\2\u00ec\u00ed\7h\2\2\u00ed\u00ee\7c\2\2\u00ee"+
		"\u00ef\7n\2\2\u00ef\u00f0\7u\2\2\u00f0\u00f1\7g\2\2\u00f1F\3\2\2\2\u00f2"+
		"\u00f3\7k\2\2\u00f3\u00f4\7h\2\2\u00f4H\3\2\2\2\u00f5\u00f6\7g\2\2\u00f6"+
		"\u00f7\7n\2\2\u00f7\u00f8\7u\2\2\u00f8\u00f9\7g\2\2\u00f9J\3\2\2\2\u00fa"+
		"\u00fb\7u\2\2\u00fb\u00fc\7y\2\2\u00fc\u00fd\7k\2\2\u00fd\u00fe\7v\2\2"+
		"\u00fe\u00ff\7e\2\2\u00ff\u0100\7j\2\2\u0100L\3\2\2\2\u0101\u0102\7e\2"+
		"\2\u0102\u0103\7c\2\2\u0103\u0104\7u\2\2\u0104\u0105\7g\2\2\u0105N\3\2"+
		"\2\2\u0106\u0107\7f\2\2\u0107\u0108\7g\2\2\u0108\u0109\7h\2\2\u0109\u010a"+
		"\7c\2\2\u010a\u010b\7w\2\2\u010b\u010c\7n\2\2\u010c\u010d\7v\2\2\u010d"+
		"P\3\2\2\2\u010e\u010f\7y\2\2\u010f\u0110\7j\2\2\u0110\u0111\7k\2\2\u0111"+
		"\u0112\7n\2\2\u0112\u0113\7g\2\2\u0113R\3\2\2\2\u0114\u0115\7h\2\2\u0115"+
		"\u0116\7q\2\2\u0116\u0117\7t\2\2\u0117T\3\2\2\2\u0118\u0119\7k\2\2\u0119"+
		"\u011a\7p\2\2\u011aV\3\2\2\2\u011b\u011c\7\60\2\2\u011c\u011d\7\60\2\2"+
		"\u011d\u011e\7\60\2\2\u011eX\3\2\2\2\u011f\u0120\7i\2\2\u0120\u0121\7"+
		"w\2\2\u0121\u0122\7c\2\2\u0122\u0123\7t\2\2\u0123\u0124\7f\2\2\u0124Z"+
		"\3\2\2\2\u0125\u0126\7e\2\2\u0126\u0127\7q\2\2\u0127\u0128\7p\2\2\u0128"+
		"\u0129\7v\2\2\u0129\u012a\7k\2\2\u012a\u012b\7p\2\2\u012b\u012c\7w\2\2"+
		"\u012c\u012d\7g\2\2\u012d\\\3\2\2\2\u012e\u012f\7t\2\2\u012f\u0130\7g"+
		"\2\2\u0130\u0131\7v\2\2\u0131\u0132\7w\2\2\u0132\u0133\7t\2\2\u0133\u0134"+
		"\7p\2\2\u0134^\3\2\2\2\u0135\u0136\7d\2\2\u0136\u0137\7t\2\2\u0137\u0138"+
		"\7g\2\2\u0138\u0139\7c\2\2\u0139\u013a\7m\2\2\u013a`\3\2\2\2\u013b\u013c"+
		"\7p\2\2\u013c\u013d\7k\2\2\u013d\u013e\7n\2\2\u013eb\3\2\2\2\u013f\u0140"+
		"\7c\2\2\u0140\u0141\7r\2\2\u0141\u0142\7r\2\2\u0142\u0143\7g\2\2\u0143"+
		"\u0144\7p\2\2\u0144\u0145\7f\2\2\u0145d\3\2\2\2\u0146\u0147\7t\2\2\u0147"+
		"\u0148\7g\2\2\u0148\u0149\7o\2\2\u0149\u014a\7q\2\2\u014a\u014b\7x\2\2"+
		"\u014b\u014c\7g\2\2\u014c\u014d\7N\2\2\u014d\u014e\7c\2\2\u014e\u014f"+
		"\7u\2\2\u014f\u0150\7v\2\2\u0150f\3\2\2\2\u0151\u0152\7t\2\2\u0152\u0153"+
		"\7g\2\2\u0153\u0154\7o\2\2\u0154\u0155\7q\2\2\u0155\u0156\7x\2\2\u0156"+
		"\u0157\7g\2\2\u0157h\3\2\2\2\u0158\u0159\7c\2\2\u0159\u015a\7v\2\2\u015a"+
		"j\3\2\2\2\u015b\u015c\7K\2\2\u015c\u015d\7u\2\2\u015d\u015e\7G\2\2\u015e"+
		"\u015f\7o\2\2\u015f\u0160\7r\2\2\u0160\u0161\7v\2\2\u0161\u0162\7{\2\2"+
		"\u0162l\3\2\2\2\u0163\u0164\7e\2\2\u0164\u0165\7q\2\2\u0165\u0166\7w\2"+
		"\2\u0166\u0167\7p\2\2\u0167\u0168\7v\2\2\u0168n\3\2\2\2\u0169\u016a\7"+
		"h\2\2\u016a\u016b\7w\2\2\u016b\u016c\7p\2\2\u016c\u016d\7e\2\2\u016dp"+
		"\3\2\2\2\u016e\u016f\7t\2\2\u016f\u0170\7g\2\2\u0170\u0171\7r\2\2\u0171"+
		"\u0172\7g\2\2\u0172\u0173\7c\2\2\u0173\u0174\7v\2\2\u0174\u0175\7k\2\2"+
		"\u0175\u0176\7p\2\2\u0176\u0177\7i\2\2\u0177r\3\2\2\2\u0178\u0179\7u\2"+
		"\2\u0179\u017a\7v\2\2\u017a\u017b\7t\2\2\u017b\u017c\7w\2\2\u017c\u017d"+
		"\7e\2\2\u017d\u017e\7v\2\2\u017et\3\2\2\2\u017f\u0180\7o\2\2\u0180\u0181"+
		"\7w\2\2\u0181\u0182\7v\2\2\u0182\u0183\7c\2\2\u0183\u0184\7v\2\2\u0184"+
		"\u0185\7k\2\2\u0185\u0186\7p\2\2\u0186\u0187\7i\2\2\u0187v\3\2\2\2\u0188"+
		"\u0189\7k\2\2\u0189\u018a\7p\2\2\u018a\u018b\7q\2\2\u018b\u018c\7w\2\2"+
		"\u018c\u018d\7v\2\2\u018dx\3\2\2\2\u018e\u018f\7K\2\2\u018f\u0190\7p\2"+
		"\2\u0190\u0191\7v\2\2\u0191z\3\2\2\2\u0192\u0193\7H\2\2\u0193\u0194\7"+
		"n\2\2\u0194\u0195\7q\2\2\u0195\u0196\7c\2\2\u0196\u0197\7v\2\2\u0197|"+
		"\3\2\2\2\u0198\u0199\7D\2\2\u0199\u019a\7q\2\2\u019a\u019b\7q\2\2\u019b"+
		"\u019c\7n\2\2\u019c~\3\2\2\2\u019d\u019e\7E\2\2\u019e\u019f\7j\2\2\u019f"+
		"\u01a0\7c\2\2\u01a0\u01a1\7t\2\2\u01a1\u01a2\7c\2\2\u01a2\u01a3\7e\2\2"+
		"\u01a3\u01a4\7v\2\2\u01a4\u01a5\7g\2\2\u01a5\u01a6\7t\2\2\u01a6\u0080"+
		"\3\2\2\2\u01a7\u01a8\7U\2\2\u01a8\u01a9\7v\2\2\u01a9\u01aa\7t\2\2\u01aa"+
		"\u01ab\7k\2\2\u01ab\u01ac\7p\2\2\u01ac\u01ad\7i\2\2\u01ad\u0082\3\2\2"+
		"\2\u01ae\u01b0\t\2\2\2\u01af\u01ae\3\2\2\2\u01b0\u01b1\3\2\2\2\u01b1\u01af"+
		"\3\2\2\2\u01b1\u01b2\3\2\2\2\u01b2\u01b3\3\2\2\2\u01b3\u01b4\bB\2\2\u01b4"+
		"\u0084\3\2\2\2\u01b5\u01b6\t\3\2\2\u01b6\u0086\3\2\2\2\u01b7\u01b9\5\u0085"+
		"C\2\u01b8\u01b7\3\2\2\2\u01b9\u01ba\3\2\2\2\u01ba\u01b8\3\2\2\2\u01ba"+
		"\u01bb\3\2\2\2\u01bb\u0088\3\2\2\2\u01bc\u01be\5\u0085C\2\u01bd\u01bc"+
		"\3\2\2\2\u01be\u01bf\3\2\2\2\u01bf\u01bd\3\2\2\2\u01bf\u01c0\3\2\2\2\u01c0"+
		"\u01c1\3\2\2\2\u01c1\u01c3\7\60\2\2\u01c2\u01c4\5\u0085C\2\u01c3\u01c2"+
		"\3\2\2\2\u01c4\u01c5\3\2\2\2\u01c5\u01c3\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6"+
		"\u008a\3\2\2\2\u01c7\u01cb\7$\2\2\u01c8\u01cc\n\4\2\2\u01c9\u01ca\7^\2"+
		"\2\u01ca\u01cc\13\2\2\2\u01cb\u01c8\3\2\2\2\u01cb\u01c9\3\2\2\2\u01cc"+
		"\u01cd\3\2\2\2\u01cd\u01ce\7$\2\2\u01ce\u008c\3\2\2\2\u01cf\u01d5\7$\2"+
		"\2\u01d0\u01d4\n\5\2\2\u01d1\u01d2\7^\2\2\u01d2\u01d4\13\2\2\2\u01d3\u01d0"+
		"\3\2\2\2\u01d3\u01d1\3\2\2\2\u01d4\u01d7\3\2\2\2\u01d5\u01d3\3\2\2\2\u01d5"+
		"\u01d6\3\2\2\2\u01d6\u01d8\3\2\2\2\u01d7\u01d5\3\2\2\2\u01d8\u01d9\7$"+
		"\2\2\u01d9\u008e\3\2\2\2\u01da\u01de\t\6\2\2\u01db\u01dd\t\7\2\2\u01dc"+
		"\u01db\3\2\2\2\u01dd\u01e0\3\2\2\2\u01de\u01dc\3\2\2\2\u01de\u01df\3\2"+
		"\2\2\u01df\u01e8\3\2\2\2\u01e0\u01de\3\2\2\2\u01e1\u01e3\7a\2\2\u01e2"+
		"\u01e4\t\7\2\2\u01e3\u01e2\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5\u01e3\3\2"+
		"\2\2\u01e5\u01e6\3\2\2\2\u01e6\u01e8\3\2\2\2\u01e7\u01da\3\2\2\2\u01e7"+
		"\u01e1\3\2\2\2\u01e8\u0090\3\2\2\2\u01e9\u01ea\7\61\2\2\u01ea\u01eb\7"+
		"\61\2\2\u01eb\u01ef\3\2\2\2\u01ec\u01ee\n\b\2\2\u01ed\u01ec\3\2\2\2\u01ee"+
		"\u01f1\3\2\2\2\u01ef\u01ed\3\2\2\2\u01ef\u01f0\3\2\2\2\u01f0\u01f2\3\2"+
		"\2\2\u01f1\u01ef\3\2\2\2\u01f2\u01f3\bI\3\2\u01f3\u0092\3\2\2\2\u01f4"+
		"\u01f5\7\61\2\2\u01f5\u01f6\7,\2\2\u01f6\u01fa\3\2\2\2\u01f7\u01f9\13"+
		"\2\2\2\u01f8\u01f7\3\2\2\2\u01f9\u01fc\3\2\2\2\u01fa\u01fb\3\2\2\2\u01fa"+
		"\u01f8\3\2\2\2\u01fb\u01fd\3\2\2\2\u01fc\u01fa\3\2\2\2\u01fd\u01fe\7,"+
		"\2\2\u01fe\u01ff\7\61\2\2\u01ff\u0200\3\2\2\2\u0200\u0201\bJ\3\2\u0201"+
		"\u0094\3\2\2\2\u0202\u0203\13\2\2\2\u0203\u0204\3\2\2\2\u0204\u0205\b"+
		"K\3\2\u0205\u0096\3\2\2\2\17\2\u01b1\u01ba\u01bf\u01c5\u01cb\u01d3\u01d5"+
		"\u01de\u01e5\u01e7\u01ef\u01fa\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}