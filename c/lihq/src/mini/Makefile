all: main 

main: file1.o file2.o
	gcc -o main file1.o file2.o

file1.o: main.c
	gcc -c main.c -o file1.o

file2.o: max.c
	gcc -c max.c -o file2.o

clean:
	rm -f main file1.o file2.o
