flux: src/flux.lex src/flux.y
	bison -d src/flux.y
	lex src/flux.lex
	gcc -o $@ flux.tab.c lex.yy.c -lfl
	rm -f lex.yy.c
	rm -f flux.tab.c
	rm -f flux.tab.h

clean:
	rm -f flux
