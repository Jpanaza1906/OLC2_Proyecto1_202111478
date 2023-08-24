package terminales

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type T_Entero struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Entero----------------------------------------------
func NewT_Entero(valor string, linea int, columna int) *T_Entero {
	return &T_Entero{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (T_ent *T_Entero) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	num, err := strconv.Atoi(T_ent.Valor)
	if err != nil {
		ctx.AddError("Error: No se puede convertir " + T_ent.Valor + " a int")
		return interprete.NewNil()
	}
	return interprete.NewIntLiteral(num)
}
