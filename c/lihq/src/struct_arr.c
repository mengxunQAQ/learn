#include <stdio.h>
#include <stdlib.h>

#define NAMESIZE 32


struct birthday_st {
    int year;
    int month;
    int day;
};

struct student_st {
    int ID;
    char name[NAMESIZE];
    struct birthday_st birth;
};

int main() {

    struct student_st stu[2] = { {101102, "Alice", {1999, 2, 10}}, {101103, "Bob", {1998, 11, 5}} };
    struct student_st *p = &stu[0];
    printf("%p %p\n", p, p+1);

    return 0;
}
