package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_DecVar struct {
	ID        string
	Tipo      string
	Expresion interprete.AbstractExpression
	linea     int
	columna   int
}

// Constructor for NT_DecVar----------------------------------------------
func NewNT_DecVar(ID string, tipo string, expresion interprete.AbstractExpression, linea int, columna int) *NT_DecVar {
	return &NT_DecVar{
		ID:        ID,
		Tipo:      tipo,
		Expresion: expresion,
		linea:     linea,
		columna:   columna,
	}
}

func (NT_DV *NT_DecVar) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if NT_DV.Expresion != nil {
		expr = NT_DV.Expresion.Interpretar(ctx)
	}
	if NT_DV.Tipo == "" || NT_DV.Tipo == expr.Tipo.String() {
		ctx.AddVariable(NT_DV.ID, expr, "var", NT_DV.linea, NT_DV.columna)
	} else {
		ctx.AddError("Error: No se puede asignar el valor de tipo " + expr.Tipo.String() + " a la variable " + NT_DV.ID + " de tipo " + NT_DV.Tipo + " en la linea:" + strconv.Itoa(NT_DV.linea) + " y columna:" + strconv.Itoa(NT_DV.columna))
	}
	return interprete.NewNil()
}
