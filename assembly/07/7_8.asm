assume cs:codesg,ds:datasg,ss:stacksg

datasg segment
    db 'ibm             '
    db 'dec             '
    db 'dos             '
    db 'vax             '
datasg ends

stacksg segment
    dw 0,0,0,0,0,0,0,0
stacksg ends

codesg segment
    start: mov ax,stacksg
           mov ss,ax
	   mov sp,16
           mov ax,datasg
           mov ds,ax
	   mov cx,4
	   mov si,3
	   mov bx,0

       s:  push cx
           mov cx,3
           mov si,0
       s0: mov al,[bx+si]
	   and al,11011111b
	   mov [bx+si],al
	   inc si
           loop s0
	   add bx,16
	   pop cx
	   loop s

	   mov ax,4c00h
	   int 21h
           
codesg ends

end start
