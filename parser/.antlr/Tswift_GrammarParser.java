// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto1_202111478/OLC2_Proyecto1_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.9.2
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
		T__0=1, PARDER=2, PARIZQ=3, LLAVEIZQ=4, LLAVEDER=5, CORCHETEIZQ=6, CORCHETEDER=7, 
		DOSPT=8, PTCOMA=9, INTERROGACION=10, PUNTO=11, MASIGUAL=12, MENOSIGUAL=13, 
		IGUAL=14, DIV=15, MOD=16, MENOS=17, MAS=18, POR=19, IGUALIGUAL=20, DIFERENTE=21, 
		MAYORIGUAL=22, MENORIGUAL=23, MAYOR=24, MENOR=25, AND=26, OR=27, NOT=28, 
		PRINT=29, VAR=30, LET=31, TRUE=32, FALSE=33, IF=34, ELSE=35, SWITCH=36, 
		CASE=37, DEFAULT=38, WHILE=39, FOR=40, IN=41, RANGO=42, GUARD=43, CONTINUE=44, 
		RETURN=45, BREAK=46, NIL=47, APPEND=48, REMOVELAST=49, REMOVE=50, AT=51, 
		ISEMPTY=52, COUNT=53, INT=54, FLOAT=55, BOOL=56, CHARACTER=57, STRING=58, 
		ENBLANCO=59, ENTERO=60, DECIMAL=61, CARACTER=62, CADENA=63, ID=64, UL_C=65, 
		ML_C=66;
	public static final int
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_declaracion = 3, 
		RULE_constante = 4, RULE_asignacion = 5, RULE_if_sentencia = 6, RULE_switch_sentencia = 7, 
		RULE_l_casos = 8, RULE_l_default = 9, RULE_guard_sentencia = 10, RULE_while_sentencia = 11, 
		RULE_for_sentencia = 12, RULE_rango_p = 13, RULE_trans_sentencia = 14, 
		RULE_dec_vector = 15, RULE_def_vector = 16, RULE_func_vector = 17, RULE_tipo = 18, 
		RULE_e = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "declaracion", "constante", "asignacion", 
			"if_sentencia", "switch_sentencia", "l_casos", "l_default", "guard_sentencia", 
			"while_sentencia", "for_sentencia", "rango_p", "trans_sentencia", "dec_vector", 
			"def_vector", "func_vector", "tipo", "e"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", 
			"'?'", "'.'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'", "'*'", 
			"'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", "'!'", 
			"'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", 
			"'continue'", "'return'", "'break'", "'nil'", "'append'", "'removeLast'", 
			"'remove'", "'at'", "'IsEmpty'", "'count'", "'Int'", "'Float'", "'Bool'", 
			"'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", 
			"CORCHETEDER", "DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "MASIGUAL", 
			"MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL", 
			"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR", 
			"NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH", 
			"CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE", 
			"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", 
			"COUNT", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO", 
			"DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C"
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
			setState(40);
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
		try {
			int _alt;
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(42);
					sentencia();
					}
					} 
				}
				setState(47);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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
	public static class S_ForContext extends SentenciaContext {
		public For_sentenciaContext for_sentencia() {
			return getRuleContext(For_sentenciaContext.class,0);
		}
		public S_ForContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_AsignacionContext extends SentenciaContext {
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_AsignacionContext(SentenciaContext ctx) { copyFrom(ctx); }
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
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_ConsolaContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_ConstanteContext extends SentenciaContext {
		public ConstanteContext constante() {
			return getRuleContext(ConstanteContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_ConstanteContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_GuardContext extends SentenciaContext {
		public Guard_sentenciaContext guard_sentencia() {
			return getRuleContext(Guard_sentenciaContext.class,0);
		}
		public S_GuardContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_WhileContext extends SentenciaContext {
		public While_sentenciaContext while_sentencia() {
			return getRuleContext(While_sentenciaContext.class,0);
		}
		public S_WhileContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Funcion_VectorContext extends SentenciaContext {
		public Func_vectorContext func_vector() {
			return getRuleContext(Func_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Funcion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Declaracion_VectorContext extends SentenciaContext {
		public Dec_vectorContext dec_vector() {
			return getRuleContext(Dec_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Declaracion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_IfContext extends SentenciaContext {
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public S_IfContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_TransicionContext extends SentenciaContext {
		public Trans_sentenciaContext trans_sentencia() {
			return getRuleContext(Trans_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_TransicionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_DeclaracionContext extends SentenciaContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_DeclaracionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(91);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				_localctx = new S_ConsolaContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(48);
				match(PRINT);
				setState(49);
				match(PARIZQ);
				setState(50);
				e(0);
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(51);
					match(T__0);
					setState(52);
					e(0);
					}
					}
					setState(57);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(58);
				match(PARDER);
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(59);
					match(PTCOMA);
					}
				}

				}
				break;
			case 2:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(62);
				declaracion();
				setState(64);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(63);
					match(PTCOMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new S_ConstanteContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(66);
				constante();
				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(67);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(70);
				asignacion();
				setState(72);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(71);
					match(PTCOMA);
					}
				}

				}
				break;
			case 5:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(74);
				if_sentencia();
				}
				break;
			case 6:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(75);
				switch_sentencia();
				}
				break;
			case 7:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(76);
				guard_sentencia();
				}
				break;
			case 8:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(77);
				while_sentencia();
				}
				break;
			case 9:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(78);
				for_sentencia();
				}
				break;
			case 10:
				_localctx = new S_TransicionContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(79);
				trans_sentencia();
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(80);
					match(PTCOMA);
					}
				}

				}
				break;
			case 11:
				_localctx = new S_Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(83);
				dec_vector();
				setState(85);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(84);
					match(PTCOMA);
					}
				}

				}
				break;
			case 12:
				_localctx = new S_Funcion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(87);
				func_vector();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(88);
					match(PTCOMA);
					}
				}

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
			setState(110);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(93);
				match(VAR);
				setState(94);
				match(ID);
				setState(95);
				match(DOSPT);
				setState(96);
				tipo();
				setState(97);
				match(IGUAL);
				setState(98);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				match(VAR);
				setState(101);
				match(ID);
				setState(102);
				match(IGUAL);
				setState(103);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(VAR);
				setState(105);
				match(ID);
				setState(106);
				match(DOSPT);
				setState(107);
				tipo();
				setState(108);
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
			setState(123);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new Constante_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(LET);
				setState(113);
				match(ID);
				setState(114);
				match(DOSPT);
				setState(115);
				tipo();
				setState(116);
				match(IGUAL);
				setState(117);
				e(0);
				}
				break;
			case 2:
				_localctx = new Constante_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(119);
				match(LET);
				setState(120);
				match(ID);
				setState(121);
				match(IGUAL);
				setState(122);
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

	public static class AsignacionContext extends ParserRuleContext {
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	 
		public AsignacionContext() { }
		public void copyFrom(AsignacionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ResAsgContext extends AsignacionContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode MENOSIGUAL() { return getToken(Tswift_GrammarParser.MENOSIGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public ResAsgContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class SumAsgContext extends AsignacionContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode MASIGUAL() { return getToken(Tswift_GrammarParser.MASIGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public SumAsgContext(AsignacionContext ctx) { copyFrom(ctx); }
	}
	public static class AsigContext extends AsignacionContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public AsigContext(AsignacionContext ctx) { copyFrom(ctx); }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_asignacion);
		try {
			setState(134);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new SumAsgContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				match(ID);
				setState(126);
				match(MASIGUAL);
				setState(127);
				e(0);
				}
				break;
			case 2:
				_localctx = new ResAsgContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(128);
				match(ID);
				setState(129);
				match(MENOSIGUAL);
				setState(130);
				e(0);
				}
				break;
			case 3:
				_localctx = new AsigContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(131);
				match(ID);
				setState(132);
				match(IGUAL);
				setState(133);
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

	public static class If_sentenciaContext extends ParserRuleContext {
		public If_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_sentencia; }
	 
		public If_sentenciaContext() { }
		public void copyFrom(If_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class If_ElseIfContext extends If_sentenciaContext {
		public TerminalNode IF() { return getToken(Tswift_GrammarParser.IF, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public TerminalNode ELSE() { return getToken(Tswift_GrammarParser.ELSE, 0); }
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public If_ElseIfContext(If_sentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class If_ElseContext extends If_sentenciaContext {
		public TerminalNode IF() { return getToken(Tswift_GrammarParser.IF, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(Tswift_GrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(Tswift_GrammarParser.LLAVEIZQ, i);
		}
		public List<L_sentenciasContext> l_sentencias() {
			return getRuleContexts(L_sentenciasContext.class);
		}
		public L_sentenciasContext l_sentencias(int i) {
			return getRuleContext(L_sentenciasContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(Tswift_GrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(Tswift_GrammarParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(Tswift_GrammarParser.ELSE, 0); }
		public If_ElseContext(If_sentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class If_SimpleContext extends If_sentenciaContext {
		public TerminalNode IF() { return getToken(Tswift_GrammarParser.IF, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public If_SimpleContext(If_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final If_sentenciaContext if_sentencia() throws RecognitionException {
		If_sentenciaContext _localctx = new If_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_if_sentencia);
		try {
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new If_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				match(IF);
				setState(137);
				e(0);
				setState(138);
				match(LLAVEIZQ);
				setState(139);
				l_sentencias();
				setState(140);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new If_ElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(142);
				match(IF);
				setState(143);
				e(0);
				setState(144);
				match(LLAVEIZQ);
				setState(145);
				l_sentencias();
				setState(146);
				match(LLAVEDER);
				setState(147);
				match(ELSE);
				setState(148);
				match(LLAVEIZQ);
				setState(149);
				l_sentencias();
				setState(150);
				match(LLAVEDER);
				}
				break;
			case 3:
				_localctx = new If_ElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(152);
				match(IF);
				setState(153);
				e(0);
				setState(154);
				match(LLAVEIZQ);
				setState(155);
				l_sentencias();
				setState(156);
				match(LLAVEDER);
				setState(157);
				match(ELSE);
				setState(158);
				if_sentencia();
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

	public static class Switch_sentenciaContext extends ParserRuleContext {
		public Switch_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_sentencia; }
	 
		public Switch_sentenciaContext() { }
		public void copyFrom(Switch_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class SwitchContext extends Switch_sentenciaContext {
		public TerminalNode SWITCH() { return getToken(Tswift_GrammarParser.SWITCH, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public List<L_casosContext> l_casos() {
			return getRuleContexts(L_casosContext.class);
		}
		public L_casosContext l_casos(int i) {
			return getRuleContext(L_casosContext.class,i);
		}
		public L_defaultContext l_default() {
			return getRuleContext(L_defaultContext.class,0);
		}
		public SwitchContext(Switch_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Switch_sentenciaContext switch_sentencia() throws RecognitionException {
		Switch_sentenciaContext _localctx = new Switch_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_switch_sentencia);
		int _la;
		try {
			_localctx = new SwitchContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(SWITCH);
			setState(163);
			e(0);
			setState(164);
			match(LLAVEIZQ);
			setState(166); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(165);
				l_casos();
				}
				}
				setState(168); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(171);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(170);
				l_default();
				}
			}

			setState(173);
			match(LLAVEDER);
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

	public static class L_casosContext extends ParserRuleContext {
		public L_casosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_casos; }
	 
		public L_casosContext() { }
		public void copyFrom(L_casosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class CaseContext extends L_casosContext {
		public TerminalNode CASE() { return getToken(Tswift_GrammarParser.CASE, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(Tswift_GrammarParser.BREAK, 0); }
		public CaseContext(L_casosContext ctx) { copyFrom(ctx); }
	}

	public final L_casosContext l_casos() throws RecognitionException {
		L_casosContext _localctx = new L_casosContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_l_casos);
		int _la;
		try {
			_localctx = new CaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(CASE);
			setState(176);
			e(0);
			setState(177);
			match(DOSPT);
			setState(178);
			l_sentencias();
			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(179);
				match(BREAK);
				}
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

	public static class L_defaultContext extends ParserRuleContext {
		public L_defaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_default; }
	 
		public L_defaultContext() { }
		public void copyFrom(L_defaultContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DefaultContext extends L_defaultContext {
		public TerminalNode DEFAULT() { return getToken(Tswift_GrammarParser.DEFAULT, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(Tswift_GrammarParser.BREAK, 0); }
		public DefaultContext(L_defaultContext ctx) { copyFrom(ctx); }
	}

	public final L_defaultContext l_default() throws RecognitionException {
		L_defaultContext _localctx = new L_defaultContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_l_default);
		int _la;
		try {
			_localctx = new DefaultContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(DEFAULT);
			setState(183);
			match(DOSPT);
			setState(184);
			l_sentencias();
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(185);
				match(BREAK);
				}
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

	public static class Guard_sentenciaContext extends ParserRuleContext {
		public Guard_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_sentencia; }
	 
		public Guard_sentenciaContext() { }
		public void copyFrom(Guard_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class GuardContext extends Guard_sentenciaContext {
		public TerminalNode GUARD() { return getToken(Tswift_GrammarParser.GUARD, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(Tswift_GrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public Trans_sentenciaContext trans_sentencia() {
			return getRuleContext(Trans_sentenciaContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public GuardContext(Guard_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Guard_sentenciaContext guard_sentencia() throws RecognitionException {
		Guard_sentenciaContext _localctx = new Guard_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_guard_sentencia);
		try {
			_localctx = new GuardContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(GUARD);
			setState(189);
			e(0);
			setState(190);
			match(ELSE);
			setState(191);
			match(LLAVEIZQ);
			setState(192);
			l_sentencias();
			setState(193);
			trans_sentencia();
			setState(194);
			match(LLAVEDER);
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

	public static class While_sentenciaContext extends ParserRuleContext {
		public While_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_sentencia; }
	 
		public While_sentenciaContext() { }
		public void copyFrom(While_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class WhileContext extends While_sentenciaContext {
		public TerminalNode WHILE() { return getToken(Tswift_GrammarParser.WHILE, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public WhileContext(While_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final While_sentenciaContext while_sentencia() throws RecognitionException {
		While_sentenciaContext _localctx = new While_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_while_sentencia);
		try {
			_localctx = new WhileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			match(WHILE);
			setState(197);
			e(0);
			setState(198);
			match(LLAVEIZQ);
			setState(199);
			l_sentencias();
			setState(200);
			match(LLAVEDER);
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

	public static class For_sentenciaContext extends ParserRuleContext {
		public For_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_sentencia; }
	 
		public For_sentenciaContext() { }
		public void copyFrom(For_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ForContext extends For_sentenciaContext {
		public TerminalNode FOR() { return getToken(Tswift_GrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(Tswift_GrammarParser.IN, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public Rango_pContext rango_p() {
			return getRuleContext(Rango_pContext.class,0);
		}
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public ForContext(For_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final For_sentenciaContext for_sentencia() throws RecognitionException {
		For_sentenciaContext _localctx = new For_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_for_sentencia);
		try {
			_localctx = new ForContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			match(FOR);
			setState(203);
			match(ID);
			setState(204);
			match(IN);
			setState(207);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				{
				setState(205);
				rango_p();
				}
				break;
			case 2:
				{
				setState(206);
				e(0);
				}
				break;
			}
			setState(209);
			match(LLAVEIZQ);
			setState(210);
			l_sentencias();
			setState(211);
			match(LLAVEDER);
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

	public static class Rango_pContext extends ParserRuleContext {
		public Rango_pContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rango_p; }
	 
		public Rango_pContext() { }
		public void copyFrom(Rango_pContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class RangoContext extends Rango_pContext {
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode RANGO() { return getToken(Tswift_GrammarParser.RANGO, 0); }
		public RangoContext(Rango_pContext ctx) { copyFrom(ctx); }
	}

	public final Rango_pContext rango_p() throws RecognitionException {
		Rango_pContext _localctx = new Rango_pContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_rango_p);
		try {
			_localctx = new RangoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
			e(0);
			setState(214);
			match(RANGO);
			setState(215);
			e(0);
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

	public static class Trans_sentenciaContext extends ParserRuleContext {
		public Trans_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_trans_sentencia; }
	 
		public Trans_sentenciaContext() { }
		public void copyFrom(Trans_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ReturnContext extends Trans_sentenciaContext {
		public TerminalNode RETURN() { return getToken(Tswift_GrammarParser.RETURN, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public ReturnContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class BreakContext extends Trans_sentenciaContext {
		public TerminalNode BREAK() { return getToken(Tswift_GrammarParser.BREAK, 0); }
		public BreakContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class ContinueContext extends Trans_sentenciaContext {
		public TerminalNode CONTINUE() { return getToken(Tswift_GrammarParser.CONTINUE, 0); }
		public ContinueContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Trans_sentenciaContext trans_sentencia() throws RecognitionException {
		Trans_sentenciaContext _localctx = new Trans_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_trans_sentencia);
		try {
			setState(223);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(219);
				match(RETURN);
				setState(221);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
				case 1:
					{
					setState(220);
					e(0);
					}
					break;
				}
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

	public static class Dec_vectorContext extends ParserRuleContext {
		public Dec_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_vector; }
	 
		public Dec_vectorContext() { }
		public void copyFrom(Dec_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declaracion_VectorContext extends Dec_vectorContext {
		public Token tipod;
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public Def_vectorContext def_vector() {
			return getRuleContext(Def_vectorContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public Declaracion_VectorContext(Dec_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Dec_vectorContext dec_vector() throws RecognitionException {
		Dec_vectorContext _localctx = new Dec_vectorContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_dec_vector);
		int _la;
		try {
			_localctx = new Declaracion_VectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			((Declaracion_VectorContext)_localctx).tipod = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((Declaracion_VectorContext)_localctx).tipod = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(226);
			match(ID);
			setState(227);
			match(DOSPT);
			setState(228);
			match(CORCHETEIZQ);
			setState(229);
			tipo();
			setState(230);
			match(CORCHETEDER);
			setState(231);
			match(IGUAL);
			setState(232);
			def_vector();
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

	public static class Def_vectorContext extends ParserRuleContext {
		public Def_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_def_vector; }
	 
		public Def_vectorContext() { }
		public void copyFrom(Def_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Def_Vector_VacioContext extends Def_vectorContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Def_Vector_VacioContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Def_VectorContext extends Def_vectorContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Def_VectorContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Def_Vector_IdContext extends Def_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public Def_Vector_IdContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Def_vectorContext def_vector() throws RecognitionException {
		Def_vectorContext _localctx = new Def_vectorContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_def_vector);
		int _la;
		try {
			setState(248);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				_localctx = new Def_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(234);
				match(CORCHETEIZQ);
				setState(235);
				e(0);
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(236);
					match(T__0);
					setState(237);
					e(0);
					}
					}
					setState(242);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(243);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Vector_VacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				match(CORCHETEIZQ);
				setState(246);
				match(CORCHETEDER);
				}
				break;
			case 3:
				_localctx = new Def_Vector_IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(247);
				match(ID);
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

	public static class Func_vectorContext extends ParserRuleContext {
		public Func_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_vector; }
	 
		public Func_vectorContext() { }
		public void copyFrom(Func_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Func_Vector_AppendContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(Tswift_GrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Func_Vector_AppendContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Func_Vector_RemoveContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode REMOVE() { return getToken(Tswift_GrammarParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode AT() { return getToken(Tswift_GrammarParser.AT, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Func_Vector_RemoveContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Func_Vector_RemoveLastContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode REMOVELAST() { return getToken(Tswift_GrammarParser.REMOVELAST, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Func_Vector_RemoveLastContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Func_Vector_isEmptyContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(Tswift_GrammarParser.ISEMPTY, 0); }
		public Func_Vector_isEmptyContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class Func_Vector_CountContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(Tswift_GrammarParser.COUNT, 0); }
		public Func_Vector_CountContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Func_vectorContext func_vector() throws RecognitionException {
		Func_vectorContext _localctx = new Func_vectorContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_func_vector);
		try {
			setState(277);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new Func_Vector_AppendContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(250);
				match(ID);
				setState(251);
				match(PUNTO);
				setState(252);
				match(APPEND);
				setState(253);
				match(PARIZQ);
				setState(254);
				e(0);
				setState(255);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Func_Vector_RemoveLastContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(257);
				match(ID);
				setState(258);
				match(PUNTO);
				setState(259);
				match(REMOVELAST);
				setState(260);
				match(PARIZQ);
				setState(261);
				match(PARDER);
				}
				break;
			case 3:
				_localctx = new Func_Vector_RemoveContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(262);
				match(ID);
				setState(263);
				match(PUNTO);
				setState(264);
				match(REMOVE);
				setState(265);
				match(PARIZQ);
				setState(266);
				match(AT);
				setState(267);
				match(DOSPT);
				setState(268);
				e(0);
				setState(269);
				match(PARDER);
				}
				break;
			case 4:
				_localctx = new Func_Vector_isEmptyContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(271);
				match(ID);
				setState(272);
				match(PUNTO);
				setState(273);
				match(ISEMPTY);
				}
				break;
			case 5:
				_localctx = new Func_Vector_CountContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(274);
				match(ID);
				setState(275);
				match(PUNTO);
				setState(276);
				match(COUNT);
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
		enterRule(_localctx, 36, RULE_tipo);
		try {
			setState(284);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(279);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(281);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(282);
				match(BOOL);
				}
				break;
			case CHARACTER:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(283);
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
	public static class Expr_VectorContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Expr_VectorContext(EContext ctx) { copyFrom(ctx); }
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
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(287);
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
				setState(288);
				e(15);
				}
				break;
			case 2:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(289);
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
			case 3:
				{
				_localctx = new Expr_NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(290);
				match(NIL);
				}
				break;
			case 4:
				{
				_localctx = new Expr_VectorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(291);
				match(ID);
				setState(292);
				match(CORCHETEIZQ);
				setState(293);
				e(0);
				setState(294);
				match(CORCHETEDER);
				}
				break;
			case 5:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(296);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(297);
				match(DECIMAL);
				}
				break;
			case 7:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(298);
				match(ENTERO);
				}
				break;
			case 8:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(299);
				match(CADENA);
				}
				break;
			case 9:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(300);
				match(CARACTER);
				}
				break;
			case 10:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(301);
				match(PARIZQ);
				setState(302);
				e(0);
				setState(303);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(324);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(322);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(307);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(308);
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
						setState(309);
						e(15);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(310);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(311);
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
						setState(312);
						e(14);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(313);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(314);
						match(MOD);
						setState(315);
						e(13);
						}
						break;
					case 4:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(316);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(317);
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
						setState(318);
						e(12);
						}
						break;
					case 5:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(319);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(320);
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
						setState(321);
						e(11);
						}
						break;
					}
					} 
				}
				setState(326);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
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
		case 19:
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
			return precpred(_ctx, 12);
		case 3:
			return precpred(_ctx, 11);
		case 4:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3D\u014a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\3\7\3.\n\3\f\3\16\3\61\13\3\3"+
		"\4\3\4\3\4\3\4\3\4\7\48\n\4\f\4\16\4;\13\4\3\4\3\4\5\4?\n\4\3\4\3\4\5"+
		"\4C\n\4\3\4\3\4\5\4G\n\4\3\4\3\4\5\4K\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\5\4T\n\4\3\4\3\4\5\4X\n\4\3\4\3\4\5\4\\\n\4\5\4^\n\4\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5q\n\5\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6~\n\6\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\5\7\u0089\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00a3\n\b\3\t"+
		"\3\t\3\t\3\t\6\t\u00a9\n\t\r\t\16\t\u00aa\3\t\5\t\u00ae\n\t\3\t\3\t\3"+
		"\n\3\n\3\n\3\n\3\n\5\n\u00b7\n\n\3\13\3\13\3\13\3\13\5\13\u00bd\n\13\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16"+
		"\3\16\3\16\5\16\u00d2\n\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20"+
		"\3\20\3\20\3\20\5\20\u00e0\n\20\5\20\u00e2\n\20\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\7\22\u00f1\n\22\f\22\16\22"+
		"\u00f4\13\22\3\22\3\22\3\22\3\22\3\22\5\22\u00fb\n\22\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0118\n\23\3\24"+
		"\3\24\3\24\3\24\3\24\5\24\u011f\n\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u0134"+
		"\n\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\7\25\u0145\n\25\f\25\16\25\u0148\13\25\3\25\2\3(\26\2\4\6\b"+
		"\n\f\16\20\22\24\26\30\32\34\36 \"$&(\2\t\3\2 !\4\2\23\23\36\36\3\2\""+
		"#\4\2\21\21\25\25\3\2\23\24\3\2\34\35\3\2\26\33\2\u0171\2*\3\2\2\2\4/"+
		"\3\2\2\2\6]\3\2\2\2\bp\3\2\2\2\n}\3\2\2\2\f\u0088\3\2\2\2\16\u00a2\3\2"+
		"\2\2\20\u00a4\3\2\2\2\22\u00b1\3\2\2\2\24\u00b8\3\2\2\2\26\u00be\3\2\2"+
		"\2\30\u00c6\3\2\2\2\32\u00cc\3\2\2\2\34\u00d7\3\2\2\2\36\u00e1\3\2\2\2"+
		" \u00e3\3\2\2\2\"\u00fa\3\2\2\2$\u0117\3\2\2\2&\u011e\3\2\2\2(\u0133\3"+
		"\2\2\2*+\5\4\3\2+\3\3\2\2\2,.\5\6\4\2-,\3\2\2\2.\61\3\2\2\2/-\3\2\2\2"+
		"/\60\3\2\2\2\60\5\3\2\2\2\61/\3\2\2\2\62\63\7\37\2\2\63\64\7\5\2\2\64"+
		"9\5(\25\2\65\66\7\3\2\2\668\5(\25\2\67\65\3\2\2\28;\3\2\2\29\67\3\2\2"+
		"\29:\3\2\2\2:<\3\2\2\2;9\3\2\2\2<>\7\4\2\2=?\7\13\2\2>=\3\2\2\2>?\3\2"+
		"\2\2?^\3\2\2\2@B\5\b\5\2AC\7\13\2\2BA\3\2\2\2BC\3\2\2\2C^\3\2\2\2DF\5"+
		"\n\6\2EG\7\13\2\2FE\3\2\2\2FG\3\2\2\2G^\3\2\2\2HJ\5\f\7\2IK\7\13\2\2J"+
		"I\3\2\2\2JK\3\2\2\2K^\3\2\2\2L^\5\16\b\2M^\5\20\t\2N^\5\26\f\2O^\5\30"+
		"\r\2P^\5\32\16\2QS\5\36\20\2RT\7\13\2\2SR\3\2\2\2ST\3\2\2\2T^\3\2\2\2"+
		"UW\5 \21\2VX\7\13\2\2WV\3\2\2\2WX\3\2\2\2X^\3\2\2\2Y[\5$\23\2Z\\\7\13"+
		"\2\2[Z\3\2\2\2[\\\3\2\2\2\\^\3\2\2\2]\62\3\2\2\2]@\3\2\2\2]D\3\2\2\2]"+
		"H\3\2\2\2]L\3\2\2\2]M\3\2\2\2]N\3\2\2\2]O\3\2\2\2]P\3\2\2\2]Q\3\2\2\2"+
		"]U\3\2\2\2]Y\3\2\2\2^\7\3\2\2\2_`\7 \2\2`a\7B\2\2ab\7\n\2\2bc\5&\24\2"+
		"cd\7\20\2\2de\5(\25\2eq\3\2\2\2fg\7 \2\2gh\7B\2\2hi\7\20\2\2iq\5(\25\2"+
		"jk\7 \2\2kl\7B\2\2lm\7\n\2\2mn\5&\24\2no\7\f\2\2oq\3\2\2\2p_\3\2\2\2p"+
		"f\3\2\2\2pj\3\2\2\2q\t\3\2\2\2rs\7!\2\2st\7B\2\2tu\7\n\2\2uv\5&\24\2v"+
		"w\7\20\2\2wx\5(\25\2x~\3\2\2\2yz\7!\2\2z{\7B\2\2{|\7\20\2\2|~\5(\25\2"+
		"}r\3\2\2\2}y\3\2\2\2~\13\3\2\2\2\177\u0080\7B\2\2\u0080\u0081\7\16\2\2"+
		"\u0081\u0089\5(\25\2\u0082\u0083\7B\2\2\u0083\u0084\7\17\2\2\u0084\u0089"+
		"\5(\25\2\u0085\u0086\7B\2\2\u0086\u0087\7\20\2\2\u0087\u0089\5(\25\2\u0088"+
		"\177\3\2\2\2\u0088\u0082\3\2\2\2\u0088\u0085\3\2\2\2\u0089\r\3\2\2\2\u008a"+
		"\u008b\7$\2\2\u008b\u008c\5(\25\2\u008c\u008d\7\6\2\2\u008d\u008e\5\4"+
		"\3\2\u008e\u008f\7\7\2\2\u008f\u00a3\3\2\2\2\u0090\u0091\7$\2\2\u0091"+
		"\u0092\5(\25\2\u0092\u0093\7\6\2\2\u0093\u0094\5\4\3\2\u0094\u0095\7\7"+
		"\2\2\u0095\u0096\7%\2\2\u0096\u0097\7\6\2\2\u0097\u0098\5\4\3\2\u0098"+
		"\u0099\7\7\2\2\u0099\u00a3\3\2\2\2\u009a\u009b\7$\2\2\u009b\u009c\5(\25"+
		"\2\u009c\u009d\7\6\2\2\u009d\u009e\5\4\3\2\u009e\u009f\7\7\2\2\u009f\u00a0"+
		"\7%\2\2\u00a0\u00a1\5\16\b\2\u00a1\u00a3\3\2\2\2\u00a2\u008a\3\2\2\2\u00a2"+
		"\u0090\3\2\2\2\u00a2\u009a\3\2\2\2\u00a3\17\3\2\2\2\u00a4\u00a5\7&\2\2"+
		"\u00a5\u00a6\5(\25\2\u00a6\u00a8\7\6\2\2\u00a7\u00a9\5\22\n\2\u00a8\u00a7"+
		"\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00a8\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab"+
		"\u00ad\3\2\2\2\u00ac\u00ae\5\24\13\2\u00ad\u00ac\3\2\2\2\u00ad\u00ae\3"+
		"\2\2\2\u00ae\u00af\3\2\2\2\u00af\u00b0\7\7\2\2\u00b0\21\3\2\2\2\u00b1"+
		"\u00b2\7\'\2\2\u00b2\u00b3\5(\25\2\u00b3\u00b4\7\n\2\2\u00b4\u00b6\5\4"+
		"\3\2\u00b5\u00b7\7\60\2\2\u00b6\u00b5\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7"+
		"\23\3\2\2\2\u00b8\u00b9\7(\2\2\u00b9\u00ba\7\n\2\2\u00ba\u00bc\5\4\3\2"+
		"\u00bb\u00bd\7\60\2\2\u00bc\u00bb\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\25"+
		"\3\2\2\2\u00be\u00bf\7-\2\2\u00bf\u00c0\5(\25\2\u00c0\u00c1\7%\2\2\u00c1"+
		"\u00c2\7\6\2\2\u00c2\u00c3\5\4\3\2\u00c3\u00c4\5\36\20\2\u00c4\u00c5\7"+
		"\7\2\2\u00c5\27\3\2\2\2\u00c6\u00c7\7)\2\2\u00c7\u00c8\5(\25\2\u00c8\u00c9"+
		"\7\6\2\2\u00c9\u00ca\5\4\3\2\u00ca\u00cb\7\7\2\2\u00cb\31\3\2\2\2\u00cc"+
		"\u00cd\7*\2\2\u00cd\u00ce\7B\2\2\u00ce\u00d1\7+\2\2\u00cf\u00d2\5\34\17"+
		"\2\u00d0\u00d2\5(\25\2\u00d1\u00cf\3\2\2\2\u00d1\u00d0\3\2\2\2\u00d2\u00d3"+
		"\3\2\2\2\u00d3\u00d4\7\6\2\2\u00d4\u00d5\5\4\3\2\u00d5\u00d6\7\7\2\2\u00d6"+
		"\33\3\2\2\2\u00d7\u00d8\5(\25\2\u00d8\u00d9\7,\2\2\u00d9\u00da\5(\25\2"+
		"\u00da\35\3\2\2\2\u00db\u00e2\7\60\2\2\u00dc\u00e2\7.\2\2\u00dd\u00df"+
		"\7/\2\2\u00de\u00e0\5(\25\2\u00df\u00de\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0"+
		"\u00e2\3\2\2\2\u00e1\u00db\3\2\2\2\u00e1\u00dc\3\2\2\2\u00e1\u00dd\3\2"+
		"\2\2\u00e2\37\3\2\2\2\u00e3\u00e4\t\2\2\2\u00e4\u00e5\7B\2\2\u00e5\u00e6"+
		"\7\n\2\2\u00e6\u00e7\7\b\2\2\u00e7\u00e8\5&\24\2\u00e8\u00e9\7\t\2\2\u00e9"+
		"\u00ea\7\20\2\2\u00ea\u00eb\5\"\22\2\u00eb!\3\2\2\2\u00ec\u00ed\7\b\2"+
		"\2\u00ed\u00f2\5(\25\2\u00ee\u00ef\7\3\2\2\u00ef\u00f1\5(\25\2\u00f0\u00ee"+
		"\3\2\2\2\u00f1\u00f4\3\2\2\2\u00f2\u00f0\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3"+
		"\u00f5\3\2\2\2\u00f4\u00f2\3\2\2\2\u00f5\u00f6\7\t\2\2\u00f6\u00fb\3\2"+
		"\2\2\u00f7\u00f8\7\b\2\2\u00f8\u00fb\7\t\2\2\u00f9\u00fb\7B\2\2\u00fa"+
		"\u00ec\3\2\2\2\u00fa\u00f7\3\2\2\2\u00fa\u00f9\3\2\2\2\u00fb#\3\2\2\2"+
		"\u00fc\u00fd\7B\2\2\u00fd\u00fe\7\r\2\2\u00fe\u00ff\7\62\2\2\u00ff\u0100"+
		"\7\5\2\2\u0100\u0101\5(\25\2\u0101\u0102\7\4\2\2\u0102\u0118\3\2\2\2\u0103"+
		"\u0104\7B\2\2\u0104\u0105\7\r\2\2\u0105\u0106\7\63\2\2\u0106\u0107\7\5"+
		"\2\2\u0107\u0118\7\4\2\2\u0108\u0109\7B\2\2\u0109\u010a\7\r\2\2\u010a"+
		"\u010b\7\64\2\2\u010b\u010c\7\5\2\2\u010c\u010d\7\65\2\2\u010d\u010e\7"+
		"\n\2\2\u010e\u010f\5(\25\2\u010f\u0110\7\4\2\2\u0110\u0118\3\2\2\2\u0111"+
		"\u0112\7B\2\2\u0112\u0113\7\r\2\2\u0113\u0118\7\66\2\2\u0114\u0115\7B"+
		"\2\2\u0115\u0116\7\r\2\2\u0116\u0118\7\67\2\2\u0117\u00fc\3\2\2\2\u0117"+
		"\u0103\3\2\2\2\u0117\u0108\3\2\2\2\u0117\u0111\3\2\2\2\u0117\u0114\3\2"+
		"\2\2\u0118%\3\2\2\2\u0119\u011f\78\2\2\u011a\u011f\79\2\2\u011b\u011f"+
		"\7<\2\2\u011c\u011f\7:\2\2\u011d\u011f\7;\2\2\u011e\u0119\3\2\2\2\u011e"+
		"\u011a\3\2\2\2\u011e\u011b\3\2\2\2\u011e\u011c\3\2\2\2\u011e\u011d\3\2"+
		"\2\2\u011f\'\3\2\2\2\u0120\u0121\b\25\1\2\u0121\u0122\t\3\2\2\u0122\u0134"+
		"\5(\25\21\u0123\u0134\t\4\2\2\u0124\u0134\7\61\2\2\u0125\u0126\7B\2\2"+
		"\u0126\u0127\7\b\2\2\u0127\u0128\5(\25\2\u0128\u0129\7\t\2\2\u0129\u0134"+
		"\3\2\2\2\u012a\u0134\7B\2\2\u012b\u0134\7?\2\2\u012c\u0134\7>\2\2\u012d"+
		"\u0134\7A\2\2\u012e\u0134\7@\2\2\u012f\u0130\7\5\2\2\u0130\u0131\5(\25"+
		"\2\u0131\u0132\7\4\2\2\u0132\u0134\3\2\2\2\u0133\u0120\3\2\2\2\u0133\u0123"+
		"\3\2\2\2\u0133\u0124\3\2\2\2\u0133\u0125\3\2\2\2\u0133\u012a\3\2\2\2\u0133"+
		"\u012b\3\2\2\2\u0133\u012c\3\2\2\2\u0133\u012d\3\2\2\2\u0133\u012e\3\2"+
		"\2\2\u0133\u012f\3\2\2\2\u0134\u0146\3\2\2\2\u0135\u0136\f\20\2\2\u0136"+
		"\u0137\t\5\2\2\u0137\u0145\5(\25\21\u0138\u0139\f\17\2\2\u0139\u013a\t"+
		"\6\2\2\u013a\u0145\5(\25\20\u013b\u013c\f\16\2\2\u013c\u013d\7\22\2\2"+
		"\u013d\u0145\5(\25\17\u013e\u013f\f\r\2\2\u013f\u0140\t\7\2\2\u0140\u0145"+
		"\5(\25\16\u0141\u0142\f\f\2\2\u0142\u0143\t\b\2\2\u0143\u0145\5(\25\r"+
		"\u0144\u0135\3\2\2\2\u0144\u0138\3\2\2\2\u0144\u013b\3\2\2\2\u0144\u013e"+
		"\3\2\2\2\u0144\u0141\3\2\2\2\u0145\u0148\3\2\2\2\u0146\u0144\3\2\2\2\u0146"+
		"\u0147\3\2\2\2\u0147)\3\2\2\2\u0148\u0146\3\2\2\2\36/9>BFJSW[]p}\u0088"+
		"\u00a2\u00aa\u00ad\u00b6\u00bc\u00d1\u00df\u00e1\u00f2\u00fa\u0117\u011e"+
		"\u0133\u0144\u0146";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}