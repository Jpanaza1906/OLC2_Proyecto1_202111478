package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_DecConst struct {
	ID        string
	Tipo      string
	Expresion interprete.AbstractExpression
	linea     int
	columna   int
}

// Constructor for NT_DecConst----------------------------------------------
func NewNT_DecConst(ID string, tipo string, expresion interprete.AbstractExpression, linea int, columna int) *NT_DecConst {
	return &NT_DecConst{
		ID:        ID,
		Tipo:      tipo,
		Expresion: expresion,
		linea:     linea,
		columna:   columna,
	}
}

func (NT_DC *NT_DecConst) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()

	//Se ve que la expresion no sea nula
	if NT_DC.Expresion != nil {
		expr = NT_DC.Expresion.Interpretar(ctx)
	}

	//Si no tiene tipo, se le asigna el tipo de la expresion

	if NT_DC.Tipo == "" {
		NT_DC.Tipo = expr.Tipo.String()
	}

	//Si el tipo de la expresion es igual al tipo de la variable, se agrega la variable al contexto

	if NT_DC.Tipo == expr.Tipo.String() {
		if ctx.AddConstante(NT_DC.ID, expr.Tipo, expr, NT_DC.linea, NT_DC.columna) {
			return interprete.NewNil()
		} else {
			ctx.AddError("Error: No se puede asignar el valor de tipo " + expr.Tipo.String() + " a la constante " + NT_DC.ID + " de tipo " + NT_DC.Tipo + " en la linea:" + strconv.Itoa(NT_DC.linea) + " y columna:" + strconv.Itoa(NT_DC.columna))
			return interprete.NewNil()
		}
	} else {
		ctx.AddError("Error: No se puede asignar el valor de tipo " + expr.Tipo.String() + " a la variable " + NT_DC.ID + " de tipo " + NT_DC.Tipo + " en la linea:" + strconv.Itoa(NT_DC.linea) + " y columna:" + strconv.Itoa(NT_DC.columna))
	}
	return interprete.NewNil()

}
