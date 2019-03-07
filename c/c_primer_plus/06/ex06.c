/*打印表格*/
#include <stdio.h>

int main(void)
{
    int start, end, i;

    printf("enter limit: ");
    scanf("%d %d", &start, &end);

    for (i = start; i <= end; i++)
        printf("%5d %5d %5d\n", i, i*i, i*i*i);

    return 0;
}
