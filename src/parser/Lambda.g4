grammar Lambda;

prog            
                :       (stat (NEWLINE*|EOF))+;

stat            
                :       ID '=' value ';'?                                       # assignmentStatement
                |       ID '=' KEY '(' ( value ( ',' value )* )? ')' ';'?       # functionCallStatement
                |       COMMENT                                                 # commentStatement
                ;     

obj
                :       '{' (NEWLINE)* pair ((NEWLINE)* pair)*  '}'
                |       '{' (NEWLINE*) '}'
                ;

pair
                :       KEY '=' value ';'? (NEWLINE*)? 
                ;

array
                :       '[' ( value ( ',' value )* )? ']'
                ;

value
                :       STRING
                |       NUMBER
                |       BOOLEAN
                |       obj
                |       array
                ;

STRING
                :       '"' .*? '"'
                ;

NUMBER
                :       '-'? INT ('.' [0-9] +)?
                ;

BOOLEAN
                :       'true'
                |       'false'
                ;

fragment INT
                :       '0' | [1-9] [0-9]*
                ;

ID              
                :       '[' .+? ']';

COMMENT         
                :       '#' ~[\r\n]*;

KEY
                :       [a-zA-Z0-9]+;

NEWLINE         :       '\r'? '\n';

WS
                :       [ \t] + -> skip
                ;