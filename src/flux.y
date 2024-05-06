%{
#include <stdio.h>

int yylex(void);
void yyerror(char *s);
extern FILE *yyin;
%}

%token PLUS MINUS STAR SLASH
%token EQ NE LT LE GT GE AND OR NOT
%token LPAREN RPAREN SEMICOLON LBRACE RBRACE
%token ASSIGN PIPE ARROW COMMA
%token IDENT NUMBER STRING
%token IF ELSE FOR RETURN BREAK CONTINUE
%token INT STR BOOL VOID
%token TRUE FALSE

%%

block:
    | statement SEMICOLON block

statement:
    IDENT ASSIGN bool_expression
    | IDENT ASSIGN LBRACE block_with_return RBRACE
    | IDENT bool_expression
    | IDENT ARROW IDENT LBRACE block_with_return RBRACE
    | if_block
    | loop_block
    ;


loop_block:
    FOR statement_with_empty SEMICOLON
    bool_expression_with_empty SEMICOLON
    statement_with_empty
    LBRACE block_with_break_continue RBRACE

block_with_return:
    | statement SEMICOLON block_with_return
    | RETURN bool_expression SEMICOLON block_with_return
    ;

block_with_break_continue:
    | statement SEMICOLON block_with_break_continue
    | BREAK SEMICOLON block_with_break_continue
    | CONTINUE SEMICOLON block_with_break_continue
    ;

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
    | expression NE expression
    | expression LT expression
    | expression LE expression
    | expression GT expression
    | expression GE expression
    ;

expression: term
    | expression PLUS term
    | expression MINUS term
    ;

term: factor
    | term STAR factor
    | term SLASH factor
    ;

factor: NUMBER
    | IDENT
    | STRING
    | TRUE | FALSE
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