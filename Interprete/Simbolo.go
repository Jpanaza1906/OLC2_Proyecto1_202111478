package interprete

// Clase simbolo --------------------------------------------------------------------------------------------

type Simbolo struct {
	Nombre         string
	Ambito         string
	Categoria      string
	Tipo           TipoE
	Tipo_Compuesto TipoE
	Num_Parametros int
	Parametros     []*Simbolo
	Resultado      *Resultado
	Tipo_Retorno   *Tipos
	Linea          int
	Columna        int
}

// Constructor for Simbolo----------------------------------------------

func NewSimbolo(nombre string, ambito string, categoria string, tipo TipoE, tipo_compuesto TipoE, num_parametros int, parametros []*Simbolo, resultado *Resultado, tipo_retorno *Tipos, linea int, columna int) *Simbolo {
	return &Simbolo{
		Nombre:         nombre,
		Ambito:         ambito,
		Categoria:      categoria,
		Tipo:           tipo,
		Tipo_Compuesto: tipo_compuesto,
		Num_Parametros: num_parametros,
		Parametros:     parametros,
		Resultado:      resultado,
		Tipo_Retorno:   tipo_retorno,
		Linea:          linea,
		Columna:        columna,
	}
}
