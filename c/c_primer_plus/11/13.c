/* 不要模仿 */
#include <stdio.h>

int main(void)
{
    char side_a[] = "Side A";
    char dont[] = {'W', 'O', 'W', '!'};
    char side_b[] = "Side B";

    puts(dont); // dont不是字符串，只是一个字符数组（没有\0空白符）

    return 0;
}
