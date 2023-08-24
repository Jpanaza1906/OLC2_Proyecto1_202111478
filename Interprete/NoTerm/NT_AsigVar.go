package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_AsigVar struct {
	ID        string
	Expresion interprete.AbstractExpression
	linea     int
	columna   int
}

func NewNT_AsigVar(ID string, expr interprete.AbstractExpression, linea int, columna int) *NT_AsigVar {
	return &NT_AsigVar{
		ID:        ID,
		Expresion: expr,
		linea:     linea,
		columna:   columna,
	}
}

func (NT_AVar *NT_AsigVar) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := NT_AVar.Expresion.Interpretar(ctx)

	_, ok := ctx.GetVariable(NT_AVar.ID)

	if !ok {
		ctx.AddError("La variable " + NT_AVar.ID + " no existe" + " en la linea:" + strconv.Itoa(NT_AVar.linea) + " y columna:" + strconv.Itoa(NT_AVar.columna))
		return interprete.NewNil()
	}
	ctx.AsigVariable(NT_AVar.ID, expr)
	return interprete.NewNil()
}
