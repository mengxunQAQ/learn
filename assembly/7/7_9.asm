assume cs:codesg,ds:stacksg,ss:datasg

datasg segment
    dw 0,0,0,0,0,0,0,0
datasg ends

stacksg segment
    db '1. display      '
    db '2. brows        '
    db '3. replace      '
    db '4. modify       '
stacksg ends

codesg segment
    start: mov ax,stacksg
           mov ds,ax
	   mov sp,16
	   mov ax,datasg
	   mov ss,ax
	   mov cx,4
	   mov bx,0
        s: push cx
	   mov cx,3
	   mov si,3
       s0: mov al,[bx+si]
           and al,11011111b
	   mov [bx+si],al
	   inc si
	   loop s0
           pop cx
	   add bx,16
	   loop s

codesg ends

end start
