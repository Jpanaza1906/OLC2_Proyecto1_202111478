// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Tswift_GrammarParser.
type Tswift_GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Tswift_GrammarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#A0.
	VisitA0(ctx *A0Context) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#AENTERO.
	VisitAENTERO(ctx *AENTEROContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#ADECIMAL.
	VisitADECIMAL(ctx *ADECIMALContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#AID.
	VisitAID(ctx *AIDContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarParser#AEPSILUM.
	VisitAEPSILUM(ctx *AEPSILUMContext) interface{}
}
