grammar Tswift_Grammar;
import Tswift_Lexer;
//Produccion inicial de donde salen las sentencias--------------------------------

s: l_sentencias #SLSentencias
    ;

//Lista de sentencias-------------------------------------------------------------
l_sentencias:
    sentencia* #L_Sentencia
;

//Sentencias---------------------------------------------------------------------
sentencia:
    print_sentencia PTCOMA? #S_Print
    |declaracion PTCOMA?#S_Declaracion
    |constante PTCOMA?#S_Constante
    |asignacion PTCOMA?#S_Asignacion
    |if_sentencia #S_If
    |switch_sentencia #S_Switch
    |guard_sentencia #S_Guard
    |while_sentencia #S_While
    |for_sentencia #S_For
    |trans_sentencia PTCOMA?#S_Transicion
    |dec_vector PTCOMA?#S_Declaracion_Vector
    |func_vector PTCOMA?#S_Funcion_Vector
    |asig_vector PTCOMA?#S_Asignacion_Vector
    |declaracion_metodo #S_Declaracion_Metodo
    |declaracion_funcion #S_Declaracion_Funcion
    |llamada_funciones PTCOMA?#S_Llamada_Funcion
    |declaracion_matrices PTCOMA?#S_Declaracion_Matriz
    |asig_mat PTCOMA?#S_Asignacion_Matriz
    ;

//Sentencia de impresion---------------------------------------------------------
print_sentencia:
    PRINT PARIZQ e (',' e)* PARDER PTCOMA?#Print
    ;

//Declaracion--------------------------------------------------------------------
declaracion:
    VAR ID DOSPT tipo IGUAL e #Declaracion_Tipo_Val
    |VAR ID IGUAL e #Declaracion_Val
    |VAR ID DOSPT tipo INTERROGACION #Declaracion_Tipo_noVal
    ;

//Constante----------------------------------------------------------------------
constante:
    LET ID DOSPT tipo IGUAL e #Constante_Tipo_Val
    |LET ID IGUAL e #Constante_Val
    ;

//Asignacion---------------------------------------------------------------------

asignacion:
    ID MASIGUAL e #SumAsg
    |ID MENOSIGUAL e #ResAsg
    |ID IGUAL  e #Asig
    ;

//Sentencia if-------------------------------------------------------------------

if_sentencia:
    IF e LLAVEIZQ l_sentencias LLAVEDER #If_Simple
    |IF e LLAVEIZQ l_sentencias LLAVEDER ELSE LLAVEIZQ l_sentencias LLAVEDER #If_Else
    |IF e LLAVEIZQ l_sentencias LLAVEDER ELSE if_sentencia #If_ElseIf
    ;

//Sentencia switch----------------------------------------------------------------

switch_sentencia:
    SWITCH e LLAVEIZQ l_casos+ l_default? LLAVEDER #Switch
    ;
l_casos:
    CASE e DOSPT l_sentencias BREAK? #Case
    ;
l_default:
    DEFAULT DOSPT l_sentencias BREAK? #Default
    ;

//Sentencia guard----------------------------------------------------------------

guard_sentencia:
    GUARD e ELSE LLAVEIZQ l_sentencias trans_sentencia LLAVEDER #Guard
    ;

//Ciclo While--------------------------------------------------------------------

while_sentencia:
    WHILE e LLAVEIZQ l_sentencias LLAVEDER #While
    ;

//Ciclo For----------------------------------------------------------------------

for_sentencia:
    FOR ID IN (rango_p|e) LLAVEIZQ l_sentencias LLAVEDER #For
    ;

rango_p:
    e RANGO e #Rango
    ;

//Sentencias de Transicion-------------------------------------------------------

trans_sentencia:
    BREAK #Break
    |CONTINUE #Continue
    |RETURN (e)? #Return
    ;

//Vectores ----------------------------------------------------------------------

dec_vector:
    tipod=(VAR|LET) ID DOSPT CORCHETEIZQ tipo CORCHETEDER IGUAL def_vector #Declaracion_Vector
    ;

def_vector:
    CORCHETEIZQ e (',' e)* CORCHETEDER #Def_Vector
    |CORCHETEIZQ CORCHETEDER #Def_Vector_Vacio
    |ID #Def_Vector_Id
    ;
asig_vector:
    ID CORCHETEIZQ e CORCHETEDER IGUAL e #Asig_Vector
    |ID CORCHETEIZQ e CORCHETEDER MASIGUAL e #SumAsg_Vector
    |ID CORCHETEIZQ e CORCHETEDER MENOSIGUAL e #ResAsg_Vector
    ;

func_vector:
    ID PUNTO APPEND PARIZQ e PARDER #Func_Vector_Append
    |ID PUNTO REMOVELAST PARIZQ PARDER #Func_Vector_RemoveLast
    |ID PUNTO REMOVE PARIZQ AT DOSPT e PARDER #Func_Vector_Remove
    ;

//Matrices ----------------------------------------------------------------------

declaracion_matrices:
    (VAR|LET) ID (DOSPT tipo_matriz)? IGUAL def_matriz #Declaracion_Matriz
    ;

tipo_matriz:
    CORCHETEIZQ tipo_matriz CORCHETEDER #Tipo_Matriz
    |CORCHETEIZQ tipo CORCHETEDER  #Tipo_Matriz_Simple
    ;

def_matriz:
    listaval_mat #Def_Matriz
    | simple_mat #Def_Matriz_Simple
    ;
listaval_mat:
    CORCHETEIZQ listaval_mat2 CORCHETEDER #ListaCompletaVal
    ;
listaval_mat2:
    listaval_mat2 ',' listaval_mat #ListaValoresHermanos
    | listaval_mat #ListaValorSig
    | e ',' e (',' e)* #ListaExpr
    ;

simple_mat:
    tipo_matriz PARIZQ REPEATING DOSPT simple_mat ',' COUNT DOSPT ENTERO PARDER #Simple_Mat
    | tipo_matriz PARIZQ REPEATING DOSPT e ',' COUNT DOSPT ENTERO PARDER #Simple_Mat_Expr
    ;

asig_mat:
    ID CORCHETEIZQ ENTERO CORCHETEDER CORCHETEIZQ ENTERO CORCHETEDER (CORCHETEIZQ ENTERO CORCHETEDER)* IGUAL e #Asig_Mat
    ;

//Funciones ---------------------------------------------------------------------
declaracion_metodo:
    FUNC ID PARIZQ l_parametros* PARDER LLAVEIZQ l_sentencias LLAVEDER #Declaracion_Metodo
    ;
declaracion_funcion:
    FUNC ID PARIZQ l_parametros* PARDER MENOS MAYOR tipo LLAVEIZQ l_sentencias trans_sentencia PTCOMA? LLAVEDER #Declaracion_Funcion 
    ;

l_parametros: prim=(ID | GUIONBAJO)? ID DOSPT INOUT? tipo ','?#L_Parametros
    ;

llamada_funciones:
    ID PARIZQ l_argumentos* PARDER #Llamada_Funcion
    ;

l_argumentos:
    (ID ':')? DIR? e ','? #L_Argumentos
    ;

//Structs -----------------------------------------------------------------------

def_struct:
    STRUCT ID LLAVEIZQ l_sentencias_struct* LLAVEDER #Def_Struct
    ;

l_sentencias_struct:
    (VAR|LET) ID (DOSPT tipo)? (IGUAL e)? PTCOMA? #L_Atributos
    | MUTATING? declaracion_funcion #L_Funciones
    ;

creacion_struct:
    (VAR|LET) ID (DOSPT ID)? IGUAL ID PARIZQ l_argumentos? PARDER #Creacion_Struct
    | (VAR|LET) ID (DOSPT ID)? IGUAL ID #Creacion_Struct_Simple
    ;


//Tipos de datos-----------------------------------------------------------------
tipo:
    INT #Tipo_Int
    |FLOAT #Tipo_Float
    |STRING #Tipo_String
    |BOOL #Tipo_Bool
    |CHARACTER #Tipo_Character
    |CORCHETEIZQ tipo CORCHETEDER #Tipo_Vector
    ;

//Expresiones--------------------------------------------------------------------
e
    : op=(MENOS | NOT) e      # Expr_Neg
    | e op=(POR | DIV) e    # Expr_MulDiv
    | e op=(MAS | MENOS) e  # Expr_SumRes
    | e MOD e           # Expr_Mod
    | e op=(AND | OR) e     # Expr_Logica
    | e op=(IGUALIGUAL | DIFERENTE | MAYORIGUAL | MAYOR | MENORIGUAL | MENOR) e     # Expr_Rel
    | op=(TRUE | FALSE)          # Expr_Booleano
    | llamada_funciones # Expr_Llamada_Funcion
    | ID CORCHETEIZQ ENTERO CORCHETEDER CORCHETEIZQ ENTERO CORCHETEDER (CORCHETEIZQ ENTERO CORCHETEDER)* # Expr_Matriz
    | ID CORCHETEIZQ e CORCHETEDER # Expr_Vector
    | ID PUNTO ISEMPTY #Func_Vector_isEmpty
    | ID PUNTO COUNT #Func_Vector_Count
    | ID                # Expr_Id
    | NIL               # Expr_Nil
    | DECIMAL           # Expr_Decimal
    | ENTERO            # Expr_Entero
    | CADENA            # Expr_Cadena
    | CARACTER          # Expr_Caracter
    | PARIZQ e PARDER   # Expr_Par
    ;
