package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_DefStruct struct {
	Id               string
	StructSentencias []interprete.AbstractExpression
	Linea            int
	Columna          int
}

// Constructor for NT_DefStruct----------------------------------------------
func NewNT_DefStruct(id string, structsentencias []interprete.AbstractExpression, linea int, columna int) *NT_DefStruct {
	return &NT_DefStruct{
		Id:               id,
		StructSentencias: structsentencias,
		Linea:            linea,
		Columna:          columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NT_DS *NT_DefStruct) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	// Se crea el contexto del struct
	structContext := interprete.NewContexto(NT_DS.Id)
	// Se agregan las sentencias al struct
	for _, sentencia := range NT_DS.StructSentencias {
		sentencia.Interpretar(structContext)
	}
	if ctx.ExistStruct(NT_DS.Id) {
		ctx.AddError("Ya existe un struct con el nombre " + NT_DS.Id + " en la Linea:" + strconv.Itoa(NT_DS.Linea) + " y Columna:" + strconv.Itoa(NT_DS.Columna))
		return interprete.NewNil()
	}
	// Se agrega el struct al contexto actual
	ctx.Structs[NT_DS.Id] = structContext
	return interprete.NewNil()
}
