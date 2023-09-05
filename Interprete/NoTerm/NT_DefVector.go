package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_DefVector struct {
	ID       string
	ListaNum []interprete.AbstractExpression
}

// Constructor for NT_DefVector----------------------------------------------

func NewNT_DefVector(ID string, ListaNum []interprete.AbstractExpression) *NT_DefVector {
	return &NT_DefVector{
		ID:       ID,
		ListaNum: ListaNum,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NTdv *NT_DefVector) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	// Si viene ID, se obtiene la expresion del ID
	if NTdv.ID != "" {
		expr, ok := ctx.GetVariable(NTdv.ID)
		if !ok {
			ctx.AddError("La variable " + NTdv.ID + " no existe")
			return interprete.NewNil()
		}
		return expr
	}
	//se crea un vector del tipo de la expresion
	var vector []interprete.Resultado

	//se recorre la lista de expresiones
	for _, expresion := range NTdv.ListaNum {
		//se obtiene el valor de la expresion
		valor := expresion.Interpretar(ctx)
		//se agrega el valor al vector
		vector = append(vector, *valor)
	}
	//se retorna el vector
	return interprete.NewVectorLiteral(vector)
}
