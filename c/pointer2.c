#include <stdio.h>
#include <string.h>

int main(void)
{
    char string[15] = "ShiXP love C!";
    char c = string[0];
    char *pointer;

    pointer = &string[0];
    printf("指针pointer的值是%p \n", pointer);

    for(int index = 0;index < strlen(string); index++)
    {
        printf("第%d个字符是：%c \n", index + 1, *(pointer + index));
    }

    return 0;
}
