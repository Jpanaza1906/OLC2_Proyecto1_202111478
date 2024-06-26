package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_While struct {
	Condicion  interprete.AbstractExpression
	Sentencias interprete.AbstractExpression
	Linea      int
	Columna    int
}

// Constructor for NT_While----------------------------------------------
func NewNT_While(condicion interprete.AbstractExpression, sentencias interprete.AbstractExpression, linea int, columna int) *NT_While {
	return &NT_While{
		Condicion:  condicion,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTw *NT_While) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for {
		ctx.PushAmbito("While")
		condicion := NTw.Condicion.Interpretar(ctx)
		if condicion.Tipo != interprete.Bool {
			ctx.AddError("La condicion del while debe ser de tipo booleano")
			return interprete.NewNil()
		}
		if condicion.ValorB {
			resul := NTw.Sentencias.Interpretar(ctx)
			if !resul.Nil {
				ctx.PopAmbito()
				return resul
			}
		} else {
			ctx.PopAmbito()
			break
		}

		if len(ctx.TransState) > 0 {
			transt := ctx.TransState[len(ctx.TransState)-1]
			ctx.RemTransSentencia()
			if transt.ValorS == "break" {
				ctx.PopAmbito()
				break
			} else if transt.ValorS == "continue" {
				ctx.PopAmbito()
				continue
			}
		}
		//Se saca el ambito
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
