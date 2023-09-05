package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_Switch struct {
	Expresion interprete.AbstractExpression
	Casos     []interprete.AbstractExpression
	Default   interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_Switch----------------------------------------------
func NewNT_Switch(expresion interprete.AbstractExpression, casos []interprete.AbstractExpression, def interprete.AbstractExpression, linea int, columna int) *NT_Switch {
	return &NT_Switch{
		Expresion: expresion,
		Casos:     casos,
		Default:   def,
		Linea:     linea,
		Columna:   columna,
	}
}

type NT_Caso struct {
	Expresion  interprete.AbstractExpression
	Sentencias interprete.AbstractExpression
}

// Constructor for NT_Caso----------------------------------------------
func NewNT_Caso(expresion interprete.AbstractExpression, sentencias interprete.AbstractExpression) *NT_Caso {
	return &NT_Caso{
		Expresion:  expresion,
		Sentencias: sentencias,
	}
}

//implementacion de la interfaz de AbstractExpression----------------
func (NTc *NT_Caso) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return NTc.Sentencias.Interpretar(ctx)
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTs *NT_Switch) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expresion := NTs.Expresion.Interpretar(ctx)
	goodcase := false
	for _, caso := range NTs.Casos {
		caso := caso.(*NT_Caso)
		expresionCaso := caso.Expresion.Interpretar(ctx)
		if expresionCaso.ValorS == expresion.ValorS {
			ctx.PushAmbito()
			caso.Interpretar(ctx)
			ctx.PopAmbito()
			goodcase = true
			break
		}
		if ctx.Memoria.Anterior == nil {
			if len(ctx.TransState) > 0 {
				ctx.AddError("No se puede colocar una expresion de transicion fuera de un ciclo o función en el IF ubicado en la linea " + strconv.Itoa(NTs.Linea) + " y columna " + strconv.Itoa(NTs.Columna))
				return interprete.NewNil()
			}
		}
	}
	if NTs.Default != nil && !goodcase {
		ctx.PushAmbito()
		NTs.Default.Interpretar(ctx)
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
