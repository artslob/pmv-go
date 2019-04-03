grammar Lang;

fragment LOWERCASE  : [a-z] ;
fragment UPPERCASE  : [A-Z] ;
fragment DIGIT      : [0-9] ;

EMPTY               : (' ' | '\t' | '\n' | '\r')+ -> skip;

IDENTIFIER          : (LOWERCASE | UPPERCASE | '_') (LOWERCASE | UPPERCASE | '_' | DIGIT)* ;
STR                 : '"' [^"\\]* (.[^"\\]*)* '"';
CHAR                : '\'' ~('\'') '\'' ;
HEX                 : '0'[xX][0-9A-Fa-f]+ ;
BITS                : '0'[bB][01]+ ;
DEC                 : DIGIT+ ;
BOOL                : 'true' | 'false' ;

source : sourceItem* EOF;

sourceItem: funcDef;

funcDef: 'def' funcSignature statement* 'end';

funcSignature: IDENTIFIER '(' (arg (',' arg)*)? ')' ('of' typeRef)?;

arg: IDENTIFIER ('of' typeRef)?;

typeRef: built                        # builtin
       | IDENTIFIER                   # custom
       | typeRef 'array' '[' DEC ']'  # array
       ;

built: 'bool'|'byte'|'int'|'uint'|'long'|'ulong'|'char'|'string';

statement: 'if' expr 'then' statement ('else' statement)? 'end'  # if
         | ('while'|'until') expr statement* 'end'               # loop
         | statement ('while'|'until') expr ';'                  # repeat
         | 'break' ';'                                           # break
         | expr ';'                                              # expression
         | ('begin'|'{') statement* ('end'|'}')                  # block
         ;

expr: expr BIN_OP expr                      # binary
    | UN_OP expr                            # unary
    | '(' expr ')'                          # braces
    | expr '(' (expr (',' expr)*)? ')'      # call
    | expr '[' (ranges (',' ranges)*)? ']'  # slice
    | (BOOL|STR|CHAR|HEX|BITS|DEC)          # literal
    | IDENTIFIER                            # place
    ;

ranges: expr ('..' expr)?;

BIN_OP: '<<=' | '>>='
      | '&=' | '|='
      | '^=' | '%='
      | '='
      | '==' | '!='
      | '*=' | '/='
      | '+=' | '-='
      | '<<' | '>>'
      | '&&' | '||'
      | '*' | '%'
      | '^' | '/'
      | '&' | '|'
      | '+' | '-'
      | '<=' | '>='
      | '<' | '>'
      ;

UN_OP: '+' | '-' | '~' | '!' ;
