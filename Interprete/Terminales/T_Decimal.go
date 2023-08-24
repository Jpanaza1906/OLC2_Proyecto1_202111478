package terminales

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type T_Decimal struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Decimal----------------------------------------------

func NewT_Decimal(valor string, linea int, columna int) *T_Decimal {
	return &T_Decimal{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (T_dec *T_Decimal) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	num, err := strconv.ParseFloat(T_dec.Valor, 64)
	if err != nil {
		ctx.AddError("Error: No se puede convertir " + T_dec.Valor + " a float")
		return interprete.NewNil()
	}
	return interprete.NewFloatLiteral(num)
}
