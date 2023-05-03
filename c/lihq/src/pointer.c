#include <stdio.h>
#include <stdlib.h>


int main() {
//    static int *p;
//    int *p = NULL;
    int *p = NULL;
    *p = 1;
    printf("%d\n", p);  // pointer's value is address
    printf("%p\n", *p);

    return 0;
}
