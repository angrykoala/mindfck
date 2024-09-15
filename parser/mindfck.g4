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
	| '(' expression ')'
	| NOT expression
	| expression op = (TIMES | DIVIDE | AND) expression
	| expression op = (PLUS | MINUS | OR) expression
	| expression op = (GT | GE | LT | LE | EQUALS | DEQUALS) expression;

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
AND: 'and';
OR: 'or';
NOT: 'not ' | '!';
GT: '>';
GE: '>=';
LT: '<';
LE: '<=';
IDENTIFIER: [a-zA-Z]+;
NUMBER: [0-9]+;
