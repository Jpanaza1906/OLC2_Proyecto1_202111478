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
	Retorno     interprete.AbstractExpression
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
		Retorno:     retorno,
		Linea:       linea,
		Columna:     columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_DF *NT_DecFuncion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	// si la funcion tiene tipo de retorno se verifica
	if NT_DF.TipoRetorno != "" && NT_DF.Retorno != nil {
		// se obtiene el tipo de dato del retorno
		resul := NT_DF.Retorno.Interpretar(ctx)
		// se verifica que el tipo de dato del retorno sea igual al tipo de retorno de la funcion
		if len(ctx.ReturnState) > 0 {
			ctx.RemReturnSentencia()
			if resul.Tipo != interprete.Nil {
				if resul.Tipo.String() != NT_DF.TipoRetorno {
					ctx.AddError("El tipo de dato del retorno no coincide con el tipo de dato de la funcion en la linea " + strconv.Itoa(NT_DF.Linea) + " y columna " + strconv.Itoa(NT_DF.Columna))
					return interprete.NewNil()
				}
			}
		}
	}

	if ctx.AddFuncion(NT_DF.Id, NT_DF.Parametros, NT_DF.TipoFuncion, NT_DF.TipoRetorno, NT_DF.Sentencias, NT_DF.Retorno, NT_DF.Linea, NT_DF.Columna) {
		return interprete.NewNil()
	} else {
		ctx.AddError("La funcion " + NT_DF.Id + " ya existe en la linea " + strconv.Itoa(NT_DF.Linea) + " y columna " + strconv.Itoa(NT_DF.Columna))
		return interprete.NewNil()
	}
}
