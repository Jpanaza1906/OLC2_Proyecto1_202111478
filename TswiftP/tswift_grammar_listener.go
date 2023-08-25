// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

// Tswift_GrammarListener is a complete listener for a parse tree produced by Tswift_GrammarParser.
type Tswift_GrammarListener interface {
	antlr.ParseTreeListener

	// EnterSLSentencias is called when entering the SLSentencias production.
	EnterSLSentencias(c *SLSentenciasContext)

	// EnterL_Sentencia is called when entering the L_Sentencia production.
	EnterL_Sentencia(c *L_SentenciaContext)

	// EnterS_Consola is called when entering the S_Consola production.
	EnterS_Consola(c *S_ConsolaContext)

	// EnterS_Declaracion is called when entering the S_Declaracion production.
	EnterS_Declaracion(c *S_DeclaracionContext)

	// EnterS_Constante is called when entering the S_Constante production.
	EnterS_Constante(c *S_ConstanteContext)

	// EnterDeclaracion_Tipo_Val is called when entering the Declaracion_Tipo_Val production.
	EnterDeclaracion_Tipo_Val(c *Declaracion_Tipo_ValContext)

	// EnterDeclaracion_Val is called when entering the Declaracion_Val production.
	EnterDeclaracion_Val(c *Declaracion_ValContext)

	// EnterDeclaracion_Tipo_noVal is called when entering the Declaracion_Tipo_noVal production.
	EnterDeclaracion_Tipo_noVal(c *Declaracion_Tipo_noValContext)

	// EnterConstante_Tipo_Val is called when entering the Constante_Tipo_Val production.
	EnterConstante_Tipo_Val(c *Constante_Tipo_ValContext)

	// EnterConstante_Val is called when entering the Constante_Val production.
	EnterConstante_Val(c *Constante_ValContext)

	// EnterConstante_Tipo_noVal is called when entering the Constante_Tipo_noVal production.
	EnterConstante_Tipo_noVal(c *Constante_Tipo_noValContext)

	// EnterTipo_Int is called when entering the Tipo_Int production.
	EnterTipo_Int(c *Tipo_IntContext)

	// EnterTipo_Float is called when entering the Tipo_Float production.
	EnterTipo_Float(c *Tipo_FloatContext)

	// EnterTipo_String is called when entering the Tipo_String production.
	EnterTipo_String(c *Tipo_StringContext)

	// EnterTipo_Bool is called when entering the Tipo_Bool production.
	EnterTipo_Bool(c *Tipo_BoolContext)

	// EnterTipo_Character is called when entering the Tipo_Character production.
	EnterTipo_Character(c *Tipo_CharacterContext)

	// EnterExpr_Rel is called when entering the Expr_Rel production.
	EnterExpr_Rel(c *Expr_RelContext)

	// EnterExpr_Decimal is called when entering the Expr_Decimal production.
	EnterExpr_Decimal(c *Expr_DecimalContext)

	// EnterExpr_Caracter is called when entering the Expr_Caracter production.
	EnterExpr_Caracter(c *Expr_CaracterContext)

	// EnterExpr_SumRes is called when entering the Expr_SumRes production.
	EnterExpr_SumRes(c *Expr_SumResContext)

	// EnterExpr_Neg is called when entering the Expr_Neg production.
	EnterExpr_Neg(c *Expr_NegContext)

	// EnterExpr_MulDiv is called when entering the Expr_MulDiv production.
	EnterExpr_MulDiv(c *Expr_MulDivContext)

	// EnterExpr_Nil is called when entering the Expr_Nil production.
	EnterExpr_Nil(c *Expr_NilContext)

	// EnterExpr_Cadena is called when entering the Expr_Cadena production.
	EnterExpr_Cadena(c *Expr_CadenaContext)

	// EnterExpr_Id is called when entering the Expr_Id production.
	EnterExpr_Id(c *Expr_IdContext)

	// EnterExpr_Mod is called when entering the Expr_Mod production.
	EnterExpr_Mod(c *Expr_ModContext)

	// EnterExpr_Par is called when entering the Expr_Par production.
	EnterExpr_Par(c *Expr_ParContext)

	// EnterExpr_Logica is called when entering the Expr_Logica production.
	EnterExpr_Logica(c *Expr_LogicaContext)

	// EnterExpr_Booleano is called when entering the Expr_Booleano production.
	EnterExpr_Booleano(c *Expr_BooleanoContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitS_Consola is called when exiting the S_Consola production.
	ExitS_Consola(c *S_ConsolaContext)

	// ExitS_Declaracion is called when exiting the S_Declaracion production.
	ExitS_Declaracion(c *S_DeclaracionContext)

	// ExitS_Constante is called when exiting the S_Constante production.
	ExitS_Constante(c *S_ConstanteContext)

	// ExitDeclaracion_Tipo_Val is called when exiting the Declaracion_Tipo_Val production.
	ExitDeclaracion_Tipo_Val(c *Declaracion_Tipo_ValContext)

	// ExitDeclaracion_Val is called when exiting the Declaracion_Val production.
	ExitDeclaracion_Val(c *Declaracion_ValContext)

	// ExitDeclaracion_Tipo_noVal is called when exiting the Declaracion_Tipo_noVal production.
	ExitDeclaracion_Tipo_noVal(c *Declaracion_Tipo_noValContext)

	// ExitConstante_Tipo_Val is called when exiting the Constante_Tipo_Val production.
	ExitConstante_Tipo_Val(c *Constante_Tipo_ValContext)

	// ExitConstante_Val is called when exiting the Constante_Val production.
	ExitConstante_Val(c *Constante_ValContext)

	// ExitConstante_Tipo_noVal is called when exiting the Constante_Tipo_noVal production.
	ExitConstante_Tipo_noVal(c *Constante_Tipo_noValContext)

	// ExitTipo_Int is called when exiting the Tipo_Int production.
	ExitTipo_Int(c *Tipo_IntContext)

	// ExitTipo_Float is called when exiting the Tipo_Float production.
	ExitTipo_Float(c *Tipo_FloatContext)

	// ExitTipo_String is called when exiting the Tipo_String production.
	ExitTipo_String(c *Tipo_StringContext)

	// ExitTipo_Bool is called when exiting the Tipo_Bool production.
	ExitTipo_Bool(c *Tipo_BoolContext)

	// ExitTipo_Character is called when exiting the Tipo_Character production.
	ExitTipo_Character(c *Tipo_CharacterContext)

	// ExitExpr_Rel is called when exiting the Expr_Rel production.
	ExitExpr_Rel(c *Expr_RelContext)

	// ExitExpr_Decimal is called when exiting the Expr_Decimal production.
	ExitExpr_Decimal(c *Expr_DecimalContext)

	// ExitExpr_Caracter is called when exiting the Expr_Caracter production.
	ExitExpr_Caracter(c *Expr_CaracterContext)

	// ExitExpr_SumRes is called when exiting the Expr_SumRes production.
	ExitExpr_SumRes(c *Expr_SumResContext)

	// ExitExpr_Neg is called when exiting the Expr_Neg production.
	ExitExpr_Neg(c *Expr_NegContext)

	// ExitExpr_MulDiv is called when exiting the Expr_MulDiv production.
	ExitExpr_MulDiv(c *Expr_MulDivContext)

	// ExitExpr_Nil is called when exiting the Expr_Nil production.
	ExitExpr_Nil(c *Expr_NilContext)

	// ExitExpr_Cadena is called when exiting the Expr_Cadena production.
	ExitExpr_Cadena(c *Expr_CadenaContext)

	// ExitExpr_Id is called when exiting the Expr_Id production.
	ExitExpr_Id(c *Expr_IdContext)

	// ExitExpr_Mod is called when exiting the Expr_Mod production.
	ExitExpr_Mod(c *Expr_ModContext)

	// ExitExpr_Par is called when exiting the Expr_Par production.
	ExitExpr_Par(c *Expr_ParContext)

	// ExitExpr_Logica is called when exiting the Expr_Logica production.
	ExitExpr_Logica(c *Expr_LogicaContext)

	// ExitExpr_Booleano is called when exiting the Expr_Booleano production.
	ExitExpr_Booleano(c *Expr_BooleanoContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)
}
