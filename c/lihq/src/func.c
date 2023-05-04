#include <stdio.h>
#include <stdlib.h>


int add(int, int);

int main() {
     int (*p) (int, int) = add;
     printf("%d %d\n", add(1, 2), p(1, 2));
}


int add(int x, int y) {
     return x+y;
}
