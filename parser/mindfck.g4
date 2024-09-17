grammar mindfck;

statements: statement*;

statement:
	declaration
	| assignment
	| print
	| ifConditional
	| whileLoop;

declaration: type = (BYTE | INT) identifier;

assignment: identifier EQUALS expression;

print: PRINT expression;

ifConditional:
	IF '(' expression ')' '{' block '}' (ELSE '{' block '}')?;

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
INT: 'int';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
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
