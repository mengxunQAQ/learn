// 判断使用的机器是不是算术右移
#include <stdio.h>

int int_shifts_are_arithmetic();

int main(){
	int result = int_shifts_are_arithmetic();
	printf("%d", result);
}

int int_shifts_are_arithmetic(){
	int x = -1;
	return x << 1 == -1;
}



