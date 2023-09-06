package interprete

// Clase Contexto --------------------------------------------------------------------------------------------

type Contexto struct {
	Memoria    *Memoria
	ZGlobal    []*Memoria
	Consola    string
	TransState []*Resultado
	Errores    []string
	Conversor  *Conversor
}

func NewContexto() *Contexto {
	c := &Contexto{
		Memoria:    NewMemoria(nil, "Global"),
		ZGlobal:    make([]*Memoria, 0),
		Consola:    "",
		TransState: make([]*Resultado, 0, 10),
		Errores:    make([]string, 0, 10),
		Conversor:  nil,
	}
	c.Conversor = NewConversor(c)
	return c
}

// Se crea un nuevo ambito ----------------------------------------------

func (c *Contexto) PushAmbito(nombre string) {
	head_Memoria := NewMemoria(c.Memoria, c.Memoria.Nambito+">"+nombre)
	c.Memoria = head_Memoria
	c.ZGlobal = append(c.ZGlobal, c.Memoria)
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

func (c *Contexto) AddVariable(nombre string, tipo TipoE, resultado *Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false
	}
	return c.Memoria.CreateSimbolo(nombre, "var", tipo, 0, -1, nil, resultado, nil, linea, columna)
}

// Se crea una nueva constante en el ambito actual ----------------------------------------------
func (c *Contexto) AddConstante(nombre string, tipo TipoE, resultado *Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false
	}
	return c.Memoria.CreateSimbolo(nombre, "const", tipo, 0, -1, nil, resultado, nil, linea, columna)
}

// Se crea un vector en el ambito actual ----------------------------------------------
func (c *Contexto) AddVector(nombre string, categoria string, tipo TipoE, tipoV TipoE, resultado *Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false
	}
	return c.Memoria.CreateSimbolo(nombre, categoria, tipo, tipoV, -1, nil, resultado, nil, linea, columna)
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

//Agregar sentencia
func (c *Contexto) AddTransSentencia(entrada *Resultado) {
	if len(c.TransState) == 0 {
		c.TransState = append(c.TransState, entrada)
	}
}

//Remover sentencia
func (c *Contexto) RemTransSentencia() {
	c.TransState = c.TransState[:len(c.TransState)-1]
}
