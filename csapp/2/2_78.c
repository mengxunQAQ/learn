// 计算x/2^k
#include <stdio.h>

int divide_power2(int x, int k){
	return x >> k;
}


int main(){
	printf("%d\n", divide_power2(2, 1));
	printf("%d\n", divide_power2(8, 2));
	
	return 0;
}

