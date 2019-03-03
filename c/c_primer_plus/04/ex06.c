//编程练习第四章第6题
#include <stdio.h>
#include <string.h>

int main(void)
{
    char name1[20], name2[20];
    printf("enter your name:\n");
    scanf("%s %s", name1, name2);
    printf("%s %s\n", name1, name2);
    printf("%*lu %*lu", (int) strlen(name1), strlen(name1), (int) strlen(name2), strlen(name2));
    return 0;
}
