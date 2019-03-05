/*转换时间*/
#include <stdio.h>
const int NUM = 60;

int min_input(void){
    int min;
    printf("enter min: ");
    scanf("%d", &min);
    return min;
}

int main(void){
    int hour;
    int min;
    min = min_input();
    while (min > 0){
        hour = min / NUM;
        min = min % NUM;
        printf("%d hour(s) %d min(s)\n", hour, min);
        min = min_input();
    }
    
    return 0;
}
