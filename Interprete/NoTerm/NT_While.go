package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_While struct {
	Condicion  interprete.AbstractExpression
	Sentencias interprete.AbstractExpression
}

// Constructor for NT_While----------------------------------------------
func NewNT_While(condicion interprete.AbstractExpression, sentencias interprete.AbstractExpression) *NT_While {
	return &NT_While{
		Condicion:  condicion,
		Sentencias: sentencias,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTw *NT_While) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for {
		ctx.PushAmbito()
		condicion := NTw.Condicion.Interpretar(ctx)
		if condicion.Tipo != interprete.Bool {
			ctx.AddError("La condicion del while debe ser de tipo booleano")
			return interprete.NewNil()
		}
		if condicion.ValorB {
			NTw.Sentencias.Interpretar(ctx)
		} else {
			break
		}

		if len(ctx.TransState) > 0 {
			transt := ctx.TransState[len(ctx.TransState)-1]
			ctx.RemTransSentencia()
			if transt == "break" {
				ctx.PopAmbito()
				break
			} else if transt == "continue" {
				ctx.PopAmbito()
				continue
			} //else if transt == "return" {
			//ctx.PopAmbito()
			//	return interprete.NewNil()
			//}
		}
		//Se saca el ambito
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
