package interprete

import "strconv"

type Conversor struct {
	ctx *Contexto
}

// Constructor for Conversor----------------------------------------------

func NewConversor(ctx *Contexto) *Conversor {
	return &Conversor{
		ctx: ctx,
	}
}

func (c *Conversor) Ampliar(res *Resultado, tipo TipoE) *Resultado {
	switch tipo {
	case Nil:
		c.ctx.AddError("Error: No se puede convertir nil")
		return NewNil()
	case Bool:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNil()
		case Bool:
			return res
		case Integer:
			c.ctx.AddError("Error: No se puede convertir int a bool")
			return NewNil()
		case Float:
			c.ctx.AddError("Error: No se puede convertir float a bool")
			return NewNil()
		case String:
			c.ctx.AddError("Error: No se puede convertir string a bool")
			return NewNil()
		}
	case Integer:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNil()
		case Bool:
			c.ctx.AddError("Error: No se puede convertir bool a int")
			return NewNil()
		case Integer:
			return res
		case Float:
			respuesta := NewIntLiteral(int(res.ValorF))
			return respuesta
		case String:
			valori, err := strconv.Atoi(res.ValorS)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.ValorS + " a int")
				return NewNil()
			}
			respuesta := NewIntLiteral(valori)
			return respuesta
		}
	case Float:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNil()
		case Bool:
			c.ctx.AddError("Error: No se puede convertir bool a float")
			return NewNil()
		case Integer:
			respuesta := NewFloatLiteral(float64(res.Valor))
			return respuesta
		case Float:
			return res
		case String:
			valorf, err := strconv.ParseFloat(res.ValorS, 64)
			if err != nil {
				c.ctx.AddError("Error: No se puede convertir " + res.ValorS + " a float")
				return NewNil()
			}
			respuesta := NewFloatLiteral(valorf)
			return respuesta
		}
	case String:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("Error: No se puede convertir nil")
			return NewNil()
		case Bool:
			if res.ValorB {
				return NewStringLiteral("true")
			} else {
				return NewStringLiteral("false")
			}
		case Integer:
			respuesta := NewStringLiteral(strconv.Itoa(res.Valor))
			return respuesta
		case Float:
			respuesta := NewStringLiteral(strconv.FormatFloat(res.ValorF, 'f', 2, 64))
			return respuesta
		case String:
			return res
		}
	}
	c.ctx.AddError("Error: No se puede convertir " + res.ValorS + " a " + tipo.String())
	return NewNil()
}

/*
switch tipo {
case Nil:
case Bool:
case Integer:
case Float:
case String:
default:
}*/
