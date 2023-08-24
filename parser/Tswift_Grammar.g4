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
    PRINT PARIZQ e (',' e)* PARDER #S_Consola
    |declaracion #S_Declaracion
    ;

//Declaracion--------------------------------------------------------------------

declaracion:
    VAR ID DOSPT tipo IGUAL e #Declaracion_Tipo_Val
    |VAR ID IGUAL e #Declaracion_Val
    |VAR ID DOSPT tipo INTERROGACION #Declaracion_Tipo_noVal
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
    : e op=(IGUALIGUAL | DIFERENTE | MAYORIGUAL | MAYOR | MENORIGUAL | MENOR) e     # Expr_Rel
    | e op=(AND | OR) e     # Expr_Logica
    | op=(MENOS | NOT) e      # Expr_Neg
    | e op=(POR | DIV) e    # Expr_MulDiv
    | e op=(MAS | MENOS) e  # Expr_SumRes
    | e MOD e           # Expr_Mod
    | BOOLEANO          # Expr_Booleano
    | NIL               # Expr_Nil
    | ID                # Expr_Id
    | DECIMAL           # Expr_Decimal
    | ENTERO            # Expr_Entero
    | CADENA            # Expr_Cadena
    | CARACTER          # Expr_Caracter
    | PARIZQ e PARDER   # Expr_Par
    ;
