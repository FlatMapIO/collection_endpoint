grammar Filter;
import Base;

expr : filter ((condition) filter)*;

filter
   : field '[' operator ']' '=' value
   ;

field: WORD;
value
  : primaryValue
  | NULL
  | arrayValue
  ;


arrayValue
   : numberArrary
   | stringArray
   ;

numberArrary: '(' NUMBER (',' NUMBER)* ')';
stringArray: '(' STRING (',' STRING)* ')';

primaryValue
   : STRING
   | NUMBER
   | BOOL_TRUR
   | BOOL_FALSE
   ;
condition: OR | AND;
BOOL_TRUR: 'true' | 'y';
BOOL_FALSE: 'false' | 'n';


OR: '|';
AND: '&';
NULL
  : 'null'
  ;

operator
  :
  |'eq'
  |'ne'
  |'gt'
  |'ge'
  |'lt'
  |'le'
  |'in'
  |'nin'
  |'like'
  |'nlike'
  |'between'
  |'nbetween'
 ;