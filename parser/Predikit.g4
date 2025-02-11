grammar Predikit;

pk_toplevel: kids+=pk_toplevel_child+;

pk_toplevel_child:
    group=pk_group | test=pk_test | tool=pk_tool
    ;

pk_group: agg_fn=pk_group_agg
        PK_LCURLY
            group_children+=pk_group_child+
        PK_RCURLY
        ;

pk_group_child:
    test=pk_test |
    group=pk_group | // nested children
    actual_param=pk_actual_param // params for the group
    ;

pk_group_agg: PK_ALL | PK_ANY | PK_NONE;

pk_test :
    PK_RETRYING? PK_TEST  PK_NOT? testname=PK_ID PK_LCURLY
        aps+=pk_actual_param+
    PK_RCURLY
    ;

pk_tool: PK_TOOL tool_name=PK_ID PK_LCURLY
            pk_tool_child+
         PK_RCURLY
         ;


pk_tool_child:
    pk_actual_param |
    pk_tool_metaparam
    ;

pk_tool_metaparam:
    PK_DOLLAR toolname=PK_ID PK_LCURLY
        pk_tool_actual_param+
    PK_RCURLY
    ;

pk_actual_param:
    param_name=PK_ID PK_COLON param_value=pk_actual_param_value
    ;

pk_actual_param_value:
    vi=PK_INT | vs=PK_STRING_LIT | vb=pk_bool | vc=pk_conversion_fn
    ;

pk_tool_actual_param:
    param_name=PK_ID PK_COLON param_value=pk_tool_actual_param_value
    ;

pk_tool_actual_param_value:
    PK_INT | PK_STRING_LIT | pk_bool | pk_conversion_fn | PK_TYPENAME
    ;

pk_conversion_fn: PK_ID
        PK_FN_LIT
        ;

pk_bool:
    PK_TRUE | PK_FALSE
    ;

PK_ALL         : 'all';
PK_ANY         : 'any';
PK_NONE        : 'none';
PK_TEST        : 'test';
PK_NOT         : 'not';
PK_TOOL        : 'tool';
PK_RETRYING    : '@';
PK_COLON       : ':';
PK_DOLLAR      : '$';
PK_LCURLY      : '{';
PK_RCURLY      : '}';
PK_LPAREN      : '(';
PK_RPAREN      : ')';
PK_TRUE        : 'true';
PK_FALSE       : 'false';

PK_CMP_EQ      : '=';
PK_CMP_NEQ     : '!=';
PK_CMP_RE      : '=~';
PK_CMP_GT      : '>';
PK_CMP_GTE     : '>=';
PK_CMP_LT      : '<';
PK_CMP_LTE     : '<=';

PK_FN_LIT  :  '(' (.)*? ')';

PK_STRING_LIT  :  '"' (ESC|.)*? '"';
fragment ESC : '\\"' | '\\\\' ;

PK_ID          :       LOWER (UPPER | LOWER | DIGIT | '_' | '?' | '!' | '.')*;

// could probably just make a list of valid types here instead
PK_TYPENAME    :       UPPER (LOWER | UPPER)+;

PK_INT             :   DIGIT+;



fragment DIGIT  : '0' .. '9';
fragment LOWER : 'a' .. 'z';
fragment UPPER : 'A' .. 'Z';

LINE_COMMENT  : '//' .*? '\r'? '\n' -> skip ;
COMMENT       : '/*' .*? '*/'       -> skip ;

WS      :       [ \t\r\n]+ -> skip;
