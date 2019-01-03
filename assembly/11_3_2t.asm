assume cs:code

code segment
    start: mov ax,0f000h
           mov ds,ax

	   mov bx,0
	   mov dx,0
	   mov cx,32
	s: mov al,[bx]
	   cmp al,32
	   jna s0
	   cmp al,128
	   jnb s0
	   inc dx
       s0: inc bx
           loop s
code ends

end start
