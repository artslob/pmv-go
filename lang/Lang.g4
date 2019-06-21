grammar Lang;

EMPTY       : (' ' | '\t' | '\n' | '\r')+ -> skip;

IDENTIFIER  : [a-zA-Z_]([a-zA-Z_0-9])* ;
STR         : '"' [^"\\]* (.[^"\\]*)* '"' ;
CHAR        : '\'' ~('\'') '\'' ;
HEX         : '0'[xX][0-9A-Fa-f]+ ;
BITS        : '0'[bB][01]+ ;
DEC         : [0-9]+ ;
BOOL        : 'true' | 'false' ;

source : sourceItem* EOF;

sourceItem: funcDef;

funcDef: 'def' funcSignature statement* 'end';

/******** FUNC SIGNATURE ********/

funcSignature: IDENTIFIER '(' (arg (',' arg)*)? ')' ('of' typeRef)?;

arg: IDENTIFIER ('of' typeRef)?;

typeRef: built                        # builtin
       | IDENTIFIER                   # custom
       | typeRef 'array' '[' DEC ']'  # array
       ;

built: 'bool'|'byte'|'int'|'uint'|'long'|'ulong'|'char'|'string';

/******** STATEMENT ********/

statement: ifExpr ifThen ifElse? 'end'          # if
         | whileExpr whileBody 'end'            # loop
         | untilStatement untilExpr ';'         # repeat
         | 'break' ';'                          # break
         | expr ';'                             # expression
         | ('begin'|'{') blockBody ('end'|'}')  # block
         ;

ifExpr: 'if' expr ;
ifThen: 'then' statement ;
ifElse: 'else' statement ;

whileExpr: ('while'|'until') expr ;
whileBody: statement* ;

untilStatement: 'do' statement ;
untilExpr: ('while'|'until') expr ;

blockBody: statement* ;

/******** EXPR ********/

expr: expr '(' (expr (',' expr)*)? ')'      # call
    | expr '[' (ranges (',' ranges)*)? ']'  # slice
    | UN_OP expr                            # unary
    | expr op=('*'|'/'|'%') expr            # mulDivMod
    | expr op=('+'|'-')     expr            # addSub
    | expr op=SHIFT_OP      expr            # shift
    | expr op=CMP_ARR_OP    expr            # compareArrow
    | expr op=('=='|'!=')   expr            # compareEqual
    | expr op='&'  expr                     # and
    | expr op='^'  expr                     # xor
    | expr op='|'  expr                     # or
    | expr op='&&' expr                     # andLogical
    | expr op='||' expr                     # orLogical
    | expr op='='  expr                     # assign
    // TODO: add assign and var creation
    | '(' expr ')'                          # braces
    | (BOOL|STR|CHAR|HEX|BITS|DEC)          # literal
    | IDENTIFIER                            # place
    ;

ranges: expr ('..' expr)?;

SHIFT_OP: '<<' | '>>' ;

CMP_ARR_OP: '<' | '<='
          | '>' | '>='
          ;

UN_OP: '+' | '-' | '~' | '!' ;
