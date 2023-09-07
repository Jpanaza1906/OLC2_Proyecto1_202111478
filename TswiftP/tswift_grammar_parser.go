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
		"'Int'", "'Float'", "'Bool'", "'Character'", "'String'",
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
		"IOTA", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO",
		"DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR",
	}
	staticData.RuleNames = []string{
		"s", "l_sentencias", "sentencia", "print_sentencia", "declaracion",
		"constante", "asignacion", "if_sentencia", "switch_sentencia", "l_casos",
		"l_default", "guard_sentencia", "while_sentencia", "for_sentencia",
		"rango_p", "trans_sentencia", "dec_vector", "def_vector", "asig_vector",
		"func_vector", "declaracion_matrices", "tipo_matriz", "def_matriz",
		"listaval_mat", "listaval_mat2", "simple_mat", "asig_mat", "declaracion_metodo",
		"declaracion_funcion", "l_parametros", "llamada_funciones", "l_argumentos",
		"funciones_embebidas", "def_struct", "l_sentencias_struct", "creacion_struct",
		"tipo", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 76, 655, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 1, 0, 1, 1, 5, 1, 80, 8, 1, 10, 1, 12, 1, 83, 9, 1,
		1, 2, 1, 2, 3, 2, 87, 8, 2, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 2, 1, 2, 3,
		2, 95, 8, 2, 1, 2, 1, 2, 3, 2, 99, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 108, 8, 2, 1, 2, 1, 2, 3, 2, 112, 8, 2, 1, 2, 1, 2, 3,
		2, 116, 8, 2, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		126, 8, 2, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 2, 1, 2, 3, 2, 134, 8, 2, 3,
		2, 136, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 143, 8, 3, 10, 3, 12,
		3, 146, 9, 3, 1, 3, 1, 3, 3, 3, 150, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 169, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 182, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 193, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 219, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		4, 8, 225, 8, 8, 11, 8, 12, 8, 226, 1, 8, 3, 8, 230, 8, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 239, 8, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 245, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 266, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 280, 8, 15, 3, 15, 282,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 5, 17, 297, 8, 17, 10, 17, 12, 17, 300, 9, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 307, 8, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 330, 8,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 353, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 359, 8, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 372, 8, 21, 1, 22, 1, 22, 3, 22, 376, 8, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 389, 8,
		24, 10, 24, 12, 24, 392, 9, 24, 3, 24, 394, 8, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 399, 8, 24, 10, 24, 12, 24, 402, 9, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 426, 8,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		5, 26, 438, 8, 26, 10, 26, 12, 26, 441, 9, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 5, 27, 450, 8, 27, 10, 27, 12, 27, 453, 9, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 464,
		8, 28, 10, 28, 12, 28, 467, 9, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 3, 29, 478, 8, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		483, 8, 29, 1, 29, 1, 29, 3, 29, 487, 8, 29, 1, 30, 1, 30, 1, 30, 5, 30,
		492, 8, 30, 10, 30, 12, 30, 495, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 3,
		31, 501, 8, 31, 1, 31, 3, 31, 504, 8, 31, 1, 31, 1, 31, 3, 31, 508, 8,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 525, 8, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 5, 33, 531, 8, 33, 10, 33, 12, 33, 534, 9, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 542, 8, 34, 1, 34, 1, 34, 3, 34, 546,
		8, 34, 1, 34, 3, 34, 549, 8, 34, 1, 34, 3, 34, 552, 8, 34, 1, 34, 3, 34,
		555, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 561, 8, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 3, 35, 567, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3,
		35, 574, 8, 35, 1, 35, 1, 35, 3, 35, 578, 8, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 589, 8, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 5, 37, 607, 8, 37, 10, 37, 12, 37, 610, 9, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		3, 37, 633, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 650, 8, 37,
		10, 37, 12, 37, 653, 9, 37, 1, 37, 0, 2, 48, 74, 38, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 0, 10, 2, 0, 12,
		12, 73, 73, 1, 0, 32, 33, 2, 0, 61, 61, 63, 63, 2, 0, 62, 62, 67, 67, 2,
		0, 19, 19, 30, 30, 1, 0, 34, 35, 2, 0, 17, 17, 21, 21, 1, 0, 19, 20, 1,
		0, 22, 27, 1, 0, 28, 29, 723, 0, 76, 1, 0, 0, 0, 2, 81, 1, 0, 0, 0, 4,
		135, 1, 0, 0, 0, 6, 137, 1, 0, 0, 0, 8, 168, 1, 0, 0, 0, 10, 181, 1, 0,
		0, 0, 12, 192, 1, 0, 0, 0, 14, 218, 1, 0, 0, 0, 16, 220, 1, 0, 0, 0, 18,
		233, 1, 0, 0, 0, 20, 240, 1, 0, 0, 0, 22, 246, 1, 0, 0, 0, 24, 254, 1,
		0, 0, 0, 26, 260, 1, 0, 0, 0, 28, 271, 1, 0, 0, 0, 30, 281, 1, 0, 0, 0,
		32, 283, 1, 0, 0, 0, 34, 306, 1, 0, 0, 0, 36, 329, 1, 0, 0, 0, 38, 352,
		1, 0, 0, 0, 40, 354, 1, 0, 0, 0, 42, 371, 1, 0, 0, 0, 44, 375, 1, 0, 0,
		0, 46, 377, 1, 0, 0, 0, 48, 393, 1, 0, 0, 0, 50, 425, 1, 0, 0, 0, 52, 427,
		1, 0, 0, 0, 54, 445, 1, 0, 0, 0, 56, 459, 1, 0, 0, 0, 58, 477, 1, 0, 0,
		0, 60, 488, 1, 0, 0, 0, 62, 500, 1, 0, 0, 0, 64, 524, 1, 0, 0, 0, 66, 526,
		1, 0, 0, 0, 68, 554, 1, 0, 0, 0, 70, 577, 1, 0, 0, 0, 72, 588, 1, 0, 0,
		0, 74, 632, 1, 0, 0, 0, 76, 77, 3, 2, 1, 0, 77, 1, 1, 0, 0, 0, 78, 80,
		3, 4, 2, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 3, 6,
		3, 0, 85, 87, 5, 9, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 136,
		1, 0, 0, 0, 88, 90, 3, 8, 4, 0, 89, 91, 5, 9, 0, 0, 90, 89, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 136, 1, 0, 0, 0, 92, 94, 3, 10, 5, 0, 93, 95, 5,
		9, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 136, 1, 0, 0, 0, 96,
		98, 3, 12, 6, 0, 97, 99, 5, 9, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 99, 136, 1, 0, 0, 0, 100, 136, 3, 14, 7, 0, 101, 136, 3, 16, 8, 0,
		102, 136, 3, 22, 11, 0, 103, 136, 3, 24, 12, 0, 104, 136, 3, 26, 13, 0,
		105, 107, 3, 30, 15, 0, 106, 108, 5, 9, 0, 0, 107, 106, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 136, 1, 0, 0, 0, 109, 111, 3, 32, 16, 0, 110, 112,
		5, 9, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 136, 1, 0,
		0, 0, 113, 115, 3, 38, 19, 0, 114, 116, 5, 9, 0, 0, 115, 114, 1, 0, 0,
		0, 115, 116, 1, 0, 0, 0, 116, 136, 1, 0, 0, 0, 117, 119, 3, 36, 18, 0,
		118, 120, 5, 9, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		136, 1, 0, 0, 0, 121, 136, 3, 54, 27, 0, 122, 136, 3, 56, 28, 0, 123, 125,
		3, 60, 30, 0, 124, 126, 5, 9, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 136, 1, 0, 0, 0, 127, 129, 3, 40, 20, 0, 128, 130, 5, 9,
		0, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 136, 1, 0, 0, 0,
		131, 133, 3, 52, 26, 0, 132, 134, 5, 9, 0, 0, 133, 132, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 84, 1, 0, 0, 0, 135, 88, 1,
		0, 0, 0, 135, 92, 1, 0, 0, 0, 135, 96, 1, 0, 0, 0, 135, 100, 1, 0, 0, 0,
		135, 101, 1, 0, 0, 0, 135, 102, 1, 0, 0, 0, 135, 103, 1, 0, 0, 0, 135,
		104, 1, 0, 0, 0, 135, 105, 1, 0, 0, 0, 135, 109, 1, 0, 0, 0, 135, 113,
		1, 0, 0, 0, 135, 117, 1, 0, 0, 0, 135, 121, 1, 0, 0, 0, 135, 122, 1, 0,
		0, 0, 135, 123, 1, 0, 0, 0, 135, 127, 1, 0, 0, 0, 135, 131, 1, 0, 0, 0,
		136, 5, 1, 0, 0, 0, 137, 138, 5, 31, 0, 0, 138, 139, 5, 3, 0, 0, 139, 144,
		3, 74, 37, 0, 140, 141, 5, 1, 0, 0, 141, 143, 3, 74, 37, 0, 142, 140, 1,
		0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0,
		0, 145, 147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 149, 5, 2, 0, 0, 148,
		150, 5, 9, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 7, 1,
		0, 0, 0, 151, 152, 5, 32, 0, 0, 152, 153, 5, 73, 0, 0, 153, 154, 5, 8,
		0, 0, 154, 155, 3, 72, 36, 0, 155, 156, 5, 16, 0, 0, 156, 157, 3, 74, 37,
		0, 157, 169, 1, 0, 0, 0, 158, 159, 5, 32, 0, 0, 159, 160, 5, 73, 0, 0,
		160, 161, 5, 16, 0, 0, 161, 169, 3, 74, 37, 0, 162, 163, 5, 32, 0, 0, 163,
		164, 5, 73, 0, 0, 164, 165, 5, 8, 0, 0, 165, 166, 3, 72, 36, 0, 166, 167,
		5, 10, 0, 0, 167, 169, 1, 0, 0, 0, 168, 151, 1, 0, 0, 0, 168, 158, 1, 0,
		0, 0, 168, 162, 1, 0, 0, 0, 169, 9, 1, 0, 0, 0, 170, 171, 5, 33, 0, 0,
		171, 172, 5, 73, 0, 0, 172, 173, 5, 8, 0, 0, 173, 174, 3, 72, 36, 0, 174,
		175, 5, 16, 0, 0, 175, 176, 3, 74, 37, 0, 176, 182, 1, 0, 0, 0, 177, 178,
		5, 33, 0, 0, 178, 179, 5, 73, 0, 0, 179, 180, 5, 16, 0, 0, 180, 182, 3,
		74, 37, 0, 181, 170, 1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 182, 11, 1, 0, 0,
		0, 183, 184, 5, 73, 0, 0, 184, 185, 5, 14, 0, 0, 185, 193, 3, 74, 37, 0,
		186, 187, 5, 73, 0, 0, 187, 188, 5, 15, 0, 0, 188, 193, 3, 74, 37, 0, 189,
		190, 5, 73, 0, 0, 190, 191, 5, 16, 0, 0, 191, 193, 3, 74, 37, 0, 192, 183,
		1, 0, 0, 0, 192, 186, 1, 0, 0, 0, 192, 189, 1, 0, 0, 0, 193, 13, 1, 0,
		0, 0, 194, 195, 5, 36, 0, 0, 195, 196, 3, 74, 37, 0, 196, 197, 5, 4, 0,
		0, 197, 198, 3, 2, 1, 0, 198, 199, 5, 5, 0, 0, 199, 219, 1, 0, 0, 0, 200,
		201, 5, 36, 0, 0, 201, 202, 3, 74, 37, 0, 202, 203, 5, 4, 0, 0, 203, 204,
		3, 2, 1, 0, 204, 205, 5, 5, 0, 0, 205, 206, 5, 37, 0, 0, 206, 207, 5, 4,
		0, 0, 207, 208, 3, 2, 1, 0, 208, 209, 5, 5, 0, 0, 209, 219, 1, 0, 0, 0,
		210, 211, 5, 36, 0, 0, 211, 212, 3, 74, 37, 0, 212, 213, 5, 4, 0, 0, 213,
		214, 3, 2, 1, 0, 214, 215, 5, 5, 0, 0, 215, 216, 5, 37, 0, 0, 216, 217,
		3, 14, 7, 0, 217, 219, 1, 0, 0, 0, 218, 194, 1, 0, 0, 0, 218, 200, 1, 0,
		0, 0, 218, 210, 1, 0, 0, 0, 219, 15, 1, 0, 0, 0, 220, 221, 5, 38, 0, 0,
		221, 222, 3, 74, 37, 0, 222, 224, 5, 4, 0, 0, 223, 225, 3, 18, 9, 0, 224,
		223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 230, 3, 20, 10, 0, 229, 228, 1,
		0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 5, 5, 0,
		0, 232, 17, 1, 0, 0, 0, 233, 234, 5, 39, 0, 0, 234, 235, 3, 74, 37, 0,
		235, 236, 5, 8, 0, 0, 236, 238, 3, 2, 1, 0, 237, 239, 5, 48, 0, 0, 238,
		237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 19, 1, 0, 0, 0, 240, 241, 5,
		40, 0, 0, 241, 242, 5, 8, 0, 0, 242, 244, 3, 2, 1, 0, 243, 245, 5, 48,
		0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 21, 1, 0, 0, 0,
		246, 247, 5, 45, 0, 0, 247, 248, 3, 74, 37, 0, 248, 249, 5, 37, 0, 0, 249,
		250, 5, 4, 0, 0, 250, 251, 3, 2, 1, 0, 251, 252, 3, 30, 15, 0, 252, 253,
		5, 5, 0, 0, 253, 23, 1, 0, 0, 0, 254, 255, 5, 41, 0, 0, 255, 256, 3, 74,
		37, 0, 256, 257, 5, 4, 0, 0, 257, 258, 3, 2, 1, 0, 258, 259, 5, 5, 0, 0,
		259, 25, 1, 0, 0, 0, 260, 261, 5, 42, 0, 0, 261, 262, 7, 0, 0, 0, 262,
		265, 5, 43, 0, 0, 263, 266, 3, 28, 14, 0, 264, 266, 3, 74, 37, 0, 265,
		263, 1, 0, 0, 0, 265, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268,
		5, 4, 0, 0, 268, 269, 3, 2, 1, 0, 269, 270, 5, 5, 0, 0, 270, 27, 1, 0,
		0, 0, 271, 272, 3, 74, 37, 0, 272, 273, 5, 44, 0, 0, 273, 274, 3, 74, 37,
		0, 274, 29, 1, 0, 0, 0, 275, 282, 5, 48, 0, 0, 276, 282, 5, 46, 0, 0, 277,
		279, 5, 47, 0, 0, 278, 280, 3, 74, 37, 0, 279, 278, 1, 0, 0, 0, 279, 280,
		1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 275, 1, 0, 0, 0, 281, 276, 1, 0,
		0, 0, 281, 277, 1, 0, 0, 0, 282, 31, 1, 0, 0, 0, 283, 284, 7, 1, 0, 0,
		284, 285, 5, 73, 0, 0, 285, 286, 5, 8, 0, 0, 286, 287, 5, 6, 0, 0, 287,
		288, 3, 72, 36, 0, 288, 289, 5, 7, 0, 0, 289, 290, 5, 16, 0, 0, 290, 291,
		3, 34, 17, 0, 291, 33, 1, 0, 0, 0, 292, 293, 5, 6, 0, 0, 293, 298, 3, 74,
		37, 0, 294, 295, 5, 1, 0, 0, 295, 297, 3, 74, 37, 0, 296, 294, 1, 0, 0,
		0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302, 5, 7, 0, 0, 302, 307,
		1, 0, 0, 0, 303, 304, 5, 6, 0, 0, 304, 307, 5, 7, 0, 0, 305, 307, 5, 73,
		0, 0, 306, 292, 1, 0, 0, 0, 306, 303, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0,
		307, 35, 1, 0, 0, 0, 308, 309, 5, 73, 0, 0, 309, 310, 5, 6, 0, 0, 310,
		311, 3, 74, 37, 0, 311, 312, 5, 7, 0, 0, 312, 313, 5, 16, 0, 0, 313, 314,
		3, 74, 37, 0, 314, 330, 1, 0, 0, 0, 315, 316, 5, 73, 0, 0, 316, 317, 5,
		6, 0, 0, 317, 318, 3, 74, 37, 0, 318, 319, 5, 7, 0, 0, 319, 320, 5, 14,
		0, 0, 320, 321, 3, 74, 37, 0, 321, 330, 1, 0, 0, 0, 322, 323, 5, 73, 0,
		0, 323, 324, 5, 6, 0, 0, 324, 325, 3, 74, 37, 0, 325, 326, 5, 7, 0, 0,
		326, 327, 5, 15, 0, 0, 327, 328, 3, 74, 37, 0, 328, 330, 1, 0, 0, 0, 329,
		308, 1, 0, 0, 0, 329, 315, 1, 0, 0, 0, 329, 322, 1, 0, 0, 0, 330, 37, 1,
		0, 0, 0, 331, 332, 5, 73, 0, 0, 332, 333, 5, 11, 0, 0, 333, 334, 5, 50,
		0, 0, 334, 335, 5, 3, 0, 0, 335, 336, 3, 74, 37, 0, 336, 337, 5, 2, 0,
		0, 337, 353, 1, 0, 0, 0, 338, 339, 5, 73, 0, 0, 339, 340, 5, 11, 0, 0,
		340, 341, 5, 51, 0, 0, 341, 342, 5, 3, 0, 0, 342, 353, 5, 2, 0, 0, 343,
		344, 5, 73, 0, 0, 344, 345, 5, 11, 0, 0, 345, 346, 5, 52, 0, 0, 346, 347,
		5, 3, 0, 0, 347, 348, 5, 53, 0, 0, 348, 349, 5, 8, 0, 0, 349, 350, 3, 74,
		37, 0, 350, 351, 5, 2, 0, 0, 351, 353, 1, 0, 0, 0, 352, 331, 1, 0, 0, 0,
		352, 338, 1, 0, 0, 0, 352, 343, 1, 0, 0, 0, 353, 39, 1, 0, 0, 0, 354, 355,
		7, 1, 0, 0, 355, 358, 5, 73, 0, 0, 356, 357, 5, 8, 0, 0, 357, 359, 3, 42,
		21, 0, 358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0,
		360, 361, 5, 16, 0, 0, 361, 362, 3, 44, 22, 0, 362, 41, 1, 0, 0, 0, 363,
		364, 5, 6, 0, 0, 364, 365, 3, 42, 21, 0, 365, 366, 5, 7, 0, 0, 366, 372,
		1, 0, 0, 0, 367, 368, 5, 6, 0, 0, 368, 369, 3, 72, 36, 0, 369, 370, 5,
		7, 0, 0, 370, 372, 1, 0, 0, 0, 371, 363, 1, 0, 0, 0, 371, 367, 1, 0, 0,
		0, 372, 43, 1, 0, 0, 0, 373, 376, 3, 46, 23, 0, 374, 376, 3, 50, 25, 0,
		375, 373, 1, 0, 0, 0, 375, 374, 1, 0, 0, 0, 376, 45, 1, 0, 0, 0, 377, 378,
		5, 6, 0, 0, 378, 379, 3, 48, 24, 0, 379, 380, 5, 7, 0, 0, 380, 47, 1, 0,
		0, 0, 381, 382, 6, 24, -1, 0, 382, 394, 3, 46, 23, 0, 383, 384, 3, 74,
		37, 0, 384, 385, 5, 1, 0, 0, 385, 390, 3, 74, 37, 0, 386, 387, 5, 1, 0,
		0, 387, 389, 3, 74, 37, 0, 388, 386, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0,
		390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392,
		390, 1, 0, 0, 0, 393, 381, 1, 0, 0, 0, 393, 383, 1, 0, 0, 0, 394, 400,
		1, 0, 0, 0, 395, 396, 10, 3, 0, 0, 396, 397, 5, 1, 0, 0, 397, 399, 3, 46,
		23, 0, 398, 395, 1, 0, 0, 0, 399, 402, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0,
		400, 401, 1, 0, 0, 0, 401, 49, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 403, 404,
		3, 42, 21, 0, 404, 405, 5, 3, 0, 0, 405, 406, 5, 57, 0, 0, 406, 407, 5,
		8, 0, 0, 407, 408, 3, 50, 25, 0, 408, 409, 5, 1, 0, 0, 409, 410, 5, 55,
		0, 0, 410, 411, 5, 8, 0, 0, 411, 412, 5, 69, 0, 0, 412, 413, 5, 2, 0, 0,
		413, 426, 1, 0, 0, 0, 414, 415, 3, 42, 21, 0, 415, 416, 5, 3, 0, 0, 416,
		417, 5, 57, 0, 0, 417, 418, 5, 8, 0, 0, 418, 419, 3, 74, 37, 0, 419, 420,
		5, 1, 0, 0, 420, 421, 5, 55, 0, 0, 421, 422, 5, 8, 0, 0, 422, 423, 5, 69,
		0, 0, 423, 424, 5, 2, 0, 0, 424, 426, 1, 0, 0, 0, 425, 403, 1, 0, 0, 0,
		425, 414, 1, 0, 0, 0, 426, 51, 1, 0, 0, 0, 427, 428, 5, 73, 0, 0, 428,
		429, 5, 6, 0, 0, 429, 430, 5, 69, 0, 0, 430, 431, 5, 7, 0, 0, 431, 432,
		5, 6, 0, 0, 432, 433, 5, 69, 0, 0, 433, 439, 5, 7, 0, 0, 434, 435, 5, 6,
		0, 0, 435, 436, 5, 69, 0, 0, 436, 438, 5, 7, 0, 0, 437, 434, 1, 0, 0, 0,
		438, 441, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440,
		442, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 442, 443, 5, 16, 0, 0, 443, 444,
		3, 74, 37, 0, 444, 53, 1, 0, 0, 0, 445, 446, 5, 56, 0, 0, 446, 447, 5,
		73, 0, 0, 447, 451, 5, 3, 0, 0, 448, 450, 3, 58, 29, 0, 449, 448, 1, 0,
		0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0,
		452, 454, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 5, 2, 0, 0, 455,
		456, 5, 4, 0, 0, 456, 457, 3, 2, 1, 0, 457, 458, 5, 5, 0, 0, 458, 55, 1,
		0, 0, 0, 459, 460, 5, 56, 0, 0, 460, 461, 5, 73, 0, 0, 461, 465, 5, 3,
		0, 0, 462, 464, 3, 58, 29, 0, 463, 462, 1, 0, 0, 0, 464, 467, 1, 0, 0,
		0, 465, 463, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 468, 1, 0, 0, 0, 467,
		465, 1, 0, 0, 0, 468, 469, 5, 2, 0, 0, 469, 470, 5, 19, 0, 0, 470, 471,
		5, 26, 0, 0, 471, 472, 3, 72, 36, 0, 472, 473, 5, 4, 0, 0, 473, 474, 3,
		2, 1, 0, 474, 475, 5, 5, 0, 0, 475, 57, 1, 0, 0, 0, 476, 478, 7, 0, 0,
		0, 477, 476, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479,
		480, 5, 73, 0, 0, 480, 482, 5, 8, 0, 0, 481, 483, 5, 60, 0, 0, 482, 481,
		1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 486, 3, 72,
		36, 0, 485, 487, 5, 1, 0, 0, 486, 485, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0,
		487, 59, 1, 0, 0, 0, 488, 489, 5, 73, 0, 0, 489, 493, 5, 3, 0, 0, 490,
		492, 3, 62, 31, 0, 491, 490, 1, 0, 0, 0, 492, 495, 1, 0, 0, 0, 493, 491,
		1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 496, 1, 0, 0, 0, 495, 493, 1, 0,
		0, 0, 496, 497, 5, 2, 0, 0, 497, 61, 1, 0, 0, 0, 498, 499, 5, 73, 0, 0,
		499, 501, 5, 8, 0, 0, 500, 498, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501,
		503, 1, 0, 0, 0, 502, 504, 5, 13, 0, 0, 503, 502, 1, 0, 0, 0, 503, 504,
		1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 507, 3, 74, 37, 0, 506, 508, 5,
		1, 0, 0, 507, 506, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 63, 1, 0, 0,
		0, 509, 510, 7, 2, 0, 0, 510, 511, 5, 3, 0, 0, 511, 512, 3, 74, 37, 0,
		512, 513, 5, 2, 0, 0, 513, 525, 1, 0, 0, 0, 514, 515, 5, 64, 0, 0, 515,
		516, 5, 3, 0, 0, 516, 517, 3, 74, 37, 0, 517, 518, 5, 2, 0, 0, 518, 525,
		1, 0, 0, 0, 519, 520, 7, 3, 0, 0, 520, 521, 5, 3, 0, 0, 521, 522, 3, 74,
		37, 0, 522, 523, 5, 2, 0, 0, 523, 525, 1, 0, 0, 0, 524, 509, 1, 0, 0, 0,
		524, 514, 1, 0, 0, 0, 524, 519, 1, 0, 0, 0, 525, 65, 1, 0, 0, 0, 526, 527,
		5, 58, 0, 0, 527, 528, 5, 73, 0, 0, 528, 532, 5, 4, 0, 0, 529, 531, 3,
		68, 34, 0, 530, 529, 1, 0, 0, 0, 531, 534, 1, 0, 0, 0, 532, 530, 1, 0,
		0, 0, 532, 533, 1, 0, 0, 0, 533, 535, 1, 0, 0, 0, 534, 532, 1, 0, 0, 0,
		535, 536, 5, 5, 0, 0, 536, 67, 1, 0, 0, 0, 537, 538, 7, 1, 0, 0, 538, 541,
		5, 73, 0, 0, 539, 540, 5, 8, 0, 0, 540, 542, 3, 72, 36, 0, 541, 539, 1,
		0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 545, 1, 0, 0, 0, 543, 544, 5, 16, 0,
		0, 544, 546, 3, 74, 37, 0, 545, 543, 1, 0, 0, 0, 545, 546, 1, 0, 0, 0,
		546, 548, 1, 0, 0, 0, 547, 549, 5, 9, 0, 0, 548, 547, 1, 0, 0, 0, 548,
		549, 1, 0, 0, 0, 549, 555, 1, 0, 0, 0, 550, 552, 5, 59, 0, 0, 551, 550,
		1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553, 555, 3, 56,
		28, 0, 554, 537, 1, 0, 0, 0, 554, 551, 1, 0, 0, 0, 555, 69, 1, 0, 0, 0,
		556, 557, 7, 1, 0, 0, 557, 560, 5, 73, 0, 0, 558, 559, 5, 8, 0, 0, 559,
		561, 5, 73, 0, 0, 560, 558, 1, 0, 0, 0, 560, 561, 1, 0, 0, 0, 561, 562,
		1, 0, 0, 0, 562, 563, 5, 16, 0, 0, 563, 564, 5, 73, 0, 0, 564, 566, 5,
		3, 0, 0, 565, 567, 3, 62, 31, 0, 566, 565, 1, 0, 0, 0, 566, 567, 1, 0,
		0, 0, 567, 568, 1, 0, 0, 0, 568, 578, 5, 2, 0, 0, 569, 570, 7, 1, 0, 0,
		570, 573, 5, 73, 0, 0, 571, 572, 5, 8, 0, 0, 572, 574, 5, 73, 0, 0, 573,
		571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 575, 1, 0, 0, 0, 575, 576,
		5, 16, 0, 0, 576, 578, 5, 73, 0, 0, 577, 556, 1, 0, 0, 0, 577, 569, 1,
		0, 0, 0, 578, 71, 1, 0, 0, 0, 579, 589, 5, 63, 0, 0, 580, 589, 5, 64, 0,
		0, 581, 589, 5, 67, 0, 0, 582, 589, 5, 65, 0, 0, 583, 589, 5, 66, 0, 0,
		584, 585, 5, 6, 0, 0, 585, 586, 3, 72, 36, 0, 586, 587, 5, 7, 0, 0, 587,
		589, 1, 0, 0, 0, 588, 579, 1, 0, 0, 0, 588, 580, 1, 0, 0, 0, 588, 581,
		1, 0, 0, 0, 588, 582, 1, 0, 0, 0, 588, 583, 1, 0, 0, 0, 588, 584, 1, 0,
		0, 0, 589, 73, 1, 0, 0, 0, 590, 591, 6, 37, -1, 0, 591, 592, 7, 4, 0, 0,
		592, 633, 3, 74, 37, 20, 593, 633, 7, 5, 0, 0, 594, 633, 3, 64, 32, 0,
		595, 633, 3, 60, 30, 0, 596, 597, 5, 73, 0, 0, 597, 598, 5, 6, 0, 0, 598,
		599, 5, 69, 0, 0, 599, 600, 5, 7, 0, 0, 600, 601, 5, 6, 0, 0, 601, 602,
		5, 69, 0, 0, 602, 608, 5, 7, 0, 0, 603, 604, 5, 6, 0, 0, 604, 605, 5, 69,
		0, 0, 605, 607, 5, 7, 0, 0, 606, 603, 1, 0, 0, 0, 607, 610, 1, 0, 0, 0,
		608, 606, 1, 0, 0, 0, 608, 609, 1, 0, 0, 0, 609, 633, 1, 0, 0, 0, 610,
		608, 1, 0, 0, 0, 611, 612, 5, 73, 0, 0, 612, 613, 5, 6, 0, 0, 613, 614,
		3, 74, 37, 0, 614, 615, 5, 7, 0, 0, 615, 633, 1, 0, 0, 0, 616, 617, 5,
		73, 0, 0, 617, 618, 5, 11, 0, 0, 618, 633, 5, 54, 0, 0, 619, 620, 5, 73,
		0, 0, 620, 621, 5, 11, 0, 0, 621, 633, 5, 55, 0, 0, 622, 633, 5, 73, 0,
		0, 623, 633, 5, 49, 0, 0, 624, 633, 5, 70, 0, 0, 625, 633, 5, 69, 0, 0,
		626, 633, 5, 72, 0, 0, 627, 633, 5, 71, 0, 0, 628, 629, 5, 3, 0, 0, 629,
		630, 3, 74, 37, 0, 630, 631, 5, 2, 0, 0, 631, 633, 1, 0, 0, 0, 632, 590,
		1, 0, 0, 0, 632, 593, 1, 0, 0, 0, 632, 594, 1, 0, 0, 0, 632, 595, 1, 0,
		0, 0, 632, 596, 1, 0, 0, 0, 632, 611, 1, 0, 0, 0, 632, 616, 1, 0, 0, 0,
		632, 619, 1, 0, 0, 0, 632, 622, 1, 0, 0, 0, 632, 623, 1, 0, 0, 0, 632,
		624, 1, 0, 0, 0, 632, 625, 1, 0, 0, 0, 632, 626, 1, 0, 0, 0, 632, 627,
		1, 0, 0, 0, 632, 628, 1, 0, 0, 0, 633, 651, 1, 0, 0, 0, 634, 635, 10, 19,
		0, 0, 635, 636, 7, 6, 0, 0, 636, 650, 3, 74, 37, 20, 637, 638, 10, 18,
		0, 0, 638, 639, 7, 7, 0, 0, 639, 650, 3, 74, 37, 19, 640, 641, 10, 17,
		0, 0, 641, 642, 5, 18, 0, 0, 642, 650, 3, 74, 37, 18, 643, 644, 10, 16,
		0, 0, 644, 645, 7, 8, 0, 0, 645, 650, 3, 74, 37, 17, 646, 647, 10, 15,
		0, 0, 647, 648, 7, 9, 0, 0, 648, 650, 3, 74, 37, 16, 649, 634, 1, 0, 0,
		0, 649, 637, 1, 0, 0, 0, 649, 640, 1, 0, 0, 0, 649, 643, 1, 0, 0, 0, 649,
		646, 1, 0, 0, 0, 650, 653, 1, 0, 0, 0, 651, 649, 1, 0, 0, 0, 651, 652,
		1, 0, 0, 0, 652, 75, 1, 0, 0, 0, 653, 651, 1, 0, 0, 0, 63, 81, 86, 90,
		94, 98, 107, 111, 115, 119, 125, 129, 133, 135, 144, 149, 168, 181, 192,
		218, 226, 229, 238, 244, 265, 279, 281, 298, 306, 329, 352, 358, 371, 375,
		390, 393, 400, 425, 439, 451, 465, 477, 482, 486, 493, 500, 503, 507, 524,
		532, 541, 545, 548, 551, 554, 560, 566, 573, 577, 588, 608, 632, 649, 651,
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
	Tswift_GrammarParserINT           = 63
	Tswift_GrammarParserFLOAT         = 64
	Tswift_GrammarParserBOOL          = 65
	Tswift_GrammarParserCHARACTER     = 66
	Tswift_GrammarParserSTRING        = 67
	Tswift_GrammarParserENBLANCO      = 68
	Tswift_GrammarParserENTERO        = 69
	Tswift_GrammarParserDECIMAL       = 70
	Tswift_GrammarParserCARACTER      = 71
	Tswift_GrammarParserCADENA        = 72
	Tswift_GrammarParserID            = 73
	Tswift_GrammarParserUL_C          = 74
	Tswift_GrammarParserML_C          = 75
	Tswift_GrammarParserERROR         = 76
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
	Tswift_GrammarParserRULE_creacion_struct      = 35
	Tswift_GrammarParserRULE_tipo                 = 36
	Tswift_GrammarParserRULE_e                    = 37
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
		p.SetState(76)
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
	p.SetState(81)
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
				p.SetState(78)
				p.Sentencia()
			}

		}
		p.SetState(83)
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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewS_PrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Print_sentencia()
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(85)
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
			p.SetState(88)
			p.Declaracion()
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(89)
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
			p.SetState(92)
			p.Constante()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(93)
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
			p.SetState(96)
			p.Asignacion()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(97)
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
			p.SetState(100)
			p.If_sentencia()
		}

	case 6:
		localctx = NewS_SwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.Switch_sentencia()
		}

	case 7:
		localctx = NewS_GuardContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.Guard_sentencia()
		}

	case 8:
		localctx = NewS_WhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.While_sentencia()
		}

	case 9:
		localctx = NewS_ForContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(104)
			p.For_sentencia()
		}

	case 10:
		localctx = NewS_TransicionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(105)
			p.Trans_sentencia()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(106)
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
			p.SetState(109)
			p.Dec_vector()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(110)
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
			p.SetState(113)
			p.Func_vector()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(114)
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
			p.SetState(117)
			p.Asig_vector()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(118)
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
			p.SetState(121)
			p.Declaracion_metodo()
		}

	case 15:
		localctx = NewS_Declaracion_FuncionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(122)
			p.Declaracion_funcion()
		}

	case 16:
		localctx = NewS_Llamada_FuncionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(123)
			p.Llamada_funciones()
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

	case 17:
		localctx = NewS_Declaracion_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(127)
			p.Declaracion_matrices()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(128)
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
			p.SetState(131)
			p.Asig_mat()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(132)
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
		p.SetState(137)
		p.Match(Tswift_GrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.e(0)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(140)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.e(0)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Tipo()
		}
		{
			p.SetState(155)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.e(0)
		}

	case 2:
		localctx = NewDeclaracion_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.e(0)
		}

	case 3:
		localctx = NewDeclaracion_Tipo_noValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(Tswift_GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Tipo()
		}
		{
			p.SetState(166)
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
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstante_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(Tswift_GrammarParserLET)
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
		localctx = NewConstante_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(Tswift_GrammarParserLET)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSumAsgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(Tswift_GrammarParserMASIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.e(0)
		}

	case 2:
		localctx = NewResAsgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(Tswift_GrammarParserMENOSIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.e(0)
		}

	case 3:
		localctx = NewAsigContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
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
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIf_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.e(0)
		}
		{
			p.SetState(196)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.L_sentencias()
		}
		{
			p.SetState(198)
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
			p.SetState(200)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.e(0)
		}
		{
			p.SetState(202)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.L_sentencias()
		}
		{
			p.SetState(204)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(Tswift_GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.L_sentencias()
		}
		{
			p.SetState(208)
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
			p.SetState(210)
			p.Match(Tswift_GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.e(0)
		}
		{
			p.SetState(212)
			p.Match(Tswift_GrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.L_sentencias()
		}
		{
			p.SetState(214)
			p.Match(Tswift_GrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(Tswift_GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
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
		p.SetState(220)
		p.Match(Tswift_GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.e(0)
	}
	{
		p.SetState(222)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Tswift_GrammarParserCASE {
		{
			p.SetState(223)
			p.L_casos()
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDEFAULT {
		{
			p.SetState(228)
			p.L_default()
		}

	}
	{
		p.SetState(231)
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
		p.SetState(233)
		p.Match(Tswift_GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.e(0)
	}
	{
		p.SetState(235)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.L_sentencias()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserBREAK {
		{
			p.SetState(237)
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
		p.SetState(240)
		p.Match(Tswift_GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.L_sentencias()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserBREAK {
		{
			p.SetState(243)
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
		p.SetState(246)
		p.Match(Tswift_GrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.e(0)
	}
	{
		p.SetState(248)
		p.Match(Tswift_GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.L_sentencias()
	}
	{
		p.SetState(251)
		p.Trans_sentencia()
	}
	{
		p.SetState(252)
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
		p.SetState(254)
		p.Match(Tswift_GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.e(0)
	}
	{
		p.SetState(256)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.L_sentencias()
	}
	{
		p.SetState(258)
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
		p.SetState(260)
		p.Match(Tswift_GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)

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
		p.SetState(262)
		p.Match(Tswift_GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(263)
			p.Rango_p()
		}

	case 2:
		{
			p.SetState(264)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(267)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.L_sentencias()
	}
	{
		p.SetState(269)
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
		p.SetState(271)
		p.e(0)
	}
	{
		p.SetState(272)
		p.Match(Tswift_GrammarParserRANGO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
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
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserBREAK:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
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
			p.SetState(276)
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
			p.SetState(277)
			p.Match(Tswift_GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(278)
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
		p.SetState(283)

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
		p.SetState(284)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Tipo()
	}
	{
		p.SetState(288)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
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

	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.e(0)
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarParserT__0 {
			{
				p.SetState(294)
				p.Match(Tswift_GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(295)
				p.e(0)
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(301)
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
			p.SetState(303)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
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
			p.SetState(305)
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
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsig_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.e(0)
		}
		{
			p.SetState(311)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.e(0)
		}

	case 2:
		localctx = NewSumAsg_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(317)
			p.e(0)
		}
		{
			p.SetState(318)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(Tswift_GrammarParserMASIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.e(0)
		}

	case 3:
		localctx = NewResAsg_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.e(0)
		}
		{
			p.SetState(325)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(Tswift_GrammarParserMENOSIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
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
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunc_Vector_AppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(Tswift_GrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.e(0)
		}
		{
			p.SetState(336)
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
			p.SetState(338)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(Tswift_GrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
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
			p.SetState(343)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(Tswift_GrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Match(Tswift_GrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.e(0)
		}
		{
			p.SetState(350)
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
		p.SetState(354)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(355)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDOSPT {
		{
			p.SetState(356)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Tipo_matriz()
		}

	}
	{
		p.SetState(360)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
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
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTipo_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Tipo_matriz()
		}
		{
			p.SetState(365)
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
			p.SetState(367)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Tipo()
		}
		{
			p.SetState(369)
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
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.Listaval_mat()
		}

	case 2:
		localctx = NewDef_Matriz_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
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
		p.SetState(377)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.listaval_mat2(0)
	}
	{
		p.SetState(379)
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
	p.SetState(393)
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
			p.SetState(382)
			p.Listaval_mat()
		}

	case Tswift_GrammarParserPARIZQ, Tswift_GrammarParserMENOS, Tswift_GrammarParserNOT, Tswift_GrammarParserTRUE, Tswift_GrammarParserFALSE, Tswift_GrammarParserNIL, Tswift_GrammarParserATOI, Tswift_GrammarParserIOTA, Tswift_GrammarParserINT, Tswift_GrammarParserFLOAT, Tswift_GrammarParserSTRING, Tswift_GrammarParserENTERO, Tswift_GrammarParserDECIMAL, Tswift_GrammarParserCARACTER, Tswift_GrammarParserCADENA, Tswift_GrammarParserID:
		localctx = NewListaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(383)
			p.e(0)
		}
		{
			p.SetState(384)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.e(0)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(386)
					p.Match(Tswift_GrammarParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(387)
					p.e(0)
				}

			}
			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
			p.SetState(395)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(396)
				p.Match(Tswift_GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(397)
				p.Listaval_mat()
			}

		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimple_MatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)
			p.Tipo_matriz()
		}
		{
			p.SetState(404)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Match(Tswift_GrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Simple_mat()
		}
		{
			p.SetState(408)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
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
			p.SetState(414)
			p.Tipo_matriz()
		}
		{
			p.SetState(415)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(Tswift_GrammarParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.e(0)
		}
		{
			p.SetState(419)
			p.Match(Tswift_GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(Tswift_GrammarParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
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
		p.SetState(427)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(428)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)
		p.Match(Tswift_GrammarParserENTERO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(430)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(Tswift_GrammarParserCORCHETEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(Tswift_GrammarParserENTERO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(Tswift_GrammarParserCORCHETEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserCORCHETEIZQ {
		{
			p.SetState(434)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(442)
		p.Match(Tswift_GrammarParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
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
		p.SetState(445)
		p.Match(Tswift_GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID {
		{
			p.SetState(448)
			p.L_parametros()
		}

		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(454)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.L_sentencias()
	}
	{
		p.SetState(457)
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
		p.SetState(459)
		p.Match(Tswift_GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(460)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarParserGUIONBAJO || _la == Tswift_GrammarParserID {
		{
			p.SetState(462)
			p.L_parametros()
		}

		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(468)
		p.Match(Tswift_GrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(Tswift_GrammarParserMENOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.Match(Tswift_GrammarParserMAYOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(471)
		p.Tipo()
	}
	{
		p.SetState(472)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
		p.L_sentencias()
	}
	{
		p.SetState(474)
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
	p.SetState(477)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(476)

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
		p.SetState(479)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(Tswift_GrammarParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserINOUT {
		{
			p.SetState(481)
			p.Match(Tswift_GrammarParserINOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(484)
		p.Tipo()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(485)
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
		p.SetState(488)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Match(Tswift_GrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305280006646390776) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1001) != 0) {
		{
			p.SetState(490)
			p.L_argumentos()
		}

		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(496)
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
	p.SetState(500)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
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

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserDIR {
		{
			p.SetState(502)
			p.Match(Tswift_GrammarParserDIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(505)
		p.e(0)
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarParserT__0 {
		{
			p.SetState(506)
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

	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserATOI, Tswift_GrammarParserINT:
		localctx = NewFunc_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserATOI || _la == Tswift_GrammarParserINT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(510)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.e(0)
		}
		{
			p.SetState(512)
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
			p.SetState(514)
			p.Match(Tswift_GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.e(0)
		}
		{
			p.SetState(517)
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
			p.SetState(519)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserIOTA || _la == Tswift_GrammarParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(520)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.e(0)
		}
		{
			p.SetState(522)
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
		p.SetState(526)
		p.Match(Tswift_GrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Match(Tswift_GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(Tswift_GrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&648518359226253312) != 0 {
		{
			p.SetState(529)
			p.L_sentencias_struct()
		}

		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(535)
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
}

func NewL_AtributosContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_AtributosContext {
	var p = new(L_AtributosContext)

	InitEmptyL_sentencias_structContext(&p.L_sentencias_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentencias_structContext))

	return p
}

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

	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserVAR, Tswift_GrammarParserLET:
		localctx = NewL_AtributosContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(538)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserDOSPT {
			{
				p.SetState(539)
				p.Match(Tswift_GrammarParserDOSPT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(540)
				p.Tipo()
			}

		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserIGUAL {
			{
				p.SetState(543)
				p.Match(Tswift_GrammarParserIGUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(544)
				p.e(0)
			}

		}
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserPTCOMA {
			{
				p.SetState(547)
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
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserMUTATING {
			{
				p.SetState(550)
				p.Match(Tswift_GrammarParserMUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(553)
			p.Declaracion_funcion()
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

// ICreacion_structContext is an interface to support dynamic dispatch.
type ICreacion_structContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCreacion_structContext differentiates from other interfaces.
	IsCreacion_structContext()
}

type Creacion_structContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreacion_structContext() *Creacion_structContext {
	var p = new(Creacion_structContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_creacion_struct
	return p
}

func InitEmptyCreacion_structContext(p *Creacion_structContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_creacion_struct
}

func (*Creacion_structContext) IsCreacion_structContext() {}

func NewCreacion_structContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Creacion_structContext {
	var p = new(Creacion_structContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_creacion_struct

	return p
}

func (s *Creacion_structContext) GetParser() antlr.Parser { return s.parser }

func (s *Creacion_structContext) CopyAll(ctx *Creacion_structContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Creacion_structContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Creacion_structContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Creacion_Struct_SimpleContext struct {
	Creacion_structContext
}

func NewCreacion_Struct_SimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Creacion_Struct_SimpleContext {
	var p = new(Creacion_Struct_SimpleContext)

	InitEmptyCreacion_structContext(&p.Creacion_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Creacion_structContext))

	return p
}

func (s *Creacion_Struct_SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Creacion_Struct_SimpleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserID)
}

func (s *Creacion_Struct_SimpleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, i)
}

func (s *Creacion_Struct_SimpleContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Creacion_Struct_SimpleContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Creacion_Struct_SimpleContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Creacion_Struct_SimpleContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Creacion_Struct_SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterCreacion_Struct_Simple(s)
	}
}

func (s *Creacion_Struct_SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitCreacion_Struct_Simple(s)
	}
}

func (s *Creacion_Struct_SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitCreacion_Struct_Simple(s)

	default:
		return t.VisitChildren(s)
	}
}

type Creacion_StructContext struct {
	Creacion_structContext
}

func NewCreacion_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Creacion_StructContext {
	var p = new(Creacion_StructContext)

	InitEmptyCreacion_structContext(&p.Creacion_structContext)
	p.parser = parser
	p.CopyAll(ctx.(*Creacion_structContext))

	return p
}

func (s *Creacion_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Creacion_StructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarParserID)
}

func (s *Creacion_StructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, i)
}

func (s *Creacion_StructContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUAL, 0)
}

func (s *Creacion_StructContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *Creacion_StructContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *Creacion_StructContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserVAR, 0)
}

func (s *Creacion_StructContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserLET, 0)
}

func (s *Creacion_StructContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDOSPT, 0)
}

func (s *Creacion_StructContext) L_argumentos() IL_argumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_argumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_argumentosContext)
}

func (s *Creacion_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterCreacion_Struct(s)
	}
}

func (s *Creacion_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitCreacion_Struct(s)
	}
}

func (s *Creacion_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarVisitor:
		return t.VisitCreacion_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarParser) Creacion_struct() (localctx ICreacion_structContext) {
	localctx = NewCreacion_structContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Tswift_GrammarParserRULE_creacion_struct)
	var _la int

	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCreacion_StructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(556)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
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
				p.Match(Tswift_GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
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
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305280006646390776) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1001) != 0) {
			{
				p.SetState(565)
				p.L_argumentos()
			}

		}
		{
			p.SetState(568)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewCreacion_Struct_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(569)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarParserVAR || _la == Tswift_GrammarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(570)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarParserDOSPT {
			{
				p.SetState(571)
				p.Match(Tswift_GrammarParserDOSPT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(572)
				p.Match(Tswift_GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(575)
			p.Match(Tswift_GrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
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
	p.EnterRule(localctx, 72, Tswift_GrammarParserRULE_tipo)
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserINT:
		localctx = NewTipo_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(579)
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
			p.SetState(580)
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
			p.SetState(581)
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
			p.SetState(582)
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
			p.SetState(583)
			p.Match(Tswift_GrammarParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarParserCORCHETEIZQ:
		localctx = NewTipo_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(584)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Tipo()
		}
		{
			p.SetState(586)
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
	_startState := 74
	p.EnterRecursionRule(localctx, 74, Tswift_GrammarParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(632)
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
			p.SetState(591)

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
			p.SetState(592)
			p.e(20)
		}

	case 2:
		localctx = NewExpr_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(593)

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
		localctx = NewExpr_Funciones_EmbebidasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(594)
			p.Funciones_embebidas()
		}

	case 4:
		localctx = NewExpr_Llamada_FuncionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(595)
			p.Llamada_funciones()
		}

	case 5:
		localctx = NewExpr_MatrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(596)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(608)
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
					p.SetState(603)
					p.Match(Tswift_GrammarParserCORCHETEIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(604)
					p.Match(Tswift_GrammarParserENTERO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(605)
					p.Match(Tswift_GrammarParserCORCHETEDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(610)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpr_VectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(611)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)
			p.Match(Tswift_GrammarParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(613)
			p.e(0)
		}
		{
			p.SetState(614)
			p.Match(Tswift_GrammarParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFunc_Vector_isEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(616)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(617)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)
			p.Match(Tswift_GrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewFunc_Vector_CountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(619)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Match(Tswift_GrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(Tswift_GrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewExpr_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(622)
			p.Match(Tswift_GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExpr_NilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(623)
			p.Match(Tswift_GrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewExpr_DecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(624)
			p.Match(Tswift_GrammarParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewExpr_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(625)
			p.Match(Tswift_GrammarParserENTERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewExpr_CadenaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(626)
			p.Match(Tswift_GrammarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewExpr_CaracterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(627)
			p.Match(Tswift_GrammarParserCARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewExpr_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(628)
			p.Match(Tswift_GrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(629)
			p.e(0)
		}
		{
			p.SetState(630)
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
	p.SetState(651)
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
			p.SetState(649)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_MulDivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(634)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(635)

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
					p.SetState(636)
					p.e(20)
				}

			case 2:
				localctx = NewExpr_SumResContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(637)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(638)

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
					p.SetState(639)
					p.e(19)
				}

			case 3:
				localctx = NewExpr_ModContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(640)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(641)
					p.Match(Tswift_GrammarParserMOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(642)
					p.e(18)
				}

			case 4:
				localctx = NewExpr_RelContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(643)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(644)

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
					p.SetState(645)
					p.e(17)
				}

			case 5:
				localctx = NewExpr_LogicaContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_e)
				p.SetState(646)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(647)

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
					p.SetState(648)
					p.e(16)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(653)
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

	case 37:
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
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
