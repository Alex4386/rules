grammar JsonQuery;

query
   : NOT? SP? '(' SP? query SP? ')'                                                       #parenExp
   | query SP LOGICAL_OPERATOR SP query                                                   #logicalExp
   | attrPath SP 'pr'                                                                     #presentExp
   | attrPath SP op=( EQ | NE | GT | LT | GE | LE | CO | SW | EW | IN ) SP value          #compareExp
   | attrPath SP op=MT SP regexValue                                                      #regexExp
   | attrPath SP op=( EQ | NE | IN ) SP ipValue                                           #ipCompareExp
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
EQ : 'eq' | 'EQ' | 'equals' | 'EQUALS' | '==';
NE : 'ne' | 'NE' | 'noteq' | 'NOTEQ' | '!=';
GT : 'gt' | 'GT' | '>';
LT : 'lt' | 'LT' | '<';
GE : 'ge' | 'GE' | 'gte' | 'GTE' | '>=';
LE : 'le' | 'LE' | 'lte' | 'LTE' | '<=';
CO : 'co' | 'CO' | 'contains';
SW : 'sw' | 'SW' | 'startsWith';
EW : 'ew' | 'EW' | 'endsWith';
MT : 'mt' | 'MT' | 'matches' | 'MATCHES' | '~=';

attrPath
   : ATTRNAME subAttr?
   ;

valueAttrPath
   : ATTRNAME valueSubAttr?
   ;

subAttr
   : '.' attrPath
   ;

valueSubAttr
   : '.' valueAttrPath
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
   | valueAttrPath     #variable
   ;

VERSION
   : INT '.' INT '.' INT
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

regexValue
    : REGEX
    | valueAttrPath
    ;

REGEX
   : ('/' (REGEX_ESC | ~ [/\\]*) '/' REGEX_FLAGS? REGEX_FLAGS? REGEX_FLAGS?)
   ;

fragment REGEX_ESC
   : ESC
   | '\\' ([\\/bBdDwWsS])
   ;

fragment REGEX_FLAGS
   : [igm]
   ;

ipValue
    : IP_ADDRESS
    | IP_CIDR
    ;

listIPs
    : '[' subListOfIPs
    ;

subListOfIPs
    : ipValue COMMA subListOfIPs
    | ipValue ']';

IP_ADDRESS
    :   IPv4
    |   IPv6
    ;

IP_CIDR
    :   IPv4 '/' INT
    |   IPv6 '/' INT
    ;

fragment IPv4
    :   OCTET '.' OCTET '.' OCTET '.' OCTET
    ;

fragment OCTET
    :   '25' [0-5]         // Match 250-255
    |   '2' [0-4] [0-9]       // Match 200-249
    |   '1' [0-9] [0-9]          // Match 100-199
    |   [1-9] [0-9]           // Match 10-99
    |   [0-9]                 // Match 0-9
    ;

fragment IPv6
    :   (HEX_QUARTET ':') (HEX_QUARTET ':') (HEX_QUARTET ':') (HEX_QUARTET ':') (HEX_QUARTET ':')  HEX_QUARTET (':' HEX_QUARTET)?  // Basic IPv6 (six quartets + one or two more)
    |   '::' (HEX_QUARTET ':')* HEX_QUARTET                    // Leading ::
    |   (HEX_QUARTET ':')* '::' (HEX_QUARTET ':')* HEX_QUARTET // Embedded ::
    ;

fragment HEX_QUARTET
    : HEX HEX? HEX? HEX? // Matches 1 to 4 hexadecimal characters
    ;

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
