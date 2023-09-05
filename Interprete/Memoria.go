package interprete

type Memoria struct {
	variables map[string]*Simbolo
	Anterior  *Memoria
}

// Constructor for Memoria----------------------------------------------

func NewMemoria(anterior *Memoria) *Memoria {
	return &Memoria{
		variables: make(map[string]*Simbolo),
		Anterior:  anterior,
	}
}

// Crear simbolo en memoria----------------------------------------------

func (m *Memoria) CreateSimbolo(nombre string, categoria string, tipo TipoE, tipo_compuesto TipoE, num_parametros int, parametros []*Simbolo, resultado *Resultado, tipo_retorno *Tipos, linea int, columna int) bool {
	_, ok := m.variables[nombre]
	if ok {
		return false
	}
	m.variables[nombre] = NewSimbolo(nombre, categoria, tipo, tipo_compuesto, num_parametros, parametros, resultado, tipo_retorno, linea, columna)
	return true
}

// Actualizar simbolo en memoria----------------------------------------------

func (m *Memoria) SetSimbolo(nombre string, resultado *Resultado) bool {
	simbolo, ok := m.variables[nombre]
	//Si es una constante no se puede modificar
	if simbolo.Categoria == "const" {
		return false
	}
	//Si no existe el simbolo
	if !ok {
		return false
	}
	//Controlar los tipos
	if simbolo.Tipo == resultado.Tipo {
		simbolo.Resultado = resultado
	}
	return true
}

// Verificar si existe el simbolo en memoria----------------------------------------------

func (m *Memoria) Exist(nombre string) bool {
	_, ok := m.variables[nombre]
	return ok
}

// Obtener simbolo en memoria----------------------------------------------

func (m *Memoria) GetSimbolo(nombre string) (*Resultado, bool) {
	simbolo, ok := m.variables[nombre]
	if ok {
		return simbolo.Resultado, true
	}
	return NewNil(), false
}
