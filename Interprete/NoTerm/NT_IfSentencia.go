package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_IfSentencia struct {
	Condicion interprete.AbstractExpression
	Si        interprete.AbstractExpression
	Sino      interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_IfSentencia----------------------------------------------
func NewNT_IfSentencia(condicion interprete.AbstractExpression, si interprete.AbstractExpression, sino interprete.AbstractExpression, linea int, columna int) *NT_IfSentencia {
	return &NT_IfSentencia{
		Condicion: condicion,
		Si:        si,
		Sino:      sino,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTis *NT_IfSentencia) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	condicion := NTis.Condicion.Interpretar(ctx)
	if condicion.Tipo != interprete.Bool {
		ctx.AddError("La condicion del if debe ser de tipo booleano")
		return interprete.NewNil()
	}
	//Se coloca un nuevo ambito

	if condicion.ValorB {
		ctx.PushAmbito("If")
		NTis.Si.Interpretar(ctx)
		//Se saca el ambito
		ctx.PopAmbito()
	} else if NTis.Sino != nil {
		ctx.PushAmbito("Else")
		NTis.Sino.Interpretar(ctx)
		//Se saca el ambito
		ctx.PopAmbito()
	}

	//Se evalua si hay un ambito padre que no sea el global. Y si hay una expresion de transicion habra un error
	if ctx.Memoria.Anterior == nil {
		if len(ctx.TransState) > 0 {
			ctx.AddError("No se puede colocar una expresion de transicion fuera de un ciclo o función en el IF ubicado en la linea " + strconv.Itoa(NTis.Linea) + " y columna " + strconv.Itoa(NTis.Columna))
			return interprete.NewNil()
		}
	}

	return interprete.NewNil()
}
