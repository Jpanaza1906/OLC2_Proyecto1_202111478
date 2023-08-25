package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_DecConst struct {
	ID        string
	Tipo      string
	Expresion interprete.AbstractExpression
	linea     int
	columna   int
}

// Constructor for NT_DecConst----------------------------------------------
func NewNT_DecConst(ID string, tipo string, expresion interprete.AbstractExpression, linea int, columna int) *NT_DecConst {
	return &NT_DecConst{
		ID:        ID,
		Tipo:      tipo,
		Expresion: expresion,
		linea:     linea,
		columna:   columna,
	}
}

func (NT_DC *NT_DecConst) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	var expr = interprete.NewNil()
	if NT_DC.Expresion != nil {
		expr = NT_DC.Expresion.Interpretar(ctx)
	}
	ctx.AddVariable(NT_DC.ID, expr, "const", NT_DC.linea, NT_DC.columna)
	return interprete.NewNil()
}
