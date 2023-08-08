// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import "github.com/antlr4-go/antlr/v4"

// Tswift_GrammarListener is a complete listener for a parse tree produced by Tswift_GrammarParser.
type Tswift_GrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterA0 is called when entering the A0 production.
	EnterA0(c *A0Context)

	// EnterAENTERO is called when entering the AENTERO production.
	EnterAENTERO(c *AENTEROContext)

	// EnterADECIMAL is called when entering the ADECIMAL production.
	EnterADECIMAL(c *ADECIMALContext)

	// EnterAID is called when entering the AID production.
	EnterAID(c *AIDContext)

	// EnterAEPSILUM is called when entering the AEPSILUM production.
	EnterAEPSILUM(c *AEPSILUMContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitA0 is called when exiting the A0 production.
	ExitA0(c *A0Context)

	// ExitAENTERO is called when exiting the AENTERO production.
	ExitAENTERO(c *AENTEROContext)

	// ExitADECIMAL is called when exiting the ADECIMAL production.
	ExitADECIMAL(c *ADECIMALContext)

	// ExitAID is called when exiting the AID production.
	ExitAID(c *AIDContext)

	// ExitAEPSILUM is called when exiting the AEPSILUM production.
	ExitAEPSILUM(c *AEPSILUMContext)
}
