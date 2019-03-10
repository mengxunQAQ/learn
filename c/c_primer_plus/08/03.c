//打开文件并显示文件
#include <stdio.h>
#include <stdlib.h>

int main()
{
    int ch;
    FILE *fp;
    char fname[50];  // 储存文件名
    
    printf("Enter the name of the file: ");
    scanf("%s", fname);

    fp = fopen(fname, "r");  // 打开待读取文件
    
    if (fp == NULL)  // 如果失败
    {
        printf("Failed to open file. Bye\n");
        exit(1);
    }
    
    while ((ch = getc(fp)) != EOF)
        putchar(ch);
    
    fclose(fp);  // 关闭文件
    return 0;
}



