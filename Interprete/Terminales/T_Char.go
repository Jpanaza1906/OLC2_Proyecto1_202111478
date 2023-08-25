package terminales

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type T_Char struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Char----------------------------------------------
func NewT_Char(valor string, linea int, columna int) *T_Char {
	return &T_Char{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (Tch *T_Char) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	valorChar := Tch.Valor[1 : len(Tch.Valor)-1]
	if len(valorChar) == 1 {
		return interprete.NewCharLiteral(valorChar)
	}
	ctx.AddError("Error: El caracter '" + Tch.Valor + "' no es valido en la linea:" + strconv.Itoa(Tch.Linea) + " y columna:" + strconv.Itoa(Tch.Columna))
	return interprete.NewNil()
}
