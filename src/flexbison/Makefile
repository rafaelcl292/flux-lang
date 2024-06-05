flux_syntax: flux.lex flux.y
	bison -d flux.y
	lex flux.lex
	gcc -o $@ flux.tab.c lex.yy.c -lfl

clean:
	rm -f flux_syntax lex.yy.c flux.tab.c flux.tab.h
