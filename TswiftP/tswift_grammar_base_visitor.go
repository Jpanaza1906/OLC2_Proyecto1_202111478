// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

type BaseTswift_GrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTswift_GrammarVisitor) VisitSLSentencias(ctx *SLSentenciasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitL_Sentencia(ctx *L_SentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Consola(ctx *S_ConsolaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Declaracion(ctx *S_DeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Constante(ctx *S_ConstanteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Asignacion(ctx *S_AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_If(ctx *S_IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Switch(ctx *S_SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Guard(ctx *S_GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_While(ctx *S_WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_For(ctx *S_ForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitS_Transicion(ctx *S_TransicionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitDeclaracion_Val(ctx *Declaracion_ValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitConstante_Tipo_Val(ctx *Constante_Tipo_ValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitConstante_Val(ctx *Constante_ValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitSumAsg(ctx *SumAsgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitResAsg(ctx *ResAsgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitAsig(ctx *AsigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitIf_Simple(ctx *If_SimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitIf_Else(ctx *If_ElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitIf_ElseIf(ctx *If_ElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitSwitch(ctx *SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitCase(ctx *CaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitDefault(ctx *DefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitGuard(ctx *GuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitFor(ctx *ForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitRango(ctx *RangoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitBreak(ctx *BreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitContinue(ctx *ContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitReturn(ctx *ReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitTipo_Int(ctx *Tipo_IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitTipo_Float(ctx *Tipo_FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitTipo_String(ctx *Tipo_StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitTipo_Bool(ctx *Tipo_BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitTipo_Character(ctx *Tipo_CharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Rel(ctx *Expr_RelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Caracter(ctx *Expr_CaracterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_SumRes(ctx *Expr_SumResContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Neg(ctx *Expr_NegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_MulDiv(ctx *Expr_MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Nil(ctx *Expr_NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Cadena(ctx *Expr_CadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Id(ctx *Expr_IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Mod(ctx *Expr_ModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Par(ctx *Expr_ParContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Logica(ctx *Expr_LogicaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Booleano(ctx *Expr_BooleanoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTswift_GrammarVisitor) VisitExpr_Entero(ctx *Expr_EnteroContext) interface{} {
	return v.VisitChildren(ctx)
}
