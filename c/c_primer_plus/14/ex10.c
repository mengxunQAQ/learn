/* 函数指针菜单 */
#include <stdio.h>
#define SIZE 4

typedef void (*FP)(void); // FP为指向函数的指针

void chips(void);
void coke(void);
void rice(void);
void water(void);
void excute(FP *, char *);

int main(void)
{
    FP choose[SIZE] = {chips, coke, rice, water}; 
    char option[SIZE] = {'a', 'b', 'c', 'd'};

    printf("Please choose one option(enter a/b/c/d):\n");
    printf("a. chips    b. coke\n");
    printf("c. rice     d. water\n");
    
    excute(choose, option);
    
    return 0;
}

void excute(FP *choose, char *option){
    int count;
    char input = getchar();
    
    for (count = 0; count < SIZE; count++){
        if (input == option[count])
        {    choose[count]();  // 执行函数
             break;
        }
    }

    if (count == SIZE)
        printf("nothing is selected!\n");
}

void chips(void){
    printf("chips is selected!\n");
}

void coke(void){
    printf("coke is selected!\n");
}

void rice(void){
    printf("rice is selected!\n");
}

void water(void){
    printf("water is selected!\n");
}

    
