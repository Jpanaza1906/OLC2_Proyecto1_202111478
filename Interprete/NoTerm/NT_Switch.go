package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Switch struct {
	Expresion interprete.AbstractExpression
	Casos     []interprete.AbstractExpression
	Default   interprete.AbstractExpression
}

// Constructor for NT_Switch----------------------------------------------
func NewNT_Switch(expresion interprete.AbstractExpression, casos []interprete.AbstractExpression, def interprete.AbstractExpression) *NT_Switch {
	return &NT_Switch{
		Expresion: expresion,
		Casos:     casos,
		Default:   def,
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
	}
	if NTs.Default != nil && !goodcase {
		ctx.PushAmbito()
		NTs.Default.Interpretar(ctx)
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
