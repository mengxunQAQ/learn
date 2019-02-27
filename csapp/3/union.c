/*利用union转换类型，拥有一样的位表示*/
#include <stdio.h>

unsigned long double2bits(double d){
	union {
		double d;
		unsigned long u;
	} temp;
	temp.d = d;
	return temp.u;
}

int main(){
	double d = 1.1;
	unsigned long u = double2bits(d);
	printf("%#lx %#lx", *(unsigned long *)&d, u);
}
