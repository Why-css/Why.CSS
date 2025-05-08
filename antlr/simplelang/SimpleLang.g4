grammar SimpleLang;

program: statement+;

statement:
	ID '=' expression ';'
	| 'print' '(' expression ')' ';';

expression:
	expression op = ('*' | '/') expression		# MulDiv
	| expression op = ('+' | '-') expression	# AddSub
	| INT										# Int
	| ID										# Id
	| '(' expression ')'						# Parens;

ID: [$][a-zA-Z_][a-zA-Z_0-9]*;
INT: [0-9]+;
WS: [ \t\r\n]+ -> skip;