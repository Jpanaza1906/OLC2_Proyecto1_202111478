package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Fargumento struct {
	Id         string
	Referencia bool
	Expresion  interprete.AbstractExpression
}

// Constructor for NT_Fargumento----------------------------------------------
func NewNT_Fargumento(id string, referencia bool, expresion interprete.AbstractExpression) *NT_Fargumento {
	return &NT_Fargumento{
		Id:         id,
		Referencia: referencia,
		Expresion:  expresion,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_FA *NT_Fargumento) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewNil()
}
