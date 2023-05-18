#include <stdio.h>
#include <stdlib.h>


int main() {
//    static int *p;
//    int *p = NULL;
//    int *p = NULL;
//    *p = 1;
//    printf("%d\n", p);  // pointer's value is address
//    printf("%p\n", *p);
    int a;
    extern int b;   // extern not exist in memory
    int c;
    printf("%p\n", &a);
    printf("%p\n", &c);

    return 0;
}
