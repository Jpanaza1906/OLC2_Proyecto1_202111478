package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_DecVar struct {
	ID        string
	Tipo      string
	Expresion interprete.AbstractExpression
	linea     int
	columna   int
}

// Constructor for NT_DecVar----------------------------------------------
func NewNT_DecVar(ID string, tipo string, expresion interprete.AbstractExpression, linea int, columna int) *NT_DecVar {
	return &NT_DecVar{
		ID:        ID,
		Tipo:      tipo,
		Expresion: expresion,
		linea:     linea,
		columna:   columna,
	}
}

func (NT_DV *NT_DecVar) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()

	//Se ve que la expresion no sea nula
	if NT_DV.Expresion != nil {
		expr = NT_DV.Expresion.Interpretar(ctx)
	}

	//Si no tiene tipo, se le asigna el tipo de la expresion

	if NT_DV.Tipo == "" {
		NT_DV.Tipo = expr.Tipo.String()
	} else {
		switch NT_DV.Tipo {
		case "Int":
			expr.Tipo = interprete.Integer
		case "Float":
			expr.Tipo = interprete.Float
		case "String":
			expr.Tipo = interprete.String
		case "Bool":
			expr.Tipo = interprete.Bool
		case "Char":
			expr.Tipo = interprete.String
		case "Vector":
			expr.Tipo = interprete.Vector
		default:
			expr.Tipo = interprete.StructT
		}
	}

	//Si no tiene expresion, solo tipo y variable. Se le asigna un valor por defecto segun el tipo
	if NT_DV.Expresion == nil {
		if ctx.AddVariable(NT_DV.ID, expr.Tipo, expr, NT_DV.linea, NT_DV.columna) {
			return interprete.NewNil()
		} else {
			//agregar error de que ya existe la variable
			ctx.AddError("Error: Ya existe una variable con el nombre " + NT_DV.ID + " en la linea:" + strconv.Itoa(NT_DV.linea) + " y columna:" + strconv.Itoa(NT_DV.columna))
			return interprete.NewNil()
		}
	}

	//Si el tipo de la expresion es igual al tipo de la variable, se agrega la variable al contexto

	if NT_DV.Tipo == expr.Tipo.String() {
		if ctx.AddVariable(NT_DV.ID, expr.Tipo, expr, NT_DV.linea, NT_DV.columna) {
			return interprete.NewNil()
		} else {
			//agregar error de que ya existe la variable
			ctx.AddError("Error: Ya existe una variable con el nombre " + NT_DV.ID + " en la linea:" + strconv.Itoa(NT_DV.linea) + " y columna:" + strconv.Itoa(NT_DV.columna))
			return interprete.NewNil()
		}
	} else if NT_DV.Tipo == interprete.Float.String() && expr.Tipo.String() == interprete.Integer.String() {
		expr = interprete.NewFloatLiteral(float64(expr.Valor))
		if ctx.AddVariable(NT_DV.ID, expr.Tipo, expr, NT_DV.linea, NT_DV.columna) {
			return interprete.NewNil()
		} else {
			//agregar error de que ya existe la variable
			ctx.AddError("Error: Ya existe una variable con el nombre " + NT_DV.ID + " en la linea:" + strconv.Itoa(NT_DV.linea) + " y columna:" + strconv.Itoa(NT_DV.columna))
			return interprete.NewNil()
		}
	} else {
		ctx.AddError("Error: No se puede asignar el valor de tipo " + expr.Tipo.String() + " a la variable " + NT_DV.ID + " de tipo " + NT_DV.Tipo + " en la linea:" + strconv.Itoa(NT_DV.linea) + " y columna:" + strconv.Itoa(NT_DV.columna))
	}
	return interprete.NewNil()
}
