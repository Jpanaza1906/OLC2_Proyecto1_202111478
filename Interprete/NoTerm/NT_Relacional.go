package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

// IGUAL IGUAL ------------------------------------------------------------------

type NT_IgualIgual struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_II----------------------------------------------
func NewNT_IgualIgual(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_IgualIgual {
	return &NT_IgualIgual{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (nt_ii *NT_IgualIgual) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_ii.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_ii.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		if resultadoIzq.Valor == resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Integer:
		if resultadoIzq.Valor == resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF == resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS == resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// DIFERENTE ------------------------------------------------------------------

type NT_Diferente struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Diferente----------------------------------------------

func NewNT_Diferente(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Diferente {
	return &NT_Diferente{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression------------------------

func (nt_dif *NT_Diferente) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_dif.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_dif.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		if resultadoIzq.Valor != resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Integer:
		if resultadoIzq.Valor != resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF != resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS != resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MAYOR IGUAL ------------------------------------------------------------------

type NT_MayorIgual struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_MayorIgual----------------------------------------------

func NewNT_MayorIgual(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_MayorIgual {
	return &NT_MayorIgual{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression------------------------

func (nt_mayor *NT_MayorIgual) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_mayor.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_mayor.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede comparar bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoIzq.Valor >= resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF >= resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS >= resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MAYOR QUE ------------------------------------------------------------------

type NT_MayorQue struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_MayorQue----------------------------------------------

func NewNT_MayorQue(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_MayorQue {
	return &NT_MayorQue{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression------------------------

func (nt_mayorq *NT_MayorQue) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_mayorq.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_mayorq.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede comparar bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoIzq.Valor > resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF > resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS > resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MENOR IGUAL ------------------------------------------------------------------

type NT_MenorIgual struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_MenorIgual----------------------------------------------

func NewNT_MenorIgual(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_MenorIgual {
	return &NT_MenorIgual{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression------------------------

func (nt_menor *NT_MenorIgual) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_menor.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_menor.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede comparar bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoIzq.Valor <= resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF <= resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS <= resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MENOR QUE ------------------------------------------------------------------

type NT_MenorQue struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_MenorQue----------------------------------------------

func NewNT_MenorQue(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_MenorQue {
	return &NT_MenorQue{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression------------------------

func (nt_menorq *NT_MenorQue) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_menorq.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_menorq.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede comparar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede comparar bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoIzq.Valor < resultadoDer.Valor {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.Float:
		if resultadoIzq.ValorF < resultadoDer.ValorF {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	case interprete.String:
		if resultadoIzq.ValorS < resultadoDer.ValorS {
			return interprete.NewBoolLiteral(true)
		} else {
			return interprete.NewBoolLiteral(false)
		}
	}
	ctx.AddError("Error: No se puede comparar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}
