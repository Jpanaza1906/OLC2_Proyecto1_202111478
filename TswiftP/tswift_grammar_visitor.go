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

	// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion_Metodo.
	VisitS_Declaracion_Metodo(ctx *S_Declaracion_MetodoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion_Funcion.
	VisitS_Declaracion_Funcion(ctx *S_Declaracion_FuncionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Llamada_Funcion.
	VisitS_Llamada_Funcion(ctx *S_Llamada_FuncionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion_Matriz.
	VisitS_Declaracion_Matriz(ctx *S_Declaracion_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Asignacion_Matriz.
	VisitS_Asignacion_Matriz(ctx *S_Asignacion_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Def_Struct.
	VisitS_Def_Struct(ctx *S_Def_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Self_Data.
	VisitS_Self_Data(ctx *S_Self_DataContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Struct_Data.
	VisitS_Struct_Data(ctx *S_Struct_DataContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#S_Struct_Llamada_Func.
	VisitS_Struct_Llamada_Func(ctx *S_Struct_Llamada_FuncContext) interface{}

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

	// Visit a parse tree produced by Tswift_GrammarParser#SumAsg_Vector.
	VisitSumAsg_Vector(ctx *SumAsg_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ResAsg_Vector.
	VisitResAsg_Vector(ctx *ResAsg_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Append.
	VisitFunc_Vector_Append(ctx *Func_Vector_AppendContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_RemoveLast.
	VisitFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Remove.
	VisitFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Matriz.
	VisitDeclaracion_Matriz(ctx *Declaracion_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Matriz.
	VisitTipo_Matriz(ctx *Tipo_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Matriz_Simple.
	VisitTipo_Matriz_Simple(ctx *Tipo_Matriz_SimpleContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Matriz.
	VisitDef_Matriz(ctx *Def_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Matriz_Simple.
	VisitDef_Matriz_Simple(ctx *Def_Matriz_SimpleContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ListaCompletaVal.
	VisitListaCompletaVal(ctx *ListaCompletaValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ListaValorSig.
	VisitListaValorSig(ctx *ListaValorSigContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ListaValoresHermanos.
	VisitListaValoresHermanos(ctx *ListaValoresHermanosContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ListaExpr.
	VisitListaExpr(ctx *ListaExprContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Simple_Mat.
	VisitSimple_Mat(ctx *Simple_MatContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Simple_Mat_Expr.
	VisitSimple_Mat_Expr(ctx *Simple_Mat_ExprContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Asig_Mat.
	VisitAsig_Mat(ctx *Asig_MatContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Metodo.
	VisitDeclaracion_Metodo(ctx *Declaracion_MetodoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Funcion.
	VisitDeclaracion_Funcion(ctx *Declaracion_FuncionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#L_Parametros.
	VisitL_Parametros(ctx *L_ParametrosContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Llamada_Funcion.
	VisitLlamada_Funcion(ctx *Llamada_FuncionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#L_Argumentos.
	VisitL_Argumentos(ctx *L_ArgumentosContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Int.
	VisitFunc_Int(ctx *Func_IntContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Float.
	VisitFunc_Float(ctx *Func_FloatContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_String.
	VisitFunc_String(ctx *Func_StringContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Def_Struct.
	VisitDef_Struct(ctx *Def_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#L_Atributos.
	VisitL_Atributos(ctx *L_AtributosContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#L_Funciones.
	VisitL_Funciones(ctx *L_FuncionesContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Self_Data.
	VisitSelf_Data(ctx *Self_DataContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Struct_Data.
	VisitStruct_Data(ctx *Struct_DataContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Id_Struct.
	VisitId_Struct(ctx *Id_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Struct_Llamada_Func.
	VisitStruct_Llamada_Func(ctx *Struct_Llamada_FuncContext) interface{}

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

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Struct.
	VisitTipo_Struct(ctx *Tipo_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Vector.
	VisitTipo_Vector(ctx *Tipo_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Rel.
	VisitExpr_Rel(ctx *Expr_RelContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Funciones_Embebidas.
	VisitExpr_Funciones_Embebidas(ctx *Expr_Funciones_EmbebidasContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Decimal.
	VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Caracter.
	VisitExpr_Caracter(ctx *Expr_CaracterContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_SumRes.
	VisitExpr_SumRes(ctx *Expr_SumResContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Llamada_Funcion_Struct.
	VisitExpr_Llamada_Funcion_Struct(ctx *Expr_Llamada_Funcion_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Matriz.
	VisitExpr_Matriz(ctx *Expr_MatrizContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Neg.
	VisitExpr_Neg(ctx *Expr_NegContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_MulDiv.
	VisitExpr_MulDiv(ctx *Expr_MulDivContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Nil.
	VisitExpr_Nil(ctx *Expr_NilContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_isEmpty.
	VisitFunc_Vector_isEmpty(ctx *Func_Vector_isEmptyContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Llamada_Funcion.
	VisitExpr_Llamada_Funcion(ctx *Expr_Llamada_FuncionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Func_Vector_Count.
	VisitFunc_Vector_Count(ctx *Func_Vector_CountContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Struct.
	VisitExpr_Struct(ctx *Expr_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Cadena.
	VisitExpr_Cadena(ctx *Expr_CadenaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#Expr_Self.
	VisitExpr_Self(ctx *Expr_SelfContext) interface{}

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
