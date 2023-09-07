package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

// IDENTIFICADOR ------------------------------------------------------------------

type NT_Identificador struct {
	ID      string
	linea   int
	columna int
}

// Constructor for NT_Identificador----------------------------------------------

func NewNT_Identificador(ID string, linea int, columna int) *NT_Identificador {
	return &NT_Identificador{
		ID:      ID,
		linea:   linea,
		columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_Id *NT_Identificador) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr1, ok := ctx.GetVariable(NT_Id.ID)
	if !ok {
		ctx.AddError("La variable " + NT_Id.ID + " no existe" + " en la linea:" + strconv.Itoa(NT_Id.linea) + " y columna:" + strconv.Itoa(NT_Id.columna))
		return interprete.NewNil()
	}
	expr := expr1.Resultado
	return expr
}
