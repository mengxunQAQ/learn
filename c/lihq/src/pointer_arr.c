#include <stdio.h>
#include <stdlib.h>


int main() {
//    int i;
    char *name[5] = {"follo\0w me", "Basic", "Fortran", "Computer", "Apple"};
    int arr[3] = {1, 2, 3};
    
    // name = char **   name[0] = char *    *a = a[0]
    printf("sizeof name = %ld\n", sizeof(name));
    printf("name addr = %p\n", &name);
    printf("name addr = %p\n", name);
    printf("name[0] addr = %p\n", name[0]);
    printf("name[0] = %c\n", *name[0]);
    printf("name[0] = %s\n", name[0]);
    printf("arr addr = %p\n", arr);
    printf("arr[0] addr = %p\n", &arr[0]);

    //for (i = 0 ; i < 5 ; i++) {
    //    puts(name[i]);
    //    printf("%p\n", name[i]);
    //    printf("%c\n", *name[i]);
    //}

    return 0;

}

