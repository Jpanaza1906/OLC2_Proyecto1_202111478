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

func (m *Memoria) CreateSimbolo(nombre string, resultado *Resultado, tiposimb string, linea int, columna int) bool {
	_, ok := m.variables[nombre]
	if ok {
		return false
	}
	m.variables[nombre] = NewSimbolo(nombre, resultado, tiposimb, linea, columna)
	return true
}

// Actualizar simbolo en memoria----------------------------------------------

func (m *Memoria) SetSimbolo(nombre string, resultado *Resultado) bool {
	simbolo, ok := m.variables[nombre]
	if simbolo.TipoSimb != "var" {
		return false
	}

	if !ok {
		return false
	}
	//Controlar los tipos
	if simbolo.Resultado.Tipo == resultado.Tipo {
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
