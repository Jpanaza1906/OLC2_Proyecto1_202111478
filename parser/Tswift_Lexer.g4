lexer grammar Tswift_Lexer;

//******************* SIMBOLOS DEFINIDOS ***************************
//Simbolos
PARDER: ')';
PARIZQ: '(';
LLAVEIZQ: '{';
LLAVEDER: '}';
DOSPT: ':';
PTCOMA: ';';
INTERROGACION: '?';
//Operadores aritmeticos
MASIGUAL: '+=';
MENOSIGUAL: '-=';
IGUAL: '=';
DIV : '/';
MOD : '%';
MENOS : '-';
MAS: '+';
POR: '*';
//Operadores relacionales
IGUALIGUAL: '==';
DIFERENTE: '!=';
MAYORIGUAL: '>=';
MENORIGUAL: '<=';
MAYOR: '>';
MENOR: '<';
//Operadores logicos
AND: '&&';
OR: '||';
NOT: '!';

//******************** PALABRAS RESERVADAS ***************************
//Palabras del lenSguaje
PRINT: 'print';
VAR: 'var';
LET: 'let';
TRUE : 'true';
FALSE : 'false';
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
WHILE: 'while';
FOR: 'for';
IN: 'in';
RANGO: '...';
GUARD: 'guard';
CONTINUE: 'continue';
RETURN: 'return';
BREAK: 'break';
NIL: 'nil';
//Tipos de datos reservadas
INT: 'Int';
FLOAT: 'Float';
BOOL: 'Bool';
CHARACTER: 'Character';
STRING: 'String';


// ********************** REGLAS LEXICAS ***************************
//En blanco
ENBLANCO: [ \t\n\r]+ -> skip ;

fragment DIG: [0-9];

//tipos de datos
ENTERO: (DIG)+ ;
DECIMAL: DIG+ '.' DIG+ ;
CARACTER: '"' ( ~('\'' | '\\') | '\\' . ) '"' ;
CADENA: '"' ( ~('"' | '\\') | '\\' . )* '"' ;

ID
	: [a-zA-Z_][a-zA-Z0-9_]*|'_'[a-zA-Z0-9_]+
	;

//Comentarios
UL_C
  :  '//' ~('\r' | '\n')* -> channel(HIDDEN)
  ;
ML_C
  :  '/*' .*? '*/' -> channel(HIDDEN)
  ;