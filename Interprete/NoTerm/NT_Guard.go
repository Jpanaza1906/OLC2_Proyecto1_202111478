package noterm

import interprete "OLC2_Proyecto1_202111478/Interprete"

type NT_Guard struct {
	Expresion interprete.AbstractExpression
	Sentencia interprete.AbstractExpression
	TransSent interprete.AbstractExpression
	Linea     int
	Columna   int
}

// Constructor for NT_Guard----------------------------------------------
func NewNT_Guard(expresion interprete.AbstractExpression, sentencia interprete.AbstractExpression, transsent interprete.AbstractExpression, linea int, columna int) *NT_Guard {
	return &NT_Guard{
		Expresion: expresion,
		Sentencia: sentencia,
		TransSent: transsent,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion de la interfaz de AbstractExpression----------------
func (NTg *NT_Guard) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	expresion := NTg.Expresion.Interpretar(ctx)
	if expresion.Tipo != interprete.Bool {
		ctx.AddError("La expresion del guard debe ser de tipo booleano")
		return interprete.NewNil()
	}
	//Se coloca un nuevo ambito
	if !expresion.ValorB {
		ctx.PushAmbito("Guard")
		NTg.Sentencia.Interpretar(ctx)
		//Se coloca la transicion de sentencia
		ctx.AddTransSentencia(NTg.TransSent.Interpretar(ctx))
		//Se saca el ambito
		ctx.PopAmbito()
	}
	return interprete.NewNil()
}
