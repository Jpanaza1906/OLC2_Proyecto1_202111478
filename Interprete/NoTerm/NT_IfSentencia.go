package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_IfSentencia struct {
	Condicion interprete.AbstractExpression
	Si        interprete.AbstractExpression
	Sino      interprete.AbstractExpression
}

// Constructor for NT_IfSentencia----------------------------------------------
func NewNT_IfSentencia(condicion interprete.AbstractExpression, si interprete.AbstractExpression, sino interprete.AbstractExpression) *NT_IfSentencia {
	return &NT_IfSentencia{
		Condicion: condicion,
		Si:        si,
		Sino:      sino,
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
		ctx.PushAmbito()
		NTis.Si.Interpretar(ctx)
		//Se saca el ambito
		ctx.PopAmbito()
	} else if NTis.Sino != nil {
		ctx.PushAmbito()
		NTis.Sino.Interpretar(ctx)
		//Se saca el ambito
		ctx.PopAmbito()
	}

	return interprete.NewNil()
}
