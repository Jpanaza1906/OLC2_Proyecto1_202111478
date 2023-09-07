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
	public static final int
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_print_sentencia = 3, 
		RULE_declaracion = 4, RULE_constante = 5, RULE_asignacion = 6, RULE_if_sentencia = 7, 
		RULE_switch_sentencia = 8, RULE_l_casos = 9, RULE_l_default = 10, RULE_guard_sentencia = 11, 
		RULE_while_sentencia = 12, RULE_for_sentencia = 13, RULE_rango_p = 14, 
		RULE_trans_sentencia = 15, RULE_dec_vector = 16, RULE_def_vector = 17, 
		RULE_asig_vector = 18, RULE_func_vector = 19, RULE_declaracion_matrices = 20, 
		RULE_tipo_matriz = 21, RULE_def_matriz = 22, RULE_listaval_mat = 23, RULE_listaval_mat2 = 24, 
		RULE_simple_mat = 25, RULE_asig_mat = 26, RULE_declaracion_metodo = 27, 
		RULE_declaracion_funcion = 28, RULE_l_parametros = 29, RULE_llamada_funciones = 30, 
		RULE_l_argumentos = 31, RULE_funciones_embebidas = 32, RULE_def_struct = 33, 
		RULE_l_sentencias_struct = 34, RULE_creacion_struct = 35, RULE_tipo = 36, 
		RULE_e = 37;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "print_sentencia", "declaracion", "constante", 
			"asignacion", "if_sentencia", "switch_sentencia", "l_casos", "l_default", 
			"guard_sentencia", "while_sentencia", "for_sentencia", "rango_p", "trans_sentencia", 
			"dec_vector", "def_vector", "asig_vector", "func_vector", "declaracion_matrices", 
			"tipo_matriz", "def_matriz", "listaval_mat", "listaval_mat2", "simple_mat", 
			"asig_mat", "declaracion_metodo", "declaracion_funcion", "l_parametros", 
			"llamada_funciones", "l_argumentos", "funciones_embebidas", "def_struct", 
			"l_sentencias_struct", "creacion_struct", "tipo", "e"
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
			setState(76);
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
			setState(81);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(78);
					sentencia();
					}
					} 
				}
				setState(83);
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
	public static class S_Declaracion_FuncionContext extends SentenciaContext {
		public Declaracion_funcionContext declaracion_funcion() {
			return getRuleContext(Declaracion_funcionContext.class,0);
		}
		public S_Declaracion_FuncionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Declaracion_MatrizContext extends SentenciaContext {
		public Declaracion_matricesContext declaracion_matrices() {
			return getRuleContext(Declaracion_matricesContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Declaracion_MatrizContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Asignacion_VectorContext extends SentenciaContext {
		public Asig_vectorContext asig_vector() {
			return getRuleContext(Asig_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Asignacion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Declaracion_MetodoContext extends SentenciaContext {
		public Declaracion_metodoContext declaracion_metodo() {
			return getRuleContext(Declaracion_metodoContext.class,0);
		}
		public S_Declaracion_MetodoContext(SentenciaContext ctx) { copyFrom(ctx); }
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
	public static class S_Asignacion_MatrizContext extends SentenciaContext {
		public Asig_matContext asig_mat() {
			return getRuleContext(Asig_matContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Asignacion_MatrizContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_Llamada_FuncionContext extends SentenciaContext {
		public Llamada_funcionesContext llamada_funciones() {
			return getRuleContext(Llamada_funcionesContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public S_Llamada_FuncionContext(SentenciaContext ctx) { copyFrom(ctx); }
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
			setState(135);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(84);
				print_sentencia();
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(85);
					match(PTCOMA);
					}
				}

				}
				break;
			case 2:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(88);
				declaracion();
				setState(90);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(89);
					match(PTCOMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new S_ConstanteContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				constante();
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(93);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(96);
				asignacion();
				setState(98);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(97);
					match(PTCOMA);
					}
				}

				}
				break;
			case 5:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(100);
				if_sentencia();
				}
				break;
			case 6:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(101);
				switch_sentencia();
				}
				break;
			case 7:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(102);
				guard_sentencia();
				}
				break;
			case 8:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(103);
				while_sentencia();
				}
				break;
			case 9:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(104);
				for_sentencia();
				}
				break;
			case 10:
				_localctx = new S_TransicionContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(105);
				trans_sentencia();
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(106);
					match(PTCOMA);
					}
				}

				}
				break;
			case 11:
				_localctx = new S_Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(109);
				dec_vector();
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(110);
					match(PTCOMA);
					}
				}

				}
				break;
			case 12:
				_localctx = new S_Funcion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(113);
				func_vector();
				setState(115);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(114);
					match(PTCOMA);
					}
				}

				}
				break;
			case 13:
				_localctx = new S_Asignacion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(117);
				asig_vector();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(118);
					match(PTCOMA);
					}
				}

				}
				break;
			case 14:
				_localctx = new S_Declaracion_MetodoContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(121);
				declaracion_metodo();
				}
				break;
			case 15:
				_localctx = new S_Declaracion_FuncionContext(_localctx);
				enterOuterAlt(_localctx, 15);
				{
				setState(122);
				declaracion_funcion();
				}
				break;
			case 16:
				_localctx = new S_Llamada_FuncionContext(_localctx);
				enterOuterAlt(_localctx, 16);
				{
				setState(123);
				llamada_funciones();
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(124);
					match(PTCOMA);
					}
				}

				}
				break;
			case 17:
				_localctx = new S_Declaracion_MatrizContext(_localctx);
				enterOuterAlt(_localctx, 17);
				{
				setState(127);
				declaracion_matrices();
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(128);
					match(PTCOMA);
					}
				}

				}
				break;
			case 18:
				_localctx = new S_Asignacion_MatrizContext(_localctx);
				enterOuterAlt(_localctx, 18);
				{
				setState(131);
				asig_mat();
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(132);
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
			setState(137);
			match(PRINT);
			setState(138);
			match(PARIZQ);
			setState(139);
			e(0);
			setState(144);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(140);
				match(T__0);
				setState(141);
				e(0);
				}
				}
				setState(146);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(147);
			match(PARDER);
			setState(149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(148);
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
			setState(168);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(151);
				match(VAR);
				setState(152);
				match(ID);
				setState(153);
				match(DOSPT);
				setState(154);
				tipo();
				setState(155);
				match(IGUAL);
				setState(156);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(158);
				match(VAR);
				setState(159);
				match(ID);
				setState(160);
				match(IGUAL);
				setState(161);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(162);
				match(VAR);
				setState(163);
				match(ID);
				setState(164);
				match(DOSPT);
				setState(165);
				tipo();
				setState(166);
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
			setState(181);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new Constante_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(170);
				match(LET);
				setState(171);
				match(ID);
				setState(172);
				match(DOSPT);
				setState(173);
				tipo();
				setState(174);
				match(IGUAL);
				setState(175);
				e(0);
				}
				break;
			case 2:
				_localctx = new Constante_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(177);
				match(LET);
				setState(178);
				match(ID);
				setState(179);
				match(IGUAL);
				setState(180);
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
			setState(192);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				_localctx = new SumAsgContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(183);
				match(ID);
				setState(184);
				match(MASIGUAL);
				setState(185);
				e(0);
				}
				break;
			case 2:
				_localctx = new ResAsgContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
				match(ID);
				setState(187);
				match(MENOSIGUAL);
				setState(188);
				e(0);
				}
				break;
			case 3:
				_localctx = new AsigContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(189);
				match(ID);
				setState(190);
				match(IGUAL);
				setState(191);
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
			setState(218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				_localctx = new If_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(194);
				match(IF);
				setState(195);
				e(0);
				setState(196);
				match(LLAVEIZQ);
				setState(197);
				l_sentencias();
				setState(198);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new If_ElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(200);
				match(IF);
				setState(201);
				e(0);
				setState(202);
				match(LLAVEIZQ);
				setState(203);
				l_sentencias();
				setState(204);
				match(LLAVEDER);
				setState(205);
				match(ELSE);
				setState(206);
				match(LLAVEIZQ);
				setState(207);
				l_sentencias();
				setState(208);
				match(LLAVEDER);
				}
				break;
			case 3:
				_localctx = new If_ElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(210);
				match(IF);
				setState(211);
				e(0);
				setState(212);
				match(LLAVEIZQ);
				setState(213);
				l_sentencias();
				setState(214);
				match(LLAVEDER);
				setState(215);
				match(ELSE);
				setState(216);
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
			setState(220);
			match(SWITCH);
			setState(221);
			e(0);
			setState(222);
			match(LLAVEIZQ);
			setState(224); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(223);
				l_casos();
				}
				}
				setState(226); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(228);
				l_default();
				}
			}

			setState(231);
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
			setState(233);
			match(CASE);
			setState(234);
			e(0);
			setState(235);
			match(DOSPT);
			setState(236);
			l_sentencias();
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(237);
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
			setState(240);
			match(DEFAULT);
			setState(241);
			match(DOSPT);
			setState(242);
			l_sentencias();
			setState(244);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(243);
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
			setState(246);
			match(GUARD);
			setState(247);
			e(0);
			setState(248);
			match(ELSE);
			setState(249);
			match(LLAVEIZQ);
			setState(250);
			l_sentencias();
			setState(251);
			trans_sentencia();
			setState(252);
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
			setState(254);
			match(WHILE);
			setState(255);
			e(0);
			setState(256);
			match(LLAVEIZQ);
			setState(257);
			l_sentencias();
			setState(258);
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
		public Token id;
		public TerminalNode FOR() { return getToken(Tswift_GrammarParser.FOR, 0); }
		public TerminalNode IN() { return getToken(Tswift_GrammarParser.IN, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode GUIONBAJO() { return getToken(Tswift_GrammarParser.GUIONBAJO, 0); }
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
		int _la;
		try {
			_localctx = new ForContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			match(FOR);
			setState(261);
			((ForContext)_localctx).id = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==GUIONBAJO || _la==ID) ) {
				((ForContext)_localctx).id = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(262);
			match(IN);
			setState(265);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				{
				setState(263);
				rango_p();
				}
				break;
			case 2:
				{
				setState(264);
				e(0);
				}
				break;
			}
			setState(267);
			match(LLAVEIZQ);
			setState(268);
			l_sentencias();
			setState(269);
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
			setState(271);
			e(0);
			setState(272);
			match(RANGO);
			setState(273);
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
			setState(281);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(275);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(276);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(277);
				match(RETURN);
				setState(279);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
				case 1:
					{
					setState(278);
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
			setState(283);
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
			setState(284);
			match(ID);
			setState(285);
			match(DOSPT);
			setState(286);
			match(CORCHETEIZQ);
			setState(287);
			tipo();
			setState(288);
			match(CORCHETEDER);
			setState(289);
			match(IGUAL);
			setState(290);
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
			setState(306);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				_localctx = new Def_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(292);
				match(CORCHETEIZQ);
				setState(293);
				e(0);
				setState(298);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(294);
					match(T__0);
					setState(295);
					e(0);
					}
					}
					setState(300);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(301);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Vector_VacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(303);
				match(CORCHETEIZQ);
				setState(304);
				match(CORCHETEDER);
				}
				break;
			case 3:
				_localctx = new Def_Vector_IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(305);
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
	public static class ResAsg_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public TerminalNode MENOSIGUAL() { return getToken(Tswift_GrammarParser.MENOSIGUAL, 0); }
		public ResAsg_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}
	public static class SumAsg_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public TerminalNode MASIGUAL() { return getToken(Tswift_GrammarParser.MASIGUAL, 0); }
		public SumAsg_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Asig_vectorContext asig_vector() throws RecognitionException {
		Asig_vectorContext _localctx = new Asig_vectorContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_asig_vector);
		try {
			setState(329);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new Asig_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(308);
				match(ID);
				setState(309);
				match(CORCHETEIZQ);
				setState(310);
				e(0);
				setState(311);
				match(CORCHETEDER);
				setState(312);
				match(IGUAL);
				setState(313);
				e(0);
				}
				break;
			case 2:
				_localctx = new SumAsg_VectorContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(315);
				match(ID);
				setState(316);
				match(CORCHETEIZQ);
				setState(317);
				e(0);
				setState(318);
				match(CORCHETEDER);
				setState(319);
				match(MASIGUAL);
				setState(320);
				e(0);
				}
				break;
			case 3:
				_localctx = new ResAsg_VectorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(322);
				match(ID);
				setState(323);
				match(CORCHETEIZQ);
				setState(324);
				e(0);
				setState(325);
				match(CORCHETEDER);
				setState(326);
				match(MENOSIGUAL);
				setState(327);
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
			setState(352);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				_localctx = new Func_Vector_AppendContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(331);
				match(ID);
				setState(332);
				match(PUNTO);
				setState(333);
				match(APPEND);
				setState(334);
				match(PARIZQ);
				setState(335);
				e(0);
				setState(336);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Func_Vector_RemoveLastContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(338);
				match(ID);
				setState(339);
				match(PUNTO);
				setState(340);
				match(REMOVELAST);
				setState(341);
				match(PARIZQ);
				setState(342);
				match(PARDER);
				}
				break;
			case 3:
				_localctx = new Func_Vector_RemoveContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(343);
				match(ID);
				setState(344);
				match(PUNTO);
				setState(345);
				match(REMOVE);
				setState(346);
				match(PARIZQ);
				setState(347);
				match(AT);
				setState(348);
				match(DOSPT);
				setState(349);
				e(0);
				setState(350);
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

	public static class Declaracion_matricesContext extends ParserRuleContext {
		public Declaracion_matricesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_matrices; }
	 
		public Declaracion_matricesContext() { }
		public void copyFrom(Declaracion_matricesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declaracion_MatrizContext extends Declaracion_matricesContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public Def_matrizContext def_matriz() {
			return getRuleContext(Def_matrizContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public Tipo_matrizContext tipo_matriz() {
			return getRuleContext(Tipo_matrizContext.class,0);
		}
		public Declaracion_MatrizContext(Declaracion_matricesContext ctx) { copyFrom(ctx); }
	}

	public final Declaracion_matricesContext declaracion_matrices() throws RecognitionException {
		Declaracion_matricesContext _localctx = new Declaracion_matricesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_declaracion_matrices);
		int _la;
		try {
			_localctx = new Declaracion_MatrizContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(354);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(355);
			match(ID);
			setState(358);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPT) {
				{
				setState(356);
				match(DOSPT);
				setState(357);
				tipo_matriz();
				}
			}

			setState(360);
			match(IGUAL);
			setState(361);
			def_matriz();
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

	public static class Tipo_matrizContext extends ParserRuleContext {
		public Tipo_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_matriz; }
	 
		public Tipo_matrizContext() { }
		public void copyFrom(Tipo_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Tipo_MatrizContext extends Tipo_matrizContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public Tipo_matrizContext tipo_matriz() {
			return getRuleContext(Tipo_matrizContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Tipo_MatrizContext(Tipo_matrizContext ctx) { copyFrom(ctx); }
	}
	public static class Tipo_Matriz_SimpleContext extends Tipo_matrizContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Tipo_Matriz_SimpleContext(Tipo_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Tipo_matrizContext tipo_matriz() throws RecognitionException {
		Tipo_matrizContext _localctx = new Tipo_matrizContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_tipo_matriz);
		try {
			setState(371);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new Tipo_MatrizContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(363);
				match(CORCHETEIZQ);
				setState(364);
				tipo_matriz();
				setState(365);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Tipo_Matriz_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(367);
				match(CORCHETEIZQ);
				setState(368);
				tipo();
				setState(369);
				match(CORCHETEDER);
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

	public static class Def_matrizContext extends ParserRuleContext {
		public Def_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_def_matriz; }
	 
		public Def_matrizContext() { }
		public void copyFrom(Def_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Def_MatrizContext extends Def_matrizContext {
		public Listaval_matContext listaval_mat() {
			return getRuleContext(Listaval_matContext.class,0);
		}
		public Def_MatrizContext(Def_matrizContext ctx) { copyFrom(ctx); }
	}
	public static class Def_Matriz_SimpleContext extends Def_matrizContext {
		public Simple_matContext simple_mat() {
			return getRuleContext(Simple_matContext.class,0);
		}
		public Def_Matriz_SimpleContext(Def_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Def_matrizContext def_matriz() throws RecognitionException {
		Def_matrizContext _localctx = new Def_matrizContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_def_matriz);
		try {
			setState(375);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				_localctx = new Def_MatrizContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(373);
				listaval_mat();
				}
				break;
			case 2:
				_localctx = new Def_Matriz_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(374);
				simple_mat();
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

	public static class Listaval_matContext extends ParserRuleContext {
		public Listaval_matContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaval_mat; }
	 
		public Listaval_matContext() { }
		public void copyFrom(Listaval_matContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ListaCompletaValContext extends Listaval_matContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public Listaval_mat2Context listaval_mat2() {
			return getRuleContext(Listaval_mat2Context.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public ListaCompletaValContext(Listaval_matContext ctx) { copyFrom(ctx); }
	}

	public final Listaval_matContext listaval_mat() throws RecognitionException {
		Listaval_matContext _localctx = new Listaval_matContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_listaval_mat);
		try {
			_localctx = new ListaCompletaValContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(377);
			match(CORCHETEIZQ);
			setState(378);
			listaval_mat2(0);
			setState(379);
			match(CORCHETEDER);
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

	public static class Listaval_mat2Context extends ParserRuleContext {
		public Listaval_mat2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaval_mat2; }
	 
		public Listaval_mat2Context() { }
		public void copyFrom(Listaval_mat2Context ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ListaValorSigContext extends Listaval_mat2Context {
		public Listaval_matContext listaval_mat() {
			return getRuleContext(Listaval_matContext.class,0);
		}
		public ListaValorSigContext(Listaval_mat2Context ctx) { copyFrom(ctx); }
	}
	public static class ListaValoresHermanosContext extends Listaval_mat2Context {
		public Listaval_mat2Context listaval_mat2() {
			return getRuleContext(Listaval_mat2Context.class,0);
		}
		public Listaval_matContext listaval_mat() {
			return getRuleContext(Listaval_matContext.class,0);
		}
		public ListaValoresHermanosContext(Listaval_mat2Context ctx) { copyFrom(ctx); }
	}
	public static class ListaExprContext extends Listaval_mat2Context {
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public ListaExprContext(Listaval_mat2Context ctx) { copyFrom(ctx); }
	}

	public final Listaval_mat2Context listaval_mat2() throws RecognitionException {
		return listaval_mat2(0);
	}

	private Listaval_mat2Context listaval_mat2(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Listaval_mat2Context _localctx = new Listaval_mat2Context(_ctx, _parentState);
		Listaval_mat2Context _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_listaval_mat2, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(393);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CORCHETEIZQ:
				{
				_localctx = new ListaValorSigContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(382);
				listaval_mat();
				}
				break;
			case PARIZQ:
			case MENOS:
			case NOT:
			case TRUE:
			case FALSE:
			case NIL:
			case ATOI:
			case IOTA:
			case INT:
			case FLOAT:
			case STRING:
			case ENTERO:
			case DECIMAL:
			case CARACTER:
			case CADENA:
			case ID:
				{
				_localctx = new ListaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(383);
				e(0);
				setState(384);
				match(T__0);
				setState(385);
				e(0);
				setState(390);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(386);
						match(T__0);
						setState(387);
						e(0);
						}
						} 
					}
					setState(392);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(400);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListaValoresHermanosContext(new Listaval_mat2Context(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_listaval_mat2);
					setState(395);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(396);
					match(T__0);
					setState(397);
					listaval_mat();
					}
					} 
				}
				setState(402);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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

	public static class Simple_matContext extends ParserRuleContext {
		public Simple_matContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simple_mat; }
	 
		public Simple_matContext() { }
		public void copyFrom(Simple_matContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Simple_MatContext extends Simple_matContext {
		public Tipo_matrizContext tipo_matriz() {
			return getRuleContext(Tipo_matrizContext.class,0);
		}
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode REPEATING() { return getToken(Tswift_GrammarParser.REPEATING, 0); }
		public List<TerminalNode> DOSPT() { return getTokens(Tswift_GrammarParser.DOSPT); }
		public TerminalNode DOSPT(int i) {
			return getToken(Tswift_GrammarParser.DOSPT, i);
		}
		public Simple_matContext simple_mat() {
			return getRuleContext(Simple_matContext.class,0);
		}
		public TerminalNode COUNT() { return getToken(Tswift_GrammarParser.COUNT, 0); }
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarParser.ENTERO, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Simple_MatContext(Simple_matContext ctx) { copyFrom(ctx); }
	}
	public static class Simple_Mat_ExprContext extends Simple_matContext {
		public Tipo_matrizContext tipo_matriz() {
			return getRuleContext(Tipo_matrizContext.class,0);
		}
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode REPEATING() { return getToken(Tswift_GrammarParser.REPEATING, 0); }
		public List<TerminalNode> DOSPT() { return getTokens(Tswift_GrammarParser.DOSPT); }
		public TerminalNode DOSPT(int i) {
			return getToken(Tswift_GrammarParser.DOSPT, i);
		}
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode COUNT() { return getToken(Tswift_GrammarParser.COUNT, 0); }
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarParser.ENTERO, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Simple_Mat_ExprContext(Simple_matContext ctx) { copyFrom(ctx); }
	}

	public final Simple_matContext simple_mat() throws RecognitionException {
		Simple_matContext _localctx = new Simple_matContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_simple_mat);
		try {
			setState(425);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				_localctx = new Simple_MatContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(403);
				tipo_matriz();
				setState(404);
				match(PARIZQ);
				setState(405);
				match(REPEATING);
				setState(406);
				match(DOSPT);
				setState(407);
				simple_mat();
				setState(408);
				match(T__0);
				setState(409);
				match(COUNT);
				setState(410);
				match(DOSPT);
				setState(411);
				match(ENTERO);
				setState(412);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Simple_Mat_ExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(414);
				tipo_matriz();
				setState(415);
				match(PARIZQ);
				setState(416);
				match(REPEATING);
				setState(417);
				match(DOSPT);
				setState(418);
				e(0);
				setState(419);
				match(T__0);
				setState(420);
				match(COUNT);
				setState(421);
				match(DOSPT);
				setState(422);
				match(ENTERO);
				setState(423);
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

	public static class Asig_matContext extends ParserRuleContext {
		public Asig_matContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asig_mat; }
	 
		public Asig_matContext() { }
		public void copyFrom(Asig_matContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Asig_MatContext extends Asig_matContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public List<TerminalNode> CORCHETEIZQ() { return getTokens(Tswift_GrammarParser.CORCHETEIZQ); }
		public TerminalNode CORCHETEIZQ(int i) {
			return getToken(Tswift_GrammarParser.CORCHETEIZQ, i);
		}
		public List<TerminalNode> ENTERO() { return getTokens(Tswift_GrammarParser.ENTERO); }
		public TerminalNode ENTERO(int i) {
			return getToken(Tswift_GrammarParser.ENTERO, i);
		}
		public List<TerminalNode> CORCHETEDER() { return getTokens(Tswift_GrammarParser.CORCHETEDER); }
		public TerminalNode CORCHETEDER(int i) {
			return getToken(Tswift_GrammarParser.CORCHETEDER, i);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Asig_MatContext(Asig_matContext ctx) { copyFrom(ctx); }
	}

	public final Asig_matContext asig_mat() throws RecognitionException {
		Asig_matContext _localctx = new Asig_matContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_asig_mat);
		int _la;
		try {
			_localctx = new Asig_MatContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(427);
			match(ID);
			setState(428);
			match(CORCHETEIZQ);
			setState(429);
			match(ENTERO);
			setState(430);
			match(CORCHETEDER);
			setState(431);
			match(CORCHETEIZQ);
			setState(432);
			match(ENTERO);
			setState(433);
			match(CORCHETEDER);
			setState(439);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CORCHETEIZQ) {
				{
				{
				setState(434);
				match(CORCHETEIZQ);
				setState(435);
				match(ENTERO);
				setState(436);
				match(CORCHETEDER);
				}
				}
				setState(441);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(442);
			match(IGUAL);
			setState(443);
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

	public static class Declaracion_metodoContext extends ParserRuleContext {
		public Declaracion_metodoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_metodo; }
	 
		public Declaracion_metodoContext() { }
		public void copyFrom(Declaracion_metodoContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declaracion_MetodoContext extends Declaracion_metodoContext {
		public TerminalNode FUNC() { return getToken(Tswift_GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public List<L_parametrosContext> l_parametros() {
			return getRuleContexts(L_parametrosContext.class);
		}
		public L_parametrosContext l_parametros(int i) {
			return getRuleContext(L_parametrosContext.class,i);
		}
		public Declaracion_MetodoContext(Declaracion_metodoContext ctx) { copyFrom(ctx); }
	}

	public final Declaracion_metodoContext declaracion_metodo() throws RecognitionException {
		Declaracion_metodoContext _localctx = new Declaracion_metodoContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_declaracion_metodo);
		int _la;
		try {
			_localctx = new Declaracion_MetodoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(445);
			match(FUNC);
			setState(446);
			match(ID);
			setState(447);
			match(PARIZQ);
			setState(451);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==GUIONBAJO || _la==ID) {
				{
				{
				setState(448);
				l_parametros();
				}
				}
				setState(453);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(454);
			match(PARDER);
			setState(455);
			match(LLAVEIZQ);
			setState(456);
			l_sentencias();
			setState(457);
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

	public static class Declaracion_funcionContext extends ParserRuleContext {
		public Declaracion_funcionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_funcion; }
	 
		public Declaracion_funcionContext() { }
		public void copyFrom(Declaracion_funcionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declaracion_FuncionContext extends Declaracion_funcionContext {
		public TerminalNode FUNC() { return getToken(Tswift_GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode MENOS() { return getToken(Tswift_GrammarParser.MENOS, 0); }
		public TerminalNode MAYOR() { return getToken(Tswift_GrammarParser.MAYOR, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public List<L_parametrosContext> l_parametros() {
			return getRuleContexts(L_parametrosContext.class);
		}
		public L_parametrosContext l_parametros(int i) {
			return getRuleContext(L_parametrosContext.class,i);
		}
		public Declaracion_FuncionContext(Declaracion_funcionContext ctx) { copyFrom(ctx); }
	}

	public final Declaracion_funcionContext declaracion_funcion() throws RecognitionException {
		Declaracion_funcionContext _localctx = new Declaracion_funcionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_declaracion_funcion);
		int _la;
		try {
			_localctx = new Declaracion_FuncionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(459);
			match(FUNC);
			setState(460);
			match(ID);
			setState(461);
			match(PARIZQ);
			setState(465);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==GUIONBAJO || _la==ID) {
				{
				{
				setState(462);
				l_parametros();
				}
				}
				setState(467);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(468);
			match(PARDER);
			setState(469);
			match(MENOS);
			setState(470);
			match(MAYOR);
			setState(471);
			tipo();
			setState(472);
			match(LLAVEIZQ);
			setState(473);
			l_sentencias();
			setState(474);
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

	public static class L_parametrosContext extends ParserRuleContext {
		public L_parametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_parametros; }
	 
		public L_parametrosContext() { }
		public void copyFrom(L_parametrosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class L_ParametrosContext extends L_parametrosContext {
		public Token prim;
		public List<TerminalNode> ID() { return getTokens(Tswift_GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Tswift_GrammarParser.ID, i);
		}
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode INOUT() { return getToken(Tswift_GrammarParser.INOUT, 0); }
		public TerminalNode GUIONBAJO() { return getToken(Tswift_GrammarParser.GUIONBAJO, 0); }
		public L_ParametrosContext(L_parametrosContext ctx) { copyFrom(ctx); }
	}

	public final L_parametrosContext l_parametros() throws RecognitionException {
		L_parametrosContext _localctx = new L_parametrosContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_l_parametros);
		int _la;
		try {
			_localctx = new L_ParametrosContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(477);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				{
				setState(476);
				((L_ParametrosContext)_localctx).prim = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==GUIONBAJO || _la==ID) ) {
					((L_ParametrosContext)_localctx).prim = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			}
			setState(479);
			match(ID);
			setState(480);
			match(DOSPT);
			setState(482);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT) {
				{
				setState(481);
				match(INOUT);
				}
			}

			setState(484);
			tipo();
			setState(486);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(485);
				match(T__0);
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

	public static class Llamada_funcionesContext extends ParserRuleContext {
		public Llamada_funcionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamada_funciones; }
	 
		public Llamada_funcionesContext() { }
		public void copyFrom(Llamada_funcionesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Llamada_FuncionContext extends Llamada_funcionesContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public List<L_argumentosContext> l_argumentos() {
			return getRuleContexts(L_argumentosContext.class);
		}
		public L_argumentosContext l_argumentos(int i) {
			return getRuleContext(L_argumentosContext.class,i);
		}
		public Llamada_FuncionContext(Llamada_funcionesContext ctx) { copyFrom(ctx); }
	}

	public final Llamada_funcionesContext llamada_funciones() throws RecognitionException {
		Llamada_funcionesContext _localctx = new Llamada_funcionesContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_llamada_funciones);
		int _la;
		try {
			_localctx = new Llamada_FuncionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
			match(ID);
			setState(489);
			match(PARIZQ);
			setState(493);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PARIZQ) | (1L << DIR) | (1L << MENOS) | (1L << NOT) | (1L << TRUE) | (1L << FALSE) | (1L << NIL) | (1L << ATOI) | (1L << IOTA) | (1L << INT))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (FLOAT - 64)) | (1L << (STRING - 64)) | (1L << (ENTERO - 64)) | (1L << (DECIMAL - 64)) | (1L << (CARACTER - 64)) | (1L << (CADENA - 64)) | (1L << (ID - 64)))) != 0)) {
				{
				{
				setState(490);
				l_argumentos();
				}
				}
				setState(495);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(496);
			match(PARDER);
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

	public static class L_argumentosContext extends ParserRuleContext {
		public L_argumentosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_argumentos; }
	 
		public L_argumentosContext() { }
		public void copyFrom(L_argumentosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class L_ArgumentosContext extends L_argumentosContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TerminalNode DIR() { return getToken(Tswift_GrammarParser.DIR, 0); }
		public L_ArgumentosContext(L_argumentosContext ctx) { copyFrom(ctx); }
	}

	public final L_argumentosContext l_argumentos() throws RecognitionException {
		L_argumentosContext _localctx = new L_argumentosContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_l_argumentos);
		int _la;
		try {
			_localctx = new L_ArgumentosContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(500);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				{
				setState(498);
				match(ID);
				setState(499);
				match(DOSPT);
				}
				break;
			}
			setState(503);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DIR) {
				{
				setState(502);
				match(DIR);
				}
			}

			setState(505);
			e(0);
			setState(507);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(506);
				match(T__0);
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

	public static class Funciones_embebidasContext extends ParserRuleContext {
		public Funciones_embebidasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciones_embebidas; }
	 
		public Funciones_embebidasContext() { }
		public void copyFrom(Funciones_embebidasContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Func_StringContext extends Funciones_embebidasContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode STRING() { return getToken(Tswift_GrammarParser.STRING, 0); }
		public TerminalNode IOTA() { return getToken(Tswift_GrammarParser.IOTA, 0); }
		public Func_StringContext(Funciones_embebidasContext ctx) { copyFrom(ctx); }
	}
	public static class Func_IntContext extends Funciones_embebidasContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode INT() { return getToken(Tswift_GrammarParser.INT, 0); }
		public TerminalNode ATOI() { return getToken(Tswift_GrammarParser.ATOI, 0); }
		public Func_IntContext(Funciones_embebidasContext ctx) { copyFrom(ctx); }
	}
	public static class Func_FloatContext extends Funciones_embebidasContext {
		public TerminalNode FLOAT() { return getToken(Tswift_GrammarParser.FLOAT, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public Func_FloatContext(Funciones_embebidasContext ctx) { copyFrom(ctx); }
	}

	public final Funciones_embebidasContext funciones_embebidas() throws RecognitionException {
		Funciones_embebidasContext _localctx = new Funciones_embebidasContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_funciones_embebidas);
		int _la;
		try {
			setState(524);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ATOI:
			case INT:
				_localctx = new Func_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(509);
				_la = _input.LA(1);
				if ( !(_la==ATOI || _la==INT) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(510);
				match(PARIZQ);
				setState(511);
				e(0);
				setState(512);
				match(PARDER);
				}
				break;
			case FLOAT:
				_localctx = new Func_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(514);
				match(FLOAT);
				setState(515);
				match(PARIZQ);
				setState(516);
				e(0);
				setState(517);
				match(PARDER);
				}
				break;
			case IOTA:
			case STRING:
				_localctx = new Func_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(519);
				_la = _input.LA(1);
				if ( !(_la==IOTA || _la==STRING) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(520);
				match(PARIZQ);
				setState(521);
				e(0);
				setState(522);
				match(PARDER);
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

	public static class Def_structContext extends ParserRuleContext {
		public Def_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_def_struct; }
	 
		public Def_structContext() { }
		public void copyFrom(Def_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Def_StructContext extends Def_structContext {
		public TerminalNode STRUCT() { return getToken(Tswift_GrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarParser.LLAVEDER, 0); }
		public List<L_sentencias_structContext> l_sentencias_struct() {
			return getRuleContexts(L_sentencias_structContext.class);
		}
		public L_sentencias_structContext l_sentencias_struct(int i) {
			return getRuleContext(L_sentencias_structContext.class,i);
		}
		public Def_StructContext(Def_structContext ctx) { copyFrom(ctx); }
	}

	public final Def_structContext def_struct() throws RecognitionException {
		Def_structContext _localctx = new Def_structContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_def_struct);
		int _la;
		try {
			_localctx = new Def_StructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			match(STRUCT);
			setState(527);
			match(ID);
			setState(528);
			match(LLAVEIZQ);
			setState(532);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << FUNC) | (1L << MUTATING))) != 0)) {
				{
				{
				setState(529);
				l_sentencias_struct();
				}
				}
				setState(534);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(535);
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

	public static class L_sentencias_structContext extends ParserRuleContext {
		public L_sentencias_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_sentencias_struct; }
	 
		public L_sentencias_structContext() { }
		public void copyFrom(L_sentencias_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class L_FuncionesContext extends L_sentencias_structContext {
		public Declaracion_funcionContext declaracion_funcion() {
			return getRuleContext(Declaracion_funcionContext.class,0);
		}
		public TerminalNode MUTATING() { return getToken(Tswift_GrammarParser.MUTATING, 0); }
		public L_FuncionesContext(L_sentencias_structContext ctx) { copyFrom(ctx); }
	}
	public static class L_AtributosContext extends L_sentencias_structContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarParser.PTCOMA, 0); }
		public L_AtributosContext(L_sentencias_structContext ctx) { copyFrom(ctx); }
	}

	public final L_sentencias_structContext l_sentencias_struct() throws RecognitionException {
		L_sentencias_structContext _localctx = new L_sentencias_structContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_l_sentencias_struct);
		int _la;
		try {
			setState(554);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
			case LET:
				_localctx = new L_AtributosContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(537);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(538);
				match(ID);
				setState(541);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPT) {
					{
					setState(539);
					match(DOSPT);
					setState(540);
					tipo();
					}
				}

				setState(545);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IGUAL) {
					{
					setState(543);
					match(IGUAL);
					setState(544);
					e(0);
					}
				}

				setState(548);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(547);
					match(PTCOMA);
					}
				}

				}
				break;
			case FUNC:
			case MUTATING:
				_localctx = new L_FuncionesContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(551);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MUTATING) {
					{
					setState(550);
					match(MUTATING);
					}
				}

				setState(553);
				declaracion_funcion();
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

	public static class Creacion_structContext extends ParserRuleContext {
		public Creacion_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_creacion_struct; }
	 
		public Creacion_structContext() { }
		public void copyFrom(Creacion_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Creacion_Struct_SimpleContext extends Creacion_structContext {
		public List<TerminalNode> ID() { return getTokens(Tswift_GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Tswift_GrammarParser.ID, i);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public Creacion_Struct_SimpleContext(Creacion_structContext ctx) { copyFrom(ctx); }
	}
	public static class Creacion_StructContext extends Creacion_structContext {
		public List<TerminalNode> ID() { return getTokens(Tswift_GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(Tswift_GrammarParser.ID, i);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarParser.IGUAL, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode VAR() { return getToken(Tswift_GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarParser.LET, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarParser.DOSPT, 0); }
		public L_argumentosContext l_argumentos() {
			return getRuleContext(L_argumentosContext.class,0);
		}
		public Creacion_StructContext(Creacion_structContext ctx) { copyFrom(ctx); }
	}

	public final Creacion_structContext creacion_struct() throws RecognitionException {
		Creacion_structContext _localctx = new Creacion_structContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_creacion_struct);
		int _la;
		try {
			setState(577);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
			case 1:
				_localctx = new Creacion_StructContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(556);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(557);
				match(ID);
				setState(560);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPT) {
					{
					setState(558);
					match(DOSPT);
					setState(559);
					match(ID);
					}
				}

				setState(562);
				match(IGUAL);
				setState(563);
				match(ID);
				setState(564);
				match(PARIZQ);
				setState(566);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PARIZQ) | (1L << DIR) | (1L << MENOS) | (1L << NOT) | (1L << TRUE) | (1L << FALSE) | (1L << NIL) | (1L << ATOI) | (1L << IOTA) | (1L << INT))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (FLOAT - 64)) | (1L << (STRING - 64)) | (1L << (ENTERO - 64)) | (1L << (DECIMAL - 64)) | (1L << (CARACTER - 64)) | (1L << (CADENA - 64)) | (1L << (ID - 64)))) != 0)) {
					{
					setState(565);
					l_argumentos();
					}
				}

				setState(568);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Creacion_Struct_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(569);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(570);
				match(ID);
				setState(573);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPT) {
					{
					setState(571);
					match(DOSPT);
					setState(572);
					match(ID);
					}
				}

				setState(575);
				match(IGUAL);
				setState(576);
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
	public static class Tipo_VectorContext extends TipoContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarParser.CORCHETEDER, 0); }
		public Tipo_VectorContext(TipoContext ctx) { copyFrom(ctx); }
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
		enterRule(_localctx, 72, RULE_tipo);
		try {
			setState(588);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(579);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(580);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(581);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(582);
				match(BOOL);
				}
				break;
			case CHARACTER:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(583);
				match(CHARACTER);
				}
				break;
			case CORCHETEIZQ:
				_localctx = new Tipo_VectorContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(584);
				match(CORCHETEIZQ);
				setState(585);
				tipo();
				setState(586);
				match(CORCHETEDER);
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
	public static class Expr_Funciones_EmbebidasContext extends EContext {
		public Funciones_embebidasContext funciones_embebidas() {
			return getRuleContext(Funciones_embebidasContext.class,0);
		}
		public Expr_Funciones_EmbebidasContext(EContext ctx) { copyFrom(ctx); }
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
	public static class Expr_MatrizContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public List<TerminalNode> CORCHETEIZQ() { return getTokens(Tswift_GrammarParser.CORCHETEIZQ); }
		public TerminalNode CORCHETEIZQ(int i) {
			return getToken(Tswift_GrammarParser.CORCHETEIZQ, i);
		}
		public List<TerminalNode> ENTERO() { return getTokens(Tswift_GrammarParser.ENTERO); }
		public TerminalNode ENTERO(int i) {
			return getToken(Tswift_GrammarParser.ENTERO, i);
		}
		public List<TerminalNode> CORCHETEDER() { return getTokens(Tswift_GrammarParser.CORCHETEDER); }
		public TerminalNode CORCHETEDER(int i) {
			return getToken(Tswift_GrammarParser.CORCHETEDER, i);
		}
		public Expr_MatrizContext(EContext ctx) { copyFrom(ctx); }
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
	public static class Expr_Llamada_FuncionContext extends EContext {
		public Llamada_funcionesContext llamada_funciones() {
			return getRuleContext(Llamada_funcionesContext.class,0);
		}
		public Expr_Llamada_FuncionContext(EContext ctx) { copyFrom(ctx); }
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
		int _startState = 74;
		enterRecursionRule(_localctx, 74, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(632);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,60,_ctx) ) {
			case 1:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(591);
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
				setState(592);
				e(20);
				}
				break;
			case 2:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(593);
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
				_localctx = new Expr_Funciones_EmbebidasContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(594);
				funciones_embebidas();
				}
				break;
			case 4:
				{
				_localctx = new Expr_Llamada_FuncionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(595);
				llamada_funciones();
				}
				break;
			case 5:
				{
				_localctx = new Expr_MatrizContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(596);
				match(ID);
				setState(597);
				match(CORCHETEIZQ);
				setState(598);
				match(ENTERO);
				setState(599);
				match(CORCHETEDER);
				setState(600);
				match(CORCHETEIZQ);
				setState(601);
				match(ENTERO);
				setState(602);
				match(CORCHETEDER);
				setState(608);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,59,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(603);
						match(CORCHETEIZQ);
						setState(604);
						match(ENTERO);
						setState(605);
						match(CORCHETEDER);
						}
						} 
					}
					setState(610);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,59,_ctx);
				}
				}
				break;
			case 6:
				{
				_localctx = new Expr_VectorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(611);
				match(ID);
				setState(612);
				match(CORCHETEIZQ);
				setState(613);
				e(0);
				setState(614);
				match(CORCHETEDER);
				}
				break;
			case 7:
				{
				_localctx = new Func_Vector_isEmptyContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(616);
				match(ID);
				setState(617);
				match(PUNTO);
				setState(618);
				match(ISEMPTY);
				}
				break;
			case 8:
				{
				_localctx = new Func_Vector_CountContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(619);
				match(ID);
				setState(620);
				match(PUNTO);
				setState(621);
				match(COUNT);
				}
				break;
			case 9:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(622);
				match(ID);
				}
				break;
			case 10:
				{
				_localctx = new Expr_NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(623);
				match(NIL);
				}
				break;
			case 11:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(624);
				match(DECIMAL);
				}
				break;
			case 12:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(625);
				match(ENTERO);
				}
				break;
			case 13:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(626);
				match(CADENA);
				}
				break;
			case 14:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(627);
				match(CARACTER);
				}
				break;
			case 15:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(628);
				match(PARIZQ);
				setState(629);
				e(0);
				setState(630);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(651);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(649);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,61,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(634);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(635);
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
						setState(636);
						e(20);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(637);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(638);
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
						setState(639);
						e(19);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(640);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(641);
						match(MOD);
						setState(642);
						e(18);
						}
						break;
					case 4:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(643);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(644);
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
						setState(645);
						e(17);
						}
						break;
					case 5:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(646);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(647);
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
						setState(648);
						e(16);
						}
						break;
					}
					} 
				}
				setState(653);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
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
		case 24:
			return listaval_mat2_sempred((Listaval_mat2Context)_localctx, predIndex);
		case 37:
			return e_sempred((EContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listaval_mat2_sempred(Listaval_mat2Context _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean e_sempred(EContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 19);
		case 2:
			return precpred(_ctx, 18);
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 5:
			return precpred(_ctx, 15);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u0291\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\3\3\7\3R\n\3\f\3\16"+
		"\3U\13\3\3\4\3\4\5\4Y\n\4\3\4\3\4\5\4]\n\4\3\4\3\4\5\4a\n\4\3\4\3\4\5"+
		"\4e\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4n\n\4\3\4\3\4\5\4r\n\4\3\4\3\4"+
		"\5\4v\n\4\3\4\3\4\5\4z\n\4\3\4\3\4\3\4\3\4\5\4\u0080\n\4\3\4\3\4\5\4\u0084"+
		"\n\4\3\4\3\4\5\4\u0088\n\4\5\4\u008a\n\4\3\5\3\5\3\5\3\5\3\5\7\5\u0091"+
		"\n\5\f\5\16\5\u0094\13\5\3\5\3\5\5\5\u0098\n\5\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00ab\n\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00b8\n\7\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\5\b\u00c3\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00dd\n\t\3\n"+
		"\3\n\3\n\3\n\6\n\u00e3\n\n\r\n\16\n\u00e4\3\n\5\n\u00e8\n\n\3\n\3\n\3"+
		"\13\3\13\3\13\3\13\3\13\5\13\u00f1\n\13\3\f\3\f\3\f\3\f\5\f\u00f7\n\f"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3"+
		"\17\3\17\3\17\3\17\5\17\u010c\n\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\21\3\21\3\21\3\21\5\21\u011a\n\21\5\21\u011c\n\21\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\7\23\u012b\n\23"+
		"\f\23\16\23\u012e\13\23\3\23\3\23\3\23\3\23\3\23\5\23\u0135\n\23\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u014c\n\24\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\5\25\u0163\n\25\3\26\3\26\3\26\3\26\5\26\u0169\n\26\3\26\3"+
		"\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0176\n\27\3\30"+
		"\3\30\5\30\u017a\n\30\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\7\32\u0187\n\32\f\32\16\32\u018a\13\32\5\32\u018c\n\32\3\32\3\32"+
		"\3\32\7\32\u0191\n\32\f\32\16\32\u0194\13\32\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\5\33\u01ac\n\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\7\34\u01b8\n\34\f\34\16\34\u01bb\13\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\7\35\u01c4\n\35\f\35\16\35\u01c7\13\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\36\3\36\3\36\3\36\7\36\u01d2\n\36\f\36\16\36\u01d5\13\36"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\5\37\u01e0\n\37\3\37\3\37"+
		"\3\37\5\37\u01e5\n\37\3\37\3\37\5\37\u01e9\n\37\3 \3 \3 \7 \u01ee\n \f"+
		" \16 \u01f1\13 \3 \3 \3!\3!\5!\u01f7\n!\3!\5!\u01fa\n!\3!\3!\5!\u01fe"+
		"\n!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\5\"\u020f"+
		"\n\"\3#\3#\3#\3#\7#\u0215\n#\f#\16#\u0218\13#\3#\3#\3$\3$\3$\3$\5$\u0220"+
		"\n$\3$\3$\5$\u0224\n$\3$\5$\u0227\n$\3$\5$\u022a\n$\3$\5$\u022d\n$\3%"+
		"\3%\3%\3%\5%\u0233\n%\3%\3%\3%\3%\5%\u0239\n%\3%\3%\3%\3%\3%\5%\u0240"+
		"\n%\3%\3%\5%\u0244\n%\3&\3&\3&\3&\3&\3&\3&\3&\3&\5&\u024f\n&\3\'\3\'\3"+
		"\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\7\'\u0261\n\'\f"+
		"\'\16\'\u0264\13\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'"+
		"\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\5\'\u027b\n\'\3\'\3\'\3\'\3\'\3\'\3\'"+
		"\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\7\'\u028c\n\'\f\'\16\'\u028f\13\'"+
		"\3\'\2\4\62L(\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:<>@BDFHJL\2\f\4\2\16\16KK\3\2\"#\4\2??AA\4\2@@EE\4\2\25\25  \3\2"+
		"$%\4\2\23\23\27\27\3\2\25\26\3\2\30\35\3\2\36\37\2\u02d5\2N\3\2\2\2\4"+
		"S\3\2\2\2\6\u0089\3\2\2\2\b\u008b\3\2\2\2\n\u00aa\3\2\2\2\f\u00b7\3\2"+
		"\2\2\16\u00c2\3\2\2\2\20\u00dc\3\2\2\2\22\u00de\3\2\2\2\24\u00eb\3\2\2"+
		"\2\26\u00f2\3\2\2\2\30\u00f8\3\2\2\2\32\u0100\3\2\2\2\34\u0106\3\2\2\2"+
		"\36\u0111\3\2\2\2 \u011b\3\2\2\2\"\u011d\3\2\2\2$\u0134\3\2\2\2&\u014b"+
		"\3\2\2\2(\u0162\3\2\2\2*\u0164\3\2\2\2,\u0175\3\2\2\2.\u0179\3\2\2\2\60"+
		"\u017b\3\2\2\2\62\u018b\3\2\2\2\64\u01ab\3\2\2\2\66\u01ad\3\2\2\28\u01bf"+
		"\3\2\2\2:\u01cd\3\2\2\2<\u01df\3\2\2\2>\u01ea\3\2\2\2@\u01f6\3\2\2\2B"+
		"\u020e\3\2\2\2D\u0210\3\2\2\2F\u022c\3\2\2\2H\u0243\3\2\2\2J\u024e\3\2"+
		"\2\2L\u027a\3\2\2\2NO\5\4\3\2O\3\3\2\2\2PR\5\6\4\2QP\3\2\2\2RU\3\2\2\2"+
		"SQ\3\2\2\2ST\3\2\2\2T\5\3\2\2\2US\3\2\2\2VX\5\b\5\2WY\7\13\2\2XW\3\2\2"+
		"\2XY\3\2\2\2Y\u008a\3\2\2\2Z\\\5\n\6\2[]\7\13\2\2\\[\3\2\2\2\\]\3\2\2"+
		"\2]\u008a\3\2\2\2^`\5\f\7\2_a\7\13\2\2`_\3\2\2\2`a\3\2\2\2a\u008a\3\2"+
		"\2\2bd\5\16\b\2ce\7\13\2\2dc\3\2\2\2de\3\2\2\2e\u008a\3\2\2\2f\u008a\5"+
		"\20\t\2g\u008a\5\22\n\2h\u008a\5\30\r\2i\u008a\5\32\16\2j\u008a\5\34\17"+
		"\2km\5 \21\2ln\7\13\2\2ml\3\2\2\2mn\3\2\2\2n\u008a\3\2\2\2oq\5\"\22\2"+
		"pr\7\13\2\2qp\3\2\2\2qr\3\2\2\2r\u008a\3\2\2\2su\5(\25\2tv\7\13\2\2ut"+
		"\3\2\2\2uv\3\2\2\2v\u008a\3\2\2\2wy\5&\24\2xz\7\13\2\2yx\3\2\2\2yz\3\2"+
		"\2\2z\u008a\3\2\2\2{\u008a\58\35\2|\u008a\5:\36\2}\177\5> \2~\u0080\7"+
		"\13\2\2\177~\3\2\2\2\177\u0080\3\2\2\2\u0080\u008a\3\2\2\2\u0081\u0083"+
		"\5*\26\2\u0082\u0084\7\13\2\2\u0083\u0082\3\2\2\2\u0083\u0084\3\2\2\2"+
		"\u0084\u008a\3\2\2\2\u0085\u0087\5\66\34\2\u0086\u0088\7\13\2\2\u0087"+
		"\u0086\3\2\2\2\u0087\u0088\3\2\2\2\u0088\u008a\3\2\2\2\u0089V\3\2\2\2"+
		"\u0089Z\3\2\2\2\u0089^\3\2\2\2\u0089b\3\2\2\2\u0089f\3\2\2\2\u0089g\3"+
		"\2\2\2\u0089h\3\2\2\2\u0089i\3\2\2\2\u0089j\3\2\2\2\u0089k\3\2\2\2\u0089"+
		"o\3\2\2\2\u0089s\3\2\2\2\u0089w\3\2\2\2\u0089{\3\2\2\2\u0089|\3\2\2\2"+
		"\u0089}\3\2\2\2\u0089\u0081\3\2\2\2\u0089\u0085\3\2\2\2\u008a\7\3\2\2"+
		"\2\u008b\u008c\7!\2\2\u008c\u008d\7\5\2\2\u008d\u0092\5L\'\2\u008e\u008f"+
		"\7\3\2\2\u008f\u0091\5L\'\2\u0090\u008e\3\2\2\2\u0091\u0094\3\2\2\2\u0092"+
		"\u0090\3\2\2\2\u0092\u0093\3\2\2\2\u0093\u0095\3\2\2\2\u0094\u0092\3\2"+
		"\2\2\u0095\u0097\7\4\2\2\u0096\u0098\7\13\2\2\u0097\u0096\3\2\2\2\u0097"+
		"\u0098\3\2\2\2\u0098\t\3\2\2\2\u0099\u009a\7\"\2\2\u009a\u009b\7K\2\2"+
		"\u009b\u009c\7\n\2\2\u009c\u009d\5J&\2\u009d\u009e\7\22\2\2\u009e\u009f"+
		"\5L\'\2\u009f\u00ab\3\2\2\2\u00a0\u00a1\7\"\2\2\u00a1\u00a2\7K\2\2\u00a2"+
		"\u00a3\7\22\2\2\u00a3\u00ab\5L\'\2\u00a4\u00a5\7\"\2\2\u00a5\u00a6\7K"+
		"\2\2\u00a6\u00a7\7\n\2\2\u00a7\u00a8\5J&\2\u00a8\u00a9\7\f\2\2\u00a9\u00ab"+
		"\3\2\2\2\u00aa\u0099\3\2\2\2\u00aa\u00a0\3\2\2\2\u00aa\u00a4\3\2\2\2\u00ab"+
		"\13\3\2\2\2\u00ac\u00ad\7#\2\2\u00ad\u00ae\7K\2\2\u00ae\u00af\7\n\2\2"+
		"\u00af\u00b0\5J&\2\u00b0\u00b1\7\22\2\2\u00b1\u00b2\5L\'\2\u00b2\u00b8"+
		"\3\2\2\2\u00b3\u00b4\7#\2\2\u00b4\u00b5\7K\2\2\u00b5\u00b6\7\22\2\2\u00b6"+
		"\u00b8\5L\'\2\u00b7\u00ac\3\2\2\2\u00b7\u00b3\3\2\2\2\u00b8\r\3\2\2\2"+
		"\u00b9\u00ba\7K\2\2\u00ba\u00bb\7\20\2\2\u00bb\u00c3\5L\'\2\u00bc\u00bd"+
		"\7K\2\2\u00bd\u00be\7\21\2\2\u00be\u00c3\5L\'\2\u00bf\u00c0\7K\2\2\u00c0"+
		"\u00c1\7\22\2\2\u00c1\u00c3\5L\'\2\u00c2\u00b9\3\2\2\2\u00c2\u00bc\3\2"+
		"\2\2\u00c2\u00bf\3\2\2\2\u00c3\17\3\2\2\2\u00c4\u00c5\7&\2\2\u00c5\u00c6"+
		"\5L\'\2\u00c6\u00c7\7\6\2\2\u00c7\u00c8\5\4\3\2\u00c8\u00c9\7\7\2\2\u00c9"+
		"\u00dd\3\2\2\2\u00ca\u00cb\7&\2\2\u00cb\u00cc\5L\'\2\u00cc\u00cd\7\6\2"+
		"\2\u00cd\u00ce\5\4\3\2\u00ce\u00cf\7\7\2\2\u00cf\u00d0\7\'\2\2\u00d0\u00d1"+
		"\7\6\2\2\u00d1\u00d2\5\4\3\2\u00d2\u00d3\7\7\2\2\u00d3\u00dd\3\2\2\2\u00d4"+
		"\u00d5\7&\2\2\u00d5\u00d6\5L\'\2\u00d6\u00d7\7\6\2\2\u00d7\u00d8\5\4\3"+
		"\2\u00d8\u00d9\7\7\2\2\u00d9\u00da\7\'\2\2\u00da\u00db\5\20\t\2\u00db"+
		"\u00dd\3\2\2\2\u00dc\u00c4\3\2\2\2\u00dc\u00ca\3\2\2\2\u00dc\u00d4\3\2"+
		"\2\2\u00dd\21\3\2\2\2\u00de\u00df\7(\2\2\u00df\u00e0\5L\'\2\u00e0\u00e2"+
		"\7\6\2\2\u00e1\u00e3\5\24\13\2\u00e2\u00e1\3\2\2\2\u00e3\u00e4\3\2\2\2"+
		"\u00e4\u00e2\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00e7\3\2\2\2\u00e6\u00e8"+
		"\5\26\f\2\u00e7\u00e6\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8\u00e9\3\2\2\2"+
		"\u00e9\u00ea\7\7\2\2\u00ea\23\3\2\2\2\u00eb\u00ec\7)\2\2\u00ec\u00ed\5"+
		"L\'\2\u00ed\u00ee\7\n\2\2\u00ee\u00f0\5\4\3\2\u00ef\u00f1\7\62\2\2\u00f0"+
		"\u00ef\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\25\3\2\2\2\u00f2\u00f3\7*\2\2"+
		"\u00f3\u00f4\7\n\2\2\u00f4\u00f6\5\4\3\2\u00f5\u00f7\7\62\2\2\u00f6\u00f5"+
		"\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7\27\3\2\2\2\u00f8\u00f9\7/\2\2\u00f9"+
		"\u00fa\5L\'\2\u00fa\u00fb\7\'\2\2\u00fb\u00fc\7\6\2\2\u00fc\u00fd\5\4"+
		"\3\2\u00fd\u00fe\5 \21\2\u00fe\u00ff\7\7\2\2\u00ff\31\3\2\2\2\u0100\u0101"+
		"\7+\2\2\u0101\u0102\5L\'\2\u0102\u0103\7\6\2\2\u0103\u0104\5\4\3\2\u0104"+
		"\u0105\7\7\2\2\u0105\33\3\2\2\2\u0106\u0107\7,\2\2\u0107\u0108\t\2\2\2"+
		"\u0108\u010b\7-\2\2\u0109\u010c\5\36\20\2\u010a\u010c\5L\'\2\u010b\u0109"+
		"\3\2\2\2\u010b\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010e\7\6\2\2\u010e"+
		"\u010f\5\4\3\2\u010f\u0110\7\7\2\2\u0110\35\3\2\2\2\u0111\u0112\5L\'\2"+
		"\u0112\u0113\7.\2\2\u0113\u0114\5L\'\2\u0114\37\3\2\2\2\u0115\u011c\7"+
		"\62\2\2\u0116\u011c\7\60\2\2\u0117\u0119\7\61\2\2\u0118\u011a\5L\'\2\u0119"+
		"\u0118\3\2\2\2\u0119\u011a\3\2\2\2\u011a\u011c\3\2\2\2\u011b\u0115\3\2"+
		"\2\2\u011b\u0116\3\2\2\2\u011b\u0117\3\2\2\2\u011c!\3\2\2\2\u011d\u011e"+
		"\t\3\2\2\u011e\u011f\7K\2\2\u011f\u0120\7\n\2\2\u0120\u0121\7\b\2\2\u0121"+
		"\u0122\5J&\2\u0122\u0123\7\t\2\2\u0123\u0124\7\22\2\2\u0124\u0125\5$\23"+
		"\2\u0125#\3\2\2\2\u0126\u0127\7\b\2\2\u0127\u012c\5L\'\2\u0128\u0129\7"+
		"\3\2\2\u0129\u012b\5L\'\2\u012a\u0128\3\2\2\2\u012b\u012e\3\2\2\2\u012c"+
		"\u012a\3\2\2\2\u012c\u012d\3\2\2\2\u012d\u012f\3\2\2\2\u012e\u012c\3\2"+
		"\2\2\u012f\u0130\7\t\2\2\u0130\u0135\3\2\2\2\u0131\u0132\7\b\2\2\u0132"+
		"\u0135\7\t\2\2\u0133\u0135\7K\2\2\u0134\u0126\3\2\2\2\u0134\u0131\3\2"+
		"\2\2\u0134\u0133\3\2\2\2\u0135%\3\2\2\2\u0136\u0137\7K\2\2\u0137\u0138"+
		"\7\b\2\2\u0138\u0139\5L\'\2\u0139\u013a\7\t\2\2\u013a\u013b\7\22\2\2\u013b"+
		"\u013c\5L\'\2\u013c\u014c\3\2\2\2\u013d\u013e\7K\2\2\u013e\u013f\7\b\2"+
		"\2\u013f\u0140\5L\'\2\u0140\u0141\7\t\2\2\u0141\u0142\7\20\2\2\u0142\u0143"+
		"\5L\'\2\u0143\u014c\3\2\2\2\u0144\u0145\7K\2\2\u0145\u0146\7\b\2\2\u0146"+
		"\u0147\5L\'\2\u0147\u0148\7\t\2\2\u0148\u0149\7\21\2\2\u0149\u014a\5L"+
		"\'\2\u014a\u014c\3\2\2\2\u014b\u0136\3\2\2\2\u014b\u013d\3\2\2\2\u014b"+
		"\u0144\3\2\2\2\u014c\'\3\2\2\2\u014d\u014e\7K\2\2\u014e\u014f\7\r\2\2"+
		"\u014f\u0150\7\64\2\2\u0150\u0151\7\5\2\2\u0151\u0152\5L\'\2\u0152\u0153"+
		"\7\4\2\2\u0153\u0163\3\2\2\2\u0154\u0155\7K\2\2\u0155\u0156\7\r\2\2\u0156"+
		"\u0157\7\65\2\2\u0157\u0158\7\5\2\2\u0158\u0163\7\4\2\2\u0159\u015a\7"+
		"K\2\2\u015a\u015b\7\r\2\2\u015b\u015c\7\66\2\2\u015c\u015d\7\5\2\2\u015d"+
		"\u015e\7\67\2\2\u015e\u015f\7\n\2\2\u015f\u0160\5L\'\2\u0160\u0161\7\4"+
		"\2\2\u0161\u0163\3\2\2\2\u0162\u014d\3\2\2\2\u0162\u0154\3\2\2\2\u0162"+
		"\u0159\3\2\2\2\u0163)\3\2\2\2\u0164\u0165\t\3\2\2\u0165\u0168\7K\2\2\u0166"+
		"\u0167\7\n\2\2\u0167\u0169\5,\27\2\u0168\u0166\3\2\2\2\u0168\u0169\3\2"+
		"\2\2\u0169\u016a\3\2\2\2\u016a\u016b\7\22\2\2\u016b\u016c\5.\30\2\u016c"+
		"+\3\2\2\2\u016d\u016e\7\b\2\2\u016e\u016f\5,\27\2\u016f\u0170\7\t\2\2"+
		"\u0170\u0176\3\2\2\2\u0171\u0172\7\b\2\2\u0172\u0173\5J&\2\u0173\u0174"+
		"\7\t\2\2\u0174\u0176\3\2\2\2\u0175\u016d\3\2\2\2\u0175\u0171\3\2\2\2\u0176"+
		"-\3\2\2\2\u0177\u017a\5\60\31\2\u0178\u017a\5\64\33\2\u0179\u0177\3\2"+
		"\2\2\u0179\u0178\3\2\2\2\u017a/\3\2\2\2\u017b\u017c\7\b\2\2\u017c\u017d"+
		"\5\62\32\2\u017d\u017e\7\t\2\2\u017e\61\3\2\2\2\u017f\u0180\b\32\1\2\u0180"+
		"\u018c\5\60\31\2\u0181\u0182\5L\'\2\u0182\u0183\7\3\2\2\u0183\u0188\5"+
		"L\'\2\u0184\u0185\7\3\2\2\u0185\u0187\5L\'\2\u0186\u0184\3\2\2\2\u0187"+
		"\u018a\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189\u018c\3\2"+
		"\2\2\u018a\u0188\3\2\2\2\u018b\u017f\3\2\2\2\u018b\u0181\3\2\2\2\u018c"+
		"\u0192\3\2\2\2\u018d\u018e\f\5\2\2\u018e\u018f\7\3\2\2\u018f\u0191\5\60"+
		"\31\2\u0190\u018d\3\2\2\2\u0191\u0194\3\2\2\2\u0192\u0190\3\2\2\2\u0192"+
		"\u0193\3\2\2\2\u0193\63\3\2\2\2\u0194\u0192\3\2\2\2\u0195\u0196\5,\27"+
		"\2\u0196\u0197\7\5\2\2\u0197\u0198\7;\2\2\u0198\u0199\7\n\2\2\u0199\u019a"+
		"\5\64\33\2\u019a\u019b\7\3\2\2\u019b\u019c\79\2\2\u019c\u019d\7\n\2\2"+
		"\u019d\u019e\7G\2\2\u019e\u019f\7\4\2\2\u019f\u01ac\3\2\2\2\u01a0\u01a1"+
		"\5,\27\2\u01a1\u01a2\7\5\2\2\u01a2\u01a3\7;\2\2\u01a3\u01a4\7\n\2\2\u01a4"+
		"\u01a5\5L\'\2\u01a5\u01a6\7\3\2\2\u01a6\u01a7\79\2\2\u01a7\u01a8\7\n\2"+
		"\2\u01a8\u01a9\7G\2\2\u01a9\u01aa\7\4\2\2\u01aa\u01ac\3\2\2\2\u01ab\u0195"+
		"\3\2\2\2\u01ab\u01a0\3\2\2\2\u01ac\65\3\2\2\2\u01ad\u01ae\7K\2\2\u01ae"+
		"\u01af\7\b\2\2\u01af\u01b0\7G\2\2\u01b0\u01b1\7\t\2\2\u01b1\u01b2\7\b"+
		"\2\2\u01b2\u01b3\7G\2\2\u01b3\u01b9\7\t\2\2\u01b4\u01b5\7\b\2\2\u01b5"+
		"\u01b6\7G\2\2\u01b6\u01b8\7\t\2\2\u01b7\u01b4\3\2\2\2\u01b8\u01bb\3\2"+
		"\2\2\u01b9\u01b7\3\2\2\2\u01b9\u01ba\3\2\2\2\u01ba\u01bc\3\2\2\2\u01bb"+
		"\u01b9\3\2\2\2\u01bc\u01bd\7\22\2\2\u01bd\u01be\5L\'\2\u01be\67\3\2\2"+
		"\2\u01bf\u01c0\7:\2\2\u01c0\u01c1\7K\2\2\u01c1\u01c5\7\5\2\2\u01c2\u01c4"+
		"\5<\37\2\u01c3\u01c2\3\2\2\2\u01c4\u01c7\3\2\2\2\u01c5\u01c3\3\2\2\2\u01c5"+
		"\u01c6\3\2\2\2\u01c6\u01c8\3\2\2\2\u01c7\u01c5\3\2\2\2\u01c8\u01c9\7\4"+
		"\2\2\u01c9\u01ca\7\6\2\2\u01ca\u01cb\5\4\3\2\u01cb\u01cc\7\7\2\2\u01cc"+
		"9\3\2\2\2\u01cd\u01ce\7:\2\2\u01ce\u01cf\7K\2\2\u01cf\u01d3\7\5\2\2\u01d0"+
		"\u01d2\5<\37\2\u01d1\u01d0\3\2\2\2\u01d2\u01d5\3\2\2\2\u01d3\u01d1\3\2"+
		"\2\2\u01d3\u01d4\3\2\2\2\u01d4\u01d6\3\2\2\2\u01d5\u01d3\3\2\2\2\u01d6"+
		"\u01d7\7\4\2\2\u01d7\u01d8\7\25\2\2\u01d8\u01d9\7\34\2\2\u01d9\u01da\5"+
		"J&\2\u01da\u01db\7\6\2\2\u01db\u01dc\5\4\3\2\u01dc\u01dd\7\7\2\2\u01dd"+
		";\3\2\2\2\u01de\u01e0\t\2\2\2\u01df\u01de\3\2\2\2\u01df\u01e0\3\2\2\2"+
		"\u01e0\u01e1\3\2\2\2\u01e1\u01e2\7K\2\2\u01e2\u01e4\7\n\2\2\u01e3\u01e5"+
		"\7>\2\2\u01e4\u01e3\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5\u01e6\3\2\2\2\u01e6"+
		"\u01e8\5J&\2\u01e7\u01e9\7\3\2\2\u01e8\u01e7\3\2\2\2\u01e8\u01e9\3\2\2"+
		"\2\u01e9=\3\2\2\2\u01ea\u01eb\7K\2\2\u01eb\u01ef\7\5\2\2\u01ec\u01ee\5"+
		"@!\2\u01ed\u01ec\3\2\2\2\u01ee\u01f1\3\2\2\2\u01ef\u01ed\3\2\2\2\u01ef"+
		"\u01f0\3\2\2\2\u01f0\u01f2\3\2\2\2\u01f1\u01ef\3\2\2\2\u01f2\u01f3\7\4"+
		"\2\2\u01f3?\3\2\2\2\u01f4\u01f5\7K\2\2\u01f5\u01f7\7\n\2\2\u01f6\u01f4"+
		"\3\2\2\2\u01f6\u01f7\3\2\2\2\u01f7\u01f9\3\2\2\2\u01f8\u01fa\7\17\2\2"+
		"\u01f9\u01f8\3\2\2\2\u01f9\u01fa\3\2\2\2\u01fa\u01fb\3\2\2\2\u01fb\u01fd"+
		"\5L\'\2\u01fc\u01fe\7\3\2\2\u01fd\u01fc\3\2\2\2\u01fd\u01fe\3\2\2\2\u01fe"+
		"A\3\2\2\2\u01ff\u0200\t\4\2\2\u0200\u0201\7\5\2\2\u0201\u0202\5L\'\2\u0202"+
		"\u0203\7\4\2\2\u0203\u020f\3\2\2\2\u0204\u0205\7B\2\2\u0205\u0206\7\5"+
		"\2\2\u0206\u0207\5L\'\2\u0207\u0208\7\4\2\2\u0208\u020f\3\2\2\2\u0209"+
		"\u020a\t\5\2\2\u020a\u020b\7\5\2\2\u020b\u020c\5L\'\2\u020c\u020d\7\4"+
		"\2\2\u020d\u020f\3\2\2\2\u020e\u01ff\3\2\2\2\u020e\u0204\3\2\2\2\u020e"+
		"\u0209\3\2\2\2\u020fC\3\2\2\2\u0210\u0211\7<\2\2\u0211\u0212\7K\2\2\u0212"+
		"\u0216\7\6\2\2\u0213\u0215\5F$\2\u0214\u0213\3\2\2\2\u0215\u0218\3\2\2"+
		"\2\u0216\u0214\3\2\2\2\u0216\u0217\3\2\2\2\u0217\u0219\3\2\2\2\u0218\u0216"+
		"\3\2\2\2\u0219\u021a\7\7\2\2\u021aE\3\2\2\2\u021b\u021c\t\3\2\2\u021c"+
		"\u021f\7K\2\2\u021d\u021e\7\n\2\2\u021e\u0220\5J&\2\u021f\u021d\3\2\2"+
		"\2\u021f\u0220\3\2\2\2\u0220\u0223\3\2\2\2\u0221\u0222\7\22\2\2\u0222"+
		"\u0224\5L\'\2\u0223\u0221\3\2\2\2\u0223\u0224\3\2\2\2\u0224\u0226\3\2"+
		"\2\2\u0225\u0227\7\13\2\2\u0226\u0225\3\2\2\2\u0226\u0227\3\2\2\2\u0227"+
		"\u022d\3\2\2\2\u0228\u022a\7=\2\2\u0229\u0228\3\2\2\2\u0229\u022a\3\2"+
		"\2\2\u022a\u022b\3\2\2\2\u022b\u022d\5:\36\2\u022c\u021b\3\2\2\2\u022c"+
		"\u0229\3\2\2\2\u022dG\3\2\2\2\u022e\u022f\t\3\2\2\u022f\u0232\7K\2\2\u0230"+
		"\u0231\7\n\2\2\u0231\u0233\7K\2\2\u0232\u0230\3\2\2\2\u0232\u0233\3\2"+
		"\2\2\u0233\u0234\3\2\2\2\u0234\u0235\7\22\2\2\u0235\u0236\7K\2\2\u0236"+
		"\u0238\7\5\2\2\u0237\u0239\5@!\2\u0238\u0237\3\2\2\2\u0238\u0239\3\2\2"+
		"\2\u0239\u023a\3\2\2\2\u023a\u0244\7\4\2\2\u023b\u023c\t\3\2\2\u023c\u023f"+
		"\7K\2\2\u023d\u023e\7\n\2\2\u023e\u0240\7K\2\2\u023f\u023d\3\2\2\2\u023f"+
		"\u0240\3\2\2\2\u0240\u0241\3\2\2\2\u0241\u0242\7\22\2\2\u0242\u0244\7"+
		"K\2\2\u0243\u022e\3\2\2\2\u0243\u023b\3\2\2\2\u0244I\3\2\2\2\u0245\u024f"+
		"\7A\2\2\u0246\u024f\7B\2\2\u0247\u024f\7E\2\2\u0248\u024f\7C\2\2\u0249"+
		"\u024f\7D\2\2\u024a\u024b\7\b\2\2\u024b\u024c\5J&\2\u024c\u024d\7\t\2"+
		"\2\u024d\u024f\3\2\2\2\u024e\u0245\3\2\2\2\u024e\u0246\3\2\2\2\u024e\u0247"+
		"\3\2\2\2\u024e\u0248\3\2\2\2\u024e\u0249\3\2\2\2\u024e\u024a\3\2\2\2\u024f"+
		"K\3\2\2\2\u0250\u0251\b\'\1\2\u0251\u0252\t\6\2\2\u0252\u027b\5L\'\26"+
		"\u0253\u027b\t\7\2\2\u0254\u027b\5B\"\2\u0255\u027b\5> \2\u0256\u0257"+
		"\7K\2\2\u0257\u0258\7\b\2\2\u0258\u0259\7G\2\2\u0259\u025a\7\t\2\2\u025a"+
		"\u025b\7\b\2\2\u025b\u025c\7G\2\2\u025c\u0262\7\t\2\2\u025d\u025e\7\b"+
		"\2\2\u025e\u025f\7G\2\2\u025f\u0261\7\t\2\2\u0260\u025d\3\2\2\2\u0261"+
		"\u0264\3\2\2\2\u0262\u0260\3\2\2\2\u0262\u0263\3\2\2\2\u0263\u027b\3\2"+
		"\2\2\u0264\u0262\3\2\2\2\u0265\u0266\7K\2\2\u0266\u0267\7\b\2\2\u0267"+
		"\u0268\5L\'\2\u0268\u0269\7\t\2\2\u0269\u027b\3\2\2\2\u026a\u026b\7K\2"+
		"\2\u026b\u026c\7\r\2\2\u026c\u027b\78\2\2\u026d\u026e\7K\2\2\u026e\u026f"+
		"\7\r\2\2\u026f\u027b\79\2\2\u0270\u027b\7K\2\2\u0271\u027b\7\63\2\2\u0272"+
		"\u027b\7H\2\2\u0273\u027b\7G\2\2\u0274\u027b\7J\2\2\u0275\u027b\7I\2\2"+
		"\u0276\u0277\7\5\2\2\u0277\u0278\5L\'\2\u0278\u0279\7\4\2\2\u0279\u027b"+
		"\3\2\2\2\u027a\u0250\3\2\2\2\u027a\u0253\3\2\2\2\u027a\u0254\3\2\2\2\u027a"+
		"\u0255\3\2\2\2\u027a\u0256\3\2\2\2\u027a\u0265\3\2\2\2\u027a\u026a\3\2"+
		"\2\2\u027a\u026d\3\2\2\2\u027a\u0270\3\2\2\2\u027a\u0271\3\2\2\2\u027a"+
		"\u0272\3\2\2\2\u027a\u0273\3\2\2\2\u027a\u0274\3\2\2\2\u027a\u0275\3\2"+
		"\2\2\u027a\u0276\3\2\2\2\u027b\u028d\3\2\2\2\u027c\u027d\f\25\2\2\u027d"+
		"\u027e\t\b\2\2\u027e\u028c\5L\'\26\u027f\u0280\f\24\2\2\u0280\u0281\t"+
		"\t\2\2\u0281\u028c\5L\'\25\u0282\u0283\f\23\2\2\u0283\u0284\7\24\2\2\u0284"+
		"\u028c\5L\'\24\u0285\u0286\f\22\2\2\u0286\u0287\t\n\2\2\u0287\u028c\5"+
		"L\'\23\u0288\u0289\f\21\2\2\u0289\u028a\t\13\2\2\u028a\u028c\5L\'\22\u028b"+
		"\u027c\3\2\2\2\u028b\u027f\3\2\2\2\u028b\u0282\3\2\2\2\u028b\u0285\3\2"+
		"\2\2\u028b\u0288\3\2\2\2\u028c\u028f\3\2\2\2\u028d\u028b\3\2\2\2\u028d"+
		"\u028e\3\2\2\2\u028eM\3\2\2\2\u028f\u028d\3\2\2\2ASX\\`dmquy\177\u0083"+
		"\u0087\u0089\u0092\u0097\u00aa\u00b7\u00c2\u00dc\u00e4\u00e7\u00f0\u00f6"+
		"\u010b\u0119\u011b\u012c\u0134\u014b\u0162\u0168\u0175\u0179\u0188\u018b"+
		"\u0192\u01ab\u01b9\u01c5\u01d3\u01df\u01e4\u01e8\u01ef\u01f6\u01f9\u01fd"+
		"\u020e\u0216\u021f\u0223\u0226\u0229\u022c\u0232\u0238\u023f\u0243\u024e"+
		"\u0262\u027a\u028b\u028d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}