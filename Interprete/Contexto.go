package interprete

// Clase Contexto --------------------------------------------------------------------------------------------

type Contexto struct {
	Nombre      string
	Memoria     *Memoria
	ZGlobal     []*Memoria
	Consola     string
	TransState  []*Resultado
	ReturnState []AbstractExpression
	Errores     []string
	Funciones   []*Funcion
	Structs     map[string]*Contexto
	Conversor   *Conversor
}

func NewContexto(nombre string) *Contexto {
	c := &Contexto{
		Nombre:      nombre,
		Memoria:     NewMemoria(nil, "Global"),
		ZGlobal:     make([]*Memoria, 0),
		Consola:     "",
		TransState:  make([]*Resultado, 0, 10),
		ReturnState: make([]AbstractExpression, 0, 10),
		Errores:     make([]string, 0),
		Funciones:   make([]*Funcion, 0),
		Structs:     make(map[string]*Contexto, 0),
		Conversor:   nil,
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

// Se agrega una nueva funcion al vector de funciones
func (c *Contexto) AddFuncion(nombre string, parametros []AbstractExpression, tipo_funcion string, tipo_retorno string, sentencias AbstractExpression, linea int, columna int) bool {
	_, existe := c.ExistFuncion(nombre)
	if existe {
		return false
	}
	c.Funciones = append(c.Funciones, NewFuncion(nombre, "Global", len(parametros), parametros, tipo_funcion, tipo_retorno, sentencias, linea, columna))
	return true
}

// Se verifica si existe una funcion en el vector de funciones
func (c *Contexto) ExistFuncion(nombre string) (*Funcion, bool) {
	for _, funcion := range c.Funciones {
		if funcion.Nombre == nombre {
			return funcion, true
		}
	}
	return nil, false
}

// Se agrega un struct al vector de structs
func (c *Contexto) ExistStruct(nombre string) bool {
	_, existe := c.Structs[nombre]
	return existe
}

//Se agrega una funcion para devolver el struct por el nombre
func (c *Contexto) GetStruct(nombre string) *Contexto {
	return c.Structs[nombre]
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

//se hace una copia del contexto para no afectar el original
func (c *Contexto) Copy() *Contexto {
	var copia = NewContexto(c.Nombre)
	copia.Memoria = c.Memoria.Copy()
	copia.ZGlobal = c.ZGlobal
	copia.Consola = c.Consola
	copia.TransState = c.TransState
	copia.ReturnState = c.ReturnState
	copia.Errores = c.Errores
	copia.Funciones = c.Funciones
	copia.Structs = c.Structs
	copia.Conversor = c.Conversor
	return copia
}

// Se obtiene el valor de una variable en el ambito actual ----------------------------------------------

func (c *Contexto) GetVariable(nombre string) (*Simbolo, bool) {
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
		return nil, false
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

//Agregar sentencia
func (c *Contexto) AddReturnSentencia(entrada AbstractExpression) {
	if len(c.ReturnState) == 0 {
		c.ReturnState = append(c.ReturnState, entrada)
	}
}

//Remover sentencia
func (c *Contexto) RemReturnSentencia() {
	c.ReturnState = c.ReturnState[:len(c.ReturnState)-1]
}

//limpiar variables del contexto
func (c *Contexto) LimpiarVariables() {
	c.Memoria.LimpiarVariables()
}

//mutear funcion
func (c *Contexto) MuteFuncion(nombre string) {
	for _, funcion := range c.Funciones {
		if funcion.Nombre == nombre {
			funcion.Mutating = true
		}
	}
}

//Concatenar consolas de los structs
func (c *Contexto) ConcatenarConsola() {
	for _, struct_ := range c.Structs {
		c.Consola += struct_.Consola
	}
}

//limpiar consolas de los structs
func (c *Contexto) LimpiarConsola() {
	for _, struct_ := range c.Structs {
		struct_.Consola = ""
	}
}
