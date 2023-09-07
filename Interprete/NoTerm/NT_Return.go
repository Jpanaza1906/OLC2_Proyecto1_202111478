package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Return struct {
	Exp interprete.AbstractExpression
}

// Constructor for NT_Return----------------------------------------------

func NewNT_Return(exp interprete.AbstractExpression) *NT_Return {
	return &NT_Return{
		Exp: exp,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTR *NT_Return) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	if NTR.Exp != nil {
		return NTR.Exp.Interpretar(ctx)
	}
	return interprete.NewNil()
}
