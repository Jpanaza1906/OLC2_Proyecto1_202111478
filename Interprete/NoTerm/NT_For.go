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
			ctx.PushAmbito("For")
			// se agrega la constante al contexto
			ctx.AddVariable(NTf.Id, expr.Tipo, interprete.NewStringLiteral(""), NTf.Linea, NTf.Columa)
			for _, c := range expr.ValorS {
				// se asigna el valor de la constante
				ctx.AsigVariable(NTf.Id, interprete.NewStringLiteral(string(c)))
				// se ejecuta la sentencia
				resul := NTf.Sentencia.Interpretar(ctx)
				if !resul.Nil {
					ctx.PopAmbito()
					return resul
				}

				if len(ctx.TransState) > 0 {
					transt := ctx.TransState[len(ctx.TransState)-1]
					ctx.RemTransSentencia()
					if transt.ValorS == "break" {
						break
					} else if transt.ValorS == "continue" {
						continue
					}
				}
			}
			// se saca el ambito
			ctx.PopAmbito()
		} else if expr.Tipo == interprete.Vector {
			// se agrega nuevo ambito
			ctx.PushAmbito("For")
			// se obtiene el tipo de dato del vector
			tipo := expr.ValorV[0].Tipo

			// se agrega la constante del tipo del vector al contexto
			ctx.AddVariable(NTf.Id, tipo, interprete.NewNil(), NTf.Linea, NTf.Columa)
			for _, v := range expr.ValorV {
				// se asigna el valor de la constante
				ctx.AsigVariable(NTf.Id, &v)
				// se ejecuta la sentencia
				NTf.Sentencia.Interpretar(ctx)

				if len(ctx.TransState) > 0 {
					transt := ctx.TransState[len(ctx.TransState)-1]
					ctx.RemTransSentencia()
					if transt.ValorS == "break" {
						break
					} else if transt.ValorS == "continue" {
						continue
					}
				}
				if len(ctx.ReturnState) > 0 {
					break
				}
			}
			// se saca el ambito
			ctx.PopAmbito()
		} else {
			ctx.AddError("La expresion debe ser de tipo string o vector")
			return interprete.NewNil()
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

		// se agrega nuevo ambito
		ctx.PushAmbito("For")

		// se agrega la variable al contexto
		ctx.AddVariable(NTf.Id, interprete.Integer, interprete.NewIntLiteral(0), NTf.Linea, NTf.Columa)

		for {
			//se le asigna el valor de inicio a la variable
			ctx.AsigVariable(NTf.Id, interprete.NewIntLiteral(inicio.Valor))
			// se obtiene el valor de la variable
			valor1, ok := ctx.GetVariable(NTf.Id)
			if !ok {
				ctx.AddError("La variable " + NTf.Id + " no existe")
				return interprete.NewNil()
			}
			valor := valor1.Resultado
			// se verifica que el valor sea menor que el fin
			if valor.Valor <= fin.Valor {
				// se ejecuta la sentencia
				ctx.PushAmbito("ForSentencias")
				NTf.Sentencia.Interpretar(ctx)
				ctx.PopAmbito()
				//se limpian las variables del contexto
				//ctx.LimpiarVariables()
				// se aumenta el valor de la variable
				inicio.Valor++
			} else {
				ctx.PopAmbito()
				break
			}
			if len(ctx.TransState) > 0 {
				transt := ctx.TransState[len(ctx.TransState)-1]
				ctx.RemTransSentencia()
				if transt.ValorS == "break" {
					ctx.PopAmbito()
					break
				} else if transt.ValorS == "continue" {
					ctx.PopAmbito()
					continue
				} // else if transt == "return" {
				//ctx.PopAmbito()
				//return interprete.NewNil()
				//}
			}

		}
	}
	return interprete.NewNil()

}
