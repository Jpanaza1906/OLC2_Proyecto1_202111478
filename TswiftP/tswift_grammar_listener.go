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

	// EnterS_Print is called when entering the S_Print production.
	EnterS_Print(c *S_PrintContext)

	// EnterS_Declaracion is called when entering the S_Declaracion production.
	EnterS_Declaracion(c *S_DeclaracionContext)

	// EnterS_Constante is called when entering the S_Constante production.
	EnterS_Constante(c *S_ConstanteContext)

	// EnterS_Asignacion is called when entering the S_Asignacion production.
	EnterS_Asignacion(c *S_AsignacionContext)

	// EnterS_If is called when entering the S_If production.
	EnterS_If(c *S_IfContext)

	// EnterS_Switch is called when entering the S_Switch production.
	EnterS_Switch(c *S_SwitchContext)

	// EnterS_Guard is called when entering the S_Guard production.
	EnterS_Guard(c *S_GuardContext)

	// EnterS_While is called when entering the S_While production.
	EnterS_While(c *S_WhileContext)

	// EnterS_For is called when entering the S_For production.
	EnterS_For(c *S_ForContext)

	// EnterS_Transicion is called when entering the S_Transicion production.
	EnterS_Transicion(c *S_TransicionContext)

	// EnterS_Declaracion_Vector is called when entering the S_Declaracion_Vector production.
	EnterS_Declaracion_Vector(c *S_Declaracion_VectorContext)

	// EnterS_Funcion_Vector is called when entering the S_Funcion_Vector production.
	EnterS_Funcion_Vector(c *S_Funcion_VectorContext)

	// EnterS_Asignacion_Vector is called when entering the S_Asignacion_Vector production.
	EnterS_Asignacion_Vector(c *S_Asignacion_VectorContext)

	// EnterS_Declaracion_Metodo is called when entering the S_Declaracion_Metodo production.
	EnterS_Declaracion_Metodo(c *S_Declaracion_MetodoContext)

	// EnterS_Declaracion_Funcion is called when entering the S_Declaracion_Funcion production.
	EnterS_Declaracion_Funcion(c *S_Declaracion_FuncionContext)

	// EnterS_Llamada_Funcion is called when entering the S_Llamada_Funcion production.
	EnterS_Llamada_Funcion(c *S_Llamada_FuncionContext)

	// EnterS_Declaracion_Matriz is called when entering the S_Declaracion_Matriz production.
	EnterS_Declaracion_Matriz(c *S_Declaracion_MatrizContext)

	// EnterS_Asignacion_Matriz is called when entering the S_Asignacion_Matriz production.
	EnterS_Asignacion_Matriz(c *S_Asignacion_MatrizContext)

	// EnterS_Def_Struct is called when entering the S_Def_Struct production.
	EnterS_Def_Struct(c *S_Def_StructContext)

	// EnterS_Self_Data is called when entering the S_Self_Data production.
	EnterS_Self_Data(c *S_Self_DataContext)

	// EnterS_Struct_Data is called when entering the S_Struct_Data production.
	EnterS_Struct_Data(c *S_Struct_DataContext)

	// EnterS_Struct_Llamada_Func is called when entering the S_Struct_Llamada_Func production.
	EnterS_Struct_Llamada_Func(c *S_Struct_Llamada_FuncContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

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

	// EnterSumAsg is called when entering the SumAsg production.
	EnterSumAsg(c *SumAsgContext)

	// EnterResAsg is called when entering the ResAsg production.
	EnterResAsg(c *ResAsgContext)

	// EnterAsig is called when entering the Asig production.
	EnterAsig(c *AsigContext)

	// EnterIf_Simple is called when entering the If_Simple production.
	EnterIf_Simple(c *If_SimpleContext)

	// EnterIf_Else is called when entering the If_Else production.
	EnterIf_Else(c *If_ElseContext)

	// EnterIf_ElseIf is called when entering the If_ElseIf production.
	EnterIf_ElseIf(c *If_ElseIfContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterCase is called when entering the Case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the Default production.
	EnterDefault(c *DefaultContext)

	// EnterGuard is called when entering the Guard production.
	EnterGuard(c *GuardContext)

	// EnterWhile is called when entering the While production.
	EnterWhile(c *WhileContext)

	// EnterFor is called when entering the For production.
	EnterFor(c *ForContext)

	// EnterRango is called when entering the Rango production.
	EnterRango(c *RangoContext)

	// EnterBreak is called when entering the Break production.
	EnterBreak(c *BreakContext)

	// EnterContinue is called when entering the Continue production.
	EnterContinue(c *ContinueContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterDeclaracion_Vector is called when entering the Declaracion_Vector production.
	EnterDeclaracion_Vector(c *Declaracion_VectorContext)

	// EnterDef_Vector is called when entering the Def_Vector production.
	EnterDef_Vector(c *Def_VectorContext)

	// EnterDef_Vector_Vacio is called when entering the Def_Vector_Vacio production.
	EnterDef_Vector_Vacio(c *Def_Vector_VacioContext)

	// EnterDef_Vector_Id is called when entering the Def_Vector_Id production.
	EnterDef_Vector_Id(c *Def_Vector_IdContext)

	// EnterAsig_Vector is called when entering the Asig_Vector production.
	EnterAsig_Vector(c *Asig_VectorContext)

	// EnterSumAsg_Vector is called when entering the SumAsg_Vector production.
	EnterSumAsg_Vector(c *SumAsg_VectorContext)

	// EnterResAsg_Vector is called when entering the ResAsg_Vector production.
	EnterResAsg_Vector(c *ResAsg_VectorContext)

	// EnterFunc_Vector_Append is called when entering the Func_Vector_Append production.
	EnterFunc_Vector_Append(c *Func_Vector_AppendContext)

	// EnterFunc_Vector_RemoveLast is called when entering the Func_Vector_RemoveLast production.
	EnterFunc_Vector_RemoveLast(c *Func_Vector_RemoveLastContext)

	// EnterFunc_Vector_Remove is called when entering the Func_Vector_Remove production.
	EnterFunc_Vector_Remove(c *Func_Vector_RemoveContext)

	// EnterDeclaracion_Matriz is called when entering the Declaracion_Matriz production.
	EnterDeclaracion_Matriz(c *Declaracion_MatrizContext)

	// EnterTipo_Matriz is called when entering the Tipo_Matriz production.
	EnterTipo_Matriz(c *Tipo_MatrizContext)

	// EnterTipo_Matriz_Simple is called when entering the Tipo_Matriz_Simple production.
	EnterTipo_Matriz_Simple(c *Tipo_Matriz_SimpleContext)

	// EnterDef_Matriz is called when entering the Def_Matriz production.
	EnterDef_Matriz(c *Def_MatrizContext)

	// EnterDef_Matriz_Simple is called when entering the Def_Matriz_Simple production.
	EnterDef_Matriz_Simple(c *Def_Matriz_SimpleContext)

	// EnterListaCompletaVal is called when entering the ListaCompletaVal production.
	EnterListaCompletaVal(c *ListaCompletaValContext)

	// EnterListaValorSig is called when entering the ListaValorSig production.
	EnterListaValorSig(c *ListaValorSigContext)

	// EnterListaValoresHermanos is called when entering the ListaValoresHermanos production.
	EnterListaValoresHermanos(c *ListaValoresHermanosContext)

	// EnterListaExpr is called when entering the ListaExpr production.
	EnterListaExpr(c *ListaExprContext)

	// EnterSimple_Mat is called when entering the Simple_Mat production.
	EnterSimple_Mat(c *Simple_MatContext)

	// EnterSimple_Mat_Expr is called when entering the Simple_Mat_Expr production.
	EnterSimple_Mat_Expr(c *Simple_Mat_ExprContext)

	// EnterAsig_Mat is called when entering the Asig_Mat production.
	EnterAsig_Mat(c *Asig_MatContext)

	// EnterDeclaracion_Metodo is called when entering the Declaracion_Metodo production.
	EnterDeclaracion_Metodo(c *Declaracion_MetodoContext)

	// EnterDeclaracion_Funcion is called when entering the Declaracion_Funcion production.
	EnterDeclaracion_Funcion(c *Declaracion_FuncionContext)

	// EnterL_Parametros is called when entering the L_Parametros production.
	EnterL_Parametros(c *L_ParametrosContext)

	// EnterLlamada_Funcion is called when entering the Llamada_Funcion production.
	EnterLlamada_Funcion(c *Llamada_FuncionContext)

	// EnterL_Argumentos is called when entering the L_Argumentos production.
	EnterL_Argumentos(c *L_ArgumentosContext)

	// EnterFunc_Int is called when entering the Func_Int production.
	EnterFunc_Int(c *Func_IntContext)

	// EnterFunc_Float is called when entering the Func_Float production.
	EnterFunc_Float(c *Func_FloatContext)

	// EnterFunc_String is called when entering the Func_String production.
	EnterFunc_String(c *Func_StringContext)

	// EnterDef_Struct is called when entering the Def_Struct production.
	EnterDef_Struct(c *Def_StructContext)

	// EnterL_Atributos is called when entering the L_Atributos production.
	EnterL_Atributos(c *L_AtributosContext)

	// EnterL_Funciones is called when entering the L_Funciones production.
	EnterL_Funciones(c *L_FuncionesContext)

	// EnterSelf_Data is called when entering the Self_Data production.
	EnterSelf_Data(c *Self_DataContext)

	// EnterStruct_Data is called when entering the Struct_Data production.
	EnterStruct_Data(c *Struct_DataContext)

	// EnterId_Struct is called when entering the Id_Struct production.
	EnterId_Struct(c *Id_StructContext)

	// EnterStruct_Llamada_Func is called when entering the Struct_Llamada_Func production.
	EnterStruct_Llamada_Func(c *Struct_Llamada_FuncContext)

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

	// EnterTipo_Struct is called when entering the Tipo_Struct production.
	EnterTipo_Struct(c *Tipo_StructContext)

	// EnterTipo_Vector is called when entering the Tipo_Vector production.
	EnterTipo_Vector(c *Tipo_VectorContext)

	// EnterExpr_Rel is called when entering the Expr_Rel production.
	EnterExpr_Rel(c *Expr_RelContext)

	// EnterExpr_Funciones_Embebidas is called when entering the Expr_Funciones_Embebidas production.
	EnterExpr_Funciones_Embebidas(c *Expr_Funciones_EmbebidasContext)

	// EnterExpr_Decimal is called when entering the Expr_Decimal production.
	EnterExpr_Decimal(c *Expr_DecimalContext)

	// EnterExpr_Caracter is called when entering the Expr_Caracter production.
	EnterExpr_Caracter(c *Expr_CaracterContext)

	// EnterExpr_SumRes is called when entering the Expr_SumRes production.
	EnterExpr_SumRes(c *Expr_SumResContext)

	// EnterExpr_Llamada_Funcion_Struct is called when entering the Expr_Llamada_Funcion_Struct production.
	EnterExpr_Llamada_Funcion_Struct(c *Expr_Llamada_Funcion_StructContext)

	// EnterExpr_Matriz is called when entering the Expr_Matriz production.
	EnterExpr_Matriz(c *Expr_MatrizContext)

	// EnterExpr_Neg is called when entering the Expr_Neg production.
	EnterExpr_Neg(c *Expr_NegContext)

	// EnterExpr_MulDiv is called when entering the Expr_MulDiv production.
	EnterExpr_MulDiv(c *Expr_MulDivContext)

	// EnterExpr_Nil is called when entering the Expr_Nil production.
	EnterExpr_Nil(c *Expr_NilContext)

	// EnterFunc_Vector_isEmpty is called when entering the Func_Vector_isEmpty production.
	EnterFunc_Vector_isEmpty(c *Func_Vector_isEmptyContext)

	// EnterExpr_Llamada_Funcion is called when entering the Expr_Llamada_Funcion production.
	EnterExpr_Llamada_Funcion(c *Expr_Llamada_FuncionContext)

	// EnterFunc_Vector_Count is called when entering the Func_Vector_Count production.
	EnterFunc_Vector_Count(c *Func_Vector_CountContext)

	// EnterExpr_Struct is called when entering the Expr_Struct production.
	EnterExpr_Struct(c *Expr_StructContext)

	// EnterExpr_Cadena is called when entering the Expr_Cadena production.
	EnterExpr_Cadena(c *Expr_CadenaContext)

	// EnterExpr_Self is called when entering the Expr_Self production.
	EnterExpr_Self(c *Expr_SelfContext)

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

	// EnterExpr_Vector is called when entering the Expr_Vector production.
	EnterExpr_Vector(c *Expr_VectorContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitS_Print is called when exiting the S_Print production.
	ExitS_Print(c *S_PrintContext)

	// ExitS_Declaracion is called when exiting the S_Declaracion production.
	ExitS_Declaracion(c *S_DeclaracionContext)

	// ExitS_Constante is called when exiting the S_Constante production.
	ExitS_Constante(c *S_ConstanteContext)

	// ExitS_Asignacion is called when exiting the S_Asignacion production.
	ExitS_Asignacion(c *S_AsignacionContext)

	// ExitS_If is called when exiting the S_If production.
	ExitS_If(c *S_IfContext)

	// ExitS_Switch is called when exiting the S_Switch production.
	ExitS_Switch(c *S_SwitchContext)

	// ExitS_Guard is called when exiting the S_Guard production.
	ExitS_Guard(c *S_GuardContext)

	// ExitS_While is called when exiting the S_While production.
	ExitS_While(c *S_WhileContext)

	// ExitS_For is called when exiting the S_For production.
	ExitS_For(c *S_ForContext)

	// ExitS_Transicion is called when exiting the S_Transicion production.
	ExitS_Transicion(c *S_TransicionContext)

	// ExitS_Declaracion_Vector is called when exiting the S_Declaracion_Vector production.
	ExitS_Declaracion_Vector(c *S_Declaracion_VectorContext)

	// ExitS_Funcion_Vector is called when exiting the S_Funcion_Vector production.
	ExitS_Funcion_Vector(c *S_Funcion_VectorContext)

	// ExitS_Asignacion_Vector is called when exiting the S_Asignacion_Vector production.
	ExitS_Asignacion_Vector(c *S_Asignacion_VectorContext)

	// ExitS_Declaracion_Metodo is called when exiting the S_Declaracion_Metodo production.
	ExitS_Declaracion_Metodo(c *S_Declaracion_MetodoContext)

	// ExitS_Declaracion_Funcion is called when exiting the S_Declaracion_Funcion production.
	ExitS_Declaracion_Funcion(c *S_Declaracion_FuncionContext)

	// ExitS_Llamada_Funcion is called when exiting the S_Llamada_Funcion production.
	ExitS_Llamada_Funcion(c *S_Llamada_FuncionContext)

	// ExitS_Declaracion_Matriz is called when exiting the S_Declaracion_Matriz production.
	ExitS_Declaracion_Matriz(c *S_Declaracion_MatrizContext)

	// ExitS_Asignacion_Matriz is called when exiting the S_Asignacion_Matriz production.
	ExitS_Asignacion_Matriz(c *S_Asignacion_MatrizContext)

	// ExitS_Def_Struct is called when exiting the S_Def_Struct production.
	ExitS_Def_Struct(c *S_Def_StructContext)

	// ExitS_Self_Data is called when exiting the S_Self_Data production.
	ExitS_Self_Data(c *S_Self_DataContext)

	// ExitS_Struct_Data is called when exiting the S_Struct_Data production.
	ExitS_Struct_Data(c *S_Struct_DataContext)

	// ExitS_Struct_Llamada_Func is called when exiting the S_Struct_Llamada_Func production.
	ExitS_Struct_Llamada_Func(c *S_Struct_Llamada_FuncContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

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

	// ExitSumAsg is called when exiting the SumAsg production.
	ExitSumAsg(c *SumAsgContext)

	// ExitResAsg is called when exiting the ResAsg production.
	ExitResAsg(c *ResAsgContext)

	// ExitAsig is called when exiting the Asig production.
	ExitAsig(c *AsigContext)

	// ExitIf_Simple is called when exiting the If_Simple production.
	ExitIf_Simple(c *If_SimpleContext)

	// ExitIf_Else is called when exiting the If_Else production.
	ExitIf_Else(c *If_ElseContext)

	// ExitIf_ElseIf is called when exiting the If_ElseIf production.
	ExitIf_ElseIf(c *If_ElseIfContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitCase is called when exiting the Case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the Default production.
	ExitDefault(c *DefaultContext)

	// ExitGuard is called when exiting the Guard production.
	ExitGuard(c *GuardContext)

	// ExitWhile is called when exiting the While production.
	ExitWhile(c *WhileContext)

	// ExitFor is called when exiting the For production.
	ExitFor(c *ForContext)

	// ExitRango is called when exiting the Rango production.
	ExitRango(c *RangoContext)

	// ExitBreak is called when exiting the Break production.
	ExitBreak(c *BreakContext)

	// ExitContinue is called when exiting the Continue production.
	ExitContinue(c *ContinueContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitDeclaracion_Vector is called when exiting the Declaracion_Vector production.
	ExitDeclaracion_Vector(c *Declaracion_VectorContext)

	// ExitDef_Vector is called when exiting the Def_Vector production.
	ExitDef_Vector(c *Def_VectorContext)

	// ExitDef_Vector_Vacio is called when exiting the Def_Vector_Vacio production.
	ExitDef_Vector_Vacio(c *Def_Vector_VacioContext)

	// ExitDef_Vector_Id is called when exiting the Def_Vector_Id production.
	ExitDef_Vector_Id(c *Def_Vector_IdContext)

	// ExitAsig_Vector is called when exiting the Asig_Vector production.
	ExitAsig_Vector(c *Asig_VectorContext)

	// ExitSumAsg_Vector is called when exiting the SumAsg_Vector production.
	ExitSumAsg_Vector(c *SumAsg_VectorContext)

	// ExitResAsg_Vector is called when exiting the ResAsg_Vector production.
	ExitResAsg_Vector(c *ResAsg_VectorContext)

	// ExitFunc_Vector_Append is called when exiting the Func_Vector_Append production.
	ExitFunc_Vector_Append(c *Func_Vector_AppendContext)

	// ExitFunc_Vector_RemoveLast is called when exiting the Func_Vector_RemoveLast production.
	ExitFunc_Vector_RemoveLast(c *Func_Vector_RemoveLastContext)

	// ExitFunc_Vector_Remove is called when exiting the Func_Vector_Remove production.
	ExitFunc_Vector_Remove(c *Func_Vector_RemoveContext)

	// ExitDeclaracion_Matriz is called when exiting the Declaracion_Matriz production.
	ExitDeclaracion_Matriz(c *Declaracion_MatrizContext)

	// ExitTipo_Matriz is called when exiting the Tipo_Matriz production.
	ExitTipo_Matriz(c *Tipo_MatrizContext)

	// ExitTipo_Matriz_Simple is called when exiting the Tipo_Matriz_Simple production.
	ExitTipo_Matriz_Simple(c *Tipo_Matriz_SimpleContext)

	// ExitDef_Matriz is called when exiting the Def_Matriz production.
	ExitDef_Matriz(c *Def_MatrizContext)

	// ExitDef_Matriz_Simple is called when exiting the Def_Matriz_Simple production.
	ExitDef_Matriz_Simple(c *Def_Matriz_SimpleContext)

	// ExitListaCompletaVal is called when exiting the ListaCompletaVal production.
	ExitListaCompletaVal(c *ListaCompletaValContext)

	// ExitListaValorSig is called when exiting the ListaValorSig production.
	ExitListaValorSig(c *ListaValorSigContext)

	// ExitListaValoresHermanos is called when exiting the ListaValoresHermanos production.
	ExitListaValoresHermanos(c *ListaValoresHermanosContext)

	// ExitListaExpr is called when exiting the ListaExpr production.
	ExitListaExpr(c *ListaExprContext)

	// ExitSimple_Mat is called when exiting the Simple_Mat production.
	ExitSimple_Mat(c *Simple_MatContext)

	// ExitSimple_Mat_Expr is called when exiting the Simple_Mat_Expr production.
	ExitSimple_Mat_Expr(c *Simple_Mat_ExprContext)

	// ExitAsig_Mat is called when exiting the Asig_Mat production.
	ExitAsig_Mat(c *Asig_MatContext)

	// ExitDeclaracion_Metodo is called when exiting the Declaracion_Metodo production.
	ExitDeclaracion_Metodo(c *Declaracion_MetodoContext)

	// ExitDeclaracion_Funcion is called when exiting the Declaracion_Funcion production.
	ExitDeclaracion_Funcion(c *Declaracion_FuncionContext)

	// ExitL_Parametros is called when exiting the L_Parametros production.
	ExitL_Parametros(c *L_ParametrosContext)

	// ExitLlamada_Funcion is called when exiting the Llamada_Funcion production.
	ExitLlamada_Funcion(c *Llamada_FuncionContext)

	// ExitL_Argumentos is called when exiting the L_Argumentos production.
	ExitL_Argumentos(c *L_ArgumentosContext)

	// ExitFunc_Int is called when exiting the Func_Int production.
	ExitFunc_Int(c *Func_IntContext)

	// ExitFunc_Float is called when exiting the Func_Float production.
	ExitFunc_Float(c *Func_FloatContext)

	// ExitFunc_String is called when exiting the Func_String production.
	ExitFunc_String(c *Func_StringContext)

	// ExitDef_Struct is called when exiting the Def_Struct production.
	ExitDef_Struct(c *Def_StructContext)

	// ExitL_Atributos is called when exiting the L_Atributos production.
	ExitL_Atributos(c *L_AtributosContext)

	// ExitL_Funciones is called when exiting the L_Funciones production.
	ExitL_Funciones(c *L_FuncionesContext)

	// ExitSelf_Data is called when exiting the Self_Data production.
	ExitSelf_Data(c *Self_DataContext)

	// ExitStruct_Data is called when exiting the Struct_Data production.
	ExitStruct_Data(c *Struct_DataContext)

	// ExitId_Struct is called when exiting the Id_Struct production.
	ExitId_Struct(c *Id_StructContext)

	// ExitStruct_Llamada_Func is called when exiting the Struct_Llamada_Func production.
	ExitStruct_Llamada_Func(c *Struct_Llamada_FuncContext)

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

	// ExitTipo_Struct is called when exiting the Tipo_Struct production.
	ExitTipo_Struct(c *Tipo_StructContext)

	// ExitTipo_Vector is called when exiting the Tipo_Vector production.
	ExitTipo_Vector(c *Tipo_VectorContext)

	// ExitExpr_Rel is called when exiting the Expr_Rel production.
	ExitExpr_Rel(c *Expr_RelContext)

	// ExitExpr_Funciones_Embebidas is called when exiting the Expr_Funciones_Embebidas production.
	ExitExpr_Funciones_Embebidas(c *Expr_Funciones_EmbebidasContext)

	// ExitExpr_Decimal is called when exiting the Expr_Decimal production.
	ExitExpr_Decimal(c *Expr_DecimalContext)

	// ExitExpr_Caracter is called when exiting the Expr_Caracter production.
	ExitExpr_Caracter(c *Expr_CaracterContext)

	// ExitExpr_SumRes is called when exiting the Expr_SumRes production.
	ExitExpr_SumRes(c *Expr_SumResContext)

	// ExitExpr_Llamada_Funcion_Struct is called when exiting the Expr_Llamada_Funcion_Struct production.
	ExitExpr_Llamada_Funcion_Struct(c *Expr_Llamada_Funcion_StructContext)

	// ExitExpr_Matriz is called when exiting the Expr_Matriz production.
	ExitExpr_Matriz(c *Expr_MatrizContext)

	// ExitExpr_Neg is called when exiting the Expr_Neg production.
	ExitExpr_Neg(c *Expr_NegContext)

	// ExitExpr_MulDiv is called when exiting the Expr_MulDiv production.
	ExitExpr_MulDiv(c *Expr_MulDivContext)

	// ExitExpr_Nil is called when exiting the Expr_Nil production.
	ExitExpr_Nil(c *Expr_NilContext)

	// ExitFunc_Vector_isEmpty is called when exiting the Func_Vector_isEmpty production.
	ExitFunc_Vector_isEmpty(c *Func_Vector_isEmptyContext)

	// ExitExpr_Llamada_Funcion is called when exiting the Expr_Llamada_Funcion production.
	ExitExpr_Llamada_Funcion(c *Expr_Llamada_FuncionContext)

	// ExitFunc_Vector_Count is called when exiting the Func_Vector_Count production.
	ExitFunc_Vector_Count(c *Func_Vector_CountContext)

	// ExitExpr_Struct is called when exiting the Expr_Struct production.
	ExitExpr_Struct(c *Expr_StructContext)

	// ExitExpr_Cadena is called when exiting the Expr_Cadena production.
	ExitExpr_Cadena(c *Expr_CadenaContext)

	// ExitExpr_Self is called when exiting the Expr_Self production.
	ExitExpr_Self(c *Expr_SelfContext)

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

	// ExitExpr_Vector is called when exiting the Expr_Vector production.
	ExitExpr_Vector(c *Expr_VectorContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)
}
