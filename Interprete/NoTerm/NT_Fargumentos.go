package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_Fargumento struct {
	Id         string
	Referencia bool
	Expresion  interprete.AbstractExpression
	Linea      int
	Columna    int
}

// Constructor for NT_Fargumento----------------------------------------------
func NewNT_Fargumento(id string, referencia bool, expresion interprete.AbstractExpression, linea int, columna int) *NT_Fargumento {
	return &NT_Fargumento{
		Id:         id,
		Referencia: referencia,
		Expresion:  expresion,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_FA *NT_Fargumento) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := NT_FA.Expresion.Interpretar(ctx)

	_, ok := ctx.GetVariable(NT_FA.Id)

	if !ok {
		ctx.AddError("La variable " + NT_FA.Id + " no existe" + " en la linea:" + strconv.Itoa(NT_FA.Linea) + " y columna:" + strconv.Itoa(NT_FA.Columna))
		return interprete.NewNil()
	}
	ctx.AsigVariable(NT_FA.Id, expr)
	return interprete.NewNil()
}
