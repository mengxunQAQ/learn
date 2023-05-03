/*
 * practice how to use <define>
 * */
#include <stdio.h>
#include <stdlib.h>

#define PI 3.14159
//#define MAX(a, b) (a > b ? a : b)
#define MAX(a, b) \
    ({typeof(a) A=a,B=b; (A > B ? A : B);})


#if 0
int max(int a, int b) {
    return a > b ? a : b;
}
#endif

int main(void) {
    //PI = 5;
    int a = 3, b = 5;
    printf("%d\n", MAX(a++, b++)); 
    printf("%d\t%d\n", a, b);

    return 0;  
}
