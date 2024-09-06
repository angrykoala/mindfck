grammar mindfck;

statements: statement*;

statement: 
    declaration | assignment | print;

declaration: 
    BYTE identifier;

assignment: 
    identifier EQUALS expression;

expression:
    identifier | literal | expression operand expression;

print:
    PRINT expression;

operand:
    PLUS | MINUS | TIMES | DIVIDE | DEQUALS | GT | GE | LT | LE | AND | OR;

identifier: 
    IDENTIFIER;

literal:
    NUMBER;

WS: [ \n\t\r]+ -> channel(HIDDEN);
BYTE: 'byte';
PRINT: 'print';
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
