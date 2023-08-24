package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Error struct {
	errorS string
}

// Constructor for NT_Error----------------------------------------------

func NewNT_Error(errorS string) *NT_Error {
	return &NT_Error{
		errorS: errorS,
	}
}

func (NT_E *NT_Error) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ctx.AddError(NT_E.errorS)
	return interprete.NewNil()
}
