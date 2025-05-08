grammar SimpleLang;

program: statement+;

statement:
	'@use' '"' IMPORTNAMES '"' ';'		# Import
	| ID '=' expression ';'				# Assignment
	| 'print' '(' expression ')' ';'	# Print;

expression:
	expression op = ('*' | '/') expression		# MulDiv
	| expression op = ('+' | '-') expression	# AddSub
	| INT										# Int
	| ID										# Id
	| '(' expression ')'						# Parens;

ID: [$][a-zA-Z_][a-zA-Z_0-9]*;
INT: [0-9]+;
IMPORTNAMES: [a-zA-Z_][a-zA-Z_0-9]*;
WS: [ \t\r\n]+ -> skip;