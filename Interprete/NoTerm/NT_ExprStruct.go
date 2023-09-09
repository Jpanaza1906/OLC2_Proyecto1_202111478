package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_ExprStruct struct {
	IdVar   interprete.AbstractExpression
	IdAt    string
	Linea   int
	Columna int
}

// Constructor for NT_ExprStruct----------------------------------------------

func NewNT_ExprStruct(idVar interprete.AbstractExpression, idAt string, linea int, columna int) *NT_ExprStruct {
	return &NT_ExprStruct{
		IdVar:   idVar,
		IdAt:    idAt,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_ES *NT_ExprStruct) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	exp := NT_ES.IdVar.Interpretar(ctx)
	if exp.Nil {
		ctx.AddError("La variable " + exp.ValorS + " no existe" + " en la linea:" + strconv.Itoa(NT_ES.Linea) + " y columna:" + strconv.Itoa(NT_ES.Columna))
		return interprete.NewNil()
	}
	if exp.Tipo != interprete.StructT {
		ctx.AddError("La variable " + exp.ValorS + " no es de tipo struct" + " en la linea:" + strconv.Itoa(NT_ES.Linea) + " y columna:" + strconv.Itoa(NT_ES.Columna))
		return interprete.NewNil()
	}
	structexp := exp.StructC

	exp2, ok := structexp.GetVariable(NT_ES.IdAt)

	if !ok {
		ctx.AddError("La variable " + NT_ES.IdAt + " no existe" + " en la linea:" + strconv.Itoa(NT_ES.Linea) + " y columna:" + strconv.Itoa(NT_ES.Columna))
		return interprete.NewNil()
	}
	expfinal := exp2.Resultado
	return expfinal

}
