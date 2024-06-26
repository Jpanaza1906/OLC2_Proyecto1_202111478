package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_LSentencias struct {
	LSentencias []interprete.AbstractExpression
}

// Constructor for NT_LSentencias----------------------------------------------
func NewNT_LSentencias() *NT_LSentencias {
	return &NT_LSentencias{
		LSentencias: make([]interprete.AbstractExpression, 0),
	}
}

// Funcion que añade sentencias -----------------------------------------------
func (ls *NT_LSentencias) AddSentencia(sentencia interprete.AbstractExpression) {
	ls.LSentencias = append(ls.LSentencias, sentencia)
}

// Implementacion de la interfaz de AbstractExpression----------------
func (ls *NT_LSentencias) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, sentencia := range ls.LSentencias {
		if len(ctx.TransState) > 0 {
			break
		}
		resul := sentencia.Interpretar(ctx)
		if !resul.Nil {
			return resul
		}
	}
	return interprete.NewNil()
}

//Implementacion para interpretar la ultima sentencia de una funcion
func (ls *NT_LSentencias) InterpretarUltima(ctx *interprete.Contexto) *interprete.Resultado {
	if len(ls.LSentencias) > 0 {
		return ls.LSentencias[len(ls.LSentencias)-1].Interpretar(ctx)
	}
	return interprete.NewNil()
}
