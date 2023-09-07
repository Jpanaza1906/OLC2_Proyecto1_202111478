package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
)

type NT_DecVector struct {
	Tipodec string
	ID      string
	TipoVec string
	Def_Vec interprete.AbstractExpression
	Linea   int
	Columna int
}

// Constructor for NT_DecVector----------------------------------------------
func NewNT_DecVector(tipodec string, id string, tipoVec string, def_vec interprete.AbstractExpression, linea int, columna int) *NT_DecVector {
	return &NT_DecVector{
		Tipodec: tipodec,
		ID:      id,
		TipoVec: tipoVec,
		Def_Vec: def_vec,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NTdv *NT_DecVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var defvect *interprete.Resultado
	//Se ve que la expresion no sea nula
	if NTdv.Def_Vec != nil {
		defvect = NTdv.Def_Vec.Interpretar(ctx)
	}
	//se crea una copia del vector
	copiavec := make([]interprete.Resultado, len(defvect.ValorV))
	copy(copiavec, defvect.ValorV)
	//Si no tiene tipo, se le asigna el tipo que contiene el vector
	if NTdv.TipoVec == "" {
		//se obtiene el primer valor del vector
		val := defvect.ValorV[0]
		//se obtiene el tipo del valor
		NTdv.TipoVec = val.Tipo.String()
	}
	var tipovector interprete.TipoE
	switch NTdv.TipoVec {
	case "Bool":
		tipovector = interprete.Bool
	case "Int":
		tipovector = interprete.Integer
	case "Float":
		tipovector = interprete.Float
	case "String":
		tipovector = interprete.String
	case "Character":
		tipovector = interprete.String
	}

	//Se verifica cada elemento del vector si coincide con el tipo del vector
	for _, val := range defvect.ValorV {
		if NTdv.TipoVec != val.Tipo.String() {
			ctx.AddError("Error: El tipo de dato del vector no coincide con el tipo de dato que se le asigno")
			return interprete.NewNil()
		}
	}
	//se crea el vector en el contexto
	if ctx.AddVector(NTdv.ID, NTdv.Tipodec, interprete.Vector, tipovector, interprete.NewVectorLiteral(copiavec), NTdv.Linea, NTdv.Columna) {
		return interprete.NewNil()
	} else {
		//agregar error de que ya existe la variable
		ctx.AddError("Error: Ya existe un vector con el nombre " + NTdv.ID)
		return interprete.NewNil()
	}

}
