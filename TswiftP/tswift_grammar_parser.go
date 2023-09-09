// Code generated from Tswift_Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftP // Tswift_Grammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Tswift_GrammarParser struct {
	*antlr.BaseParser
}

var Tswift_GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswift_grammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'",
		"'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'",
		"'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'",
		"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'",
		"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'",
		"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'",
		"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'",
		"'repeating'", "'struct'", "'mutating'", "'inout'", "'atoi'", "'iota'",
		"'self'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'",
	}
	staticData.SymbolicNames = []string{
		"", "", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER",
		"DOSPT", "PTCOMA", "INTERROGACION", "PUNTO", "GUIONBAJO", "DIR", "MASIGUAL",
		"MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", "IGUALIGUAL",
		"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "AND", "OR",
		"NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", "SWITCH",
		"CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE",
		"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY",
		"COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", "INOUT", "ATOI",
		"IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO",
		"ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR",
	}
	staticData.RuleNames = []string{
		"s", "l_sentencias", "sentencia", "print_sentencia", "declaracion",
		"constante", "asignacion", "if_sentencia", "switch_sentencia", "l_casos",
		"l_default", "guard_sentencia", "while_sentencia", "for_sentencia",
		"rango_p", "trans_sentencia", "dec_vector", "def_vector", "asig_vector",
		"func_vector", "declaracion_matrices", "tipo_matriz", "def_matriz",
		"listaval_mat", "listaval_mat2", "simple_mat", "asig_mat", "declaracion_metodo",
		"declaracion_funcion", "l_parametros", "llamada_funciones", "l_argumentos",
		"funciones_embebidas", "def_struct", "l_sentencias_struct", "self_data",
		"struct_data", "idstruct", "struct_llamadafunc", "tipo", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 77, 680, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 1,
		1, 5, 1, 86, 8, 1, 10, 1, 12, 1, 89, 9, 1, 1, 2, 1, 2, 3, 2, 93, 8, 2,
		1, 2, 1, 2, 3, 2, 97, 8, 2, 1, 2, 1, 2, 3, 2, 101, 8, 2, 1, 2, 1, 2, 3,
		2, 105, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 114, 8, 2,
		1, 2, 1, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 2, 1, 2, 3,
		2, 126, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 132, 8, 2, 1, 2, 1, 2, 3, 2,
		136, 8, 2, 1, 2, 1, 2, 3, 2, 140, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 145, 8,
		2, 1, 2, 1, 2, 3, 2, 149, 8, 2, 1, 2, 1, 2, 3, 2, 153, 8, 2, 3, 2, 155,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 162, 8, 3, 10, 3, 12, 3, 165,
		9, 3, 1, 3, 1, 3, 3, 3, 169, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		188, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 201, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 212, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 238, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8,
		244, 8, 8, 11, 8, 12, 8, 245, 1, 8, 3, 8, 249, 8, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 258, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 264, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 285, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 299, 8, 15, 3, 15, 301, 8, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 5, 17, 316, 8, 17, 10, 17, 12, 17, 319, 9, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 326, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 349, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 372, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 378, 8, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21,
		391, 8, 21, 1, 22, 1, 22, 3, 22, 395, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 408, 8, 24, 10,
		24, 12, 24, 411, 9, 24, 3, 24, 413, 8, 24, 1, 24, 1, 24, 1, 24, 5, 24,
		418, 8, 24, 10, 24, 12, 24, 421, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 445, 8, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26,
		457, 8, 26, 10, 26, 12, 26, 460, 9, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 5, 27, 469, 8, 27, 10, 27, 12, 27, 472, 9, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 483, 8,
		28, 10, 28, 12, 28, 486, 9, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 29, 3, 29, 497, 8, 29, 1, 29, 1, 29, 1, 29, 3, 29, 502,
		8, 29, 1, 29, 1, 29, 3, 29, 506, 8, 29, 1, 30, 1, 30, 1, 30, 5, 30, 511,
		8, 30, 10, 30, 12, 30, 514, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 3, 31, 520,
		8, 31, 1, 31, 3, 31, 523, 8, 31, 1, 31, 1, 31, 3, 31, 527, 8, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 3, 32, 544, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		5, 33, 550, 8, 33, 10, 33, 12, 33, 553, 9, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 3, 34, 561, 8, 34, 1, 34, 1, 34, 3, 34, 565, 8, 34, 1,
		34, 3, 34, 568, 8, 34, 1, 34, 3, 34, 571, 8, 34, 1, 34, 1, 34, 3, 34, 575,
		8, 34, 3, 34, 577, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 605,
		8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40,
		626, 8, 40, 10, 40, 12, 40, 629, 9, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 652, 8, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 675,
		8, 40, 10, 40, 12, 40, 678, 9, 40, 1, 40, 0, 2, 48, 80, 41, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 0, 11, 2, 0, 12, 12, 74, 74, 1, 0, 32, 33, 2, 0, 61, 61, 64, 64, 2,
		0, 62, 62, 68, 68, 1, 0, 14, 16, 2, 0, 19, 19, 30, 30, 1, 0, 34, 35, 2,
		0, 17, 17, 21, 21, 1, 0, 19, 20, 1, 0, 22, 27, 1, 0, 28, 29, 753, 0, 82,
		1, 0, 0, 0, 2, 87, 1, 0, 0, 0, 4, 154, 1, 0, 0, 0, 6, 156, 1, 0, 0, 0,
		8, 187, 1, 0, 0, 0, 10, 200, 1, 0, 0, 0, 12, 211, 1, 0, 0, 0, 14, 237,
		1, 0, 0, 0, 16, 239, 1, 0, 0, 0, 18, 252, 1, 0, 0, 0, 20, 259, 1, 0, 0,
		0, 22, 265, 1, 0, 0, 0, 24, 273, 1, 0, 0, 0, 26, 279, 1, 0, 0, 0, 28, 290,
		1, 0, 0, 0, 30, 300, 1, 0, 0, 0, 32, 302, 1, 0, 0, 0, 34, 325, 1, 0, 0,
		0, 36, 348, 1, 0, 0, 0, 38, 371, 1, 0, 0, 0, 40, 373, 1, 0, 0, 0, 42, 390,
		1, 0, 0, 0, 44, 394, 1, 0, 0, 0, 46, 396, 1, 0, 0, 0, 48, 412, 1, 0, 0,
		0, 50, 444, 1, 0, 0, 0, 52, 446, 1, 0, 0, 0, 54, 464, 1, 0, 0, 0, 56, 478,
		1, 0, 0, 0, 58, 496, 1, 0, 0, 0, 60, 507, 1, 0, 0, 0, 62, 519, 1, 0, 0,
		0, 64, 543, 1, 0, 0, 0, 66, 545, 1, 0, 0, 0, 68, 576, 1, 0, 0, 0, 70, 578,
		1, 0, 0, 0, 72, 582, 1, 0, 0, 0, 74, 586, 1, 0, 0, 0, 76, 590, 1, 0, 0,
		0, 78, 604, 1, 0, 0, 0, 80, 651, 1, 0, 0, 0, 82, 83, 3, 2, 1, 0, 83, 1,
		1, 0, 0, 0, 84, 86, 3, 4, 2, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0,
		87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 3, 1, 0, 0, 0, 89, 87, 1, 0,
		0, 0, 90, 92, 3, 6, 3, 0, 91, 93, 5, 9, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 155, 1, 0, 0, 0, 94, 96, 3, 8, 4, 0, 95, 97, 5, 9, 0, 0,
		96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 155, 1, 0, 0, 0, 98, 100, 3,
		10, 5, 0, 99, 101, 5, 9, 0, 0, 100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 155, 1, 0, 0, 0, 102, 104, 3, 12, 6, 0, 103, 105, 5, 9, 0, 0, 104,
		103, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 155, 1, 0, 0, 0, 106, 155,
		3, 14, 7, 0, 107, 155, 3, 16, 8, 0, 108, 155, 3, 22, 11, 0, 109, 155, 3,
		24, 12, 0, 110, 155, 3, 26, 13, 0, 111, 113, 3, 30, 15, 0, 112, 114, 5,
		9, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 155, 1, 0, 0,
		0, 115, 117, 3, 32, 16, 0, 116, 118, 5, 9, 0, 0, 117, 116, 1, 0, 0, 0,
		117, 118, 1, 0, 0, 0, 118, 155, 1, 0, 0, 0, 119, 121, 3, 38, 19, 0, 120,
		122, 5, 9, 0, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 155,
		1, 0, 0, 0, 123, 125, 3, 36, 18, 0, 124, 126, 5, 9, 0, 0, 125, 124, 1,
		0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 155, 1, 0, 0, 0, 127, 155, 3, 54, 27,
		0, 128, 155, 3, 56, 28, 0, 129, 131, 3, 60, 30, 0, 130, 132, 5, 9, 0, 0,
		131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 155, 1, 0, 0, 0, 133,
		135, 3, 40, 20, 0, 134, 136, 5, 9, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136,
		1, 0, 0, 0, 136, 155, 1, 0, 0, 0, 137, 139, 3, 52, 26, 0, 138, 140, 5,
		9, 0, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 155, 1, 0, 0,
		0, 141, 155, 3, 66, 33, 0, 142, 144, 3, 70, 35, 0, 143, 145, 5, 9, 0, 0,
		144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 155, 1, 0, 0, 0, 146,
		148, 3, 72, 36, 0, 147, 149, 5, 9, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 155, 1, 0, 0, 0, 150, 152, 3, 76, 38, 0, 151, 153, 5,
		9, 0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 1, 0, 0,
		0, 154, 90, 1, 0, 0, 0, 154, 94, 1, 0, 0, 0, 154, 98, 1, 0, 0, 0, 154,
		102, 1, 0, 0, 0, 154, 106, 1, 0, 0, 0, 154, 107, 1, 0, 0, 0, 154, 108,
		1, 0, 0, 0, 154, 109, 1, 0, 0, 0, 154, 110, 1, 0, 0, 0, 154, 111, 1, 0,
		0, 0, 154, 115, 1, 0, 0, 0, 154, 119, 1, 0, 0, 0, 154, 123, 1, 0, 0, 0,
		154, 127, 1, 0, 0, 0, 154, 128, 1, 0, 0, 0, 154, 129, 1, 0, 0, 0, 154,
		133, 1, 0, 0, 0, 154, 137, 1, 0, 0, 0, 154, 141, 1, 0, 0, 0, 154, 142,
		1, 0, 0, 0, 154, 146, 1, 0, 0, 0, 154, 150, 1, 0, 0, 0, 155, 5, 1, 0, 0,
		0, 156, 157, 5, 31, 0, 0, 157, 158, 5, 3, 0, 0, 158, 163, 3, 80, 40, 0,
		159, 160, 5, 1, 0, 0, 160, 162, 3, 80, 40, 0, 161, 159, 1, 0, 0, 0, 162,
		165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166,
		1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 168, 5, 2, 0, 0, 167, 169, 5, 9,
		0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 7, 1, 0, 0, 0, 170,
		171, 5, 32, 0, 0, 171, 172, 5, 74, 0, 0, 172, 173, 5, 8, 0, 0, 173, 174,
		3, 78, 39, 0, 174, 175, 5, 16, 0, 0, 175, 176, 3, 80, 40, 0, 176, 188,
		1, 0, 0, 0, 177, 178, 5, 32, 0, 0, 178, 179, 5, 74, 0, 0, 179, 180, 5,
		16, 0, 0, 180, 188, 3, 80, 40, 0, 181, 182, 5, 32, 0, 0, 182, 183, 5, 74,
		0, 0, 183, 184, 5, 8, 0, 0, 184, 185, 3, 78, 39, 0, 185, 186, 5, 10, 0,
		0, 186, 188, 1, 0, 0, 0, 187, 170, 1, 0, 0, 0, 187, 177, 1, 0, 0, 0, 187,
		181, 1, 0, 0, 0, 188, 9, 1, 0, 0, 0, 189, 190, 5, 33, 0, 0, 190, 191, 5,
		74, 0, 0, 191, 192, 5, 8, 0, 0, 192, 193, 3, 78, 39, 0, 193, 194, 5, 16,
		0, 0, 194, 195, 3, 80, 40, 0, 195, 201, 1, 0, 0, 0, 196, 197, 5, 33, 0,
		0, 197, 198, 5, 74, 0, 0, 198, 199, 5, 16, 0, 0, 199, 201, 3, 80, 40, 0,
		200, 189, 1, 0, 0, 0, 200, 196, 1, 0, 0, 0, 201, 11, 1, 0, 0, 0, 202, 203,
		5, 74, 0, 0, 203, 204, 5, 14, 0, 0, 204, 212, 3, 80, 40, 0, 205, 206, 5,
		74, 0, 0, 206, 207, 5, 15, 0, 0, 207, 212, 3, 80, 40, 0, 208, 209, 5, 74,
		0, 0, 209, 210, 5, 16, 0, 0, 210, 212, 3, 80, 40, 0, 211, 202, 1, 0, 0,
		0, 211, 205, 1, 0, 0, 0, 211, 208, 1, 0, 0, 0, 212, 13, 1, 0, 0, 0, 213,
		214, 5, 36, 0, 0, 214, 215, 3, 80, 40, 0, 215, 216, 5, 4, 0, 0, 216, 217,
		3, 2, 1, 0, 217, 218, 5, 5, 0, 0, 218, 238, 1, 0, 0, 0, 219, 220, 5, 36,
		0, 0, 220, 221, 3, 80, 40, 0, 221, 222, 5, 4, 0, 0, 222, 223, 3, 2, 1,
		0, 223, 224, 5, 5, 0, 0, 224, 225, 5, 37, 0, 0, 225, 226, 5, 4, 0, 0, 226,
		227, 3, 2, 1, 0, 227, 228, 5, 5, 0, 0, 228, 238, 1, 0, 0, 0, 229, 230,
		5, 36, 0, 0, 230, 231, 3, 80, 40, 0, 231, 232, 5, 4, 0, 0, 232, 233, 3,
		2, 1, 0, 233, 234, 5, 5, 0, 0, 234, 235, 5, 37, 0, 0, 235, 236, 3, 14,
		7, 0, 236, 238, 1, 0, 0, 0, 237, 213, 1, 0, 0, 0, 237, 219, 1, 0, 0, 0,
		237, 229, 1, 0, 0, 0, 238, 15, 1, 0, 0, 0, 239, 240, 5, 38, 0, 0, 240,
		241, 3, 80, 40, 0, 241, 243, 5, 4, 0, 0, 242, 244, 3, 18, 9, 0, 243, 242,
		1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0,
		0, 0, 246, 248, 1, 0, 0, 0, 247, 249, 3, 20, 10, 0, 248, 247, 1, 0, 0,
		0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 5, 0, 0, 251,
		17, 1, 0, 0, 0, 252, 253, 5, 39, 0, 0, 253, 254, 3, 80, 40, 0, 254, 255,
		5, 8, 0, 0, 255, 257, 3, 2, 1, 0, 256, 258, 5, 48, 0, 0, 257, 256, 1, 0,
		0, 0, 257, 258, 1, 0, 0, 0, 258, 19, 1, 0, 0, 0, 259, 260, 5, 40, 0, 0,
		260, 261, 5, 8, 0, 0, 261, 263, 3, 2, 1, 0, 262, 264, 5, 48, 0, 0, 263,
		262, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 21, 1, 0, 0, 0, 265, 266, 5,
		45, 0, 0, 266, 267, 3, 80, 40, 0, 267, 268, 5, 37, 0, 0, 268, 269, 5, 4,
		0, 0, 269, 270, 3, 2, 1, 0, 270, 271, 3, 30, 15, 0, 271, 272, 5, 5, 0,
		0, 272, 23, 1, 0, 0, 0, 273, 274, 5, 41, 0, 0, 274, 275, 3, 80, 40, 0,
		275, 276, 5, 4, 0, 0, 276, 277, 3, 2, 1, 0, 277, 278, 5, 5, 0, 0, 278,
		25, 1, 0, 0, 0, 279, 280, 5, 42, 0, 0, 280, 281, 7, 0, 0, 0, 281, 284,
		5, 43, 0, 0, 282, 285, 3, 28, 14, 0, 283, 285, 3, 80, 40, 0, 284, 282,
		1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 5, 4,
		0, 0, 287, 288, 3, 2, 1, 0, 288, 289, 5, 5, 0, 0, 289, 27, 1, 0, 0, 0,
		290, 291, 3, 80, 40, 0, 291, 292, 5, 44, 0, 0, 292, 293, 3, 80, 40, 0,
		293, 29, 1, 0, 0, 0, 294, 301, 5, 48, 0, 0, 295, 301, 5, 46, 0, 0, 296,
		298, 5, 47, 0, 0, 297, 299, 3, 80, 40, 0, 298, 297, 1, 0, 0, 0, 298, 299,
		1, 0, 0, 0, 299, 301, 1, 0, 0, 0, 300, 294, 1, 0, 0, 0, 300, 295, 1, 0,
		0, 0, 300, 296, 1, 0, 0, 0, 301, 31, 1, 0, 0, 0, 302, 303, 7, 1, 0, 0,
		303, 304, 5, 74, 0, 0, 304, 305, 5, 8, 0, 0, 305, 306, 5, 6, 0, 0, 306,
		307, 3, 78, 39, 0, 307, 308, 5, 7, 0, 0, 308, 309, 5, 16, 0, 0, 309, 310,
		3, 34, 17, 0, 310, 33, 1, 0, 0, 0, 311, 312, 5, 6, 0, 0, 312, 317, 3, 80,
		40, 0, 313, 314, 5, 1, 0, 0, 314, 316, 3, 80, 40, 0, 315, 313, 1, 0, 0,
		0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318,
		320, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321, 5, 7, 0, 0, 321, 326,
		1, 0, 0, 0, 322, 323, 5, 6, 0, 0, 323, 326, 5, 7, 0, 0, 324, 326, 5, 74,
		0, 0, 325, 311, 1, 0, 0, 0, 325, 322, 1, 0, 0, 0, 325, 324, 1, 0, 0, 0,
		326, 35, 1, 0, 0, 0, 327, 328, 5, 74, 0, 0, 328, 329, 5, 6, 0, 0, 329,
		330, 3, 80, 40, 0, 330, 331, 5, 7, 0, 0, 331, 332, 5, 16, 0, 0, 332, 333,
		3, 80, 40, 0, 333, 349, 1, 0, 0, 0, 334, 335, 5, 74, 0, 0, 335, 336, 5,
		6, 0, 0, 336, 337, 3, 80, 40, 0, 337, 338, 5, 7, 0, 0, 338, 339, 5, 14,
		0, 0, 339, 340, 3, 80, 40, 0, 340, 349, 1, 0, 0, 0, 341, 342, 5, 74, 0,
		0, 342, 343, 5, 6, 0, 0, 343, 344, 3, 80, 40, 0, 344, 345, 5, 7, 0, 0,
		345, 346, 5, 15, 0, 0, 346, 347, 3, 80, 40, 0, 347, 349, 1, 0, 0, 0, 348,
		327, 1, 0, 0, 0, 348, 334, 1, 0, 0, 0, 348, 341, 1, 0, 0, 0, 349, 37, 1,
		0, 0, 0, 350, 351, 5, 74, 0, 0, 351, 352, 5, 11, 0, 0, 352, 353, 5, 50,
		0, 0, 353, 354, 5, 3, 0, 0, 354, 355, 3, 80, 40, 0, 355, 356, 5, 2, 0,
		0, 356, 372, 1, 0, 0, 0, 357, 358, 5, 74, 0, 0, 358, 359, 5, 11, 0, 0,
		359, 360, 5, 51, 0, 0, 360, 361, 5, 3, 0, 0, 361, 372, 5, 2, 0, 0, 362,
		363, 5, 74, 0, 0, 363, 364, 5, 11, 0, 0, 364, 365, 5, 52, 0, 0, 365, 366,
		5, 3, 0, 0, 366, 367, 5, 53, 0, 0, 367, 368, 5, 8, 0, 0, 368, 369, 3, 80,
		40, 0, 369, 370, 5, 2, 0, 0, 370, 372, 1, 0, 0, 0, 371, 350, 1, 0, 0, 0,
		371, 357, 1, 0, 0, 0, 371, 362, 1, 0, 0, 0, 372, 39, 1, 0, 0, 0, 373, 374,
		7, 1, 0, 0, 374, 377, 5, 74, 0, 0, 375, 376, 5, 8, 0, 0, 376, 378, 3, 42,
		21, 0, 377, 375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0,
		379, 380, 5, 16, 0, 0, 380, 381, 3, 44, 22, 0, 381, 41, 1, 0, 0, 0, 382,
		383, 5, 6, 0, 0, 383, 384, 3, 42, 21, 0, 384, 385, 5, 7, 0, 0, 385, 391,
		1, 0, 0, 0, 386, 387, 5, 6, 0, 0, 387, 388, 3, 78, 39, 0, 388, 389, 5,
		7, 0, 0, 389, 391, 1, 0, 0, 0, 390, 382, 1, 0, 0, 0, 390, 386, 1, 0, 0,
		0, 391, 43, 1, 0, 0, 0, 392, 395, 3, 46, 23, 0, 393, 395, 3, 50, 25, 0,
		394, 392, 1, 0, 0, 0, 394, 393, 1, 0, 0, 0, 395, 45, 1, 0, 0, 0, 396, 397,
		5, 6, 0, 0, 397, 398, 3, 48, 24, 0, 398, 399, 5, 7, 0, 0, 399, 47, 1, 0,
		0, 0, 400, 401, 6, 24, -1, 0, 401, 413, 3, 46, 23, 0, 402, 403, 3, 80,
		40, 0, 403, 404, 5, 1, 0, 0, 404, 409, 3, 80, 40, 0, 405, 406, 5, 1, 0,
		0, 406, 408, 3, 80, 40, 0, 407, 405, 1, 0, 0, 0, 408, 411, 1, 0, 0, 0,
		409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 413, 1, 0, 0, 0, 411,
		409, 1, 0, 0, 0, 412, 400, 1, 0, 0, 0, 412, 402, 1, 0, 0, 0, 413, 419,
		1, 0, 0, 0, 414, 415, 10, 3, 0, 0, 415, 416, 5, 1, 0, 0, 416, 418, 3, 46,
		23, 0, 417, 414, 1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0,
		419, 420, 1, 0, 0, 0, 420, 49, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 423,
		3, 42, 21, 0, 423, 424, 5, 3, 0, 0, 424, 425, 5, 57, 0, 0, 425, 426, 5,
		8, 0, 0, 426, 427, 3, 50, 25, 0, 427, 428, 5, 1, 0, 0, 428, 429, 5, 55,
		0, 0, 429, 430, 5, 8, 0, 0, 430, 431, 5, 70, 0, 0, 431, 432, 5, 2, 0, 0,
		432, 445, 1, 0, 0, 0, 433, 434, 3, 42, 21, 0, 434, 435, 5, 3, 0, 0, 435,
		436, 5, 57, 0, 0, 436, 437, 5, 8, 0, 0, 437, 438, 3, 80, 40, 0, 438, 439,
		5, 1, 0, 0, 439, 440, 5, 55, 0, 0, 440, 441, 5, 8, 0, 0, 441, 442, 5, 70,
		0, 0, 442, 443, 5, 2, 0, 0, 443, 445, 1, 0, 0, 0, 444, 422, 1, 0, 0, 0,
		444, 433, 1, 0, 0, 0, 445, 51, 1, 0, 0, 0, 446, 447, 5, 74, 0, 0, 447,
		448, 5, 6, 0, 0, 448, 449, 5, 70, 0, 0, 449, 450, 5, 7, 0, 0, 450, 451,
		5, 6, 0, 0, 451, 452, 5, 70, 0, 0, 452, 458, 5, 7, 0, 0, 453, 454, 5, 6,
		0, 0, 454, 455, 5, 70, 0, 0, 455, 457, 5, 7, 0, 0, 456, 453, 1, 0, 0, 0,
		457, 460, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459,
		461, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 461, 462, 5, 16, 0, 0, 462, 463,
		3, 80, 40, 0, 463, 53, 1, 0, 0, 0, 464, 465, 5, 56, 0, 0, 465, 466, 5,
		74, 0, 0, 466, 470, 5, 3, 0, 0, 467, 469, 3, 58, 29, 0, 468, 467, 1, 0,
		0, 0, 469, 472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0,
		471, 473, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 473, 474, 5, 2, 0, 0, 474,
		475, 5, 4, 0, 0, 475, 476, 3, 2, 1, 0, 476, 477, 5, 5, 0, 0, 477, 55, 1,
		0, 0, 0, 478, 479, 5, 56, 0, 0, 479, 480, 5, 74, 0, 0, 480, 484, 5, 3,
		0, 0, 481, 483, 3, 58, 29, 0, 482, 481, 1, 0, 0, 0, 483, 486, 1, 0, 0,
		0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 487, 1, 0, 0, 0, 486,
		484, 1, 0, 0, 0, 487, 488, 5, 2, 0, 0, 488, 489, 5, 19, 0, 0, 489, 490,
		5, 26, 0, 0, 490, 491, 3, 78, 39, 0, 491, 492, 5, 4, 0, 0, 492, 493, 3,
		2, 1, 0, 493, 494, 5, 5, 0, 0, 494, 57, 1, 0, 0, 0, 495, 497, 7, 0, 0,
		0, 496, 495, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498,
		499, 5, 74, 0, 0, 499, 501, 5, 8, 0, 0, 500, 502, 5, 60, 0, 0, 501, 500,
		1, 0, 0, 0, 501, 502, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 505, 3, 78,
		39, 0, 504, 506, 5, 1, 0, 0, 505, 504, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0,
		506, 59, 1, 0, 0, 0, 507, 508, 5, 74, 0, 0, 508, 512, 5, 3, 0, 0, 509,
		511, 3, 62, 31, 0, 510, 509, 1, 0, 0, 0, 511, 514, 1, 0, 0, 0, 512, 510,
		1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 515, 1, 0, 0, 0, 514, 512, 1, 0,
		0, 0, 515, 516, 5, 2, 0, 0, 516, 61, 1, 0, 0, 0, 517, 518, 5, 74, 0, 0,
		518, 520, 5, 8, 0, 0, 519, 517, 1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520,
		522, 1, 0, 0, 0, 521, 523, 5, 13, 0, 0, 522, 521, 1, 0, 0, 0, 522, 523,
		1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 526, 3, 80, 40, 0, 525, 527, 5,
		1, 0, 0, 526, 525, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 63, 1, 0, 0,
		0, 528, 529, 7, 2, 0, 0, 529, 530, 5, 3, 0, 0, 530, 531, 3, 80, 40, 0,
		531, 532, 5, 2, 0, 0, 532, 544, 1, 0, 0, 0, 533, 534, 5, 65, 0, 0, 534,
		535, 5, 3, 0, 0, 535, 536, 3, 80, 40, 0, 536, 537, 5, 2, 0, 0, 537, 544,
		1, 0, 0, 0, 538, 539, 7, 3, 0, 0, 539, 540, 5, 3, 0, 0, 540, 541, 3, 80,
		40, 0, 541, 542, 5, 2, 0, 0, 542, 544, 1, 0, 0, 0, 543, 528, 1, 0, 0, 0,
		543, 533, 1, 0, 0, 0, 543, 538, 1, 0, 0, 0, 544, 65, 1, 0, 0, 0, 545, 546,
		5, 58, 0, 0, 546, 547, 5, 74, 0, 0, 547, 551, 5, 4, 0, 0, 548, 550, 3,
		68, 34, 0, 549, 548, 1, 0, 0, 0, 550, 553, 1, 0, 0, 0, 551, 549, 1, 0,
		0, 0, 551, 552, 1, 0, 0, 0, 552, 554, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0,
		554, 555, 5, 5, 0, 0, 555, 67, 1, 0, 0, 0, 556, 557, 7, 1, 0, 0, 557, 560,
		5, 74, 0, 0, 558, 559, 5, 8, 0, 0, 559, 561, 3, 78, 39, 0, 560, 558, 1,
		0, 0, 0, 560, 561, 1, 0, 0, 0, 561, 564, 1, 0, 0, 0, 562, 563, 5, 16, 0,
		0, 563, 565, 3, 80, 40, 0, 564, 562, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0,
		565, 567, 1, 0, 0, 0, 566, 568, 5, 9, 0, 0, 567, 566, 1, 0, 0, 0, 567,
		568, 1, 0, 0, 0, 568, 577, 1, 0, 0, 0, 569, 571, 5, 59, 0, 0, 570, 569,
		1, 0, 0, 0, 570, 571, 1, 0, 0, 0, 571, 574, 1, 0, 0, 0, 572, 575, 3, 54,
		27, 0, 573, 575, 3, 56, 28, 0, 574, 572, 1, 0, 0, 0, 574, 573, 1, 0, 0,
		0, 575, 577, 1, 0, 0, 0, 576, 556, 1, 0, 0, 0, 576, 570, 1, 0, 0, 0, 577,
		69, 1, 0, 0, 0, 578, 579, 5, 63, 0, 0, 579, 580, 5, 11, 0, 0, 580, 581,
		3, 12, 6, 0, 581, 71, 1, 0, 0, 0, 582, 583, 3, 74, 37, 0, 583, 584, 7,
		4, 0, 0, 584, 585, 3, 80, 40, 0, 585, 73, 1, 0, 0, 0, 586, 587, 3, 80,
		40, 0, 587, 588, 5, 11, 0, 0, 588, 589, 5, 74, 0, 0, 589, 75, 1, 0, 0,
		0, 590, 591, 3, 80, 40, 0, 591, 592, 5, 11, 0, 0, 592, 593, 3, 60, 30,
		0, 593, 77, 1, 0, 0, 0, 594, 605, 5, 64, 0, 0, 595, 605, 5, 65, 0, 0, 596,
		605, 5, 68, 0, 0, 597, 605, 5, 66, 0, 0, 598, 605, 5, 67, 0, 0, 599, 605,
		5, 74, 0, 0, 600, 601, 5, 6, 0, 0, 601, 602, 3, 78, 39, 0, 602, 603, 5,
		7, 0, 0, 603, 605, 1, 0, 0, 0, 604, 594, 1, 0, 0, 0, 604, 595, 1, 0, 0,
		0, 604, 596, 1, 0, 0, 0, 604, 597, 1, 0, 0, 0, 604, 598, 1, 0, 0, 0, 604,
		599, 1, 0, 0, 0, 604, 600, 1, 0, 0, 0, 605, 79, 1, 0, 0, 0, 606, 607, 6,
		40, -1, 0, 607, 608, 7, 5, 0, 0, 608, 652, 3, 80, 40, 23, 609, 652, 7,
		6, 0, 0, 610, 611, 5, 63, 0, 0, 611, 612, 5, 11, 0, 0, 612, 652, 5, 74,
		0, 0, 613, 652, 3, 64, 32, 0, 614, 652, 3, 60, 30, 0, 615, 616, 5, 74,
		0, 0, 616, 617, 5, 6, 0, 0, 617, 618, 5, 70, 0, 0, 618, 619, 5, 7, 0, 0,
		619, 620, 5, 6, 0, 0, 620, 621, 5, 70, 0, 0, 621, 627, 5, 7, 0, 0, 622,
		623, 5, 6, 0, 0, 623, 624, 5, 70, 0, 0, 624, 626, 5, 7, 0, 0, 625, 622,
		1, 0, 0, 0, 626, 629, 1, 0, 0, 0, 627, 625, 1, 0, 0, 0, 627, 628, 1, 0,
		0, 0, 628, 652, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 630, 631, 5, 74, 0, 0,
		631, 632, 5, 6, 0, 0, 632, 633, 3, 80, 40, 0, 633, 634, 5, 7, 0, 0, 634,
		652, 1, 0, 0, 0, 635, 636, 5, 74, 0, 0, 636, 637, 5, 11, 0, 0, 637, 652,
		5, 54, 0, 0, 638, 639, 5, 74, 0, 0, 639, 640, 5, 11, 0, 0, 640, 652, 5,
		55, 0, 0, 641, 652, 5, 74, 0, 0, 642, 652, 5, 49, 0, 0, 643, 652, 5, 71,
		0, 0, 644, 652, 5, 70, 0, 0, 645, 652, 5, 73, 0, 0, 646, 652, 5, 72, 0,
		0, 647, 648, 5, 3, 0, 0, 648, 649, 3, 80, 40, 0, 649, 650, 5, 2, 0, 0,
		650, 652, 1, 0, 0, 0, 651, 606, 1, 0, 0, 0, 651, 609, 1, 0, 0, 0, 651,
		610, 1, 0, 0, 0, 651, 613, 1, 0, 0, 0, 651, 614, 1, 0, 0, 0, 651, 615,
		1, 0, 0, 0, 651, 630, 1, 0, 0, 0, 651, 635, 1, 0, 0, 0, 651, 638, 1, 0,
		0, 0, 651, 641, 1, 0, 0, 0, 651, 642, 1, 0, 0, 0, 651, 643, 1, 0, 0, 0,
		651, 644, 1, 0, 0, 0, 651, 645, 1, 0, 0, 0, 651, 646, 1, 0, 0, 0, 651,
		647, 1, 0, 0, 0, 652, 676, 1, 0, 0, 0, 653, 654, 10, 22, 0, 0, 654, 655,
		7, 7, 0, 0, 655, 675, 3, 80, 40, 23, 656, 657, 10, 21, 0, 0, 657, 658,
		7, 8, 0, 0, 658, 675, 3, 80, 40, 22, 659, 660, 10, 20, 0, 0, 660, 661,
		5, 18, 0, 0, 661, 675, 3, 80, 40, 21, 662, 663, 10, 19, 0, 0, 663, 664,
		7, 9, 0, 0, 664, 675, 3, 80, 40, 20, 665, 666, 10, 18, 0, 0, 666, 667,
		7, 10, 0, 0, 667, 675, 3, 80, 40, 19, 668, 669, 10, 15, 0, 0, 669, 670,
		5, 11, 0, 0, 670, 675, 5, 74, 0, 0, 671, 672, 10, 12, 0, 0, 672, 673, 5,
		11, 0, 0, 673, 675, 3, 60, 30, 0, 674, 653, 1, 0, 0, 0, 674, 656, 1, 0,
		0, 0, 674, 659, 1, 0, 0, 0, 674, 662, 1, 0, 0, 0, 674, 665, 1, 0, 0, 0,
		674, 668, 1, 0, 0, 0, 674, 671, 1, 0, 0, 0, 675, 678, 1, 0, 0, 0, 676,
		674, 1, 0, 0, 0, 676, 677, 1, 0, 0, 0, 677, 81, 1, 0, 0, 0, 678, 676, 1,
		0, 0, 0, 63, 87, 92, 96, 100, 104, 113, 117, 121, 125, 131, 135, 139, 144,
		148, 152, 154, 163, 168, 187, 200, 211, 237, 245, 248, 257, 263, 284, 298,
		300, 317, 325, 348, 371, 377, 390, 394, 409, 412, 419, 444, 458, 470, 484,
		496, 501, 505, 512, 519, 522, 526, 543, 551, 560, 564, 567, 570, 574, 576,
		604, 627, 651, 674, 676,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// Tswift_GrammarParserInit initializes any static state used to implement Tswift_GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTswift_GrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tswift_GrammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.once.Do(tswift_grammarParserInit)
}

// NewTswift_GrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTswift_GrammarParser(input antlr.TokenStream) *Tswift_GrammarParser {
	Tswift_GrammarParserInit()
	this := new(Tswift_GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Tswift_GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tswift_Grammar.g4"

	return this
}

// Tswift_GrammarParser tokens.
const (
	Tswift_GrammarParserEOF           = antlr.TokenEOF
	Tswift_GrammarParserT__0          = 1
	Tswift_GrammarParserPARDER        = 2
	Tswift_GrammarParserPARIZQ        = 3
	Tswift_GrammarParserLLAVEIZQ      = 4
	Tswift_GrammarParserLLAVEDER      = 5
	Tswift_GrammarParserCORCHETEIZQ   = 6
	Tswift_GrammarParserCORCHETEDER   = 7
	Tswift_GrammarParserDOSPT         = 8
	Tswift_GrammarParserPTCOMA        = 9
	Tswift_GrammarParserINTERROGACION = 10
	Tswift_GrammarParserPUNTO         = 11
	Tswift_GrammarParserGUIONBAJO     = 12
	Tswift_GrammarParserDIR           = 13
	Tswift_GrammarParserMASIGUAL      = 14
	Tswift_GrammarParserMENOSIGUAL    = 15
	Tswift_GrammarParserIGUAL         = 16
	Tswift_GrammarParserDIV           = 17
	Tswift_GrammarParserMOD           = 18
	Tswift_GrammarParserMENOS         = 19
	Tswift_GrammarParserMAS           = 20
	Tswift_GrammarParserPOR           = 21
	Tswift_GrammarParserIGUALIGUAL    = 22
	Tswift_GrammarParserDIFERENTE     = 23
	Tswift_GrammarParserMAYORIGUAL    = 24
	Tswift_GrammarParserMENORIGUAL    = 25
	Tswift_GrammarParserMAYOR         = 26
	Tswift_GrammarParserMENOR         = 27
	Tswift_GrammarParserAND           = 28
	Tswift_GrammarParserOR            = 29
	Tswift_GrammarParserNOT           = 30
	Tswift_GrammarParserPRINT         = 31
	Tswift_GrammarParserVAR           = 32
	Tswift_GrammarParserLET           = 33
	Tswift_GrammarParserTRUE          = 34
	Tswift_GrammarParserFALSE         = 35
	Tswift_GrammarParserIF            = 36
	Tswift_GrammarParserELSE          = 37
	Tswift_GrammarParserSWITCH        = 38
	Tswift_GrammarParserCASE          = 39
	Tswift_GrammarParserDEFAULT       = 40
	Tswift_GrammarParserWHILE         = 41
	Tswift_GrammarParserFOR           = 42
	Tswift_GrammarParserIN            = 43
	Tswift_GrammarParserRANGO         = 44
	Tswift_GrammarParserGUARD         = 45
	Tswift_GrammarParserCONTINUE      = 46
	Tswift_GrammarParserRETURN        = 47
	Tswift_GrammarParserBREAK         = 48
	Tswift_GrammarParserNIL           = 49
	Tswift_GrammarParserAPPEND        = 50
	Tswift_GrammarParserREMOVELAST    = 51
	Tswift_GrammarParserREMOVE        = 52
	Tswift_GrammarParserAT            = 53
	Tswift_GrammarParserISEMPTY       = 54
	Tswift_GrammarParserCOUNT         = 55
	Tswift_GrammarParserFUNC          = 56
	Tswift_GrammarParserREPEATING     = 57
	Tswift_GrammarParserSTRUCT        = 58
	Tswift_GrammarParserMUTATING      = 59
	Tswift_GrammarParserINOUT         = 60
	Tswift_GrammarParserATOI          = 61
	Tswift_GrammarParserIOTA          = 62
	Tswift_GrammarParserSELF          = 63
	Tswift_GrammarParserINT           = 64
	Tswift_GrammarParserFLOAT         = 65
	Tswift_GrammarParserBOOL          = 66
	Tswift_GrammarParserCHARACTER     = 67
	Tswift_GrammarParserSTRING        = 68
	Tswift_GrammarParserENBLANCO      = 69
	Tswift_GrammarParserENTERO        = 70
	Tswift_GrammarParserDECIMAL       = 71
	Tswift_GrammarParserCARACTER      = 72
	Tswift_GrammarParserCADENA        = 73
	Tswift_GrammarParserID            = 74
	Tswift_GrammarParserUL_C          = 75
	Tswift_GrammarParserML_C          = 76
	Tswift_GrammarParserERROR         = 77
)

// Tswift_GrammarParser rules.
const (
	Tswift_GrammarParserRULE_s                    = 0
	Tswift_GrammarParserRULE_l_sentencias         = 1
	Tswift_GrammarParserRULE_sentencia            = 2
	Tswift_GrammarParserRULE_print_sentencia      = 3
	Tswift_GrammarParserRULE_declaracion          = 4
	Tswift_GrammarParserRULE_constante            = 5
	Tswift_GrammarParserRULE_asignacion           = 6
	Tswift_GrammarParserRULE_if_sentencia         = 7
	Tswift_GrammarParserRULE_switch_sentencia     = 8
	Tswift_GrammarParserRULE_l_casos              = 9
	Tswift_GrammarParserRULE_l_default            = 10
	Tswift_GrammarParserRULE_guard_sentencia      = 11
	Tswift_GrammarParserRULE_while_sentencia      = 12
	Tswift_GrammarParserRULE_for_sentencia        = 13
	Tswift_GrammarParserRULE_rango_p              = 14
	Tswift_GrammarParserRULE_trans_sentencia      = 15
	Tswift_GrammarParserRULE_dec_vector           = 16
	Tswift_GrammarParserRULE_def_vector           = 17
	Tswift_GrammarParserRULE_asig_vector          = 18
	Tswift_GrammarParserRULE_func_vector          = 19
	Tswift_GrammarParserRULE_declaracion_matrices = 20
	Tswift_GrammarParserRULE_tipo_matriz          = 21
	Tswift_GrammarParserRULE_def_matriz           = 22
	Tswift_GrammarParserRULE_listaval_mat         = 23
	Tswift_GrammarParserRULE_listaval_mat2        = 24
	Tswift_GrammarParserRULE_simple_mat           = 25
	Tswift_GrammarParserRULE_asig_mat             = 26
	Tswift_GrammarParserRULE_declaracion_metodo   = 27
	Tswift_GrammarParserRULE_declaracion_funcion  = 28
	Tswift_GrammarParserRULE_l_parametros         = 29
	Tswift_GrammarParserRULE_llamada_funciones    = 30
	Tswift_GrammarParserRULE_l_argumentos         = 31
	Tswift_GrammarParserRULE_funciones_embebidas  = 32
	Tswift_GrammarParserRULE_def_struct           = 33
	Tswift_GrammarParserRULE_l_sentencias_struct  = 34
	Tswift_GrammarParserRULE_self_data            = 35
	Tswift_GrammarParserRULE_struct_data          = 36
	Tswift_GrammarParserRULE_idstruct             = 37
	Tswift_GrammarParserRULE_struct_llamadafunc   = 38
	Tswift_GrammarParserRULE_tipo                 = 39
	Tswift_GrammarParserRULE_e                    = 40
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) CopyAll(ctx *SContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SLSentenciasContext struct {
	SContext
}

func NewSLSentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SLSentenciasContext {
	var p = new(SLSentenciasContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *SLSentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SLSentenciasContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *SLSentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSLSentencias(s)
	}
}

func (s *SLSentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSLSentencias(s)
	}
}

func (s *SLSentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSLSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tswift_GrammarParserRULE_s)
	localctx = NewSLSentenciasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.L_sentencias()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_sentenciasContext is an interface to support dynamic dispatch.
type IL_sentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_sentenciasContext differentiates from other interfaces.
	IsL_sentenciasContext()
}

type L_sentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_sentenciasContext() *L_sentenciasContext {
	var p = new(L_sentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias
	return p
}

func InitEmptyL_sentenciasContext(p *L_sentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias
}

func (*L_sentenciasContext) IsL_sentenciasContext() {}

func NewL_sentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_sentenciasContext {
	var p = new(L_sentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias

	return p
}

func (s *L_sentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *L_sentenciasContext) CopyAll(ctx *L_sentenciasContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_sentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_sentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_SentenciaContext struct {
	L_sentenciasContext
}

func NewL_SentenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_SentenciaContext {
	var p = new(L_SentenciaContext)

	InitEmptyL_sentenciasContext(&p.L_sentenciasContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentenciasContext))

	return p
}

func (s *L_SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_SentenciaContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *L_SentenciaContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISentenciaContext)
}

func (s *L_SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Sentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_sentencias() (localctx IL_sentenciasContext) {
	localctx = NewL_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Tswift_GrammarParserRULE_l_sentencias)
	var _alt int

	localctx = NewL_SentenciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.Sentencia()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) CopyAll(ctx *SentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type S_ForContext struct {
	SentenciaContext
}

func NewS_ForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ForContext {
	var p = new(S_ForContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ForContext) For_sentencia() IFor_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_sentenciaContext)
}

func (s *S_ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_For(s)
	}
}

func (s *S_ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_For(s)
	}
}

func (s *S_ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_For(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_FuncionContext struct {
	SentenciaContext
}

func NewS_Declaracion_FuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_FuncionContext {
	var p = new(S_Declaracion_FuncionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_FuncionContext) Declaracion_funcion() IDeclaracion_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_funcionContext)
}

func (s *S_Declaracion_FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion_Funcion(s)
	}
}

func (s *S_Declaracion_FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion_Funcion(s)
	}
}

func (s *S_Declaracion_FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion_Funcion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_MatrizContext struct {
	SentenciaContext
}

func NewS_Declaracion_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_MatrizContext {
	var p = new(S_Declaracion_MatrizContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_MatrizContext) Declaracion_matrices() IDeclaracion_matricesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_matricesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_matricesContext)
}

func (s *S_Declaracion_MatrizContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Declaracion_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion_Matriz(s)
	}
}

func (s *S_Declaracion_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion_Matriz(s)
	}
}

func (s *S_Declaracion_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Def_StructContext struct {
	SentenciaContext
}

func NewS_Def_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Def_StructContext {
	var p = new(S_Def_StructContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Def_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Def_StructContext) Def_struct() IDef_structContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_structContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_structContext)
}

func (s *S_Def_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Def_Struct(s)
	}
}

func (s *S_Def_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Def_Struct(s)
	}
}

func (s *S_Def_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Def_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Asignacion_VectorContext struct {
	SentenciaContext
}

func NewS_Asignacion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Asignacion_VectorContext {
	var p = new(S_Asignacion_VectorContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Asignacion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Asignacion_VectorContext) Asig_vector() IAsig_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsig_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsig_vectorContext)
}

func (s *S_Asignacion_VectorContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Asignacion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Asignacion_Vector(s)
	}
}

func (s *S_Asignacion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Asignacion_Vector(s)
	}
}

func (s *S_Asignacion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Asignacion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_MetodoContext struct {
	SentenciaContext
}

func NewS_Declaracion_MetodoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_MetodoContext {
	var p = new(S_Declaracion_MetodoContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_MetodoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_MetodoContext) Declaracion_metodo() IDeclaracion_metodoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_metodoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_metodoContext)
}

func (s *S_Declaracion_MetodoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion_Metodo(s)
	}
}

func (s *S_Declaracion_MetodoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion_Metodo(s)
	}
}

func (s *S_Declaracion_MetodoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion_Metodo(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Struct_DataContext struct {
	SentenciaContext
}

func NewS_Struct_DataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Struct_DataContext {
	var p = new(S_Struct_DataContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Struct_DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Struct_DataContext) Struct_data() IStruct_dataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_dataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_dataContext)
}

func (s *S_Struct_DataContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Struct_DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Struct_Data(s)
	}
}

func (s *S_Struct_DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Struct_Data(s)
	}
}

func (s *S_Struct_DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Struct_Data(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_AsignacionContext struct {
	SentenciaContext
}

func NewS_AsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_AsignacionContext {
	var p = new(S_AsignacionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_AsignacionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *S_AsignacionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Asignacion(s)
	}
}

func (s *S_AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Asignacion(s)
	}
}

func (s *S_AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Asignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Self_DataContext struct {
	SentenciaContext
}

func NewS_Self_DataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Self_DataContext {
	var p = new(S_Self_DataContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Self_DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Self_DataContext) Self_data() ISelf_dataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelf_dataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelf_dataContext)
}

func (s *S_Self_DataContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Self_DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Self_Data(s)
	}
}

func (s *S_Self_DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Self_Data(s)
	}
}

func (s *S_Self_DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Self_Data(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Funcion_VectorContext struct {
	SentenciaContext
}

func NewS_Funcion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Funcion_VectorContext {
	var p = new(S_Funcion_VectorContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Funcion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Funcion_VectorContext) Func_vector() IFunc_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_vectorContext)
}

func (s *S_Funcion_VectorContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Funcion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Funcion_Vector(s)
	}
}

func (s *S_Funcion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Funcion_Vector(s)
	}
}

func (s *S_Funcion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Funcion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_VectorContext struct {
	SentenciaContext
}

func NewS_Declaracion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_VectorContext {
	var p = new(S_Declaracion_VectorContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_VectorContext) Dec_vector() IDec_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_vectorContext)
}

func (s *S_Declaracion_VectorContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Declaracion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion_Vector(s)
	}
}

func (s *S_Declaracion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion_Vector(s)
	}
}

func (s *S_Declaracion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_IfContext struct {
	SentenciaContext
}

func NewS_IfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_IfContext {
	var p = new(S_IfContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_IfContext) If_sentencia() IIf_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_sentenciaContext)
}

func (s *S_IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_If(s)
	}
}

func (s *S_IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_If(s)
	}
}

func (s *S_IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_If(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_TransicionContext struct {
	SentenciaContext
}

func NewS_TransicionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_TransicionContext {
	var p = new(S_TransicionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_TransicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_TransicionContext) Trans_sentencia() ITrans_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrans_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrans_sentenciaContext)
}

func (s *S_TransicionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_TransicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Transicion(s)
	}
}

func (s *S_TransicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Transicion(s)
	}
}

func (s *S_TransicionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Transicion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Asignacion_MatrizContext struct {
	SentenciaContext
}

func NewS_Asignacion_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Asignacion_MatrizContext {
	var p = new(S_Asignacion_MatrizContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Asignacion_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Asignacion_MatrizContext) Asig_mat() IAsig_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsig_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsig_matContext)
}

func (s *S_Asignacion_MatrizContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Asignacion_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Asignacion_Matriz(s)
	}
}

func (s *S_Asignacion_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Asignacion_Matriz(s)
	}
}

func (s *S_Asignacion_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Asignacion_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_SwitchContext struct {
	SentenciaContext
}

func NewS_SwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_SwitchContext {
	var p = new(S_SwitchContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_SwitchContext) Switch_sentencia() ISwitch_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_sentenciaContext)
}

func (s *S_SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Switch(s)
	}
}

func (s *S_SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Switch(s)
	}
}

func (s *S_SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Switch(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Llamada_FuncionContext struct {
	SentenciaContext
}

func NewS_Llamada_FuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Llamada_FuncionContext {
	var p = new(S_Llamada_FuncionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Llamada_FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Llamada_FuncionContext) Llamada_funciones() ILlamada_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionesContext)
}

func (s *S_Llamada_FuncionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Llamada_FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Llamada_Funcion(s)
	}
}

func (s *S_Llamada_FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Llamada_Funcion(s)
	}
}

func (s *S_Llamada_FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Llamada_Funcion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Struct_Llamada_FuncContext struct {
	SentenciaContext
}

func NewS_Struct_Llamada_FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Struct_Llamada_FuncContext {
	var p = new(S_Struct_Llamada_FuncContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Struct_Llamada_FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Struct_Llamada_FuncContext) Struct_llamadafunc() IStruct_llamadafuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_llamadafuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_llamadafuncContext)
}

func (s *S_Struct_Llamada_FuncContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_Struct_Llamada_FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Struct_Llamada_Func(s)
	}
}

func (s *S_Struct_Llamada_FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Struct_Llamada_Func(s)
	}
}

func (s *S_Struct_Llamada_FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Struct_Llamada_Func(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_ConstanteContext struct {
	SentenciaContext
}

func NewS_ConstanteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ConstanteContext {
	var p = new(S_ConstanteContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ConstanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ConstanteContext) Constante() IConstanteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstanteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstanteContext)
}

func (s *S_ConstanteContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_ConstanteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Constante(s)
	}
}

func (s *S_ConstanteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Constante(s)
	}
}

func (s *S_ConstanteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Constante(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_GuardContext struct {
	SentenciaContext
}

func NewS_GuardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_GuardContext {
	var p = new(S_GuardContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_GuardContext) Guard_sentencia() IGuard_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_sentenciaContext)
}

func (s *S_GuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Guard(s)
	}
}

func (s *S_GuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Guard(s)
	}
}

func (s *S_GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Guard(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_WhileContext struct {
	SentenciaContext
}

func NewS_WhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_WhileContext {
	var p = new(S_WhileContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_WhileContext) While_sentencia() IWhile_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_sentenciaContext)
}

func (s *S_WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_While(s)
	}
}

func (s *S_WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_While(s)
	}
}

func (s *S_WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_While(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_PrintContext struct {
	SentenciaContext
}

func NewS_PrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_PrintContext {
	var p = new(S_PrintContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_PrintContext) Print_sentencia() IPrint_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_sentenciaContext)
}

func (s *S_PrintContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Print(s)
	}
}

func (s *S_PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Print(s)
	}
}

func (s *S_PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Print(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_DeclaracionContext struct {
	SentenciaContext
}

func NewS_DeclaracionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_DeclaracionContext {
	var p = new(S_DeclaracionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_DeclaracionContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *S_DeclaracionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *S_DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitS_Declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Tswift_GrammarParserRULE_sentencia)
	var _la int

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewS_PrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Print_sentencia()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(91)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewS_DeclaracionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Declaracion()
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(95)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewS_ConstanteContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Constante()
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(99)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		localctx = NewS_AsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Asignacion()
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(103)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		localctx = NewS_IfContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.If_sentencia()
		}

	case 6:
		localctx = NewS_SwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Switch_sentencia()
		}

	case 7:
		localctx = NewS_GuardContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(108)
			p.Guard_sentencia()
		}

	case 8:
		localctx = NewS_WhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(109)
			p.While_sentencia()
		}

	case 9:
		localctx = NewS_ForContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(110)
			p.For_sentencia()
		}

	case 10:
		localctx = NewS_TransicionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)
			p.Trans_sentencia()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(112)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 11:
		localctx = NewS_Declaracion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(115)
			p.Dec_vector()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(116)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		localctx = NewS_Funcion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(119)
			p.Func_vector()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(120)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 13:
		localctx = NewS_Asignacion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(123)
			p.Asig_vector()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(124)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 14:
		localctx = NewS_Declaracion_MetodoContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(127)
			p.Declaracion_metodo()
		}

	case 15:
		localctx = NewS_Declaracion_FuncionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(128)
			p.Declaracion_funcion()
		}

	case 16:
		localctx = NewS_Llamada_FuncionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(129)
			p.Llamada_funciones()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(130)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 17:
		localctx = NewS_Declaracion_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(133)
			p.Declaracion_matrices()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(134)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 18:
		localctx = NewS_Asignacion_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(137)
			p.Asig_mat()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(138)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 19:
		localctx = NewS_Def_StructContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(141)
			p.Def_struct()
		}

	case 20:
		localctx = NewS_Self_DataContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(142)
			p.Self_data()
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(143)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 21:
		localctx = NewS_Struct_DataContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(146)
			p.Struct_data()
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(147)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 22:
		localctx = NewS_Struct_Llamada_FuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(150)
			p.Struct_llamadafunc()
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(151)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrint_sentenciaContext is an interface to support dynamic dispatch.
type IPrint_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrint_sentenciaContext differentiates from other interfaces.
	IsPrint_sentenciaContext()
}

type Print_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_sentenciaContext() *Print_sentenciaContext {
	var p = new(Print_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_print_sentencia
	return p
}

func InitEmptyPrint_sentenciaContext(p *Print_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_print_sentencia
}

func (*Print_sentenciaContext) IsPrint_sentenciaContext() {}

func NewPrint_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_sentenciaContext {
	var p = new(Print_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_print_sentencia

	return p
}

func (s *Print_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_sentenciaContext) CopyAll(ctx *Print_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Print_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	Print_sentenciaContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	InitEmptyPrint_sentenciaContext(&p.Print_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Print_sentenciaContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPRINT, 0)
}

func (s *PrintContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *PrintContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *PrintContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *PrintContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Print_sentencia() (localctx IPrint_sentenciaContext) {
	localctx = NewPrint_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Tswift_GrammarParserRULE_print_sentencia)
	var _la int

	localctx = NewPrintContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(Tswift_GrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.e(0)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(159)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.e(0)
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(167)
			p.Match(Tswift_GrammarParserPTCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) CopyAll(ctx *DeclaracionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_Tipo_ValContext struct {
	DeclaracionContext
}

func NewDeclaracion_Tipo_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_ValContext {
	var p = new(Declaracion_Tipo_ValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_Tipo_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_Tipo_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_Tipo_ValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_Tipo_ValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_Tipo_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Declaracion_Tipo_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Tipo_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_Tipo_noValContext struct {
	DeclaracionContext
}

func NewDeclaracion_Tipo_noValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_noValContext {
	var p = new(Declaracion_Tipo_noValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_Tipo_noValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_noValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_Tipo_noValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_Tipo_noValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_Tipo_noValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_noValContext) INTERROGACION() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINTERROGACION, 0)
}

func (s *Declaracion_Tipo_noValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Tipo_noVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_ValContext struct {
	DeclaracionContext
}

func NewDeclaracion_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_ValContext {
	var p = new(Declaracion_ValContext)

	InitEmptyDeclaracionContext(&p.DeclaracionContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclaracionContext))

	return p
}

func (s *Declaracion_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Declaracion_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Tswift_GrammarParserRULE_declaracion)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Tipo()
		}
		{
			p.SetState(174)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.e(0)
		}

	case 2:
		localctx = NewDeclaracion_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.e(0)
		}

	case 3:
		localctx = NewDeclaracion_Tipo_noValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Tipo()
		}
		{
			p.SetState(185)
			p.Match(Tswift_GrammarParserINTERROGACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstanteContext is an interface to support dynamic dispatch.
type IConstanteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstanteContext differentiates from other interfaces.
	IsConstanteContext()
}

type ConstanteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstanteContext() *ConstanteContext {
	var p = new(ConstanteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_constante
	return p
}

func InitEmptyConstanteContext(p *ConstanteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_constante
}

func (*ConstanteContext) IsConstanteContext() {}

func NewConstanteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstanteContext {
	var p = new(ConstanteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_constante

	return p
}

func (s *ConstanteContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstanteContext) CopyAll(ctx *ConstanteContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConstanteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstanteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Constante_ValContext struct {
	ConstanteContext
}

func NewConstante_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Constante_ValContext {
	var p = new(Constante_ValContext)

	InitEmptyConstanteContext(&p.ConstanteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstanteContext))

	return p
}

func (s *Constante_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constante_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Constante_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Constante_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Constante_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Constante_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterConstante_Val(s)
	}
}

func (s *Constante_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitConstante_Val(s)
	}
}

func (s *Constante_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitConstante_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Constante_Tipo_ValContext struct {
	ConstanteContext
}

func NewConstante_Tipo_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Constante_Tipo_ValContext {
	var p = new(Constante_Tipo_ValContext)

	InitEmptyConstanteContext(&p.ConstanteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConstanteContext))

	return p
}

func (s *Constante_Tipo_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constante_Tipo_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Constante_Tipo_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Constante_Tipo_ValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Constante_Tipo_ValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Constante_Tipo_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Constante_Tipo_ValContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Constante_Tipo_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterConstante_Tipo_Val(s)
	}
}

func (s *Constante_Tipo_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitConstante_Tipo_Val(s)
	}
}

func (s *Constante_Tipo_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitConstante_Tipo_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Constante() (localctx IConstanteContext) {
	localctx = NewConstanteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Tswift_GrammarParserRULE_constante)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstante_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(Tswift_GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Tipo()
		}
		{
			p.SetState(193)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.e(0)
		}

	case 2:
		localctx = NewConstante_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Match(Tswift_GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) CopyAll(ctx *AsignacionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ResAsgContext struct {
	AsignacionContext
}

func NewResAsgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ResAsgContext {
	var p = new(ResAsgContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *ResAsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResAsgContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *ResAsgContext) MENOSIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOSIGUAL, 0)
}

func (s *ResAsgContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ResAsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterResAsg(s)
	}
}

func (s *ResAsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitResAsg(s)
	}
}

func (s *ResAsgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitResAsg(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumAsgContext struct {
	AsignacionContext
}

func NewSumAsgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumAsgContext {
	var p = new(SumAsgContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *SumAsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumAsgContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *SumAsgContext) MASIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMASIGUAL, 0)
}

func (s *SumAsgContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumAsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSumAsg(s)
	}
}

func (s *SumAsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSumAsg(s)
	}
}

func (s *SumAsgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSumAsg(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsigContext struct {
	AsignacionContext
}

func NewAsigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsigContext {
	var p = new(AsigContext)

	InitEmptyAsignacionContext(&p.AsignacionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignacionContext))

	return p
}

func (s *AsigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsigContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *AsigContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *AsigContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *AsigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterAsig(s)
	}
}

func (s *AsigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitAsig(s)
	}
}

func (s *AsigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitAsig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Tswift_GrammarParserRULE_asignacion)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSumAsgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(Tswift_GrammarParserMASIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.e(0)
		}

	case 2:
		localctx = NewResAsgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(Tswift_GrammarParserMENOSIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.e(0)
		}

	case 3:
		localctx = NewAsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_sentenciaContext is an interface to support dynamic dispatch.
type IIf_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_sentenciaContext differentiates from other interfaces.
	IsIf_sentenciaContext()
}

type If_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_sentenciaContext() *If_sentenciaContext {
	var p = new(If_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_if_sentencia
	return p
}

func InitEmptyIf_sentenciaContext(p *If_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_if_sentencia
}

func (*If_sentenciaContext) IsIf_sentenciaContext() {}

func NewIf_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_sentenciaContext {
	var p = new(If_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_if_sentencia

	return p
}

func (s *If_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *If_sentenciaContext) CopyAll(ctx *If_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type If_ElseIfContext struct {
	If_sentenciaContext
}

func NewIf_ElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_ElseIfContext {
	var p = new(If_ElseIfContext)

	InitEmptyIf_sentenciaContext(&p.If_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_sentenciaContext))

	return p
}

func (s *If_ElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_ElseIfContext) IF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIF, 0)
}

func (s *If_ElseIfContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *If_ElseIfContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *If_ElseIfContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *If_ElseIfContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *If_ElseIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserELSE, 0)
}

func (s *If_ElseIfContext) If_sentencia() IIf_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_sentenciaContext)
}

func (s *If_ElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterIf_ElseIf(s)
	}
}

func (s *If_ElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitIf_ElseIf(s)
	}
}

func (s *If_ElseIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitIf_ElseIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_ElseContext struct {
	If_sentenciaContext
}

func NewIf_ElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_ElseContext {
	var p = new(If_ElseContext)

	InitEmptyIf_sentenciaContext(&p.If_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_sentenciaContext))

	return p
}

func (s *If_ElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_ElseContext) IF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIF, 0)
}

func (s *If_ElseContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *If_ElseContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserLLAVEIZQ)
}

func (s *If_ElseContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, i)
}

func (s *If_ElseContext) AllL_sentencias() []IL_sentenciasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			len++
		}
	}

	tst := make([]IL_sentenciasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_sentenciasContext); ok {
			tst[i] = t.(IL_sentenciasContext)
			i++
		}
	}

	return tst
}

func (s *If_ElseContext) L_sentencias(i int) IL_sentenciasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *If_ElseContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserLLAVEDER)
}

func (s *If_ElseContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, i)
}

func (s *If_ElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserELSE, 0)
}

func (s *If_ElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterIf_Else(s)
	}
}

func (s *If_ElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitIf_Else(s)
	}
}

func (s *If_ElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitIf_Else(s)

	default:
		return t.VisitChildren(s)
	}
}

type If_SimpleContext struct {
	If_sentenciaContext
}

func NewIf_SimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *If_SimpleContext {
	var p = new(If_SimpleContext)

	InitEmptyIf_sentenciaContext(&p.If_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_sentenciaContext))

	return p
}

func (s *If_SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_SimpleContext) IF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIF, 0)
}

func (s *If_SimpleContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *If_SimpleContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *If_SimpleContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *If_SimpleContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *If_SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterIf_Simple(s)
	}
}

func (s *If_SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitIf_Simple(s)
	}
}

func (s *If_SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitIf_Simple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) If_sentencia() (localctx IIf_sentenciaContext) {
	localctx = NewIf_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Tswift_GrammarParserRULE_if_sentencia)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.e(0)
		}
		{
			p.SetState(215)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.L_sentencias()
		}
		{
			p.SetState(217)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIf_ElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.e(0)
		}
		{
			p.SetState(221)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.L_sentencias()
		}
		{
			p.SetState(223)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(Tswift_GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.L_sentencias()
		}
		{
			p.SetState(227)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIf_ElseIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.e(0)
		}
		{
			p.SetState(231)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.L_sentencias()
		}
		{
			p.SetState(233)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(Tswift_GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.If_sentencia()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_sentenciaContext is an interface to support dynamic dispatch.
type ISwitch_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_sentenciaContext differentiates from other interfaces.
	IsSwitch_sentenciaContext()
}

type Switch_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_sentenciaContext() *Switch_sentenciaContext {
	var p = new(Switch_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_switch_sentencia
	return p
}

func InitEmptySwitch_sentenciaContext(p *Switch_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_switch_sentencia
}

func (*Switch_sentenciaContext) IsSwitch_sentenciaContext() {}

func NewSwitch_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_sentenciaContext {
	var p = new(Switch_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_switch_sentencia

	return p
}

func (s *Switch_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_sentenciaContext) CopyAll(ctx *Switch_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchContext struct {
	Switch_sentenciaContext
}

func NewSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchContext {
	var p = new(SwitchContext)

	InitEmptySwitch_sentenciaContext(&p.Switch_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_sentenciaContext))

	return p
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSWITCH, 0)
}

func (s *SwitchContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SwitchContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *SwitchContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *SwitchContext) AllL_casos() []IL_casosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_casosContext); ok {
			len++
		}
	}

	tst := make([]IL_casosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_casosContext); ok {
			tst[i] = t.(IL_casosContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) L_casos(i int) IL_casosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_casosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_casosContext)
}

func (s *SwitchContext) L_default() IL_defaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_defaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_defaultContext)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSwitch(s)
	}
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Switch_sentencia() (localctx ISwitch_sentenciaContext) {
	localctx = NewSwitch_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Tswift_GrammarParserRULE_switch_sentencia)
	var _la int

	localctx = NewSwitchContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(Tswift_GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.e(0)
	}
	{
		p.SetState(241)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Tswift_GrammarParserCASE {
		{
			p.SetState(242)
			p.L_casos()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDEFAULT {
		{
			p.SetState(247)
			p.L_default()
		}

	}
	{
		p.SetState(250)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_casosContext is an interface to support dynamic dispatch.
type IL_casosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_casosContext differentiates from other interfaces.
	IsL_casosContext()
}

type L_casosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_casosContext() *L_casosContext {
	var p = new(L_casosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_casos
	return p
}

func InitEmptyL_casosContext(p *L_casosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_casos
}

func (*L_casosContext) IsL_casosContext() {}

func NewL_casosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_casosContext {
	var p = new(L_casosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_casos

	return p
}

func (s *L_casosContext) GetParser() antlr.Parser { return s.parser }

func (s *L_casosContext) CopyAll(ctx *L_casosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_casosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_casosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CaseContext struct {
	L_casosContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	InitEmptyL_casosContext(&p.L_casosContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_casosContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCASE, 0)
}

func (s *CaseContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *CaseContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *CaseContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *CaseContext) BREAK() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserBREAK, 0)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_casos() (localctx IL_casosContext) {
	localctx = NewL_casosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Tswift_GrammarParserRULE_l_casos)
	var _la int

	localctx = NewCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(Tswift_GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.e(0)
	}
	{
		p.SetState(254)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.L_sentencias()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserBREAK {
		{
			p.SetState(256)
			p.Match(Tswift_GrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_defaultContext is an interface to support dynamic dispatch.
type IL_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_defaultContext differentiates from other interfaces.
	IsL_defaultContext()
}

type L_defaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_defaultContext() *L_defaultContext {
	var p = new(L_defaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_default
	return p
}

func InitEmptyL_defaultContext(p *L_defaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_default
}

func (*L_defaultContext) IsL_defaultContext() {}

func NewL_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_defaultContext {
	var p = new(L_defaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_default

	return p
}

func (s *L_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *L_defaultContext) CopyAll(ctx *L_defaultContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultContext struct {
	L_defaultContext
}

func NewDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultContext {
	var p = new(DefaultContext)

	InitEmptyL_defaultContext(&p.L_defaultContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_defaultContext))

	return p
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDEFAULT, 0)
}

func (s *DefaultContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *DefaultContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *DefaultContext) BREAK() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserBREAK, 0)
}

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_default() (localctx IL_defaultContext) {
	localctx = NewL_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Tswift_GrammarParserRULE_l_default)
	var _la int

	localctx = NewDefaultContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(Tswift_GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.L_sentencias()
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserBREAK {
		{
			p.SetState(262)
			p.Match(Tswift_GrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuard_sentenciaContext is an interface to support dynamic dispatch.
type IGuard_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGuard_sentenciaContext differentiates from other interfaces.
	IsGuard_sentenciaContext()
}

type Guard_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_sentenciaContext() *Guard_sentenciaContext {
	var p = new(Guard_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_guard_sentencia
	return p
}

func InitEmptyGuard_sentenciaContext(p *Guard_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_guard_sentencia
}

func (*Guard_sentenciaContext) IsGuard_sentenciaContext() {}

func NewGuard_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_sentenciaContext {
	var p = new(Guard_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_guard_sentencia

	return p
}

func (s *Guard_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_sentenciaContext) CopyAll(ctx *Guard_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Guard_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GuardContext struct {
	Guard_sentenciaContext
}

func NewGuardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardContext {
	var p = new(GuardContext)

	InitEmptyGuard_sentenciaContext(&p.Guard_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Guard_sentenciaContext))

	return p
}

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) GUARD() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserGUARD, 0)
}

func (s *GuardContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *GuardContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserELSE, 0)
}

func (s *GuardContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *GuardContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *GuardContext) Trans_sentencia() ITrans_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrans_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrans_sentenciaContext)
}

func (s *GuardContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *GuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterGuard(s)
	}
}

func (s *GuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitGuard(s)
	}
}

func (s *GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Guard_sentencia() (localctx IGuard_sentenciaContext) {
	localctx = NewGuard_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Tswift_GrammarParserRULE_guard_sentencia)
	localctx = NewGuardContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(Tswift_GrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.e(0)
	}
	{
		p.SetState(267)
		p.Match(Tswift_GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.L_sentencias()
	}
	{
		p.SetState(270)
		p.Trans_sentencia()
	}
	{
		p.SetState(271)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_sentenciaContext is an interface to support dynamic dispatch.
type IWhile_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_sentenciaContext differentiates from other interfaces.
	IsWhile_sentenciaContext()
}

type While_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_sentenciaContext() *While_sentenciaContext {
	var p = new(While_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_while_sentencia
	return p
}

func InitEmptyWhile_sentenciaContext(p *While_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_while_sentencia
}

func (*While_sentenciaContext) IsWhile_sentenciaContext() {}

func NewWhile_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_sentenciaContext {
	var p = new(While_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_while_sentencia

	return p
}

func (s *While_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *While_sentenciaContext) CopyAll(ctx *While_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileContext struct {
	While_sentenciaContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	InitEmptyWhile_sentenciaContext(&p.While_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_sentenciaContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserWHILE, 0)
}

func (s *WhileContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *WhileContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *WhileContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *WhileContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) While_sentencia() (localctx IWhile_sentenciaContext) {
	localctx = NewWhile_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Tswift_GrammarParserRULE_while_sentencia)
	localctx = NewWhileContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(Tswift_GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.e(0)
	}
	{
		p.SetState(275)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.L_sentencias()
	}
	{
		p.SetState(277)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_sentenciaContext is an interface to support dynamic dispatch.
type IFor_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_sentenciaContext differentiates from other interfaces.
	IsFor_sentenciaContext()
}

type For_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_sentenciaContext() *For_sentenciaContext {
	var p = new(For_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_for_sentencia
	return p
}

func InitEmptyFor_sentenciaContext(p *For_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_for_sentencia
}

func (*For_sentenciaContext) IsFor_sentenciaContext() {}

func NewFor_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_sentenciaContext {
	var p = new(For_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_for_sentencia

	return p
}

func (s *For_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *For_sentenciaContext) CopyAll(ctx *For_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForContext struct {
	For_sentenciaContext
	id antlr.Token
}

func NewForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForContext {
	var p = new(ForContext)

	InitEmptyFor_sentenciaContext(&p.For_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_sentenciaContext))

	return p
}

func (s *ForContext) GetId() antlr.Token { return s.id }

func (s *ForContext) SetId(v antlr.Token) { s.id = v }

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFOR, 0)
}

func (s *ForContext) IN() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIN, 0)
}

func (s *ForContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *ForContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *ForContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *ForContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *ForContext) GUIONBAJO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserGUIONBAJO, 0)
}

func (s *ForContext) Rango_p() IRango_pContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRango_pContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRango_pContext)
}

func (s *ForContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFor(s)
	}
}

func (s *ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) For_sentencia() (localctx IFor_sentenciaContext) {
	localctx = NewFor_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Tswift_GrammarParserRULE_for_sentencia)
	var _la int

	localctx = NewForContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(Tswift_GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ForContext).id = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ForContext).id = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(281)
		p.Match(Tswift_GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(282)
			p.Rango_p()
		}

	case 2:
		{
			p.SetState(283)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(286)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.L_sentencias()
	}
	{
		p.SetState(288)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRango_pContext is an interface to support dynamic dispatch.
type IRango_pContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRango_pContext differentiates from other interfaces.
	IsRango_pContext()
}

type Rango_pContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRango_pContext() *Rango_pContext {
	var p = new(Rango_pContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_rango_p
	return p
}

func InitEmptyRango_pContext(p *Rango_pContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_rango_p
}

func (*Rango_pContext) IsRango_pContext() {}

func NewRango_pContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rango_pContext {
	var p = new(Rango_pContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_rango_p

	return p
}

func (s *Rango_pContext) GetParser() antlr.Parser { return s.parser }

func (s *Rango_pContext) CopyAll(ctx *Rango_pContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Rango_pContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rango_pContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RangoContext struct {
	Rango_pContext
}

func NewRangoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangoContext {
	var p = new(RangoContext)

	InitEmptyRango_pContext(&p.Rango_pContext)
	p.parser = parser
	p.CopyAll(ctx.(*Rango_pContext))

	return p
}

func (s *RangoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *RangoContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *RangoContext) RANGO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserRANGO, 0)
}

func (s *RangoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterRango(s)
	}
}

func (s *RangoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitRango(s)
	}
}

func (s *RangoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitRango(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Rango_p() (localctx IRango_pContext) {
	localctx = NewRango_pContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Tswift_GrammarParserRULE_rango_p)
	localctx = NewRangoContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.e(0)
	}
	{
		p.SetState(291)
		p.Match(Tswift_GrammarParserRANGO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.e(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITrans_sentenciaContext is an interface to support dynamic dispatch.
type ITrans_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTrans_sentenciaContext differentiates from other interfaces.
	IsTrans_sentenciaContext()
}

type Trans_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrans_sentenciaContext() *Trans_sentenciaContext {
	var p = new(Trans_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_trans_sentencia
	return p
}

func InitEmptyTrans_sentenciaContext(p *Trans_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_trans_sentencia
}

func (*Trans_sentenciaContext) IsTrans_sentenciaContext() {}

func NewTrans_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Trans_sentenciaContext {
	var p = new(Trans_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_trans_sentencia

	return p
}

func (s *Trans_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Trans_sentenciaContext) CopyAll(ctx *Trans_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Trans_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Trans_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnContext struct {
	Trans_sentenciaContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserRETURN, 0)
}

func (s *ReturnContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakContext struct {
	Trans_sentenciaContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserBREAK, 0)
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitBreak(s)
	}
}

func (s *BreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueContext struct {
	Trans_sentenciaContext
}

func NewContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueContext {
	var p = new(ContinueContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCONTINUE, 0)
}

func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitContinue(s)
	}
}

func (s *ContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Trans_sentencia() (localctx ITrans_sentenciaContext) {
	localctx = NewTrans_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Tswift_GrammarParserRULE_trans_sentencia)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserBREAK:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(Tswift_GrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCONTINUE:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(Tswift_GrammarParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserRETURN:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(Tswift_GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(297)
				p.e(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDec_vectorContext is an interface to support dynamic dispatch.
type IDec_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDec_vectorContext differentiates from other interfaces.
	IsDec_vectorContext()
}

type Dec_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_vectorContext() *Dec_vectorContext {
	var p = new(Dec_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_dec_vector
	return p
}

func InitEmptyDec_vectorContext(p *Dec_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_dec_vector
}

func (*Dec_vectorContext) IsDec_vectorContext() {}

func NewDec_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_vectorContext {
	var p = new(Dec_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_dec_vector

	return p
}

func (s *Dec_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_vectorContext) CopyAll(ctx *Dec_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Dec_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_VectorContext struct {
	Dec_vectorContext
	tipod antlr.Token
}

func NewDeclaracion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_VectorContext {
	var p = new(Declaracion_VectorContext)

	InitEmptyDec_vectorContext(&p.Dec_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_vectorContext))

	return p
}

func (s *Declaracion_VectorContext) GetTipod() antlr.Token { return s.tipod }

func (s *Declaracion_VectorContext) SetTipod(v antlr.Token) { s.tipod = v }

func (s *Declaracion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_VectorContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Declaracion_VectorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Declaracion_VectorContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_VectorContext) Def_vector() IDef_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_vectorContext)
}

func (s *Declaracion_VectorContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_VectorContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Declaracion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Vector(s)
	}
}

func (s *Declaracion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Vector(s)
	}
}

func (s *Declaracion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Dec_vector() (localctx IDec_vectorContext) {
	localctx = NewDec_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Tswift_GrammarParserRULE_dec_vector)
	var _la int

	localctx = NewDeclaracion_VectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Declaracion_VectorContext).tipod = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Declaracion_VectorContext).tipod = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(303)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Tipo()
	}
	{
		p.SetState(307)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Def_vector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDef_vectorContext is an interface to support dynamic dispatch.
type IDef_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDef_vectorContext differentiates from other interfaces.
	IsDef_vectorContext()
}

type Def_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_vectorContext() *Def_vectorContext {
	var p = new(Def_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_vector
	return p
}

func InitEmptyDef_vectorContext(p *Def_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_vector
}

func (*Def_vectorContext) IsDef_vectorContext() {}

func NewDef_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_vectorContext {
	var p = new(Def_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_def_vector

	return p
}

func (s *Def_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_vectorContext) CopyAll(ctx *Def_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Def_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_Vector_VacioContext struct {
	Def_vectorContext
}

func NewDef_Vector_VacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Vector_VacioContext {
	var p = new(Def_Vector_VacioContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_Vector_VacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Vector_VacioContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Def_Vector_VacioContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Def_Vector_VacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Vector_Vacio(s)
	}
}

func (s *Def_Vector_VacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Vector_Vacio(s)
	}
}

func (s *Def_Vector_VacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Vector_Vacio(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_VectorContext struct {
	Def_vectorContext
}

func NewDef_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_VectorContext {
	var p = new(Def_VectorContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Def_VectorContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Def_VectorContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Def_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Def_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Vector(s)
	}
}

func (s *Def_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Vector(s)
	}
}

func (s *Def_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Vector_IdContext struct {
	Def_vectorContext
}

func NewDef_Vector_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Vector_IdContext {
	var p = new(Def_Vector_IdContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_Vector_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Vector_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Def_Vector_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Vector_Id(s)
	}
}

func (s *Def_Vector_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Vector_Id(s)
	}
}

func (s *Def_Vector_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Vector_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Def_vector() (localctx IDef_vectorContext) {
	localctx = NewDef_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Tswift_GrammarParserRULE_def_vector)
	var _la int

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.e(0)
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarParserT__0 {
			{
				p.SetState(313)
				p.Match(Tswift_GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(314)
				p.e(0)
			}

			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(320)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDef_Vector_VacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDef_Vector_IdContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsig_vectorContext is an interface to support dynamic dispatch.
type IAsig_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsig_vectorContext differentiates from other interfaces.
	IsAsig_vectorContext()
}

type Asig_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsig_vectorContext() *Asig_vectorContext {
	var p = new(Asig_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asig_vector
	return p
}

func InitEmptyAsig_vectorContext(p *Asig_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asig_vector
}

func (*Asig_vectorContext) IsAsig_vectorContext() {}

func NewAsig_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asig_vectorContext {
	var p = new(Asig_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_asig_vector

	return p
}

func (s *Asig_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Asig_vectorContext) CopyAll(ctx *Asig_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asig_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asig_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Asig_VectorContext struct {
	Asig_vectorContext
}

func NewAsig_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asig_VectorContext {
	var p = new(Asig_VectorContext)

	InitEmptyAsig_vectorContext(&p.Asig_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asig_vectorContext))

	return p
}

func (s *Asig_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asig_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Asig_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Asig_VectorContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Asig_VectorContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Asig_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Asig_VectorContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Asig_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterAsig_Vector(s)
	}
}

func (s *Asig_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitAsig_Vector(s)
	}
}

func (s *Asig_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitAsig_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type ResAsg_VectorContext struct {
	Asig_vectorContext
}

func NewResAsg_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ResAsg_VectorContext {
	var p = new(ResAsg_VectorContext)

	InitEmptyAsig_vectorContext(&p.Asig_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asig_vectorContext))

	return p
}

func (s *ResAsg_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResAsg_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *ResAsg_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *ResAsg_VectorContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *ResAsg_VectorContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ResAsg_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *ResAsg_VectorContext) MENOSIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOSIGUAL, 0)
}

func (s *ResAsg_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterResAsg_Vector(s)
	}
}

func (s *ResAsg_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitResAsg_Vector(s)
	}
}

func (s *ResAsg_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitResAsg_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumAsg_VectorContext struct {
	Asig_vectorContext
}

func NewSumAsg_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumAsg_VectorContext {
	var p = new(SumAsg_VectorContext)

	InitEmptyAsig_vectorContext(&p.Asig_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asig_vectorContext))

	return p
}

func (s *SumAsg_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumAsg_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *SumAsg_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *SumAsg_VectorContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *SumAsg_VectorContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumAsg_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *SumAsg_VectorContext) MASIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMASIGUAL, 0)
}

func (s *SumAsg_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSumAsg_Vector(s)
	}
}

func (s *SumAsg_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSumAsg_Vector(s)
	}
}

func (s *SumAsg_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSumAsg_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Asig_vector() (localctx IAsig_vectorContext) {
	localctx = NewAsig_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Tswift_GrammarParserRULE_asig_vector)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsig_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.e(0)
		}
		{
			p.SetState(330)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.e(0)
		}

	case 2:
		localctx = NewSumAsg_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.e(0)
		}
		{
			p.SetState(337)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(Tswift_GrammarParserMASIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.e(0)
		}

	case 3:
		localctx = NewResAsg_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.e(0)
		}
		{
			p.SetState(344)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(Tswift_GrammarParserMENOSIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_vectorContext is an interface to support dynamic dispatch.
type IFunc_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_vectorContext differentiates from other interfaces.
	IsFunc_vectorContext()
}

type Func_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_vectorContext() *Func_vectorContext {
	var p = new(Func_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_func_vector
	return p
}

func InitEmptyFunc_vectorContext(p *Func_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_func_vector
}

func (*Func_vectorContext) IsFunc_vectorContext() {}

func NewFunc_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_vectorContext {
	var p = new(Func_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_func_vector

	return p
}

func (s *Func_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_vectorContext) CopyAll(ctx *Func_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Func_Vector_AppendContext struct {
	Func_vectorContext
}

func NewFunc_Vector_AppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_AppendContext {
	var p = new(Func_Vector_AppendContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_AppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_AppendContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Func_Vector_AppendContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Func_Vector_AppendContext) APPEND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserAPPEND, 0)
}

func (s *Func_Vector_AppendContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_Vector_AppendContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Func_Vector_AppendContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_Vector_AppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Vector_Append(s)
	}
}

func (s *Func_Vector_AppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Vector_Append(s)
	}
}

func (s *Func_Vector_AppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Vector_Append(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_RemoveContext struct {
	Func_vectorContext
}

func NewFunc_Vector_RemoveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_RemoveContext {
	var p = new(Func_Vector_RemoveContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_RemoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_RemoveContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Func_Vector_RemoveContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Func_Vector_RemoveContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserREMOVE, 0)
}

func (s *Func_Vector_RemoveContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_Vector_RemoveContext) AT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserAT, 0)
}

func (s *Func_Vector_RemoveContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Func_Vector_RemoveContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Func_Vector_RemoveContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_Vector_RemoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Vector_Remove(s)
	}
}

func (s *Func_Vector_RemoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Vector_Remove(s)
	}
}

func (s *Func_Vector_RemoveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Vector_Remove(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_RemoveLastContext struct {
	Func_vectorContext
}

func NewFunc_Vector_RemoveLastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_RemoveLastContext {
	var p = new(Func_Vector_RemoveLastContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_RemoveLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_RemoveLastContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Func_Vector_RemoveLastContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Func_Vector_RemoveLastContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserREMOVELAST, 0)
}

func (s *Func_Vector_RemoveLastContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_Vector_RemoveLastContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_Vector_RemoveLastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Vector_RemoveLast(s)
	}
}

func (s *Func_Vector_RemoveLastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Vector_RemoveLast(s)
	}
}

func (s *Func_Vector_RemoveLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Vector_RemoveLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Func_vector() (localctx IFunc_vectorContext) {
	localctx = NewFunc_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Tswift_GrammarParserRULE_func_vector)
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunc_Vector_AppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(Tswift_GrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.e(0)
		}
		{
			p.SetState(355)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunc_Vector_RemoveLastContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Match(Tswift_GrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunc_Vector_RemoveContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(362)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(Tswift_GrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(Tswift_GrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.e(0)
		}
		{
			p.SetState(369)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_matricesContext is an interface to support dynamic dispatch.
type IDeclaracion_matricesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_matricesContext differentiates from other interfaces.
	IsDeclaracion_matricesContext()
}

type Declaracion_matricesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_matricesContext() *Declaracion_matricesContext {
	var p = new(Declaracion_matricesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_matrices
	return p
}

func InitEmptyDeclaracion_matricesContext(p *Declaracion_matricesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_matrices
}

func (*Declaracion_matricesContext) IsDeclaracion_matricesContext() {}

func NewDeclaracion_matricesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_matricesContext {
	var p = new(Declaracion_matricesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_matrices

	return p
}

func (s *Declaracion_matricesContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_matricesContext) CopyAll(ctx *Declaracion_matricesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_matricesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_matricesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_MatrizContext struct {
	Declaracion_matricesContext
}

func NewDeclaracion_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_MatrizContext {
	var p = new(Declaracion_MatrizContext)

	InitEmptyDeclaracion_matricesContext(&p.Declaracion_matricesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_matricesContext))

	return p
}

func (s *Declaracion_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_MatrizContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_MatrizContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Declaracion_MatrizContext) Def_matriz() IDef_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_matrizContext)
}

func (s *Declaracion_MatrizContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Declaracion_MatrizContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Declaracion_MatrizContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Declaracion_MatrizContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Declaracion_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Matriz(s)
	}
}

func (s *Declaracion_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Matriz(s)
	}
}

func (s *Declaracion_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Declaracion_matrices() (localctx IDeclaracion_matricesContext) {
	localctx = NewDeclaracion_matricesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Tswift_GrammarParserRULE_declaracion_matrices)
	var _la int

	localctx = NewDeclaracion_MatrizContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(374)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDOSPT {
		{
			p.SetState(375)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.Tipo_matriz()
		}

	}
	{
		p.SetState(379)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)
		p.Def_matriz()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipo_matrizContext is an interface to support dynamic dispatch.
type ITipo_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipo_matrizContext differentiates from other interfaces.
	IsTipo_matrizContext()
}

type Tipo_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipo_matrizContext() *Tipo_matrizContext {
	var p = new(Tipo_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo_matriz
	return p
}

func InitEmptyTipo_matrizContext(p *Tipo_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo_matriz
}

func (*Tipo_matrizContext) IsTipo_matrizContext() {}

func NewTipo_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_matrizContext {
	var p = new(Tipo_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_tipo_matriz

	return p
}

func (s *Tipo_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_matrizContext) CopyAll(ctx *Tipo_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Tipo_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Tipo_MatrizContext struct {
	Tipo_matrizContext
}

func NewTipo_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_MatrizContext {
	var p = new(Tipo_MatrizContext)

	InitEmptyTipo_matrizContext(&p.Tipo_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipo_matrizContext))

	return p
}

func (s *Tipo_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_MatrizContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Tipo_MatrizContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Tipo_MatrizContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Tipo_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Matriz(s)
	}
}

func (s *Tipo_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Matriz(s)
	}
}

func (s *Tipo_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_Matriz_SimpleContext struct {
	Tipo_matrizContext
}

func NewTipo_Matriz_SimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_Matriz_SimpleContext {
	var p = new(Tipo_Matriz_SimpleContext)

	InitEmptyTipo_matrizContext(&p.Tipo_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tipo_matrizContext))

	return p
}

func (s *Tipo_Matriz_SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_Matriz_SimpleContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Tipo_Matriz_SimpleContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Tipo_Matriz_SimpleContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Tipo_Matriz_SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Matriz_Simple(s)
	}
}

func (s *Tipo_Matriz_SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Matriz_Simple(s)
	}
}

func (s *Tipo_Matriz_SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Matriz_Simple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Tipo_matriz() (localctx ITipo_matrizContext) {
	localctx = NewTipo_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Tswift_GrammarParserRULE_tipo_matriz)
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTipo_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Tipo_matriz()
		}
		{
			p.SetState(384)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTipo_Matriz_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Tipo()
		}
		{
			p.SetState(388)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDef_matrizContext is an interface to support dynamic dispatch.
type IDef_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDef_matrizContext differentiates from other interfaces.
	IsDef_matrizContext()
}

type Def_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_matrizContext() *Def_matrizContext {
	var p = new(Def_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_matriz
	return p
}

func InitEmptyDef_matrizContext(p *Def_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_matriz
}

func (*Def_matrizContext) IsDef_matrizContext() {}

func NewDef_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_matrizContext {
	var p = new(Def_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_def_matriz

	return p
}

func (s *Def_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_matrizContext) CopyAll(ctx *Def_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Def_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_MatrizContext struct {
	Def_matrizContext
}

func NewDef_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_MatrizContext {
	var p = new(Def_MatrizContext)

	InitEmptyDef_matrizContext(&p.Def_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_matrizContext))

	return p
}

func (s *Def_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_MatrizContext) Listaval_mat() IListaval_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaval_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaval_matContext)
}

func (s *Def_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Matriz(s)
	}
}

func (s *Def_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Matriz(s)
	}
}

func (s *Def_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Matriz_SimpleContext struct {
	Def_matrizContext
}

func NewDef_Matriz_SimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_SimpleContext {
	var p = new(Def_Matriz_SimpleContext)

	InitEmptyDef_matrizContext(&p.Def_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_matrizContext))

	return p
}

func (s *Def_Matriz_SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_SimpleContext) Simple_mat() ISimple_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matContext)
}

func (s *Def_Matriz_SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Matriz_Simple(s)
	}
}

func (s *Def_Matriz_SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Matriz_Simple(s)
	}
}

func (s *Def_Matriz_SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Matriz_Simple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Def_matriz() (localctx IDef_matrizContext) {
	localctx = NewDef_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Tswift_GrammarParserRULE_def_matriz)
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Listaval_mat()
		}

	case 2:
		localctx = NewDef_Matriz_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Simple_mat()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaval_matContext is an interface to support dynamic dispatch.
type IListaval_matContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListaval_matContext differentiates from other interfaces.
	IsListaval_matContext()
}

type Listaval_matContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaval_matContext() *Listaval_matContext {
	var p = new(Listaval_matContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat
	return p
}

func InitEmptyListaval_matContext(p *Listaval_matContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat
}

func (*Listaval_matContext) IsListaval_matContext() {}

func NewListaval_matContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listaval_matContext {
	var p = new(Listaval_matContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat

	return p
}

func (s *Listaval_matContext) GetParser() antlr.Parser { return s.parser }

func (s *Listaval_matContext) CopyAll(ctx *Listaval_matContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Listaval_matContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listaval_matContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaCompletaValContext struct {
	Listaval_matContext
}

func NewListaCompletaValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaCompletaValContext {
	var p = new(ListaCompletaValContext)

	InitEmptyListaval_matContext(&p.Listaval_matContext)
	p.parser = parser
	p.CopyAll(ctx.(*Listaval_matContext))

	return p
}

func (s *ListaCompletaValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaCompletaValContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *ListaCompletaValContext) Listaval_mat2() IListaval_mat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaval_mat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaval_mat2Context)
}

func (s *ListaCompletaValContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *ListaCompletaValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterListaCompletaVal(s)
	}
}

func (s *ListaCompletaValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitListaCompletaVal(s)
	}
}

func (s *ListaCompletaValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitListaCompletaVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Listaval_mat() (localctx IListaval_matContext) {
	localctx = NewListaval_matContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Tswift_GrammarParserRULE_listaval_mat)
	localctx = NewListaCompletaValContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.listaval_mat2(0)
	}
	{
		p.SetState(398)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaval_mat2Context is an interface to support dynamic dispatch.
type IListaval_mat2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListaval_mat2Context differentiates from other interfaces.
	IsListaval_mat2Context()
}

type Listaval_mat2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaval_mat2Context() *Listaval_mat2Context {
	var p = new(Listaval_mat2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat2
	return p
}

func InitEmptyListaval_mat2Context(p *Listaval_mat2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat2
}

func (*Listaval_mat2Context) IsListaval_mat2Context() {}

func NewListaval_mat2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listaval_mat2Context {
	var p = new(Listaval_mat2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_listaval_mat2

	return p
}

func (s *Listaval_mat2Context) GetParser() antlr.Parser { return s.parser }

func (s *Listaval_mat2Context) CopyAll(ctx *Listaval_mat2Context) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Listaval_mat2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listaval_mat2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaValorSigContext struct {
	Listaval_mat2Context
}

func NewListaValorSigContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaValorSigContext {
	var p = new(ListaValorSigContext)

	InitEmptyListaval_mat2Context(&p.Listaval_mat2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Listaval_mat2Context))

	return p
}

func (s *ListaValorSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaValorSigContext) Listaval_mat() IListaval_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaval_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaval_matContext)
}

func (s *ListaValorSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterListaValorSig(s)
	}
}

func (s *ListaValorSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitListaValorSig(s)
	}
}

func (s *ListaValorSigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitListaValorSig(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListaValoresHermanosContext struct {
	Listaval_mat2Context
}

func NewListaValoresHermanosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaValoresHermanosContext {
	var p = new(ListaValoresHermanosContext)

	InitEmptyListaval_mat2Context(&p.Listaval_mat2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Listaval_mat2Context))

	return p
}

func (s *ListaValoresHermanosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaValoresHermanosContext) Listaval_mat2() IListaval_mat2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaval_mat2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaval_mat2Context)
}

func (s *ListaValoresHermanosContext) Listaval_mat() IListaval_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaval_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaval_matContext)
}

func (s *ListaValoresHermanosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterListaValoresHermanos(s)
	}
}

func (s *ListaValoresHermanosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitListaValoresHermanos(s)
	}
}

func (s *ListaValoresHermanosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitListaValoresHermanos(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListaExprContext struct {
	Listaval_mat2Context
}

func NewListaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaExprContext {
	var p = new(ListaExprContext)

	InitEmptyListaval_mat2Context(&p.Listaval_mat2Context)
	p.parser = parser
	p.CopyAll(ctx.(*Listaval_mat2Context))

	return p
}

func (s *ListaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaExprContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *ListaExprContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *ListaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterListaExpr(s)
	}
}

func (s *ListaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitListaExpr(s)
	}
}

func (s *ListaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitListaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Listaval_mat2() (localctx IListaval_mat2Context) {
	return p.listaval_mat2(0)
}

func (p *Tswift_GrammarParser) listaval_mat2(_p int) (localctx IListaval_mat2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListaval_mat2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaval_mat2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, Tswift_GrammarParserRULE_listaval_mat2, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserCORCHETEIZQ:
		localctx = NewListaValorSigContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(401)
			p.Listaval_mat()
		}

	case Tswift_GrammarParserPARIZQ, Tswift_GrammarParserMENOS, Tswift_GrammarParserNOT, Tswift_GrammarParserTRUE, Tswift_GrammarParserFALSE, Tswift_GrammarParserNIL, Tswift_GrammarParserATOI, Tswift_GrammarParserIOTA, Tswift_GrammarParserSELF, Tswift_GrammarParserINT, Tswift_GrammarParserFLOAT, Tswift_GrammarParserSTRING, Tswift_GrammarParserENTERO, Tswift_GrammarParserDECIMAL, Tswift_GrammarParserCARACTER, Tswift_GrammarParserCADENA, Tswift_GrammarParserID:
		localctx = NewListaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(402)
			p.e(0)
		}
		{
			p.SetState(403)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.e(0)
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(405)
					p.Match(Tswift_GrammarParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(406)
					p.e(0)
				}

			}
			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaValoresHermanosContext(p, NewListaval_mat2Context(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_listaval_mat2)
			p.SetState(414)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(415)
				p.Match(Tswift_GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(416)
				p.Listaval_mat()
			}

		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_matContext is an interface to support dynamic dispatch.
type ISimple_matContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimple_matContext differentiates from other interfaces.
	IsSimple_matContext()
}

type Simple_matContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_matContext() *Simple_matContext {
	var p = new(Simple_matContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_simple_mat
	return p
}

func InitEmptySimple_matContext(p *Simple_matContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_simple_mat
}

func (*Simple_matContext) IsSimple_matContext() {}

func NewSimple_matContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_matContext {
	var p = new(Simple_matContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_simple_mat

	return p
}

func (s *Simple_matContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_matContext) CopyAll(ctx *Simple_matContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Simple_matContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_matContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Simple_MatContext struct {
	Simple_matContext
}

func NewSimple_MatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Simple_MatContext {
	var p = new(Simple_MatContext)

	InitEmptySimple_matContext(&p.Simple_matContext)
	p.parser = parser
	p.CopyAll(ctx.(*Simple_matContext))

	return p
}

func (s *Simple_MatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_MatContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Simple_MatContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Simple_MatContext) REPEATING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserREPEATING, 0)
}

func (s *Simple_MatContext) AllDOSPT() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserDOSPT)
}

func (s *Simple_MatContext) DOSPT(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, i)
}

func (s *Simple_MatContext) Simple_mat() ISimple_matContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matContext)
}

func (s *Simple_MatContext) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCOUNT, 0)
}

func (s *Simple_MatContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, 0)
}

func (s *Simple_MatContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Simple_MatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSimple_Mat(s)
	}
}

func (s *Simple_MatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSimple_Mat(s)
	}
}

func (s *Simple_MatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSimple_Mat(s)

	default:
		return t.VisitChildren(s)
	}
}

type Simple_Mat_ExprContext struct {
	Simple_matContext
}

func NewSimple_Mat_ExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Simple_Mat_ExprContext {
	var p = new(Simple_Mat_ExprContext)

	InitEmptySimple_matContext(&p.Simple_matContext)
	p.parser = parser
	p.CopyAll(ctx.(*Simple_matContext))

	return p
}

func (s *Simple_Mat_ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_Mat_ExprContext) Tipo_matriz() ITipo_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipo_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipo_matrizContext)
}

func (s *Simple_Mat_ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Simple_Mat_ExprContext) REPEATING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserREPEATING, 0)
}

func (s *Simple_Mat_ExprContext) AllDOSPT() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserDOSPT)
}

func (s *Simple_Mat_ExprContext) DOSPT(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, i)
}

func (s *Simple_Mat_ExprContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Simple_Mat_ExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCOUNT, 0)
}

func (s *Simple_Mat_ExprContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, 0)
}

func (s *Simple_Mat_ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Simple_Mat_ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSimple_Mat_Expr(s)
	}
}

func (s *Simple_Mat_ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSimple_Mat_Expr(s)
	}
}

func (s *Simple_Mat_ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSimple_Mat_Expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Simple_mat() (localctx ISimple_matContext) {
	localctx = NewSimple_matContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Tswift_GrammarParserRULE_simple_mat)
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimple_MatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.Tipo_matriz()
		}
		{
			p.SetState(423)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(Tswift_GrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Simple_mat()
		}
		{
			p.SetState(427)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSimple_Mat_ExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Tipo_matriz()
		}
		{
			p.SetState(434)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(Tswift_GrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.e(0)
		}
		{
			p.SetState(438)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsig_matContext is an interface to support dynamic dispatch.
type IAsig_matContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsig_matContext differentiates from other interfaces.
	IsAsig_matContext()
}

type Asig_matContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsig_matContext() *Asig_matContext {
	var p = new(Asig_matContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asig_mat
	return p
}

func InitEmptyAsig_matContext(p *Asig_matContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_asig_mat
}

func (*Asig_matContext) IsAsig_matContext() {}

func NewAsig_matContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asig_matContext {
	var p = new(Asig_matContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_asig_mat

	return p
}

func (s *Asig_matContext) GetParser() antlr.Parser { return s.parser }

func (s *Asig_matContext) CopyAll(ctx *Asig_matContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asig_matContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asig_matContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Asig_MatContext struct {
	Asig_matContext
}

func NewAsig_MatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asig_MatContext {
	var p = new(Asig_MatContext)

	InitEmptyAsig_matContext(&p.Asig_matContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asig_matContext))

	return p
}

func (s *Asig_MatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asig_MatContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Asig_MatContext) AllCORCHETEIZQ() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserCORCHETEIZQ)
}

func (s *Asig_MatContext) CORCHETEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, i)
}

func (s *Asig_MatContext) AllENTERO() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserENTERO)
}

func (s *Asig_MatContext) ENTERO(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, i)
}

func (s *Asig_MatContext) AllCORCHETEDER() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserCORCHETEDER)
}

func (s *Asig_MatContext) CORCHETEDER(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, i)
}

func (s *Asig_MatContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Asig_MatContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Asig_MatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterAsig_Mat(s)
	}
}

func (s *Asig_MatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitAsig_Mat(s)
	}
}

func (s *Asig_MatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitAsig_Mat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Asig_mat() (localctx IAsig_matContext) {
	localctx = NewAsig_matContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Tswift_GrammarParserRULE_asig_mat)
	var _la int

	localctx = NewAsig_MatContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(447)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)
		p.Match(Tswift_GrammarParserENTERO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Match(Tswift_GrammarParserENTERO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserCORCHETEIZQ {
		{
			p.SetState(453)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(461)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(462)
		p.e(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_metodoContext is an interface to support dynamic dispatch.
type IDeclaracion_metodoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_metodoContext differentiates from other interfaces.
	IsDeclaracion_metodoContext()
}

type Declaracion_metodoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_metodoContext() *Declaracion_metodoContext {
	var p = new(Declaracion_metodoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_metodo
	return p
}

func InitEmptyDeclaracion_metodoContext(p *Declaracion_metodoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_metodo
}

func (*Declaracion_metodoContext) IsDeclaracion_metodoContext() {}

func NewDeclaracion_metodoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_metodoContext {
	var p = new(Declaracion_metodoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_metodo

	return p
}

func (s *Declaracion_metodoContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_metodoContext) CopyAll(ctx *Declaracion_metodoContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_metodoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_metodoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_MetodoContext struct {
	Declaracion_metodoContext
}

func NewDeclaracion_MetodoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_MetodoContext {
	var p = new(Declaracion_MetodoContext)

	InitEmptyDeclaracion_metodoContext(&p.Declaracion_metodoContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_metodoContext))

	return p
}

func (s *Declaracion_MetodoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_MetodoContext) FUNC() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFUNC, 0)
}

func (s *Declaracion_MetodoContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_MetodoContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Declaracion_MetodoContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Declaracion_MetodoContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *Declaracion_MetodoContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *Declaracion_MetodoContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *Declaracion_MetodoContext) AllL_parametros() []IL_parametrosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_parametrosContext); ok {
			len++
		}
	}

	tst := make([]IL_parametrosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_parametrosContext); ok {
			tst[i] = t.(IL_parametrosContext)
			i++
		}
	}

	return tst
}

func (s *Declaracion_MetodoContext) L_parametros(i int) IL_parametrosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_parametrosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_parametrosContext)
}

func (s *Declaracion_MetodoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Metodo(s)
	}
}

func (s *Declaracion_MetodoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Metodo(s)
	}
}

func (s *Declaracion_MetodoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Metodo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Declaracion_metodo() (localctx IDeclaracion_metodoContext) {
	localctx = NewDeclaracion_metodoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Tswift_GrammarParserRULE_declaracion_metodo)
	var _la int

	localctx = NewDeclaracion_MetodoContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(Tswift_GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID {
		{
			p.SetState(467)
			p.L_parametros()
		}

		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(473)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(474)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(475)
		p.L_sentencias()
	}
	{
		p.SetState(476)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaracion_funcionContext is an interface to support dynamic dispatch.
type IDeclaracion_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_funcionContext differentiates from other interfaces.
	IsDeclaracion_funcionContext()
}

type Declaracion_funcionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_funcionContext() *Declaracion_funcionContext {
	var p = new(Declaracion_funcionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_funcion
	return p
}

func InitEmptyDeclaracion_funcionContext(p *Declaracion_funcionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_funcion
}

func (*Declaracion_funcionContext) IsDeclaracion_funcionContext() {}

func NewDeclaracion_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_funcionContext {
	var p = new(Declaracion_funcionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_declaracion_funcion

	return p
}

func (s *Declaracion_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_funcionContext) CopyAll(ctx *Declaracion_funcionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_FuncionContext struct {
	Declaracion_funcionContext
}

func NewDeclaracion_FuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_FuncionContext {
	var p = new(Declaracion_FuncionContext)

	InitEmptyDeclaracion_funcionContext(&p.Declaracion_funcionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_funcionContext))

	return p
}

func (s *Declaracion_FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_FuncionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFUNC, 0)
}

func (s *Declaracion_FuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Declaracion_FuncionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Declaracion_FuncionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Declaracion_FuncionContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *Declaracion_FuncionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYOR, 0)
}

func (s *Declaracion_FuncionContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_FuncionContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *Declaracion_FuncionContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *Declaracion_FuncionContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *Declaracion_FuncionContext) AllL_parametros() []IL_parametrosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_parametrosContext); ok {
			len++
		}
	}

	tst := make([]IL_parametrosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_parametrosContext); ok {
			tst[i] = t.(IL_parametrosContext)
			i++
		}
	}

	return tst
}

func (s *Declaracion_FuncionContext) L_parametros(i int) IL_parametrosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_parametrosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_parametrosContext)
}

func (s *Declaracion_FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDeclaracion_Funcion(s)
	}
}

func (s *Declaracion_FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDeclaracion_Funcion(s)
	}
}

func (s *Declaracion_FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDeclaracion_Funcion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Declaracion_funcion() (localctx IDeclaracion_funcionContext) {
	localctx = NewDeclaracion_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Tswift_GrammarParserRULE_declaracion_funcion)
	var _la int

	localctx = NewDeclaracion_FuncionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		p.Match(Tswift_GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(479)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID {
		{
			p.SetState(481)
			p.L_parametros()
		}

		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(487)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(Tswift_GrammarParserMENOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Match(Tswift_GrammarParserMAYOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
		p.Tipo()
	}
	{
		p.SetState(491)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.L_sentencias()
	}
	{
		p.SetState(493)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_parametrosContext is an interface to support dynamic dispatch.
type IL_parametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_parametrosContext differentiates from other interfaces.
	IsL_parametrosContext()
}

type L_parametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_parametrosContext() *L_parametrosContext {
	var p = new(L_parametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_parametros
	return p
}

func InitEmptyL_parametrosContext(p *L_parametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_parametros
}

func (*L_parametrosContext) IsL_parametrosContext() {}

func NewL_parametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_parametrosContext {
	var p = new(L_parametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_parametros

	return p
}

func (s *L_parametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *L_parametrosContext) CopyAll(ctx *L_parametrosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_parametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_parametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_ParametrosContext struct {
	L_parametrosContext
	prim antlr.Token
}

func NewL_ParametrosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_ParametrosContext {
	var p = new(L_ParametrosContext)

	InitEmptyL_parametrosContext(&p.L_parametrosContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_parametrosContext))

	return p
}

func (s *L_ParametrosContext) GetPrim() antlr.Token { return s.prim }

func (s *L_ParametrosContext) SetPrim(v antlr.Token) { s.prim = v }

func (s *L_ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_ParametrosContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserID)
}

func (s *L_ParametrosContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, i)
}

func (s *L_ParametrosContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *L_ParametrosContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *L_ParametrosContext) INOUT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINOUT, 0)
}

func (s *L_ParametrosContext) GUIONBAJO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserGUIONBAJO, 0)
}

func (s *L_ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Parametros(s)
	}
}

func (s *L_ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Parametros(s)
	}
}

func (s *L_ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Parametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_parametros() (localctx IL_parametrosContext) {
	localctx = NewL_parametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Tswift_GrammarParserRULE_l_parametros)
	var _la int

	localctx = NewL_ParametrosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(496)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(495)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*L_ParametrosContext).prim = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*L_ParametrosContext).prim = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(498)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(499)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserINOUT {
		{
			p.SetState(500)
			p.Match(Tswift_GrammarParserINOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(503)
		p.Tipo()
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(504)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamada_funcionesContext is an interface to support dynamic dispatch.
type ILlamada_funcionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLlamada_funcionesContext differentiates from other interfaces.
	IsLlamada_funcionesContext()
}

type Llamada_funcionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamada_funcionesContext() *Llamada_funcionesContext {
	var p = new(Llamada_funcionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_llamada_funciones
	return p
}

func InitEmptyLlamada_funcionesContext(p *Llamada_funcionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_llamada_funciones
}

func (*Llamada_funcionesContext) IsLlamada_funcionesContext() {}

func NewLlamada_funcionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_funcionesContext {
	var p = new(Llamada_funcionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_llamada_funciones

	return p
}

func (s *Llamada_funcionesContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_funcionesContext) CopyAll(ctx *Llamada_funcionesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Llamada_funcionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_funcionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Llamada_FuncionContext struct {
	Llamada_funcionesContext
}

func NewLlamada_FuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Llamada_FuncionContext {
	var p = new(Llamada_FuncionContext)

	InitEmptyLlamada_funcionesContext(&p.Llamada_funcionesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Llamada_funcionesContext))

	return p
}

func (s *Llamada_FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_FuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Llamada_FuncionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Llamada_FuncionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Llamada_FuncionContext) AllL_argumentos() []IL_argumentosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_argumentosContext); ok {
			len++
		}
	}

	tst := make([]IL_argumentosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_argumentosContext); ok {
			tst[i] = t.(IL_argumentosContext)
			i++
		}
	}

	return tst
}

func (s *Llamada_FuncionContext) L_argumentos(i int) IL_argumentosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_argumentosContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_argumentosContext)
}

func (s *Llamada_FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterLlamada_Funcion(s)
	}
}

func (s *Llamada_FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitLlamada_Funcion(s)
	}
}

func (s *Llamada_FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitLlamada_Funcion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Llamada_funciones() (localctx ILlamada_funcionesContext) {
	localctx = NewLlamada_funcionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Tswift_GrammarParserRULE_llamada_funciones)
	var _la int

	localctx = NewLlamada_FuncionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305280006646390776) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2003) != 0) {
		{
			p.SetState(509)
			p.L_argumentos()
		}

		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(515)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_argumentosContext is an interface to support dynamic dispatch.
type IL_argumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_argumentosContext differentiates from other interfaces.
	IsL_argumentosContext()
}

type L_argumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_argumentosContext() *L_argumentosContext {
	var p = new(L_argumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_argumentos
	return p
}

func InitEmptyL_argumentosContext(p *L_argumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_argumentos
}

func (*L_argumentosContext) IsL_argumentosContext() {}

func NewL_argumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_argumentosContext {
	var p = new(L_argumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_argumentos

	return p
}

func (s *L_argumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *L_argumentosContext) CopyAll(ctx *L_argumentosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_argumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_argumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_ArgumentosContext struct {
	L_argumentosContext
}

func NewL_ArgumentosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_ArgumentosContext {
	var p = new(L_ArgumentosContext)

	InitEmptyL_argumentosContext(&p.L_argumentosContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_argumentosContext))

	return p
}

func (s *L_ArgumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_ArgumentosContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *L_ArgumentosContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *L_ArgumentosContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *L_ArgumentosContext) DIR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIR, 0)
}

func (s *L_ArgumentosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Argumentos(s)
	}
}

func (s *L_ArgumentosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Argumentos(s)
	}
}

func (s *L_ArgumentosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Argumentos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_argumentos() (localctx IL_argumentosContext) {
	localctx = NewL_argumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Tswift_GrammarParserRULE_l_argumentos)
	var _la int

	localctx = NewL_ArgumentosContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(519)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(517)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDIR {
		{
			p.SetState(521)
			p.Match(Tswift_GrammarParserDIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(524)
		p.e(0)
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(525)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunciones_embebidasContext is an interface to support dynamic dispatch.
type IFunciones_embebidasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunciones_embebidasContext differentiates from other interfaces.
	IsFunciones_embebidasContext()
}

type Funciones_embebidasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunciones_embebidasContext() *Funciones_embebidasContext {
	var p = new(Funciones_embebidasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_funciones_embebidas
	return p
}

func InitEmptyFunciones_embebidasContext(p *Funciones_embebidasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_funciones_embebidas
}

func (*Funciones_embebidasContext) IsFunciones_embebidasContext() {}

func NewFunciones_embebidasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funciones_embebidasContext {
	var p = new(Funciones_embebidasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_funciones_embebidas

	return p
}

func (s *Funciones_embebidasContext) GetParser() antlr.Parser { return s.parser }

func (s *Funciones_embebidasContext) CopyAll(ctx *Funciones_embebidasContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Funciones_embebidasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funciones_embebidasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Func_StringContext struct {
	Funciones_embebidasContext
}

func NewFunc_StringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_StringContext {
	var p = new(Func_StringContext)

	InitEmptyFunciones_embebidasContext(&p.Funciones_embebidasContext)
	p.parser = parser
	p.CopyAll(ctx.(*Funciones_embebidasContext))

	return p
}

func (s *Func_StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_StringContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_StringContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Func_StringContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSTRING, 0)
}

func (s *Func_StringContext) IOTA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIOTA, 0)
}

func (s *Func_StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_String(s)
	}
}

func (s *Func_StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_String(s)
	}
}

func (s *Func_StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_String(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_IntContext struct {
	Funciones_embebidasContext
}

func NewFunc_IntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_IntContext {
	var p = new(Func_IntContext)

	InitEmptyFunciones_embebidasContext(&p.Funciones_embebidasContext)
	p.parser = parser
	p.CopyAll(ctx.(*Funciones_embebidasContext))

	return p
}

func (s *Func_IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_IntContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_IntContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Func_IntContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_IntContext) INT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINT, 0)
}

func (s *Func_IntContext) ATOI() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserATOI, 0)
}

func (s *Func_IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Int(s)
	}
}

func (s *Func_IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Int(s)
	}
}

func (s *Func_IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Int(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_FloatContext struct {
	Funciones_embebidasContext
}

func NewFunc_FloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_FloatContext {
	var p = new(Func_FloatContext)

	InitEmptyFunciones_embebidasContext(&p.Funciones_embebidasContext)
	p.parser = parser
	p.CopyAll(ctx.(*Funciones_embebidasContext))

	return p
}

func (s *Func_FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFLOAT, 0)
}

func (s *Func_FloatContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Func_FloatContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Func_FloatContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Func_FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Float(s)
	}
}

func (s *Func_FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Float(s)
	}
}

func (s *Func_FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Float(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Funciones_embebidas() (localctx IFunciones_embebidasContext) {
	localctx = NewFunciones_embebidasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Tswift_GrammarParserRULE_funciones_embebidas)
	var _la int

	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserATOI, Tswift_GrammarParserINT:
		localctx = NewFunc_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(528)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserATOI || _la == Tswift_GrammarParserINT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(529)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.e(0)
		}
		{
			p.SetState(531)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserFLOAT:
		localctx = NewFunc_FloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
			p.Match(Tswift_GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(534)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.e(0)
		}
		{
			p.SetState(536)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserIOTA, Tswift_GrammarParserSTRING:
		localctx = NewFunc_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(538)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserIOTA || _la == Tswift_GrammarParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(539)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.e(0)
		}
		{
			p.SetState(541)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDef_structContext is an interface to support dynamic dispatch.
type IDef_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDef_structContext differentiates from other interfaces.
	IsDef_structContext()
}

type Def_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_structContext() *Def_structContext {
	var p = new(Def_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_struct
	return p
}

func InitEmptyDef_structContext(p *Def_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_def_struct
}

func (*Def_structContext) IsDef_structContext() {}

func NewDef_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_structContext {
	var p = new(Def_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_def_struct

	return p
}

func (s *Def_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_structContext) CopyAll(ctx *Def_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Def_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_StructContext struct {
	Def_structContext
}

func NewDef_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_StructContext {
	var p = new(Def_StructContext)

	InitEmptyDef_structContext(&p.Def_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_structContext))

	return p
}

func (s *Def_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_StructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSTRUCT, 0)
}

func (s *Def_StructContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Def_StructContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEIZQ, 0)
}

func (s *Def_StructContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLLAVEDER, 0)
}

func (s *Def_StructContext) AllL_sentencias_struct() []IL_sentencias_structContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_sentencias_structContext); ok {
			len++
		}
	}

	tst := make([]IL_sentencias_structContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_sentencias_structContext); ok {
			tst[i] = t.(IL_sentencias_structContext)
			i++
		}
	}

	return tst
}

func (s *Def_StructContext) L_sentencias_struct(i int) IL_sentencias_structContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentencias_structContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentencias_structContext)
}

func (s *Def_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterDef_Struct(s)
	}
}

func (s *Def_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitDef_Struct(s)
	}
}

func (s *Def_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitDef_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Def_struct() (localctx IDef_structContext) {
	localctx = NewDef_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Tswift_GrammarParserRULE_def_struct)
	var _la int

	localctx = NewDef_StructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		p.Match(Tswift_GrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(546)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(547)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&648518359226253312) != 0 {
		{
			p.SetState(548)
			p.L_sentencias_struct()
		}

		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(554)
		p.Match(Tswift_GrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IL_sentencias_structContext is an interface to support dynamic dispatch.
type IL_sentencias_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_sentencias_structContext differentiates from other interfaces.
	IsL_sentencias_structContext()
}

type L_sentencias_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_sentencias_structContext() *L_sentencias_structContext {
	var p = new(L_sentencias_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias_struct
	return p
}

func InitEmptyL_sentencias_structContext(p *L_sentencias_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias_struct
}

func (*L_sentencias_structContext) IsL_sentencias_structContext() {}

func NewL_sentencias_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_sentencias_structContext {
	var p = new(L_sentencias_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_l_sentencias_struct

	return p
}

func (s *L_sentencias_structContext) GetParser() antlr.Parser { return s.parser }

func (s *L_sentencias_structContext) CopyAll(ctx *L_sentencias_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_sentencias_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_sentencias_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_FuncionesContext struct {
	L_sentencias_structContext
}

func NewL_FuncionesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_FuncionesContext {
	var p = new(L_FuncionesContext)

	InitEmptyL_sentencias_structContext(&p.L_sentencias_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentencias_structContext))

	return p
}

func (s *L_FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_FuncionesContext) Declaracion_metodo() IDeclaracion_metodoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_metodoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_metodoContext)
}

func (s *L_FuncionesContext) Declaracion_funcion() IDeclaracion_funcionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_funcionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_funcionContext)
}

func (s *L_FuncionesContext) MUTATING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMUTATING, 0)
}

func (s *L_FuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Funciones(s)
	}
}

func (s *L_FuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Funciones(s)
	}
}

func (s *L_FuncionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Funciones(s)

	default:
		return t.VisitChildren(s)
	}
}

type L_AtributosContext struct {
	L_sentencias_structContext
	tipod antlr.Token
}

func NewL_AtributosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_AtributosContext {
	var p = new(L_AtributosContext)

	InitEmptyL_sentencias_structContext(&p.L_sentencias_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentencias_structContext))

	return p
}

func (s *L_AtributosContext) GetTipod() antlr.Token { return s.tipod }

func (s *L_AtributosContext) SetTipod(v antlr.Token) { s.tipod = v }

func (s *L_AtributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_AtributosContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *L_AtributosContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *L_AtributosContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *L_AtributosContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *L_AtributosContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *L_AtributosContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *L_AtributosContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *L_AtributosContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPTCOMA, 0)
}

func (s *L_AtributosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterL_Atributos(s)
	}
}

func (s *L_AtributosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitL_Atributos(s)
	}
}

func (s *L_AtributosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitL_Atributos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) L_sentencias_struct() (localctx IL_sentencias_structContext) {
	localctx = NewL_sentencias_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Tswift_GrammarParserRULE_l_sentencias_struct)
	var _la int

	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserVAR, Tswift_GrammarParserLET:
		localctx = NewL_AtributosContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(556)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*L_AtributosContext).tipod = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*L_AtributosContext).tipod = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(557)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserDOSPT {
			{
				p.SetState(558)
				p.Match(Tswift_GrammarParserDOSPT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(559)
				p.Tipo()
			}

		}
		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserIGUAL {
			{
				p.SetState(562)
				p.Match(Tswift_GrammarParserIGUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(563)
				p.e(0)
			}

		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(566)
				p.Match(Tswift_GrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case Tswift_GrammarParserFUNC, Tswift_GrammarParserMUTATING:
		localctx = NewL_FuncionesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserMUTATING {
			{
				p.SetState(569)
				p.Match(Tswift_GrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(572)
				p.Declaracion_metodo()
			}

		case 2:
			{
				p.SetState(573)
				p.Declaracion_funcion()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelf_dataContext is an interface to support dynamic dispatch.
type ISelf_dataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelf_dataContext differentiates from other interfaces.
	IsSelf_dataContext()
}

type Self_dataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelf_dataContext() *Self_dataContext {
	var p = new(Self_dataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_self_data
	return p
}

func InitEmptySelf_dataContext(p *Self_dataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_self_data
}

func (*Self_dataContext) IsSelf_dataContext() {}

func NewSelf_dataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Self_dataContext {
	var p = new(Self_dataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_self_data

	return p
}

func (s *Self_dataContext) GetParser() antlr.Parser { return s.parser }

func (s *Self_dataContext) CopyAll(ctx *Self_dataContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Self_dataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Self_dataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Self_DataContext struct {
	Self_dataContext
}

func NewSelf_DataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Self_DataContext {
	var p = new(Self_DataContext)

	InitEmptySelf_dataContext(&p.Self_dataContext)
	p.parser = parser
	p.CopyAll(ctx.(*Self_dataContext))

	return p
}

func (s *Self_DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Self_DataContext) SELF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSELF, 0)
}

func (s *Self_DataContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Self_DataContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Self_DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterSelf_Data(s)
	}
}

func (s *Self_DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitSelf_Data(s)
	}
}

func (s *Self_DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitSelf_Data(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Self_data() (localctx ISelf_dataContext) {
	localctx = NewSelf_dataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Tswift_GrammarParserRULE_self_data)
	localctx = NewSelf_DataContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)
		p.Match(Tswift_GrammarParserSELF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(579)
		p.Match(Tswift_GrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(580)
		p.Asignacion()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_dataContext is an interface to support dynamic dispatch.
type IStruct_dataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_dataContext differentiates from other interfaces.
	IsStruct_dataContext()
}

type Struct_dataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_dataContext() *Struct_dataContext {
	var p = new(Struct_dataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_struct_data
	return p
}

func InitEmptyStruct_dataContext(p *Struct_dataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_struct_data
}

func (*Struct_dataContext) IsStruct_dataContext() {}

func NewStruct_dataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_dataContext {
	var p = new(Struct_dataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_struct_data

	return p
}

func (s *Struct_dataContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_dataContext) CopyAll(ctx *Struct_dataContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_dataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_dataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Struct_DataContext struct {
	Struct_dataContext
	oper antlr.Token
}

func NewStruct_DataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Struct_DataContext {
	var p = new(Struct_DataContext)

	InitEmptyStruct_dataContext(&p.Struct_dataContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_dataContext))

	return p
}

func (s *Struct_DataContext) GetOper() antlr.Token { return s.oper }

func (s *Struct_DataContext) SetOper(v antlr.Token) { s.oper = v }

func (s *Struct_DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_DataContext) Idstruct() IIdstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdstructContext)
}

func (s *Struct_DataContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Struct_DataContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Struct_DataContext) MASIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMASIGUAL, 0)
}

func (s *Struct_DataContext) MENOSIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOSIGUAL, 0)
}

func (s *Struct_DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterStruct_Data(s)
	}
}

func (s *Struct_DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitStruct_Data(s)
	}
}

func (s *Struct_DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitStruct_Data(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Struct_data() (localctx IStruct_dataContext) {
	localctx = NewStruct_dataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Tswift_GrammarParserRULE_struct_data)
	var _la int

	localctx = NewStruct_DataContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(582)
		p.Idstruct()
	}
	{
		p.SetState(583)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Struct_DataContext).oper = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114688) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Struct_DataContext).oper = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(584)
		p.e(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdstructContext is an interface to support dynamic dispatch.
type IIdstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdstructContext differentiates from other interfaces.
	IsIdstructContext()
}

type IdstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdstructContext() *IdstructContext {
	var p = new(IdstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_idstruct
	return p
}

func InitEmptyIdstructContext(p *IdstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_idstruct
}

func (*IdstructContext) IsIdstructContext() {}

func NewIdstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdstructContext {
	var p = new(IdstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_idstruct

	return p
}

func (s *IdstructContext) GetParser() antlr.Parser { return s.parser }

func (s *IdstructContext) CopyAll(ctx *IdstructContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Id_StructContext struct {
	IdstructContext
}

func NewId_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Id_StructContext {
	var p = new(Id_StructContext)

	InitEmptyIdstructContext(&p.IdstructContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdstructContext))

	return p
}

func (s *Id_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_StructContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Id_StructContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Id_StructContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Id_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterId_Struct(s)
	}
}

func (s *Id_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitId_Struct(s)
	}
}

func (s *Id_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitId_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Idstruct() (localctx IIdstructContext) {
	localctx = NewIdstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Tswift_GrammarParserRULE_idstruct)
	localctx = NewId_StructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(586)
		p.e(0)
	}
	{
		p.SetState(587)
		p.Match(Tswift_GrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(588)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_llamadafuncContext is an interface to support dynamic dispatch.
type IStruct_llamadafuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_llamadafuncContext differentiates from other interfaces.
	IsStruct_llamadafuncContext()
}

type Struct_llamadafuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_llamadafuncContext() *Struct_llamadafuncContext {
	var p = new(Struct_llamadafuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_struct_llamadafunc
	return p
}

func InitEmptyStruct_llamadafuncContext(p *Struct_llamadafuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_struct_llamadafunc
}

func (*Struct_llamadafuncContext) IsStruct_llamadafuncContext() {}

func NewStruct_llamadafuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_llamadafuncContext {
	var p = new(Struct_llamadafuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_struct_llamadafunc

	return p
}

func (s *Struct_llamadafuncContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_llamadafuncContext) CopyAll(ctx *Struct_llamadafuncContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_llamadafuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_llamadafuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Struct_Llamada_FuncContext struct {
	Struct_llamadafuncContext
}

func NewStruct_Llamada_FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Struct_Llamada_FuncContext {
	var p = new(Struct_Llamada_FuncContext)

	InitEmptyStruct_llamadafuncContext(&p.Struct_llamadafuncContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_llamadafuncContext))

	return p
}

func (s *Struct_Llamada_FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_Llamada_FuncContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Struct_Llamada_FuncContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Struct_Llamada_FuncContext) Llamada_funciones() ILlamada_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionesContext)
}

func (s *Struct_Llamada_FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterStruct_Llamada_Func(s)
	}
}

func (s *Struct_Llamada_FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitStruct_Llamada_Func(s)
	}
}

func (s *Struct_Llamada_FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitStruct_Llamada_Func(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Struct_llamadafunc() (localctx IStruct_llamadafuncContext) {
	localctx = NewStruct_llamadafuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Tswift_GrammarParserRULE_struct_llamadafunc)
	localctx = NewStruct_Llamada_FuncContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.e(0)
	}
	{
		p.SetState(591)
		p.Match(Tswift_GrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(592)
		p.Llamada_funciones()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) CopyAll(ctx *TipoContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Tipo_VectorContext struct {
	TipoContext
}

func NewTipo_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_VectorContext {
	var p = new(Tipo_VectorContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Tipo_VectorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Tipo_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Tipo_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Vector(s)
	}
}

func (s *Tipo_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Vector(s)
	}
}

func (s *Tipo_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_StructContext struct {
	TipoContext
}

func NewTipo_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_StructContext {
	var p = new(Tipo_StructContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_StructContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Tipo_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Struct(s)
	}
}

func (s *Tipo_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Struct(s)
	}
}

func (s *Tipo_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_FloatContext struct {
	TipoContext
}

func NewTipo_FloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_FloatContext {
	var p = new(Tipo_FloatContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFLOAT, 0)
}

func (s *Tipo_FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Float(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_CharacterContext struct {
	TipoContext
}

func NewTipo_CharacterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_CharacterContext {
	var p = new(Tipo_CharacterContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_CharacterContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCHARACTER, 0)
}

func (s *Tipo_CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Character(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_StringContext struct {
	TipoContext
}

func NewTipo_StringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_StringContext {
	var p = new(Tipo_StringContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSTRING, 0)
}

func (s *Tipo_StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_String(s)
	}
}

func (s *Tipo_StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_String(s)
	}
}

func (s *Tipo_StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_String(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_BoolContext struct {
	TipoContext
}

func NewTipo_BoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_BoolContext {
	var p = new(Tipo_BoolContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserBOOL, 0)
}

func (s *Tipo_BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Bool(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_IntContext struct {
	TipoContext
}

func NewTipo_IntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_IntContext {
	var p = new(Tipo_IntContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_IntContext) INT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserINT, 0)
}

func (s *Tipo_IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterTipo_Int(s)
	}
}

func (s *Tipo_IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitTipo_Int(s)
	}
}

func (s *Tipo_IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitTipo_Int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Tswift_GrammarParserRULE_tipo)
	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserINT:
		localctx = NewTipo_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(594)
			p.Match(Tswift_GrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserFLOAT:
		localctx = NewTipo_FloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(Tswift_GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserSTRING:
		localctx = NewTipo_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(596)
			p.Match(Tswift_GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserBOOL:
		localctx = NewTipo_BoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(597)
			p.Match(Tswift_GrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCHARACTER:
		localctx = NewTipo_CharacterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(598)
			p.Match(Tswift_GrammarParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserID:
		localctx = NewTipo_StructContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(599)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCORCHETEIZQ:
		localctx = NewTipo_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(600)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Tipo()
		}
		{
			p.SetState(602)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_RelContext struct {
	EContext
	op antlr.Token
}

func NewExpr_RelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_RelContext {
	var p = new(Expr_RelContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_RelContext) GetOp() antlr.Token { return s.op }

func (s *Expr_RelContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_RelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_RelContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_RelContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_RelContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUALIGUAL, 0)
}

func (s *Expr_RelContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIFERENTE, 0)
}

func (s *Expr_RelContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYORIGUAL, 0)
}

func (s *Expr_RelContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYOR, 0)
}

func (s *Expr_RelContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENORIGUAL, 0)
}

func (s *Expr_RelContext) MENOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOR, 0)
}

func (s *Expr_RelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Rel(s)
	}
}

func (s *Expr_RelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Rel(s)
	}
}

func (s *Expr_RelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Rel(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_Funciones_EmbebidasContext struct {
	EContext
}

func NewExpr_Funciones_EmbebidasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_Funciones_EmbebidasContext {
	var p = new(Expr_Funciones_EmbebidasContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_Funciones_EmbebidasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_Funciones_EmbebidasContext) Funciones_embebidas() IFunciones_embebidasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunciones_embebidasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunciones_embebidasContext)
}

func (s *Expr_Funciones_EmbebidasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Funciones_Embebidas(s)
	}
}

func (s *Expr_Funciones_EmbebidasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Funciones_Embebidas(s)
	}
}

func (s *Expr_Funciones_EmbebidasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Funciones_Embebidas(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_DecimalContext struct {
	EContext
}

func NewExpr_DecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_DecimalContext {
	var p = new(Expr_DecimalContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDECIMAL, 0)
}

func (s *Expr_DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Decimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CaracterContext struct {
	EContext
}

func NewExpr_CaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CaracterContext {
	var p = new(Expr_CaracterContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCARACTER, 0)
}

func (s *Expr_CaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Caracter(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_SumResContext struct {
	EContext
	op antlr.Token
}

func NewExpr_SumResContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_SumResContext {
	var p = new(Expr_SumResContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_SumResContext) GetOp() antlr.Token { return s.op }

func (s *Expr_SumResContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_SumResContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_SumResContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_SumResContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_SumResContext) MAS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAS, 0)
}

func (s *Expr_SumResContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *Expr_SumResContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_SumRes(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_Llamada_Funcion_StructContext struct {
	EContext
}

func NewExpr_Llamada_Funcion_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_Llamada_Funcion_StructContext {
	var p = new(Expr_Llamada_Funcion_StructContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_Llamada_Funcion_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_Llamada_Funcion_StructContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_Llamada_Funcion_StructContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Expr_Llamada_Funcion_StructContext) Llamada_funciones() ILlamada_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionesContext)
}

func (s *Expr_Llamada_Funcion_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Llamada_Funcion_Struct(s)
	}
}

func (s *Expr_Llamada_Funcion_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Llamada_Funcion_Struct(s)
	}
}

func (s *Expr_Llamada_Funcion_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Llamada_Funcion_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_MatrizContext struct {
	EContext
}

func NewExpr_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_MatrizContext {
	var p = new(Expr_MatrizContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_MatrizContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_MatrizContext) AllCORCHETEIZQ() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserCORCHETEIZQ)
}

func (s *Expr_MatrizContext) CORCHETEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, i)
}

func (s *Expr_MatrizContext) AllENTERO() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserENTERO)
}

func (s *Expr_MatrizContext) ENTERO(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, i)
}

func (s *Expr_MatrizContext) AllCORCHETEDER() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserCORCHETEDER)
}

func (s *Expr_MatrizContext) CORCHETEDER(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, i)
}

func (s *Expr_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Matriz(s)
	}
}

func (s *Expr_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Matriz(s)
	}
}

func (s *Expr_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_NegContext struct {
	EContext
	op antlr.Token
}

func NewExpr_NegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NegContext {
	var p = new(Expr_NegContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NegContext) GetOp() antlr.Token { return s.op }

func (s *Expr_NegContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_NegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NegContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_NegContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *Expr_NegContext) NOT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserNOT, 0)
}

func (s *Expr_NegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Neg(s)
	}
}

func (s *Expr_NegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Neg(s)
	}
}

func (s *Expr_NegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Neg(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_MulDivContext struct {
	EContext
	op antlr.Token
}

func NewExpr_MulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_MulDivContext {
	var p = new(Expr_MulDivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_MulDivContext) GetOp() antlr.Token { return s.op }

func (s *Expr_MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_MulDivContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_MulDivContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_MulDivContext) POR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPOR, 0)
}

func (s *Expr_MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIV, 0)
}

func (s *Expr_MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_MulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_NilContext struct {
	EContext
}

func NewExpr_NilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NilContext {
	var p = new(Expr_NilContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserNIL, 0)
}

func (s *Expr_NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Nil(s)
	}
}

func (s *Expr_NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Nil(s)
	}
}

func (s *Expr_NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Nil(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_isEmptyContext struct {
	EContext
}

func NewFunc_Vector_isEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_isEmptyContext {
	var p = new(Func_Vector_isEmptyContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Func_Vector_isEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_isEmptyContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Func_Vector_isEmptyContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Func_Vector_isEmptyContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserISEMPTY, 0)
}

func (s *Func_Vector_isEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Vector_isEmpty(s)
	}
}

func (s *Func_Vector_isEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Vector_isEmpty(s)
	}
}

func (s *Func_Vector_isEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Vector_isEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_Llamada_FuncionContext struct {
	EContext
}

func NewExpr_Llamada_FuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_Llamada_FuncionContext {
	var p = new(Expr_Llamada_FuncionContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_Llamada_FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_Llamada_FuncionContext) Llamada_funciones() ILlamada_funcionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamada_funcionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionesContext)
}

func (s *Expr_Llamada_FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Llamada_Funcion(s)
	}
}

func (s *Expr_Llamada_FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Llamada_Funcion(s)
	}
}

func (s *Expr_Llamada_FuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Llamada_Funcion(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_CountContext struct {
	EContext
}

func NewFunc_Vector_CountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_CountContext {
	var p = new(Func_Vector_CountContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Func_Vector_CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_CountContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Func_Vector_CountContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Func_Vector_CountContext) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCOUNT, 0)
}

func (s *Func_Vector_CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterFunc_Vector_Count(s)
	}
}

func (s *Func_Vector_CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitFunc_Vector_Count(s)
	}
}

func (s *Func_Vector_CountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitFunc_Vector_Count(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_StructContext struct {
	EContext
}

func NewExpr_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_StructContext {
	var p = new(Expr_StructContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_StructContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_StructContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Expr_StructContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Struct(s)
	}
}

func (s *Expr_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Struct(s)
	}
}

func (s *Expr_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CadenaContext struct {
	EContext
}

func NewExpr_CadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CadenaContext {
	var p = new(Expr_CadenaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCADENA, 0)
}

func (s *Expr_CadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Cadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_SelfContext struct {
	EContext
}

func NewExpr_SelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_SelfContext {
	var p = new(Expr_SelfContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_SelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_SelfContext) SELF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserSELF, 0)
}

func (s *Expr_SelfContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPUNTO, 0)
}

func (s *Expr_SelfContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_SelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Self(s)
	}
}

func (s *Expr_SelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Self(s)
	}
}

func (s *Expr_SelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Self(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_IdContext struct {
	EContext
}

func NewExpr_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_IdContext {
	var p = new(Expr_IdContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Id(s)
	}
}

func (s *Expr_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Id(s)
	}
}

func (s *Expr_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ModContext struct {
	EContext
}

func NewExpr_ModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ModContext {
	var p = new(Expr_ModContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ModContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_ModContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMOD, 0)
}

func (s *Expr_ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Mod(s)
	}
}

func (s *Expr_ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Mod(s)
	}
}

func (s *Expr_ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Mod(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ParContext struct {
	EContext
}

func NewExpr_ParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ParContext {
	var p = new(Expr_ParContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ParContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Expr_ParContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_ParContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Expr_ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Par(s)
	}
}

func (s *Expr_ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Par(s)
	}
}

func (s *Expr_ParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Par(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_LogicaContext struct {
	EContext
	op antlr.Token
}

func NewExpr_LogicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_LogicaContext {
	var p = new(Expr_LogicaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_LogicaContext) GetOp() antlr.Token { return s.op }

func (s *Expr_LogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_LogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_LogicaContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_LogicaContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_LogicaContext) AND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserAND, 0)
}

func (s *Expr_LogicaContext) OR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserOR, 0)
}

func (s *Expr_LogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Logica(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_BooleanoContext struct {
	EContext
	op antlr.Token
}

func NewExpr_BooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_BooleanoContext {
	var p = new(Expr_BooleanoContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_BooleanoContext) GetOp() antlr.Token { return s.op }

func (s *Expr_BooleanoContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_BooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_BooleanoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserTRUE, 0)
}

func (s *Expr_BooleanoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFALSE, 0)
}

func (s *Expr_BooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Booleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_VectorContext struct {
	EContext
}

func NewExpr_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_VectorContext {
	var p = new(Expr_VectorContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *Expr_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEIZQ, 0)
}

func (s *Expr_VectorContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserCORCHETEDER, 0)
}

func (s *Expr_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Vector(s)
	}
}

func (s *Expr_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Vector(s)
	}
}

func (s *Expr_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_EnteroContext struct {
	EContext
}

func NewExpr_EnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_EnteroContext {
	var p = new(Expr_EnteroContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_EnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_EnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, 0)
}

func (s *Expr_EnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitExpr_Entero(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *Tswift_GrammarParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, Tswift_GrammarParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(651)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpr_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(607)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_NegContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserMENOS || _la == Tswift_GrammarParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_NegContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(608)
			p.e(23)
		}

	case 2:
		localctx = NewExpr_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(609)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_BooleanoContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserTRUE || _la == Tswift_GrammarParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_BooleanoContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewExpr_SelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(610)
			p.Match(Tswift_GrammarParserSELF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(611)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpr_Funciones_EmbebidasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(613)
			p.Funciones_embebidas()
		}

	case 5:
		localctx = NewExpr_Llamada_FuncionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(614)
			p.Llamada_funciones()
		}

	case 6:
		localctx = NewExpr_MatrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(615)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(617)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(619)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(627)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(622)
					p.Match(Tswift_GrammarParserCORCHETEIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(623)
					p.Match(Tswift_GrammarParserENTERO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(624)
					p.Match(Tswift_GrammarParserCORCHETEDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(629)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpr_VectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(630)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(631)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(632)
			p.e(0)
		}
		{
			p.SetState(633)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewFunc_Vector_isEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(635)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(637)
			p.Match(Tswift_GrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewFunc_Vector_CountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(638)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(639)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(640)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExpr_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(641)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewExpr_NilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(642)
			p.Match(Tswift_GrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewExpr_DecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(643)
			p.Match(Tswift_GrammarParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewExpr_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(644)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewExpr_CadenaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(645)
			p.Match(Tswift_GrammarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewExpr_CaracterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(646)
			p.Match(Tswift_GrammarParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewExpr_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(647)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.e(0)
		}
		{
			p.SetState(649)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(676)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(674)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_MulDivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(653)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(654)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserDIV || _la == Tswift_GrammarParserPOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(655)
					p.e(23)
				}

			case 2:
				localctx = NewExpr_SumResContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(656)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(657)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_SumResContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserMENOS || _la == Tswift_GrammarParserMAS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_SumResContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(658)
					p.e(22)
				}

			case 3:
				localctx = NewExpr_ModContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(659)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(660)
					p.Match(Tswift_GrammarParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(661)
					p.e(21)
				}

			case 4:
				localctx = NewExpr_RelContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(662)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(663)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_RelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264241152) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_RelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(664)
					p.e(20)
				}

			case 5:
				localctx = NewExpr_LogicaContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(665)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(666)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_LogicaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserAND || _la == Tswift_GrammarParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_LogicaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(667)
					p.e(19)
				}

			case 6:
				localctx = NewExpr_StructContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(668)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(669)
					p.Match(Tswift_GrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(670)
					p.Match(Tswift_GrammarParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 7:
				localctx = NewExpr_Llamada_Funcion_StructContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(671)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(672)
					p.Match(Tswift_GrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(673)
					p.Llamada_funciones()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(678)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *Tswift_GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 24:
		var t *Listaval_mat2Context = nil
		if localctx != nil {
			t = localctx.(*Listaval_mat2Context)
		}
		return p.Listaval_mat2_Sempred(t, predIndex)

	case 40:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Tswift_GrammarParser) Listaval_mat2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Tswift_GrammarParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 12)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
