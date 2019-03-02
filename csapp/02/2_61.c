//判断x的最高有效字节等于0
#include <stdio.h>

int judge(int num);

int main(){
	int num = 0x00abcdef;
	printf("%x\n", judge(num));

	return 0;
}

int judge(int num){
	return !(num >> ((sizeof(int)-1) << 3));
}
