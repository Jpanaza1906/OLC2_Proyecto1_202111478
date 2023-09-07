package interprete

type Funcion struct {
	Nombre       string
	Ambito       string
	Nparametros  int
	Parametros   []AbstractExpression
	Tipo_Funcion string
	Tipo_Retorno string
	Sentencias   AbstractExpression
	Retorno      AbstractExpression
	Linea        int
	Columna      int
}

// Constructor for Funcion----------------------------------------------

func NewFuncion(nombre string, ambito string, nparametros int, parametros []AbstractExpression, tipo_funcion string, tipo_retorno string, sentencias AbstractExpression, retorno AbstractExpression, linea int, columna int) *Funcion {
	return &Funcion{
		Nombre:       nombre,
		Ambito:       ambito,
		Nparametros:  nparametros,
		Parametros:   parametros,
		Tipo_Funcion: tipo_funcion,
		Tipo_Retorno: tipo_retorno,
		Sentencias:   sentencias,
		Retorno:      retorno,
		Linea:        linea,
		Columna:      columna,
	}
}
