// 位模式下进行旋转，比如x=0x12345678, n=4->0x2345678, n=20->0x67812345
#include <stdio.h>

unsigned rotate_left(unsigned x, int n){
    int t1 = x << n;
    int w = sizeof(int) << 3;
    int t2 = x >> (w-n);
    return t1 | t2;
}

int main(){
    printf("%#x\n", rotate_left(0x12345678, 4));
    printf("%#x\n", rotate_left(0x12345678, 20));
    return 0;
}
