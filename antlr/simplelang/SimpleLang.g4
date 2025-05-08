grammar SimpleLang;

prog:   stat+ ;

stat:   ID '=' expr ';'
    |   'print' '(' expr ')' ';'
    ;

expr:   expr op=('*'|'/') expr   # MulDiv
    |   expr op=('+'|'-') expr   # AddSub
    |   INT                      # Int
    |   ID                       # Id
    |   '(' expr ')'             # Parens
    ;

ID  :   [a-zA-Z_][a-zA-Z_0-9]* ;
INT :   [0-9]+ ;
WS  :   [ \t\r\n]+ -> skip ;