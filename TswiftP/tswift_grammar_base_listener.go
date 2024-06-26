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

// EnterS_Print is called when production S_Print is entered.
func (s *BaseTswift_GrammarListener) EnterS_Print(ctx *S_PrintContext) {}

// ExitS_Print is called when production S_Print is exited.
func (s *BaseTswift_GrammarListener) ExitS_Print(ctx *S_PrintContext) {}

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

// EnterS_Asignacion_Vector is called when production S_Asignacion_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterS_Asignacion_Vector(ctx *S_Asignacion_VectorContext) {}

// ExitS_Asignacion_Vector is called when production S_Asignacion_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitS_Asignacion_Vector(ctx *S_Asignacion_VectorContext) {}

// EnterS_Declaracion_Metodo is called when production S_Declaracion_Metodo is entered.
func (s *BaseTswift_GrammarListener) EnterS_Declaracion_Metodo(ctx *S_Declaracion_MetodoContext) {}

// ExitS_Declaracion_Metodo is called when production S_Declaracion_Metodo is exited.
func (s *BaseTswift_GrammarListener) ExitS_Declaracion_Metodo(ctx *S_Declaracion_MetodoContext) {}

// EnterS_Declaracion_Funcion is called when production S_Declaracion_Funcion is entered.
func (s *BaseTswift_GrammarListener) EnterS_Declaracion_Funcion(ctx *S_Declaracion_FuncionContext) {}

// ExitS_Declaracion_Funcion is called when production S_Declaracion_Funcion is exited.
func (s *BaseTswift_GrammarListener) ExitS_Declaracion_Funcion(ctx *S_Declaracion_FuncionContext) {}

// EnterS_Llamada_Funcion is called when production S_Llamada_Funcion is entered.
func (s *BaseTswift_GrammarListener) EnterS_Llamada_Funcion(ctx *S_Llamada_FuncionContext) {}

// ExitS_Llamada_Funcion is called when production S_Llamada_Funcion is exited.
func (s *BaseTswift_GrammarListener) ExitS_Llamada_Funcion(ctx *S_Llamada_FuncionContext) {}

// EnterS_Declaracion_Matriz is called when production S_Declaracion_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterS_Declaracion_Matriz(ctx *S_Declaracion_MatrizContext) {}

// ExitS_Declaracion_Matriz is called when production S_Declaracion_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitS_Declaracion_Matriz(ctx *S_Declaracion_MatrizContext) {}

// EnterS_Asignacion_Matriz is called when production S_Asignacion_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterS_Asignacion_Matriz(ctx *S_Asignacion_MatrizContext) {}

// ExitS_Asignacion_Matriz is called when production S_Asignacion_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitS_Asignacion_Matriz(ctx *S_Asignacion_MatrizContext) {}

// EnterS_Def_Struct is called when production S_Def_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterS_Def_Struct(ctx *S_Def_StructContext) {}

// ExitS_Def_Struct is called when production S_Def_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitS_Def_Struct(ctx *S_Def_StructContext) {}

// EnterS_Self_Data is called when production S_Self_Data is entered.
func (s *BaseTswift_GrammarListener) EnterS_Self_Data(ctx *S_Self_DataContext) {}

// ExitS_Self_Data is called when production S_Self_Data is exited.
func (s *BaseTswift_GrammarListener) ExitS_Self_Data(ctx *S_Self_DataContext) {}

// EnterS_Struct_Data is called when production S_Struct_Data is entered.
func (s *BaseTswift_GrammarListener) EnterS_Struct_Data(ctx *S_Struct_DataContext) {}

// ExitS_Struct_Data is called when production S_Struct_Data is exited.
func (s *BaseTswift_GrammarListener) ExitS_Struct_Data(ctx *S_Struct_DataContext) {}

// EnterS_Struct_Llamada_Func is called when production S_Struct_Llamada_Func is entered.
func (s *BaseTswift_GrammarListener) EnterS_Struct_Llamada_Func(ctx *S_Struct_Llamada_FuncContext) {}

// ExitS_Struct_Llamada_Func is called when production S_Struct_Llamada_Func is exited.
func (s *BaseTswift_GrammarListener) ExitS_Struct_Llamada_Func(ctx *S_Struct_Llamada_FuncContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseTswift_GrammarListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseTswift_GrammarListener) ExitPrint(ctx *PrintContext) {}

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

// EnterAsig_Vector is called when production Asig_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterAsig_Vector(ctx *Asig_VectorContext) {}

// ExitAsig_Vector is called when production Asig_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitAsig_Vector(ctx *Asig_VectorContext) {}

// EnterSumAsg_Vector is called when production SumAsg_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterSumAsg_Vector(ctx *SumAsg_VectorContext) {}

// ExitSumAsg_Vector is called when production SumAsg_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitSumAsg_Vector(ctx *SumAsg_VectorContext) {}

// EnterResAsg_Vector is called when production ResAsg_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterResAsg_Vector(ctx *ResAsg_VectorContext) {}

// ExitResAsg_Vector is called when production ResAsg_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitResAsg_Vector(ctx *ResAsg_VectorContext) {}

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

// EnterDeclaracion_Matriz is called when production Declaracion_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Matriz(ctx *Declaracion_MatrizContext) {}

// ExitDeclaracion_Matriz is called when production Declaracion_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Matriz(ctx *Declaracion_MatrizContext) {}

// EnterTipo_Matriz is called when production Tipo_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Matriz(ctx *Tipo_MatrizContext) {}

// ExitTipo_Matriz is called when production Tipo_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Matriz(ctx *Tipo_MatrizContext) {}

// EnterTipo_Matriz_Simple is called when production Tipo_Matriz_Simple is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Matriz_Simple(ctx *Tipo_Matriz_SimpleContext) {}

// ExitTipo_Matriz_Simple is called when production Tipo_Matriz_Simple is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Matriz_Simple(ctx *Tipo_Matriz_SimpleContext) {}

// EnterDef_Matriz is called when production Def_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Matriz(ctx *Def_MatrizContext) {}

// ExitDef_Matriz is called when production Def_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Matriz(ctx *Def_MatrizContext) {}

// EnterDef_Matriz_Simple is called when production Def_Matriz_Simple is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Matriz_Simple(ctx *Def_Matriz_SimpleContext) {}

// ExitDef_Matriz_Simple is called when production Def_Matriz_Simple is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Matriz_Simple(ctx *Def_Matriz_SimpleContext) {}

// EnterListaCompletaVal is called when production ListaCompletaVal is entered.
func (s *BaseTswift_GrammarListener) EnterListaCompletaVal(ctx *ListaCompletaValContext) {}

// ExitListaCompletaVal is called when production ListaCompletaVal is exited.
func (s *BaseTswift_GrammarListener) ExitListaCompletaVal(ctx *ListaCompletaValContext) {}

// EnterListaValorSig is called when production ListaValorSig is entered.
func (s *BaseTswift_GrammarListener) EnterListaValorSig(ctx *ListaValorSigContext) {}

// ExitListaValorSig is called when production ListaValorSig is exited.
func (s *BaseTswift_GrammarListener) ExitListaValorSig(ctx *ListaValorSigContext) {}

// EnterListaValoresHermanos is called when production ListaValoresHermanos is entered.
func (s *BaseTswift_GrammarListener) EnterListaValoresHermanos(ctx *ListaValoresHermanosContext) {}

// ExitListaValoresHermanos is called when production ListaValoresHermanos is exited.
func (s *BaseTswift_GrammarListener) ExitListaValoresHermanos(ctx *ListaValoresHermanosContext) {}

// EnterListaExpr is called when production ListaExpr is entered.
func (s *BaseTswift_GrammarListener) EnterListaExpr(ctx *ListaExprContext) {}

// ExitListaExpr is called when production ListaExpr is exited.
func (s *BaseTswift_GrammarListener) ExitListaExpr(ctx *ListaExprContext) {}

// EnterSimple_Mat is called when production Simple_Mat is entered.
func (s *BaseTswift_GrammarListener) EnterSimple_Mat(ctx *Simple_MatContext) {}

// ExitSimple_Mat is called when production Simple_Mat is exited.
func (s *BaseTswift_GrammarListener) ExitSimple_Mat(ctx *Simple_MatContext) {}

// EnterSimple_Mat_Expr is called when production Simple_Mat_Expr is entered.
func (s *BaseTswift_GrammarListener) EnterSimple_Mat_Expr(ctx *Simple_Mat_ExprContext) {}

// ExitSimple_Mat_Expr is called when production Simple_Mat_Expr is exited.
func (s *BaseTswift_GrammarListener) ExitSimple_Mat_Expr(ctx *Simple_Mat_ExprContext) {}

// EnterAsig_Mat is called when production Asig_Mat is entered.
func (s *BaseTswift_GrammarListener) EnterAsig_Mat(ctx *Asig_MatContext) {}

// ExitAsig_Mat is called when production Asig_Mat is exited.
func (s *BaseTswift_GrammarListener) ExitAsig_Mat(ctx *Asig_MatContext) {}

// EnterDeclaracion_Metodo is called when production Declaracion_Metodo is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Metodo(ctx *Declaracion_MetodoContext) {}

// ExitDeclaracion_Metodo is called when production Declaracion_Metodo is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Metodo(ctx *Declaracion_MetodoContext) {}

// EnterDeclaracion_Funcion is called when production Declaracion_Funcion is entered.
func (s *BaseTswift_GrammarListener) EnterDeclaracion_Funcion(ctx *Declaracion_FuncionContext) {}

// ExitDeclaracion_Funcion is called when production Declaracion_Funcion is exited.
func (s *BaseTswift_GrammarListener) ExitDeclaracion_Funcion(ctx *Declaracion_FuncionContext) {}

// EnterL_Parametros is called when production L_Parametros is entered.
func (s *BaseTswift_GrammarListener) EnterL_Parametros(ctx *L_ParametrosContext) {}

// ExitL_Parametros is called when production L_Parametros is exited.
func (s *BaseTswift_GrammarListener) ExitL_Parametros(ctx *L_ParametrosContext) {}

// EnterLlamada_Funcion is called when production Llamada_Funcion is entered.
func (s *BaseTswift_GrammarListener) EnterLlamada_Funcion(ctx *Llamada_FuncionContext) {}

// ExitLlamada_Funcion is called when production Llamada_Funcion is exited.
func (s *BaseTswift_GrammarListener) ExitLlamada_Funcion(ctx *Llamada_FuncionContext) {}

// EnterL_Argumentos is called when production L_Argumentos is entered.
func (s *BaseTswift_GrammarListener) EnterL_Argumentos(ctx *L_ArgumentosContext) {}

// ExitL_Argumentos is called when production L_Argumentos is exited.
func (s *BaseTswift_GrammarListener) ExitL_Argumentos(ctx *L_ArgumentosContext) {}

// EnterFunc_Int is called when production Func_Int is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Int(ctx *Func_IntContext) {}

// ExitFunc_Int is called when production Func_Int is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Int(ctx *Func_IntContext) {}

// EnterFunc_Float is called when production Func_Float is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Float(ctx *Func_FloatContext) {}

// ExitFunc_Float is called when production Func_Float is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Float(ctx *Func_FloatContext) {}

// EnterFunc_String is called when production Func_String is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_String(ctx *Func_StringContext) {}

// ExitFunc_String is called when production Func_String is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_String(ctx *Func_StringContext) {}

// EnterDef_Struct is called when production Def_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterDef_Struct(ctx *Def_StructContext) {}

// ExitDef_Struct is called when production Def_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitDef_Struct(ctx *Def_StructContext) {}

// EnterL_Atributos is called when production L_Atributos is entered.
func (s *BaseTswift_GrammarListener) EnterL_Atributos(ctx *L_AtributosContext) {}

// ExitL_Atributos is called when production L_Atributos is exited.
func (s *BaseTswift_GrammarListener) ExitL_Atributos(ctx *L_AtributosContext) {}

// EnterL_Funciones is called when production L_Funciones is entered.
func (s *BaseTswift_GrammarListener) EnterL_Funciones(ctx *L_FuncionesContext) {}

// ExitL_Funciones is called when production L_Funciones is exited.
func (s *BaseTswift_GrammarListener) ExitL_Funciones(ctx *L_FuncionesContext) {}

// EnterSelf_Data is called when production Self_Data is entered.
func (s *BaseTswift_GrammarListener) EnterSelf_Data(ctx *Self_DataContext) {}

// ExitSelf_Data is called when production Self_Data is exited.
func (s *BaseTswift_GrammarListener) ExitSelf_Data(ctx *Self_DataContext) {}

// EnterStruct_Data is called when production Struct_Data is entered.
func (s *BaseTswift_GrammarListener) EnterStruct_Data(ctx *Struct_DataContext) {}

// ExitStruct_Data is called when production Struct_Data is exited.
func (s *BaseTswift_GrammarListener) ExitStruct_Data(ctx *Struct_DataContext) {}

// EnterId_Struct is called when production Id_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterId_Struct(ctx *Id_StructContext) {}

// ExitId_Struct is called when production Id_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitId_Struct(ctx *Id_StructContext) {}

// EnterStruct_Llamada_Func is called when production Struct_Llamada_Func is entered.
func (s *BaseTswift_GrammarListener) EnterStruct_Llamada_Func(ctx *Struct_Llamada_FuncContext) {}

// ExitStruct_Llamada_Func is called when production Struct_Llamada_Func is exited.
func (s *BaseTswift_GrammarListener) ExitStruct_Llamada_Func(ctx *Struct_Llamada_FuncContext) {}

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

// EnterTipo_Struct is called when production Tipo_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Struct(ctx *Tipo_StructContext) {}

// ExitTipo_Struct is called when production Tipo_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Struct(ctx *Tipo_StructContext) {}

// EnterTipo_Vector is called when production Tipo_Vector is entered.
func (s *BaseTswift_GrammarListener) EnterTipo_Vector(ctx *Tipo_VectorContext) {}

// ExitTipo_Vector is called when production Tipo_Vector is exited.
func (s *BaseTswift_GrammarListener) ExitTipo_Vector(ctx *Tipo_VectorContext) {}

// EnterExpr_Rel is called when production Expr_Rel is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Rel(ctx *Expr_RelContext) {}

// ExitExpr_Rel is called when production Expr_Rel is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Rel(ctx *Expr_RelContext) {}

// EnterExpr_Funciones_Embebidas is called when production Expr_Funciones_Embebidas is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Funciones_Embebidas(ctx *Expr_Funciones_EmbebidasContext) {
}

// ExitExpr_Funciones_Embebidas is called when production Expr_Funciones_Embebidas is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Funciones_Embebidas(ctx *Expr_Funciones_EmbebidasContext) {
}

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

// EnterExpr_Llamada_Funcion_Struct is called when production Expr_Llamada_Funcion_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Llamada_Funcion_Struct(ctx *Expr_Llamada_Funcion_StructContext) {
}

// ExitExpr_Llamada_Funcion_Struct is called when production Expr_Llamada_Funcion_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Llamada_Funcion_Struct(ctx *Expr_Llamada_Funcion_StructContext) {
}

// EnterExpr_Matriz is called when production Expr_Matriz is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Matriz(ctx *Expr_MatrizContext) {}

// ExitExpr_Matriz is called when production Expr_Matriz is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Matriz(ctx *Expr_MatrizContext) {}

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

// EnterFunc_Vector_isEmpty is called when production Func_Vector_isEmpty is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) {}

// ExitFunc_Vector_isEmpty is called when production Func_Vector_isEmpty is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) {}

// EnterExpr_Llamada_Funcion is called when production Expr_Llamada_Funcion is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Llamada_Funcion(ctx *Expr_Llamada_FuncionContext) {}

// ExitExpr_Llamada_Funcion is called when production Expr_Llamada_Funcion is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Llamada_Funcion(ctx *Expr_Llamada_FuncionContext) {}

// EnterFunc_Vector_Count is called when production Func_Vector_Count is entered.
func (s *BaseTswift_GrammarListener) EnterFunc_Vector_Count(ctx *Func_Vector_CountContext) {}

// ExitFunc_Vector_Count is called when production Func_Vector_Count is exited.
func (s *BaseTswift_GrammarListener) ExitFunc_Vector_Count(ctx *Func_Vector_CountContext) {}

// EnterExpr_Struct is called when production Expr_Struct is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Struct(ctx *Expr_StructContext) {}

// ExitExpr_Struct is called when production Expr_Struct is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Struct(ctx *Expr_StructContext) {}

// EnterExpr_Cadena is called when production Expr_Cadena is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Cadena(ctx *Expr_CadenaContext) {}

// ExitExpr_Cadena is called when production Expr_Cadena is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Cadena(ctx *Expr_CadenaContext) {}

// EnterExpr_Self is called when production Expr_Self is entered.
func (s *BaseTswift_GrammarListener) EnterExpr_Self(ctx *Expr_SelfContext) {}

// ExitExpr_Self is called when production Expr_Self is exited.
func (s *BaseTswift_GrammarListener) ExitExpr_Self(ctx *Expr_SelfContext) {}

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
