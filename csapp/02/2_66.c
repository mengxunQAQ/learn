//只保留最高位的1，比如0xFF00->0x8000,0x6600->0x4000
#include <stdio.h>

int leftmost_one(unsigned x);

int main(){
    unsigned int x = 0xff00;
    printf("%d\n", leftmost_one(x));
    printf("%#x\n", leftmost_one(x));
    
    return 0;
}

int leftmost_one(unsigned x){
	x |= x>>1;
	x |= x>>2;
	x |= x>>4;
	x |= x>>8;
	x |= x>>16;

        return x^(x>>1);
}
