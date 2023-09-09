package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_Atributo struct {
	TipoAtr   string
	Id        string
	Tipo      string
	Expresion interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_Atributo----------------------------------------------

func NewNT_Atributo(tipoatr string, Id string, tipo string, expr interprete.AbstractExpression, Linea int, Columna int) *NT_Atributo {
	return &NT_Atributo{
		TipoAtr:   tipoatr,
		Id:        Id,
		Tipo:      tipo,
		Expresion: expr,
		Linea:     Linea,
		Columna:   Columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_Atr *NT_Atributo) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()

	//Se ve que la expresion no sea nula
	if NT_Atr.Expresion != nil {
		expr = NT_Atr.Expresion.Interpretar(ctx)
	}

	//Si no tiene tipo, se le asigna el tipo de la expresion

	if NT_Atr.Tipo == "" {
		NT_Atr.Tipo = expr.Tipo.String()
	} else {
		switch NT_Atr.Tipo {
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
	if NT_Atr.Expresion == nil {
		//Se evalua el tipo del atributo, si es var o let
		if NT_Atr.TipoAtr == "var" {
			//Se agrega la variable al contexto
			if ctx.AddVariable(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una variable con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		} else if NT_Atr.TipoAtr == "let" {
			// se agrega la constante al contexto
			if ctx.AddConstante(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una constante con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		}
	}

	//Si el tipo de la expresion es igual al tipo de la variable, se agrega la variable al contexto

	if NT_Atr.Tipo == expr.Tipo.String() {
		if NT_Atr.TipoAtr == "var" {
			if ctx.AddVariable(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una variable con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		} else if NT_Atr.TipoAtr == "let" {
			if ctx.AddConstante(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una constante con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		}
	} else if NT_Atr.Tipo == interprete.Float.String() && expr.Tipo.String() == interprete.Integer.String() {
		expr = interprete.NewFloatLiteral(float64(expr.Valor))
		if NT_Atr.TipoAtr == "var" {
			if ctx.AddVariable(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una variable con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		} else if NT_Atr.TipoAtr == "let" {
			if ctx.AddConstante(NT_Atr.Id, expr.Tipo, expr, NT_Atr.Linea, NT_Atr.Columna) {
				return interprete.NewNil()
			} else {
				//agregar error de que ya existe la variable
				ctx.AddError("Error: Ya existe una constante con el nombre " + NT_Atr.Id + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
				return interprete.NewNil()
			}
		}
	} else {
		ctx.AddError("Error: No se puede asignar el valor de tipo " + expr.Tipo.String() + " a la variable " + NT_Atr.Id + " de tipo " + NT_Atr.Tipo + " en la Linea:" + strconv.Itoa(NT_Atr.Linea) + " y Columna:" + strconv.Itoa(NT_Atr.Columna))
	}
	return interprete.NewNil()
}
