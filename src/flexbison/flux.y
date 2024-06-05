%{
#include <stdio.h>

int yylex(void);
void yyerror(char *s);
extern FILE *yyin;
%}

%token PLUS MINUS STAR SLASH
%token EQ LT GT AND OR NOT CONCAT
%token LPAREN RPAREN SEMICOLON LBRACE RBRACE
%token ASSIGN ARROW COMMA
%token IDENT NUMBER STRING
%token IF ELSE FOR RETURN PRINTLN

%%

block:
    | statement SEMICOLON block

statement:
    IDENT ASSIGN bool_expression
    | IDENT ASSIGN LBRACE block RBRACE
    | IDENT ARROW ident_list LBRACE block RBRACE
    | IDENT LPAREN arg_list RPAREN
    | if_block
    | loop_block
    | RETURN bool_expression
    | RETURN
    | PRINTLN bool_expression
    ;

arg_list:
    | bool_expression
    | bool_expression COMMA arg_list_2

arg_list_2: bool_expression
    | bool_expression COMMA arg_list_2


ident_list:
    | IDENT
    | IDENT COMMA ident_list_2

ident_list_2: IDENT
    | IDENT COMMA ident_list_2


loop_block:
    FOR statement_with_empty SEMICOLON
    bool_expression_with_empty SEMICOLON
    statement_with_empty
    LBRACE block RBRACE

if_block:
    IF bool_expression LBRACE block RBRACE else_block;

else_block:
    | ELSE LBRACE block RBRACE
    | ELSE if_block
    ;

bool_expression_with_empty:
    | bool_expression
    ;

statement_with_empty:
    | statement
    ;

bool_expression: bool_term
    | bool_expression OR bool_term
    ;

bool_term: rel_expression
    | bool_term AND rel_expression
    ;

rel_expression: expression
    | expression EQ expression
    | expression LT expression
    | expression GT expression
    ;

expression: term
    | expression PLUS term
    | expression MINUS term
    | expression CONCAT term
    ;

term: factor
    | term STAR factor
    | term SLASH factor
    ;

factor: NUMBER
    | IDENT
    | IDENT LPAREN arg_list RPAREN
    | STRING
    | LPAREN bool_expression RPAREN
    | MINUS factor
    | NOT factor
    ;


%%

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <input_file>\n", argv[0]);
        return 1;
    }

    FILE *input_file = fopen(argv[1], "r");
    if (!input_file) {
        fprintf(stderr, "Error: could not open file %s\n", argv[1]);
        return 1;
    }

    yyin = input_file;

    int result = yyparse();

    fclose(input_file);

    if (result == 0) {
        printf("Valid syntax\n");
    }

    return 0;
}

void yyerror(char *s) {
    fprintf(stderr, "error: %s\n", s);
}
