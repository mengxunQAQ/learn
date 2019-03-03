//身高
#include <stdio.h>

int main(void)
{
    int height;
    char name[20];
    printf("enter your height:\n");
    scanf("%d", &height);
    printf("enter your name:\n");
    scanf("%s", name);
    printf("%s, you are %.2f meter tall\n", name, height/100.0);
    return 0;
}
