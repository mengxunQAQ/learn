// 2.93 求浮点数的绝对值，如果等于NaN，则直接返回浮点数本身
#include <stdio.h>

typedef unsigned float_bits;

float float_absval(float_bits f){
	//printf("%f\n", *(float *)&f);

	int exp = f >> 23 & 0xff;  // 对整数的位进行移位操作就相当于对浮点数的位进行移位操作
	int frac = f & 0x7fffff;
	if (exp==0xff && frac != 0){
		return *(float *)&f;  // 把整数f转换为float，转换前后二进制不变
	}
	int f2 = ~(1 << 31) & f;
	//printf("%d %d\n",f, f2);
	return *(float *)&f2;
}

int main(){
	float x = -32.1;
	//printf("%#x\n", *(int *)&x);
	float f = float_absval(*(int *)&x);  // 把浮点数x类型转换成整数，转换前后二进制不变，这样操作整数就相当于操作浮点数的底层二进制
	printf("%f", f);
	return 0;
}
