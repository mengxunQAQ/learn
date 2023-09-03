#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
    char str[] = "hello";
    char dest_str[] = "world";

    printf("str address is %p\n", str);
    printf("dest_str address is %p\n", dest_str);

    printf("str is %s\n", str);
    
    memcpy(dest_str, "world!", 6);  // no end flag
    printf("dest_str is %s\n", dest_str);  // will print world!hello

#if 0
    memcpy(dest_str, "world!", 7);  // overflow

    printf("str address is %p\n", str);
    printf("str is %s\n", str);  // print `!`
#endif


}
