package interprete

//Clase tipos de datos--------------------------------------------------------------------------------------------

type Tipos struct {
	Cod       int
	Nombre    string
	TipoBase  int
	Padre     int
	Dimension int
	Minimo    int
	Maximo    int
	Ambito    int
}

// Constructor for Tipos----------------------------------------------
func NewTipos(cod int, nombre string, tipoBase int, padre int, dimension int, minimo int, maximo int, ambito int) *Tipos {
	return &Tipos{
		Cod:       cod,
		Nombre:    nombre,
		TipoBase:  tipoBase,
		Padre:     padre,
		Dimension: dimension,
		Minimo:    minimo,
		Maximo:    maximo,
		Ambito:    ambito,
	}
}
