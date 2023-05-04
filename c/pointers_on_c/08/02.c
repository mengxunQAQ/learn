#include <stdio.h>
#include <stdlib.h>


int main(void) {

//    bus error
    char *str = "Hello World";
    *str = 'h';

//    no problem
//    char str[12] = "Hello World";
//    *str = 'h';

    printf("%s\n", str);

    return 0;
}
