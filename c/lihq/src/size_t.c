#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int main(void) {
    char str[] = "hello";
    char *dest_str;
    dest_str = (char *)malloc(sizeof(str));
    memcpy(dest_str, str, sizeof(str));

    printf("dest str is %s\n", dest_str);  

    size_t i = 10;
    unsigned long j = 10;
    printf("size_t number is %ld\n", i);
    printf("unsigned number is %ld\n", j);

    char str2[] = "world";
    char *dest_str2;

    dest_str2 = str2;
    printf("str2 address is %p\n", str2);
    printf("dest_str2 address is %p\n", dest_str2);

    //memcpy(dest_str2, str, sizeof(str));
    memcpy(dest_str2, str, sizeof(str)-2);
    printf("str2 and dest_str2 is %s %s\n", str2, dest_str2);

}
