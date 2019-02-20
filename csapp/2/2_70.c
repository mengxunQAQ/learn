// 判断x是否可以被表示为n位
#include <stdio.h>

int fits_bits(int x, int n){
	int w = sizeof(int) << 3;
	int x1 = x << (w-n) >> (w-n);
	return !(x ^ x1);
}

int main(){
	printf("%d\n", fits_bits(0x1234, 8));
	printf("%d\n", fits_bits(0x1234, 16));
	return 0;
}

