package interprete

// Clase simbolo --------------------------------------------------------------------------------------------

type Simbolo struct {
	Nombre         string
	Categoria      string
	Tipo           TipoE
	Tipo_Compuesto *Tipos
	Num_Parametros int
	Parametros     []*Simbolo
	Resultado      *Resultado
	Resultados     []*Resultado
	Tipo_Retorno   *Tipos
	Linea          int
	Columna        int
}

// Constructor for Simbolo----------------------------------------------

func NewSimbolo(nombre string, categoria string, tipo TipoE, tipo_compuesto *Tipos, num_parametros int, parametros []*Simbolo, resultado *Resultado, resultados []*Resultado, tipo_retorno *Tipos, linea int, columna int) *Simbolo {
	return &Simbolo{
		Nombre:         nombre,
		Categoria:      categoria,
		Tipo:           tipo,
		Tipo_Compuesto: tipo_compuesto,
		Num_Parametros: num_parametros,
		Parametros:     parametros,
		Resultado:      resultado,
		Resultados:     resultados,
		Tipo_Retorno:   tipo_retorno,
		Linea:          linea,
		Columna:        columna,
	}
}
