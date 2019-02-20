// 计算3/4x的值
#include <stdio.h>
#include <limits.h>

int threefourths(int x){

	int f = x & ~0x3; // 巧妙的拆为两部分，否则后两位会被舍去 
	int l = x & 0x3;
	int t1  = f >> 2;
	int t2 = (t1 << 1) + t1;

	int p1 = (l << 1) + l;
	int p2 = p1 >> 2;

	return t2+p2;

}


int main(){
	printf("%d\n", threefourths(8));
	printf("%d\n", threefourths(7));
	return 0;
}

