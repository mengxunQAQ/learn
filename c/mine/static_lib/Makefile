all: main

main: main.o
	gcc -o main main.o -L./lib -lcalc

main.o:
	gcc -c main.c -o main.o -I./lib/include

clean:
	rm -f main main.o
