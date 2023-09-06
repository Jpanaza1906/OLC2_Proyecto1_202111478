// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Tswift_GrammarParser.
type Tswift_GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Tswift_GrammarParser#SLSentencias.
	VisitSLSentencias(ctx *SLSentenciasContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#L_Sentencia.
	VisitL_Sentencia(ctx *L_SentenciaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Print.
	VisitS_Print(ctx *S_PrintContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion.
	VisitS_Declaracion(ctx *S_DeclaracionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Constante.
	VisitS_Constante(ctx *S_ConstanteContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Asignacion.
	VisitS_Asignacion(ctx *S_AsignacionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_If.
	VisitS_If(ctx *S_IfContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Switch.
	VisitS_Switch(ctx *S_SwitchContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Guard.
	VisitS_Guard(ctx *S_GuardContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_While.
	VisitS_While(ctx *S_WhileContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_For.
	VisitS_For(ctx *S_ForContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Transicion.
	VisitS_Transicion(ctx *S_TransicionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion_Vector.
	VisitS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Funcion_Vector.
	VisitS_Funcion_Vector(ctx *S_Funcion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Asignacion_Vector.
	VisitS_Asignacion_Vector(ctx *S_Asignacion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Tipo_Val.
	VisitDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Val.
	VisitDeclaracion_Val(ctx *Declaracion_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Tipo_noVal.
	VisitDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Constante_Tipo_Val.
	VisitConstante_Tipo_Val(ctx *Constante_Tipo_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Constante_Val.
	VisitConstante_Val(ctx *Constante_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#SumAsg.
	VisitSumAsg(ctx *SumAsgContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ResAsg.
	VisitResAsg(ctx *ResAsgContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Asig.
	VisitAsig(ctx *AsigContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#If_Simple.
	VisitIf_Simple(ctx *If_SimpleContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#If_Else.
	VisitIf_Else(ctx *If_ElseContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#If_ElseIf.
	VisitIf_ElseIf(ctx *If_ElseIfContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#While.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#For.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Rango.
	VisitRango(ctx *RangoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Break.
	VisitBreak(ctx *BreakContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Continue.
	VisitContinue(ctx *ContinueContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Vector.
	VisitDeclaracion_Vector(ctx *Declaracion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Vector.
	VisitDef_Vector(ctx *Def_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Vector_Vacio.
	VisitDef_Vector_Vacio(ctx *Def_Vector_VacioContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Vector_Id.
	VisitDef_Vector_Id(ctx *Def_Vector_IdContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Asig_Vector.
	VisitAsig_Vector(ctx *Asig_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Append.
	VisitFunc_Vector_Append(ctx *Func_Vector_AppendContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_RemoveLast.
	VisitFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Remove.
	VisitFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Int.
	VisitTipo_Int(ctx *Tipo_IntContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Float.
	VisitTipo_Float(ctx *Tipo_FloatContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_String.
	VisitTipo_String(ctx *Tipo_StringContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Bool.
	VisitTipo_Bool(ctx *Tipo_BoolContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Character.
	VisitTipo_Character(ctx *Tipo_CharacterContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Rel.
	VisitExpr_Rel(ctx *Expr_RelContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Decimal.
	VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Caracter.
	VisitExpr_Caracter(ctx *Expr_CaracterContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_SumRes.
	VisitExpr_SumRes(ctx *Expr_SumResContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Neg.
	VisitExpr_Neg(ctx *Expr_NegContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_MulDiv.
	VisitExpr_MulDiv(ctx *Expr_MulDivContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Nil.
	VisitExpr_Nil(ctx *Expr_NilContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_isEmpty.
	VisitFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Count.
	VisitFunc_Vector_Count(ctx *Func_Vector_CountContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Cadena.
	VisitExpr_Cadena(ctx *Expr_CadenaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Id.
	VisitExpr_Id(ctx *Expr_IdContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Mod.
	VisitExpr_Mod(ctx *Expr_ModContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Par.
	VisitExpr_Par(ctx *Expr_ParContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Logica.
	VisitExpr_Logica(ctx *Expr_LogicaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Booleano.
	VisitExpr_Booleano(ctx *Expr_BooleanoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Vector.
	VisitExpr_Vector(ctx *Expr_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Entero.
	VisitExpr_Entero(ctx *Expr_EnteroContext) interface{}
}
