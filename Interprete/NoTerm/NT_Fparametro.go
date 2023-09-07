package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Fparametro struct {
	Id_prim  string
	Id       string
	Refencia bool
	Tipo     string
}

// Constructor for NT_Fparametro----------------------------------------------
func NewNT_Fparametro(id_prim string, id string, referencia bool, tipo string) *NT_Fparametro {
	return &NT_Fparametro{
		Id_prim:  id_prim,
		Id:       id,
		Refencia: referencia,
		Tipo:     tipo,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_FP *NT_Fparametro) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewNil()
}
