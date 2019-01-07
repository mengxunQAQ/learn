assume cs:code

code segment
    s1: db 'Good,better,best,','$'
    s2: db 'Never let it reset,','$'
    s3: db 'Till good is better,','$'
    s4: db 'And better,best.','$'
    s:  dw offset s1,offset s2,offset s3,offset s4
    row: db 2,4,6,8

    start: mov ax,cs
           mov ds,ax
           mov bx,offset s
           mov si,offset row
           mov cx,4
       ok: mov bh,0 ;第0页
           mov dh,[si] ;第si行
	   mov dl,0 ;第0列
	   mov ah,2 ;调用2号子程序，置光标
	   int 10h

	   mov dx,[bx] ;字符串首地址
	   mov ah,9 ;9号子程序，显示字符串
	   int 21h  ;21号中断
	   add bx,2
	   inc si
       loop ok

	   mov ax,4c00h ;21号中断例程的21h号子程序，程序返回功能
	   int 21h
code ends
end start

