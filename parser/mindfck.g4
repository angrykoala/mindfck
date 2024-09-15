grammar mindfck;

statements: statement*;

statement:
	declaration
	| assignment
	| print
	| ifConditional
	| whileLoop
	| read;

declaration: BYTE identifier;

assignment: identifier EQUALS expression;

print: PRINT expression;

ifConditional:
	IF '(' expression ')' '{' block '}' (ELSE '{' block '}')?;

whileLoop: WHILE '(' expression ')' '{' block '}';

read: READ identifier;

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

literal: NUMBER | CHAR;

WS: [ \n\t\r]+ -> channel(HIDDEN);
BYTE: 'byte';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
READ: 'read';
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

CHAR: '\'' EXT_ASCII_CHAR '\'';

fragment EXT_ASCII_CHAR:
	[\u0000-\u00FF] ; // Matches characters in the range 0–255 (extended ASCII).
