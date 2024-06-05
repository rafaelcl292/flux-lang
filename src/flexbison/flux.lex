%{
# include "flux.tab.h"
%}

WS     [ \t\n]
IDENT  [a-zA-Z_][a-zA-Z0-9_]*
NUMBER [0-9]+
STRING \"[^\"]*\"
COMMENT "//"[^\n]*\n

%%

{WS}
{COMMENT}

if    { return IF; }
else  { return ELSE; }
for   { return FOR; }
return { return RETURN; }

and   { return AND; }
or    { return OR; }
not   { return NOT; }
println { return PRINTLN; }

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
"<"  { return LT; }
">"  { return GT; }
"->" { return ARROW; }
","  { return COMMA; }
".." { return CONCAT; }

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
