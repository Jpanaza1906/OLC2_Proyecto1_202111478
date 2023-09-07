package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_LlamadaFuncion struct {
	Id        string
	Atributos []interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_LlamadaFuncion----------------------------------------------

func NewNT_LlamadaFuncion(id string, atributos []interprete.AbstractExpression, linea int, columna int) *NT_LlamadaFuncion {
	return &NT_LlamadaFuncion{
		Id:        id,
		Atributos: atributos,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_LF *NT_LlamadaFuncion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	// se verifica que la funcion exista
	funcionL, existe := ctx.ExistFuncion(NT_LF.Id)
	if !existe {
		ctx.AddError("La funcion " + NT_LF.Id + " no existe en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
		return interprete.NewNil()
	}
	//Se crea un vector de resultados para guardar los valores de los parametros

	var valores []*interprete.Resultado

	for i := 0; i < funcionL.Nparametros; i++ {
		parametro := funcionL.Parametros[i].(*NT_Fparametro)
		atributo := NT_LF.Atributos[i].(*NT_Fargumento)

		if parametro.Id_prim != "" {
			if parametro.Id_prim == "_" {
				if atributo.Id != "" {
					ctx.AddError("El parametro " + parametro.Id_prim + " indica que no debe haber id en el atributo en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
					return interprete.NewNil()
				}
			} else if parametro.Id_prim != atributo.Id {
				ctx.AddError("El parametro " + parametro.Id_prim + " no coincide con el atributo en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
				return interprete.NewNil()
			}
		} else if parametro.Id != atributo.Id {
			ctx.AddError("El parametro " + parametro.Id + " no coincide con el atributo en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
			return interprete.NewNil()
		}
		var valorparam *interprete.Resultado
		//Se evalua si el parametro es por referencia
		if parametro.Refencia {
			// se verifica que el atributo sea una variable
			if atributo.Referencia {
				expresion := atributo.Expresion.(*NT_Identificador)
				// se obtiene el valor de la variable
				valor, ok := ctx.GetVariable(expresion.ID)
				if !ok {
					ctx.AddError("La variable " + expresion.ID + " no existe en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
					return interprete.NewNil()
				}
				valorparam = valor.Resultado
				if parametro.Tipo != valorparam.Tipo.String() {
					ctx.AddError("El tipo de dato del parametro " + parametro.Id + " no coincide con el tipo de dato del atributo en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
					return interprete.NewNil()
				}
				// se agrega una variable que referencie a la variable del contexto
				valores = append(valores, valorparam)

			} else {
				ctx.AddError("El parametro " + parametro.Id + " es por referencia y no se le puede asignar un valor en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
				return interprete.NewNil()
			}
		} else {
			valor := atributo.Expresion.Interpretar(ctx)
			valorparam = valor
			//el problema es que opera el primer valor de M, y al meterlo a la otra funcion lo opera de una vez, asi que toma el valor de la M en la nueva funcion
			if parametro.Tipo != valorparam.Tipo.String() {
				ctx.AddError("El tipo de dato del parametro " + parametro.Id + " no coincide con el tipo de dato del atributo en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
				return interprete.NewNil()
			}
			//se agrega una nueva variable al contexto
			valores = append(valores, valorparam)
		}

	}
	// se evalua que el vector de valores sea del mismo tamaño al numero de parametros
	if len(valores) != funcionL.Nparametros {
		ctx.AddError("El numero de parametros no coincide con el numero de atributos en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
		return interprete.NewNil()
	}

	// se crea un nuevo contexto para la funcion
	ctx.PushAmbito("Funcion " + NT_LF.Id)
	// se agregan las variables al contexto
	for i := 0; i < len(funcionL.Parametros); i++ {
		parametro := funcionL.Parametros[i].(*NT_Fparametro)
		valorparam := valores[i]
		ctx.AddVariable(parametro.Id, valorparam.Tipo, valorparam, NT_LF.Linea, NT_LF.Columna)
	}

	var resul *interprete.Resultado
	if funcionL.Sentencias != nil {
		// se ejecutan las sentencias de la funcion
		resul = funcionL.Sentencias.Interpretar(ctx)
	}
	// si tiene valor de retorno se retorna
	if !resul.Nil {
		ctx.PopAmbito()
		//se verifica que el resultado sea del mismo tipo que el de la funcion
		if resul.Tipo.String() != funcionL.Tipo_Retorno {
			ctx.AddError("El tipo de dato del retorno no coincide con el tipo de dato de la funcion en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
			return interprete.NewNil()
		}
		return resul
	}
	variables := ctx.Memoria.Variables
	ctx.PopAmbito()
	// se actualiza el valor de las variables
	for i := 0; i < len(funcionL.Parametros); i++ {
		atributo := NT_LF.Atributos[i].(*NT_Fargumento)
		parametro := funcionL.Parametros[i].(*NT_Fparametro)
		if atributo.Referencia {
			expresion := atributo.Expresion.(*NT_Identificador)
			// se obtiene el valor de la variable
			valor, ok := ctx.GetVariable(expresion.ID)
			if !ok {
				ctx.AddError("La variable " + expresion.ID + " no existe en la linea " + strconv.Itoa(NT_LF.Linea) + " y columna " + strconv.Itoa(NT_LF.Columna))
				return interprete.NewNil()
			}
			valor.Resultado = variables[parametro.Id].Resultado
			ctx.AsigVariable(expresion.ID, valor.Resultado)
		}
	}
	return interprete.NewNil()
}
