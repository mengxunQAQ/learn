#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int main(void) {
    char msg[] = "Hello";  // end with \0
    // print 6 5
    printf("%d %d\n", sizeof(msg),  strlen(msg));

    // print 5 5
    char msg2[] = {'H', 'e', 'l', 'l', 'o'};
    printf("%d %d\n", sizeof(msg2),  strlen(msg2));


    return 0;
}
