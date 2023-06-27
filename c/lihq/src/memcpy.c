#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    char str1[3] = "ab";
    char str2[3];
    char *str3;

    memcpy(str2, str1, sizeof(str1)); 
    str3 = str1;

    
    printf("%s %s %s\n", str1, str2, str3);
    printf("%p %p %p\n", str1, str2, str3);
    memcpy(str1, "cd", 3); 
    printf("%s %s %s\n", str1, str2, str3);


    char s1[7] = "abcdef";
    char s2[3] = "de";
    char *s3;

    s3 = "a";
    //memcpy(s3, s1, 5);
    printf("%s\n", s3); 
    int *i;
    i = 1;

}
