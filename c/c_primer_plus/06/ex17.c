/*100万美元，年利率8%，每年取10万，多少年取完*/
#include <stdio.h>

int main(void){
    const float RATE = 1.08;
    int total, reduce, year;
    total = 100;
    reduce = 10;
    year = 0;

    for (; total > 0; total = total - reduce)
    {
        total *= RATE;
        year += 1;
    }
    
    printf("%d\n", year);

    return 0;
}

