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
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_print_sentencia = 3, 
		RULE_declaracion = 4, RULE_constante = 5, RULE_asignacion = 6, RULE_if_sentencia = 7, 
		RULE_switch_sentencia = 8, RULE_l_casos = 9, RULE_l_default = 10, RULE_guard_sentencia = 11, 
		RULE_while_sentencia = 12, RULE_for_sentencia = 13, RULE_rango_p = 14, 
		RULE_trans_sentencia = 15, RULE_dec_vector = 16, RULE_def_vector = 17, 
		RULE_asig_vector = 18, RULE_func_vector = 19, RULE_tipo = 20, RULE_e = 21;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "print_sentencia", "declaracion", "constante", 
			"asignacion", "if_sentencia", "switch_sentencia", "l_casos", "l_default", 
			"guard_sentencia", "while_sentencia", "for_sentencia", "rango_p", "trans_sentencia", 
			"dec_vector", "def_vector", "asig_vector", "func_vector", "tipo", "e"
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
			setState(44);
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
			setState(49);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(46);
					sentencia();
					}
					} 
				}
				setState(51);
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
	public static class S_Asignacion_VectorContext extends SentenciaContext {
		public Asig_vectorContext asig_vector() {
			return getRuleContext(Asig_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Asignacion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_AsignacionContext extends SentenciaContext {
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_AsignacionContext(SentenciaContext ctx) { copyFrom(ctx); }
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
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
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
	public static class S_PrintContext extends SentenciaContext {
		public Print_sentenciaContext print_sentencia() {
			return getRuleContext(Print_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_PrintContext(SentenciaContext ctx) { copyFrom(ctx); }
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
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				print_sentencia();
				setState(54);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(53);
					match(PTCOMA);
					}
				}

				}
				break;
			case 2:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				declaracion();
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(57);
					match(PTCOMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new S_ConstanteContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(60);
				constante();
				setState(62);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(61);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(64);
				asignacion();
				setState(66);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(65);
					match(PTCOMA);
					}
				}

				}
				break;
			case 5:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(68);
				if_sentencia();
				}
				break;
			case 6:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(69);
				switch_sentencia();
				}
				break;
			case 7:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(70);
				guard_sentencia();
				}
				break;
			case 8:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(71);
				while_sentencia();
				}
				break;
			case 9:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(72);
				for_sentencia();
				}
				break;
			case 10:
				_localctx = new S_TransicionContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(73);
				trans_sentencia();
				setState(75);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(74);
					match(PTCOMA);
					}
				}

				}
				break;
			case 11:
				_localctx = new S_Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(77);
				dec_vector();
				setState(79);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(78);
					match(PTCOMA);
					}
				}

				}
				break;
			case 12:
				_localctx = new S_Funcion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(81);
				func_vector();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(82);
					match(PTCOMA);
					}
				}

				}
				break;
			case 13:
				_localctx = new S_Asignacion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(85);
				asig_vector();
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(86);
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

	public static class Print_sentenciaContext extends ParserRuleContext {
		public Print_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_sentencia; }
	 
		public Print_sentenciaContext() { }
		public void copyFrom(Print_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PrintContext extends Print_sentenciaContext {
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
		public PrintContext(Print_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Print_sentenciaContext print_sentencia() throws RecognitionException {
		Print_sentenciaContext _localctx = new Print_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			match(PRINT);
			setState(92);
			match(PARIZQ);
			setState(93);
			e(0);
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(94);
				match(T__0);
				setState(95);
				e(0);
				}
				}
				setState(100);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(101);
			match(PARDER);
			setState(103);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(102);
				match(PTCOMA);
				}
				break;
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
		enterRule(_localctx, 8, RULE_declaracion);
		try {
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(105);
				match(VAR);
				setState(106);
				match(ID);
				setState(107);
				match(DOSPT);
				setState(108);
				tipo();
				setState(109);
				match(IGUAL);
				setState(110);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(VAR);
				setState(113);
				match(ID);
				setState(114);
				match(IGUAL);
				setState(115);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(116);
				match(VAR);
				setState(117);
				match(ID);
				setState(118);
				match(DOSPT);
				setState(119);
				tipo();
				setState(120);
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
		enterRule(_localctx, 10, RULE_constante);
		try {
			setState(135);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new Constante_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(LET);
				setState(125);
				match(ID);
				setState(126);
				match(DOSPT);
				setState(127);
				tipo();
				setState(128);
				match(IGUAL);
				setState(129);
				e(0);
				}
				break;
			case 2:
				_localctx = new Constante_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(131);
				match(LET);
				setState(132);
				match(ID);
				setState(133);
				match(IGUAL);
				setState(134);
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
		enterRule(_localctx, 12, RULE_asignacion);
		try {
			setState(146);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				_localctx = new SumAsgContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(137);
				match(ID);
				setState(138);
				match(MASIGUAL);
				setState(139);
				e(0);
				}
				break;
			case 2:
				_localctx = new ResAsgContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				match(ID);
				setState(141);
				match(MENOSIGUAL);
				setState(142);
				e(0);
				}
				break;
			case 3:
				_localctx = new AsigContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(143);
				match(ID);
				setState(144);
				match(IGUAL);
				setState(145);
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
		enterRule(_localctx, 14, RULE_if_sentencia);
		try {
			setState(172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				_localctx = new If_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(148);
				match(IF);
				setState(149);
				e(0);
				setState(150);
				match(LLAVEIZQ);
				setState(151);
				l_sentencias();
				setState(152);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new If_ElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				match(IF);
				setState(155);
				e(0);
				setState(156);
				match(LLAVEIZQ);
				setState(157);
				l_sentencias();
				setState(158);
				match(LLAVEDER);
				setState(159);
				match(ELSE);
				setState(160);
				match(LLAVEIZQ);
				setState(161);
				l_sentencias();
				setState(162);
				match(LLAVEDER);
				}
				break;
			case 3:
				_localctx = new If_ElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(164);
				match(IF);
				setState(165);
				e(0);
				setState(166);
				match(LLAVEIZQ);
				setState(167);
				l_sentencias();
				setState(168);
				match(LLAVEDER);
				setState(169);
				match(ELSE);
				setState(170);
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
		enterRule(_localctx, 16, RULE_switch_sentencia);
		int _la;
		try {
			_localctx = new SwitchContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			match(SWITCH);
			setState(175);
			e(0);
			setState(176);
			match(LLAVEIZQ);
			setState(178); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(177);
				l_casos();
				}
				}
				setState(180); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(182);
				l_default();
				}
			}

			setState(185);
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
		enterRule(_localctx, 18, RULE_l_casos);
		int _la;
		try {
			_localctx = new CaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			match(CASE);
			setState(188);
			e(0);
			setState(189);
			match(DOSPT);
			setState(190);
			l_sentencias();
			setState(192);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(191);
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
		enterRule(_localctx, 20, RULE_l_default);
		int _la;
		try {
			_localctx = new DefaultContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			match(DEFAULT);
			setState(195);
			match(DOSPT);
			setState(196);
			l_sentencias();
			setState(198);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(197);
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
		enterRule(_localctx, 22, RULE_guard_sentencia);
		try {
			_localctx = new GuardContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(GUARD);
			setState(201);
			e(0);
			setState(202);
			match(ELSE);
			setState(203);
			match(LLAVEIZQ);
			setState(204);
			l_sentencias();
			setState(205);
			trans_sentencia();
			setState(206);
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
		enterRule(_localctx, 24, RULE_while_sentencia);
		try {
			_localctx = new WhileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(WHILE);
			setState(209);
			e(0);
			setState(210);
			match(LLAVEIZQ);
			setState(211);
			l_sentencias();
			setState(212);
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
		enterRule(_localctx, 26, RULE_for_sentencia);
		try {
			_localctx = new ForContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(FOR);
			setState(215);
			match(ID);
			setState(216);
			match(IN);
			setState(219);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				{
				setState(217);
				rango_p();
				}
				break;
			case 2:
				{
				setState(218);
				e(0);
				}
				break;
			}
			setState(221);
			match(LLAVEIZQ);
			setState(222);
			l_sentencias();
			setState(223);
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
		enterRule(_localctx, 28, RULE_rango_p);
		try {
			_localctx = new RangoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			e(0);
			setState(226);
			match(RANGO);
			setState(227);
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
		enterRule(_localctx, 30, RULE_trans_sentencia);
		try {
			setState(235);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(229);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(230);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(231);
				match(RETURN);
				setState(233);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
				case 1:
					{
					setState(232);
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
		enterRule(_localctx, 32, RULE_dec_vector);
		int _la;
		try {
			_localctx = new Declaracion_VectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
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
			setState(238);
			match(ID);
			setState(239);
			match(DOSPT);
			setState(240);
			match(CORCHETEIZQ);
			setState(241);
			tipo();
			setState(242);
			match(CORCHETEDER);
			setState(243);
			match(IGUAL);
			setState(244);
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
		enterRule(_localctx, 34, RULE_def_vector);
		int _la;
		try {
			setState(260);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new Def_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(246);
				match(CORCHETEIZQ);
				setState(247);
				e(0);
				setState(252);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(248);
					match(T__0);
					setState(249);
					e(0);
					}
					}
					setState(254);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(255);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Vector_VacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(257);
				match(CORCHETEIZQ);
				setState(258);
				match(CORCHETEDER);
				}
				break;
			case 3:
				_localctx = new Def_Vector_IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(259);
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

	public static class Asig_vectorContext extends ParserRuleContext {
		public Asig_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asig_vector; }
	 
		public Asig_vectorContext() { }
		public void copyFrom(Asig_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Asig_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public Asig_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Asig_vectorContext asig_vector() throws RecognitionException {
		Asig_vectorContext _localctx = new Asig_vectorContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_asig_vector);
		try {
			_localctx = new Asig_VectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(ID);
			setState(263);
			match(CORCHETEIZQ);
			setState(264);
			e(0);
			setState(265);
			match(CORCHETEDER);
			setState(266);
			match(IGUAL);
			setState(267);
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

	public final Func_vectorContext func_vector() throws RecognitionException {
		Func_vectorContext _localctx = new Func_vectorContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_func_vector);
		try {
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new Func_Vector_AppendContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(269);
				match(ID);
				setState(270);
				match(PUNTO);
				setState(271);
				match(APPEND);
				setState(272);
				match(PARIZQ);
				setState(273);
				e(0);
				setState(274);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Func_Vector_RemoveLastContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(276);
				match(ID);
				setState(277);
				match(PUNTO);
				setState(278);
				match(REMOVELAST);
				setState(279);
				match(PARIZQ);
				setState(280);
				match(PARDER);
				}
				break;
			case 3:
				_localctx = new Func_Vector_RemoveContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(281);
				match(ID);
				setState(282);
				match(PUNTO);
				setState(283);
				match(REMOVE);
				setState(284);
				match(PARIZQ);
				setState(285);
				match(AT);
				setState(286);
				match(DOSPT);
				setState(287);
				e(0);
				setState(288);
				match(PARDER);
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
		enterRule(_localctx, 40, RULE_tipo);
		try {
			setState(297);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(292);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(293);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(294);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(295);
				match(BOOL);
				}
				break;
			case CHARACTER:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(296);
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
	public static class Func_Vector_isEmptyContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(Tswift_GrammarParser.ISEMPTY, 0); }
		public Func_Vector_isEmptyContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Func_Vector_CountContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(Tswift_GrammarParser.COUNT, 0); }
		public Func_Vector_CountContext(EContext ctx) { copyFrom(ctx); }
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
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(300);
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
				setState(301);
				e(17);
				}
				break;
			case 2:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(302);
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
				setState(303);
				match(NIL);
				}
				break;
			case 4:
				{
				_localctx = new Expr_VectorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(304);
				match(ID);
				setState(305);
				match(CORCHETEIZQ);
				setState(306);
				e(0);
				setState(307);
				match(CORCHETEDER);
				}
				break;
			case 5:
				{
				_localctx = new Func_Vector_isEmptyContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(309);
				match(ID);
				setState(310);
				match(PUNTO);
				setState(311);
				match(ISEMPTY);
				}
				break;
			case 6:
				{
				_localctx = new Func_Vector_CountContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(312);
				match(ID);
				setState(313);
				match(PUNTO);
				setState(314);
				match(COUNT);
				}
				break;
			case 7:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(315);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(316);
				match(DECIMAL);
				}
				break;
			case 9:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(317);
				match(ENTERO);
				}
				break;
			case 10:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(318);
				match(CADENA);
				}
				break;
			case 11:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
				match(CARACTER);
				}
				break;
			case 12:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(320);
				match(PARIZQ);
				setState(321);
				e(0);
				setState(322);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(343);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(341);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(326);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(327);
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
						setState(328);
						e(17);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(329);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(330);
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
						setState(331);
						e(16);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(332);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(333);
						match(MOD);
						setState(334);
						e(15);
						}
						break;
					case 4:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(335);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(336);
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
						setState(337);
						e(14);
						}
						break;
					case 5:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(338);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(339);
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
						setState(340);
						e(13);
						}
						break;
					}
					} 
				}
				setState(345);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
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
		case 21:
			return e_sempred((EContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean e_sempred(EContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 16);
		case 1:
			return precpred(_ctx, 15);
		case 2:
			return precpred(_ctx, 14);
		case 3:
			return precpred(_ctx, 13);
		case 4:
			return precpred(_ctx, 12);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3D\u015d\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\3\2\3\3\7\3\62"+
		"\n\3\f\3\16\3\65\13\3\3\4\3\4\5\49\n\4\3\4\3\4\5\4=\n\4\3\4\3\4\5\4A\n"+
		"\4\3\4\3\4\5\4E\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4N\n\4\3\4\3\4\5\4R"+
		"\n\4\3\4\3\4\5\4V\n\4\3\4\3\4\5\4Z\n\4\5\4\\\n\4\3\5\3\5\3\5\3\5\3\5\7"+
		"\5c\n\5\f\5\16\5f\13\5\3\5\3\5\5\5j\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6}\n\6\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\5\7\u008a\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\5\b\u0095\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00af\n\t\3\n\3\n\3\n\3"+
		"\n\6\n\u00b5\n\n\r\n\16\n\u00b6\3\n\5\n\u00ba\n\n\3\n\3\n\3\13\3\13\3"+
		"\13\3\13\3\13\5\13\u00c3\n\13\3\f\3\f\3\f\3\f\5\f\u00c9\n\f\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3"+
		"\17\3\17\5\17\u00de\n\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\21\3\21\5\21\u00ec\n\21\5\21\u00ee\n\21\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\7\23\u00fd\n\23\f\23\16\23"+
		"\u0100\13\23\3\23\3\23\3\23\3\23\3\23\5\23\u0107\n\23\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u0125\n\25"+
		"\3\26\3\26\3\26\3\26\3\26\5\26\u012c\n\26\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\5\27\u0147\n\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u0158\n\27\f\27\16"+
		"\27\u015b\13\27\3\27\2\3,\30\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&(*,\2\t\3\2 !\4\2\23\23\36\36\3\2\"#\4\2\21\21\25\25\3\2\23\24\3\2\34"+
		"\35\3\2\26\33\2\u0185\2.\3\2\2\2\4\63\3\2\2\2\6[\3\2\2\2\b]\3\2\2\2\n"+
		"|\3\2\2\2\f\u0089\3\2\2\2\16\u0094\3\2\2\2\20\u00ae\3\2\2\2\22\u00b0\3"+
		"\2\2\2\24\u00bd\3\2\2\2\26\u00c4\3\2\2\2\30\u00ca\3\2\2\2\32\u00d2\3\2"+
		"\2\2\34\u00d8\3\2\2\2\36\u00e3\3\2\2\2 \u00ed\3\2\2\2\"\u00ef\3\2\2\2"+
		"$\u0106\3\2\2\2&\u0108\3\2\2\2(\u0124\3\2\2\2*\u012b\3\2\2\2,\u0146\3"+
		"\2\2\2./\5\4\3\2/\3\3\2\2\2\60\62\5\6\4\2\61\60\3\2\2\2\62\65\3\2\2\2"+
		"\63\61\3\2\2\2\63\64\3\2\2\2\64\5\3\2\2\2\65\63\3\2\2\2\668\5\b\5\2\67"+
		"9\7\13\2\28\67\3\2\2\289\3\2\2\29\\\3\2\2\2:<\5\n\6\2;=\7\13\2\2<;\3\2"+
		"\2\2<=\3\2\2\2=\\\3\2\2\2>@\5\f\7\2?A\7\13\2\2@?\3\2\2\2@A\3\2\2\2A\\"+
		"\3\2\2\2BD\5\16\b\2CE\7\13\2\2DC\3\2\2\2DE\3\2\2\2E\\\3\2\2\2F\\\5\20"+
		"\t\2G\\\5\22\n\2H\\\5\30\r\2I\\\5\32\16\2J\\\5\34\17\2KM\5 \21\2LN\7\13"+
		"\2\2ML\3\2\2\2MN\3\2\2\2N\\\3\2\2\2OQ\5\"\22\2PR\7\13\2\2QP\3\2\2\2QR"+
		"\3\2\2\2R\\\3\2\2\2SU\5(\25\2TV\7\13\2\2UT\3\2\2\2UV\3\2\2\2V\\\3\2\2"+
		"\2WY\5&\24\2XZ\7\13\2\2YX\3\2\2\2YZ\3\2\2\2Z\\\3\2\2\2[\66\3\2\2\2[:\3"+
		"\2\2\2[>\3\2\2\2[B\3\2\2\2[F\3\2\2\2[G\3\2\2\2[H\3\2\2\2[I\3\2\2\2[J\3"+
		"\2\2\2[K\3\2\2\2[O\3\2\2\2[S\3\2\2\2[W\3\2\2\2\\\7\3\2\2\2]^\7\37\2\2"+
		"^_\7\5\2\2_d\5,\27\2`a\7\3\2\2ac\5,\27\2b`\3\2\2\2cf\3\2\2\2db\3\2\2\2"+
		"de\3\2\2\2eg\3\2\2\2fd\3\2\2\2gi\7\4\2\2hj\7\13\2\2ih\3\2\2\2ij\3\2\2"+
		"\2j\t\3\2\2\2kl\7 \2\2lm\7B\2\2mn\7\n\2\2no\5*\26\2op\7\20\2\2pq\5,\27"+
		"\2q}\3\2\2\2rs\7 \2\2st\7B\2\2tu\7\20\2\2u}\5,\27\2vw\7 \2\2wx\7B\2\2"+
		"xy\7\n\2\2yz\5*\26\2z{\7\f\2\2{}\3\2\2\2|k\3\2\2\2|r\3\2\2\2|v\3\2\2\2"+
		"}\13\3\2\2\2~\177\7!\2\2\177\u0080\7B\2\2\u0080\u0081\7\n\2\2\u0081\u0082"+
		"\5*\26\2\u0082\u0083\7\20\2\2\u0083\u0084\5,\27\2\u0084\u008a\3\2\2\2"+
		"\u0085\u0086\7!\2\2\u0086\u0087\7B\2\2\u0087\u0088\7\20\2\2\u0088\u008a"+
		"\5,\27\2\u0089~\3\2\2\2\u0089\u0085\3\2\2\2\u008a\r\3\2\2\2\u008b\u008c"+
		"\7B\2\2\u008c\u008d\7\16\2\2\u008d\u0095\5,\27\2\u008e\u008f\7B\2\2\u008f"+
		"\u0090\7\17\2\2\u0090\u0095\5,\27\2\u0091\u0092\7B\2\2\u0092\u0093\7\20"+
		"\2\2\u0093\u0095\5,\27\2\u0094\u008b\3\2\2\2\u0094\u008e\3\2\2\2\u0094"+
		"\u0091\3\2\2\2\u0095\17\3\2\2\2\u0096\u0097\7$\2\2\u0097\u0098\5,\27\2"+
		"\u0098\u0099\7\6\2\2\u0099\u009a\5\4\3\2\u009a\u009b\7\7\2\2\u009b\u00af"+
		"\3\2\2\2\u009c\u009d\7$\2\2\u009d\u009e\5,\27\2\u009e\u009f\7\6\2\2\u009f"+
		"\u00a0\5\4\3\2\u00a0\u00a1\7\7\2\2\u00a1\u00a2\7%\2\2\u00a2\u00a3\7\6"+
		"\2\2\u00a3\u00a4\5\4\3\2\u00a4\u00a5\7\7\2\2\u00a5\u00af\3\2\2\2\u00a6"+
		"\u00a7\7$\2\2\u00a7\u00a8\5,\27\2\u00a8\u00a9\7\6\2\2\u00a9\u00aa\5\4"+
		"\3\2\u00aa\u00ab\7\7\2\2\u00ab\u00ac\7%\2\2\u00ac\u00ad\5\20\t\2\u00ad"+
		"\u00af\3\2\2\2\u00ae\u0096\3\2\2\2\u00ae\u009c\3\2\2\2\u00ae\u00a6\3\2"+
		"\2\2\u00af\21\3\2\2\2\u00b0\u00b1\7&\2\2\u00b1\u00b2\5,\27\2\u00b2\u00b4"+
		"\7\6\2\2\u00b3\u00b5\5\24\13\2\u00b4\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2"+
		"\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00ba"+
		"\5\26\f\2\u00b9\u00b8\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bb\3\2\2\2"+
		"\u00bb\u00bc\7\7\2\2\u00bc\23\3\2\2\2\u00bd\u00be\7\'\2\2\u00be\u00bf"+
		"\5,\27\2\u00bf\u00c0\7\n\2\2\u00c0\u00c2\5\4\3\2\u00c1\u00c3\7\60\2\2"+
		"\u00c2\u00c1\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\25\3\2\2\2\u00c4\u00c5"+
		"\7(\2\2\u00c5\u00c6\7\n\2\2\u00c6\u00c8\5\4\3\2\u00c7\u00c9\7\60\2\2\u00c8"+
		"\u00c7\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\27\3\2\2\2\u00ca\u00cb\7-\2\2"+
		"\u00cb\u00cc\5,\27\2\u00cc\u00cd\7%\2\2\u00cd\u00ce\7\6\2\2\u00ce\u00cf"+
		"\5\4\3\2\u00cf\u00d0\5 \21\2\u00d0\u00d1\7\7\2\2\u00d1\31\3\2\2\2\u00d2"+
		"\u00d3\7)\2\2\u00d3\u00d4\5,\27\2\u00d4\u00d5\7\6\2\2\u00d5\u00d6\5\4"+
		"\3\2\u00d6\u00d7\7\7\2\2\u00d7\33\3\2\2\2\u00d8\u00d9\7*\2\2\u00d9\u00da"+
		"\7B\2\2\u00da\u00dd\7+\2\2\u00db\u00de\5\36\20\2\u00dc\u00de\5,\27\2\u00dd"+
		"\u00db\3\2\2\2\u00dd\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00df\u00e0\7\6"+
		"\2\2\u00e0\u00e1\5\4\3\2\u00e1\u00e2\7\7\2\2\u00e2\35\3\2\2\2\u00e3\u00e4"+
		"\5,\27\2\u00e4\u00e5\7,\2\2\u00e5\u00e6\5,\27\2\u00e6\37\3\2\2\2\u00e7"+
		"\u00ee\7\60\2\2\u00e8\u00ee\7.\2\2\u00e9\u00eb\7/\2\2\u00ea\u00ec\5,\27"+
		"\2\u00eb\u00ea\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ee\3\2\2\2\u00ed\u00e7"+
		"\3\2\2\2\u00ed\u00e8\3\2\2\2\u00ed\u00e9\3\2\2\2\u00ee!\3\2\2\2\u00ef"+
		"\u00f0\t\2\2\2\u00f0\u00f1\7B\2\2\u00f1\u00f2\7\n\2\2\u00f2\u00f3\7\b"+
		"\2\2\u00f3\u00f4\5*\26\2\u00f4\u00f5\7\t\2\2\u00f5\u00f6\7\20\2\2\u00f6"+
		"\u00f7\5$\23\2\u00f7#\3\2\2\2\u00f8\u00f9\7\b\2\2\u00f9\u00fe\5,\27\2"+
		"\u00fa\u00fb\7\3\2\2\u00fb\u00fd\5,\27\2\u00fc\u00fa\3\2\2\2\u00fd\u0100"+
		"\3\2\2\2\u00fe\u00fc\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ff\u0101\3\2\2\2\u0100"+
		"\u00fe\3\2\2\2\u0101\u0102\7\t\2\2\u0102\u0107\3\2\2\2\u0103\u0104\7\b"+
		"\2\2\u0104\u0107\7\t\2\2\u0105\u0107\7B\2\2\u0106\u00f8\3\2\2\2\u0106"+
		"\u0103\3\2\2\2\u0106\u0105\3\2\2\2\u0107%\3\2\2\2\u0108\u0109\7B\2\2\u0109"+
		"\u010a\7\b\2\2\u010a\u010b\5,\27\2\u010b\u010c\7\t\2\2\u010c\u010d\7\20"+
		"\2\2\u010d\u010e\5,\27\2\u010e\'\3\2\2\2\u010f\u0110\7B\2\2\u0110\u0111"+
		"\7\r\2\2\u0111\u0112\7\62\2\2\u0112\u0113\7\5\2\2\u0113\u0114\5,\27\2"+
		"\u0114\u0115\7\4\2\2\u0115\u0125\3\2\2\2\u0116\u0117\7B\2\2\u0117\u0118"+
		"\7\r\2\2\u0118\u0119\7\63\2\2\u0119\u011a\7\5\2\2\u011a\u0125\7\4\2\2"+
		"\u011b\u011c\7B\2\2\u011c\u011d\7\r\2\2\u011d\u011e\7\64\2\2\u011e\u011f"+
		"\7\5\2\2\u011f\u0120\7\65\2\2\u0120\u0121\7\n\2\2\u0121\u0122\5,\27\2"+
		"\u0122\u0123\7\4\2\2\u0123\u0125\3\2\2\2\u0124\u010f\3\2\2\2\u0124\u0116"+
		"\3\2\2\2\u0124\u011b\3\2\2\2\u0125)\3\2\2\2\u0126\u012c\78\2\2\u0127\u012c"+
		"\79\2\2\u0128\u012c\7<\2\2\u0129\u012c\7:\2\2\u012a\u012c\7;\2\2\u012b"+
		"\u0126\3\2\2\2\u012b\u0127\3\2\2\2\u012b\u0128\3\2\2\2\u012b\u0129\3\2"+
		"\2\2\u012b\u012a\3\2\2\2\u012c+\3\2\2\2\u012d\u012e\b\27\1\2\u012e\u012f"+
		"\t\3\2\2\u012f\u0147\5,\27\23\u0130\u0147\t\4\2\2\u0131\u0147\7\61\2\2"+
		"\u0132\u0133\7B\2\2\u0133\u0134\7\b\2\2\u0134\u0135\5,\27\2\u0135\u0136"+
		"\7\t\2\2\u0136\u0147\3\2\2\2\u0137\u0138\7B\2\2\u0138\u0139\7\r\2\2\u0139"+
		"\u0147\7\66\2\2\u013a\u013b\7B\2\2\u013b\u013c\7\r\2\2\u013c\u0147\7\67"+
		"\2\2\u013d\u0147\7B\2\2\u013e\u0147\7?\2\2\u013f\u0147\7>\2\2\u0140\u0147"+
		"\7A\2\2\u0141\u0147\7@\2\2\u0142\u0143\7\5\2\2\u0143\u0144\5,\27\2\u0144"+
		"\u0145\7\4\2\2\u0145\u0147\3\2\2\2\u0146\u012d\3\2\2\2\u0146\u0130\3\2"+
		"\2\2\u0146\u0131\3\2\2\2\u0146\u0132\3\2\2\2\u0146\u0137\3\2\2\2\u0146"+
		"\u013a\3\2\2\2\u0146\u013d\3\2\2\2\u0146\u013e\3\2\2\2\u0146\u013f\3\2"+
		"\2\2\u0146\u0140\3\2\2\2\u0146\u0141\3\2\2\2\u0146\u0142\3\2\2\2\u0147"+
		"\u0159\3\2\2\2\u0148\u0149\f\22\2\2\u0149\u014a\t\5\2\2\u014a\u0158\5"+
		",\27\23\u014b\u014c\f\21\2\2\u014c\u014d\t\6\2\2\u014d\u0158\5,\27\22"+
		"\u014e\u014f\f\20\2\2\u014f\u0150\7\22\2\2\u0150\u0158\5,\27\21\u0151"+
		"\u0152\f\17\2\2\u0152\u0153\t\7\2\2\u0153\u0158\5,\27\20\u0154\u0155\f"+
		"\16\2\2\u0155\u0156\t\b\2\2\u0156\u0158\5,\27\17\u0157\u0148\3\2\2\2\u0157"+
		"\u014b\3\2\2\2\u0157\u014e\3\2\2\2\u0157\u0151\3\2\2\2\u0157\u0154\3\2"+
		"\2\2\u0158\u015b\3\2\2\2\u0159\u0157\3\2\2\2\u0159\u015a\3\2\2\2\u015a"+
		"-\3\2\2\2\u015b\u0159\3\2\2\2 \638<@DMQUY[di|\u0089\u0094\u00ae\u00b6"+
		"\u00b9\u00c2\u00c8\u00dd\u00eb\u00ed\u00fe\u0106\u0124\u012b\u0146\u0157"+
		"\u0159";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}