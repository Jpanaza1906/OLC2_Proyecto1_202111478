package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_FuncVector struct {
	ID        string
	Expresion interprete.AbstractExpression
	Tipo      string
	Linea     int
	Columna   int
}

// Constructor for NT_FuncVector----------------------------------------------

func NewNT_FuncVector(ID string, expresion interprete.AbstractExpression, tipo string, linea int, columna int) *NT_FuncVector {
	return &NT_FuncVector{
		ID:        ID,
		Expresion: expresion,
		Tipo:      tipo,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_FV *NT_FuncVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//se verifica si el id existe
	expr, ok := ctx.GetVariable(NT_FV.ID)
	if !ok {
		ctx.AddError("La variable " + NT_FV.ID + " no existe en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
		return interprete.NewNil()
	}

	//se verifica que la variable sea un vector
	if expr.Tipo != interprete.Vector {
		ctx.AddError("La variable " + NT_FV.ID + " no es un vector en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
		return interprete.NewNil()
	}

	// si el tipo de funcion es append
	if NT_FV.Tipo == "append" {
		// se obtiene el valor de la expresion
		valor := NT_FV.Expresion.Interpretar(ctx)
		//Se verifica que el tipo del valor coincida con cada uno de los valores del vector
		for _, val := range expr.ValorV {
			if val.Tipo.String() != valor.Tipo.String() {
				ctx.AddError("El tipo de dato del vector no coincide con el tipo de dato que se le asigno en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
				return interprete.NewNil()
			}
		}
		// se agrega el valor al vector
		expr.ValorV = append(expr.ValorV, *valor)
		// se actualiza el valor de la variable
		if ctx.AsigVariable(NT_FV.ID, expr) {
			return interprete.NewNil()
		} else {
			ctx.AddError("Error al actualizar el valor de la variable " + NT_FV.ID + " en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}
	} else if NT_FV.Tipo == "removeLast" {
		//Se verifica que el vector no este vacio
		if len(expr.ValorV) == 0 {
			ctx.AddError("El vector " + NT_FV.ID + " esta vacio en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}
		// se obtiene el ultimo valor del vector
		expr.ValorV = expr.ValorV[:len(expr.ValorV)-1]
		// se actualiza el valor de la variable
		if ctx.AsigVariable(NT_FV.ID, expr) {
			return interprete.NewNil()
		} else {
			ctx.AddError("Error al actualizar el valor de la variable " + NT_FV.ID + " en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}
	} else if NT_FV.Tipo == "remove" {
		// se obtiene el valor de la expresion
		valor := NT_FV.Expresion.Interpretar(ctx)

		//se verifica que sea un entero
		if valor.Tipo != interprete.Integer {
			ctx.AddError("El valor de la expresion debe ser de tipo int en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}

		//Se verifica que el valor este dentro del rango del vector
		if valor.Valor >= len(expr.ValorV) {
			ctx.AddError("El valor " + strconv.Itoa(valor.Valor) + " no existe en el vector " + NT_FV.ID + " en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}

		//se remueve el valor del vector en la posicion indicada
		expr.ValorV = append(expr.ValorV[:valor.Valor], expr.ValorV[valor.Valor+1:]...)

		// se actualiza el valor de la variable
		if ctx.AsigVariable(NT_FV.ID, expr) {
			return interprete.NewNil()
		} else {
			ctx.AddError("Error al actualizar el valor de la variable " + NT_FV.ID + " en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
			return interprete.NewNil()
		}
	} else if NT_FV.Tipo == "isEmpty" {
		// se verifica si el vector esta vacio
		if len(expr.ValorV) == 0 {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	} else if NT_FV.Tipo == "count" {
		// se obtiene el tamaño del vector
		return interprete.NewIntLiteral(len(expr.ValorV))
	}
	//Se agrega un error de que no existe la funcion
	ctx.AddError("La funcion " + NT_FV.Tipo + " no existe en la linea " + strconv.Itoa(NT_FV.Linea) + " y columna " + strconv.Itoa(NT_FV.Columna))
	return interprete.NewNil()
}
