package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_For struct {
	Id        string
	Expresion interprete.AbstractExpression
	Sentencia interprete.AbstractExpression
	Rango     []interprete.AbstractExpression
	Linea     int
	Columa    int
}

// Constructor for NT_For----------------------------------------------
func NewNT_For(id string, expresion interprete.AbstractExpression, rango []interprete.AbstractExpression, sentencia interprete.AbstractExpression, linea int, columna int) *NT_For {
	return &NT_For{
		Id:        id,
		Expresion: expresion,
		Rango:     rango,
		Sentencia: sentencia,
		Linea:     linea,
		Columa:    columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTf *NT_For) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	// se verifica si viene una expresion o un rango
	if NTf.Expresion != nil {
		// se obtiene el valor de la expresion
		expr := NTf.Expresion.Interpretar(ctx)

		// se verifica que el id no exista en el contexto
		_, ok := ctx.GetVariable(NTf.Id)
		if ok {
			ctx.AddError("La variable " + NTf.Id + " ya existe")
			return interprete.NewNil()
		}

		// si la expresion es de tipo string se recorre caracter por caracter
		if expr.Tipo == interprete.String {
			// se agrega nuevo ambito
			ctx.PushAmbito()
			// se agrega la constante al contexto
			ctx.AddVariable(NTf.Id, expr.Tipo, interprete.NewStringLiteral(""), NTf.Linea, NTf.Columa)
			for _, c := range expr.ValorS {
				// se asigna el valor de la constante
				ctx.AsigVariable(NTf.Id, interprete.NewStringLiteral(string(c)))
				// se ejecuta la sentencia
				NTf.Sentencia.Interpretar(ctx)
			}
			// se saca el ambito
			ctx.PopAmbito()
		}

	} else {
		// se obtienen los rangos
		inicio := NTf.Rango[0].Interpretar(ctx)
		fin := NTf.Rango[1].Interpretar(ctx)

		// se verifica que los rangos sean de tipo int
		if inicio.Tipo != interprete.Integer || fin.Tipo != interprete.Integer {
			ctx.AddError("Los rangos deben ser de tipo int")
			return interprete.NewNil()
		}

		// se verifica que el id no exista en el contexto
		_, ok := ctx.GetVariable(NTf.Id)
		if ok {
			ctx.AddError("La variable " + NTf.Id + " ya existe")
			return interprete.NewNil()
		}
		// se verifica que el inicio sea menor que el fin
		if inicio.Valor > fin.Valor {
			ctx.AddError("El inicio del for debe ser menor que el fin")
			return interprete.NewNil()
		}

		for {
			// se agrega nuevo ambito
			ctx.PushAmbito()
			// se agrega la variable al contexto
			ctx.AddVariable(NTf.Id, interprete.Integer, interprete.NewIntLiteral(0), NTf.Linea, NTf.Columa)
			//se le asigna el valor de inicio a la variable
			ctx.AsigVariable(NTf.Id, interprete.NewIntLiteral(inicio.Valor))
			// se obtiene el valor de la variable
			valor, _ := ctx.GetVariable(NTf.Id)
			// se verifica que el valor sea menor que el fin
			if valor.Valor <= fin.Valor {
				// se ejecuta la sentencia
				NTf.Sentencia.Interpretar(ctx)
				// se aumenta el valor de la variable
				inicio.Valor++
			} else {
				break
			}
			if len(ctx.TransState) > 0 {
				transt := ctx.TransState[len(ctx.TransState)-1]
				ctx.RemTransSentencia()
				if transt == "break" {
					ctx.PopAmbito()
					break
				} else if transt == "continue" {
					ctx.PopAmbito()
					continue
				} // else if transt == "return" {
				//ctx.PopAmbito()
				//return interprete.NewNil()
				//}
			}
			ctx.PopAmbito()

		}
	}
	return interprete.NewNil()

}