#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int a;
    int b[2];
} Ex2;

typedef struct EX {
    int a;
    char b[3];
    Ex2 c;
    struct EX *d;
} Ex;


int main(void) {
    Ex x = {1, {'H', 'i', 'x'}, {2, {3, 4}}, NULL};
    Ex y = {2, {'H', 'i', 'y'}, {2, {3, 4}}, NULL};
    Ex *px = &x;
    x.d = &y;

//    int *pi = &px->a;
//    printf("%d\n", px->d->a);
//    printf("%d\n", px->d->c.b[0]);
    Ex2 *pc = &px->d->c;
    printf("%d\n", pc->b[0]);
    return 0;
}
