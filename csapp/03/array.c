#include <stdio.h>

int main(){
	char A[4] = {1, 2, 3, 4};
	printf("*A = %x\n", *A);
	printf("A = %x\n", A);
	printf("&A[0] = %x\n", &A[0]);
	printf("A[0] = %x\n", A[0]);
}
