#include <stdio.h>
#include <stdlib.h>


int main(void) {
    printf("1\n");
    goto out;
    printf("2\n");
out:
    printf("3\n");
out2:
    printf("4\n");
}
