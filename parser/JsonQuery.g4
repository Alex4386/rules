grammar JsonQuery;

query
   : NOT? SP? '(' SP? query SP? ')'                                                                         #parenExp
   | query SP LOGICAL_OPERATOR SP query                                                             #logicalExp
   | attrPath SP 'pr'                                                                               #presentExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) SP value       #compareExp
   | attrPath SP MT SP Regex                                                                 #regexExp
   | attrPath SP op=( IN | MT ) SP ipCidr                                                  #ipExp
   ;

NOT
   : 'not' | 'NOT'
   ;

LOGICAL_OPERATOR
   : 'and' | 'or'
   ;

BOOLEAN
   : 'true' | 'false'
   ;

NULL
   : 'null'
   ;

IN:  'IN' | 'in';
EQ : 'eq' | 'EQ' | 'equals' | '==';
NE : 'ne' | 'NE' | 'noteq' | '!=';
GT : 'gt' | 'GT' | '>';
LT : 'lt' | 'LT' | '<';
GE : 'ge' | 'gte' | 'GE' | 'GTE' | '>=';
LE : 'le' | 'lte' | 'LE' | 'LTE' | '<=';
CO : 'co' | 'CO';
SW : 'sw' | 'SW';
EW : 'ew' | 'EW';
MT : 'mt' | 'matches' | 'MT' | 'MATCHES' | '~=';

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | VERSION           #version
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   | listInts          #listOfInts
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   | IPV4              #ipv4
   | IPV6              #ipv6
   ;

ipCidr
   : IPV4Cidr
   | IPV6Cidr
   ;

VERSION
   : INT '.' INT '.' INT
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;


// list, sublist, and elements
listStrings
   : '[' subListOfStrings
   ;

subListOfStrings
   : STRING COMMA subListOfStrings
   | STRING ']';

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

listDoubles
   : '[' subListOfDoubles
   ;

subListOfDoubles
   : DOUBLE COMMA subListOfDoubles
   | DOUBLE ']';

listInts
   : '[' subListOfInts
   ;

subListOfInts
   : INT COMMA subListOfInts
   | INT ']';

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;
SP
   : ' ' NEWLINE*
   ;

// IPv4 and IPv6 are defined in RFC 3986
IPV4 : IPV4Octet '.' IPV4Octet '.' IPV4Octet '.' IPV4Octet ;
IPV6 : IPV6HexSequence ('::' IPV6HexSequence?)? | '::' (IPV6HexSequence?) ;

fragment IPV4Octet
   : '25' [0-5]
   | '2' [0-4] [0-9]
   | '1' [0-9] [0-9]
   | [1-9] [0-9]
   | [0-9]
   ;

fragment IPV4Mask
   : '3' [0-2]
   | [1-2] [0-9]
   | [0-9]
   ;

fragment IPV6HexSequence : IPV6HexPart (':' IPV6HexPart)* ;
fragment IPV6HexPart : HEX+ ;

fragment IPV6Mask
   : '12' [0-8]
   | '1' [0-1] [0-9]
   | [1-9] [0-9]
   | [0-9]
   ;

IPV4Cidr
   : IPV4 ('/' IPV4Mask)
   ;

IPV6Cidr
   : IPV6 ('/' IPV6Mask)
   ;

// Regular expression support
Regex
   : '/' RegexPattern '/' RegexFlags?
   ;

fragment RegexPattern
   : (RegexChar | RegexEscape)+
   ;

fragment RegexFlags
   : [gimsuy]+
   ;

fragment RegexChar
   : ~[/\\]+
   ;

fragment RegexEscape
   : '\\/'
   ;
