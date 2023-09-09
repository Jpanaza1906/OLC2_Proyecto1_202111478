package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_SelfData struct {
	Asignacion interprete.AbstractExpression
	Linea      int
	Columna    int
}

// Constructor for NT_SelfData----------------------------------------------

func NewNT_SelfData(asignacion interprete.AbstractExpression, linea int, columna int) *NT_SelfData {
	return &NT_SelfData{
		Asignacion: asignacion,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_SD *NT_SelfData) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	NT_SD.Asignacion.Interpretar(ctx)
	return interprete.NewNil()
}
