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
	panic("not implemented") // TODO: Implement
}

func (vI *VisitorInterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Decimal.
func (vI *VisitorInterprete) VisitExpr_Decimal(ctx *TswiftP.Expr_DecimalContext) interface{} {
	return terminales.NewT_Decimal(ctx.DECIMAL().GetText(),
		ctx.DECIMAL().GetSymbol().GetLine(),
		ctx.DECIMAL().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Caracter.
func (vI *VisitorInterprete) VisitExpr_Caracter(ctx *TswiftP.Expr_CaracterContext) interface{} {
	panic("not implemented") // TODO: Implement
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
		//return noterm.NewNT_Resta(exprIzq, exprDer)
		panic("not implemented")
	}
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Neg.
func (vI *VisitorInterprete) VisitExpr_Neg(ctx *TswiftP.Expr_NegContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_MulDiv.
func (vI *VisitorInterprete) VisitExpr_MulDiv(ctx *TswiftP.Expr_MulDivContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Nil.
func (vI *VisitorInterprete) VisitExpr_Nil(ctx *TswiftP.Expr_NilContext) interface{} {
	return ctx.NIL().GetText()
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Cadena.
func (vI *VisitorInterprete) VisitExpr_Cadena(ctx *TswiftP.Expr_CadenaContext) interface{} {
	return terminales.NewT_Cad(ctx.CADENA().GetText(),
		ctx.CADENA().GetSymbol().GetLine(),
		ctx.CADENA().GetSymbol().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Id.
func (vI *VisitorInterprete) VisitExpr_Id(ctx *TswiftP.Expr_IdContext) interface{} {
	return noterm.NewNT_Identificador(ctx.ID().GetText(),
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Mod.
func (vI *VisitorInterprete) VisitExpr_Mod(ctx *TswiftP.Expr_ModContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Par.
func (vI *VisitorInterprete) VisitExpr_Par(ctx *TswiftP.Expr_ParContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Logica.
func (vI *VisitorInterprete) VisitExpr_Logica(ctx *TswiftP.Expr_LogicaContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Booleano.
func (vI *VisitorInterprete) VisitExpr_Booleano(ctx *TswiftP.Expr_BooleanoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarParser#Expr_Entero.
func (vI *VisitorInterprete) VisitExpr_Entero(ctx *TswiftP.Expr_EnteroContext) interface{} {
	nodoEntero := terminales.NewT_Entero(
		ctx.ENTERO().GetText(),
		ctx.ENTERO().GetSymbol().GetLine(),
		ctx.ENTERO().GetSymbol().GetColumn(),
	)
	return nodoEntero
}
