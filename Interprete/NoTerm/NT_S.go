package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_S struct {
	Sentencias interprete.AbstractExpression
}

// Constructor for NT_S----------------------------------------------
func NewNT_S(sentencias interprete.AbstractExpression) *NT_S {
	o := new(NT_S)
	o.Sentencias = sentencias
	return o
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTs *NT_S) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return NTs.Sentencias.Interpretar(ctx)
}
