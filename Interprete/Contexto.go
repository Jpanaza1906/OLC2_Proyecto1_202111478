package interprete

// Clase simbolo --------------------------------------------------------------------------------------------

type Simbolo struct {
	Nombre    string
	Resultado *Resultado
	TipoSimb  string
	linea     int
	columna   int
}

// Constructor for Simbolo----------------------------------------------

func NewSimbolo(nombre string, resultado *Resultado, tiposimb string, linea int, columna int) *Simbolo {
	return &Simbolo{
		Nombre:    nombre,
		Resultado: resultado,
		TipoSimb:  tiposimb,
		linea:     linea,
		columna:   columna,
	}
}

// Clase Contexto --------------------------------------------------------------------------------------------

type Contexto struct {
	Memoria   *Memoria
	zGlobal   *Memoria
	Consola   string
	Errores   []string
	Conversor *Conversor
}

func NewContexto() *Contexto {
	c := &Contexto{
		Memoria:   NewMemoria(nil),
		Consola:   "",
		Errores:   make([]string, 0, 10),
		Conversor: nil,
	}
	c.Conversor = NewConversor(c)
	return c
}

// Se crea un nuevo ambito ----------------------------------------------

func (c *Contexto) PushAmbito() {
	head_Memoria := NewMemoria(c.Memoria)
	c.Memoria = head_Memoria
}

// Se elimina el ambito actual ----------------------------------------------

func (c *Contexto) PopAmbito() {
	if c.Memoria != nil {
		c.Memoria = c.Memoria.Anterior
	} else {
		c.AddError("Error: No se puede eliminar el ambito global")
	}
}

// Se crea una nueva variable en el ambito actual ----------------------------------------------

func (c *Contexto) AddVariable(nombre string, expresion *Resultado, tiposimb string, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false
	}
	return c.Memoria.CreateSimbolo(nombre, expresion, tiposimb, linea, columna)
}

// Se asigna un valor a una variable en el ambito actual ---------------------------------------------

func (c *Contexto) AsigVariable(nombre string, expresion *Resultado) bool {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Mem = c.Memoria.Anterior
		for aux_Mem != nil {
			existe = aux_Mem.Exist(nombre)
			if existe {
				if aux_Mem.SetSimbolo(nombre, expresion) {
					return true
				} else {
					c.AddError("Error: No se le puede asignar un valor a una constante.")
					return false
				}
			}
			aux_Mem = aux_Mem.Anterior
		}
		return false
	}
	if c.Memoria.SetSimbolo(nombre, expresion) {
		return true
	} else {
		c.AddError("Error: No se le puede asignar un valor a una constante.")
		return false
	}
}

// Se obtiene el valor de una variable en el ambito actual ----------------------------------------------

func (c *Contexto) GetVariable(nombre string) (*Resultado, bool) {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Mem = c.Memoria.Anterior
		for aux_Mem != nil {
			existe = aux_Mem.Exist(nombre)
			if existe {
				return aux_Mem.GetSimbolo(nombre)
			}
			aux_Mem = aux_Mem.Anterior
		}
		return NewNil(), false
	}
	return c.Memoria.GetSimbolo(nombre)
}

// Print en consola ----------------------------------------------

func (c *Contexto) Print(entrada string) {
	c.Consola += entrada
}

//Se agrega un error al contexto----------------------------------------------

func (c *Contexto) AddError(entrada string) {
	c.Errores = append(c.Errores, entrada)
}
