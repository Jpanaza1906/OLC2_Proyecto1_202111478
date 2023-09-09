package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_StLlamadaFunc struct {
	IdVar   interprete.AbstractExpression
	Funcion interprete.AbstractExpression
	Linea   int
	Columna int
}

// Constructor for NT_StLlamadaFunc----------------------------------------------
func NewNT_StLlamadaFunc(idVar interprete.AbstractExpression, funcion interprete.AbstractExpression, linea int, columna int) *NT_StLlamadaFunc {
	return &NT_StLlamadaFunc{
		IdVar:   idVar,
		Funcion: funcion,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NT_SLF *NT_StLlamadaFunc) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	str := NT_SLF.IdVar.Interpretar(ctx)
	if str.Nil {
		ctx.AddError("La variable " + str.ValorS + " no existe" + " en la linea:" + strconv.Itoa(NT_SLF.Linea) + " y columna:" + strconv.Itoa(NT_SLF.Columna))
		return interprete.NewNil()
	}
	structc := str.StructC

	//se evalua la funcion en su contexto

	resul := NT_SLF.Funcion.Interpretar(structc)
	ctx.Consola = ctx.Consola + structc.Consola
	structc.Consola = ""

	return resul
}
