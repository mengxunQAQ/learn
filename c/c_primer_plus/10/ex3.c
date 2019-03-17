/* 返回数组里的最大值 */
#include <stdio.h>

int max(int *ar, int);

int main(void)
{
    int ar[4] = {3, 1, 2, 4};
    int value = max(ar, 4);

    printf("the max value is %d\n", value);

    return 0;
}

int max(int *ar, int n)
{
    int i;
    int value;

    for (i = 1, value = ar[0]; i < n; i++)
        value = ar[i] > value ? ar[i] : value;
    
    return value;
}
