package main

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	noterm "OLC2_Proyecto1_202111478/Interprete/NoTerm"
	terminales "OLC2_Proyecto1_202111478/Interprete/Terminales"
	"OLC2_Proyecto1_202111478/TswiftP"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorInterprete struct {
	antlr.ParseTreeVisitor
	Raiz interprete.AbstractExpression
}

func NewVisitorInterprete() TswiftP.Tswift_GrammarVisitor {
	return &VisitorInterprete{
		ParseTreeVisitor: &TswiftP.BaseTswift_GrammarVisitor{},
		Raiz:             nil,
	}
}

func (vI *VisitorInterprete) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return noterm.NewNT_Error(val.GetText())
	default:
		nodoInterprete, ok := tree.Accept(vI).(interprete.AbstractExpression)
		if ok {
			return nodoInterprete
		}
		return noterm.NewNT_Error("Nodo desconocido.")
	}
}

func (vI *VisitorInterprete) VisitChildren(node antlr.RuleNode) interface{} {
	//panic("not implemented") // TODO: Implement
	return ""
}

func (vI *VisitorInterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
	//panic("not implemented") // TODO: Implement
	return ""
}

func (vI *VisitorInterprete) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return noterm.NewNT_Error(node.GetText())
}

// Visit a parse tree produced by Tswift_GrammarParser#SLSentencias.
func (vI *VisitorInterprete) VisitSLSentencias(ctx *TswiftP.SLSentenciasContext) interface{} {
	lsentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)
	nodoS := noterm.NewNT_S(lsentencias)
	vI.Raiz = nodoS
	return nodoS
}

// SENTENCIAS --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#L_Sentencia.
func (vI *VisitorInterprete) VisitL_Sentencia(ctx *TswiftP.L_SentenciaContext) interface{} {
	lsentencias := noterm.NewNT_LSentencias()
	sentenciasAntlr := ctx.AllSentencia()
	for _, sentenciaAntlr := range sentenciasAntlr {
		nodoSentencia := sentenciaAntlr.Accept(vI).(interprete.AbstractExpression)
		lsentencias.AddSentencia(nodoSentencia)
	}
	return lsentencias
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Consola.
func (vI *VisitorInterprete) VisitS_Consola(ctx *TswiftP.S_ConsolaContext) interface{} {
	consola := noterm.NewNT_Print()
	lExpresiones := ctx.AllE()

	for _, expresion := range lExpresiones {
		nodoExpresion := expresion.Accept(vI).(interprete.AbstractExpression)
		consola.AddExpresion(nodoExpresion)
	}
	return consola
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Declaracion.
func (vI *VisitorInterprete) VisitS_Declaracion(ctx *TswiftP.S_DeclaracionContext) interface{} {
	return ctx.Declaracion().Accept(vI).(interprete.AbstractExpression)

}

// Visit a parse tree produced by Tswift_GrammarParser#S_Constante.
func (vI *VisitorInterprete) VisitS_Constante(ctx *TswiftP.S_ConstanteContext) interface{} {
	return ctx.Constante().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Asignacion.
func (vI *VisitorInterprete) VisitS_Asignacion(ctx *TswiftP.S_AsignacionContext) interface{} {
	return ctx.Asignacion().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_If.
func (vI *VisitorInterprete) VisitS_If(ctx *TswiftP.S_IfContext) interface{} {
	return ctx.If_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Switch.
func (vI *VisitorInterprete) VisitS_Switch(ctx *TswiftP.S_SwitchContext) interface{} {
	return ctx.Switch_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Guard.
func (vI *VisitorInterprete) VisitS_Guard(ctx *TswiftP.S_GuardContext) interface{} {
	return ctx.Guard_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_While.
func (vI *VisitorInterprete) VisitS_While(ctx *TswiftP.S_WhileContext) interface{} {
	return ctx.While_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_For.
func (vI *VisitorInterprete) VisitS_For(ctx *TswiftP.S_ForContext) interface{} {
	return ctx.For_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#S_Transicion.
func (vI *VisitorInterprete) VisitS_Transicion(ctx *TswiftP.S_TransicionContext) interface{} {
	return ctx.Trans_sentencia().Accept(vI).(interprete.AbstractExpression)
}

// DECLARACIONES --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Tipo_Val.
func (vI *VisitorInterprete) VisitDeclaracion_Tipo_Val(ctx *TswiftP.Declaracion_Tipo_ValContext) interface{} {
	// tipo
	tipo := ctx.Tipo().GetText()

	// id
	id := ctx.ID().GetText()

	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_DecVar(id, tipo, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Val.
func (vI *VisitorInterprete) VisitDeclaracion_Val(ctx *TswiftP.Declaracion_ValContext) interface{} {
	// id
	id := ctx.ID().GetText()

	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_DecVar(id, "", expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Declaracion_Tipo_noVal.
func (vI *VisitorInterprete) VisitDeclaracion_Tipo_noVal(ctx *TswiftP.Declaracion_Tipo_noValContext) interface{} {
	// tipo
	tipo := ctx.Tipo().GetText()

	// id
	id := ctx.ID().GetText()

	return noterm.NewNT_DecVar(id, tipo, nil, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// CONTANTES --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Constante_Tipo_Val.
func (vI *VisitorInterprete) VisitConstante_Tipo_Val(ctx *TswiftP.Constante_Tipo_ValContext) interface{} {
	// tipo
	tipo := ctx.Tipo().GetText()

	// id
	id := ctx.ID().GetText()

	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_DecConst(id, tipo, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Constante_Val.
func (vI *VisitorInterprete) VisitConstante_Val(ctx *TswiftP.Constante_ValContext) interface{} {
	// id
	id := ctx.ID().GetText()

	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_DecConst(id, "", expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// ASIGNACIONES --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#SumAsg.
func (vI *VisitorInterprete) VisitSumAsg(ctx *TswiftP.SumAsgContext) interface{} {
	//Se obtiene el ID
	id := ctx.ID().GetText()
	//la expresion seria +=, se obtiene la expresion con su ID sumado
	expr := noterm.NewNT_Suma(noterm.NewNT_Identificador(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()), ctx.E().Accept(vI).(interprete.AbstractExpression))

	return noterm.NewNT_AsigVar(id, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarParser#ResAsg.
func (vI *VisitorInterprete) VisitResAsg(ctx *TswiftP.ResAsgContext) interface{} {
	//Se obtiene el ID
	id := ctx.ID().GetText()
	//la expresion seria -=, se obtiene la expresion con su ID restado
	expr := noterm.NewNT_Resta(noterm.NewNT_Identificador(id, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn()), ctx.E().Accept(vI).(interprete.AbstractExpression))

	return noterm.NewNT_AsigVar(id, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Asig.
func (vI *VisitorInterprete) VisitAsig(ctx *TswiftP.AsigContext) interface{} {
	//Se obtiene el ID
	id := ctx.ID().GetText()
	//Se obtiene la expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_AsigVar(id, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
}

// IF SENTENCIAS --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#If_Simple.
func (vI *VisitorInterprete) VisitIf_Simple(ctx *TswiftP.If_SimpleContext) interface{} {

	// condicion
	condicion := ctx.E().Accept(vI).(interprete.AbstractExpression)

	//si
	si := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_IfSentencia(condicion, si, nil)

}

// Visit a parse tree produced by Tswift_GrammarParser#If_Else.
func (vI *VisitorInterprete) VisitIf_Else(ctx *TswiftP.If_ElseContext) interface{} {

	// condicion
	condicion := ctx.E().Accept(vI).(interprete.AbstractExpression)

	//si
	si := ctx.L_sentencias(0).Accept(vI).(interprete.AbstractExpression)

	//sino
	sino := ctx.L_sentencias(1).Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_IfSentencia(condicion, si, sino)
}

// Visit a parse tree produced by Tswift_GrammarParser#If_ElseIf.
func (vI *VisitorInterprete) VisitIf_ElseIf(ctx *TswiftP.If_ElseIfContext) interface{} {

	// condicion
	condicion := ctx.E().Accept(vI).(interprete.AbstractExpression)

	//si
	si := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)

	//sino
	sino := ctx.If_sentencia().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_IfSentencia(condicion, si, sino)

}

// SWITCH CASE SENTENCIA

// Visit a parse tree produced by Tswift_GrammarParser#Switch.
func (vI *VisitorInterprete) VisitSwitch(ctx *TswiftP.SwitchContext) interface{} {
	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	// casos
	casos := []interprete.AbstractExpression{}
	casosAntlr := ctx.AllL_casos()
	for _, casoAntlr := range casosAntlr {
		nodoCaso := casoAntlr.Accept(vI).(interprete.AbstractExpression)
		casos = append(casos, nodoCaso)
	}
	// default
	//si no viene default guardar nil
	var def interprete.AbstractExpression
	if ctx.L_default() != nil {
		def = ctx.L_default().Accept(vI).(interprete.AbstractExpression)
	} else {
		def = nil
	}

	return noterm.NewNT_Switch(expr, casos, def)

}

// Visit a parse tree produced by Tswift_GrammarParser#Case.
func (vI *VisitorInterprete) VisitCase(ctx *TswiftP.CaseContext) interface{} {
	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	// sentencias
	sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_Caso(expr, sentencias)
}

// Visit a parse tree produced by Tswift_GrammarParser#Default.
func (vI *VisitorInterprete) VisitDefault(ctx *TswiftP.DefaultContext) interface{} {
	// sentencias
	sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)

	return sentencias
}

//GUARD SENTENCIA --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Guard.
func (vI *VisitorInterprete) VisitGuard(ctx *TswiftP.GuardContext) interface{} {
	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	// sentencias
	sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)
	//verifica si viene una transentencia
	trans := ctx.Trans_sentencia().GetText()
	if trans == "" {
		return noterm.NewNT_Error("Error: Se esperaba una sentencia de transicion en el guard")
	}

	//fin
	transSent := ctx.Trans_sentencia().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_Guard(expr, sentencias, transSent)
}

//CICLO WHILE --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#While.
func (vI *VisitorInterprete) VisitWhile(ctx *TswiftP.WhileContext) interface{} {
	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	// sentencias
	sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_While(expr, sentencias)
}

//CICLO FOR--------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#For.
func (vI *VisitorInterprete) VisitFor(ctx *TswiftP.ForContext) interface{} {
	//id
	id := ctx.ID().GetText()

	if ctx.E() != nil {
		//expresion
		expr := ctx.E().Accept(vI).(interprete.AbstractExpression)
		//sentencias
		sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)
		return noterm.NewNT_For(id, expr, nil, sentencias, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	} else {
		//rango
		rango := ctx.Rango_p().Accept(vI).([]interprete.AbstractExpression)
		//sentencias
		sentencias := ctx.L_sentencias().Accept(vI).(interprete.AbstractExpression)
		return noterm.NewNT_For(id, nil, rango, sentencias, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	}

}

// Visit a parse tree produced by Tswift_GrammarParser#Rango.
func (vI *VisitorInterprete) VisitRango(ctx *TswiftP.RangoContext) interface{} {
	//expresion
	expr1 := ctx.E(0).Accept(vI).(interprete.AbstractExpression)
	expr2 := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	return []interprete.AbstractExpression{expr1, expr2}
}

//SENTENCIAS DE TRANSICION --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Break.
func (vI *VisitorInterprete) VisitBreak(ctx *TswiftP.BreakContext) interface{} {
	return noterm.NewNT_TransSentencia("break")
}

// Visit a parse tree produced by Tswift_GrammarParser#Continue.
func (vI *VisitorInterprete) VisitContinue(ctx *TswiftP.ContinueContext) interface{} {
	return noterm.NewNT_TransSentencia("continue")
}

// Visit a parse tree produced by Tswift_GrammarParser#Return.
func (vI *VisitorInterprete) VisitReturn(ctx *TswiftP.ReturnContext) interface{} {
	return noterm.NewNT_TransSentencia("return")
}

// TIPOS DE DATOS --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Int.
func (vI *VisitorInterprete) VisitTipo_Int(ctx *TswiftP.Tipo_IntContext) interface{} {
	return ctx.INT().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Double.
func (vI *VisitorInterprete) VisitTipo_Float(ctx *TswiftP.Tipo_FloatContext) interface{} {
	return ctx.FLOAT().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Tipo_String.
func (vI *VisitorInterprete) VisitTipo_String(ctx *TswiftP.Tipo_StringContext) interface{} {
	return ctx.STRING().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Bool.
func (vI *VisitorInterprete) VisitTipo_Bool(ctx *TswiftP.Tipo_BoolContext) interface{} {
	return ctx.BOOL().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Tipo_Character.
func (vI *VisitorInterprete) VisitTipo_Character(ctx *TswiftP.Tipo_CharacterContext) interface{} {
	return ctx.CHARACTER().GetText()
}

// EXPRESIONES --------------------------------------------------------------------------------------------

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Rel.
func (vI *VisitorInterprete) VisitExpr_Rel(ctx *TswiftP.Expr_RelContext) interface{} {
	// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstractExpression)

	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	if ctx.GetOp().GetText() == ">" {
		return noterm.NewNT_MayorQue(exprIzq, exprDer)
	} else if ctx.GetOp().GetText() == "<" {
		return noterm.NewNT_MenorQue(exprIzq, exprDer)
	} else if ctx.GetOp().GetText() == ">=" {
		return noterm.NewNT_MayorIgual(exprIzq, exprDer)
	} else if ctx.GetOp().GetText() == "<=" {
		return noterm.NewNT_MenorIgual(exprIzq, exprDer)
	} else if ctx.GetOp().GetText() == "==" {
		return noterm.NewNT_IgualIgual(exprIzq, exprDer)
	} else {
		return noterm.NewNT_Diferente(exprIzq, exprDer)
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Logica.
func (vI *VisitorInterprete) VisitExpr_Logica(ctx *TswiftP.Expr_LogicaContext) interface{} {
	// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstractExpression)

	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	if ctx.GetOp().GetText() == "&&" {
		return noterm.NewNT_And(exprIzq, exprDer)
	} else {
		return noterm.NewNT_Or(exprIzq, exprDer)
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Neg.
func (vI *VisitorInterprete) VisitExpr_Neg(ctx *TswiftP.Expr_NegContext) interface{} {
	// expresion
	expr := ctx.E().Accept(vI).(interprete.AbstractExpression)

	if ctx.GetOp().GetText() == "!" {
		return noterm.NewNT_Not(expr)
	} else {
		return noterm.NewNT_MenosUnario(expr)
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_MulDiv.
func (vI *VisitorInterprete) VisitExpr_MulDiv(ctx *TswiftP.Expr_MulDivContext) interface{} {
	// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstractExpression)

	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	if ctx.GetOp().GetText() == "*" {
		return noterm.NewNT_Multiplicacion(exprIzq, exprDer)
	} else {
		return noterm.NewNT_Division(exprIzq, exprDer)
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_SumRes.
func (vI *VisitorInterprete) VisitExpr_SumRes(ctx *TswiftP.Expr_SumResContext) interface{} {
	// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstractExpression)

	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	if ctx.GetOp().GetText() == "+" {
		return noterm.NewNT_Suma(exprIzq, exprDer)
	} else {
		return noterm.NewNT_Resta(exprIzq, exprDer)
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Mod.
func (vI *VisitorInterprete) VisitExpr_Mod(ctx *TswiftP.Expr_ModContext) interface{} {
	// izquierda
	exprIzq := ctx.E(0).Accept(vI).(interprete.AbstractExpression)

	// derecha
	exprDer := ctx.E(1).Accept(vI).(interprete.AbstractExpression)

	return noterm.NewNT_Modulo(exprIzq, exprDer)
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Par.
func (vI *VisitorInterprete) VisitExpr_Par(ctx *TswiftP.Expr_ParContext) interface{} {
	return ctx.E().Accept(vI).(interprete.AbstractExpression)
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Nil.
func (vI *VisitorInterprete) VisitExpr_Nil(ctx *TswiftP.Expr_NilContext) interface{} {
	return ctx.NIL().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Booleano.
func (vI *VisitorInterprete) VisitExpr_Booleano(ctx *TswiftP.Expr_BooleanoContext) interface{} {
	return terminales.NewT_Bool(ctx.GetOp().GetText(),
		ctx.GetOp().GetLine(),
		ctx.GetOp().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Id.
func (vI *VisitorInterprete) VisitExpr_Id(ctx *TswiftP.Expr_IdContext) interface{} {
	return noterm.NewNT_Identificador(ctx.ID().GetText(),
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Decimal.
func (vI *VisitorInterprete) VisitExpr_Decimal(ctx *TswiftP.Expr_DecimalContext) interface{} {
	return terminales.NewT_Decimal(ctx.DECIMAL().GetText(),
		ctx.DECIMAL().GetSymbol().GetLine(),
		ctx.DECIMAL().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Entero.
func (vI *VisitorInterprete) VisitExpr_Entero(ctx *TswiftP.Expr_EnteroContext) interface{} {
	return terminales.NewT_Entero(ctx.ENTERO().GetText(),
		ctx.ENTERO().GetSymbol().GetLine(),
		ctx.ENTERO().GetSymbol().GetColumn(),
	)
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Cadena.
func (vI *VisitorInterprete) VisitExpr_Cadena(ctx *TswiftP.Expr_CadenaContext) interface{} {
	return terminales.NewT_Cad(ctx.CADENA().GetText(),
		ctx.CADENA().GetSymbol().GetLine(),
		ctx.CADENA().GetSymbol().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Caracter.
func (vI *VisitorInterprete) VisitExpr_Caracter(ctx *TswiftP.Expr_CaracterContext) interface{} {
	return terminales.NewT_Char(ctx.CARACTER().GetText(),
		ctx.CARACTER().GetSymbol().GetLine(),
		ctx.CARACTER().GetSymbol().GetColumn())
}
