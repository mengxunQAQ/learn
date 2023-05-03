#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#if 0
int main() {
    int i,ret;
    do {
        printf("Please enter: ");
        ret = scanf("%d", &i);
        if (ret == 0) {
            printf("not number\n");
            break;
        }
        
        printf("%d\n", i);

    } while (1);

    return 0;

}
#endif

# define STRSIZE 4

int main() {
    char str[STRSIZE];
    int i;
    
    strcpy(str, "abcd");

    for (i=0;i<STRSIZE;i++) {

         printf("%c\n", str[i]);
    }

    return 0;


}
