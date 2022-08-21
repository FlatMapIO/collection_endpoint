grammar Base;


STRING
   : '\'' (ESC | SAFECODEPOINT)* '\''
   ;

fragment ESC
   : '\\' (['\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
   : [0-9a-fA-F]
   ;
fragment SAFECODEPOINT
   : ~ ['\\\u0000-\u001F]
   ;

WORD: [a-zA-Z_][a-zA-Z0-9_]*;


NUMBER
   : INT | '-'? INT ('.' [0-9] +)? EXP?
   ;


fragment INT
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   : [Ee] [+\-]? INT
   ;


WS
   : [ \t\n\r] + -> skip
   ;