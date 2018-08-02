/*
求第 m 个到第 n 个素数的和 

*/ 

# include <stdio.h>

int main(){
	
	int m,n;
	int num,num2;
	int number = 0;
	int sum=0;
	
	scanf("%d %d", &n, &m);
	
	for (num=2; number<m; num++){
		for (num2=2; num2<num; num2++){
			if (num%num2 == 0) break;
		}
		if (num2 == num){
			number++;
			if (n<=number){
				sum += num;
			}
		} 
		
	};
	
	printf("%d", sum);
	return 0;
}
