package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_FuncionStruct struct {
	Mutating bool
	Funcion  interprete.AbstractExpression
	Linea    int
	Columna  int
}

// Constructor for NT_FuncionStruct----------------------------------------------

func NewNT_FuncionStruct(mutating bool, funcion interprete.AbstractExpression, linea int, columna int) *NT_FuncionStruct {
	return &NT_FuncionStruct{
		Mutating: mutating,
		Funcion:  funcion,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_FS *NT_FuncionStruct) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	NT_FS.Funcion.Interpretar(ctx)
	vfuncion := NT_FS.Funcion.(*NT_DecFuncion)
	if NT_FS.Mutating {
		ctx.MuteFuncion(vfuncion.Id)
	}
	return interprete.NewNil()
}
