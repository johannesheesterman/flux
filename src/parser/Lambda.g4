grammar Lambda;

prog            
                :       (stat (NEWLINE*|EOF))+;

stat            
                :       ID '=' value?                                         # assignmentStatement
                |       ID obj?                                             # variableDeclarationStatement
                |       COMMENT                                             # commentStatement
                ;     

obj
                :       '{' (NEWLINE)* pair ((NEWLINE)* pair)*  '}'
                |       '{' (NEWLINE*) '}'
                ;

pair
                :       MODIFIER? KEY '=' value ';'? (NEWLINE*)? 
                ;

array
                :       '[' value (',' value)* ']'
                |       '[' ']'
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

MODIFIER        
                :       'public' | 'private';

KEY
                :       [a-zA-Z0-9]+;

NEWLINE         :       '\r'? '\n';

WS
                :       [ \t] + -> skip
                ;