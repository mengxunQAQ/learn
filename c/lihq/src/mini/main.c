#include <stdlib.h>
#include <stdio.h>
#include "max.h"

/*
 *  main函数里的void表示无参数 不可以传参数，如果main()没有void 则表示不知道参数有多少
 *  extern int gAll 是声明   int gAll 是定义；声明不算代码不占用空间
 * */
int main(void) {

    int i = 1;
    int max_i;
    max_i = max(i, gAll);
    printf("max int is %d\n", max_i);

}
