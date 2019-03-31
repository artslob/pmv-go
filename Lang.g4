grammar Lang;

fragment LOWERCASE  : [a-z] ;
fragment UPPERCASE  : [A-Z] ;
fragment DIGIT      : [0-9] ;

EMPTY               : (' ' | '\t' | '\n' | '\r')+ -> skip;
IDENTIFIER          : (LOWERCASE | UPPERCASE | '_') (LOWERCASE | UPPERCASE | '_' | DIGIT)* ;
DEC                 : DIGIT+ ;

source : sourceItem* EOF;

sourceItem: funcDef;

funcDef: 'def' funcSignature 'end';

funcSignature: IDENTIFIER '(' (arg (',' arg)*)? ')' ('of' typeRef)?;

arg: IDENTIFIER ('of' typeRef)?;

typeRef: built                        # builtin
       | IDENTIFIER                   # custom
       | typeRef 'array' '[' DEC ']'  # array
       ;

built: 'bool'|'byte'|'int'|'uint'|'long'|'ulong'|'char'|'string';
