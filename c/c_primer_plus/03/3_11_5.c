/*根据输入的年龄显示对应的秒数*/
#include <stdio.h>

int main(void){
    int age;
    printf("enter ages: ");
    scanf("%d", &age);
    long seconds = age * 3.156e7;
    printf("%ld\n",  seconds);

    return 0;
}
