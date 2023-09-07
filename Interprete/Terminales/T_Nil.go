package terminales

import interprete "OLC2_Proyecto1_202111478/Interprete"

type T_Nil struct {
	Valor   string
	Linea   int
	Columna int
}

// Constructor for T_Nil----------------------------------------------
func NewT_Nil(valor string, linea int, columna int) *T_Nil {
	return &T_Nil{
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (Tn *T_Nil) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewNil()
}
