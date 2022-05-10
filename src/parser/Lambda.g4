grammar Lambda;

prog            
                :       (stat (NEWLINE*|EOF))+;

stat            
                :       ID '=' value ';'?                                       # assignmentStatement
                |       ID '=' func                                             # functionCallAssignmentStatement
                |       func                                                    # functionCallStatement
                |       COMMENT                                                 # commentStatement
                ;     

func            :        KEY '(' ( value ( ',' value )* )? ')' ';'? 
                ;

obj
                :       '{' (NEWLINE)* pair ((NEWLINE)* pair)* (NEWLINE)*  '}'
                |       '{' (NEWLINE*) '}'
                ;

pair
                :       KEY '=' value ';'? 
                ;

array
                :       '[' ( value ( ',' value )* )? ']'
                ;

value
                :       STRING
                |       NUMBER
                |       BOOLEAN
                |       ID
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