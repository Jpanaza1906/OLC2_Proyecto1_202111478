// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

// BaseTswift_GrammarListener is a complete listener for a parse tree produced by Tswift_GrammarParser.
type BaseTswift_GrammarListener struct{}

var _ Tswift_GrammarListener = &BaseTswift_GrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTswift_GrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTswift_GrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTswift_GrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTswift_GrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSLSentencias is called when production SLSentencias is entered.
func (s *BaseTswift_GrammarListener) EnterSLSentencias(ctx *SLSentenciasContext) {}

// ExitSLSentencias is called when production SLSentencias is exited.
func (s *BaseTswift_GrammarListener) ExitSLSentencias(ctx *SLSentenciasContext) {}

// EnterL_Sentencia is called when production L_Sentencia is entered.
func (s *BaseTswift_GrammarListener) EnterL_Sentencia(ctx *L_SentenciaContext) {}

// ExitL_Sentencia is called when production L_Sentencia is exited.
func (s *BaseTswift_GrammarListener) ExitL_Sentencia(ctx *L_SentenciaContext) {}

// EnterS_Consola is called when production S_Consola is entered.
func (s *BaseTswift_GrammarListener) EnterS_Consola(ctx *S_ConsolaContext) {}

// ExitS_Consola is called when production S_Consola is exited.
func (s *BaseTswift_GrammarListener) ExitS_Consola(ctx *S_ConsolaContext) {}

// EnterS_Declaracion is called when production S_Declaracion is entered.
func (s *BaseTswift_GrammarListener) EnterS_Declaracion(ctx *S_DeclaracionContext) {}

// ExitS_Declaracion is called when production S_Declaracion is exited.
func (s *BaseTswift_GrammarListener) ExitS_Declaracion(ctx *S_DeclaracionContext) {}

// EnterDeclaracion_Tipo_Val is called when production Declaracion_Tipo_Val is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) {}

// ExitDeclaracion_Tipo_Val is called when production Declaracion_Tipo_Val is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) {}

// EnterDeclaracion_Val is called when production Declaracion_Val is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Val(ctx *Declaracion_ValContext) {}

// ExitDeclaracion_Val is called when production Declaracion_Val is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Val(ctx *Declaracion_ValContext) {}

// EnterDeclaracion_Tipo_noVal is called when production Declaracion_Tipo_noVal is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) {
}

// ExitDeclaracion_Tipo_noVal is called when production Declaracion_Tipo_noVal is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) {}

// EnterTipo_Int is called when production Tipo_Int is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Int(ctx *Tipo_IntContext) {}

// ExitTipo_Int is called when production Tipo_Int is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Int(ctx *Tipo_IntContext) {}

// EnterTipo_Float is called when production Tipo_Float is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Float(ctx *Tipo_FloatContext) {}

// ExitTipo_Float is called when production Tipo_Float is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Float(ctx *Tipo_FloatContext) {}

// EnterTipo_String is called when production Tipo_String is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_String(ctx *Tipo_StringContext) {}

// ExitTipo_String is called when production Tipo_String is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_String(ctx *Tipo_StringContext) {}

// EnterTipo_Bool is called when production Tipo_Bool is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Bool(ctx *Tipo_BoolContext) {}

// ExitTipo_Bool is called when production Tipo_Bool is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Bool(ctx *Tipo_BoolContext) {}

// EnterTipo_Character is called when production Tipo_Character is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Character(ctx *Tipo_CharacterContext) {}

// ExitTipo_Character is called when production Tipo_Character is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Character(ctx *Tipo_CharacterContext) {}

// EnterExpr_Rel is called when production Expr_Rel is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Rel(ctx *Expr_RelContext) {}

// ExitExpr_Rel is called when production Expr_Rel is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Rel(ctx *Expr_RelContext) {}

// EnterExpr_Decimal is called when production Expr_Decimal is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Decimal(ctx *Expr_DecimalContext) {}

// ExitExpr_Decimal is called when production Expr_Decimal is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Decimal(ctx *Expr_DecimalContext) {}

// EnterExpr_Caracter is called when production Expr_Caracter is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Caracter(ctx *Expr_CaracterContext) {}

// ExitExpr_Caracter is called when production Expr_Caracter is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Caracter(ctx *Expr_CaracterContext) {}

// EnterExpr_SumRes is called when production Expr_SumRes is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_SumRes(ctx *Expr_SumResContext) {}

// ExitExpr_SumRes is called when production Expr_SumRes is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_SumRes(ctx *Expr_SumResContext) {}

// EnterExpr_Neg is called when production Expr_Neg is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Neg(ctx *Expr_NegContext) {}

// ExitExpr_Neg is called when production Expr_Neg is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Neg(ctx *Expr_NegContext) {}

// EnterExpr_MulDiv is called when production Expr_MulDiv is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_MulDiv(ctx *Expr_MulDivContext) {}

// ExitExpr_MulDiv is called when production Expr_MulDiv is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_MulDiv(ctx *Expr_MulDivContext) {}

// EnterExpr_Nil is called when production Expr_Nil is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Nil(ctx *Expr_NilContext) {}

// ExitExpr_Nil is called when production Expr_Nil is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Nil(ctx *Expr_NilContext) {}

// EnterExpr_Cadena is called when production Expr_Cadena is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Cadena(ctx *Expr_CadenaContext) {}

// ExitExpr_Cadena is called when production Expr_Cadena is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Cadena(ctx *Expr_CadenaContext) {}

// EnterExpr_Id is called when production Expr_Id is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Id(ctx *Expr_IdContext) {}

// ExitExpr_Id is called when production Expr_Id is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Id(ctx *Expr_IdContext) {}

// EnterExpr_Mod is called when production Expr_Mod is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Mod(ctx *Expr_ModContext) {}

// ExitExpr_Mod is called when production Expr_Mod is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Mod(ctx *Expr_ModContext) {}

// EnterExpr_Par is called when production Expr_Par is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Par(ctx *Expr_ParContext) {}

// ExitExpr_Par is called when production Expr_Par is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Par(ctx *Expr_ParContext) {}

// EnterExpr_Logica is called when production Expr_Logica is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Logica(ctx *Expr_LogicaContext) {}

// ExitExpr_Logica is called when production Expr_Logica is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Logica(ctx *Expr_LogicaContext) {}

// EnterExpr_Booleano is called when production Expr_Booleano is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Booleano(ctx *Expr_BooleanoContext) {}

// ExitExpr_Booleano is called when production Expr_Booleano is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Booleano(ctx *Expr_BooleanoContext) {}

// EnterExpr_Entero is called when production Expr_Entero is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Entero(ctx *Expr_EnteroContext) {}

// ExitExpr_Entero is called when production Expr_Entero is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Entero(ctx *Expr_EnteroContext) {}
