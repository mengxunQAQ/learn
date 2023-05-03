#include <stdio.h>
#include <stdlib.h>


int main() {
    int i;
    char *name[5] = {"follo\0w me", "Basic", "Fortran", "Computer", "Apple"};

    for (i = 0 ; i < 5 ; i++) {
        puts(name[i]);
    }

    return 0;

}

