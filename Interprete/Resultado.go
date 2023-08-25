package interprete

import "strconv"

type TipoE uint

const (
	Nil     TipoE = 0
	Bool    TipoE = 1
	Integer TipoE = 2
	Float   TipoE = 3
	String  TipoE = 4
)

func (t TipoE) String() string {
	switch t {
	case Nil:
		return "Nil"
	case Bool:
		return "Bool"
	case Integer:
		return "Int"
	case Float:
		return "Float"
	case String:
		return "String"
	default:
		return "Unknown"
	}
}

type Resultado struct {
	Nil    bool
	ValorB bool
	Valor  int
	ValorF float64
	ValorS string
	Tipo   TipoE
}

// Constructor para NIL RESULTADO----------------------------------------------

func NewNil() *Resultado {
	return &Resultado{
		Nil:    true,
		ValorF: 0.00,
		Valor:  0,
		ValorS: "Nil",
		Tipo:   Nil,
	}
}

// Constructor para BOOL RESULTADO----------------------------------------------

func NewBoolLiteral(valor bool) *Resultado {
	return &Resultado{
		Nil:    false,
		ValorB: valor,
		Valor:  0,
		ValorF: 0.00,
		ValorS: strconv.FormatBool(valor),
		Tipo:   Bool,
	}
}

// Constructor para INT RESULTADO----------------------------------------------

func NewIntLiteral(valor int) *Resultado {
	return &Resultado{
		Nil:    false,
		Valor:  valor,
		ValorF: float64(valor),
		ValorS: strconv.Itoa(valor),
		Tipo:   Integer,
	}
}

// Constructor para FLOAT RESULTADO----------------------------------------------

func NewFloatLiteral(valor float64) *Resultado {
	return &Resultado{
		Nil:    false,
		ValorF: valor,
		Valor:  0,
		ValorS: strconv.FormatFloat(valor, 'f', 2, 64),
		Tipo:   Float,
	}
}

// Constructor para STRING RESULTADO----------------------------------------------

func NewStringLiteral(valor string) *Resultado {
	return &Resultado{
		Nil:    false,
		ValorF: 0.00,
		Valor:  0,
		ValorS: valor,
		Tipo:   String,
	}
}

// Constructor para CHAR RESULTADO----------------------------------------------

func NewCharLiteral(valor string) *Resultado {
	return &Resultado{
		Nil:    false,
		ValorF: 0.00,
		Valor:  0,
		ValorS: valor,
		Tipo:   String,
	}
}
