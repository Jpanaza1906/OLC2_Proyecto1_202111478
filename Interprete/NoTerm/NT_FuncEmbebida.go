package noterm

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"strconv"
)

type NT_FuncEmbebida struct {
	Nfunc     string
	Expresion interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_FuncEmbebida----------------------------------------------

func NewNT_FuncEmbebida(nfunc string, expresion interprete.AbstractExpression, linea int, columna int) *NT_FuncEmbebida {
	return &NT_FuncEmbebida{
		Nfunc:     nfunc,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------

func (fe *NT_FuncEmbebida) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expr := fe.Expresion.Interpretar(ctx)

	switch fe.Nfunc {
	case "Int":
		switch expr.Tipo {
		case interprete.Bool:
			if expr.ValorB {
				return interprete.NewIntLiteral(1)
			} else {
				return interprete.NewIntLiteral(0)
			}
		case interprete.Integer:
			return interprete.NewIntLiteral(expr.Valor)
		case interprete.Float:
			return interprete.NewIntLiteral(int(expr.ValorF))
		case interprete.String:
			num, err := strconv.Atoi(expr.ValorS)
			if err != nil {
				num2, err2 := strconv.ParseFloat(expr.ValorS, 64)
				if err2 != nil {
					ctx.AddError("Error: No se puede convertir el String a Int en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
					return interprete.NewNil()
				}
				return interprete.NewIntLiteral(int(num2))
			}
			return interprete.NewIntLiteral(num)
		default:
			ctx.AddError("Error: No se puede convertir el tipo de dato a Int en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
			return interprete.NewNil()
		}
	case "Float":
		switch expr.Tipo {
		case interprete.Bool:
			if expr.ValorB {
				return interprete.NewFloatLiteral(1.00)
			} else {
				return interprete.NewFloatLiteral(0.00)
			}
		case interprete.Integer:
			return interprete.NewFloatLiteral(float64(expr.Valor))
		case interprete.Float:
			return interprete.NewFloatLiteral(expr.ValorF)
		case interprete.String:
			num, err := strconv.ParseFloat(expr.ValorS, 64)
			if err != nil {
				ctx.AddError("Error: No se puede convertir el String a Float en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
				return interprete.NewNil()
			}
			return interprete.NewFloatLiteral(num)
		default:
			ctx.AddError("Error: No se puede convertir el tipo de dato a Int en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
			return interprete.NewNil()
		}
	case "String":
		switch expr.Tipo {
		case interprete.Bool:
			return interprete.NewStringLiteral(expr.ValorS)
		case interprete.Integer:
			return interprete.NewStringLiteral(expr.ValorS)
		case interprete.Float:
			return interprete.NewStringLiteral(expr.ValorS)
		case interprete.String:
			return interprete.NewStringLiteral(expr.ValorS)
		default:
			ctx.AddError("Error: No se puede convertir el tipo de dato a Int en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
			return interprete.NewNil()
		}
	default:
		ctx.AddError("Error: No se reconoce la funcion embebida en la linea:" + strconv.Itoa(fe.Linea) + " y columna:" + strconv.Itoa(fe.Columna))
	}
	return interprete.NewNil()
}
