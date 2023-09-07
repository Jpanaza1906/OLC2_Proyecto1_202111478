package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_AsigIgualVector struct {
	ID        string
	Tipo      string
	Posicion  interprete.AbstractExpression
	Expresion interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_AsigIgualVector----------------------------------------------
func NewNT_AsigIgualVector(id string, tipo string, posicion interprete.AbstractExpression, expresion interprete.AbstractExpression, linea int, columna int) *NT_AsigIgualVector {
	return &NT_AsigIgualVector{
		ID:        id,
		Tipo:      tipo,
		Posicion:  posicion,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_AIV *NT_AsigIgualVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	vec1, ok := ctx.GetVariable(NT_AIV.ID)

	if !ok {
		ctx.AddError("La variable " + NT_AIV.ID + " no existe en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
		return interprete.NewNil()
	}

	vec := vec1.Resultado

	//Se verifica que la variable sea un vector
	if vec.Tipo != interprete.Vector {
		ctx.AddError("La variable " + NT_AIV.ID + " no es un vector en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
		return interprete.NewNil()
	}

	//Se obtiene el valor de la posicion
	pos := NT_AIV.Posicion.Interpretar(ctx)

	//Se verifica que la posicion sea un entero
	if pos.Tipo != interprete.Integer {
		ctx.AddError("La posicion debe ser un entero en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
		return interprete.NewNil()
	}

	//Se obtiene el valor de la expresion
	valor := NT_AIV.Expresion.Interpretar(ctx)

	//Se verifica que el tipo de la expresion sea igual al tipo del vector
	for _, val := range vec.ValorV {
		if val.Tipo.String() != valor.Tipo.String() {
			ctx.AddError("El tipo de dato del vector no coincide con el tipo de dato que se le asigno en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
			return interprete.NewNil()
		}
	}

	//Se verifica que la posicion sea menor al tamaño del vector
	if pos.Valor >= len(vec.ValorV) {
		ctx.AddError("La posicion " + strconv.Itoa(pos.Valor) + " no existe en el vector " + NT_AIV.ID + " en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
		return interprete.NewNil()
	}

	// Segun el tipo de operacion se suma o se resta
	if NT_AIV.Tipo == "+=" {
		valor := vec.ValorV[pos.Valor].Valor + valor.Valor
		vec.ValorV[pos.Valor] = *interprete.NewIntLiteral(valor)
	} else if NT_AIV.Tipo == "-=" {
		valor := vec.ValorV[pos.Valor].Valor - valor.Valor
		vec.ValorV[pos.Valor] = *interprete.NewIntLiteral(valor)
	}

	// se actualiza el valor de la variable
	if ctx.AsigVariable(NT_AIV.ID, interprete.NewVectorLiteral(vec.ValorV)) {
		return interprete.NewNil()
	} else {
		ctx.AddError("Error al actualizar el valor de la variable " + NT_AIV.ID + " en la linea " + strconv.Itoa(NT_AIV.Linea) + " y columna " + strconv.Itoa(NT_AIV.Columna))
		return interprete.NewNil()
	}

}
