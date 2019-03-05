/*几种常见的语句*/
#include <stdio.h>

int main(void)
{
    int count, sum;  // 声明,不是表达式
    count = 0;
    sum = 0;
    while (count++ < 20)
    sum = sum + count;
    printf("sum = %d\n", sum);
    return 0;
}
