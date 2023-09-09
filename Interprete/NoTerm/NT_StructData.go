package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_StructData struct {
	IdVar     interprete.AbstractExpression
	Operador  string
	Expresion interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_StructData----------------------------------------------

func NewNT_StructData(idVar interprete.AbstractExpression, operador string, expr interprete.AbstractExpression, linea int, columna int) *NT_StructData {
	return &NT_StructData{
		IdVar:     idVar,
		Operador:  operador,
		Expresion: expr,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_SD *NT_StructData) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := NT_SD.Expresion.Interpretar(ctx)
	stv := NT_SD.IdVar.(*NT_IdStruct)

	structv := stv.IdVar.Interpretar(ctx)

	if structv.Nil {
		ctx.AddError("La variable " + structv.ValorS + " no existe" + " en la linea:" + strconv.Itoa(NT_SD.Linea) + " y columna:" + strconv.Itoa(NT_SD.Columna))
		return interprete.NewNil()
	}
	if structv.Tipo != interprete.StructT {
		ctx.AddError("La variable " + structv.ValorS + " no es de tipo struct" + " en la linea:" + strconv.Itoa(NT_SD.Linea) + " y columna:" + strconv.Itoa(NT_SD.Columna))
		return interprete.NewNil()
	}

	if NT_SD.Operador == "=" {
		structv.StructC.AsigVariable(stv.IdAt, expr)
		return interprete.NewNil()
	}
	return interprete.NewNil()
}
