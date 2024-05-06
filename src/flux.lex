%{
# include "flux.tab.h"
%}

WS     [ \t\n]
IDENT  [a-zA-Z_][a-zA-Z0-9_]*
NUMBER [0-9]+
STRING \".*\"

%%

{WS}

if    { return IF; }
else  { return ELSE; }
for   { return FOR; }
return { return RETURN; }
break { return BREAK; }
continue { return CONTINUE; }

true  { return TRUE; }
false { return FALSE; }

int   { return INT; }
str   { return STR; }
bool  { return BOOL; }
void  { return VOID; }

"+"  { return PLUS; }
"-"  { return MINUS; }
"*"  { return STAR; }
"/"  { return SLASH; }
"="  { return ASSIGN; }
"("  { return LPAREN; }
")"  { return RPAREN; }
"{"  { return LBRACE; }
"}"  { return RBRACE; }
";"  { return SEMICOLON; }
"==" { return EQ; }
"!=" { return NE; }
"<"  { return LT; }
"<=" { return LE; }
">"  { return GT; }
">=" { return GE; }
"|>" { return PIPE; }
"->" { return ARROW; }
","  { return COMMA; }

{IDENT} {
    return IDENT;
}

{NUMBER} {
    yylval = atoi(yytext);
    return NUMBER;
}

{STRING} {
    return STRING;
}

. { return yytext[0]; }

%%
