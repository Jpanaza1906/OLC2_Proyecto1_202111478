package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_DecFuncion struct {
	Id          string
	Parametros  []interprete.AbstractExpression
	TipoFuncion string
	TipoRetorno string
	Sentencias  interprete.AbstractExpression
	Linea       int
	Columna     int
}

// Constructor for NT_DecFuncion----------------------------------------------
func NewNT_DecFuncion(id string, parametros []interprete.AbstractExpression, tipofuncion string, tiporetorno string, sentencias interprete.AbstractExpression, retorno interprete.AbstractExpression, linea int, columna int) *NT_DecFuncion {
	return &NT_DecFuncion{
		Id:          id,
		Parametros:  parametros,
		TipoFuncion: tipofuncion,
		TipoRetorno: tiporetorno,
		Sentencias:  sentencias,
		Linea:       linea,
		Columna:     columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_DF *NT_DecFuncion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {

	if ctx.AddFuncion(NT_DF.Id, NT_DF.Parametros, NT_DF.TipoFuncion, NT_DF.TipoRetorno, NT_DF.Sentencias, NT_DF.Linea, NT_DF.Columna) {
		return interprete.NewNil()
	} else {
		ctx.AddError("La funcion " + NT_DF.Id + " ya existe en la linea " + strconv.Itoa(NT_DF.Linea) + " y columna " + strconv.Itoa(NT_DF.Columna))
		return interprete.NewNil()
	}
}
