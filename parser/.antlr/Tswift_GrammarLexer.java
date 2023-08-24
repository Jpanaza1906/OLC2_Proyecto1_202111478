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
		T__0=1, PARDER=2, PARIZQ=3, LLAVEIZQ=4, LLAVEDER=5, DOSPT=6, INTERROGACION=7, 
		MASIGUAL=8, MENOSIGUAL=9, IGUAL=10, DIV=11, MOD=12, MENOS=13, MAS=14, 
		POR=15, IGUALIGUAL=16, DIFERENTE=17, MAYORIGUAL=18, MENORIGUAL=19, MAYOR=20, 
		MENOR=21, AND=22, OR=23, NOT=24, PRINT=25, VAR=26, LET=27, TRUE=28, FALSE=29, 
		RETURN=30, IF=31, ELSE=32, SWITCH=33, CASE=34, BREAK=35, DEFAULT=36, WHILE=37, 
		FOR=38, IN=39, RANGO=40, GUARD=41, CONTINUE=42, NIL=43, INT=44, FLOAT=45, 
		BOOL=46, CHARACTER=47, STRING=48, ENBLANCO=49, ENTERO=50, DECIMAL=51, 
		BOOLEANO=52, CARACTER=53, CADENA=54, ID=55, UL_C=56, ML_C=57;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "INTERROGACION", 
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
			null, "','", "')'", "'('", "'{'", "'}'", "':'", "'?'", "'+='", "'-='", 
			"'='", "'/'", "'%'", "'-'", "'+'", "'*'", "'=='", "'!='", "'>='", "'<='", 
			"'>'", "'<'", "'&&'", "'||'", "'!'", "'print'", "'var'", "'let'", "'true'", 
			"'false'", "'return'", "'if'", "'else'", "'switch'", "'case'", "'break'", 
			"'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", "'continue'", 
			"'nil'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "INTERROGACION", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2;\u018d\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3\3"+
		"\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\26\3\26\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33"+
		"\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3!\3"+
		"!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3"+
		"$\3%\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3"+
		")\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3"+
		"-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3"+
		"\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62\6\62\u0137"+
		"\n\62\r\62\16\62\u0138\3\62\3\62\3\63\3\63\3\64\6\64\u0140\n\64\r\64\16"+
		"\64\u0141\3\65\6\65\u0145\n\65\r\65\16\65\u0146\3\65\3\65\6\65\u014b\n"+
		"\65\r\65\16\65\u014c\3\66\3\66\5\66\u0151\n\66\3\67\3\67\3\67\3\67\5\67"+
		"\u0157\n\67\3\67\3\67\38\38\38\38\78\u015f\n8\f8\168\u0162\138\38\38\3"+
		"9\39\79\u0168\n9\f9\169\u016b\139\39\39\69\u016f\n9\r9\169\u0170\59\u0173"+
		"\n9\3:\3:\3:\3:\7:\u0179\n:\f:\16:\u017c\13:\3:\3:\3;\3;\3;\3;\7;\u0184"+
		"\n;\f;\16;\u0187\13;\3;\3;\3;\3;\3;\3\u0185\2<\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\2g\64i\65k\66m\67o8q9s:u;\3\2\t\5\2\13"+
		"\f\17\17\"\"\3\2\62;\4\2))^^\4\2$$^^\5\2C\\aac|\6\2\62;C\\aac|\4\2\f\f"+
		"\17\17\2\u0198\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3"+
		"\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2"+
		"\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2"+
		"_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3"+
		"\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\3w\3\2\2\2\5y\3\2\2"+
		"\2\7{\3\2\2\2\t}\3\2\2\2\13\177\3\2\2\2\r\u0081\3\2\2\2\17\u0083\3\2\2"+
		"\2\21\u0085\3\2\2\2\23\u0088\3\2\2\2\25\u008b\3\2\2\2\27\u008d\3\2\2\2"+
		"\31\u008f\3\2\2\2\33\u0091\3\2\2\2\35\u0093\3\2\2\2\37\u0095\3\2\2\2!"+
		"\u0097\3\2\2\2#\u009a\3\2\2\2%\u009d\3\2\2\2\'\u00a0\3\2\2\2)\u00a3\3"+
		"\2\2\2+\u00a5\3\2\2\2-\u00a7\3\2\2\2/\u00aa\3\2\2\2\61\u00ad\3\2\2\2\63"+
		"\u00af\3\2\2\2\65\u00b5\3\2\2\2\67\u00b9\3\2\2\29\u00bd\3\2\2\2;\u00c2"+
		"\3\2\2\2=\u00c8\3\2\2\2?\u00cf\3\2\2\2A\u00d2\3\2\2\2C\u00d7\3\2\2\2E"+
		"\u00de\3\2\2\2G\u00e3\3\2\2\2I\u00e9\3\2\2\2K\u00f1\3\2\2\2M\u00f7\3\2"+
		"\2\2O\u00fb\3\2\2\2Q\u00fe\3\2\2\2S\u0102\3\2\2\2U\u0108\3\2\2\2W\u0111"+
		"\3\2\2\2Y\u0115\3\2\2\2[\u0119\3\2\2\2]\u011f\3\2\2\2_\u0124\3\2\2\2a"+
		"\u012e\3\2\2\2c\u0136\3\2\2\2e\u013c\3\2\2\2g\u013f\3\2\2\2i\u0144\3\2"+
		"\2\2k\u0150\3\2\2\2m\u0152\3\2\2\2o\u015a\3\2\2\2q\u0172\3\2\2\2s\u0174"+
		"\3\2\2\2u\u017f\3\2\2\2wx\7.\2\2x\4\3\2\2\2yz\7+\2\2z\6\3\2\2\2{|\7*\2"+
		"\2|\b\3\2\2\2}~\7}\2\2~\n\3\2\2\2\177\u0080\7\177\2\2\u0080\f\3\2\2\2"+
		"\u0081\u0082\7<\2\2\u0082\16\3\2\2\2\u0083\u0084\7A\2\2\u0084\20\3\2\2"+
		"\2\u0085\u0086\7-\2\2\u0086\u0087\7?\2\2\u0087\22\3\2\2\2\u0088\u0089"+
		"\7/\2\2\u0089\u008a\7?\2\2\u008a\24\3\2\2\2\u008b\u008c\7?\2\2\u008c\26"+
		"\3\2\2\2\u008d\u008e\7\61\2\2\u008e\30\3\2\2\2\u008f\u0090\7\'\2\2\u0090"+
		"\32\3\2\2\2\u0091\u0092\7/\2\2\u0092\34\3\2\2\2\u0093\u0094\7-\2\2\u0094"+
		"\36\3\2\2\2\u0095\u0096\7,\2\2\u0096 \3\2\2\2\u0097\u0098\7?\2\2\u0098"+
		"\u0099\7?\2\2\u0099\"\3\2\2\2\u009a\u009b\7#\2\2\u009b\u009c\7?\2\2\u009c"+
		"$\3\2\2\2\u009d\u009e\7@\2\2\u009e\u009f\7?\2\2\u009f&\3\2\2\2\u00a0\u00a1"+
		"\7>\2\2\u00a1\u00a2\7?\2\2\u00a2(\3\2\2\2\u00a3\u00a4\7@\2\2\u00a4*\3"+
		"\2\2\2\u00a5\u00a6\7>\2\2\u00a6,\3\2\2\2\u00a7\u00a8\7(\2\2\u00a8\u00a9"+
		"\7(\2\2\u00a9.\3\2\2\2\u00aa\u00ab\7~\2\2\u00ab\u00ac\7~\2\2\u00ac\60"+
		"\3\2\2\2\u00ad\u00ae\7#\2\2\u00ae\62\3\2\2\2\u00af\u00b0\7r\2\2\u00b0"+
		"\u00b1\7t\2\2\u00b1\u00b2\7k\2\2\u00b2\u00b3\7p\2\2\u00b3\u00b4\7v\2\2"+
		"\u00b4\64\3\2\2\2\u00b5\u00b6\7x\2\2\u00b6\u00b7\7c\2\2\u00b7\u00b8\7"+
		"t\2\2\u00b8\66\3\2\2\2\u00b9\u00ba\7n\2\2\u00ba\u00bb\7g\2\2\u00bb\u00bc"+
		"\7v\2\2\u00bc8\3\2\2\2\u00bd\u00be\7v\2\2\u00be\u00bf\7t\2\2\u00bf\u00c0"+
		"\7w\2\2\u00c0\u00c1\7g\2\2\u00c1:\3\2\2\2\u00c2\u00c3\7h\2\2\u00c3\u00c4"+
		"\7c\2\2\u00c4\u00c5\7n\2\2\u00c5\u00c6\7u\2\2\u00c6\u00c7\7g\2\2\u00c7"+
		"<\3\2\2\2\u00c8\u00c9\7t\2\2\u00c9\u00ca\7g\2\2\u00ca\u00cb\7v\2\2\u00cb"+
		"\u00cc\7w\2\2\u00cc\u00cd\7t\2\2\u00cd\u00ce\7p\2\2\u00ce>\3\2\2\2\u00cf"+
		"\u00d0\7k\2\2\u00d0\u00d1\7h\2\2\u00d1@\3\2\2\2\u00d2\u00d3\7g\2\2\u00d3"+
		"\u00d4\7n\2\2\u00d4\u00d5\7u\2\2\u00d5\u00d6\7g\2\2\u00d6B\3\2\2\2\u00d7"+
		"\u00d8\7u\2\2\u00d8\u00d9\7y\2\2\u00d9\u00da\7k\2\2\u00da\u00db\7v\2\2"+
		"\u00db\u00dc\7e\2\2\u00dc\u00dd\7j\2\2\u00ddD\3\2\2\2\u00de\u00df\7e\2"+
		"\2\u00df\u00e0\7c\2\2\u00e0\u00e1\7u\2\2\u00e1\u00e2\7g\2\2\u00e2F\3\2"+
		"\2\2\u00e3\u00e4\7d\2\2\u00e4\u00e5\7t\2\2\u00e5\u00e6\7g\2\2\u00e6\u00e7"+
		"\7c\2\2\u00e7\u00e8\7m\2\2\u00e8H\3\2\2\2\u00e9\u00ea\7f\2\2\u00ea\u00eb"+
		"\7g\2\2\u00eb\u00ec\7h\2\2\u00ec\u00ed\7c\2\2\u00ed\u00ee\7w\2\2\u00ee"+
		"\u00ef\7n\2\2\u00ef\u00f0\7v\2\2\u00f0J\3\2\2\2\u00f1\u00f2\7y\2\2\u00f2"+
		"\u00f3\7j\2\2\u00f3\u00f4\7k\2\2\u00f4\u00f5\7n\2\2\u00f5\u00f6\7g\2\2"+
		"\u00f6L\3\2\2\2\u00f7\u00f8\7h\2\2\u00f8\u00f9\7q\2\2\u00f9\u00fa\7t\2"+
		"\2\u00faN\3\2\2\2\u00fb\u00fc\7k\2\2\u00fc\u00fd\7p\2\2\u00fdP\3\2\2\2"+
		"\u00fe\u00ff\7\60\2\2\u00ff\u0100\7\60\2\2\u0100\u0101\7\60\2\2\u0101"+
		"R\3\2\2\2\u0102\u0103\7i\2\2\u0103\u0104\7w\2\2\u0104\u0105\7c\2\2\u0105"+
		"\u0106\7t\2\2\u0106\u0107\7f\2\2\u0107T\3\2\2\2\u0108\u0109\7e\2\2\u0109"+
		"\u010a\7q\2\2\u010a\u010b\7p\2\2\u010b\u010c\7v\2\2\u010c\u010d\7k\2\2"+
		"\u010d\u010e\7p\2\2\u010e\u010f\7w\2\2\u010f\u0110\7g\2\2\u0110V\3\2\2"+
		"\2\u0111\u0112\7p\2\2\u0112\u0113\7k\2\2\u0113\u0114\7n\2\2\u0114X\3\2"+
		"\2\2\u0115\u0116\7K\2\2\u0116\u0117\7p\2\2\u0117\u0118\7v\2\2\u0118Z\3"+
		"\2\2\2\u0119\u011a\7H\2\2\u011a\u011b\7n\2\2\u011b\u011c\7q\2\2\u011c"+
		"\u011d\7c\2\2\u011d\u011e\7v\2\2\u011e\\\3\2\2\2\u011f\u0120\7D\2\2\u0120"+
		"\u0121\7q\2\2\u0121\u0122\7q\2\2\u0122\u0123\7n\2\2\u0123^\3\2\2\2\u0124"+
		"\u0125\7E\2\2\u0125\u0126\7j\2\2\u0126\u0127\7c\2\2\u0127\u0128\7t\2\2"+
		"\u0128\u0129\7c\2\2\u0129\u012a\7e\2\2\u012a\u012b\7v\2\2\u012b\u012c"+
		"\7g\2\2\u012c\u012d\7t\2\2\u012d`\3\2\2\2\u012e\u012f\7U\2\2\u012f\u0130"+
		"\7v\2\2\u0130\u0131\7t\2\2\u0131\u0132\7k\2\2\u0132\u0133\7p\2\2\u0133"+
		"\u0134\7i\2\2\u0134b\3\2\2\2\u0135\u0137\t\2\2\2\u0136\u0135\3\2\2\2\u0137"+
		"\u0138\3\2\2\2\u0138\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013a\3\2"+
		"\2\2\u013a\u013b\b\62\2\2\u013bd\3\2\2\2\u013c\u013d\t\3\2\2\u013df\3"+
		"\2\2\2\u013e\u0140\5e\63\2\u013f\u013e\3\2\2\2\u0140\u0141\3\2\2\2\u0141"+
		"\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142h\3\2\2\2\u0143\u0145\5e\63\2"+
		"\u0144\u0143\3\2\2\2\u0145\u0146\3\2\2\2\u0146\u0144\3\2\2\2\u0146\u0147"+
		"\3\2\2\2\u0147\u0148\3\2\2\2\u0148\u014a\7\60\2\2\u0149\u014b\5e\63\2"+
		"\u014a\u0149\3\2\2\2\u014b\u014c\3\2\2\2\u014c\u014a\3\2\2\2\u014c\u014d"+
		"\3\2\2\2\u014dj\3\2\2\2\u014e\u0151\59\35\2\u014f\u0151\5;\36\2\u0150"+
		"\u014e\3\2\2\2\u0150\u014f\3\2\2\2\u0151l\3\2\2\2\u0152\u0156\7)\2\2\u0153"+
		"\u0157\n\4\2\2\u0154\u0155\7^\2\2\u0155\u0157\13\2\2\2\u0156\u0153\3\2"+
		"\2\2\u0156\u0154\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0159\7)\2\2\u0159"+
		"n\3\2\2\2\u015a\u0160\7$\2\2\u015b\u015f\n\5\2\2\u015c\u015d\7^\2\2\u015d"+
		"\u015f\13\2\2\2\u015e\u015b\3\2\2\2\u015e\u015c\3\2\2\2\u015f\u0162\3"+
		"\2\2\2\u0160\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u0163\3\2\2\2\u0162"+
		"\u0160\3\2\2\2\u0163\u0164\7$\2\2\u0164p\3\2\2\2\u0165\u0169\t\6\2\2\u0166"+
		"\u0168\t\7\2\2\u0167\u0166\3\2\2\2\u0168\u016b\3\2\2\2\u0169\u0167\3\2"+
		"\2\2\u0169\u016a\3\2\2\2\u016a\u0173\3\2\2\2\u016b\u0169\3\2\2\2\u016c"+
		"\u016e\7a\2\2\u016d\u016f\t\7\2\2\u016e\u016d\3\2\2\2\u016f\u0170\3\2"+
		"\2\2\u0170\u016e\3\2\2\2\u0170\u0171\3\2\2\2\u0171\u0173\3\2\2\2\u0172"+
		"\u0165\3\2\2\2\u0172\u016c\3\2\2\2\u0173r\3\2\2\2\u0174\u0175\7\61\2\2"+
		"\u0175\u0176\7\61\2\2\u0176\u017a\3\2\2\2\u0177\u0179\n\b\2\2\u0178\u0177"+
		"\3\2\2\2\u0179\u017c\3\2\2\2\u017a\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017b"+
		"\u017d\3\2\2\2\u017c\u017a\3\2\2\2\u017d\u017e\b:\3\2\u017et\3\2\2\2\u017f"+
		"\u0180\7\61\2\2\u0180\u0181\7,\2\2\u0181\u0185\3\2\2\2\u0182\u0184\13"+
		"\2\2\2\u0183\u0182\3\2\2\2\u0184\u0187\3\2\2\2\u0185\u0186\3\2\2\2\u0185"+
		"\u0183\3\2\2\2\u0186\u0188\3\2\2\2\u0187\u0185\3\2\2\2\u0188\u0189\7,"+
		"\2\2\u0189\u018a\7\61\2\2\u018a\u018b\3\2\2\2\u018b\u018c\b;\3\2\u018c"+
		"v\3\2\2\2\20\2\u0138\u0141\u0146\u014c\u0150\u0156\u015e\u0160\u0169\u0170"+
		"\u0172\u017a\u0185\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}