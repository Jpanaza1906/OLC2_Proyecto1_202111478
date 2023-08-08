grammar Tswift_Grammar;
import Tswift_Lexer;

s: a;
// los labels no pueden ser igual que las producciones
a: PARIZQ a PARDER #A0
| ENTERO           #AENTERO
| DECIMAL          #ADECIMAL
| ID               #AID
|                  #AEPSILUM // espsium
;