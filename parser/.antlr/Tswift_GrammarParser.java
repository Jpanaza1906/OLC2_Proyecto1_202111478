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
		T__0=1, PARDER=2, PARIZQ=3, LLAVEIZQ=4, LLAVEDER=5, DOSPT=6, INTERROGACION=7, 
		MASIGUAL=8, MENOSIGUAL=9, IGUAL=10, DIV=11, MOD=12, MENOS=13, MAS=14, 
		POR=15, IGUALIGUAL=16, DIFERENTE=17, MAYORIGUAL=18, MENORIGUAL=19, MAYOR=20, 
		MENOR=21, AND=22, OR=23, NOT=24, PRINT=25, VAR=26, LET=27, TRUE=28, FALSE=29, 
		RETURN=30, IF=31, ELSE=32, SWITCH=33, CASE=34, BREAK=35, DEFAULT=36, WHILE=37, 
		FOR=38, IN=39, RANGO=40, GUARD=41, CONTINUE=42, NIL=43, INT=44, FLOAT=45, 
		BOOL=46, CHARACTER=47, STRING=48, ENBLANCO=49, ENTERO=50, DECIMAL=51, 
		CARACTER=52, CADENA=53, ID=54, UL_C=55, ML_C=56;
	public static final int
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_declaracion = 3, 
		RULE_constante = 4, RULE_tipo = 5, RULE_e = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "declaracion", "constante", "tipo", 
			"e"
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
			"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", 
			"UL_C", "ML_C"
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
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	 
		public SContext() { }
		public void copyFrom(SContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class SLSentenciasContext extends SContext {
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public SLSentenciasContext(SContext ctx) { copyFrom(ctx); }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			_localctx = new SLSentenciasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			l_sentencias();
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

	public static class L_sentenciasContext extends ParserRuleContext {
		public L_sentenciasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_sentencias; }
	 
		public L_sentenciasContext() { }
		public void copyFrom(L_sentenciasContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class L_SentenciaContext extends L_sentenciasContext {
		public List<SentenciaContext> sentencia() {
			return getRuleContexts(SentenciaContext.class);
		}
		public SentenciaContext sentencia(int i) {
			return getRuleContext(SentenciaContext.class,i);
		}
		public L_SentenciaContext(L_sentenciasContext ctx) { copyFrom(ctx); }
	}

	public final L_sentenciasContext l_sentencias() throws RecognitionException {
		L_sentenciasContext _localctx = new L_sentenciasContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_l_sentencias);
		int _la;
		try {
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(19);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << VAR) | (1L << LET))) != 0)) {
				{
				{
				setState(16);
				sentencia();
				}
				}
				setState(21);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class SentenciaContext extends ParserRuleContext {
		public SentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentencia; }
	 
		public SentenciaContext() { }
		public void copyFrom(SentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class S_ConsolaContext extends SentenciaContext {
		public TerminalNode PRINT() { return getToken(Tswift_GrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public S_ConsolaContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_ConstanteContext extends SentenciaContext {
		public ConstanteContext constante() {
			return getRuleContext(ConstanteContext.class,0);
		}
		public S_ConstanteContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_DeclaracionContext extends SentenciaContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public S_DeclaracionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(36);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new S_ConsolaContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(22);
				match(PRINT);
				setState(23);
				match(PARIZQ);
				setState(24);
				e(0);
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(25);
					match(T__0);
					setState(26);
					e(0);
					}
					}
					setState(31);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(32);
				match(PARDER);
				}
				break;
			case VAR:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				declaracion();
				}
				break;
			case LET:
				_localctx = new S_ConstanteContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				constante();
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

	public static class DeclaracionContext extends ParserRuleContext {
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	 
		public DeclaracionContext() { }
		public void copyFrom(DeclaracionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declaracion_Tipo_ValContext extends DeclaracionContext {
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Declaracion_Tipo_ValContext(DeclaracionContext ctx) { copyFrom(ctx); }
	}
	public static class Declaracion_Tipo_noValContext extends DeclaracionContext {
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode INTERROGACION() { return getToken(Tswift_GrammarParser.INTERROGACION, 0); }
		public Declaracion_Tipo_noValContext(DeclaracionContext ctx) { copyFrom(ctx); }
	}
	public static class Declaracion_ValContext extends DeclaracionContext {
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Declaracion_ValContext(DeclaracionContext ctx) { copyFrom(ctx); }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaracion);
		try {
			setState(55);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				match(VAR);
				setState(39);
				match(ID);
				setState(40);
				match(DOSPT);
				setState(41);
				tipo();
				setState(42);
				match(IGUAL);
				setState(43);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				match(VAR);
				setState(46);
				match(ID);
				setState(47);
				match(IGUAL);
				setState(48);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				match(VAR);
				setState(50);
				match(ID);
				setState(51);
				match(DOSPT);
				setState(52);
				tipo();
				setState(53);
				match(INTERROGACION);
				}
				break;
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

	public static class ConstanteContext extends ParserRuleContext {
		public ConstanteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constante; }
	 
		public ConstanteContext() { }
		public void copyFrom(ConstanteContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Constante_ValContext extends ConstanteContext {
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Constante_ValContext(ConstanteContext ctx) { copyFrom(ctx); }
	}
	public static class Constante_Tipo_ValContext extends ConstanteContext {
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Constante_Tipo_ValContext(ConstanteContext ctx) { copyFrom(ctx); }
	}

	public final ConstanteContext constante() throws RecognitionException {
		ConstanteContext _localctx = new ConstanteContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_constante);
		try {
			setState(68);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new Constante_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(57);
				match(LET);
				setState(58);
				match(ID);
				setState(59);
				match(DOSPT);
				setState(60);
				tipo();
				setState(61);
				match(IGUAL);
				setState(62);
				e(0);
				}
				break;
			case 2:
				_localctx = new Constante_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				match(LET);
				setState(65);
				match(ID);
				setState(66);
				match(IGUAL);
				setState(67);
				e(0);
				}
				break;
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

	public static class TipoContext extends ParserRuleContext {
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	 
		public TipoContext() { }
		public void copyFrom(TipoContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Tipo_FloatContext extends TipoContext {
		public TerminalNode FLOAT() { return getToken(Tswift_GrammarParser.FLOAT, 0); }
		public Tipo_FloatContext(TipoContext ctx) { copyFrom(ctx); }
	}
	public static class Tipo_CharacterContext extends TipoContext {
		public TerminalNode CHARACTER() { return getToken(Tswift_GrammarParser.CHARACTER, 0); }
		public Tipo_CharacterContext(TipoContext ctx) { copyFrom(ctx); }
	}
	public static class Tipo_StringContext extends TipoContext {
		public TerminalNode STRING() { return getToken(Tswift_GrammarParser.STRING, 0); }
		public Tipo_StringContext(TipoContext ctx) { copyFrom(ctx); }
	}
	public static class Tipo_BoolContext extends TipoContext {
		public TerminalNode BOOL() { return getToken(Tswift_GrammarParser.BOOL, 0); }
		public Tipo_BoolContext(TipoContext ctx) { copyFrom(ctx); }
	}
	public static class Tipo_IntContext extends TipoContext {
		public TerminalNode INT() { return getToken(Tswift_GrammarParser.INT, 0); }
		public Tipo_IntContext(TipoContext ctx) { copyFrom(ctx); }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_tipo);
		try {
			setState(75);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(70);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(72);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(73);
				match(BOOL);
				}
				break;
			case CHARACTER:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(74);
				match(CHARACTER);
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

	public static class EContext extends ParserRuleContext {
		public EContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_e; }
	 
		public EContext() { }
		public void copyFrom(EContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Expr_RelContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode IGUALIGUAL() { return getToken(Tswift_GrammarParser.IGUALIGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(Tswift_GrammarParser.DIFERENTE, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Tswift_GrammarParser.MAYORIGUAL, 0); }
		public TerminalNode MAYOR() { return getToken(Tswift_GrammarParser.MAYOR, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Tswift_GrammarParser.MENORIGUAL, 0); }
		public TerminalNode MENOR() { return getToken(Tswift_GrammarParser.MENOR, 0); }
		public Expr_RelContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_DecimalContext extends EContext {
		public TerminalNode DECIMAL() { return getToken(Tswift_GrammarParser.DECIMAL, 0); }
		public Expr_DecimalContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_CaracterContext extends EContext {
		public TerminalNode CARACTER() { return getToken(Tswift_GrammarParser.CARACTER, 0); }
		public Expr_CaracterContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_SumResContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MAS() { return getToken(Tswift_GrammarParser.MAS, 0); }
		public TerminalNode MENOS() { return getToken(Tswift_GrammarParser.MENOS, 0); }
		public Expr_SumResContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_NegContext extends EContext {
		public Token op;
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode MENOS() { return getToken(Tswift_GrammarParser.MENOS, 0); }
		public TerminalNode NOT() { return getToken(Tswift_GrammarParser.NOT, 0); }
		public Expr_NegContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_MulDivContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode POR() { return getToken(Tswift_GrammarParser.POR, 0); }
		public TerminalNode DIV() { return getToken(Tswift_GrammarParser.DIV, 0); }
		public Expr_MulDivContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_NilContext extends EContext {
		public TerminalNode NIL() { return getToken(Tswift_GrammarParser.NIL, 0); }
		public Expr_NilContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_CadenaContext extends EContext {
		public TerminalNode CADENA() { return getToken(Tswift_GrammarParser.CADENA, 0); }
		public Expr_CadenaContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_IdContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public Expr_IdContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_ModContext extends EContext {
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MOD() { return getToken(Tswift_GrammarParser.MOD, 0); }
		public Expr_ModContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_ParContext extends EContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Expr_ParContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_LogicaContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode AND() { return getToken(Tswift_GrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(Tswift_GrammarParser.OR, 0); }
		public Expr_LogicaContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_BooleanoContext extends EContext {
		public Token op;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarParser.FALSE, 0); }
		public Expr_BooleanoContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_EnteroContext extends EContext {
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarParser.ENTERO, 0); }
		public Expr_EnteroContext(EContext ctx) { copyFrom(ctx); }
	}

	public final EContext e() throws RecognitionException {
		return e(0);
	}

	private EContext e(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		EContext _localctx = new EContext(_ctx, _parentState);
		EContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
			case NOT:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(78);
				((Expr_NegContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MENOS || _la==NOT) ) {
					((Expr_NegContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(79);
				e(12);
				}
				break;
			case TRUE:
			case FALSE:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(80);
				((Expr_BooleanoContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TRUE || _la==FALSE) ) {
					((Expr_BooleanoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case NIL:
				{
				_localctx = new Expr_NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(81);
				match(NIL);
				}
				break;
			case ID:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(82);
				match(ID);
				}
				break;
			case DECIMAL:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(83);
				match(DECIMAL);
				}
				break;
			case ENTERO:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(84);
				match(ENTERO);
				}
				break;
			case CADENA:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(85);
				match(CADENA);
				}
				break;
			case CARACTER:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(86);
				match(CARACTER);
				}
				break;
			case PARIZQ:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(87);
				match(PARIZQ);
				setState(88);
				e(0);
				setState(89);
				match(PARDER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(110);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(108);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(93);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(94);
						((Expr_RelContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IGUALIGUAL) | (1L << DIFERENTE) | (1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << MAYOR) | (1L << MENOR))) != 0)) ) {
							((Expr_RelContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(95);
						e(15);
						}
						break;
					case 2:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(96);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(97);
						((Expr_LogicaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==AND || _la==OR) ) {
							((Expr_LogicaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(98);
						e(14);
						}
						break;
					case 3:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(99);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(100);
						((Expr_MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIV || _la==POR) ) {
							((Expr_MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(101);
						e(12);
						}
						break;
					case 4:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(102);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(103);
						((Expr_SumResContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MENOS || _la==MAS) ) {
							((Expr_SumResContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(104);
						e(11);
						}
						break;
					case 5:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(105);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(106);
						match(MOD);
						setState(107);
						e(10);
						}
						break;
					}
					} 
				}
				setState(112);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return e_sempred((EContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean e_sempred(EContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 11);
		case 3:
			return precpred(_ctx, 10);
		case 4:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3:t\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\3\7\3\24\n\3\f\3\16"+
		"\3\27\13\3\3\4\3\4\3\4\3\4\3\4\7\4\36\n\4\f\4\16\4!\13\4\3\4\3\4\3\4\3"+
		"\4\5\4\'\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\5\5:\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6G\n"+
		"\6\3\7\3\7\3\7\3\7\3\7\5\7N\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\5\b^\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\7\bo\n\b\f\b\16\br\13\b\3\b\2\3\16\t\2\4\6\b\n\f\16"+
		"\2\b\4\2\17\17\32\32\3\2\36\37\3\2\22\27\3\2\30\31\4\2\r\r\21\21\3\2\17"+
		"\20\2\u0084\2\20\3\2\2\2\4\25\3\2\2\2\6&\3\2\2\2\b9\3\2\2\2\nF\3\2\2\2"+
		"\fM\3\2\2\2\16]\3\2\2\2\20\21\5\4\3\2\21\3\3\2\2\2\22\24\5\6\4\2\23\22"+
		"\3\2\2\2\24\27\3\2\2\2\25\23\3\2\2\2\25\26\3\2\2\2\26\5\3\2\2\2\27\25"+
		"\3\2\2\2\30\31\7\33\2\2\31\32\7\5\2\2\32\37\5\16\b\2\33\34\7\3\2\2\34"+
		"\36\5\16\b\2\35\33\3\2\2\2\36!\3\2\2\2\37\35\3\2\2\2\37 \3\2\2\2 \"\3"+
		"\2\2\2!\37\3\2\2\2\"#\7\4\2\2#\'\3\2\2\2$\'\5\b\5\2%\'\5\n\6\2&\30\3\2"+
		"\2\2&$\3\2\2\2&%\3\2\2\2\'\7\3\2\2\2()\7\34\2\2)*\78\2\2*+\7\b\2\2+,\5"+
		"\f\7\2,-\7\f\2\2-.\5\16\b\2.:\3\2\2\2/\60\7\34\2\2\60\61\78\2\2\61\62"+
		"\7\f\2\2\62:\5\16\b\2\63\64\7\34\2\2\64\65\78\2\2\65\66\7\b\2\2\66\67"+
		"\5\f\7\2\678\7\t\2\28:\3\2\2\29(\3\2\2\29/\3\2\2\29\63\3\2\2\2:\t\3\2"+
		"\2\2;<\7\35\2\2<=\78\2\2=>\7\b\2\2>?\5\f\7\2?@\7\f\2\2@A\5\16\b\2AG\3"+
		"\2\2\2BC\7\35\2\2CD\78\2\2DE\7\f\2\2EG\5\16\b\2F;\3\2\2\2FB\3\2\2\2G\13"+
		"\3\2\2\2HN\7.\2\2IN\7/\2\2JN\7\62\2\2KN\7\60\2\2LN\7\61\2\2MH\3\2\2\2"+
		"MI\3\2\2\2MJ\3\2\2\2MK\3\2\2\2ML\3\2\2\2N\r\3\2\2\2OP\b\b\1\2PQ\t\2\2"+
		"\2Q^\5\16\b\16R^\t\3\2\2S^\7-\2\2T^\78\2\2U^\7\65\2\2V^\7\64\2\2W^\7\67"+
		"\2\2X^\7\66\2\2YZ\7\5\2\2Z[\5\16\b\2[\\\7\4\2\2\\^\3\2\2\2]O\3\2\2\2]"+
		"R\3\2\2\2]S\3\2\2\2]T\3\2\2\2]U\3\2\2\2]V\3\2\2\2]W\3\2\2\2]X\3\2\2\2"+
		"]Y\3\2\2\2^p\3\2\2\2_`\f\20\2\2`a\t\4\2\2ao\5\16\b\21bc\f\17\2\2cd\t\5"+
		"\2\2do\5\16\b\20ef\f\r\2\2fg\t\6\2\2go\5\16\b\16hi\f\f\2\2ij\t\7\2\2j"+
		"o\5\16\b\rkl\f\13\2\2lm\7\16\2\2mo\5\16\b\fn_\3\2\2\2nb\3\2\2\2ne\3\2"+
		"\2\2nh\3\2\2\2nk\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2q\17\3\2\2\2rp\3"+
		"\2\2\2\13\25\37&9FM]np";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}