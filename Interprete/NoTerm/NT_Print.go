package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Print struct {
	LExpresiones []interprete.AbstractExpression
}

// Constructor for NT_Print----------------------------------------------
func NewNT_Print() *NT_Print {
	return &NT_Print{
		LExpresiones: make([]interprete.AbstractExpression, 0),
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (le *NT_Print) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, expresion := range le.LExpresiones {
		resultado := expresion.Interpretar(ctx)
		ctx.Print(resultado.ValorS + " ")
	}
	ctx.Print("\n")
	return interprete.NewNil()
}

func (le *NT_Print) AddExpresion(expresion interprete.AbstractExpression) {
	le.LExpresiones = append(le.LExpresiones, expresion)
}
