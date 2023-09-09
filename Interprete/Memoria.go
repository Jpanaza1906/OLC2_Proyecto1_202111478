package interprete

type Memoria struct {
	Variables map[string]*Simbolo
	Nambito   string
	Anterior  *Memoria
}

// Constructor for Memoria----------------------------------------------

func NewMemoria(anterior *Memoria, nombre string) *Memoria {
	return &Memoria{
		Variables: make(map[string]*Simbolo),
		Nambito:   nombre,
		Anterior:  anterior,
	}
}

// Crear simbolo en memoria----------------------------------------------

func (m *Memoria) CreateSimbolo(nombre string, categoria string, tipo TipoE, tipo_compuesto TipoE, num_parametros int, parametros []*Simbolo, resultado *Resultado, tipo_retorno *Tipos, linea int, columna int) bool {
	_, ok := m.Variables[nombre]
	if ok {
		return false
	}
	m.Variables[nombre] = NewSimbolo(nombre, m.Nambito, categoria, tipo, tipo_compuesto, num_parametros, parametros, resultado, tipo_retorno, linea, columna)
	return true
}

// Actualizar simbolo en memoria----------------------------------------------

func (m *Memoria) SetSimbolo(nombre string, resultado *Resultado) bool {
	simbolo, ok := m.Variables[nombre]
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
		return true
	}
	return false
}

// Verificar si existe el simbolo en memoria----------------------------------------------

func (m *Memoria) Exist(nombre string) bool {
	_, ok := m.Variables[nombre]
	return ok
}

//Se hace una copia de la memoria
func (m *Memoria) Copy() *Memoria {
	memoria := NewMemoria(nil, m.Nambito)
	for key, value := range m.Variables {
		memoria.Variables[key] = value
	}
	return memoria
}

// Obtener simbolo en memoria----------------------------------------------

func (m *Memoria) GetSimbolo(nombre string) (*Simbolo, bool) {
	simbolo, ok := m.Variables[nombre]
	if ok {
		return simbolo, true
	}
	return nil, false
}

//limpiar variables de la memoria
func (m *Memoria) LimpiarVariables() {
	m.Variables = make(map[string]*Simbolo)
}
