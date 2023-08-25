package terminales

import interprete "OLC2_Proyecto1_202111478/Interprete"

type T_Bool struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Bool----------------------------------------------
func NewT_Bool(valor string, linea int, columna int) *T_Bool {
	return &T_Bool{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (Tb *T_Bool) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	if Tb.Valor == "true" {
		return interprete.NewBoolLiteral(true)
	} else {
		return interprete.NewBoolLiteral(false)
	}
}
