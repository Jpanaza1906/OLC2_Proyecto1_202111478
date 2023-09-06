package terminales

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type T_VectorPos struct {
	Id      string
	Pos     interprete.AbstractExpression
	Linea   int
	Columna int
}

// Constructor for T_VectorPos----------------------------------------------

func NewT_VectorPos(id string, pos interprete.AbstractExpression, linea int, columna int) *T_VectorPos {
	return &T_VectorPos{
		Id:      id,
		Pos:     pos,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (T_vec *T_VectorPos) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//Con el Id se busca el vector
	vec, ok := ctx.GetVariable(T_vec.Id)
	if !ok {
		ctx.AddError("La variable " + T_vec.Id + " no existe en la linea " + strconv.Itoa(T_vec.Linea) + " y columna " + strconv.Itoa(T_vec.Columna))
		return interprete.NewNil()
	}

	//Se verifica que la variable sea un vector
	if vec.Tipo != interprete.Vector {
		ctx.AddError("La variable " + T_vec.Id + " no es un vector en la linea " + strconv.Itoa(T_vec.Linea) + " y columna " + strconv.Itoa(T_vec.Columna))
		return interprete.NewNil()
	}

	//Se obtiene el valor de la posicion
	pos := T_vec.Pos.Interpretar(ctx)

	//Se verifica que la posicion sea un entero
	if pos.Tipo != interprete.Integer {
		ctx.AddError("La posicion debe ser un entero en la linea " + strconv.Itoa(T_vec.Linea) + " y columna " + strconv.Itoa(T_vec.Columna))
		return interprete.NewNil()
	}

	//Se verifica que la posicion sea menor al tamaño del vector
	if pos.Valor >= len(vec.ValorV) {
		ctx.AddError("La posicion " + strconv.Itoa(pos.Valor) + " no existe en el vector " + T_vec.Id + " en la linea " + strconv.Itoa(T_vec.Linea) + " y columna " + strconv.Itoa(T_vec.Columna))
		return interprete.NewNil()
	}

	//Se retorna el valor del vector en la posicion
	return &vec.ValorV[pos.Valor]

}
