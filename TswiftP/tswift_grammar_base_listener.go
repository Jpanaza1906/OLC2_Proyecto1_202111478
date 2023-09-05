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

// EnterS_Constante is called when production S_Constante is entered.
func (s *BaseTswift_GrammarListener) EnterS_Constante(ctx *S_ConstanteContext) {}

// ExitS_Constante is called when production S_Constante is exited.
func (s *BaseTswift_GrammarListener) ExitS_Constante(ctx *S_ConstanteContext) {}

// EnterS_Asignacion is called when production S_Asignacion is entered.
func (s *BaseTswift_GrammarListener) EnterS_Asignacion(ctx *S_AsignacionContext) {}

// ExitS_Asignacion is called when production S_Asignacion is exited.
func (s *BaseTswift_GrammarListener) ExitS_Asignacion(ctx *S_AsignacionContext) {}

// EnterS_If is called when production S_If is entered.
func (s *BaseTswift_GrammarListener) EnterS_If(ctx *S_IfContext) {}

// ExitS_If is called when production S_If is exited.
func (s *BaseTswift_GrammarListener) ExitS_If(ctx *S_IfContext) {}

// EnterS_Switch is called when production S_Switch is entered.
func (s *BaseTswift_GrammarListener) EnterS_Switch(ctx *S_SwitchContext) {}

// ExitS_Switch is called when production S_Switch is exited.
func (s *BaseTswift_GrammarListener) ExitS_Switch(ctx *S_SwitchContext) {}

// EnterS_Guard is called when production S_Guard is entered.
func (s *BaseTswift_GrammarListener) EnterS_Guard(ctx *S_GuardContext) {}

// ExitS_Guard is called when production S_Guard is exited.
func (s *BaseTswift_GrammarListener) ExitS_Guard(ctx *S_GuardContext) {}

// EnterS_While is called when production S_While is entered.
func (s *BaseTswift_GrammarListener) EnterS_While(ctx *S_WhileContext) {}

// ExitS_While is called when production S_While is exited.
func (s *BaseTswift_GrammarListener) ExitS_While(ctx *S_WhileContext) {}

// EnterS_For is called when production S_For is entered.
func (s *BaseTswift_GrammarListener) EnterS_For(ctx *S_ForContext) {}

// ExitS_For is called when production S_For is exited.
func (s *BaseTswift_GrammarListener) ExitS_For(ctx *S_ForContext) {}

// EnterS_Transicion is called when production S_Transicion is entered.
func (s *BaseTswift_GrammarListener) EnterS_Transicion(ctx *S_TransicionContext) {}

// ExitS_Transicion is called when production S_Transicion is exited.
func (s *BaseTswift_GrammarListener) ExitS_Transicion(ctx *S_TransicionContext) {}

// EnterS_Declaracion_Vector is called when production S_Declaracion_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) {}

// ExitS_Declaracion_Vector is called when production S_Declaracion_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) {}

// EnterS_Funcion_Vector is called when production S_Funcion_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterS_Funcion_Vector(ctx *S_Funcion_VectorContext) {}

// ExitS_Funcion_Vector is called when production S_Funcion_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitS_Funcion_Vector(ctx *S_Funcion_VectorContext) {}

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

// EnterConstante_Tipo_Val is called when production Constante_Tipo_Val is entered.
func (s *BaseTswift_GrammarListener) EnterConstante_Tipo_Val(ctx *Constante_Tipo_ValContext) {}

// ExitConstante_Tipo_Val is called when production Constante_Tipo_Val is exited.
func (s *BaseTswift_GrammarListener) ExitConstante_Tipo_Val(ctx *Constante_Tipo_ValContext) {}

// EnterConstante_Val is called when production Constante_Val is entered.
func (s *BaseTswift_GrammarListener) EnterConstante_Val(ctx *Constante_ValContext) {}

// ExitConstante_Val is called when production Constante_Val is exited.
func (s *BaseTswift_GrammarListener) ExitConstante_Val(ctx *Constante_ValContext) {}

// EnterSumAsg is called when production SumAsg is entered.
func (s *BaseTswift_GrammarListener) EnterSumAsg(ctx *SumAsgContext) {}

// ExitSumAsg is called when production SumAsg is exited.
func (s *BaseTswift_GrammarListener) ExitSumAsg(ctx *SumAsgContext) {}

// EnterResAsg is called when production ResAsg is entered.
func (s *BaseTswift_GrammarListener) EnterResAsg(ctx *ResAsgContext) {}

// ExitResAsg is called when production ResAsg is exited.
func (s *BaseTswift_GrammarListener) ExitResAsg(ctx *ResAsgContext) {}

// EnterAsig is called when production Asig is entered.
func (s *BaseTswift_GrammarListener) EnterAsig(ctx *AsigContext) {}

// ExitAsig is called when production Asig is exited.
func (s *BaseTswift_GrammarListener) ExitAsig(ctx *AsigContext) {}

// EnterIf_Simple is called when production If_Simple is entered.
func (s *BaseTswift_GrammarListener) EnterIf_Simple(ctx *If_SimpleContext) {}

// ExitIf_Simple is called when production If_Simple is exited.
func (s *BaseTswift_GrammarListener) ExitIf_Simple(ctx *If_SimpleContext) {}

// EnterIf_Else is called when production If_Else is entered.
func (s *BaseTswift_GrammarListener) EnterIf_Else(ctx *If_ElseContext) {}

// ExitIf_Else is called when production If_Else is exited.
func (s *BaseTswift_GrammarListener) ExitIf_Else(ctx *If_ElseContext) {}

// EnterIf_ElseIf is called when production If_ElseIf is entered.
func (s *BaseTswift_GrammarListener) EnterIf_ElseIf(ctx *If_ElseIfContext) {}

// ExitIf_ElseIf is called when production If_ElseIf is exited.
func (s *BaseTswift_GrammarListener) ExitIf_ElseIf(ctx *If_ElseIfContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BaseTswift_GrammarListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BaseTswift_GrammarListener) ExitSwitch(ctx *SwitchContext) {}

// EnterCase is called when production Case is entered.
func (s *BaseTswift_GrammarListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production Case is exited.
func (s *BaseTswift_GrammarListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production Default is entered.
func (s *BaseTswift_GrammarListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production Default is exited.
func (s *BaseTswift_GrammarListener) ExitDefault(ctx *DefaultContext) {}

// EnterGuard is called when production Guard is entered.
func (s *BaseTswift_GrammarListener) EnterGuard(ctx *GuardContext) {}

// ExitGuard is called when production Guard is exited.
func (s *BaseTswift_GrammarListener) ExitGuard(ctx *GuardContext) {}

// EnterWhile is called when production While is entered.
func (s *BaseTswift_GrammarListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production While is exited.
func (s *BaseTswift_GrammarListener) ExitWhile(ctx *WhileContext) {}

// EnterFor is called when production For is entered.
func (s *BaseTswift_GrammarListener) EnterFor(ctx *ForContext) {}

// ExitFor is called when production For is exited.
func (s *BaseTswift_GrammarListener) ExitFor(ctx *ForContext) {}

// EnterRango is called when production Rango is entered.
func (s *BaseTswift_GrammarListener) EnterRango(ctx *RangoContext) {}

// ExitRango is called when production Rango is exited.
func (s *BaseTswift_GrammarListener) ExitRango(ctx *RangoContext) {}

// EnterBreak is called when production Break is entered.
func (s *BaseTswift_GrammarListener) EnterBreak(ctx *BreakContext) {}

// ExitBreak is called when production Break is exited.
func (s *BaseTswift_GrammarListener) ExitBreak(ctx *BreakContext) {}

// EnterContinue is called when production Continue is entered.
func (s *BaseTswift_GrammarListener) EnterContinue(ctx *ContinueContext) {}

// ExitContinue is called when production Continue is exited.
func (s *BaseTswift_GrammarListener) ExitContinue(ctx *ContinueContext) {}

// EnterReturn is called when production Return is entered.
func (s *BaseTswift_GrammarListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production Return is exited.
func (s *BaseTswift_GrammarListener) ExitReturn(ctx *ReturnContext) {}

// EnterDeclaracion_Vector is called when production Declaracion_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Vector(ctx *Declaracion_VectorContext) {}

// ExitDeclaracion_Vector is called when production Declaracion_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Vector(ctx *Declaracion_VectorContext) {}

// EnterDef_Vector is called when production Def_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Vector(ctx *Def_VectorContext) {}

// ExitDef_Vector is called when production Def_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Vector(ctx *Def_VectorContext) {}

// EnterDef_Vector_Vacio is called when production Def_Vector_Vacio is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Vector_Vacio(ctx *Def_Vector_VacioContext) {}

// ExitDef_Vector_Vacio is called when production Def_Vector_Vacio is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Vector_Vacio(ctx *Def_Vector_VacioContext) {}

// EnterDef_Vector_Id is called when production Def_Vector_Id is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Vector_Id(ctx *Def_Vector_IdContext) {}

// ExitDef_Vector_Id is called when production Def_Vector_Id is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Vector_Id(ctx *Def_Vector_IdContext) {}

// EnterFunc_Vector_Append is called when production Func_Vector_Append is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_Append(ctx *Func_Vector_AppendContext) {}

// ExitFunc_Vector_Append is called when production Func_Vector_Append is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_Append(ctx *Func_Vector_AppendContext) {}

// EnterFunc_Vector_RemoveLast is called when production Func_Vector_RemoveLast is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) {
}

// ExitFunc_Vector_RemoveLast is called when production Func_Vector_RemoveLast is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) {}

// EnterFunc_Vector_Remove is called when production Func_Vector_Remove is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) {}

// ExitFunc_Vector_Remove is called when production Func_Vector_Remove is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) {}

// EnterFunc_Vector_isEmpty is called when production Func_Vector_isEmpty is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) {}

// ExitFunc_Vector_isEmpty is called when production Func_Vector_isEmpty is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) {}

// EnterFunc_Vector_Count is called when production Func_Vector_Count is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_Count(ctx *Func_Vector_CountContext) {}

// ExitFunc_Vector_Count is called when production Func_Vector_Count is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_Count(ctx *Func_Vector_CountContext) {}

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

// EnterExpr_Vector is called when production Expr_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Vector(ctx *Expr_VectorContext) {}

// ExitExpr_Vector is called when production Expr_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Vector(ctx *Expr_VectorContext) {}

// EnterExpr_Entero is called when production Expr_Entero is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Entero(ctx *Expr_EnteroContext) {}

// ExitExpr_Entero is called when production Expr_Entero is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Entero(ctx *Expr_EnteroContext) {}
