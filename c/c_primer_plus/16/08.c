// 使用头文件
#include <stdio.h>
#include "06.h"

// 记住连接07.c
int main(void)
{
    names candidate;
    get_names(&candidate);
    printf("Let's welcome ");
    show_names(&candidate);
    printf(" to this program!\n");

    return 0;
}
