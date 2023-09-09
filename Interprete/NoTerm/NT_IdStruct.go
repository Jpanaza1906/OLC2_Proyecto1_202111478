package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_IdStruct struct {
	IdVar   interprete.AbstractExpression
	IdAt    string
	Linea   int
	Columna int
}

// Constructor for NT_IdStruct----------------------------------------------
func NewNT_IdStruct(idVar interprete.AbstractExpression, idAt string, linea int, columna int) *NT_IdStruct {
	return &NT_IdStruct{
		IdVar:   idVar,
		IdAt:    idAt,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NT_IS *NT_IdStruct) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewNil()
}
