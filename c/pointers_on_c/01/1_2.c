/*
 * read line from std-input and print it to std-output with line-numbers
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void main() {
    int ch;
    int line = 1;
    int new = 1;
    
    while ((ch = getchar()) != EOF) {  // end of file, meaning end input
        if (new == 1) {
            printf("%d ", line);
            new = 0;
        }
        putchar(ch);

        if (ch == '\n') {
            line += 1; 
            new = 1;
        }
    }
}
