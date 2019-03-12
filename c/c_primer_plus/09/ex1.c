#include <stdio.h>

double min(double x, double y);

int main(void)
{
    double value;
    double x = 1.1;
    double y = 1.2;

    value = min(x, y);

    printf("The min value between double %lf and double %lf "
           "is %lf.\n", x, y, value);

    return 0;
}

double min(double x, double y)
{
    return (x<y ? x : y);
}

    
