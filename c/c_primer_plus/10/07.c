/* 计算每年的总降水量、年平均降水量和五月中每月的平均降水量 */
#include <stdio.h>
#define MONTHS 12  // 一年的月份数
#define YEARS 5  // 年数

int main(void)
{
    // 用2010～2014的降水量数据初始化数组
    const float rain[YEARS][MONTHS] = 
        {
            {4.3, 4.3, 4.3, 3.0, 2.0, 1.2, 0.2, 0.2, 0.4, 2.4, 3.5, 6.6},
            {8.5, 8.2, 1.2, 2.4, 0.0, 5.2, 0.9, 0.3, 0.1, 2.9, 6.5, 3.6},
            {3.3, 2.3, 1.5, 3.2, 1.1, 3.1, 0.3, 0.1, 1.4, 3.4, 3.3, 5.2},
            {5.3, 1.3, 5.1, 3.9, 3.1, 1.1, 5.2, 0.2, 0.9, 4.2, 3.4, 4.6},
            {7.7, 4.9, 3.6, 1.9, 2.1, 3.2, 0.1, 0.5, 0.8, 2.9, 3.5, 5.4}
        };
    int year, month;
    float subtot, total;
    
    printf(" YEAR RAINFALL (inches)\n");

    for (year = 0, total = 0; year < YEARS; year++)
    {
        for (month = 0, subtot = 0; month < MONTHS; month++)
            subtot += rain[year][month];
        printf("%5d %15.1f\n", 2010 + year, subtot);
    
        total += subtot;  // 五年的总降水量
    }
    printf("\nThe yearly average is %.1f inches.\n\n", total / YEARS);
    printf("MONTHLY AVERAGES:\n\n");
    printf("Jan Feb Mar Apr May Jun JUl Aug Sep Oct ");
    printf("Nov Dec\n");

    for (month = 0; month < MONTHS; month++)
    {   // 五年里每个月的总降水量
        for (year = 0, subtot = 0; year < YEARS; year++)
            subtot += rain[year][month];

        printf("%4.1f ", subtot / YEARS);
    }

    printf("\n");

    return 0;
}
