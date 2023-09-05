package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_TransSentencia struct {
	tipo *interprete.Resultado
}

// Constructor for NT_TransSentencia----------------------------------------------

func NewNT_TransSentencia(tipo *interprete.Resultado) *NT_TransSentencia {
	return &NT_TransSentencia{
		tipo: tipo,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTts *NT_TransSentencia) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ctx.AddTransSentencia(NTts.tipo)
	return interprete.NewNil()
}
