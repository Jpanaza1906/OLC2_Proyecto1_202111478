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
    PRINT PARIZQ e (',' e)* PARDER PTCOMA?#S_Consola
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

func_vector:
    ID PUNTO APPEND PARIZQ e PARDER #Func_Vector_Append
    |ID PUNTO REMOVELAST PARIZQ PARDER #Func_Vector_RemoveLast
    |ID PUNTO REMOVE PARIZQ AT DOSPT e PARDER #Func_Vector_Remove
    |ID PUNTO ISEMPTY #Func_Vector_isEmpty
    |ID PUNTO COUNT #Func_Vector_Count
    ;
//Tipos de datos-----------------------------------------------------------------
tipo:
    INT #Tipo_Int
    |FLOAT #Tipo_Float
    |STRING #Tipo_String
    |BOOL #Tipo_Bool
    |CHARACTER #Tipo_Character
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
    | NIL               # Expr_Nil
    | ID CORCHETEIZQ e CORCHETEDER # Expr_Vector
    | ID                # Expr_Id
    | DECIMAL           # Expr_Decimal
    | ENTERO            # Expr_Entero
    | CADENA            # Expr_Cadena
    | CARACTER          # Expr_Caracter
    | PARIZQ e PARDER   # Expr_Par
    ;
