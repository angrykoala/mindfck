grammar mindfck;

statements: statement*;

statement:
	declaration
	| assignment
	| print
	| ifConditional
	| whileLoop;

declaration: BYTE identifier;

assignment: identifier EQUALS expression;

print: PRINT expression;

ifConditional: IF '(' expression ')' '{' block '}';

whileLoop: WHILE '(' expression ')' '{' block '}';

block: statement*;

expression:
	identifier
	| literal
	| expression operand expression;

operand:
	PLUS
	| MINUS
	| TIMES
	| DIVIDE
	| DEQUALS
	| GT
	| GE
	| LT
	| LE
	| AND
	| OR;

identifier: IDENTIFIER;

literal: NUMBER;

WS: [ \n\t\r]+ -> channel(HIDDEN);
BYTE: 'byte';
PRINT: 'print';
IF: 'if';
WHILE: 'while';
PLUS: '+';
MINUS: '-';
TIMES: '*';
DIVIDE: '/';
EQUALS: '=';
DEQUALS: '==';
AND: 'AND' | 'and';
OR: 'OR' | 'or';
GT: '>';
GE: '>=';
LT: '<';
LE: '<=';
IDENTIFIER: [a-zA-Z]+;
NUMBER: [0-9]+;
