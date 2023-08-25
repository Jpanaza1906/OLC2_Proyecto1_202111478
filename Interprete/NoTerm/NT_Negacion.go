package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

// NEGACION NOT------------------------------------------------------------------

type NT_Not struct {
	Expr interprete.AbstractExpression
}

// Constructor for NT_Not----------------------------------------------

func NewNT_Not(expr interprete.AbstractExpression) *NT_Not {
	return &NT_Not{
		Expr: expr,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_No *NT_Not) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultado := NT_No.Expr.Interpretar(ctx)

	if resultado.Nil {
		ctx.AddError("Error: No se puede negar nil")
		return interprete.NewNil()
	}

	switch resultado.Tipo {
	case interprete.Bool:
		resul := !resultado.ValorB
		return interprete.NewBoolLiteral(resul)
	}
	ctx.AddError("Error: No se puede negar " + resultado.ValorS)
	return interprete.NewNil()
}

// NEGACION MENOS UNARIO------------------------------------------------------------------

type NT_MenosUnario struct {
	Expr interprete.AbstractExpression
}

// Constructor for NT_MenosUnario----------------------------------------------

func NewNT_MenosUnario(expr interprete.AbstractExpression) *NT_MenosUnario {
	return &NT_MenosUnario{
		Expr: expr,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_MU *NT_MenosUnario) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultado := NT_MU.Expr.Interpretar(ctx)

	if resultado.Nil {
		ctx.AddError("Error: No se puede negar nil")
		return interprete.NewNil()
	}

	switch resultado.Tipo {
	case interprete.Integer:
		resul := resultado.Valor * -1
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		resul := resultado.ValorF * -1
		return interprete.NewFloatLiteral(resul)
	}
	ctx.AddError("Error: No se puede negar " + resultado.ValorS)
	return interprete.NewNil()
}
