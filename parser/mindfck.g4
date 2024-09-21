grammar mindfck;

statements: statement*;

statement:
	declaration
	| assignment
	| print
	| ifConditional
	| whileLoop
	| read;

declaration: type = (BYTE | INT) identifier;

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

literal: NUMBER | CHAR | BYTE_NUMBER;

CHAR: '\'' EXT_ASCII_CHAR '\'';
BYTE_NUMBER: NUMBER 'b';

fragment EXT_ASCII_CHAR:
	[\u0000-\u00FF]; // Matches characters in the range 0–255 (extended ASCII).

WS: [ \n\t\r]+ -> channel(HIDDEN);
BYTE: 'byte';
INT: 'int';
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
NUMBER: [0-9]+;
IDENTIFIER: [a-zA-Z]+;
