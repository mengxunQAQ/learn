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

typedef struct test_s {

       int i;
       char s;


} test_t;

typedef struct bitfield {
    unsigned int a:2;
    unsigned int b:2;
    unsigned int c:2;
} my_bitfield;

int main() {

    struct student_st stu = {101102, "Alice", {1999, 2, 10}};
    printf("%d %s %d-%d-%d\n", stu.ID, stu.name, stu.birth.year, stu.birth.month, stu.birth.day);
    printf("%p %p\n", &stu.ID, stu.name);  //series

    struct student_st *p = &stu;
    printf("%d %s %d-%d-%d\n", p->ID, p->name, p->birth.year, p->birth.month, p->birth.day);

    //test_t t[] = {{1, 'a'}, {2, 'b'}};
    //printf("%lu\n", sizeof(t));

    // not give all value
    test_t my_test_t = {
        .s = 'a',
    };
    printf("my_test_t %d %c\n", my_test_t.i, my_test_t.s);
    printf("my_test_t size %lu \n", sizeof(my_test_t));

    // bitfield
    my_bitfield test_1 = {1,2,3};
    printf("my bitfield size: %lu\n", sizeof(test_1));


    return 0;
}
