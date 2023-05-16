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

    struct student_st stu = {101102, "Alice", {1999, 2, 10}};
    printf("%d %s %d-%d-%d\n", stu.ID, stu.name, stu.birth.year, stu.birth.month, stu.birth.day);
    printf("%p %p\n", &stu.ID, stu.name);  //series

    struct student_st *p = &stu;
    printf("%d %s %d-%d-%d\n", p->ID, p->name, p->birth.year, p->birth.month, p->birth.day);

    return 0;
}
