#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    char str1[10] = "Saul";
    char *str2 = str1;
    char str3[10];

    strcpy(str3, str1);
    str1[0] = 'A';
    printf("%s\n", str2);
    printf("%s\n", str3);


}
