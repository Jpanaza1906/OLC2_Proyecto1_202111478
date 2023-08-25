package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

// AND ------------------------------------------------------------------

type NT_And struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_And----------------------------------------------

func NewNT_And(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_And {
	return &NT_And{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (nt_and *NT_And) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_and.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_and.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede and nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		resul := resultadoIzq.ValorB && resultadoDer.ValorB
		return interprete.NewBoolLiteral(resul)
	}
	ctx.AddError("Error: No se puede and " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// OR ------------------------------------------------------------------

type NT_Or struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Or----------------------------------------------

func NewNT_Or(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Or {
	return &NT_Or{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (nt_or *NT_Or) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_or.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_or.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede or nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		resul := resultadoIzq.ValorB || resultadoDer.ValorB
		return interprete.NewBoolLiteral(resul)
	}
	ctx.AddError("Error: No se puede or " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}
