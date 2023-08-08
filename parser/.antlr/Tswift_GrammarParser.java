// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Tswift_GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, ASIG=3, DIV=4, INT=5, CONSOLE=6, ENBLANCO=7, ENTERO=8, 
		DECIMAL=9, ID=10, RETURN=11, UL_C=12;
	public static final int
		RULE_s = 0, RULE_a = 1;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "a"
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

	@Override
	public String getGrammarFileName() { return "Tswift_Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Tswift_GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public AContext a() {
			return getRuleContext(AContext.class,0);
		}
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(4);
			a();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AContext extends ParserRuleContext {
		public AContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_a; }
	 
		public AContext() { }
		public void copyFrom(AContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class AEPSILUMContext extends AContext {
		public AEPSILUMContext(AContext ctx) { copyFrom(ctx); }
	}
	public static class ADECIMALContext extends AContext {
		public TerminalNode DECIMAL() { return getToken(Tswift_GrammarParser.DECIMAL, 0); }
		public ADECIMALContext(AContext ctx) { copyFrom(ctx); }
	}
	public static class AIDContext extends AContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public AIDContext(AContext ctx) { copyFrom(ctx); }
	}
	public static class AENTEROContext extends AContext {
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarParser.ENTERO, 0); }
		public AENTEROContext(AContext ctx) { copyFrom(ctx); }
	}
	public static class A0Context extends AContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public AContext a() {
			return getRuleContext(AContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public A0Context(AContext ctx) { copyFrom(ctx); }
	}

	public final AContext a() throws RecognitionException {
		AContext _localctx = new AContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_a);
		try {
			setState(14);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARIZQ:
				_localctx = new A0Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(6);
				match(PARIZQ);
				setState(7);
				a();
				setState(8);
				match(PARDER);
				}
				break;
			case ENTERO:
				_localctx = new AENTEROContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(10);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ADECIMALContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(11);
				match(DECIMAL);
				}
				break;
			case ID:
				_localctx = new AIDContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(12);
				match(ID);
				}
				break;
			case EOF:
			case PARDER:
				_localctx = new AEPSILUMContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\16\23\4\2\t\2\4\3"+
		"\t\3\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3\21\n\3\3\3\2\2\4\2\4"+
		"\2\2\2\24\2\6\3\2\2\2\4\20\3\2\2\2\6\7\5\4\3\2\7\3\3\2\2\2\b\t\7\4\2\2"+
		"\t\n\5\4\3\2\n\13\7\3\2\2\13\21\3\2\2\2\f\21\7\n\2\2\r\21\7\13\2\2\16"+
		"\21\7\f\2\2\17\21\3\2\2\2\20\b\3\2\2\2\20\f\3\2\2\2\20\r\3\2\2\2\20\16"+
		"\3\2\2\2\20\17\3\2\2\2\21\5\3\2\2\2\3\20";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}