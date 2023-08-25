package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

// SUMA ------------------------------------------------------------------

type NT_Suma struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Sumar----------------------------------------------
func NewNT_Suma(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Suma {
	return &NT_Suma{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (nt_sum *NT_Suma) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_sum.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_sum.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede sumar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede sumar bool")
		return interprete.NewNil()
	case interprete.Integer:
		resul := resultadoIzq.Valor + resultadoDer.Valor
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		resul := resultadoIzq.ValorF + resultadoDer.ValorF
		return interprete.NewFloatLiteral(resul)
	case interprete.String:
		resul := resultadoIzq.ValorS + resultadoDer.ValorS
		return interprete.NewStringLiteral(resul)
	}
	ctx.AddError("Error: No se puede sumar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// RESTA ------------------------------------------------------------------

type NT_Resta struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Resta----------------------------------------------

func NewNT_Resta(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Resta {
	return &NT_Resta{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression--------------------

func (nt_res *NT_Resta) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := nt_res.ExprIzquierda.Interpretar(ctx)
	resultadoDer := nt_res.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede restar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede restar bool")
		return interprete.NewNil()
	case interprete.Integer:
		resul := resultadoIzq.Valor - resultadoDer.Valor
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		resul := resultadoIzq.ValorF - resultadoDer.ValorF
		return interprete.NewFloatLiteral(resul)
	case interprete.String:
		ctx.AddError("Error: No se puede restar string")
		return interprete.NewNil()
	}
	ctx.AddError("Error: No se puede restar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MULTIPLICACION ------------------------------------------------------------------

type NT_Multiplicacion struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Multiplicacion----------------------------------------------
func NewNT_Multiplicacion(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Multiplicacion {
	return &NT_Multiplicacion{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_mul *NT_Multiplicacion) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := NT_mul.ExprIzquierda.Interpretar(ctx)
	resultadoDer := NT_mul.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede multiplicar nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede multiplicar bool")
		return interprete.NewNil()
	case interprete.Integer:
		resul := resultadoIzq.Valor * resultadoDer.Valor
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		resul := resultadoIzq.ValorF * resultadoDer.ValorF
		return interprete.NewFloatLiteral(resul)
	case interprete.String:
		ctx.AddError("Error: No se puede multiplicar string")
		return interprete.NewNil()
	}
	ctx.AddError("Error: No se puede multiplicar " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()

}

// DIVISION ------------------------------------------------------------------

type NT_Division struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Division----------------------------------------------

func NewNT_Division(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Division {
	return &NT_Division{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression--------------------

func (NT_div *NT_Division) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := NT_div.ExprIzquierda.Interpretar(ctx)
	resultadoDer := NT_div.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede dividir nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede dividir bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoDer.Valor == 0 {
			ctx.AddError("Error: No se puede dividir entre 0")
			return interprete.NewNil()
		}
		resul := resultadoIzq.Valor / resultadoDer.Valor
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		if resultadoDer.ValorF == 0 {
			ctx.AddError("Error: No se puede dividir entre 0")
			return interprete.NewNil()
		}
		resul := resultadoIzq.ValorF / resultadoDer.ValorF
		return interprete.NewFloatLiteral(resul)
	case interprete.String:
		ctx.AddError("Error: No se puede dividir string")
		return interprete.NewNil()
	}
	ctx.AddError("Error: No se puede dividir " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}

// MODULO ------------------------------------------------------------------

type NT_Modulo struct {
	ExprIzquierda interprete.AbstractExpression
	ExprDerecha   interprete.AbstractExpression
}

// Constructor for NT_Modulo----------------------------------------------

func NewNT_Modulo(exprIzquierda interprete.AbstractExpression, exprDerecha interprete.AbstractExpression) *NT_Modulo {
	return &NT_Modulo{
		ExprIzquierda: exprIzquierda,
		ExprDerecha:   exprDerecha,
	}
}

// Implementacion de la interfaz de AbstractExpression--------------------

func (NT_mod *NT_Modulo) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	resultadoIzq := NT_mod.ExprIzquierda.Interpretar(ctx)
	resultadoDer := NT_mod.ExprDerecha.Interpretar(ctx)

	if resultadoIzq.Tipo != resultadoDer.Tipo {
		if resultadoIzq.Tipo > resultadoDer.Tipo {
			resultadoDer = ctx.Conversor.Ampliar(resultadoDer, resultadoIzq.Tipo)

		} else {
			resultadoIzq = ctx.Conversor.Ampliar(resultadoIzq, resultadoDer.Tipo)
		}
	}

	if resultadoDer.Nil || resultadoIzq.Nil {
		ctx.AddError("Error: No se puede realizar modulo nil")
		return interprete.NewNil()
	}

	switch resultadoIzq.Tipo {
	case interprete.Bool:
		ctx.AddError("Error: No se puede realizar modulo bool")
		return interprete.NewNil()
	case interprete.Integer:
		if resultadoDer.Valor == 0 {
			ctx.AddError("Error: No se puede realizar modulo entre 0")
			return interprete.NewNil()
		}
		resul := resultadoIzq.Valor % resultadoDer.Valor
		return interprete.NewIntLiteral(resul)
	case interprete.Float:
		if resultadoDer.ValorF == 0 {
			ctx.AddError("Error: No se puede realizar modulo entre 0")
			return interprete.NewNil()
		}
		resul := int(resultadoIzq.ValorF) % int(resultadoDer.ValorF)
		return interprete.NewIntLiteral(resul)
	case interprete.String:
		ctx.AddError("Error: No se puede realizar modulo string")
		return interprete.NewNil()
	}
	ctx.AddError("Error: No se puede realizar modulo " + resultadoIzq.ValorS + " y " + resultadoDer.ValorS)
	return interprete.NewNil()
}
