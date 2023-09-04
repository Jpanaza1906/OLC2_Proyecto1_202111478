package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_TransSentencia struct {
	tipo string
}

// Constructor for NT_TransSentencia----------------------------------------------

func NewNT_TransSentencia(tipo string) *NT_TransSentencia {
	return &NT_TransSentencia{
		tipo: tipo,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTts *NT_TransSentencia) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ctx.AddTransSentencia(NTts.tipo)
	return interprete.NewNil()
}
