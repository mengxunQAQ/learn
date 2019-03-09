// ei出现的次数
#include <stdio.h>
#include <stdbool.h>

int main(void)
{
    char ch;
    bool flag = false;
    int count = 0;

    printf("Enter:\n");

    while ((ch = getchar()) != '#')
    {
        if (ch == 'e')
        {
            flag = true;
            continue;
        }
        if (ch == 'i' && flag == true)
            count++;
        flag = false;
    }

    printf("count: %d\n", count);

    return 0;
}
