// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Tswift_GrammarParser struct {
	*antlr.BaseParser
}

var Tswift_GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswift_grammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "')'", "'('", "'{'", "'}'", "':'", "'?'", "'+='", "'-='",
		"'='", "'/'", "'%'", "'-'", "'+'", "'*'", "'=='", "'!='", "'>='", "'<='",
		"'>'", "'<'", "'&&'", "'||'", "'!'", "'print'", "'var'", "'let'", "'true'",
		"'false'", "'return'", "'if'", "'else'", "'switch'", "'case'", "'break'",
		"'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", "'continue'",
		"'nil'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'",
	}
	staticData.SymbolicNames = []string{
		"", "", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "DOSPT", "INTERROGACION",
		"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR",
		"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR",
		"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "RETURN",
		"IF", "ELSE", "SWITCH", "CASE", "BREAK", "DEFAULT", "WHILE", "FOR",
		"IN", "RANGO", "GUARD", "CONTINUE", "NIL", "INT", "FLOAT", "BOOL", "CHARACTER",
		"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID",
		"UL_C", "ML_C",
	}
	staticData.RuleNames = []string{
		"s", "l_sentencias", "sentencia", "declaracion", "constante", "tipo",
		"e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 5, 1, 18, 8, 1, 10, 1, 12,
		1, 21, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 28, 8, 2, 10, 2, 12, 2,
		31, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 56, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 75, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 82, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 98, 8, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 115, 8, 6, 10, 6, 12, 6, 118, 9, 6, 1, 6, 0, 1, 12,
		7, 0, 2, 4, 6, 8, 10, 12, 0, 6, 2, 0, 13, 13, 24, 24, 1, 0, 28, 29, 1,
		0, 16, 21, 1, 0, 22, 23, 2, 0, 11, 11, 15, 15, 1, 0, 13, 14, 137, 0, 14,
		1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 55, 1, 0, 0, 0, 8,
		74, 1, 0, 0, 0, 10, 81, 1, 0, 0, 0, 12, 97, 1, 0, 0, 0, 14, 15, 3, 2, 1,
		0, 15, 1, 1, 0, 0, 0, 16, 18, 3, 4, 2, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1,
		0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 3, 1, 0, 0, 0, 21,
		19, 1, 0, 0, 0, 22, 23, 5, 25, 0, 0, 23, 24, 5, 3, 0, 0, 24, 29, 3, 12,
		6, 0, 25, 26, 5, 1, 0, 0, 26, 28, 3, 12, 6, 0, 27, 25, 1, 0, 0, 0, 28,
		31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0,
		0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 2, 0, 0, 33, 37, 1, 0, 0, 0, 34, 37,
		3, 6, 3, 0, 35, 37, 3, 8, 4, 0, 36, 22, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0,
		36, 35, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38, 39, 5, 26, 0, 0, 39, 40, 5,
		54, 0, 0, 40, 41, 5, 6, 0, 0, 41, 42, 3, 10, 5, 0, 42, 43, 5, 10, 0, 0,
		43, 44, 3, 12, 6, 0, 44, 56, 1, 0, 0, 0, 45, 46, 5, 26, 0, 0, 46, 47, 5,
		54, 0, 0, 47, 48, 5, 10, 0, 0, 48, 56, 3, 12, 6, 0, 49, 50, 5, 26, 0, 0,
		50, 51, 5, 54, 0, 0, 51, 52, 5, 6, 0, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5,
		7, 0, 0, 54, 56, 1, 0, 0, 0, 55, 38, 1, 0, 0, 0, 55, 45, 1, 0, 0, 0, 55,
		49, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 27, 0, 0, 58, 59, 5, 54,
		0, 0, 59, 60, 5, 6, 0, 0, 60, 61, 3, 10, 5, 0, 61, 62, 5, 10, 0, 0, 62,
		63, 3, 12, 6, 0, 63, 75, 1, 0, 0, 0, 64, 65, 5, 27, 0, 0, 65, 66, 5, 54,
		0, 0, 66, 67, 5, 10, 0, 0, 67, 75, 3, 12, 6, 0, 68, 69, 5, 27, 0, 0, 69,
		70, 5, 54, 0, 0, 70, 71, 5, 6, 0, 0, 71, 72, 3, 10, 5, 0, 72, 73, 5, 7,
		0, 0, 73, 75, 1, 0, 0, 0, 74, 57, 1, 0, 0, 0, 74, 64, 1, 0, 0, 0, 74, 68,
		1, 0, 0, 0, 75, 9, 1, 0, 0, 0, 76, 82, 5, 44, 0, 0, 77, 82, 5, 45, 0, 0,
		78, 82, 5, 48, 0, 0, 79, 82, 5, 46, 0, 0, 80, 82, 5, 47, 0, 0, 81, 76,
		1, 0, 0, 0, 81, 77, 1, 0, 0, 0, 81, 78, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 80, 1, 0, 0, 0, 82, 11, 1, 0, 0, 0, 83, 84, 6, 6, -1, 0, 84, 85, 7,
		0, 0, 0, 85, 98, 3, 12, 6, 12, 86, 98, 7, 1, 0, 0, 87, 98, 5, 43, 0, 0,
		88, 98, 5, 54, 0, 0, 89, 98, 5, 51, 0, 0, 90, 98, 5, 50, 0, 0, 91, 98,
		5, 53, 0, 0, 92, 98, 5, 52, 0, 0, 93, 94, 5, 3, 0, 0, 94, 95, 3, 12, 6,
		0, 95, 96, 5, 2, 0, 0, 96, 98, 1, 0, 0, 0, 97, 83, 1, 0, 0, 0, 97, 86,
		1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 88, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0,
		97, 90, 1, 0, 0, 0, 97, 91, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 93, 1,
		0, 0, 0, 98, 116, 1, 0, 0, 0, 99, 100, 10, 14, 0, 0, 100, 101, 7, 2, 0,
		0, 101, 115, 3, 12, 6, 15, 102, 103, 10, 13, 0, 0, 103, 104, 7, 3, 0, 0,
		104, 115, 3, 12, 6, 14, 105, 106, 10, 11, 0, 0, 106, 107, 7, 4, 0, 0, 107,
		115, 3, 12, 6, 12, 108, 109, 10, 10, 0, 0, 109, 110, 7, 5, 0, 0, 110, 115,
		3, 12, 6, 11, 111, 112, 10, 9, 0, 0, 112, 113, 5, 12, 0, 0, 113, 115, 3,
		12, 6, 10, 114, 99, 1, 0, 0, 0, 114, 102, 1, 0, 0, 0, 114, 105, 1, 0, 0,
		0, 114, 108, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 13, 1, 0, 0, 0, 118, 116, 1,
		0, 0, 0, 9, 19, 29, 36, 55, 74, 81, 97, 114, 116,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// Tswift_GrammarParserInit initializes any static state used to implement Tswift_GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTswift_GrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tswift_GrammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.once.Do(tswift_grammarParserInit)
}

// NewTswift_GrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTswift_GrammarParser(input antlr.TokenStream) *Tswift_GrammarParser {
	Tswift_GrammarParserInit()
	this := new(Tswift_GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Tswift_GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tswift_Grammar.g4"

	return this
}

// Tswift_GrammarParser tokens.
const (
	Tswift_GrammarParserEOF           = antlr.TokenEOF
	Tswift_GrammarParserT__0          = 1
	Tswift_GrammarParserPARDER        = 2
	Tswift_GrammarParserPARIZQ        = 3
	Tswift_GrammarParserLLAVEIZQ      = 4
	Tswift_GrammarParserLLAVEDER      = 5
	Tswift_GrammarParserDOSPT         = 6
	Tswift_GrammarParserINTERROGACION = 7
	Tswift_GrammarParserMASIGUAL      = 8
	Tswift_GrammarParserMENOSIGUAL    = 9
	Tswift_GrammarParserIGUAL         = 10
	Tswift_GrammarParserDIV           = 11
	Tswift_GrammarParserMOD           = 12
	Tswift_GrammarParserMENOS         = 13
	Tswift_GrammarParserMAS           = 14
	Tswift_GrammarParserPOR           = 15
	Tswift_GrammarParserIGUALIGUAL    = 16
	Tswift_GrammarParserDIFERENTE     = 17
	Tswift_GrammarParserMAYORIGUAL    = 18
	Tswift_GrammarParserMENORIGUAL    = 19
	Tswift_GrammarParserMAYOR         = 20
	Tswift_GrammarParserMENOR         = 21
	Tswift_GrammarParserAND           = 22
	Tswift_GrammarParserOR            = 23
	Tswift_GrammarParserNOT           = 24
	Tswift_GrammarParserPRINT         = 25
	Tswift_GrammarParserVAR           = 26
	Tswift_GrammarParserLET           = 27
	Tswift_GrammarParserTRUE          = 28
	Tswift_GrammarParserFALSE         = 29
	Tswift_GrammarParserRETURN        = 30
	Tswift_GrammarParserIF            = 31
	Tswift_GrammarParserELSE          = 32
	Tswift_GrammarParserSWITCH        = 33
	Tswift_GrammarParserCASE          = 34
	Tswift_GrammarParserBREAK         = 35
	Tswift_GrammarParserDEFAULT       = 36
	Tswift_GrammarParserWHILE         = 37
	Tswift_GrammarParserFOR           = 38
	Tswift_GrammarParserIN            = 39
	Tswift_GrammarParserRANGO         = 40
	Tswift_GrammarParserGUARD         = 41
	Tswift_GrammarParserCONTINUE      = 42
	Tswift_GrammarParserNIL           = 43
	Tswift_GrammarParserINT           = 44
	Tswift_GrammarParserFLOAT         = 45
	Tswift_GrammarParserBOOL          = 46
	Tswift_GrammarParserCHARACTER     = 47
	Tswift_GrammarParserSTRING        = 48
	Tswift_GrammarParserENBLANCO      = 49
	Tswift_GrammarParserENTERO        = 50
	Tswift_GrammarParserDECIMAL       = 51
	Tswift_GrammarParserCARACTER      = 52
	Tswift_GrammarParserCADENA        = 53
	Tswift_GrammarParserID            = 54
	Tswift_GrammarParserUL_C          = 55
	Tswift_GrammarParserML_C          = 56
)

// Tswift_GrammarParser rules.
const (
	Tswift_GrammarParserRULE_s            = 0
	Tswift_GrammarParserRULE_l_sentencias = 1
	Tswift_GrammarParserRULE_sentencia    = 2
	Tswift_GrammarParserRULE_declaracion  = 3
	Tswift_GrammarParserRULE_constante    = 4
	Tswift_GrammarParserRULE_tipo         = 5
	Tswift_GrammarParserRULE_e            = 6
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) CopyAll(ctx *SContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SLSentenciasContext struct {
	SContext
}

func NewSLSentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SLSentenciasContext {
	var p = new(SLSentenciasContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *SLSentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SLSentenciasContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *SLSentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSLSentencias(s)
	}
}

func (s *SLSentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSLSentencias(s)
	}
}

func (s *SLSentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSLSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tswift_GrammarParserRULE_s)
	localctx = NewSLSentenciasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.L_sentencias()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_sentenciasContext is an interface to support dynamic dispatch.
type IL_sentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_sentenciasContext differentiates from other interfaces.
	IsL_sentenciasContext()
}

type L_sentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_sentenciasContext() *L_sentenciasContext {
	var p = new(L_sentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias
	return p
}

func InitEmptyL_sentenciasContext(p *L_sentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias
}

func (*L_sentenciasContext) IsL_sentenciasContext() {}

func NewL_sentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_sentenciasContext {
	var p = new(L_sentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias

	return p
}

func (s *L_sentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *L_sentenciasContext) CopyAll(ctx *L_sentenciasContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_sentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_sentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_SentenciaContext struct {
	L_sentenciasContext
}

func NewL_SentenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_SentenciaContext {
	var p = new(L_SentenciaContext)

	InitEmptyL_sentenciasContext(&p.L_sentenciasContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentenciasContext))

	return p
}

func (s *L_SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_SentenciaContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *L_SentenciaContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaContext)
}

func (s *L_SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Sentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_sentencias() (localctx IL_sentenciasContext) {
	localctx = NewL_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Tswift_GrammarParserRULE_l_sentencias)
	var _la int

	localctx = NewL_SentenciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&234881024) != 0 {
		{
			p.SetState(16)
			p.Sentencia()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) CopyAll(ctx *SentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type S_ConsolaContext struct {
	SentenciaContext
}

func NewS_ConsolaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ConsolaContext {
	var p = new(S_ConsolaContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ConsolaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ConsolaContext) PRINT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPRINT, 0)
}

func (s *S_ConsolaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *S_ConsolaContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *S_ConsolaContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *S_ConsolaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *S_ConsolaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Consola(s)
	}
}

func (s *S_ConsolaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Consola(s)
	}
}

func (s *S_ConsolaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Consola(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_ConstanteContext struct {
	SentenciaContext
}

func NewS_ConstanteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ConstanteContext {
	var p = new(S_ConstanteContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ConstanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ConstanteContext) Constante() IConstanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstanteContext)
}

func (s *S_ConstanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Constante(s)
	}
}

func (s *S_ConstanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Constante(s)
	}
}

func (s *S_ConstanteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Constante(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_DeclaracionContext struct {
	SentenciaContext
}

func NewS_DeclaracionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_DeclaracionContext {
	var p = new(S_DeclaracionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_DeclaracionContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *S_DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Tswift_GrammarParserRULE_sentencia)
	var _la int

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserPRINT:
		localctx = NewS_ConsolaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(Tswift_GrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.e(0)
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarParserT__0 {
			{
				p.SetState(25)
				p.Match(Tswift_GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(26)
				p.e(0)
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(32)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserVAR:
		localctx = NewS_DeclaracionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Declaracion()
		}

	case Tswift_GrammarParserLET:
		localctx = NewS_ConstanteContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Constante()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) CopyAll(ctx *DeclaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_Tipo_ValContext struct {
	DeclaracionContext
}

func NewDeclaracion_Tipo_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_ValContext {
	var p = new(Declaracion_Tipo_ValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_Tipo_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_Tipo_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_Tipo_ValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_Tipo_ValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_Tipo_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Declaracion_Tipo_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Tipo_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_Tipo_noValContext struct {
	DeclaracionContext
}

func NewDeclaracion_Tipo_noValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_noValContext {
	var p = new(Declaracion_Tipo_noValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_Tipo_noValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_noValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_Tipo_noValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_Tipo_noValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_Tipo_noValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_noValContext) INTERROGACION() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINTERROGACION, 0)
}

func (s *Declaracion_Tipo_noValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Tipo_noVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_ValContext struct {
	DeclaracionContext
}

func NewDeclaracion_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_ValContext {
	var p = new(Declaracion_ValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Declaracion_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Tswift_GrammarParserRULE_declaracion)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Tipo()
		}
		{
			p.SetState(42)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.e(0)
		}

	case 2:
		localctx = NewDeclaracion_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.e(0)
		}

	case 3:
		localctx = NewDeclaracion_Tipo_noValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Tipo()
		}
		{
			p.SetState(53)
			p.Match(Tswift_GrammarParserINTERROGACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstanteContext is an interface to support dynamic dispatch.
type IConstanteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstanteContext differentiates from other interfaces.
	IsConstanteContext()
}

type ConstanteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstanteContext() *ConstanteContext {
	var p = new(ConstanteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_constante
	return p
}

func InitEmptyConstanteContext(p *ConstanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_constante
}

func (*ConstanteContext) IsConstanteContext() {}

func NewConstanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstanteContext {
	var p = new(ConstanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_constante

	return p
}

func (s *ConstanteContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstanteContext) CopyAll(ctx *ConstanteContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConstanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Constante_ValContext struct {
	ConstanteContext
}

func NewConstante_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Constante_ValContext {
	var p = new(Constante_ValContext)

	InitEmptyConstanteContext(&p.ConstanteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstanteContext))

	return p
}

func (s *Constante_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constante_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Constante_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Constante_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Constante_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Constante_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterConstante_Val(s)
	}
}

func (s *Constante_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitConstante_Val(s)
	}
}

func (s *Constante_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitConstante_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Constante_Tipo_ValContext struct {
	ConstanteContext
}

func NewConstante_Tipo_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Constante_Tipo_ValContext {
	var p = new(Constante_Tipo_ValContext)

	InitEmptyConstanteContext(&p.ConstanteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstanteContext))

	return p
}

func (s *Constante_Tipo_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constante_Tipo_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Constante_Tipo_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Constante_Tipo_ValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Constante_Tipo_ValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Constante_Tipo_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Constante_Tipo_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Constante_Tipo_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterConstante_Tipo_Val(s)
	}
}

func (s *Constante_Tipo_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitConstante_Tipo_Val(s)
	}
}

func (s *Constante_Tipo_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitConstante_Tipo_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Constante_Tipo_noValContext struct {
	ConstanteContext
}

func NewConstante_Tipo_noValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Constante_Tipo_noValContext {
	var p = new(Constante_Tipo_noValContext)

	InitEmptyConstanteContext(&p.ConstanteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstanteContext))

	return p
}

func (s *Constante_Tipo_noValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constante_Tipo_noValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Constante_Tipo_noValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Constante_Tipo_noValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Constante_Tipo_noValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Constante_Tipo_noValContext) INTERROGACION() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINTERROGACION, 0)
}

func (s *Constante_Tipo_noValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterConstante_Tipo_noVal(s)
	}
}

func (s *Constante_Tipo_noValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitConstante_Tipo_noVal(s)
	}
}

func (s *Constante_Tipo_noValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitConstante_Tipo_noVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Constante() (localctx IConstanteContext) {
	localctx = NewConstanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Tswift_GrammarParserRULE_constante)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstante_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(Tswift_GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Tipo()
		}
		{
			p.SetState(61)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.e(0)
		}

	case 2:
		localctx = NewConstante_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(Tswift_GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.e(0)
		}

	case 3:
		localctx = NewConstante_Tipo_noValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(Tswift_GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Tipo()
		}
		{
			p.SetState(72)
			p.Match(Tswift_GrammarParserINTERROGACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) CopyAll(ctx *TipoContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Tipo_FloatContext struct {
	TipoContext
}

func NewTipo_FloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_FloatContext {
	var p = new(Tipo_FloatContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFLOAT, 0)
}

func (s *Tipo_FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Float(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_CharacterContext struct {
	TipoContext
}

func NewTipo_CharacterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_CharacterContext {
	var p = new(Tipo_CharacterContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_CharacterContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCHARACTER, 0)
}

func (s *Tipo_CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Character(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_StringContext struct {
	TipoContext
}

func NewTipo_StringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_StringContext {
	var p = new(Tipo_StringContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSTRING, 0)
}

func (s *Tipo_StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_String(s)
	}
}

func (s *Tipo_StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_String(s)
	}
}

func (s *Tipo_StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_String(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_BoolContext struct {
	TipoContext
}

func NewTipo_BoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_BoolContext {
	var p = new(Tipo_BoolContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserBOOL, 0)
}

func (s *Tipo_BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Bool(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_IntContext struct {
	TipoContext
}

func NewTipo_IntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_IntContext {
	var p = new(Tipo_IntContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_IntContext) INT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINT, 0)
}

func (s *Tipo_IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Int(s)
	}
}

func (s *Tipo_IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Int(s)
	}
}

func (s *Tipo_IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Tswift_GrammarParserRULE_tipo)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserINT:
		localctx = NewTipo_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(Tswift_GrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserFLOAT:
		localctx = NewTipo_FloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(Tswift_GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserSTRING:
		localctx = NewTipo_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(Tswift_GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserBOOL:
		localctx = NewTipo_BoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Match(Tswift_GrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCHARACTER:
		localctx = NewTipo_CharacterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)
			p.Match(Tswift_GrammarParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_RelContext struct {
	EContext
	op antlr.Token
}

func NewExpr_RelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_RelContext {
	var p = new(Expr_RelContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_RelContext) GetOp() antlr.Token { return s.op }

func (s *Expr_RelContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_RelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_RelContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_RelContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_RelContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUALIGUAL, 0)
}

func (s *Expr_RelContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIFERENTE, 0)
}

func (s *Expr_RelContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYORIGUAL, 0)
}

func (s *Expr_RelContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYOR, 0)
}

func (s *Expr_RelContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENORIGUAL, 0)
}

func (s *Expr_RelContext) MENOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOR, 0)
}

func (s *Expr_RelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Rel(s)
	}
}

func (s *Expr_RelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Rel(s)
	}
}

func (s *Expr_RelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Rel(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_DecimalContext struct {
	EContext
}

func NewExpr_DecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_DecimalContext {
	var p = new(Expr_DecimalContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDECIMAL, 0)
}

func (s *Expr_DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Decimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CaracterContext struct {
	EContext
}

func NewExpr_CaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CaracterContext {
	var p = new(Expr_CaracterContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCARACTER, 0)
}

func (s *Expr_CaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Caracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_SumResContext struct {
	EContext
	op antlr.Token
}

func NewExpr_SumResContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_SumResContext {
	var p = new(Expr_SumResContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_SumResContext) GetOp() antlr.Token { return s.op }

func (s *Expr_SumResContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_SumResContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_SumResContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_SumResContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_SumResContext) MAS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAS, 0)
}

func (s *Expr_SumResContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *Expr_SumResContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_SumRes(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_NegContext struct {
	EContext
	op antlr.Token
}

func NewExpr_NegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NegContext {
	var p = new(Expr_NegContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NegContext) GetOp() antlr.Token { return s.op }

func (s *Expr_NegContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_NegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NegContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_NegContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *Expr_NegContext) NOT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserNOT, 0)
}

func (s *Expr_NegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Neg(s)
	}
}

func (s *Expr_NegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Neg(s)
	}
}

func (s *Expr_NegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Neg(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_MulDivContext struct {
	EContext
	op antlr.Token
}

func NewExpr_MulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_MulDivContext {
	var p = new(Expr_MulDivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_MulDivContext) GetOp() antlr.Token { return s.op }

func (s *Expr_MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_MulDivContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_MulDivContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_MulDivContext) POR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPOR, 0)
}

func (s *Expr_MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIV, 0)
}

func (s *Expr_MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_MulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_NilContext struct {
	EContext
}

func NewExpr_NilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NilContext {
	var p = new(Expr_NilContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserNIL, 0)
}

func (s *Expr_NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Nil(s)
	}
}

func (s *Expr_NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Nil(s)
	}
}

func (s *Expr_NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Nil(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CadenaContext struct {
	EContext
}

func NewExpr_CadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CadenaContext {
	var p = new(Expr_CadenaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCADENA, 0)
}

func (s *Expr_CadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Cadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_IdContext struct {
	EContext
}

func NewExpr_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_IdContext {
	var p = new(Expr_IdContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Id(s)
	}
}

func (s *Expr_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Id(s)
	}
}

func (s *Expr_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ModContext struct {
	EContext
}

func NewExpr_ModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ModContext {
	var p = new(Expr_ModContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ModContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_ModContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMOD, 0)
}

func (s *Expr_ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Mod(s)
	}
}

func (s *Expr_ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Mod(s)
	}
}

func (s *Expr_ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Mod(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ParContext struct {
	EContext
}

func NewExpr_ParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ParContext {
	var p = new(Expr_ParContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ParContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Expr_ParContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_ParContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Expr_ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Par(s)
	}
}

func (s *Expr_ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Par(s)
	}
}

func (s *Expr_ParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Par(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_LogicaContext struct {
	EContext
	op antlr.Token
}

func NewExpr_LogicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_LogicaContext {
	var p = new(Expr_LogicaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_LogicaContext) GetOp() antlr.Token { return s.op }

func (s *Expr_LogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_LogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_LogicaContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_LogicaContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_LogicaContext) AND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserAND, 0)
}

func (s *Expr_LogicaContext) OR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserOR, 0)
}

func (s *Expr_LogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Logica(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_BooleanoContext struct {
	EContext
	op antlr.Token
}

func NewExpr_BooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_BooleanoContext {
	var p = new(Expr_BooleanoContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_BooleanoContext) GetOp() antlr.Token { return s.op }

func (s *Expr_BooleanoContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_BooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_BooleanoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserTRUE, 0)
}

func (s *Expr_BooleanoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFALSE, 0)
}

func (s *Expr_BooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Booleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_EnteroContext struct {
	EContext
}

func NewExpr_EnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_EnteroContext {
	var p = new(Expr_EnteroContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_EnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_EnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, 0)
}

func (s *Expr_EnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Entero(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *Tswift_GrammarParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, Tswift_GrammarParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserMENOS, Tswift_GrammarParserNOT:
		localctx = NewExpr_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(84)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_NegContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserMENOS || _la == Tswift_GrammarParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_NegContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(85)
			p.e(12)
		}

	case Tswift_GrammarParserTRUE, Tswift_GrammarParserFALSE:
		localctx = NewExpr_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_BooleanoContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserTRUE || _la == Tswift_GrammarParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_BooleanoContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case Tswift_GrammarParserNIL:
		localctx = NewExpr_NilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)
			p.Match(Tswift_GrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserID:
		localctx = NewExpr_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserDECIMAL:
		localctx = NewExpr_DecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Match(Tswift_GrammarParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserENTERO:
		localctx = NewExpr_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCADENA:
		localctx = NewExpr_CadenaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(Tswift_GrammarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCARACTER:
		localctx = NewExpr_CaracterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(Tswift_GrammarParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserPARIZQ:
		localctx = NewExpr_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.e(0)
		}
		{
			p.SetState(95)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_RelContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(100)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_RelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4128768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_RelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(101)
					p.e(15)
				}

			case 2:
				localctx = NewExpr_LogicaContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(103)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_LogicaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserAND || _la == Tswift_GrammarParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_LogicaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)
					p.e(14)
				}

			case 3:
				localctx = NewExpr_MulDivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(106)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserDIV || _la == Tswift_GrammarParserPOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)
					p.e(12)
				}

			case 4:
				localctx = NewExpr_SumResContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(109)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_SumResContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserMENOS || _la == Tswift_GrammarParserMAS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_SumResContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(110)
					p.e(11)
				}

			case 5:
				localctx = NewExpr_ModContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(Tswift_GrammarParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.e(10)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *Tswift_GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Tswift_GrammarParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
