package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_AsigVector struct {
	ID       string
	Posicion interprete.AbstractExpression
	Valor    interprete.AbstractExpression
	Linea    int
	Columna  int
}

// Constructor for NT_AsigVector----------------------------------------------
func NewNT_AsigVector(id string, posicion interprete.AbstractExpression, valor interprete.AbstractExpression, linea int, columna int) *NT_AsigVector {
	return &NT_AsigVector{
		ID:       id,
		Posicion: posicion,
		Valor:    valor,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_AV *NT_AsigVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//Con el Id se busca el vector
	vec1, ok := ctx.GetVariable(NT_AV.ID)

	if !ok {
		ctx.AddError("La variable " + NT_AV.ID + " no existe en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
		return interprete.NewNil()
	}

	vec := vec1.Resultado

	//Se verifica que la variable sea un vector
	if vec.Tipo != interprete.Vector {
		ctx.AddError("La variable " + NT_AV.ID + " no es un vector en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
		return interprete.NewNil()
	}

	//Se obtiene el valor de la posicion
	pos := NT_AV.Posicion.Interpretar(ctx)

	//Se verifica que la posicion sea un entero
	if pos.Tipo != interprete.Integer {
		ctx.AddError("La posicion debe ser un entero en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
		return interprete.NewNil()
	}

	//Se obtiene el valor de la expresion
	valor := NT_AV.Valor.Interpretar(ctx)

	//Se verifica que el tipo de la expresion sea igual al tipo del vector
	for _, val := range vec.ValorV {
		if val.Tipo.String() != valor.Tipo.String() {
			ctx.AddError("El tipo de dato del vector no coincide con el tipo de dato que se le asigno en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
			return interprete.NewNil()
		}
	}

	//Se verifica que la posicion sea menor al tamaño del vector
	if pos.Valor > len(vec.ValorV) {
		ctx.AddError("La posicion es mayor al tamaño del vector en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
		return interprete.NewNil()
	}

	//Se asigna una copia del valor a la posicion del vector
	vec.ValorV[pos.Valor] = *valor

	//Se actualiza el valor de la variable
	if ctx.AsigVariable(NT_AV.ID, interprete.NewVectorLiteral(vec.ValorV)) {
		return interprete.NewNil()
	} else {
		ctx.AddError("Error al actualizar el valor de la variable " + NT_AV.ID + " en la linea " + strconv.Itoa(NT_AV.Linea) + " y columna " + strconv.Itoa(NT_AV.Columna))
		return interprete.NewNil()
	}

}
