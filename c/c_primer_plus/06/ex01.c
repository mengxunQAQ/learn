//打印26个小写字母
#include <stdio.h>
#define SIZE 26

int main(void){
    char lcase[SIZE];
    int i;

    for (i = 0; i < SIZE; i++)
        lcase[i] = 'a' + i;
    
    for (i = 0; i < SIZE; i++)
        printf("%c", lcase[i]);
    printf("\n");
    
    return 0;
}

        
