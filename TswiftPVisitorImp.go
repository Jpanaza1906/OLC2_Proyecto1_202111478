package main

import (
	"OLC2_Proyecto1_202111478/TswiftP"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

type TswiftPVisitorImp struct {
	antlr.ParseTreeVisitor
}

//Funcion que conecta el visitor y devuelve obj visitor
func NewTswiftPVisitorImp() TswiftP.Tswift_GrammarVisitor {
	return &TswiftPVisitorImp{
		ParseTreeVisitor: &TswiftP.BaseTswift_GrammarVisitor{},
	}
}

// VisitA0 implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitA0(ctx *TswiftP.A0Context) interface{} {
	return ctx.A().Accept(l)
}

// VisitADECIMAL implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitADECIMAL(ctx *TswiftP.ADECIMALContext) interface{} {
	fmt.Println("Imprimiendo decimal")
	fmt.Println(ctx.DECIMAL().GetText())
	return nil
}

// VisitAENTERO implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitAENTERO(ctx *TswiftP.AENTEROContext) interface{} {
	fmt.Println("Imprimiento entero")
	fmt.Println(ctx.ENTERO().GetText())
	return nil
}

// VisitAEPSILUM implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitAEPSILUM(ctx *TswiftP.AEPSILUMContext) interface{} {
	return nil
}

// VisitAID implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitAID(ctx *TswiftP.AIDContext) interface{} {
	fmt.Println("ID en la linea")
	fmt.Println(ctx.ID().GetSymbol().GetLine())
	fmt.Println("ID en la columna")
	fmt.Println(ctx.ID().GetSymbol().GetColumn())
	fmt.Println(ctx.ID().GetText())
	return nil
}

// VisitS implements TswiftP.Tswift_GrammarVisitor.
func (l *TswiftPVisitorImp) VisitS(ctx *TswiftP.SContext) interface{} {
	ctx.A().Accept(l)
	return "Se ha completado el reconocimiento"
}

//FUNCION VISIT QUE PUEDA DEPURAR ERRORES
func (l *TswiftPVisitorImp) Visit(arbol antlr.ParseTree) interface{} {
	switch val := arbol.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		nodo := arbol.Accept(l)
		return nodo
	}
}
