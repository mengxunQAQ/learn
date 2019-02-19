// 判断位模式里1的个数是否是寄数,
// 证明过程：http://rex-shen.net/%E8%AE%A1%E7%AE%97unsigned%E5%8F%98%E9%87%8F%E4%BA%8C%E8%BF%9B%E5%88%B6%E8%A1%A8%E7%A4%BA%E4%B8%AD%E5%90%AB%E6%9C%89%E7%9A%841%E7%9A%84%E4%B8%AA%E6%95%B0%E7%9A%84%E5%A5%87%E5%81%B6%E6%80%A7/
#include <stdio.h>

int even_ones(unsigned x)
{
	unsigned  y = x >> 16; x ^= y;
	printf("%x\n", x);
	y = x >> 8; x ^= y;
	printf("%x\n", x);
	y = x >> 4; x ^= y;
	printf("%x\n", x);
	y = x >> 2; x ^= y; 
	printf("%x\n", x);
	y = x >> 1; x ^= y; 
        printf("%x\n", x);

	return x & 1;
	
}

int main(){
    printf("%d\n", even_ones(0xf1));
    return 0;
}
