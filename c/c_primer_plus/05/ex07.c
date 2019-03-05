/*计算立方*/
#include <stdio.h>

void cube(double n);

int main(void){
    double n;

    printf("enter double: ");
    scanf("%lf", &n);
    cube(n);

    return 0;
}

void cube(double n){
    printf("the cube of %e is %e.\n", n, n*n*n);
}
