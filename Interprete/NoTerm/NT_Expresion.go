package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

// IDENTIFICADOR ------------------------------------------------------------------

type NT_Identificador struct {
	ID      string
	linea   int
	columna int
}

// Constructor for NT_Identificador----------------------------------------------

func NewNT_Identificador(ID string, linea int, columna int) *NT_Identificador {
	return &NT_Identificador{
		ID:      ID,
		linea:   linea,
		columna: columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (NT_Id *NT_Identificador) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr, ok := ctx.GetVariable(NT_Id.ID)
	if !ok {
		ctx.AddError("La variable " + NT_Id.ID + " no existe" + " en la linea:" + strconv.Itoa(NT_Id.linea) + " y columna:" + strconv.Itoa(NT_Id.columna))
		return interprete.NewNil()
	}
	return expr
}

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
	//exprIzq := NT_mul.ExprIzquierda.Interpretar(ctx)
	//exprDer := NT_mul.ExprDerecha.Interpretar(ctx)
	return nil

}
