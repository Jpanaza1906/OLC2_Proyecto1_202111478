package terminales

import interprete "OLC2_Proyecto1_202111478/Interprete"

type T_Cad struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Cad----------------------------------------------
func NewT_Cad(valor string, linea int, columna int) *T_Cad {
	return &T_Cad{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (T_Cad *T_Cad) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewStringLiteral(T_Cad.Valor[1 : len(T_Cad.Valor)-1])
}
